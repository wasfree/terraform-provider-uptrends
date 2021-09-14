/*
Uptrends API v4

This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"encoding/json"
	"fmt"
)

// SubStepType the model 'SubStepType'
type SubStepType string

// List of SubStepType
const (
	SUBSTEPTYPE_NAVIGATE SubStepType = "Navigate"
	SUBSTEPTYPE_CLICK SubStepType = "Click"
	SUBSTEPTYPE_SET SubStepType = "Set"
	SUBSTEPTYPE_TEST SubStepType = "Test"
	SUBSTEPTYPE_SCRIPT SubStepType = "Script"
	SUBSTEPTYPE_HOVER SubStepType = "Hover"
	SUBSTEPTYPE_SCREENSHOT SubStepType = "Screenshot"
	SUBSTEPTYPE_SCROLL_TO SubStepType = "ScrollTo"
	SUBSTEPTYPE_WAIT_FOR_ELEMENT SubStepType = "WaitForElement"
	SUBSTEPTYPE_SWITCH_TO_FRAME SubStepType = "SwitchToFrame"
	SUBSTEPTYPE_SWITCH_TO_TAB SubStepType = "SwitchToTab"
)

var allowedSubStepTypeEnumValues = []SubStepType{
	"Navigate",
	"Click",
	"Set",
	"Test",
	"Script",
	"Hover",
	"Screenshot",
	"ScrollTo",
	"WaitForElement",
	"SwitchToFrame",
	"SwitchToTab",
}

func (v *SubStepType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubStepType(value)
	for _, existing := range allowedSubStepTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubStepType", value)
}

// NewSubStepTypeFromValue returns a pointer to a valid SubStepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubStepTypeFromValue(v string) (*SubStepType, error) {
	ev := SubStepType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubStepType: valid values are %v", v, allowedSubStepTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubStepType) IsValid() bool {
	for _, existing := range allowedSubStepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubStepType value
func (v SubStepType) Ptr() *SubStepType {
	return &v
}

type NullableSubStepType struct {
	value *SubStepType
	isSet bool
}

func (v NullableSubStepType) Get() *SubStepType {
	return v.value
}

func (v *NullableSubStepType) Set(val *SubStepType) {
	v.value = val
	v.isSet = true
}

func (v NullableSubStepType) IsSet() bool {
	return v.isSet
}

func (v *NullableSubStepType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubStepType(val *SubStepType) *NullableSubStepType {
	return &NullableSubStepType{value: val, isSet: true}
}

func (v NullableSubStepType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubStepType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

