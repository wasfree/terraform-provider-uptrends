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

// EscalationLevelIntegration struct for EscalationLevelIntegration
type EscalationLevelIntegration struct {
	// The unique key of this Integration.
	IntegrationGuid *string `json:"IntegrationGuid,omitempty"`
	// Extra email addresses
	ExtraEmailAddresses []string `json:"ExtraEmailAddresses,omitempty"`
	// Specified Extra EmailAddresses For Patch request
	ExtraEmailAddressesSpecified *bool `json:"ExtraEmailAddressesSpecified,omitempty"`
	// StatusHub Service Mapping
	StatusHubServiceList []IntegrationServiceMap `json:"StatusHubServiceList,omitempty"`
	// Specified StatusHubServiceList For Patch request
	StatusHubServiceListSpecified *bool `json:"StatusHubServiceListSpecified,omitempty"`
	// Indicates whether this Integration is active.
	IsActive *bool `json:"IsActive,omitempty"`
	Hash *string `json:"Hash,omitempty"`
}

// NewEscalationLevelIntegration instantiates a new EscalationLevelIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEscalationLevelIntegration() *EscalationLevelIntegration {
	this := EscalationLevelIntegration{}
	return &this
}

// NewEscalationLevelIntegrationWithDefaults instantiates a new EscalationLevelIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEscalationLevelIntegrationWithDefaults() *EscalationLevelIntegration {
	this := EscalationLevelIntegration{}
	return &this
}

// GetIntegrationGuid returns the IntegrationGuid field value if set, zero value otherwise.
func (o *EscalationLevelIntegration) GetIntegrationGuid() string {
	if o == nil || isNil(o.IntegrationGuid) {
		var ret string
		return ret
	}
	return *o.IntegrationGuid
}

// GetIntegrationGuidOk returns a tuple with the IntegrationGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevelIntegration) GetIntegrationGuidOk() (*string, bool) {
	if o == nil || isNil(o.IntegrationGuid) {
    return nil, false
	}
	return o.IntegrationGuid, true
}

// HasIntegrationGuid returns a boolean if a field has been set.
func (o *EscalationLevelIntegration) HasIntegrationGuid() bool {
	if o != nil && !isNil(o.IntegrationGuid) {
		return true
	}

	return false
}

// SetIntegrationGuid gets a reference to the given string and assigns it to the IntegrationGuid field.
func (o *EscalationLevelIntegration) SetIntegrationGuid(v string) {
	o.IntegrationGuid = &v
}

// GetExtraEmailAddresses returns the ExtraEmailAddresses field value if set, zero value otherwise.
func (o *EscalationLevelIntegration) GetExtraEmailAddresses() []string {
	if o == nil || isNil(o.ExtraEmailAddresses) {
		var ret []string
		return ret
	}
	return o.ExtraEmailAddresses
}

// GetExtraEmailAddressesOk returns a tuple with the ExtraEmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevelIntegration) GetExtraEmailAddressesOk() ([]string, bool) {
	if o == nil || isNil(o.ExtraEmailAddresses) {
    return nil, false
	}
	return o.ExtraEmailAddresses, true
}

// HasExtraEmailAddresses returns a boolean if a field has been set.
func (o *EscalationLevelIntegration) HasExtraEmailAddresses() bool {
	if o != nil && !isNil(o.ExtraEmailAddresses) {
		return true
	}

	return false
}

// SetExtraEmailAddresses gets a reference to the given []string and assigns it to the ExtraEmailAddresses field.
func (o *EscalationLevelIntegration) SetExtraEmailAddresses(v []string) {
	o.ExtraEmailAddresses = v
}

// GetExtraEmailAddressesSpecified returns the ExtraEmailAddressesSpecified field value if set, zero value otherwise.
func (o *EscalationLevelIntegration) GetExtraEmailAddressesSpecified() bool {
	if o == nil || isNil(o.ExtraEmailAddressesSpecified) {
		var ret bool
		return ret
	}
	return *o.ExtraEmailAddressesSpecified
}

// GetExtraEmailAddressesSpecifiedOk returns a tuple with the ExtraEmailAddressesSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevelIntegration) GetExtraEmailAddressesSpecifiedOk() (*bool, bool) {
	if o == nil || isNil(o.ExtraEmailAddressesSpecified) {
    return nil, false
	}
	return o.ExtraEmailAddressesSpecified, true
}

// HasExtraEmailAddressesSpecified returns a boolean if a field has been set.
func (o *EscalationLevelIntegration) HasExtraEmailAddressesSpecified() bool {
	if o != nil && !isNil(o.ExtraEmailAddressesSpecified) {
		return true
	}

	return false
}

// SetExtraEmailAddressesSpecified gets a reference to the given bool and assigns it to the ExtraEmailAddressesSpecified field.
func (o *EscalationLevelIntegration) SetExtraEmailAddressesSpecified(v bool) {
	o.ExtraEmailAddressesSpecified = &v
}

