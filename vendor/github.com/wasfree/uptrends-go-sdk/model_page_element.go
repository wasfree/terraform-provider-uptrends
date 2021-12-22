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

// PageElement Page element describes a page load/request (in a waterfall)
type PageElement struct {
	// Index of the element in the waterfall context
	Index int32 `json:"Index"`
	// Starting time in milliseconds
	StartTime int32 `json:"StartTime"`
	// Number of milliseconds this element has been queueing, when appropriate.
	QueueTime int32 `json:"QueueTime"`
	// Number of milliseconds needed to perform the DNS query for this element, when appropriate.
	ResolveTime int32 `json:"ResolveTime"`
	// Number of milliseconds needed to establish a connection.
	ConnectTime int32 `json:"ConnectTime"`
	// Number of milliseconds the connection was stale
	StaleTime int32 `json:"StaleTime"`
	// Number of milliseconds needed for the HTTPS handshake
	HttpsHandshakeTime int32 `json:"HttpsHandshakeTime"`
	// Number of milliseconds it took to send data
	SendTime int32 `json:"SendTime"`
	// Number of milliseconds the connection was in waiting state
	WaitTime int32 `json:"WaitTime"`
	// Number of milliseconds it took to retrieve the data
	ReceiveTime int32 `json:"ReceiveTime"`
	// Number of milliseconds the connection was timed-out 
	TimeoutTime int32 `json:"TimeoutTime"`
	// Total number of milliseconds it took for the connection to complete
	TotalTime int32 `json:"TotalTime"`
	// The Http status code
	HttpStatusCode int32 `json:"HttpStatusCode"`
	// The requested resource url
	Url *string `json:"Url,omitempty"`
	// Total number of downloaded bytes
	TotalBytes int32 `json:"TotalBytes"`
	// Requested resource element type, can be HTML, scripts, CSS, images, frames, objects, data and other
	ElementType *string `json:"ElementType,omitempty"`
	// The HTTP request headers used
	RequestHeaders *string `json:"RequestHeaders,omitempty"`
	// The HTTP response headers retrieved
	ResponseHeaders *string `json:"ResponseHeaders,omitempty"`
	// The IP address that was found for the specified domain name as part of this monitor check.
	ResolvedIpAddress *map[string]interface{} `json:"ResolvedIpAddress,omitempty"`
	GroupIds *[]int32 `json:"GroupIds,omitempty"`
	// Was the Url excluded from waterfall (timing) data by the user?
	UrlIsBlocked bool `json:"UrlIsBlocked"`
}

// NewPageElement instantiates a new PageElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageElement(index int32, startTime int32, queueTime int32, resolveTime int32, connectTime int32, staleTime int32, httpsHandshakeTime int32, sendTime int32, waitTime int32, receiveTime int32, timeoutTime int32, totalTime int32, httpStatusCode int32, totalBytes int32, urlIsBlocked bool) *PageElement {
	this := PageElement{}
	this.Index = index
	this.StartTime = startTime
	this.QueueTime = queueTime
	this.ResolveTime = resolveTime
	this.ConnectTime = connectTime
	this.StaleTime = staleTime
	this.HttpsHandshakeTime = httpsHandshakeTime
	this.SendTime = sendTime
	this.WaitTime = waitTime
	this.ReceiveTime = receiveTime
	this.TimeoutTime = timeoutTime
	this.TotalTime = totalTime
	this.HttpStatusCode = httpStatusCode
	this.TotalBytes = totalBytes
	this.UrlIsBlocked = urlIsBlocked
	return &this
}

// NewPageElementWithDefaults instantiates a new PageElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageElementWithDefaults() *PageElement {
	this := PageElement{}
	return &this
}

// GetIndex returns the Index field value
func (o *PageElement) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *PageElement) SetIndex(v int32) {
	o.Index = v
}

// GetStartTime returns the StartTime field value
func (o *PageElement) GetStartTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetStartTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *PageElement) SetStartTime(v int32) {
	o.StartTime = v
}

// GetQueueTime returns the QueueTime field value
func (o *PageElement) GetQueueTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QueueTime
}

// GetQueueTimeOk returns a tuple with the QueueTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetQueueTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QueueTime, true
}

// SetQueueTime sets field value
func (o *PageElement) SetQueueTime(v int32) {
	o.QueueTime = v
}

// GetResolveTime returns the ResolveTime field value
func (o *PageElement) GetResolveTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResolveTime
}

// GetResolveTimeOk returns a tuple with the ResolveTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetResolveTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResolveTime, true
}

// SetResolveTime sets field value
func (o *PageElement) SetResolveTime(v int32) {
	o.ResolveTime = v
}

// GetConnectTime returns the ConnectTime field value
func (o *PageElement) GetConnectTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConnectTime
}

// GetConnectTimeOk returns a tuple with the ConnectTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetConnectTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectTime, true
}

