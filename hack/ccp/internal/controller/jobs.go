package controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/sevein/gearmin"

	"github.com/artefactual/archivematica/hack/ccp/internal/store/sqlcmysql"
	"github.com/artefactual/archivematica/hack/ccp/internal/workflow"
)

type job struct {
	id        uuid.UUID
	createdAt time.Time
	logger    logr.Logger
	chain     *chain
	p         *Package
	wl        *workflow.Link
	gearman   *gearmin.Server
	jobRunner
}

type jobRunner interface {
	exec(context.Context) (uuid.UUID, error)
}

func newJob(logger logr.Logger, chain *chain, p *Package, gearman *gearmin.Server, wl *workflow.Link) (*job, error) {
	j := &job{
		id:        uuid.New(),
		createdAt: time.Now().UTC(),
		chain:     chain,
		p:         p,
		wl:        wl,
		gearman:   gearman,
	}

	var err error
	switch wl.Manager {

	// Decision jobs - handles workflow decision points.
	case "linkTaskManagerGetUserChoiceFromMicroserviceGeneratedList":
		j.logger = logger.WithName("outputDecisionJob")
		j.jobRunner, err = newOutputDecisionJob(j)
	case "linkTaskManagerChoice":
		j.logger = logger.WithName("nextChainDecisionJob")
		j.jobRunner, err = newNextChainDecisionJob(j)
	case "linkTaskManagerReplacementDicFromChoice":
		j.logger = logger.WithName("updateContextDecisionJob")
		j.jobRunner, err = newUpdateContextDecisionJob(j)

	// Executable jobs - dispatched to the worker pool.
	case "linkTaskManagerDirectories":
		j.logger = logger.WithName("directoryClientScriptJob")
		j.jobRunner, err = newDirectoryClientScriptJob(j)
	case "linkTaskManagerFiles":
		j.logger = logger.WithName("filesClientScriptJob")
		j.jobRunner, err = newFilesClientScriptJob(j)
	case "linkTaskManagerGetMicroserviceGeneratedListInStdOut":
		j.logger = logger.WithName("outputClientScriptJob")
		j.jobRunner, err = newOutputClientScriptJob(j)

	// Local jobs - executed directly.
	case "linkTaskManagerSetUnitVariable":
		j.logger = logger.WithName("setUnitVarLinkJob")
		j.jobRunner, err = newSetUnitVarLinkJob(j)
	case "linkTaskManagerUnitVariableLinkPull":
		j.logger = logger.WithName("getUnitVarLinkJob")
		j.jobRunner, err = newGetUnitVarLinkJob(j)

	default:
		err = fmt.Errorf("unknown job manager: %q", wl.Manager)
	}

	return j, err
}

func (j *job) save(ctx context.Context) error {
	return j.p.store.CreateJob(ctx, &sqlcmysql.CreateJobParams{
		ID:             j.id,
		Type:           j.wl.Description.String(),
		CreatedAt:      j.createdAt,
		Createdtimedec: fmt.Sprintf("%.9f", float64(j.createdAt.Nanosecond())/1e9),
		// Directory: string,        '%sharedPath%currentlyProcessing/FOOBAR2-1fa48af6-8ec0-4ce5-b282-3150fa5f0cbe/',
		SIPID:             j.p.id,
		Unittype:          j.p.jobUnitType(),
		Currentstep:       3,
		Microservicegroup: j.wl.Group.String(),
		Hidden:            false,
		Microservicechainlinkspk: sql.NullString{
			String: j.wl.ID.String(),
			Valid:  true,
		},
	})
}

func (j *job) markAwaitingDecision(ctx context.Context) error {
	err := j.p.store.UpdateJobStatus(ctx, j.id, "STATUS_AWAITING_DECISION")
	if err != nil {
		return fmt.Errorf("mark awaiting decision: %v", err)
	}

	return nil
}

func (j *job) markComplete(ctx context.Context) error {
	err := j.p.store.UpdateJobStatus(ctx, j.id, "STATUS_COMPLETED_SUCCESSFULLY")
	if err != nil {
		return fmt.Errorf("mark complete: %v", err)
	}

	return nil
}

