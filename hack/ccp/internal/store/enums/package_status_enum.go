// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package enums

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

const (
	// PackageStatusUnknown is a PackageStatus of type Unknown.
	PackageStatusUnknown PackageStatus = iota
	// PackageStatusProcessing is a PackageStatus of type Processing.
	PackageStatusProcessing
	// PackageStatusDone is a PackageStatus of type Done.
	PackageStatusDone
	// PackageStatusCompletedSuccessfully is a PackageStatus of type CompletedSuccessfully.
	PackageStatusCompletedSuccessfully
	// PackageStatusFailed is a PackageStatus of type Failed.
	PackageStatusFailed
)

var ErrInvalidPackageStatus = fmt.Errorf("not a valid PackageStatus, try [%s]", strings.Join(_PackageStatusNames, ", "))

const _PackageStatusName = "UnknownProcessingDoneCompletedSuccessfullyFailed"

var _PackageStatusNames = []string{
	_PackageStatusName[0:7],
	_PackageStatusName[7:17],
	_PackageStatusName[17:21],
	_PackageStatusName[21:42],
	_PackageStatusName[42:48],
}

// PackageStatusNames returns a list of possible string values of PackageStatus.
func PackageStatusNames() []string {
	tmp := make([]string, len(_PackageStatusNames))
	copy(tmp, _PackageStatusNames)
	return tmp
}

var _PackageStatusMap = map[PackageStatus]string{
	PackageStatusUnknown:               _PackageStatusName[0:7],
	PackageStatusProcessing:            _PackageStatusName[7:17],
	PackageStatusDone:                  _PackageStatusName[17:21],
	PackageStatusCompletedSuccessfully: _PackageStatusName[21:42],
	PackageStatusFailed:                _PackageStatusName[42:48],
}

// String implements the Stringer interface.
func (x PackageStatus) String() string {
	if str, ok := _PackageStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PackageStatus(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PackageStatus) IsValid() bool {
	_, ok := _PackageStatusMap[x]
	return ok
}

var _PackageStatusValue = map[string]PackageStatus{
	_PackageStatusName[0:7]:                    PackageStatusUnknown,
	strings.ToLower(_PackageStatusName[0:7]):   PackageStatusUnknown,
	_PackageStatusName[7:17]:                   PackageStatusProcessing,
	strings.ToLower(_PackageStatusName[7:17]):  PackageStatusProcessing,
	_PackageStatusName[17:21]:                  PackageStatusDone,
	strings.ToLower(_PackageStatusName[17:21]): PackageStatusDone,
	_PackageStatusName[21:42]:                  PackageStatusCompletedSuccessfully,
	strings.ToLower(_PackageStatusName[21:42]): PackageStatusCompletedSuccessfully,
	_PackageStatusName[42:48]:                  PackageStatusFailed,
	strings.ToLower(_PackageStatusName[42:48]): PackageStatusFailed,
}

// ParsePackageStatus attempts to convert a string to a PackageStatus.
func ParsePackageStatus(name string) (PackageStatus, error) {
	if x, ok := _PackageStatusValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _PackageStatusValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return PackageStatus(0), fmt.Errorf("%s is %w", name, ErrInvalidPackageStatus)
}

func (x PackageStatus) Ptr() *PackageStatus {
	return &x
}

// MarshalText implements the text marshaller method.
func (x PackageStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PackageStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePackageStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errPackageStatusNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *PackageStatus) Scan(value interface{}) (err error) {
	if value == nil {
		*x = PackageStatus(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = PackageStatus(v)
	case string:
		*x, err = ParsePackageStatus(v)
	case []byte:
		*x, err = ParsePackageStatus(string(v))
	case PackageStatus:
		*x = v
	case int:
		*x = PackageStatus(v)
	case *PackageStatus:
		if v == nil {
			return errPackageStatusNilPtr
		}
		*x = *v
	case uint:
		*x = PackageStatus(v)
	case uint64:
		*x = PackageStatus(v)
	case *int:
		if v == nil {
			return errPackageStatusNilPtr
		}
		*x = PackageStatus(*v)
	case *int64:
		if v == nil {
			return errPackageStatusNilPtr
		}
		*x = PackageStatus(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = PackageStatus(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errPackageStatusNilPtr
		}
		*x = PackageStatus(*v)
	case *uint:
		if v == nil {
			return errPackageStatusNilPtr
		}
		*x = PackageStatus(*v)
	case *uint64:
		if v == nil {
			return errPackageStatusNilPtr
		}
		*x = PackageStatus(*v)
	case *string:
		if v == nil {
			return errPackageStatusNilPtr
		}
		*x, err = ParsePackageStatus(*v)
	}

	return
}

// Value implements the driver Valuer interface.
func (x PackageStatus) Value() (driver.Value, error) {
	return x.String(), nil
}

// Set implements the Golang flag.Value interface func.
func (x *PackageStatus) Set(val string) error {
	v, err := ParsePackageStatus(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *PackageStatus) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *PackageStatus) Type() string {
	return "PackageStatus"
}
