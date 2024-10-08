// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/v1/permission.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GrantRolePermissionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GrantRolePermissionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrantRolePermissionRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GrantRolePermissionRequestMultiError, or nil if none found.
func (m *GrantRolePermissionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantRolePermissionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoleId

	if len(errors) > 0 {
		return GrantRolePermissionRequestMultiError(errors)
	}

	return nil
}

// GrantRolePermissionRequestMultiError is an error wrapping multiple
// validation errors returned by GrantRolePermissionRequest.ValidateAll() if
// the designated constraints aren't met.
type GrantRolePermissionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantRolePermissionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantRolePermissionRequestMultiError) AllErrors() []error { return m }

// GrantRolePermissionRequestValidationError is the validation error returned
// by GrantRolePermissionRequest.Validate if the designated constraints aren't met.
type GrantRolePermissionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantRolePermissionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrantRolePermissionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrantRolePermissionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrantRolePermissionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantRolePermissionRequestValidationError) ErrorName() string {
	return "GrantRolePermissionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GrantRolePermissionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantRolePermissionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantRolePermissionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantRolePermissionRequestValidationError{}

// Validate checks the field values on GrantRolePermissionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GrantRolePermissionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrantRolePermissionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GrantRolePermissionResponseMultiError, or nil if none found.
func (m *GrantRolePermissionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GrantRolePermissionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	if len(errors) > 0 {
		return GrantRolePermissionResponseMultiError(errors)
	}

	return nil
}

// GrantRolePermissionResponseMultiError is an error wrapping multiple
// validation errors returned by GrantRolePermissionResponse.ValidateAll() if
// the designated constraints aren't met.
type GrantRolePermissionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrantRolePermissionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrantRolePermissionResponseMultiError) AllErrors() []error { return m }

// GrantRolePermissionResponseValidationError is the validation error returned
// by GrantRolePermissionResponse.Validate if the designated constraints
// aren't met.
type GrantRolePermissionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrantRolePermissionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrantRolePermissionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrantRolePermissionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrantRolePermissionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrantRolePermissionResponseValidationError) ErrorName() string {
	return "GrantRolePermissionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GrantRolePermissionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrantRolePermissionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrantRolePermissionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrantRolePermissionResponseValidationError{}
