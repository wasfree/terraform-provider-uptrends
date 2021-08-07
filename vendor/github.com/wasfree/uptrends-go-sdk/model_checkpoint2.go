/*
 * Uptrends API v4
 *
 * This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"encoding/json"
)

// Checkpoint2 struct for Checkpoint2
type Checkpoint2 struct {
	Attributes *CheckpointAttributes `json:"Attributes,omitempty"`
	Id int32 `json:"Id"`
	Type *string `json:"Type,omitempty"`
	Relationships *[]RelationObject `json:"Relationships,omitempty"`
	Links *map[string]string `json:"Links,omitempty"`
}

// NewCheckpoint2 instantiates a new Checkpoint2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckpoint2(id int32) *Checkpoint2 {
	this := Checkpoint2{}
	this.Id = id
	return &this
}

// NewCheckpoint2WithDefaults instantiates a new Checkpoint2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckpoint2WithDefaults() *Checkpoint2 {
	this := Checkpoint2{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Checkpoint2) GetAttributes() CheckpointAttributes {
	if o == nil || o.Attributes == nil {
		var ret CheckpointAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint2) GetAttributesOk() (*CheckpointAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Checkpoint2) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CheckpointAttributes and assigns it to the Attributes field.
func (o *Checkpoint2) SetAttributes(v CheckpointAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value
func (o *Checkpoint2) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Checkpoint2) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Checkpoint2) SetId(v int32) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Checkpoint2) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint2) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Checkpoint2) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Checkpoint2) SetType(v string) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *Checkpoint2) GetRelationships() []RelationObject {
	if o == nil || o.Relationships == nil {
		var ret []RelationObject
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint2) GetRelationshipsOk() (*[]RelationObject, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *Checkpoint2) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []RelationObject and assigns it to the Relationships field.
func (o *Checkpoint2) SetRelationships(v []RelationObject) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Checkpoint2) GetLinks() map[string]string {
	if o == nil || o.Links == nil {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkpoint2) GetLinksOk() (*map[string]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Checkpoint2) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *Checkpoint2) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o Checkpoint2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["Attributes"] = o.Attributes
	}
	if true {
		toSerialize["Id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Relationships != nil {
		toSerialize["Relationships"] = o.Relationships
	}
	if o.Links != nil {
		toSerialize["Links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableCheckpoint2 struct {
	value *Checkpoint2
	isSet bool
}

func (v NullableCheckpoint2) Get() *Checkpoint2 {
	return v.value
}

func (v *NullableCheckpoint2) Set(val *Checkpoint2) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckpoint2) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckpoint2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckpoint2(val *Checkpoint2) *NullableCheckpoint2 {
	return &NullableCheckpoint2{value: val, isSet: true}
}

func (v NullableCheckpoint2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckpoint2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