// SetConnectTime sets field value
func (o *PageElement) SetConnectTime(v int32) {
	o.ConnectTime = v
}

// GetStaleTime returns the StaleTime field value
func (o *PageElement) GetStaleTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StaleTime
}

// GetStaleTimeOk returns a tuple with the StaleTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetStaleTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StaleTime, true
}

// SetStaleTime sets field value
func (o *PageElement) SetStaleTime(v int32) {
	o.StaleTime = v
}

// GetHttpsHandshakeTime returns the HttpsHandshakeTime field value
func (o *PageElement) GetHttpsHandshakeTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HttpsHandshakeTime
}

// GetHttpsHandshakeTimeOk returns a tuple with the HttpsHandshakeTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetHttpsHandshakeTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HttpsHandshakeTime, true
}

// SetHttpsHandshakeTime sets field value
func (o *PageElement) SetHttpsHandshakeTime(v int32) {
	o.HttpsHandshakeTime = v
}

// GetSendTime returns the SendTime field value
func (o *PageElement) GetSendTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SendTime
}

// GetSendTimeOk returns a tuple with the SendTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetSendTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SendTime, true
}

// SetSendTime sets field value
func (o *PageElement) SetSendTime(v int32) {
	o.SendTime = v
}

// GetWaitTime returns the WaitTime field value
func (o *PageElement) GetWaitTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WaitTime
}

// GetWaitTimeOk returns a tuple with the WaitTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetWaitTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WaitTime, true
}

// SetWaitTime sets field value
func (o *PageElement) SetWaitTime(v int32) {
	o.WaitTime = v
}

// GetReceiveTime returns the ReceiveTime field value
func (o *PageElement) GetReceiveTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReceiveTime
}

// GetReceiveTimeOk returns a tuple with the ReceiveTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetReceiveTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReceiveTime, true
}

// SetReceiveTime sets field value
func (o *PageElement) SetReceiveTime(v int32) {
	o.ReceiveTime = v
}

// GetTimeoutTime returns the TimeoutTime field value
func (o *PageElement) GetTimeoutTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeoutTime
}

// GetTimeoutTimeOk returns a tuple with the TimeoutTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetTimeoutTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeoutTime, true
}

// SetTimeoutTime sets field value
func (o *PageElement) SetTimeoutTime(v int32) {
	o.TimeoutTime = v
}

// GetTotalTime returns the TotalTime field value
func (o *PageElement) GetTotalTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetTotalTimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalTime, true
}

// SetTotalTime sets field value
func (o *PageElement) SetTotalTime(v int32) {
	o.TotalTime = v
}

// GetHttpStatusCode returns the HttpStatusCode field value
func (o *PageElement) GetHttpStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetHttpStatusCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HttpStatusCode, true
}

// SetHttpStatusCode sets field value
func (o *PageElement) SetHttpStatusCode(v int32) {
	o.HttpStatusCode = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PageElement) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageElement) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PageElement) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PageElement) SetUrl(v string) {
	o.Url = &v
}

// GetTotalBytes returns the TotalBytes field value
func (o *PageElement) GetTotalBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetTotalBytesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalBytes, true
}

// SetTotalBytes sets field value
func (o *PageElement) SetTotalBytes(v int32) {
	o.TotalBytes = v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *PageElement) GetElementType() string {
	if o == nil || o.ElementType == nil {
		var ret string
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageElement) GetElementTypeOk() (*string, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *PageElement) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given string and assigns it to the ElementType field.
func (o *PageElement) SetElementType(v string) {
	o.ElementType = &v
}

// GetRequestHeaders returns the RequestHeaders field value if set, zero value otherwise.
func (o *PageElement) GetRequestHeaders() string {
	if o == nil || o.RequestHeaders == nil {
		var ret string
		return ret
	}
	return *o.RequestHeaders
}

// GetRequestHeadersOk returns a tuple with the RequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageElement) GetRequestHeadersOk() (*string, bool) {
	if o == nil || o.RequestHeaders == nil {
		return nil, false
	}
	return o.RequestHeaders, true
}

// HasRequestHeaders returns a boolean if a field has been set.
func (o *PageElement) HasRequestHeaders() bool {
	if o != nil && o.RequestHeaders != nil {
		return true
	}

	return false
}

// SetRequestHeaders gets a reference to the given string and assigns it to the RequestHeaders field.
func (o *PageElement) SetRequestHeaders(v string) {
	o.RequestHeaders = &v
}

// GetResponseHeaders returns the ResponseHeaders field value if set, zero value otherwise.
func (o *PageElement) GetResponseHeaders() string {
	if o == nil || o.ResponseHeaders == nil {
		var ret string
		return ret
	}
	return *o.ResponseHeaders
}

