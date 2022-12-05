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

// PatternMatch struct for PatternMatch
type PatternMatch struct {
	Pattern *string `json:"Pattern,omitempty"`
	IsPositive bool `json:"IsPositive"`
	DateTime *DateTimePatternMatch `json:"DateTime,omitempty"`
}

// NewPatternMatch instantiates a new PatternMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatternMatch(isPositive bool) *PatternMatch {
	this := PatternMatch{}
	this.IsPositive = isPositive
	return &this
}

// NewPatternMatchWithDefaults instantiates a new PatternMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatternMatchWithDefaults() *PatternMatch {
	this := PatternMatch{}
	return &this
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *PatternMatch) GetPattern() string {
	if o == nil || isNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatternMatch) GetPatternOk() (*string, bool) {
	if o == nil || isNil(o.Pattern) {
    return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *PatternMatch) HasPattern() bool {
	if o != nil && !isNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *PatternMatch) SetPattern(v string) {
	o.Pattern = &v
}

// GetIsPositive returns the IsPositive field value
func (o *PatternMatch) GetIsPositive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPositive
}

// GetIsPositiveOk returns a tuple with the IsPositive field value
// and a boolean to check if the value has been set.
func (o *PatternMatch) GetIsPositiveOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsPositive, true
}

// SetIsPositive sets field value
func (o *PatternMatch) SetIsPositive(v bool) {
	o.IsPositive = v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *PatternMatch) GetDateTime() DateTimePatternMatch {
	if o == nil || isNil(o.DateTime) {
		var ret DateTimePatternMatch
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatternMatch) GetDateTimeOk() (*DateTimePatternMatch, bool) {
	if o == nil || isNil(o.DateTime) {
    return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *PatternMatch) HasDateTime() bool {
	if o != nil && !isNil(o.DateTime) {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given DateTimePatternMatch and assigns it to the DateTime field.
func (o *PatternMatch) SetDateTime(v DateTimePatternMatch) {
	o.DateTime = &v
}

func (o PatternMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Pattern) {
		toSerialize["Pattern"] = o.Pattern
	}
	if true {
		toSerialize["IsPositive"] = o.IsPositive
	}
	if !isNil(o.DateTime) {
		toSerialize["DateTime"] = o.DateTime
	}
	return json.Marshal(toSerialize)
}

type NullablePatternMatch struct {
	value *PatternMatch
	isSet bool
}

func (v NullablePatternMatch) Get() *PatternMatch {
	return v.value
}

func (v *NullablePatternMatch) Set(val *PatternMatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePatternMatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePatternMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatternMatch(val *PatternMatch) *NullablePatternMatch {
	return &NullablePatternMatch{value: val, isSet: true}
}

func (v NullablePatternMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatternMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


