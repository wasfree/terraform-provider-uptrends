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

// PublicStatusPage struct for PublicStatusPage
type PublicStatusPage struct {
	PublicDashboardGuid *string `json:"PublicDashboardGuid,omitempty"`
	Name *string `json:"Name,omitempty"`
	IsPublished *bool `json:"IsPublished,omitempty"`
	PresetPeriodType *PresetPeriodType `json:"PresetPeriodType,omitempty"`
	CustomizationInfo *CustomizationInfo `json:"CustomizationInfo,omitempty"`
	SlaGuid *string `json:"SlaGuid,omitempty"`
	SlaGuidSpecified *bool `json:"SlaGuidSpecified,omitempty"`
	MonitorGuids []string `json:"MonitorGuids,omitempty"`
	MonitorGroupGuids []string `json:"MonitorGroupGuids,omitempty"`
}

// NewPublicStatusPage instantiates a new PublicStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicStatusPage() *PublicStatusPage {
	this := PublicStatusPage{}
	return &this
}

// NewPublicStatusPageWithDefaults instantiates a new PublicStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicStatusPageWithDefaults() *PublicStatusPage {
	this := PublicStatusPage{}
	return &this
}

// GetPublicDashboardGuid returns the PublicDashboardGuid field value if set, zero value otherwise.
func (o *PublicStatusPage) GetPublicDashboardGuid() string {
	if o == nil || isNil(o.PublicDashboardGuid) {
		var ret string
		return ret
	}
	return *o.PublicDashboardGuid
}

// GetPublicDashboardGuidOk returns a tuple with the PublicDashboardGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicStatusPage) GetPublicDashboardGuidOk() (*string, bool) {
	if o == nil || isNil(o.PublicDashboardGuid) {
    return nil, false
	}
	return o.PublicDashboardGuid, true
}

// HasPublicDashboardGuid returns a boolean if a field has been set.
func (o *PublicStatusPage) HasPublicDashboardGuid() bool {
	if o != nil && !isNil(o.PublicDashboardGuid) {
		return true
	}

	return false
}

// SetPublicDashboardGuid gets a reference to the given string and assigns it to the PublicDashboardGuid field.
func (o *PublicStatusPage) SetPublicDashboardGuid(v string) {
	o.PublicDashboardGuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublicStatusPage) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicStatusPage) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublicStatusPage) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublicStatusPage) SetName(v string) {
	o.Name = &v
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *PublicStatusPage) GetIsPublished() bool {
	if o == nil || isNil(o.IsPublished) {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicStatusPage) GetIsPublishedOk() (*bool, bool) {
	if o == nil || isNil(o.IsPublished) {
    return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *PublicStatusPage) HasIsPublished() bool {
	if o != nil && !isNil(o.IsPublished) {
		return true
	}

	return false
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *PublicStatusPage) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetPresetPeriodType returns the PresetPeriodType field value if set, zero value otherwise.
func (o *PublicStatusPage) GetPresetPeriodType() PresetPeriodType {
	if o == nil || isNil(o.PresetPeriodType) {
		var ret PresetPeriodType
		return ret
	}
	return *o.PresetPeriodType
}

// GetPresetPeriodTypeOk returns a tuple with the PresetPeriodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicStatusPage) GetPresetPeriodTypeOk() (*PresetPeriodType, bool) {
	if o == nil || isNil(o.PresetPeriodType) {
    return nil, false
	}
	return o.PresetPeriodType, true
}

// HasPresetPeriodType returns a boolean if a field has been set.
func (o *PublicStatusPage) HasPresetPeriodType() bool {
	if o != nil && !isNil(o.PresetPeriodType) {
		return true
	}

	return false
}

// SetPresetPeriodType gets a reference to the given PresetPeriodType and assigns it to the PresetPeriodType field.
func (o *PublicStatusPage) SetPresetPeriodType(v PresetPeriodType) {
	o.PresetPeriodType = &v
}

// GetCustomizationInfo returns the CustomizationInfo field value if set, zero value otherwise.
func (o *PublicStatusPage) GetCustomizationInfo() CustomizationInfo {
	if o == nil || isNil(o.CustomizationInfo) {
		var ret CustomizationInfo
		return ret
	}
	return *o.CustomizationInfo
}

// GetCustomizationInfoOk returns a tuple with the CustomizationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicStatusPage) GetCustomizationInfoOk() (*CustomizationInfo, bool) {
	if o == nil || isNil(o.CustomizationInfo) {
    return nil, false
	}
	return o.CustomizationInfo, true
}

// HasCustomizationInfo returns a boolean if a field has been set.
func (o *PublicStatusPage) HasCustomizationInfo() bool {
	if o != nil && !isNil(o.CustomizationInfo) {
		return true
	}

	return false
}

// SetCustomizationInfo gets a reference to the given CustomizationInfo and assigns it to the CustomizationInfo field.
func (o *PublicStatusPage) SetCustomizationInfo(v CustomizationInfo) {
	o.CustomizationInfo = &v
}

// GetSlaGuid returns the SlaGuid field value if set, zero value otherwise.
func (o *PublicStatusPage) GetSlaGuid() string {
	if o == nil || isNil(o.SlaGuid) {
		var ret string
		return ret
	}
	return *o.SlaGuid
}