// GetResponseHeadersOk returns a tuple with the ResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageElement) GetResponseHeadersOk() (*string, bool) {
	if o == nil || o.ResponseHeaders == nil {
		return nil, false
	}
	return o.ResponseHeaders, true
}

// HasResponseHeaders returns a boolean if a field has been set.
func (o *PageElement) HasResponseHeaders() bool {
	if o != nil && o.ResponseHeaders != nil {
		return true
	}

	return false
}

// SetResponseHeaders gets a reference to the given string and assigns it to the ResponseHeaders field.
func (o *PageElement) SetResponseHeaders(v string) {
	o.ResponseHeaders = &v
}

// GetResolvedIpAddress returns the ResolvedIpAddress field value if set, zero value otherwise.
func (o *PageElement) GetResolvedIpAddress() map[string]interface{} {
	if o == nil || o.ResolvedIpAddress == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ResolvedIpAddress
}

// GetResolvedIpAddressOk returns a tuple with the ResolvedIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageElement) GetResolvedIpAddressOk() (*map[string]interface{}, bool) {
	if o == nil || o.ResolvedIpAddress == nil {
		return nil, false
	}
	return o.ResolvedIpAddress, true
}

// HasResolvedIpAddress returns a boolean if a field has been set.
func (o *PageElement) HasResolvedIpAddress() bool {
	if o != nil && o.ResolvedIpAddress != nil {
		return true
	}

	return false
}

// SetResolvedIpAddress gets a reference to the given map[string]interface{} and assigns it to the ResolvedIpAddress field.
func (o *PageElement) SetResolvedIpAddress(v map[string]interface{}) {
	o.ResolvedIpAddress = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *PageElement) GetGroupIds() []int32 {
	if o == nil || o.GroupIds == nil {
		var ret []int32
		return ret
	}
	return *o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageElement) GetGroupIdsOk() (*[]int32, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *PageElement) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []int32 and assigns it to the GroupIds field.
func (o *PageElement) SetGroupIds(v []int32) {
	o.GroupIds = &v
}

// GetUrlIsBlocked returns the UrlIsBlocked field value
func (o *PageElement) GetUrlIsBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UrlIsBlocked
}

// GetUrlIsBlockedOk returns a tuple with the UrlIsBlocked field value
// and a boolean to check if the value has been set.
func (o *PageElement) GetUrlIsBlockedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UrlIsBlocked, true
}

// SetUrlIsBlocked sets field value
func (o *PageElement) SetUrlIsBlocked(v bool) {
	o.UrlIsBlocked = v
}

func (o PageElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Index"] = o.Index
	}
	if true {
		toSerialize["StartTime"] = o.StartTime
	}
	if true {
		toSerialize["QueueTime"] = o.QueueTime
	}
	if true {
		toSerialize["ResolveTime"] = o.ResolveTime
	}
	if true {
		toSerialize["ConnectTime"] = o.ConnectTime
	}
	if true {
		toSerialize["StaleTime"] = o.StaleTime
	}
	if true {
		toSerialize["HttpsHandshakeTime"] = o.HttpsHandshakeTime
	}
	if true {
		toSerialize["SendTime"] = o.SendTime
	}
	if true {
		toSerialize["WaitTime"] = o.WaitTime
	}
	if true {
		toSerialize["ReceiveTime"] = o.ReceiveTime
	}
	if true {
		toSerialize["TimeoutTime"] = o.TimeoutTime
	}
	if true {
		toSerialize["TotalTime"] = o.TotalTime
	}
	if true {
		toSerialize["HttpStatusCode"] = o.HttpStatusCode
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}
	if true {
		toSerialize["TotalBytes"] = o.TotalBytes
	}
	if o.ElementType != nil {
		toSerialize["ElementType"] = o.ElementType
	}
	if o.RequestHeaders != nil {
		toSerialize["RequestHeaders"] = o.RequestHeaders
	}
	if o.ResponseHeaders != nil {
		toSerialize["ResponseHeaders"] = o.ResponseHeaders
	}
	if o.ResolvedIpAddress != nil {
		toSerialize["ResolvedIpAddress"] = o.ResolvedIpAddress
	}
	if o.GroupIds != nil {
		toSerialize["GroupIds"] = o.GroupIds
	}
	if true {
		toSerialize["UrlIsBlocked"] = o.UrlIsBlocked
	}
	return json.Marshal(toSerialize)
}

type NullablePageElement struct {
	value *PageElement
	isSet bool
}

func (v NullablePageElement) Get() *PageElement {
	return v.value
}

func (v *NullablePageElement) Set(val *PageElement) {
	v.value = val
	v.isSet = true
}

func (v NullablePageElement) IsSet() bool {
	return v.isSet
}

func (v *NullablePageElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageElement(val *PageElement) *NullablePageElement {
	return &NullablePageElement{value: val, isSet: true}
}

func (v NullablePageElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


