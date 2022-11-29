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

// HttpEngineCheckDetailsAttributes Object attributes 
type HttpEngineCheckDetailsAttributes struct {
	// The results of the steps 
	StepResults []HttpEngineStep `json:"StepResults,omitempty"`
	TimingInfo *HttpEngineAttributesTimingInfo `json:"TimingInfo,omitempty"`
	// Number of total steps
	TotalSteps int32 `json:"TotalSteps"`
	// Number of passed/succeed tests
	PassedSteps int32 `json:"PassedSteps"`
}

// NewHttpEngineCheckDetailsAttributes instantiates a new HttpEngineCheckDetailsAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpEngineCheckDetailsAttributes(totalSteps int32, passedSteps int32) *HttpEngineCheckDetailsAttributes {
	this := HttpEngineCheckDetailsAttributes{}
	this.TotalSteps = totalSteps
	this.PassedSteps = passedSteps
	return &this
}

// NewHttpEngineCheckDetailsAttributesWithDefaults instantiates a new HttpEngineCheckDetailsAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpEngineCheckDetailsAttributesWithDefaults() *HttpEngineCheckDetailsAttributes {
	this := HttpEngineCheckDetailsAttributes{}
	return &this
}

// GetStepResults returns the StepResults field value if set, zero value otherwise.
func (o *HttpEngineCheckDetailsAttributes) GetStepResults() []HttpEngineStep {
	if o == nil || isNil(o.StepResults) {
		var ret []HttpEngineStep
		return ret
	}
	return o.StepResults
}

// GetStepResultsOk returns a tuple with the StepResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEngineCheckDetailsAttributes) GetStepResultsOk() ([]HttpEngineStep, bool) {
	if o == nil || isNil(o.StepResults) {
    return nil, false
	}
	return o.StepResults, true
}

// HasStepResults returns a boolean if a field has been set.
func (o *HttpEngineCheckDetailsAttributes) HasStepResults() bool {
	if o != nil && !isNil(o.StepResults) {
		return true
	}

	return false
}

// SetStepResults gets a reference to the given []HttpEngineStep and assigns it to the StepResults field.
func (o *HttpEngineCheckDetailsAttributes) SetStepResults(v []HttpEngineStep) {
	o.StepResults = v
}

// GetTimingInfo returns the TimingInfo field value if set, zero value otherwise.
func (o *HttpEngineCheckDetailsAttributes) GetTimingInfo() HttpEngineAttributesTimingInfo {
	if o == nil || isNil(o.TimingInfo) {
		var ret HttpEngineAttributesTimingInfo
		return ret
	}
	return *o.TimingInfo
}

// GetTimingInfoOk returns a tuple with the TimingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEngineCheckDetailsAttributes) GetTimingInfoOk() (*HttpEngineAttributesTimingInfo, bool) {
	if o == nil || isNil(o.TimingInfo) {
    return nil, false
	}
	return o.TimingInfo, true
}

// HasTimingInfo returns a boolean if a field has been set.
func (o *HttpEngineCheckDetailsAttributes) HasTimingInfo() bool {
	if o != nil && !isNil(o.TimingInfo) {
		return true
	}

	return false
}

// SetTimingInfo gets a reference to the given HttpEngineAttributesTimingInfo and assigns it to the TimingInfo field.
func (o *HttpEngineCheckDetailsAttributes) SetTimingInfo(v HttpEngineAttributesTimingInfo) {
	o.TimingInfo = &v
}

// GetTotalSteps returns the TotalSteps field value
func (o *HttpEngineCheckDetailsAttributes) GetTotalSteps() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalSteps
}

// GetTotalStepsOk returns a tuple with the TotalSteps field value
// and a boolean to check if the value has been set.
func (o *HttpEngineCheckDetailsAttributes) GetTotalStepsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalSteps, true
}

// SetTotalSteps sets field value
func (o *HttpEngineCheckDetailsAttributes) SetTotalSteps(v int32) {
	o.TotalSteps = v
}

// GetPassedSteps returns the PassedSteps field value
func (o *HttpEngineCheckDetailsAttributes) GetPassedSteps() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PassedSteps
}

// GetPassedStepsOk returns a tuple with the PassedSteps field value
// and a boolean to check if the value has been set.
func (o *HttpEngineCheckDetailsAttributes) GetPassedStepsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PassedSteps, true
}

// SetPassedSteps sets field value
func (o *HttpEngineCheckDetailsAttributes) SetPassedSteps(v int32) {
	o.PassedSteps = v
}

func (o HttpEngineCheckDetailsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StepResults) {
		toSerialize["StepResults"] = o.StepResults
	}
	if !isNil(o.TimingInfo) {
		toSerialize["TimingInfo"] = o.TimingInfo
	}
	if true {
		toSerialize["TotalSteps"] = o.TotalSteps
	}
	if true {
		toSerialize["PassedSteps"] = o.PassedSteps
	}
	return json.Marshal(toSerialize)
}

type NullableHttpEngineCheckDetailsAttributes struct {
	value *HttpEngineCheckDetailsAttributes
	isSet bool
}

func (v NullableHttpEngineCheckDetailsAttributes) Get() *HttpEngineCheckDetailsAttributes {
	return v.value
}

func (v *NullableHttpEngineCheckDetailsAttributes) Set(val *HttpEngineCheckDetailsAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpEngineCheckDetailsAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpEngineCheckDetailsAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpEngineCheckDetailsAttributes(val *HttpEngineCheckDetailsAttributes) *NullableHttpEngineCheckDetailsAttributes {
	return &NullableHttpEngineCheckDetailsAttributes{value: val, isSet: true}
}

func (v NullableHttpEngineCheckDetailsAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpEngineCheckDetailsAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


