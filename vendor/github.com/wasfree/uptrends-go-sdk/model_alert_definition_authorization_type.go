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

// AlertDefinitionAuthorizationType 
type AlertDefinitionAuthorizationType string

// List of AlertDefinitionAuthorizationType
const (
	ALERTDEFINITIONAUTHORIZATIONTYPE_EDIT_ALERT_DEFINITION AlertDefinitionAuthorizationType = "EditAlertDefinition"
)

// All allowed values of AlertDefinitionAuthorizationType enum
var AllowedAlertDefinitionAuthorizationTypeEnumValues = []AlertDefinitionAuthorizationType{
	"EditAlertDefinition",
}

func (v *AlertDefinitionAuthorizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertDefinitionAuthorizationType(value)
	for _, existing := range AllowedAlertDefinitionAuthorizationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertDefinitionAuthorizationType", value)
}

// NewAlertDefinitionAuthorizationTypeFromValue returns a pointer to a valid AlertDefinitionAuthorizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertDefinitionAuthorizationTypeFromValue(v string) (*AlertDefinitionAuthorizationType, error) {
	ev := AlertDefinitionAuthorizationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertDefinitionAuthorizationType: valid values are %v", v, AllowedAlertDefinitionAuthorizationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertDefinitionAuthorizationType) IsValid() bool {
	for _, existing := range AllowedAlertDefinitionAuthorizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertDefinitionAuthorizationType value
func (v AlertDefinitionAuthorizationType) Ptr() *AlertDefinitionAuthorizationType {
	return &v
}

type NullableAlertDefinitionAuthorizationType struct {
	value *AlertDefinitionAuthorizationType
	isSet bool
}

func (v NullableAlertDefinitionAuthorizationType) Get() *AlertDefinitionAuthorizationType {
	return v.value
}

func (v *NullableAlertDefinitionAuthorizationType) Set(val *AlertDefinitionAuthorizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertDefinitionAuthorizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertDefinitionAuthorizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertDefinitionAuthorizationType(val *AlertDefinitionAuthorizationType) *NullableAlertDefinitionAuthorizationType {
	return &NullableAlertDefinitionAuthorizationType{value: val, isSet: true}
}

func (v NullableAlertDefinitionAuthorizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertDefinitionAuthorizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

