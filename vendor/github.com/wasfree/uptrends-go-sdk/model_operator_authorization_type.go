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

// OperatorAuthorizationType 
type OperatorAuthorizationType string

// List of OperatorAuthorizationType
const (
	OPERATORAUTHORIZATIONTYPE_ACCOUNT_ACCESS OperatorAuthorizationType = "AccountAccess"
)

var allowedOperatorAuthorizationTypeEnumValues = []OperatorAuthorizationType{
	"AccountAccess",
}

func (v *OperatorAuthorizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperatorAuthorizationType(value)
	for _, existing := range allowedOperatorAuthorizationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperatorAuthorizationType", value)
}

// NewOperatorAuthorizationTypeFromValue returns a pointer to a valid OperatorAuthorizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperatorAuthorizationTypeFromValue(v string) (*OperatorAuthorizationType, error) {
	ev := OperatorAuthorizationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperatorAuthorizationType: valid values are %v", v, allowedOperatorAuthorizationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperatorAuthorizationType) IsValid() bool {
	for _, existing := range allowedOperatorAuthorizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperatorAuthorizationType value
func (v OperatorAuthorizationType) Ptr() *OperatorAuthorizationType {
	return &v
}

type NullableOperatorAuthorizationType struct {
	value *OperatorAuthorizationType
	isSet bool
}

func (v NullableOperatorAuthorizationType) Get() *OperatorAuthorizationType {
	return v.value
}

func (v *NullableOperatorAuthorizationType) Set(val *OperatorAuthorizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorAuthorizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorAuthorizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorAuthorizationType(val *OperatorAuthorizationType) *NullableOperatorAuthorizationType {
	return &NullableOperatorAuthorizationType{value: val, isSet: true}
}

func (v NullableOperatorAuthorizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorAuthorizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

