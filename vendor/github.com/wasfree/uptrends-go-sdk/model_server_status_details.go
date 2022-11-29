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

// ServerStatusDetails ServerStatusDetails
type ServerStatusDetails struct {
	// The UTC time the server was last online
	LastOnlineUtc map[string]interface{} `json:"LastOnlineUtc,omitempty"`
}

// NewServerStatusDetails instantiates a new ServerStatusDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStatusDetails() *ServerStatusDetails {
	this := ServerStatusDetails{}
	return &this
}

// NewServerStatusDetailsWithDefaults instantiates a new ServerStatusDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStatusDetailsWithDefaults() *ServerStatusDetails {
	this := ServerStatusDetails{}
	return &this
}

// GetLastOnlineUtc returns the LastOnlineUtc field value if set, zero value otherwise.
func (o *ServerStatusDetails) GetLastOnlineUtc() map[string]interface{} {
	if o == nil || isNil(o.LastOnlineUtc) {
		var ret map[string]interface{}
		return ret
	}
	return o.LastOnlineUtc
}

// GetLastOnlineUtcOk returns a tuple with the LastOnlineUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatusDetails) GetLastOnlineUtcOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.LastOnlineUtc) {
    return map[string]interface{}{}, false
	}
	return o.LastOnlineUtc, true
}

// HasLastOnlineUtc returns a boolean if a field has been set.
func (o *ServerStatusDetails) HasLastOnlineUtc() bool {
	if o != nil && !isNil(o.LastOnlineUtc) {
		return true
	}

	return false
}

// SetLastOnlineUtc gets a reference to the given map[string]interface{} and assigns it to the LastOnlineUtc field.
func (o *ServerStatusDetails) SetLastOnlineUtc(v map[string]interface{}) {
	o.LastOnlineUtc = v
}

func (o ServerStatusDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LastOnlineUtc) {
		toSerialize["LastOnlineUtc"] = o.LastOnlineUtc
	}
	return json.Marshal(toSerialize)
}

type NullableServerStatusDetails struct {
	value *ServerStatusDetails
	isSet bool
}

func (v NullableServerStatusDetails) Get() *ServerStatusDetails {
	return v.value
}

func (v *NullableServerStatusDetails) Set(val *ServerStatusDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStatusDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStatusDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStatusDetails(val *ServerStatusDetails) *NullableServerStatusDetails {
	return &NullableServerStatusDetails{value: val, isSet: true}
}

func (v NullableServerStatusDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStatusDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


