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

// AlertDefinitionAuthorization Alert definition authorization
type AlertDefinitionAuthorization struct {
	// The unique ID of this authorization
	AuthorizationId *string `json:"AuthorizationId,omitempty"`
	// The authorization type
	AuthorizationType AlertDefinitionAuthorizationType `json:"AuthorizationType"`
	// The GUID of the operator (NULL if this authorization is for an operator group)
	OperatorGuid *string `json:"OperatorGuid,omitempty"`
	// The GUID of the operator group (NULL if this authorization is for an operator)
	OperatorGroupGuid *string `json:"OperatorGroupGuid,omitempty"`
}

// NewAlertDefinitionAuthorization instantiates a new AlertDefinitionAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertDefinitionAuthorization(authorizationType AlertDefinitionAuthorizationType) *AlertDefinitionAuthorization {
	this := AlertDefinitionAuthorization{}
	this.AuthorizationType = authorizationType
	return &this
}

// NewAlertDefinitionAuthorizationWithDefaults instantiates a new AlertDefinitionAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertDefinitionAuthorizationWithDefaults() *AlertDefinitionAuthorization {
	this := AlertDefinitionAuthorization{}
	return &this
}

// GetAuthorizationId returns the AuthorizationId field value if set, zero value otherwise.
func (o *AlertDefinitionAuthorization) GetAuthorizationId() string {
	if o == nil || isNil(o.AuthorizationId) {
		var ret string
		return ret
	}
	return *o.AuthorizationId
}

// GetAuthorizationIdOk returns a tuple with the AuthorizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionAuthorization) GetAuthorizationIdOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationId) {
    return nil, false
	}
	return o.AuthorizationId, true
}

// HasAuthorizationId returns a boolean if a field has been set.
func (o *AlertDefinitionAuthorization) HasAuthorizationId() bool {
	if o != nil && !isNil(o.AuthorizationId) {
		return true
	}

	return false
}

// SetAuthorizationId gets a reference to the given string and assigns it to the AuthorizationId field.
func (o *AlertDefinitionAuthorization) SetAuthorizationId(v string) {
	o.AuthorizationId = &v
}

// GetAuthorizationType returns the AuthorizationType field value
func (o *AlertDefinitionAuthorization) GetAuthorizationType() AlertDefinitionAuthorizationType {
	if o == nil {
		var ret AlertDefinitionAuthorizationType
		return ret
	}

	return o.AuthorizationType
}

// GetAuthorizationTypeOk returns a tuple with the AuthorizationType field value
// and a boolean to check if the value has been set.
func (o *AlertDefinitionAuthorization) GetAuthorizationTypeOk() (*AlertDefinitionAuthorizationType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthorizationType, true
}

// SetAuthorizationType sets field value
func (o *AlertDefinitionAuthorization) SetAuthorizationType(v AlertDefinitionAuthorizationType) {
	o.AuthorizationType = v
}

// GetOperatorGuid returns the OperatorGuid field value if set, zero value otherwise.
func (o *AlertDefinitionAuthorization) GetOperatorGuid() string {
	if o == nil || isNil(o.OperatorGuid) {
		var ret string
		return ret
	}
	return *o.OperatorGuid
}

// GetOperatorGuidOk returns a tuple with the OperatorGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionAuthorization) GetOperatorGuidOk() (*string, bool) {
	if o == nil || isNil(o.OperatorGuid) {
    return nil, false
	}
	return o.OperatorGuid, true
}

// HasOperatorGuid returns a boolean if a field has been set.
func (o *AlertDefinitionAuthorization) HasOperatorGuid() bool {
	if o != nil && !isNil(o.OperatorGuid) {
		return true
	}

	return false
}

// SetOperatorGuid gets a reference to the given string and assigns it to the OperatorGuid field.
func (o *AlertDefinitionAuthorization) SetOperatorGuid(v string) {
	o.OperatorGuid = &v
}

// GetOperatorGroupGuid returns the OperatorGroupGuid field value if set, zero value otherwise.
func (o *AlertDefinitionAuthorization) GetOperatorGroupGuid() string {
	if o == nil || isNil(o.OperatorGroupGuid) {
		var ret string
		return ret
	}
	return *o.OperatorGroupGuid
}

// GetOperatorGroupGuidOk returns a tuple with the OperatorGroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertDefinitionAuthorization) GetOperatorGroupGuidOk() (*string, bool) {
	if o == nil || isNil(o.OperatorGroupGuid) {
    return nil, false
	}
	return o.OperatorGroupGuid, true
}

// HasOperatorGroupGuid returns a boolean if a field has been set.
func (o *AlertDefinitionAuthorization) HasOperatorGroupGuid() bool {
	if o != nil && !isNil(o.OperatorGroupGuid) {
		return true
	}

	return false
}

// SetOperatorGroupGuid gets a reference to the given string and assigns it to the OperatorGroupGuid field.
func (o *AlertDefinitionAuthorization) SetOperatorGroupGuid(v string) {
	o.OperatorGroupGuid = &v
}

func (o AlertDefinitionAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthorizationId) {
		toSerialize["AuthorizationId"] = o.AuthorizationId
	}
	if true {
		toSerialize["AuthorizationType"] = o.AuthorizationType
	}
	if !isNil(o.OperatorGuid) {
		toSerialize["OperatorGuid"] = o.OperatorGuid
	}
	if !isNil(o.OperatorGroupGuid) {
		toSerialize["OperatorGroupGuid"] = o.OperatorGroupGuid
	}
	return json.Marshal(toSerialize)
}

type NullableAlertDefinitionAuthorization struct {
	value *AlertDefinitionAuthorization
	isSet bool
}

func (v NullableAlertDefinitionAuthorization) Get() *AlertDefinitionAuthorization {
	return v.value
}

func (v *NullableAlertDefinitionAuthorization) Set(val *AlertDefinitionAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertDefinitionAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertDefinitionAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertDefinitionAuthorization(val *AlertDefinitionAuthorization) *NullableAlertDefinitionAuthorization {
	return &NullableAlertDefinitionAuthorization{value: val, isSet: true}
}

func (v NullableAlertDefinitionAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertDefinitionAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

