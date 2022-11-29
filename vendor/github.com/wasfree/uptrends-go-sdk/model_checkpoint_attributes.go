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

// CheckpointAttributes Checkpoint attributes
type CheckpointAttributes struct {
	// The name of the checkpoint
	CheckpointName *string `json:"CheckpointName,omitempty"`
	// Location code of the checkpoint
	Code *string `json:"Code,omitempty"`
	// Ipv4 addresses of the checkpoint 
	Ipv4Addresses []string `json:"Ipv4Addresses,omitempty"`
	// Ipv6 addresses of the checkpoint
	IpV6Addresses []string `json:"IpV6Addresses,omitempty"`
	// Checkpoint is primary
	IsPrimaryCheckpoint bool `json:"IsPrimaryCheckpoint"`
	// Checkpoint supports IPv6
	SupportsIpv6 bool `json:"SupportsIpv6"`
	// Checkpoint has high availability
	HasHighAvailability bool `json:"HasHighAvailability"`
}

// NewCheckpointAttributes instantiates a new CheckpointAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckpointAttributes(isPrimaryCheckpoint bool, supportsIpv6 bool, hasHighAvailability bool) *CheckpointAttributes {
	this := CheckpointAttributes{}
	this.IsPrimaryCheckpoint = isPrimaryCheckpoint
	this.SupportsIpv6 = supportsIpv6
	this.HasHighAvailability = hasHighAvailability
	return &this
}

// NewCheckpointAttributesWithDefaults instantiates a new CheckpointAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckpointAttributesWithDefaults() *CheckpointAttributes {
	this := CheckpointAttributes{}
	return &this
}

// GetCheckpointName returns the CheckpointName field value if set, zero value otherwise.
func (o *CheckpointAttributes) GetCheckpointName() string {
	if o == nil || isNil(o.CheckpointName) {
		var ret string
		return ret
	}
	return *o.CheckpointName
}

// GetCheckpointNameOk returns a tuple with the CheckpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpointAttributes) GetCheckpointNameOk() (*string, bool) {
	if o == nil || isNil(o.CheckpointName) {
    return nil, false
	}
	return o.CheckpointName, true
}

// HasCheckpointName returns a boolean if a field has been set.
func (o *CheckpointAttributes) HasCheckpointName() bool {
	if o != nil && !isNil(o.CheckpointName) {
		return true
	}

	return false
}

// SetCheckpointName gets a reference to the given string and assigns it to the CheckpointName field.
func (o *CheckpointAttributes) SetCheckpointName(v string) {
	o.CheckpointName = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CheckpointAttributes) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpointAttributes) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CheckpointAttributes) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CheckpointAttributes) SetCode(v string) {
	o.Code = &v
}

// GetIpv4Addresses returns the Ipv4Addresses field value if set, zero value otherwise.
func (o *CheckpointAttributes) GetIpv4Addresses() []string {
	if o == nil || isNil(o.Ipv4Addresses) {
		var ret []string
		return ret
	}
	return o.Ipv4Addresses
}

// GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpointAttributes) GetIpv4AddressesOk() ([]string, bool) {
	if o == nil || isNil(o.Ipv4Addresses) {
    return nil, false
	}
	return o.Ipv4Addresses, true
}

// HasIpv4Addresses returns a boolean if a field has been set.
func (o *CheckpointAttributes) HasIpv4Addresses() bool {
	if o != nil && !isNil(o.Ipv4Addresses) {
		return true
	}

	return false
}

// SetIpv4Addresses gets a reference to the given []string and assigns it to the Ipv4Addresses field.
func (o *CheckpointAttributes) SetIpv4Addresses(v []string) {
	o.Ipv4Addresses = v
}

// GetIpV6Addresses returns the IpV6Addresses field value if set, zero value otherwise.
func (o *CheckpointAttributes) GetIpV6Addresses() []string {
	if o == nil || isNil(o.IpV6Addresses) {
		var ret []string
		return ret
	}
	return o.IpV6Addresses
}

// GetIpV6AddressesOk returns a tuple with the IpV6Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpointAttributes) GetIpV6AddressesOk() ([]string, bool) {
	if o == nil || isNil(o.IpV6Addresses) {
    return nil, false
	}
	return o.IpV6Addresses, true
}

// HasIpV6Addresses returns a boolean if a field has been set.
func (o *CheckpointAttributes) HasIpV6Addresses() bool {
	if o != nil && !isNil(o.IpV6Addresses) {
		return true
	}

	return false
}

// SetIpV6Addresses gets a reference to the given []string and assigns it to the IpV6Addresses field.
func (o *CheckpointAttributes) SetIpV6Addresses(v []string) {
	o.IpV6Addresses = v
}

// GetIsPrimaryCheckpoint returns the IsPrimaryCheckpoint field value
func (o *CheckpointAttributes) GetIsPrimaryCheckpoint() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrimaryCheckpoint
}

// GetIsPrimaryCheckpointOk returns a tuple with the IsPrimaryCheckpoint field value
// and a boolean to check if the value has been set.
func (o *CheckpointAttributes) GetIsPrimaryCheckpointOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsPrimaryCheckpoint, true
}

// SetIsPrimaryCheckpoint sets field value
func (o *CheckpointAttributes) SetIsPrimaryCheckpoint(v bool) {
	o.IsPrimaryCheckpoint = v
}

// GetSupportsIpv6 returns the SupportsIpv6 field value
func (o *CheckpointAttributes) GetSupportsIpv6() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsIpv6
}

// GetSupportsIpv6Ok returns a tuple with the SupportsIpv6 field value
// and a boolean to check if the value has been set.
func (o *CheckpointAttributes) GetSupportsIpv6Ok() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SupportsIpv6, true
}

// SetSupportsIpv6 sets field value
func (o *CheckpointAttributes) SetSupportsIpv6(v bool) {
	o.SupportsIpv6 = v
}

// GetHasHighAvailability returns the HasHighAvailability field value
func (o *CheckpointAttributes) GetHasHighAvailability() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasHighAvailability
}

// GetHasHighAvailabilityOk returns a tuple with the HasHighAvailability field value
// and a boolean to check if the value has been set.
func (o *CheckpointAttributes) GetHasHighAvailabilityOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HasHighAvailability, true
}

// SetHasHighAvailability sets field value
func (o *CheckpointAttributes) SetHasHighAvailability(v bool) {
	o.HasHighAvailability = v
}

func (o CheckpointAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CheckpointName) {
		toSerialize["CheckpointName"] = o.CheckpointName
	}
	if !isNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !isNil(o.Ipv4Addresses) {
		toSerialize["Ipv4Addresses"] = o.Ipv4Addresses
	}
	if !isNil(o.IpV6Addresses) {
		toSerialize["IpV6Addresses"] = o.IpV6Addresses
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

type NullableCheckpointAttributes struct {
	value *CheckpointAttributes
	isSet bool
}

func (v NullableCheckpointAttributes) Get() *CheckpointAttributes {
	return v.value
}

func (v *NullableCheckpointAttributes) Set(val *CheckpointAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckpointAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckpointAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckpointAttributes(val *CheckpointAttributes) *NullableCheckpointAttributes {
	return &NullableCheckpointAttributes{value: val, isSet: true}
}

func (v NullableCheckpointAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckpointAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


