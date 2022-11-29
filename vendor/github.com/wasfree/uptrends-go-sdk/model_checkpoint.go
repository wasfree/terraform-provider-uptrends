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

// Checkpoint Checkpoint
type Checkpoint struct {
	// The Id of the checkpoint
	CheckpointId int32 `json:"CheckpointId"`
	// The name of the checkpoint
	CheckpointName *string `json:"CheckpointName,omitempty"`
	// The location code of the checkpoint
	Code *string `json:"Code,omitempty"`
	// The IPv4 addresses of the checkpoint 
	Ipv4Addresses []string `json:"Ipv4Addresses,omitempty"`
	// The IPv6 addresses of the checkpoint
	Ipv6Addresses []Ipv6Address `json:"Ipv6Addresses,omitempty"`
	// Checkpoint is primary
	IsPrimaryCheckpoint bool `json:"IsPrimaryCheckpoint"`
	// Checkpoint supports IPv6
	SupportsIpv6 bool `json:"SupportsIpv6"`
	// Checkpoint has high availability
	HasHighAvailability bool `json:"HasHighAvailability"`
}

// NewCheckpoint instantiates a new Checkpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckpoint(checkpointId int32, isPrimaryCheckpoint bool, supportsIpv6 bool, hasHighAvailability bool) *Checkpoint {
	this := Checkpoint{}
	this.CheckpointId = checkpointId
	this.IsPrimaryCheckpoint = isPrimaryCheckpoint
	this.SupportsIpv6 = supportsIpv6
	this.HasHighAvailability = hasHighAvailability
	return &this
}

// NewCheckpointWithDefaults instantiates a new Checkpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckpointWithDefaults() *Checkpoint {
	this := Checkpoint{}
	return &this
}

// GetCheckpointId returns the CheckpointId field value
func (o *Checkpoint) GetCheckpointId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CheckpointId
}

// GetCheckpointIdOk returns a tuple with the CheckpointId field value
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetCheckpointIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CheckpointId, true
}

// SetCheckpointId sets field value
func (o *Checkpoint) SetCheckpointId(v int32) {
	o.CheckpointId = v
}

// GetCheckpointName returns the CheckpointName field value if set, zero value otherwise.
func (o *Checkpoint) GetCheckpointName() string {
	if o == nil || isNil(o.CheckpointName) {
		var ret string
		return ret
	}
	return *o.CheckpointName
}

// GetCheckpointNameOk returns a tuple with the CheckpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetCheckpointNameOk() (*string, bool) {
	if o == nil || isNil(o.CheckpointName) {
    return nil, false
	}
	return o.CheckpointName, true
}

// HasCheckpointName returns a boolean if a field has been set.
func (o *Checkpoint) HasCheckpointName() bool {
	if o != nil && !isNil(o.CheckpointName) {
		return true
	}

	return false
}

// SetCheckpointName gets a reference to the given string and assigns it to the CheckpointName field.
func (o *Checkpoint) SetCheckpointName(v string) {
	o.CheckpointName = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Checkpoint) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Checkpoint) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Checkpoint) SetCode(v string) {
	o.Code = &v
}

// GetIpv4Addresses returns the Ipv4Addresses field value if set, zero value otherwise.
func (o *Checkpoint) GetIpv4Addresses() []string {
	if o == nil || isNil(o.Ipv4Addresses) {
		var ret []string
		return ret
	}
	return o.Ipv4Addresses
}

// GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetIpv4AddressesOk() ([]string, bool) {
	if o == nil || isNil(o.Ipv4Addresses) {
    return nil, false
	}
	return o.Ipv4Addresses, true
}

// HasIpv4Addresses returns a boolean if a field has been set.
func (o *Checkpoint) HasIpv4Addresses() bool {
	if o != nil && !isNil(o.Ipv4Addresses) {
		return true
	}

	return false
}