func (j *job) updateStatusFromExitCode(ctx context.Context, code int) error {
	status := ""
	if ec, ok := j.wl.ExitCodes[code]; ok {
		status = ec.JobStatus
	} else {
		status = j.wl.FallbackJobStatus
	}

	err := j.p.store.UpdateJobStatus(ctx, j.id, status)
	if err != nil {
		return fmt.Errorf("update job status from exit code: %v", err)
	}

	return nil
}

func (j *job) replaceValues(cmd string, replacements replacementMapping) string {
	if cmd == "" {
		return ""
	}

	for k, v := range replacements {
		cmd = strings.ReplaceAll(cmd, k, v.escape())
	}

	return cmd
}

// outputDecisionJob.
//
// Manager: linkTaskManagerGetUserChoiceFromMicroserviceGeneratedList.
// Class: OutputDecisionJob(DecisionJob).
type outputDecisionJob struct {
	j      *job
	config *workflow.LinkStandardTaskConfig
}

var _ jobRunner = (*outputDecisionJob)(nil)

func newOutputDecisionJob(j *job) (*outputDecisionJob, error) {
	config, ok := j.wl.Config.(workflow.LinkStandardTaskConfig)
	if !ok {
		return nil, errors.New("invalid config")
	}

	return &outputDecisionJob{
		j:      j,
		config: &config,
	}, nil
}

func (l *outputDecisionJob) exec(ctx context.Context) (uuid.UUID, error) {
	return uuid.Nil, nil
}

// nextChainDecisionJob.
//
// 1. Reload *Package (pull state from db into memory since it can be changed by client scripts).
// 2. Persist job in database: https://github.com/artefactual/archivematica/blob/fbda1a91d6dff086e7124fa1d7a3c7953d8755bb/src/MCPServer/lib/server/jobs/base.py#L76.
// 3. Load preconfigured choices, and resolve the job if the decision is preconfigured.
// 4. Otherwise, mark as awaiting decision, put on hold. Decision must be made by the user via the API.
// 5. ...
//
// Manager: linkTaskManagerChoice.
// Class: NextChainDecisionJob(DecisionJob).
type nextChainDecisionJob struct {
	j      *job
	config *workflow.LinkMicroServiceChainChoice
}

var _ jobRunner = (*nextChainDecisionJob)(nil)

func newNextChainDecisionJob(j *job) (*nextChainDecisionJob, error) {
	config, ok := j.wl.Config.(workflow.LinkMicroServiceChainChoice)
	if !ok {
		return nil, errors.New("invalid config")
	}

	return &nextChainDecisionJob{
		j:      j,
		config: &config,
	}, nil
}

func (l *nextChainDecisionJob) exec(ctx context.Context) (uuid.UUID, error) {
	// When we have a preconfigured choice.
	choice, err := l.j.p.PreconfiguredChoice(l.j.wl.ID)
	if err != nil {
		return uuid.Nil, err
	}
	if choice != nil {
		if ret, err := uuid.Parse(choice.GoToChain); err != nil {
			l.j.logger.Info("Preconfigured choice is not a valid UUID.", "choice", choice.GoToChain, "err", err)
		} else {
			return ret, nil
		}
	}

	// Build decision point and await resolution.
	opts := make([]option, len(l.config.Choices))
	for i, item := range l.config.Choices {
		opts[i] = option(item.String())
	}
	if decision, err := l.j.p.AwaitDecision(ctx, opts); err != nil {
		return uuid.Nil, fmt.Errorf("await decision: %v", err)
	} else {
		return decision.uuid(), nil
	}
}

// updateContextDecisionJob is a job that updates the chain context based on a user choice.
//
// Manager: linkTaskManagerReplacementDicFromChoice (14 matches).
// Class: UpdateContextDecisionJob(DecisionJob) (decisions.py).
type updateContextDecisionJob struct {
	j      *job
	config *workflow.LinkMicroServiceChoiceReplacementDic
}

