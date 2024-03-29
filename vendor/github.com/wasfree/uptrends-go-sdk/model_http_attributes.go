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

// HttpAttributes Http details attributes
type HttpAttributes struct {
	// The HTML code retrieved from the target
	ResponseBody *string `json:"ResponseBody,omitempty"`
	// The HTTP response headers retrieved from the target 
	ResponseHeaders *string `json:"ResponseHeaders,omitempty"`
	// The URL of the HTTP Check. 
	Url *string `json:"Url,omitempty"`
}

// NewHttpAttributes instantiates a new HttpAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpAttributes() *HttpAttributes {
	this := HttpAttributes{}
	return &this
}

// NewHttpAttributesWithDefaults instantiates a new HttpAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpAttributesWithDefaults() *HttpAttributes {
	this := HttpAttributes{}
	return &this
}

// GetResponseBody returns the ResponseBody field value if set, zero value otherwise.
func (o *HttpAttributes) GetResponseBody() string {
	if o == nil || isNil(o.ResponseBody) {
		var ret string
		return ret
	}
	return *o.ResponseBody
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpAttributes) GetResponseBodyOk() (*string, bool) {
	if o == nil || isNil(o.ResponseBody) {
    return nil, false
	}
	return o.ResponseBody, true
}

// HasResponseBody returns a boolean if a field has been set.
func (o *HttpAttributes) HasResponseBody() bool {
	if o != nil && !isNil(o.ResponseBody) {
		return true
	}

	return false
}

// SetResponseBody gets a reference to the given string and assigns it to the ResponseBody field.
func (o *HttpAttributes) SetResponseBody(v string) {
	o.ResponseBody = &v
}

// GetResponseHeaders returns the ResponseHeaders field value if set, zero value otherwise.
func (o *HttpAttributes) GetResponseHeaders() string {
	if o == nil || isNil(o.ResponseHeaders) {
		var ret string
		return ret
	}
	return *o.ResponseHeaders
}

// GetResponseHeadersOk returns a tuple with the ResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpAttributes) GetResponseHeadersOk() (*string, bool) {
	if o == nil || isNil(o.ResponseHeaders) {
    return nil, false
	}
	return o.ResponseHeaders, true
}

// HasResponseHeaders returns a boolean if a field has been set.
func (o *HttpAttributes) HasResponseHeaders() bool {
	if o != nil && !isNil(o.ResponseHeaders) {
		return true
	}

	return false
}

// SetResponseHeaders gets a reference to the given string and assigns it to the ResponseHeaders field.
func (o *HttpAttributes) SetResponseHeaders(v string) {
	o.ResponseHeaders = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HttpAttributes) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpAttributes) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HttpAttributes) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HttpAttributes) SetUrl(v string) {
	o.Url = &v
}

func (o HttpAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResponseBody) {
		toSerialize["ResponseBody"] = o.ResponseBody
	}
	if !isNil(o.ResponseHeaders) {
		toSerialize["ResponseHeaders"] = o.ResponseHeaders
	}
	if !isNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableHttpAttributes struct {
	value *HttpAttributes
	isSet bool
}

func (v NullableHttpAttributes) Get() *HttpAttributes {
	return v.value
}

func (v *NullableHttpAttributes) Set(val *HttpAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpAttributes(val *HttpAttributes) *NullableHttpAttributes {
	return &NullableHttpAttributes{value: val, isSet: true}
}

func (v NullableHttpAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


