// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: receiver.proto

package v3

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

// Validate checks the field values on ReceiverRule with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReceiverRule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiverRule with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReceiverRuleMultiError, or
// nil if none found.
func (m *ReceiverRule) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiverRule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Source

	// no validation rules for Destination

	// no validation rules for Svc

	if len(errors) > 0 {
		return ReceiverRuleMultiError(errors)
	}

	return nil
}

// ReceiverRuleMultiError is an error wrapping multiple validation errors
// returned by ReceiverRule.ValidateAll() if the designated constraints aren't met.
type ReceiverRuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiverRuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiverRuleMultiError) AllErrors() []error { return m }

// ReceiverRuleValidationError is the validation error returned by
// ReceiverRule.Validate if the designated constraints aren't met.
type ReceiverRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiverRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiverRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiverRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiverRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiverRuleValidationError) ErrorName() string { return "ReceiverRuleValidationError" }

// Error satisfies the builtin error interface
func (e ReceiverRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiverRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiverRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiverRuleValidationError{}

// Validate checks the field values on Receiver with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Receiver) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Receiver with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReceiverMultiError, or nil
// if none found.
func (m *Receiver) ValidateAll() error {
	return m.validate(true)
}

func (m *Receiver) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SelfNamespace

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReceiverValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReceiverValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReceiverValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ReceiverMultiError(errors)
	}

	return nil
}

// ReceiverMultiError is an error wrapping multiple validation errors returned
// by Receiver.ValidateAll() if the designated constraints aren't met.
type ReceiverMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiverMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiverMultiError) AllErrors() []error { return m }

// ReceiverValidationError is the validation error returned by
// Receiver.Validate if the designated constraints aren't met.
type ReceiverValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiverValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiverValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiverValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiverValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiverValidationError) ErrorName() string { return "ReceiverValidationError" }

// Error satisfies the builtin error interface
func (e ReceiverValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiver.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiverValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiverValidationError{}
