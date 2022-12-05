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

// TransactionStepOResourceObject struct for TransactionStepOResourceObject
type TransactionStepOResourceObject struct {
	Attributes *TransactionStepOResourceObjectAttributes `json:"Attributes,omitempty"`
	// Identifier 
	Id int32 `json:"Id"`
	// Object type
	Type *string `json:"Type,omitempty"`
	// Relationships of the object
	Relationships []RelationObject `json:"Relationships,omitempty"`
	// Links related to the object
	Links *map[string]string `json:"Links,omitempty"`
}

// NewTransactionStepOResourceObject instantiates a new TransactionStepOResourceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionStepOResourceObject(id int32) *TransactionStepOResourceObject {
	this := TransactionStepOResourceObject{}
	this.Id = id
	return &this
}

// NewTransactionStepOResourceObjectWithDefaults instantiates a new TransactionStepOResourceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionStepOResourceObjectWithDefaults() *TransactionStepOResourceObject {
	this := TransactionStepOResourceObject{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TransactionStepOResourceObject) GetAttributes() TransactionStepOResourceObjectAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret TransactionStepOResourceObjectAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStepOResourceObject) GetAttributesOk() (*TransactionStepOResourceObjectAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TransactionStepOResourceObject) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TransactionStepOResourceObjectAttributes and assigns it to the Attributes field.
func (o *TransactionStepOResourceObject) SetAttributes(v TransactionStepOResourceObjectAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value
func (o *TransactionStepOResourceObject) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionStepOResourceObject) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionStepOResourceObject) SetId(v int32) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransactionStepOResourceObject) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStepOResourceObject) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransactionStepOResourceObject) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransactionStepOResourceObject) SetType(v string) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TransactionStepOResourceObject) GetRelationships() []RelationObject {
	if o == nil || isNil(o.Relationships) {
		var ret []RelationObject
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStepOResourceObject) GetRelationshipsOk() ([]RelationObject, bool) {
	if o == nil || isNil(o.Relationships) {
    return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TransactionStepOResourceObject) HasRelationships() bool {
	if o != nil && !isNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []RelationObject and assigns it to the Relationships field.
func (o *TransactionStepOResourceObject) SetRelationships(v []RelationObject) {
	o.Relationships = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TransactionStepOResourceObject) GetLinks() map[string]string {
	if o == nil || isNil(o.Links) {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionStepOResourceObject) GetLinksOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TransactionStepOResourceObject) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *TransactionStepOResourceObject) SetLinks(v map[string]string) {
	o.Links = &v
}

func (o TransactionStepOResourceObject) MarshalJSON() ([]byte, error) {
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

type NullableTransactionStepOResourceObject struct {
	value *TransactionStepOResourceObject
	isSet bool
}

func (v NullableTransactionStepOResourceObject) Get() *TransactionStepOResourceObject {
	return v.value
}

func (v *NullableTransactionStepOResourceObject) Set(val *TransactionStepOResourceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStepOResourceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStepOResourceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStepOResourceObject(val *TransactionStepOResourceObject) *NullableTransactionStepOResourceObject {
	return &NullableTransactionStepOResourceObject{value: val, isSet: true}
}

func (v NullableTransactionStepOResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStepOResourceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