// GetSlaGuidOk returns a tuple with the SlaGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicStatusPage) GetSlaGuidOk() (*string, bool) {
	if o == nil || isNil(o.SlaGuid) {
    return nil, false
	}
	return o.SlaGuid, true
}

// HasSlaGuid returns a boolean if a field has been set.
func (o *PublicStatusPage) HasSlaGuid() bool {
	if o != nil && !isNil(o.SlaGuid) {
		return true
	}

	return false
}

// SetSlaGuid gets a reference to the given string and assigns it to the SlaGuid field.
func (o *PublicStatusPage) SetSlaGuid(v string) {
	o.SlaGuid = &v
}

// GetSlaGuidSpecified returns the SlaGuidSpecified field value if set, zero value otherwise.
func (o *PublicStatusPage) GetSlaGuidSpecified() bool {
	if o == nil || isNil(o.SlaGuidSpecified) {
		var ret bool
		return ret
	}
	return *o.SlaGuidSpecified
}

// GetSlaGuidSpecifiedOk returns a tuple with the SlaGuidSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicStatusPage) GetSlaGuidSpecifiedOk() (*bool, bool) {
	if o == nil || isNil(o.SlaGuidSpecified) {
    return nil, false
	}
	return o.SlaGuidSpecified, true
}

// HasSlaGuidSpecified returns a boolean if a field has been set.
func (o *PublicStatusPage) HasSlaGuidSpecified() bool {
	if o != nil && !isNil(o.SlaGuidSpecified) {
		return true
	}

	return false
}

// SetSlaGuidSpecified gets a reference to the given bool and assigns it to the SlaGuidSpecified field.
func (o *PublicStatusPage) SetSlaGuidSpecified(v bool) {
	o.SlaGuidSpecified = &v
}

// GetMonitorGuids returns the MonitorGuids field value if set, zero value otherwise.
func (o *PublicStatusPage) GetMonitorGuids() []string {
	if o == nil || isNil(o.MonitorGuids) {
		var ret []string
		return ret
	}
	return o.MonitorGuids
}

// GetMonitorGuidsOk returns a tuple with the MonitorGuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicStatusPage) GetMonitorGuidsOk() ([]string, bool) {
	if o == nil || isNil(o.MonitorGuids) {
    return nil, false
	}
	return o.MonitorGuids, true
}

// HasMonitorGuids returns a boolean if a field has been set.
func (o *PublicStatusPage) HasMonitorGuids() bool {
	if o != nil && !isNil(o.MonitorGuids) {
		return true
	}

	return false
}

// SetMonitorGuids gets a reference to the given []string and assigns it to the MonitorGuids field.
func (o *PublicStatusPage) SetMonitorGuids(v []string) {
	o.MonitorGuids = v
}

// GetMonitorGroupGuids returns the MonitorGroupGuids field value if set, zero value otherwise.
func (o *PublicStatusPage) GetMonitorGroupGuids() []string {
	if o == nil || isNil(o.MonitorGroupGuids) {
		var ret []string
		return ret
	}
	return o.MonitorGroupGuids
}

// GetMonitorGroupGuidsOk returns a tuple with the MonitorGroupGuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicStatusPage) GetMonitorGroupGuidsOk() ([]string, bool) {
	if o == nil || isNil(o.MonitorGroupGuids) {
    return nil, false
	}
	return o.MonitorGroupGuids, true
}

// HasMonitorGroupGuids returns a boolean if a field has been set.
func (o *PublicStatusPage) HasMonitorGroupGuids() bool {
	if o != nil && !isNil(o.MonitorGroupGuids) {
		return true
	}

	return false
}

// SetMonitorGroupGuids gets a reference to the given []string and assigns it to the MonitorGroupGuids field.
func (o *PublicStatusPage) SetMonitorGroupGuids(v []string) {
	o.MonitorGroupGuids = v
}

func (o PublicStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PublicDashboardGuid) {
		toSerialize["PublicDashboardGuid"] = o.PublicDashboardGuid
	}
	if !isNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !isNil(o.IsPublished) {
		toSerialize["IsPublished"] = o.IsPublished
	}
	if !isNil(o.PresetPeriodType) {
		toSerialize["PresetPeriodType"] = o.PresetPeriodType
	}
	if !isNil(o.CustomizationInfo) {
		toSerialize["CustomizationInfo"] = o.CustomizationInfo
	}
	if !isNil(o.SlaGuid) {
		toSerialize["SlaGuid"] = o.SlaGuid
	}
	if !isNil(o.SlaGuidSpecified) {
		toSerialize["SlaGuidSpecified"] = o.SlaGuidSpecified
	}
	if !isNil(o.MonitorGuids) {
		toSerialize["MonitorGuids"] = o.MonitorGuids
	}
	if !isNil(o.MonitorGroupGuids) {
		toSerialize["MonitorGroupGuids"] = o.MonitorGroupGuids
	}
	return json.Marshal(toSerialize)
}

type NullablePublicStatusPage struct {
	value *PublicStatusPage
	isSet bool
}

func (v NullablePublicStatusPage) Get() *PublicStatusPage {
	return v.value
}

func (v *NullablePublicStatusPage) Set(val *PublicStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicStatusPage(val *PublicStatusPage) *NullablePublicStatusPage {
	return &NullablePublicStatusPage{value: val, isSet: true}
}

func (v NullablePublicStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


