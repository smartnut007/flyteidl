// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/plugins/sagemaker/hpojob.proto

package plugins

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _hpojob_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on HPOJobObjective with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HPOJobObjective) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ObjectiveType

	// no validation rules for MetricName

	return nil
}

// HPOJobObjectiveValidationError is the validation error returned by
// HPOJobObjective.Validate if the designated constraints aren't met.
type HPOJobObjectiveValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HPOJobObjectiveValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HPOJobObjectiveValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HPOJobObjectiveValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HPOJobObjectiveValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HPOJobObjectiveValidationError) ErrorName() string { return "HPOJobObjectiveValidationError" }

// Error satisfies the builtin error interface
func (e HPOJobObjectiveValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHPOJobObjective.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HPOJobObjectiveValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HPOJobObjectiveValidationError{}

// Validate checks the field values on HPOJob with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HPOJob) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MaxNumberOfTrainingJobs

	// no validation rules for MaxParallelTrainingJobs

	if v, ok := interface{}(m.GetTrainingJob()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HPOJobValidationError{
				field:  "TrainingJob",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HPOJobValidationError is the validation error returned by HPOJob.Validate if
// the designated constraints aren't met.
type HPOJobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HPOJobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HPOJobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HPOJobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HPOJobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HPOJobValidationError) ErrorName() string { return "HPOJobValidationError" }

// Error satisfies the builtin error interface
func (e HPOJobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHPOJob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HPOJobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HPOJobValidationError{}
