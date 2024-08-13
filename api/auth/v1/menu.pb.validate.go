// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/v1/menu.proto

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

// Validate checks the field values on CreateMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateMenuResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateMenuResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateMenuResponseMultiError, or nil if none found.
func (m *CreateMenuResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateMenuResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	if len(errors) > 0 {
		return CreateMenuResponseMultiError(errors)
	}

	return nil
}

// CreateMenuResponseMultiError is an error wrapping multiple validation errors
// returned by CreateMenuResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateMenuResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateMenuResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateMenuResponseMultiError) AllErrors() []error { return m }

// CreateMenuResponseValidationError is the validation error returned by
// CreateMenuResponse.Validate if the designated constraints aren't met.
type CreateMenuResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMenuResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMenuResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMenuResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMenuResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMenuResponseValidationError) ErrorName() string {
	return "CreateMenuResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMenuResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMenuResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMenuResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMenuResponseValidationError{}

// Validate checks the field values on ListMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMenuRequestMultiError, or nil if none found.
func (m *ListMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for ParentId

	// no validation rules for Sort

	// no validation rules for Level

	if len(errors) > 0 {
		return ListMenuRequestMultiError(errors)
	}

	return nil
}

// ListMenuRequestMultiError is an error wrapping multiple validation errors
// returned by ListMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type ListMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMenuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMenuRequestMultiError) AllErrors() []error { return m }

// ListMenuRequestValidationError is the validation error returned by
// ListMenuRequest.Validate if the designated constraints aren't met.
type ListMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMenuRequestValidationError) ErrorName() string { return "ListMenuRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMenuRequestValidationError{}

// Validate checks the field values on MenuDetail with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MenuDetail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MenuDetail with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MenuDetailMultiError, or
// nil if none found.
func (m *MenuDetail) ValidateAll() error {
	return m.validate(true)
}

func (m *MenuDetail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for ParentId

	// no validation rules for Sort

	// no validation rules for Level

	// no validation rules for Path

	// no validation rules for Component

	// no validation rules for Description

	// no validation rules for Leaf

	if len(errors) > 0 {
		return MenuDetailMultiError(errors)
	}

	return nil
}

// MenuDetailMultiError is an error wrapping multiple validation errors
// returned by MenuDetail.ValidateAll() if the designated constraints aren't met.
type MenuDetailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MenuDetailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MenuDetailMultiError) AllErrors() []error { return m }

// MenuDetailValidationError is the validation error returned by
// MenuDetail.Validate if the designated constraints aren't met.
type MenuDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MenuDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MenuDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MenuDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MenuDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MenuDetailValidationError) ErrorName() string { return "MenuDetailValidationError" }

// Error satisfies the builtin error interface
func (e MenuDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMenuDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MenuDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MenuDetailValidationError{}
