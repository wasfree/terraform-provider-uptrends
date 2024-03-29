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

// StatisticsResponseMeta Meta data
type StatisticsResponseMeta struct {
	Timestamp map[string]interface{} `json:"Timestamp,omitempty"`
	Period *PeriodMetaData `json:"Period,omitempty"`
}

// NewStatisticsResponseMeta instantiates a new StatisticsResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticsResponseMeta() *StatisticsResponseMeta {
	this := StatisticsResponseMeta{}
	return &this
}

// NewStatisticsResponseMetaWithDefaults instantiates a new StatisticsResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticsResponseMetaWithDefaults() *StatisticsResponseMeta {
	this := StatisticsResponseMeta{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *StatisticsResponseMeta) GetTimestamp() map[string]interface{} {
	if o == nil || isNil(o.Timestamp) {
		var ret map[string]interface{}
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticsResponseMeta) GetTimestampOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Timestamp) {
    return map[string]interface{}{}, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *StatisticsResponseMeta) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given map[string]interface{} and assigns it to the Timestamp field.
func (o *StatisticsResponseMeta) SetTimestamp(v map[string]interface{}) {
	o.Timestamp = v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *StatisticsResponseMeta) GetPeriod() PeriodMetaData {
	if o == nil || isNil(o.Period) {
		var ret PeriodMetaData
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticsResponseMeta) GetPeriodOk() (*PeriodMetaData, bool) {
	if o == nil || isNil(o.Period) {
    return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *StatisticsResponseMeta) HasPeriod() bool {
	if o != nil && !isNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given PeriodMetaData and assigns it to the Period field.
func (o *StatisticsResponseMeta) SetPeriod(v PeriodMetaData) {
	o.Period = &v
}

func (o StatisticsResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timestamp) {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if !isNil(o.Period) {
		toSerialize["Period"] = o.Period
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticsResponseMeta struct {
	value *StatisticsResponseMeta
	isSet bool
}

func (v NullableStatisticsResponseMeta) Get() *StatisticsResponseMeta {
	return v.value
}

func (v *NullableStatisticsResponseMeta) Set(val *StatisticsResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticsResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticsResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticsResponseMeta(val *StatisticsResponseMeta) *NullableStatisticsResponseMeta {
	return &NullableStatisticsResponseMeta{value: val, isSet: true}
}

func (v NullableStatisticsResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticsResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


