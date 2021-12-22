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

// ApiAssertionSourceType 
type ApiAssertionSourceType string

// List of ApiAssertionSourceType
const (
	APIASSERTIONSOURCETYPE_NONE ApiAssertionSourceType = "None"
	APIASSERTIONSOURCETYPE_CONSTANT ApiAssertionSourceType = "Constant"
	APIASSERTIONSOURCETYPE_VARIABLE ApiAssertionSourceType = "Variable"
	APIASSERTIONSOURCETYPE_RESPONSE_STATUS_CODE ApiAssertionSourceType = "ResponseStatusCode"
	APIASSERTIONSOURCETYPE_RESPONSE_STATUS_DESCRIPTION ApiAssertionSourceType = "ResponseStatusDescription"
	APIASSERTIONSOURCETYPE_RESPONSE_BODY_JSON ApiAssertionSourceType = "ResponseBodyJson"
	APIASSERTIONSOURCETYPE_RESPONSE_COMPLETED ApiAssertionSourceType = "ResponseCompleted"
	APIASSERTIONSOURCETYPE_RESPONSE_BODY_TEXT ApiAssertionSourceType = "ResponseBodyText"
	APIASSERTIONSOURCETYPE_FAIL ApiAssertionSourceType = "Fail"
	APIASSERTIONSOURCETYPE_CONTENT_LENGTH_CALCULATED ApiAssertionSourceType = "ContentLengthCalculated"
	APIASSERTIONSOURCETYPE_DURATION ApiAssertionSourceType = "Duration"
	APIASSERTIONSOURCETYPE_SUM ApiAssertionSourceType = "Sum"
	APIASSERTIONSOURCETYPE_CONCATENATION ApiAssertionSourceType = "Concatenation"
	APIASSERTIONSOURCETYPE_TO_BASE64 ApiAssertionSourceType = "ToBase64"
	APIASSERTIONSOURCETYPE_TO_SHA1_HEX ApiAssertionSourceType = "ToSHA1Hex"
	APIASSERTIONSOURCETYPE_TO_MD5_HEX ApiAssertionSourceType = "ToMD5Hex"
	APIASSERTIONSOURCETYPE_RESPONSE_BODY_XML ApiAssertionSourceType = "ResponseBodyXml"
	APIASSERTIONSOURCETYPE_RESPONSE_HEADER ApiAssertionSourceType = "ResponseHeader"
	APIASSERTIONSOURCETYPE_COOKIE ApiAssertionSourceType = "Cookie"
	APIASSERTIONSOURCETYPE_VARIABLES_RESOLVED ApiAssertionSourceType = "VariablesResolved"
	APIASSERTIONSOURCETYPE_CUMULATIVE_DURATION ApiAssertionSourceType = "CumulativeDuration"
	APIASSERTIONSOURCETYPE_RESPONSE_HAS_EXCEPTION ApiAssertionSourceType = "ResponseHasException"
)

var allowedApiAssertionSourceTypeEnumValues = []ApiAssertionSourceType{
	"None",
	"Constant",
	"Variable",
	"ResponseStatusCode",
	"ResponseStatusDescription",
	"ResponseBodyJson",
	"ResponseCompleted",
	"ResponseBodyText",
	"Fail",
	"ContentLengthCalculated",
	"Duration",
	"Sum",
	"Concatenation",
	"ToBase64",
	"ToSHA1Hex",
	"ToMD5Hex",
	"ResponseBodyXml",
	"ResponseHeader",
	"Cookie",
	"VariablesResolved",
	"CumulativeDuration",
	"ResponseHasException",
}

func (v *ApiAssertionSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiAssertionSourceType(value)
	for _, existing := range allowedApiAssertionSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiAssertionSourceType", value)
}

// NewApiAssertionSourceTypeFromValue returns a pointer to a valid ApiAssertionSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiAssertionSourceTypeFromValue(v string) (*ApiAssertionSourceType, error) {
	ev := ApiAssertionSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiAssertionSourceType: valid values are %v", v, allowedApiAssertionSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiAssertionSourceType) IsValid() bool {
	for _, existing := range allowedApiAssertionSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiAssertionSourceType value
func (v ApiAssertionSourceType) Ptr() *ApiAssertionSourceType {
	return &v
}

type NullableApiAssertionSourceType struct {
	value *ApiAssertionSourceType
	isSet bool
}

func (v NullableApiAssertionSourceType) Get() *ApiAssertionSourceType {
	return v.value
}

func (v *NullableApiAssertionSourceType) Set(val *ApiAssertionSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAssertionSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAssertionSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAssertionSourceType(val *ApiAssertionSourceType) *NullableApiAssertionSourceType {
	return &NullableApiAssertionSourceType{value: val, isSet: true}
}

func (v NullableApiAssertionSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAssertionSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

