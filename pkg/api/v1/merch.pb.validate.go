// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: merch.proto

package merch

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

// Validate checks the field values on InfoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InfoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InfoRequestMultiError, or
// nil if none found.
func (m *InfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return InfoRequestMultiError(errors)
	}

	return nil
}

// InfoRequestMultiError is an error wrapping multiple validation errors
// returned by InfoRequest.ValidateAll() if the designated constraints aren't met.
type InfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InfoRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InfoRequestMultiError) AllErrors() []error { return m }

// InfoRequestValidationError is the validation error returned by
// InfoRequest.Validate if the designated constraints aren't met.
type InfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoRequestValidationError) ErrorName() string { return "InfoRequestValidationError" }

// Error satisfies the builtin error interface
func (e InfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoRequestValidationError{}

// Validate checks the field values on InfoResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InfoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InfoResponseMultiError, or
// nil if none found.
func (m *InfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *InfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Coins

	for idx, item := range m.GetInventory() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, InfoResponseValidationError{
						field:  fmt.Sprintf("Inventory[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, InfoResponseValidationError{
						field:  fmt.Sprintf("Inventory[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InfoResponseValidationError{
					field:  fmt.Sprintf("Inventory[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCoinHistory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InfoResponseValidationError{
					field:  "CoinHistory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InfoResponseValidationError{
					field:  "CoinHistory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCoinHistory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InfoResponseValidationError{
				field:  "CoinHistory",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return InfoResponseMultiError(errors)
	}

	return nil
}

// InfoResponseMultiError is an error wrapping multiple validation errors
// returned by InfoResponse.ValidateAll() if the designated constraints aren't met.
type InfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InfoResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InfoResponseMultiError) AllErrors() []error { return m }

// InfoResponseValidationError is the validation error returned by
// InfoResponse.Validate if the designated constraints aren't met.
type InfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoResponseValidationError) ErrorName() string { return "InfoResponseValidationError" }

// Error satisfies the builtin error interface
func (e InfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoResponseValidationError{}

// Validate checks the field values on SendCoinRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SendCoinRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendCoinRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendCoinRequestMultiError, or nil if none found.
func (m *SendCoinRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SendCoinRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ToUser

	if m.GetAmount() <= 0 {
		err := SendCoinRequestValidationError{
			field:  "Amount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SendCoinRequestMultiError(errors)
	}

	return nil
}

// SendCoinRequestMultiError is an error wrapping multiple validation errors
// returned by SendCoinRequest.ValidateAll() if the designated constraints
// aren't met.
type SendCoinRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendCoinRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendCoinRequestMultiError) AllErrors() []error { return m }

// SendCoinRequestValidationError is the validation error returned by
// SendCoinRequest.Validate if the designated constraints aren't met.
type SendCoinRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendCoinRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendCoinRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendCoinRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendCoinRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendCoinRequestValidationError) ErrorName() string { return "SendCoinRequestValidationError" }

// Error satisfies the builtin error interface
func (e SendCoinRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendCoinRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendCoinRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendCoinRequestValidationError{}

// Validate checks the field values on BuyItemRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BuyItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BuyItemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BuyItemRequestMultiError,
// or nil if none found.
func (m *BuyItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BuyItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Item

	if len(errors) > 0 {
		return BuyItemRequestMultiError(errors)
	}

	return nil
}

// BuyItemRequestMultiError is an error wrapping multiple validation errors
// returned by BuyItemRequest.ValidateAll() if the designated constraints
// aren't met.
type BuyItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BuyItemRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BuyItemRequestMultiError) AllErrors() []error { return m }

// BuyItemRequestValidationError is the validation error returned by
// BuyItemRequest.Validate if the designated constraints aren't met.
type BuyItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BuyItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BuyItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BuyItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BuyItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BuyItemRequestValidationError) ErrorName() string { return "BuyItemRequestValidationError" }

// Error satisfies the builtin error interface
func (e BuyItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuyItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BuyItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BuyItemRequestValidationError{}

// Validate checks the field values on AuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthRequestMultiError, or
// nil if none found.
func (m *AuthRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetLogin()) < 3 {
		err := AuthRequestValidationError{
			field:  "Login",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		err := AuthRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AuthRequestMultiError(errors)
	}

	return nil
}

// AuthRequestMultiError is an error wrapping multiple validation errors
// returned by AuthRequest.ValidateAll() if the designated constraints aren't met.
type AuthRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthRequestMultiError) AllErrors() []error { return m }

// AuthRequestValidationError is the validation error returned by
// AuthRequest.Validate if the designated constraints aren't met.
type AuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthRequestValidationError) ErrorName() string { return "AuthRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthRequestValidationError{}

// Validate checks the field values on AuthResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthResponseMultiError, or
// nil if none found.
func (m *AuthResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return AuthResponseMultiError(errors)
	}

	return nil
}

