package controller

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/go-logr/logr"
	"github.com/google/uuid"

	"github.com/artefactual/archivematica/hack/ccp/internal/store"
	"github.com/artefactual/archivematica/hack/ccp/internal/workflow"
)

// A Package can be a Transfer, a SIP, or a DIP.
type Package struct {
	logger    logr.Logger
	store     store.Store
	id        uuid.UUID
	path      string
	base      string
	name      string
	isDir     bool
	watchedAt *workflow.WatchedDirectory
	decision  decision
	unit
}

func NewPackage(ctx context.Context, logger logr.Logger, store store.Store, path string, wd *workflow.WatchedDirectory) (*Package, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("stat: %v", err)
	}

	base, name := filepath.Split(path)

	p := &Package{
		logger:    logger,
		store:     store,
		path:      path,
		base:      base,
		name:      name,
		isDir:     fi.IsDir(),
		watchedAt: wd,
	}

	switch {
	case wd.UnitType == "SIP" && p.isDir:
		p.unit = &SIP{p}
	case wd.UnitType == "DIP" && p.isDir:
		p.unit = &DIP{p}
	case wd.UnitType == "Transfer":
		p.unit = &Transfer{p}
	default:
		return nil, fmt.Errorf("unexpected type given for file %q", path)
	}

	if err := p.hydrate(ctx); err != nil {
		return nil, fmt.Errorf("hydrate: %v", err)
	}

	return p, nil
}

func (p *Package) String() string {
	return p.name
}

func (p *Package) PreconfiguredChoice(linkID uuid.UUID) (*workflow.Choice, error) {
	li := linkID.String()

	// TODO: automate "Approve standard transfer" until we can submit decisions.
	if li == "0c94e6b5-4714-4bec-82c8-e187e0c04d77" {
		return &workflow.Choice{
			AppliesTo: "0c94e6b5-4714-4bec-82c8-e187e0c04d77",
			GoToChain: "b4567e89-9fea-4256-99f5-a88987026488",
		}, nil
	}

	f, err := os.Open(filepath.Join(p.path, "processingMCP.xml"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	// TODO: this could be cached if the file isn't going to change.
	choices, err := workflow.ParseConfig(f)
	if err != nil {
		return nil, err
	}

	var match *workflow.Choice
	for _, choice := range choices {
		if choice.AppliesTo == li {
			match = &choice
			break
		}
	}

	// Resort to automated config.
	// TODO: allow user to choose the system processing config to use.
	if match == nil {
		for _, choice := range workflow.AutomatedConfig.Choices.Choices {
			if choice.AppliesTo == li {
				match = &choice
				break
			}
		}
	}

	return match, nil
}

// Decide resolves an awaiting decision.
func (p *Package) Decide(opt option) error {
	return p.decision.resolve(opt)
}

// AwaitDecision builds a new decision and waits for its resolution.
func (p *Package) AwaitDecision(ctx context.Context, opts []option) (option, error) {
	p.decision.build(opts...)

	for {
		select {
		case d := <-p.decision.recv:
			return d, nil
		case <-ctx.Done():
			return option(""), ctx.Err()
		}
	}
}

// Decision provides the current awaiting decision.
func (p *Package) Decision() []option {
	return p.decision.decision()
}

type replacementMapping map[string]string

type unit interface {
	hydrate(ctx context.Context) error
	reload(ctx context.Context) error
	replacements() replacementMapping

	// Props...
	// REPLACEMENT_PATH_STRING = r"%SIPDirectory%"
	// UNIT_VARIABLE_TYPE = "DIP"
	// JOB_UNIT_TYPE = "unitDIP"
}

type Transfer struct {
	p *Package
}

func (u *Transfer) hydrate(ctx context.Context) error { return nil }
func (u *Transfer) reload(ctx context.Context) error  { return nil }
func (u *Transfer) replacements() replacementMapping  { return nil }

type SIP struct {
	p *Package
}

func (u *SIP) hydrate(ctx context.Context) error { return nil }
func (u *SIP) reload(ctx context.Context) error  { return nil }
func (u *SIP) replacements() replacementMapping  { return nil }

type DIP struct {
	p *Package
}

func (u *DIP) hydrate(ctx context.Context) error {
	return nil
}

func (u *DIP) reload(ctx context.Context) error {
	return nil // No-op.
}

func (u *DIP) replacements() replacementMapping {
	return nil
}

type decision struct {
	opts     []option
	recv     chan option
	unsolved atomic.Bool
	sync.Mutex
}

func (pd *decision) build(opts ...option) {
	pd.Lock()
	pd.opts = opts
	pd.recv = make(chan option) // is this ok?
	pd.Unlock()

	pd.unsolved.Store(true)
}

func (pd *decision) resolve(opt option) error {
	if !pd.unsolved.Load() {
		return errors.New("decision is not pending resolution")
	}

	select {
	case pd.recv <- opt:
		pd.unsolved.Store(false)
	default:
		return errors.New("resolve can't proceed because nobody is listening")
	}

	return nil
}

func (pd *decision) decision() []option {
	if !pd.unsolved.Load() {
		return nil
	}

	var opts []option
	if pd.unsolved.Load() {
		pd.Lock()
		opts = make([]option, len(pd.opts))
		copy(opts, pd.opts)
		pd.Unlock()
	}

	return opts
}

// option is a single selectable decision choice.
//
// In most cases, an option is the UUID of a workflow item, but there is one
// exception: "Store DIP location", containing a location path.
type option string

func (do option) uuid() uuid.UUID {
	id, err := uuid.Parse(string(do))
	if err != nil {
		return uuid.Nil
	}
	return id
}
