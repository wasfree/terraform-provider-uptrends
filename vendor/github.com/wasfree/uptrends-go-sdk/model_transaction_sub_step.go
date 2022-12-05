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

// TransactionSubStep struct for TransactionSubStep
type TransactionSubStep struct {
	Name *string `json:"Name,omitempty"`
	Type SubStepType `json:"Type"`
	Url *string `json:"Url,omitempty"`
	SetValue *string `json:"SetValue,omitempty"`
}

// NewTransactionSubStep instantiates a new TransactionSubStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionSubStep(type_ SubStepType) *TransactionSubStep {
	this := TransactionSubStep{}
	this.Type = type_
	return &this
}

// NewTransactionSubStepWithDefaults instantiates a new TransactionSubStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionSubStepWithDefaults() *TransactionSubStep {
	this := TransactionSubStep{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransactionSubStep) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSubStep) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransactionSubStep) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransactionSubStep) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *TransactionSubStep) GetType() SubStepType {
	if o == nil {
		var ret SubStepType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionSubStep) GetTypeOk() (*SubStepType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionSubStep) SetType(v SubStepType) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TransactionSubStep) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSubStep) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TransactionSubStep) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TransactionSubStep) SetUrl(v string) {
	o.Url = &v
}

// GetSetValue returns the SetValue field value if set, zero value otherwise.
func (o *TransactionSubStep) GetSetValue() string {
	if o == nil || isNil(o.SetValue) {
		var ret string
		return ret
	}
	return *o.SetValue
}

// GetSetValueOk returns a tuple with the SetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSubStep) GetSetValueOk() (*string, bool) {
	if o == nil || isNil(o.SetValue) {
    return nil, false
	}
	return o.SetValue, true
}

// HasSetValue returns a boolean if a field has been set.
func (o *TransactionSubStep) HasSetValue() bool {
	if o != nil && !isNil(o.SetValue) {
		return true
	}

	return false
}

// SetSetValue gets a reference to the given string and assigns it to the SetValue field.
func (o *TransactionSubStep) SetSetValue(v string) {
	o.SetValue = &v
}

func (o TransactionSubStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Type"] = o.Type
	}
	if !isNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !isNil(o.SetValue) {
		toSerialize["SetValue"] = o.SetValue
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionSubStep struct {
	value *TransactionSubStep
	isSet bool
}

func (v NullableTransactionSubStep) Get() *TransactionSubStep {
	return v.value
}

func (v *NullableTransactionSubStep) Set(val *TransactionSubStep) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionSubStep) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionSubStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionSubStep(val *TransactionSubStep) *NullableTransactionSubStep {
	return &NullableTransactionSubStep{value: val, isSet: true}
}

func (v NullableTransactionSubStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionSubStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


