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

// OperatorGroupMember struct for OperatorGroupMember
type OperatorGroupMember struct {
	// The unique identifier of this Operator
	OperatorGuid string `json:"OperatorGuid"`
}

// NewOperatorGroupMember instantiates a new OperatorGroupMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorGroupMember(operatorGuid string) *OperatorGroupMember {
	this := OperatorGroupMember{}
	this.OperatorGuid = operatorGuid
	return &this
}

// NewOperatorGroupMemberWithDefaults instantiates a new OperatorGroupMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorGroupMemberWithDefaults() *OperatorGroupMember {
	this := OperatorGroupMember{}
	return &this
}

// GetOperatorGuid returns the OperatorGuid field value
func (o *OperatorGroupMember) GetOperatorGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatorGuid
}

// GetOperatorGuidOk returns a tuple with the OperatorGuid field value
// and a boolean to check if the value has been set.
func (o *OperatorGroupMember) GetOperatorGuidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OperatorGuid, true
}

// SetOperatorGuid sets field value
func (o *OperatorGroupMember) SetOperatorGuid(v string) {
	o.OperatorGuid = v
}

func (o OperatorGroupMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["OperatorGuid"] = o.OperatorGuid
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorGroupMember struct {
	value *OperatorGroupMember
	isSet bool
}

func (v NullableOperatorGroupMember) Get() *OperatorGroupMember {
	return v.value
}

func (v *NullableOperatorGroupMember) Set(val *OperatorGroupMember) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorGroupMember) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorGroupMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorGroupMember(val *OperatorGroupMember) *NullableOperatorGroupMember {
	return &NullableOperatorGroupMember{value: val, isSet: true}
}

func (v NullableOperatorGroupMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorGroupMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


