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

// CheckpointsHealth CheckpointsHealth
type CheckpointsHealth struct {
	// The number of healthy servers
	HealthyServers int32 `json:"HealthyServers"`
	// The health of each server
	Servers []ServerHealth `json:"Servers,omitempty"`
}

// NewCheckpointsHealth instantiates a new CheckpointsHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckpointsHealth(healthyServers int32) *CheckpointsHealth {
	this := CheckpointsHealth{}
	this.HealthyServers = healthyServers
	return &this
}

// NewCheckpointsHealthWithDefaults instantiates a new CheckpointsHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckpointsHealthWithDefaults() *CheckpointsHealth {
	this := CheckpointsHealth{}
	return &this
}

// GetHealthyServers returns the HealthyServers field value
func (o *CheckpointsHealth) GetHealthyServers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HealthyServers
}

// GetHealthyServersOk returns a tuple with the HealthyServers field value
// and a boolean to check if the value has been set.
func (o *CheckpointsHealth) GetHealthyServersOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HealthyServers, true
}

// SetHealthyServers sets field value
func (o *CheckpointsHealth) SetHealthyServers(v int32) {
	o.HealthyServers = v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *CheckpointsHealth) GetServers() []ServerHealth {
	if o == nil || isNil(o.Servers) {
		var ret []ServerHealth
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpointsHealth) GetServersOk() ([]ServerHealth, bool) {
	if o == nil || isNil(o.Servers) {
    return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *CheckpointsHealth) HasServers() bool {
	if o != nil && !isNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ServerHealth and assigns it to the Servers field.
func (o *CheckpointsHealth) SetServers(v []ServerHealth) {
	o.Servers = v
}

func (o CheckpointsHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["HealthyServers"] = o.HealthyServers
	}
	if !isNil(o.Servers) {
		toSerialize["Servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableCheckpointsHealth struct {
	value *CheckpointsHealth
	isSet bool
}

func (v NullableCheckpointsHealth) Get() *CheckpointsHealth {
	return v.value
}

func (v *NullableCheckpointsHealth) Set(val *CheckpointsHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckpointsHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckpointsHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckpointsHealth(val *CheckpointsHealth) *NullableCheckpointsHealth {
	return &NullableCheckpointsHealth{value: val, isSet: true}
}

func (v NullableCheckpointsHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckpointsHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