var _ jobRunner = (*updateContextDecisionJob)(nil)

// nolint: unused
var updateContextDecisionJobChoiceMapping = map[string]string{
	// Decision point "Assign UUIDs to directories?".
	"8882bad4-561c-4126-89c9-f7f0c083d5d7": "bd899573-694e-4d33-8c9b-df0af802437d",
	"e10a31c3-56df-4986-af7e-2794ddfe8686": "bd899573-694e-4d33-8c9b-df0af802437d",
	"d6f6f5db-4cc2-4652-9283-9ec6a6d181e5": "bd899573-694e-4d33-8c9b-df0af802437d",
	"1563f22f-f5f7-4dfe-a926-6ab50d408832": "bd899573-694e-4d33-8c9b-df0af802437d",
	// Decision "Yes" (for "Assign UUIDs to directories?").
	"7e4cf404-e62d-4dc2-8d81-6141e390f66f": "2dc3f487-e4b0-4e07-a4b3-6216ed24ca14",
	"2732a043-b197-4cbc-81ab-4e2bee9b74d3": "2dc3f487-e4b0-4e07-a4b3-6216ed24ca14",
	"aa793efa-1b62-498c-8f92-cab187a99a2a": "2dc3f487-e4b0-4e07-a4b3-6216ed24ca14",
	"efd98ddb-80a6-4206-80bf-81bf00f84416": "2dc3f487-e4b0-4e07-a4b3-6216ed24ca14",
	// Decision "No" (for "Assign UUIDs to directories?").
	"0053c670-3e61-4a3e-a188-3a2dd1eda426": "891f60d0-1ba8-48d3-b39e-dd0934635d29",
	"8e93e523-86bb-47e1-a03a-4b33e13f8c5e": "891f60d0-1ba8-48d3-b39e-dd0934635d29",
	"6dfbeff8-c6b1-435b-833a-ed764229d413": "891f60d0-1ba8-48d3-b39e-dd0934635d29",
	"dc0ee6b6-ed5f-42a3-bc8f-c9c7ead03ed1": "891f60d0-1ba8-48d3-b39e-dd0934635d29",
}

func newUpdateContextDecisionJob(j *job) (*updateContextDecisionJob, error) {
	config, ok := j.wl.Config.(workflow.LinkMicroServiceChoiceReplacementDic)
	if !ok {
		return nil, errors.New("invalid config")
	}

	return &updateContextDecisionJob{
		j:      j,
		config: &config,
	}, nil
}

func (l *updateContextDecisionJob) exec(ctx context.Context) (uuid.UUID, error) {
	id := l.j.wl.ExitCodes[0].LinkID
	if id == nil || *id == uuid.Nil {
		return uuid.Nil, errors.New("ops")
	}
	return *id, nil
}

// directoryClientScriptJob.
//
// Manager: linkTaskManagerDirectories.
// Class: DirectoryClientScriptJob(DecisionJob).
type directoryClientScriptJob struct {
	j      *job
	config *workflow.LinkStandardTaskConfig
}

var _ jobRunner = (*directoryClientScriptJob)(nil)

func newDirectoryClientScriptJob(j *job) (*directoryClientScriptJob, error) {
	config, ok := j.wl.Config.(workflow.LinkStandardTaskConfig)
	if !ok {
		return nil, errors.New("invalid config")
	}

	return &directoryClientScriptJob{
		j:      j,
		config: &config,
	}, nil
}

func (l *directoryClientScriptJob) exec(ctx context.Context) (uuid.UUID, error) {
	if err := l.j.p.reload(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("reload: %v", err)
	}

	if err := l.j.save(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("save: %v", err)
	}

	replacements := l.j.p.unit.replacements(l.config.FilterSubdir)
	replacements.withContext(l.j.chain.ctx)

	args := l.j.replaceValues(l.config.Arguments, replacements)
	stdout := l.j.replaceValues(l.config.StdoutFile, replacements)
	stderr := l.j.replaceValues(l.config.StderrFile, replacements)

	tt := &tasks{Tasks: map[uuid.UUID]*task{}}
	tt.add(l.j.chain.ctx, args, false, stdout, stderr)

	res, err := submitJob(ctx, l.j.logger, l.j.gearman, l.config.Execute, tt)

	l.j.logger.Info("Job executed.", "results", res, "err", err)
	if err != nil {
		return uuid.Nil, err
	}

	task := res.Results[uuid.MustParse("09fdf5bb-3361-4323-9bd7-234fbc9b6517")]

	if err := l.j.updateStatusFromExitCode(ctx, task.ExitCode); err != nil {
		return uuid.Nil, err
	}

	return *l.j.wl.FallbackLinkID, nil
}