// AuthResponseMultiError is an error wrapping multiple validation errors
// returned by AuthResponse.ValidateAll() if the designated constraints aren't met.
type AuthResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthResponseMultiError) AllErrors() []error { return m }

// AuthResponseValidationError is the validation error returned by
// AuthResponse.Validate if the designated constraints aren't met.
type AuthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthResponseValidationError) ErrorName() string { return "AuthResponseValidationError" }

// Error satisfies the builtin error interface
func (e AuthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthResponseValidationError{}

// Validate checks the field values on InfoResponseCoinHistoryMessage with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InfoResponseCoinHistoryMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InfoResponseCoinHistoryMessage with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// InfoResponseCoinHistoryMessageMultiError, or nil if none found.
func (m *InfoResponseCoinHistoryMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *InfoResponseCoinHistoryMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSent() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, InfoResponseCoinHistoryMessageValidationError{
						field:  fmt.Sprintf("Sent[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, InfoResponseCoinHistoryMessageValidationError{
						field:  fmt.Sprintf("Sent[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InfoResponseCoinHistoryMessageValidationError{
					field:  fmt.Sprintf("Sent[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetReceived() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, InfoResponseCoinHistoryMessageValidationError{
						field:  fmt.Sprintf("Received[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, InfoResponseCoinHistoryMessageValidationError{
						field:  fmt.Sprintf("Received[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InfoResponseCoinHistoryMessageValidationError{
					field:  fmt.Sprintf("Received[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return InfoResponseCoinHistoryMessageMultiError(errors)
	}

	return nil
}

// InfoResponseCoinHistoryMessageMultiError is an error wrapping multiple
// validation errors returned by InfoResponseCoinHistoryMessage.ValidateAll()
// if the designated constraints aren't met.
type InfoResponseCoinHistoryMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InfoResponseCoinHistoryMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InfoResponseCoinHistoryMessageMultiError) AllErrors() []error { return m }

// InfoResponseCoinHistoryMessageValidationError is the validation error
// returned by InfoResponseCoinHistoryMessage.Validate if the designated
// constraints aren't met.
type InfoResponseCoinHistoryMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoResponseCoinHistoryMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoResponseCoinHistoryMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoResponseCoinHistoryMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoResponseCoinHistoryMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoResponseCoinHistoryMessageValidationError) ErrorName() string {
	return "InfoResponseCoinHistoryMessageValidationError"
}

// Error satisfies the builtin error interface
func (e InfoResponseCoinHistoryMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoResponseCoinHistoryMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoResponseCoinHistoryMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoResponseCoinHistoryMessageValidationError{}

// Validate checks the field values on InfoResponseItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *InfoResponseItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InfoResponseItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InfoResponseItemMultiError, or nil if none found.
func (m *InfoResponseItem) ValidateAll() error {
	return m.validate(true)
}

func (m *InfoResponseItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Quantity

	if len(errors) > 0 {
		return InfoResponseItemMultiError(errors)
	}

	return nil
}

// InfoResponseItemMultiError is an error wrapping multiple validation errors
// returned by InfoResponseItem.ValidateAll() if the designated constraints
// aren't met.
type InfoResponseItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InfoResponseItemMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InfoResponseItemMultiError) AllErrors() []error { return m }

// InfoResponseItemValidationError is the validation error returned by
// InfoResponseItem.Validate if the designated constraints aren't met.
type InfoResponseItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoResponseItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoResponseItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoResponseItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoResponseItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoResponseItemValidationError) ErrorName() string { return "InfoResponseItemValidationError" }

// Error satisfies the builtin error interface
func (e InfoResponseItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoResponseItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoResponseItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoResponseItemValidationError{}

// Validate checks the field values on
// InfoResponseCoinHistoryMessageSendCoinEntry with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InfoResponseCoinHistoryMessageSendCoinEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// InfoResponseCoinHistoryMessageSendCoinEntry with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// InfoResponseCoinHistoryMessageSendCoinEntryMultiError, or nil if none found.
func (m *InfoResponseCoinHistoryMessageSendCoinEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *InfoResponseCoinHistoryMessageSendCoinEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ToUser

	// no validation rules for Amount

	if len(errors) > 0 {
		return InfoResponseCoinHistoryMessageSendCoinEntryMultiError(errors)
	}

	return nil
}

// InfoResponseCoinHistoryMessageSendCoinEntryMultiError is an error wrapping
// multiple validation errors returned by
// InfoResponseCoinHistoryMessageSendCoinEntry.ValidateAll() if the designated
// constraints aren't met.
type InfoResponseCoinHistoryMessageSendCoinEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InfoResponseCoinHistoryMessageSendCoinEntryMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InfoResponseCoinHistoryMessageSendCoinEntryMultiError) AllErrors() []error { return m }

// InfoResponseCoinHistoryMessageSendCoinEntryValidationError is the validation
// error returned by InfoResponseCoinHistoryMessageSendCoinEntry.Validate if
// the designated constraints aren't met.
type InfoResponseCoinHistoryMessageSendCoinEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoResponseCoinHistoryMessageSendCoinEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoResponseCoinHistoryMessageSendCoinEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoResponseCoinHistoryMessageSendCoinEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoResponseCoinHistoryMessageSendCoinEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoResponseCoinHistoryMessageSendCoinEntryValidationError) ErrorName() string {
	return "InfoResponseCoinHistoryMessageSendCoinEntryValidationError"
}

// Error satisfies the builtin error interface
func (e InfoResponseCoinHistoryMessageSendCoinEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoResponseCoinHistoryMessageSendCoinEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoResponseCoinHistoryMessageSendCoinEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoResponseCoinHistoryMessageSendCoinEntryValidationError{}

// Validate checks the field values on
// InfoResponseCoinHistoryMessageReceiveCoinEntry with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InfoResponseCoinHistoryMessageReceiveCoinEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// InfoResponseCoinHistoryMessageReceiveCoinEntry with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// InfoResponseCoinHistoryMessageReceiveCoinEntryMultiError, or nil if none found.
func (m *InfoResponseCoinHistoryMessageReceiveCoinEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *InfoResponseCoinHistoryMessageReceiveCoinEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FromUser

	// no validation rules for Amount

	if len(errors) > 0 {
		return InfoResponseCoinHistoryMessageReceiveCoinEntryMultiError(errors)
	}

	return nil
}

// InfoResponseCoinHistoryMessageReceiveCoinEntryMultiError is an error
// wrapping multiple validation errors returned by
// InfoResponseCoinHistoryMessageReceiveCoinEntry.ValidateAll() if the
// designated constraints aren't met.
type InfoResponseCoinHistoryMessageReceiveCoinEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InfoResponseCoinHistoryMessageReceiveCoinEntryMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InfoResponseCoinHistoryMessageReceiveCoinEntryMultiError) AllErrors() []error { return m }

// InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError is the
// validation error returned by
// InfoResponseCoinHistoryMessageReceiveCoinEntry.Validate if the designated
// constraints aren't met.
type InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError) ErrorName() string {
	return "InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError"
}

// Error satisfies the builtin error interface
func (e InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoResponseCoinHistoryMessageReceiveCoinEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoResponseCoinHistoryMessageReceiveCoinEntryValidationError{}
