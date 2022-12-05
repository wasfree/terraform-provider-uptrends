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

// Ipv6Address Ipv6 address
type Ipv6Address struct {
	// The IPv6 address
	IpAddress *string `json:"IpAddress,omitempty"`
}

// NewIpv6Address instantiates a new Ipv6Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6Address() *Ipv6Address {
	this := Ipv6Address{}
	return &this
}

// NewIpv6AddressWithDefaults instantiates a new Ipv6Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6AddressWithDefaults() *Ipv6Address {
	this := Ipv6Address{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *Ipv6Address) GetIpAddress() string {
	if o == nil || isNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6Address) GetIpAddressOk() (*string, bool) {
	if o == nil || isNil(o.IpAddress) {
    return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *Ipv6Address) HasIpAddress() bool {
	if o != nil && !isNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *Ipv6Address) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o Ipv6Address) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpAddress) {
		toSerialize["IpAddress"] = o.IpAddress
	}
	return json.Marshal(toSerialize)
}

type NullableIpv6Address struct {
	value *Ipv6Address
	isSet bool
}

func (v NullableIpv6Address) Get() *Ipv6Address {
	return v.value
}

func (v *NullableIpv6Address) Set(val *Ipv6Address) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6Address) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6Address) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6Address(val *Ipv6Address) *NullableIpv6Address {
	return &NullableIpv6Address{value: val, isSet: true}
}

func (v NullableIpv6Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6Address) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