// filesClientScriptJob.
//
// Manager: linkTaskManagerFiles.
// Class: FilesClientScriptJob(DecisionJob).
type filesClientScriptJob struct {
	j      *job
	config *workflow.LinkStandardTaskConfig
}

var _ jobRunner = (*filesClientScriptJob)(nil)

func newFilesClientScriptJob(j *job) (*filesClientScriptJob, error) {
	config, ok := j.wl.Config.(workflow.LinkStandardTaskConfig)
	if !ok {
		return nil, errors.New("invalid config")
	}

	return &filesClientScriptJob{
		j:      j,
		config: &config,
	}, nil
}

func (l *filesClientScriptJob) exec(ctx context.Context) (uuid.UUID, error) {
	return uuid.Nil, nil
}

// outputClientScriptJob.
//
// Manager: linkTaskManagerGetMicroserviceGeneratedListInStdOut.
// Class: OutputClientScriptJob(DecisionJob).
type outputClientScriptJob struct {
	j      *job
	config *workflow.LinkStandardTaskConfig
}

var _ jobRunner = (*outputClientScriptJob)(nil)

func newOutputClientScriptJob(j *job) (*outputClientScriptJob, error) {
	config, ok := j.wl.Config.(workflow.LinkStandardTaskConfig)
	if !ok {
		return nil, errors.New("invalid config")
	}

	return &outputClientScriptJob{
		j:      j,
		config: &config,
	}, nil
}

func (l *outputClientScriptJob) exec(ctx context.Context) (uuid.UUID, error) {
	return uuid.Nil, nil
}

// setUnitVarLinkJob is a local job that sets the unit variable configured in the workflow.
//
// Manager: linkTaskManagerSetUnitVariable.
// Class: SetUnitVarLinkJob(DecisionJob) (decisions.py).
type setUnitVarLinkJob struct {
	j      *job
	config *workflow.LinkTaskConfigSetUnitVariable
}

var _ jobRunner = (*setUnitVarLinkJob)(nil)

func newSetUnitVarLinkJob(j *job) (*setUnitVarLinkJob, error) {
	config, ok := j.wl.Config.(workflow.LinkTaskConfigSetUnitVariable)
	if !ok {
		return nil, errors.New("invalid config")
	}

	return &setUnitVarLinkJob{
		j:      j,
		config: &config,
	}, nil
}

func (l *setUnitVarLinkJob) exec(ctx context.Context) (uuid.UUID, error) {
	return l.config.ChainID, nil
}

// getUnitVarLinkJob is a local job that gets the next link in the chain from a UnitVariable.
//
// Manager: linkTaskManagerUnitVariableLinkPull.
// Class: GetUnitVarLinkJob(DecisionJob) (decisions.py).
type getUnitVarLinkJob struct {
	j      *job
	config *workflow.LinkTaskConfigUnitVariableLinkPull
}

var _ jobRunner = (*getUnitVarLinkJob)(nil)

func newGetUnitVarLinkJob(j *job) (*getUnitVarLinkJob, error) {
	config, ok := j.wl.Config.(workflow.LinkTaskConfigUnitVariableLinkPull)
	if !ok {
		return nil, errors.New("invalid config")
	}

	return &getUnitVarLinkJob{
		j:      j,
		config: &config,
	}, nil
}

func (l *getUnitVarLinkJob) exec(ctx context.Context) (uuid.UUID, error) {
	return l.config.ChainID, nil
}
