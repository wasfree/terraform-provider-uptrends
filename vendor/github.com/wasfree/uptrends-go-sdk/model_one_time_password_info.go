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

// OneTimePasswordInfo struct for OneTimePasswordInfo
type OneTimePasswordInfo struct {
	Secret *string `json:"Secret,omitempty"`
	Digits int32 `json:"Digits"`
	Period int32 `json:"Period"`
	HashAlgorithm HashAlgorithm `json:"HashAlgorithm"`
}

// NewOneTimePasswordInfo instantiates a new OneTimePasswordInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneTimePasswordInfo(digits int32, period int32, hashAlgorithm HashAlgorithm) *OneTimePasswordInfo {
	this := OneTimePasswordInfo{}
	this.Digits = digits
	this.Period = period
	this.HashAlgorithm = hashAlgorithm
	return &this
}

// NewOneTimePasswordInfoWithDefaults instantiates a new OneTimePasswordInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneTimePasswordInfoWithDefaults() *OneTimePasswordInfo {
	this := OneTimePasswordInfo{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *OneTimePasswordInfo) GetSecret() string {
	if o == nil || isNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneTimePasswordInfo) GetSecretOk() (*string, bool) {
	if o == nil || isNil(o.Secret) {
    return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *OneTimePasswordInfo) HasSecret() bool {
	if o != nil && !isNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *OneTimePasswordInfo) SetSecret(v string) {
	o.Secret = &v
}

// GetDigits returns the Digits field value
func (o *OneTimePasswordInfo) GetDigits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Digits
}

// GetDigitsOk returns a tuple with the Digits field value
// and a boolean to check if the value has been set.
func (o *OneTimePasswordInfo) GetDigitsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Digits, true
}

// SetDigits sets field value
func (o *OneTimePasswordInfo) SetDigits(v int32) {
	o.Digits = v
}

// GetPeriod returns the Period field value
func (o *OneTimePasswordInfo) GetPeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *OneTimePasswordInfo) GetPeriodOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *OneTimePasswordInfo) SetPeriod(v int32) {
	o.Period = v
}

// GetHashAlgorithm returns the HashAlgorithm field value
func (o *OneTimePasswordInfo) GetHashAlgorithm() HashAlgorithm {
	if o == nil {
		var ret HashAlgorithm
		return ret
	}

	return o.HashAlgorithm
}

// GetHashAlgorithmOk returns a tuple with the HashAlgorithm field value
// and a boolean to check if the value has been set.
func (o *OneTimePasswordInfo) GetHashAlgorithmOk() (*HashAlgorithm, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HashAlgorithm, true
}

// SetHashAlgorithm sets field value
func (o *OneTimePasswordInfo) SetHashAlgorithm(v HashAlgorithm) {
	o.HashAlgorithm = v
}

func (o OneTimePasswordInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Secret) {
		toSerialize["Secret"] = o.Secret
	}
	if true {
		toSerialize["Digits"] = o.Digits
	}
	if true {
		toSerialize["Period"] = o.Period
	}
	if true {
		toSerialize["HashAlgorithm"] = o.HashAlgorithm
	}
	return json.Marshal(toSerialize)
}

type NullableOneTimePasswordInfo struct {
	value *OneTimePasswordInfo
	isSet bool
}

func (v NullableOneTimePasswordInfo) Get() *OneTimePasswordInfo {
	return v.value
}

func (v *NullableOneTimePasswordInfo) Set(val *OneTimePasswordInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOneTimePasswordInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOneTimePasswordInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneTimePasswordInfo(val *OneTimePasswordInfo) *NullableOneTimePasswordInfo {
	return &NullableOneTimePasswordInfo{value: val, isSet: true}
}

func (v NullableOneTimePasswordInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneTimePasswordInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