// SetIpv4Addresses gets a reference to the given []string and assigns it to the Ipv4Addresses field.
func (o *Checkpoint) SetIpv4Addresses(v []string) {
	o.Ipv4Addresses = v
}

// GetIpv6Addresses returns the Ipv6Addresses field value if set, zero value otherwise.
func (o *Checkpoint) GetIpv6Addresses() []Ipv6Address {
	if o == nil || isNil(o.Ipv6Addresses) {
		var ret []Ipv6Address
		return ret
	}
	return o.Ipv6Addresses
}

// GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetIpv6AddressesOk() ([]Ipv6Address, bool) {
	if o == nil || isNil(o.Ipv6Addresses) {
    return nil, false
	}
	return o.Ipv6Addresses, true
}

// HasIpv6Addresses returns a boolean if a field has been set.
func (o *Checkpoint) HasIpv6Addresses() bool {
	if o != nil && !isNil(o.Ipv6Addresses) {
		return true
	}

	return false
}

// SetIpv6Addresses gets a reference to the given []Ipv6Address and assigns it to the Ipv6Addresses field.
func (o *Checkpoint) SetIpv6Addresses(v []Ipv6Address) {
	o.Ipv6Addresses = v
}

// GetIsPrimaryCheckpoint returns the IsPrimaryCheckpoint field value
func (o *Checkpoint) GetIsPrimaryCheckpoint() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrimaryCheckpoint
}

// GetIsPrimaryCheckpointOk returns a tuple with the IsPrimaryCheckpoint field value
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetIsPrimaryCheckpointOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsPrimaryCheckpoint, true
}

// SetIsPrimaryCheckpoint sets field value
func (o *Checkpoint) SetIsPrimaryCheckpoint(v bool) {
	o.IsPrimaryCheckpoint = v
}

// GetSupportsIpv6 returns the SupportsIpv6 field value
func (o *Checkpoint) GetSupportsIpv6() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsIpv6
}

// GetSupportsIpv6Ok returns a tuple with the SupportsIpv6 field value
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetSupportsIpv6Ok() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SupportsIpv6, true
}

// SetSupportsIpv6 sets field value
func (o *Checkpoint) SetSupportsIpv6(v bool) {
	o.SupportsIpv6 = v
}

// GetHasHighAvailability returns the HasHighAvailability field value
func (o *Checkpoint) GetHasHighAvailability() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasHighAvailability
}

// GetHasHighAvailabilityOk returns a tuple with the HasHighAvailability field value
// and a boolean to check if the value has been set.
func (o *Checkpoint) GetHasHighAvailabilityOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HasHighAvailability, true
}

// SetHasHighAvailability sets field value
func (o *Checkpoint) SetHasHighAvailability(v bool) {
	o.HasHighAvailability = v
}

func (o Checkpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CheckpointId"] = o.CheckpointId
	}
	if !isNil(o.CheckpointName) {
		toSerialize["CheckpointName"] = o.CheckpointName
	}
	if !isNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !isNil(o.Ipv4Addresses) {
		toSerialize["Ipv4Addresses"] = o.Ipv4Addresses
	}
	if !isNil(o.Ipv6Addresses) {
		toSerialize["Ipv6Addresses"] = o.Ipv6Addresses
	}
	if true {
		toSerialize["IsPrimaryCheckpoint"] = o.IsPrimaryCheckpoint
	}
	if true {
		toSerialize["SupportsIpv6"] = o.SupportsIpv6
	}
	if true {
		toSerialize["HasHighAvailability"] = o.HasHighAvailability
	}
	return json.Marshal(toSerialize)
}

type NullableCheckpoint struct {
	value *Checkpoint
	isSet bool
}

func (v NullableCheckpoint) Get() *Checkpoint {
	return v.value
}

func (v *NullableCheckpoint) Set(val *Checkpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckpoint(val *Checkpoint) *NullableCheckpoint {
	return &NullableCheckpoint{value: val, isSet: true}
}

func (v NullableCheckpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


