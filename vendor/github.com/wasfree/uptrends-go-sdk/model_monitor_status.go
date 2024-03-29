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

// MonitorStatus struct for MonitorStatus
type MonitorStatus struct {
	Attributes *MonitorStatusAttributes `json:"Attributes,omitempty"`
	// Identifier 
	Id string `json:"Id"`
	// Object type
	Type *string `json:"Type,omitempty"`
	// Relationships of the object
	Relationships []RelationObject `json:"Relationships,omitempty"`
	// Links related to the object
	Links *map[string]string `json:"Links,omitempty"`
}

// NewMonitorStatus instantiates a new MonitorStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorStatus(id string) *MonitorStatus {
	this := MonitorStatus{}
	this.Id = id
	return &this
}

// NewMonitorStatusWithDefaults instantiates a new MonitorStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorStatusWithDefaults() *MonitorStatus {
	this := MonitorStatus{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MonitorStatus) GetAttributes() MonitorStatusAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret MonitorStatusAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetAttributesOk() (*MonitorStatusAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MonitorStatus) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MonitorStatusAttributes and assigns it to the Attributes field.
func (o *MonitorStatus) SetAttributes(v MonitorStatusAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value
func (o *MonitorStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MonitorStatus) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitorStatus) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitorStatus) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MonitorStatus) SetType(v string) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *MonitorStatus) GetRelationships() []RelationObject {
	if o == nil || isNil(o.Relationships) {
		var ret []RelationObject
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetRelationshipsOk() ([]RelationObject, bool) {
	if o == nil || isNil(o.Relationships) {
    return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *MonitorStatus) HasRelationships() bool {
	if o != nil && !isNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []RelationObject and assigns it to the Relationships field.
func (o *MonitorStatus) SetRelationships(v []RelationObject) {
	o.Relationships = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MonitorStatus) GetLinks() map[string]string {
	if o == nil || isNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStatus) GetLinksOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MonitorStatus) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *MonitorStatus) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o MonitorStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["Attributes"] = o.Attributes
	}
	if true {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !isNil(o.Relationships) {
		toSerialize["Relationships"] = o.Relationships
	}
	if !isNil(o.Links) {
		toSerialize["Links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorStatus struct {
	value *MonitorStatus
	isSet bool
}

func (v NullableMonitorStatus) Get() *MonitorStatus {
	return v.value
}

func (v *NullableMonitorStatus) Set(val *MonitorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorStatus(val *MonitorStatus) *NullableMonitorStatus {
	return &NullableMonitorStatus{value: val, isSet: true}
}

func (v NullableMonitorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


