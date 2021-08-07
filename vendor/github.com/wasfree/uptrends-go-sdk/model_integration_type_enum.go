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

// IntegrationTypeEnum 
type IntegrationTypeEnum string

// List of IntegrationTypeEnum
const (
	INTEGRATIONTYPEENUM_SLACK IntegrationTypeEnum = "Slack"
	INTEGRATIONTYPEENUM_PAGER_DUTY IntegrationTypeEnum = "PagerDuty"
	INTEGRATIONTYPEENUM_SMS IntegrationTypeEnum = "Sms"
	INTEGRATIONTYPEENUM_EMAIL IntegrationTypeEnum = "Email"
	INTEGRATIONTYPEENUM_PHONE IntegrationTypeEnum = "Phone"
	INTEGRATIONTYPEENUM_STATUSHUB IntegrationTypeEnum = "Statushub"
	INTEGRATIONTYPEENUM_GENERIC_WEBHOOK IntegrationTypeEnum = "GenericWebhook"
)

var allowedIntegrationTypeEnumEnumValues = []IntegrationTypeEnum{
	"Slack",
	"PagerDuty",
	"Sms",
	"Email",
	"Phone",
	"Statushub",
	"GenericWebhook",
}

func (v *IntegrationTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntegrationTypeEnum(value)
	for _, existing := range allowedIntegrationTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntegrationTypeEnum", value)
}

// NewIntegrationTypeEnumFromValue returns a pointer to a valid IntegrationTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntegrationTypeEnumFromValue(v string) (*IntegrationTypeEnum, error) {
	ev := IntegrationTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntegrationTypeEnum: valid values are %v", v, allowedIntegrationTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntegrationTypeEnum) IsValid() bool {
	for _, existing := range allowedIntegrationTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationTypeEnum value
func (v IntegrationTypeEnum) Ptr() *IntegrationTypeEnum {
	return &v
}

type NullableIntegrationTypeEnum struct {
	value *IntegrationTypeEnum
	isSet bool
}

func (v NullableIntegrationTypeEnum) Get() *IntegrationTypeEnum {
	return v.value
}

func (v *NullableIntegrationTypeEnum) Set(val *IntegrationTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationTypeEnum(val *IntegrationTypeEnum) *NullableIntegrationTypeEnum {
	return &NullableIntegrationTypeEnum{value: val, isSet: true}
}

func (v NullableIntegrationTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