// GetStatusHubServiceList returns the StatusHubServiceList field value if set, zero value otherwise.
func (o *EscalationLevelIntegration) GetStatusHubServiceList() []IntegrationServiceMap {
	if o == nil || isNil(o.StatusHubServiceList) {
		var ret []IntegrationServiceMap
		return ret
	}
	return o.StatusHubServiceList
}

// GetStatusHubServiceListOk returns a tuple with the StatusHubServiceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevelIntegration) GetStatusHubServiceListOk() ([]IntegrationServiceMap, bool) {
	if o == nil || isNil(o.StatusHubServiceList) {
    return nil, false
	}
	return o.StatusHubServiceList, true
}

// HasStatusHubServiceList returns a boolean if a field has been set.
func (o *EscalationLevelIntegration) HasStatusHubServiceList() bool {
	if o != nil && !isNil(o.StatusHubServiceList) {
		return true
	}

	return false
}

// SetStatusHubServiceList gets a reference to the given []IntegrationServiceMap and assigns it to the StatusHubServiceList field.
func (o *EscalationLevelIntegration) SetStatusHubServiceList(v []IntegrationServiceMap) {
	o.StatusHubServiceList = v
}

// GetStatusHubServiceListSpecified returns the StatusHubServiceListSpecified field value if set, zero value otherwise.
func (o *EscalationLevelIntegration) GetStatusHubServiceListSpecified() bool {
	if o == nil || isNil(o.StatusHubServiceListSpecified) {
		var ret bool
		return ret
	}
	return *o.StatusHubServiceListSpecified
}

// GetStatusHubServiceListSpecifiedOk returns a tuple with the StatusHubServiceListSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevelIntegration) GetStatusHubServiceListSpecifiedOk() (*bool, bool) {
	if o == nil || isNil(o.StatusHubServiceListSpecified) {
    return nil, false
	}
	return o.StatusHubServiceListSpecified, true
}

// HasStatusHubServiceListSpecified returns a boolean if a field has been set.
func (o *EscalationLevelIntegration) HasStatusHubServiceListSpecified() bool {
	if o != nil && !isNil(o.StatusHubServiceListSpecified) {
		return true
	}

	return false
}

// SetStatusHubServiceListSpecified gets a reference to the given bool and assigns it to the StatusHubServiceListSpecified field.
func (o *EscalationLevelIntegration) SetStatusHubServiceListSpecified(v bool) {
	o.StatusHubServiceListSpecified = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *EscalationLevelIntegration) GetIsActive() bool {
	if o == nil || isNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevelIntegration) GetIsActiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsActive) {
    return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *EscalationLevelIntegration) HasIsActive() bool {
	if o != nil && !isNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *EscalationLevelIntegration) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *EscalationLevelIntegration) GetHash() string {
	if o == nil || isNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevelIntegration) GetHashOk() (*string, bool) {
	if o == nil || isNil(o.Hash) {
    return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *EscalationLevelIntegration) HasHash() bool {
	if o != nil && !isNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *EscalationLevelIntegration) SetHash(v string) {
	o.Hash = &v
}

func (o EscalationLevelIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IntegrationGuid) {
		toSerialize["IntegrationGuid"] = o.IntegrationGuid
	}
	if !isNil(o.ExtraEmailAddresses) {
		toSerialize["ExtraEmailAddresses"] = o.ExtraEmailAddresses
	}
	if !isNil(o.ExtraEmailAddressesSpecified) {
		toSerialize["ExtraEmailAddressesSpecified"] = o.ExtraEmailAddressesSpecified
	}
	if !isNil(o.StatusHubServiceList) {
		toSerialize["StatusHubServiceList"] = o.StatusHubServiceList
	}
	if !isNil(o.StatusHubServiceListSpecified) {
		toSerialize["StatusHubServiceListSpecified"] = o.StatusHubServiceListSpecified
	}
	if !isNil(o.IsActive) {
		toSerialize["IsActive"] = o.IsActive
	}
	if !isNil(o.Hash) {
		toSerialize["Hash"] = o.Hash
	}
	return json.Marshal(toSerialize)
}

type NullableEscalationLevelIntegration struct {
	value *EscalationLevelIntegration
	isSet bool
}

func (v NullableEscalationLevelIntegration) Get() *EscalationLevelIntegration {
	return v.value
}

func (v *NullableEscalationLevelIntegration) Set(val *EscalationLevelIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableEscalationLevelIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableEscalationLevelIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEscalationLevelIntegration(val *EscalationLevelIntegration) *NullableEscalationLevelIntegration {
	return &NullableEscalationLevelIntegration{value: val, isSet: true}
}

func (v NullableEscalationLevelIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEscalationLevelIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


