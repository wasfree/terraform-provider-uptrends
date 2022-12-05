/*
Uptrends API v4

This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"encoding/json"
)

// ApiHttpHeaderValue struct for ApiHttpHeaderValue
type ApiHttpHeaderValue struct {
	Key *string `json:"Key,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewApiHttpHeaderValue instantiates a new ApiHttpHeaderValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHttpHeaderValue() *ApiHttpHeaderValue {
	this := ApiHttpHeaderValue{}
	return &this
}

// NewApiHttpHeaderValueWithDefaults instantiates a new ApiHttpHeaderValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHttpHeaderValueWithDefaults() *ApiHttpHeaderValue {
	this := ApiHttpHeaderValue{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ApiHttpHeaderValue) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHttpHeaderValue) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
    return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ApiHttpHeaderValue) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ApiHttpHeaderValue) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiHttpHeaderValue) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHttpHeaderValue) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiHttpHeaderValue) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiHttpHeaderValue) SetValue(v string) {
	o.Value = &v
}

func (o ApiHttpHeaderValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !isNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableApiHttpHeaderValue struct {
	value *ApiHttpHeaderValue
	isSet bool
}

func (v NullableApiHttpHeaderValue) Get() *ApiHttpHeaderValue {
	return v.value
}

func (v *NullableApiHttpHeaderValue) Set(val *ApiHttpHeaderValue) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHttpHeaderValue) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHttpHeaderValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHttpHeaderValue(val *ApiHttpHeaderValue) *NullableApiHttpHeaderValue {
	return &NullableApiHttpHeaderValue{value: val, isSet: true}
}

func (v NullableApiHttpHeaderValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHttpHeaderValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


