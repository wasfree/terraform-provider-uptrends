/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"encoding/json"
	"fmt"
)

// EscalationMode 
type EscalationMode string

// List of EscalationMode
const (
	ESCALATIONMODE_ALERT_ON_ERROR_COUNT EscalationMode = "AlertOnErrorCount"
	ESCALATIONMODE_ALERT_ON_ERROR_DURATION EscalationMode = "AlertOnErrorDuration"
)

var allowedEscalationModeEnumValues = []EscalationMode{
	"AlertOnErrorCount",
	"AlertOnErrorDuration",
}

func (v *EscalationMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EscalationMode(value)
	for _, existing := range allowedEscalationModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EscalationMode", value)
}

// NewEscalationModeFromValue returns a pointer to a valid EscalationMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEscalationModeFromValue(v string) (*EscalationMode, error) {
	ev := EscalationMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EscalationMode: valid values are %v", v, allowedEscalationModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EscalationMode) IsValid() bool {
	for _, existing := range allowedEscalationModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EscalationMode value
func (v EscalationMode) Ptr() *EscalationMode {
	return &v
}

type NullableEscalationMode struct {
	value *EscalationMode
	isSet bool
}

func (v NullableEscalationMode) Get() *EscalationMode {
	return v.value
}

func (v *NullableEscalationMode) Set(val *EscalationMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEscalationMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEscalationMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEscalationMode(val *EscalationMode) *NullableEscalationMode {
	return &NullableEscalationMode{value: val, isSet: true}
}

func (v NullableEscalationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEscalationMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

