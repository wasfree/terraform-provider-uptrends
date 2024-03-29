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

// MonitorMode 
type MonitorMode string

// List of MonitorMode
const (
	MONITORMODE_DEVELOPMENT MonitorMode = "Development"
	MONITORMODE_STAGING MonitorMode = "Staging"
	MONITORMODE_PRODUCTION MonitorMode = "Production"
)

// All allowed values of MonitorMode enum
var AllowedMonitorModeEnumValues = []MonitorMode{
	"Development",
	"Staging",
	"Production",
}

func (v *MonitorMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MonitorMode(value)
	for _, existing := range AllowedMonitorModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MonitorMode", value)
}

// NewMonitorModeFromValue returns a pointer to a valid MonitorMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitorModeFromValue(v string) (*MonitorMode, error) {
	ev := MonitorMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitorMode: valid values are %v", v, AllowedMonitorModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitorMode) IsValid() bool {
	for _, existing := range AllowedMonitorModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorMode value
func (v MonitorMode) Ptr() *MonitorMode {
	return &v
}

type NullableMonitorMode struct {
	value *MonitorMode
	isSet bool
}

func (v NullableMonitorMode) Get() *MonitorMode {
	return v.value
}

func (v *NullableMonitorMode) Set(val *MonitorMode) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorMode) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorMode(val *MonitorMode) *NullableMonitorMode {
	return &NullableMonitorMode{value: val, isSet: true}
}

func (v NullableMonitorMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

