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

// CheckpointListResponse struct for CheckpointListResponse
type CheckpointListResponse struct {
	// The resposne data/monitor checks
	Data []Checkpoint2 `json:"Data,omitempty"`
	Links *StatisticsResponseLinks `json:"Links,omitempty"`
	// Relationships of the object
	Relationships []RelationObject `json:"Relationships,omitempty"`
	Meta *StatisticsResponseMeta `json:"Meta,omitempty"`
}

// NewCheckpointListResponse instantiates a new CheckpointListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckpointListResponse() *CheckpointListResponse {
	this := CheckpointListResponse{}
	return &this
}

// NewCheckpointListResponseWithDefaults instantiates a new CheckpointListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckpointListResponseWithDefaults() *CheckpointListResponse {
	this := CheckpointListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckpointListResponse) GetData() []Checkpoint2 {
	if o == nil || isNil(o.Data) {
		var ret []Checkpoint2
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpointListResponse) GetDataOk() ([]Checkpoint2, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckpointListResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Checkpoint2 and assigns it to the Data field.
func (o *CheckpointListResponse) SetData(v []Checkpoint2) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CheckpointListResponse) GetLinks() StatisticsResponseLinks {
	if o == nil || isNil(o.Links) {
		var ret StatisticsResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpointListResponse) GetLinksOk() (*StatisticsResponseLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CheckpointListResponse) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given StatisticsResponseLinks and assigns it to the Links field.
func (o *CheckpointListResponse) SetLinks(v StatisticsResponseLinks) {
	o.Links = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CheckpointListResponse) GetRelationships() []RelationObject {
	if o == nil || isNil(o.Relationships) {
		var ret []RelationObject
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpointListResponse) GetRelationshipsOk() ([]RelationObject, bool) {
	if o == nil || isNil(o.Relationships) {
    return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CheckpointListResponse) HasRelationships() bool {
	if o != nil && !isNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []RelationObject and assigns it to the Relationships field.
func (o *CheckpointListResponse) SetRelationships(v []RelationObject) {
	o.Relationships = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CheckpointListResponse) GetMeta() StatisticsResponseMeta {
	if o == nil || isNil(o.Meta) {
		var ret StatisticsResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckpointListResponse) GetMetaOk() (*StatisticsResponseMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CheckpointListResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given StatisticsResponseMeta and assigns it to the Meta field.
func (o *CheckpointListResponse) SetMeta(v StatisticsResponseMeta) {
	o.Meta = &v
}

func (o CheckpointListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !isNil(o.Links) {
		toSerialize["Links"] = o.Links
	}
	if !isNil(o.Relationships) {
		toSerialize["Relationships"] = o.Relationships
	}
	if !isNil(o.Meta) {
		toSerialize["Meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableCheckpointListResponse struct {
	value *CheckpointListResponse
	isSet bool
}

func (v NullableCheckpointListResponse) Get() *CheckpointListResponse {
	return v.value
}

func (v *NullableCheckpointListResponse) Set(val *CheckpointListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckpointListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckpointListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckpointListResponse(val *CheckpointListResponse) *NullableCheckpointListResponse {
	return &NullableCheckpointListResponse{value: val, isSet: true}
}

func (v NullableCheckpointListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckpointListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


