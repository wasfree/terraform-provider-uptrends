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

// Operator Operator
type Operator struct {
	// The unique identifier of this operator
	OperatorGuid *string `json:"OperatorGuid,omitempty"`
	// The hash of this operator.
	Hash *string `json:"Hash,omitempty"`
	// The password is a required field if AllowNativeLogin is set to True
	Password *string `json:"Password,omitempty"`
	// The full name of this operator
	FullName *string `json:"FullName,omitempty"`
	// The email address of this operator. This also serves as the username
	Email *string `json:"Email,omitempty"`
	// The phone number of this operator to which SMS and phone alerts can be sent. Start with a plus (+) sign and your country code
	MobilePhone *string `json:"MobilePhone,omitempty"`
	// The id of the phone number that will be used to send phone alerts (See /OutgoingPhoneNumber API under Miscellaneous for available ids)
	OutgoingPhoneNumberId *int32 `json:"OutgoingPhoneNumberId,omitempty"`
	OutgoingPhoneNumberIdSpecified *bool `json:"OutgoingPhoneNumberIdSpecified,omitempty"`
	// This indicates if the operator is the account administrator.
	IsAccountAdministrator *bool `json:"IsAccountAdministrator,omitempty"`
	// The backup email address of this operator
	BackupEmail *string `json:"BackupEmail,omitempty"`
	// Indicates whether the operator is currently active
	IsOnDuty *bool `json:"IsOnDuty,omitempty"`
	// If ommitted the operator will use the account culture. If set it will override the account default
	CultureName *string `json:"CultureName,omitempty"`
	CultureNameSpecified *bool `json:"CultureNameSpecified,omitempty"`
	// The id of the timezone of this operator (See /Timezone API under Miscellaneous for available timezones)
	TimeZoneId *int32 `json:"TimeZoneId,omitempty"`
	TimeZoneIdSpecified *bool `json:"TimeZoneIdSpecified,omitempty"`
	// The SMS provider used to send out SMS alerts
	SmsProvider *SmsProvider `json:"SmsProvider,omitempty"`
	// Set to True to override the default behavior of sending SMS alerts with textual sender ID
	UseNumericSender *bool `json:"UseNumericSender,omitempty"`
	UseNumericSenderSpecified *bool `json:"UseNumericSenderSpecified,omitempty"`
	// This can only be set to false if the account has SSO enabled. Ommitting or providing null will use the account default
	AllowNativeLogin *bool `json:"AllowNativeLogin,omitempty"`
	AllowNativeLoginSpecified *bool `json:"AllowNativeLoginSpecified,omitempty"`
	// This can only be set to true if the account has SSO enabled. Ommitting or providing null will use the account default
	AllowSingleSignon *bool `json:"AllowSingleSignon,omitempty"`
	AllowSingleSignonSpecified *bool `json:"AllowSingleSignonSpecified,omitempty"`
	DefaultDashboard *string `json:"DefaultDashboard,omitempty"`
	SetupMode *OperatorSetupMode `json:"SetupMode,omitempty"`
}

// NewOperator instantiates a new Operator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperator() *Operator {
	this := Operator{}
	return &this
}

// NewOperatorWithDefaults instantiates a new Operator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorWithDefaults() *Operator {
	this := Operator{}
	return &this
}

// GetOperatorGuid returns the OperatorGuid field value if set, zero value otherwise.
func (o *Operator) GetOperatorGuid() string {
	if o == nil || o.OperatorGuid == nil {
		var ret string
		return ret
	}
	return *o.OperatorGuid
}

// GetOperatorGuidOk returns a tuple with the OperatorGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetOperatorGuidOk() (*string, bool) {
	if o == nil || o.OperatorGuid == nil {
		return nil, false
	}
	return o.OperatorGuid, true
}

// HasOperatorGuid returns a boolean if a field has been set.
func (o *Operator) HasOperatorGuid() bool {
	if o != nil && o.OperatorGuid != nil {
		return true
	}

	return false
}

// SetOperatorGuid gets a reference to the given string and assigns it to the OperatorGuid field.
func (o *Operator) SetOperatorGuid(v string) {
	o.OperatorGuid = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *Operator) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *Operator) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *Operator) SetHash(v string) {
	o.Hash = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Operator) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Operator) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Operator) SetPassword(v string) {
	o.Password = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *Operator) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *Operator) HasFullName() bool {
	if o != nil && o.FullName != nil {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *Operator) SetFullName(v string) {
	o.FullName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Operator) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Operator) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Operator) SetEmail(v string) {
	o.Email = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *Operator) GetMobilePhone() string {
	if o == nil || o.MobilePhone == nil {
		var ret string
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetMobilePhoneOk() (*string, bool) {
	if o == nil || o.MobilePhone == nil {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *Operator) HasMobilePhone() bool {
	if o != nil && o.MobilePhone != nil {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.
func (o *Operator) SetMobilePhone(v string) {
	o.MobilePhone = &v
}

// GetOutgoingPhoneNumberId returns the OutgoingPhoneNumberId field value if set, zero value otherwise.
func (o *Operator) GetOutgoingPhoneNumberId() int32 {
	if o == nil || o.OutgoingPhoneNumberId == nil {
		var ret int32
		return ret
	}
	return *o.OutgoingPhoneNumberId
}

// GetOutgoingPhoneNumberIdOk returns a tuple with the OutgoingPhoneNumberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetOutgoingPhoneNumberIdOk() (*int32, bool) {
	if o == nil || o.OutgoingPhoneNumberId == nil {
		return nil, false
	}
	return o.OutgoingPhoneNumberId, true
}

// HasOutgoingPhoneNumberId returns a boolean if a field has been set.
func (o *Operator) HasOutgoingPhoneNumberId() bool {
	if o != nil && o.OutgoingPhoneNumberId != nil {
		return true
	}

	return false
}

// SetOutgoingPhoneNumberId gets a reference to the given int32 and assigns it to the OutgoingPhoneNumberId field.
func (o *Operator) SetOutgoingPhoneNumberId(v int32) {
	o.OutgoingPhoneNumberId = &v
}

// GetOutgoingPhoneNumberIdSpecified returns the OutgoingPhoneNumberIdSpecified field value if set, zero value otherwise.
func (o *Operator) GetOutgoingPhoneNumberIdSpecified() bool {
	if o == nil || o.OutgoingPhoneNumberIdSpecified == nil {
		var ret bool
		return ret
	}
	return *o.OutgoingPhoneNumberIdSpecified
}

// GetOutgoingPhoneNumberIdSpecifiedOk returns a tuple with the OutgoingPhoneNumberIdSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetOutgoingPhoneNumberIdSpecifiedOk() (*bool, bool) {
	if o == nil || o.OutgoingPhoneNumberIdSpecified == nil {
		return nil, false
	}
	return o.OutgoingPhoneNumberIdSpecified, true
}

// HasOutgoingPhoneNumberIdSpecified returns a boolean if a field has been set.
func (o *Operator) HasOutgoingPhoneNumberIdSpecified() bool {
	if o != nil && o.OutgoingPhoneNumberIdSpecified != nil {
		return true
	}

	return false
}

// SetOutgoingPhoneNumberIdSpecified gets a reference to the given bool and assigns it to the OutgoingPhoneNumberIdSpecified field.
func (o *Operator) SetOutgoingPhoneNumberIdSpecified(v bool) {
	o.OutgoingPhoneNumberIdSpecified = &v
}

// GetIsAccountAdministrator returns the IsAccountAdministrator field value if set, zero value otherwise.
func (o *Operator) GetIsAccountAdministrator() bool {
	if o == nil || o.IsAccountAdministrator == nil {
		var ret bool
		return ret
	}
	return *o.IsAccountAdministrator
}

// GetIsAccountAdministratorOk returns a tuple with the IsAccountAdministrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetIsAccountAdministratorOk() (*bool, bool) {
	if o == nil || o.IsAccountAdministrator == nil {
		return nil, false
	}
	return o.IsAccountAdministrator, true
}

// HasIsAccountAdministrator returns a boolean if a field has been set.
func (o *Operator) HasIsAccountAdministrator() bool {
	if o != nil && o.IsAccountAdministrator != nil {
		return true
	}

	return false
}

// SetIsAccountAdministrator gets a reference to the given bool and assigns it to the IsAccountAdministrator field.
func (o *Operator) SetIsAccountAdministrator(v bool) {
	o.IsAccountAdministrator = &v
}

// GetBackupEmail returns the BackupEmail field value if set, zero value otherwise.
func (o *Operator) GetBackupEmail() string {
	if o == nil || o.BackupEmail == nil {
		var ret string
		return ret
	}
	return *o.BackupEmail
}

// GetBackupEmailOk returns a tuple with the BackupEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetBackupEmailOk() (*string, bool) {
	if o == nil || o.BackupEmail == nil {
		return nil, false
	}
	return o.BackupEmail, true
}

// HasBackupEmail returns a boolean if a field has been set.
func (o *Operator) HasBackupEmail() bool {
	if o != nil && o.BackupEmail != nil {
		return true
	}

	return false
}

// SetBackupEmail gets a reference to the given string and assigns it to the BackupEmail field.
func (o *Operator) SetBackupEmail(v string) {
	o.BackupEmail = &v
}

// GetIsOnDuty returns the IsOnDuty field value if set, zero value otherwise.
func (o *Operator) GetIsOnDuty() bool {
	if o == nil || o.IsOnDuty == nil {
		var ret bool
		return ret
	}
	return *o.IsOnDuty
}

// GetIsOnDutyOk returns a tuple with the IsOnDuty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetIsOnDutyOk() (*bool, bool) {
	if o == nil || o.IsOnDuty == nil {
		return nil, false
	}
	return o.IsOnDuty, true
}

// HasIsOnDuty returns a boolean if a field has been set.
func (o *Operator) HasIsOnDuty() bool {
	if o != nil && o.IsOnDuty != nil {
		return true
	}

	return false
}

// SetIsOnDuty gets a reference to the given bool and assigns it to the IsOnDuty field.
func (o *Operator) SetIsOnDuty(v bool) {
	o.IsOnDuty = &v
}

// GetCultureName returns the CultureName field value if set, zero value otherwise.
func (o *Operator) GetCultureName() string {
	if o == nil || o.CultureName == nil {
		var ret string
		return ret
	}
	return *o.CultureName
}

// GetCultureNameOk returns a tuple with the CultureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetCultureNameOk() (*string, bool) {
	if o == nil || o.CultureName == nil {
		return nil, false
	}
	return o.CultureName, true
}

// HasCultureName returns a boolean if a field has been set.
func (o *Operator) HasCultureName() bool {
	if o != nil && o.CultureName != nil {
		return true
	}

	return false
}

// SetCultureName gets a reference to the given string and assigns it to the CultureName field.
func (o *Operator) SetCultureName(v string) {
	o.CultureName = &v
}

// GetCultureNameSpecified returns the CultureNameSpecified field value if set, zero value otherwise.
func (o *Operator) GetCultureNameSpecified() bool {
	if o == nil || o.CultureNameSpecified == nil {
		var ret bool
		return ret
	}
	return *o.CultureNameSpecified
}

// GetCultureNameSpecifiedOk returns a tuple with the CultureNameSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetCultureNameSpecifiedOk() (*bool, bool) {
	if o == nil || o.CultureNameSpecified == nil {
		return nil, false
	}
	return o.CultureNameSpecified, true
}

// HasCultureNameSpecified returns a boolean if a field has been set.
func (o *Operator) HasCultureNameSpecified() bool {
	if o != nil && o.CultureNameSpecified != nil {
		return true
	}

	return false
}

// SetCultureNameSpecified gets a reference to the given bool and assigns it to the CultureNameSpecified field.
func (o *Operator) SetCultureNameSpecified(v bool) {
	o.CultureNameSpecified = &v
}

// GetTimeZoneId returns the TimeZoneId field value if set, zero value otherwise.
func (o *Operator) GetTimeZoneId() int32 {
	if o == nil || o.TimeZoneId == nil {
		var ret int32
		return ret
	}
	return *o.TimeZoneId
}

// GetTimeZoneIdOk returns a tuple with the TimeZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetTimeZoneIdOk() (*int32, bool) {
	if o == nil || o.TimeZoneId == nil {
		return nil, false
	}
	return o.TimeZoneId, true
}

// HasTimeZoneId returns a boolean if a field has been set.
func (o *Operator) HasTimeZoneId() bool {
	if o != nil && o.TimeZoneId != nil {
		return true
	}

	return false
}

// SetTimeZoneId gets a reference to the given int32 and assigns it to the TimeZoneId field.
func (o *Operator) SetTimeZoneId(v int32) {
	o.TimeZoneId = &v
}

// GetTimeZoneIdSpecified returns the TimeZoneIdSpecified field value if set, zero value otherwise.
func (o *Operator) GetTimeZoneIdSpecified() bool {
	if o == nil || o.TimeZoneIdSpecified == nil {
		var ret bool
		return ret
	}
	return *o.TimeZoneIdSpecified
}

// GetTimeZoneIdSpecifiedOk returns a tuple with the TimeZoneIdSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetTimeZoneIdSpecifiedOk() (*bool, bool) {
	if o == nil || o.TimeZoneIdSpecified == nil {
		return nil, false
	}
	return o.TimeZoneIdSpecified, true
}

// HasTimeZoneIdSpecified returns a boolean if a field has been set.
func (o *Operator) HasTimeZoneIdSpecified() bool {
	if o != nil && o.TimeZoneIdSpecified != nil {
		return true
	}

	return false
}

// SetTimeZoneIdSpecified gets a reference to the given bool and assigns it to the TimeZoneIdSpecified field.
func (o *Operator) SetTimeZoneIdSpecified(v bool) {
	o.TimeZoneIdSpecified = &v
}

// GetSmsProvider returns the SmsProvider field value if set, zero value otherwise.
func (o *Operator) GetSmsProvider() SmsProvider {
	if o == nil || o.SmsProvider == nil {
		var ret SmsProvider
		return ret
	}
	return *o.SmsProvider
}

// GetSmsProviderOk returns a tuple with the SmsProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetSmsProviderOk() (*SmsProvider, bool) {
	if o == nil || o.SmsProvider == nil {
		return nil, false
	}
	return o.SmsProvider, true
}

// HasSmsProvider returns a boolean if a field has been set.
func (o *Operator) HasSmsProvider() bool {
	if o != nil && o.SmsProvider != nil {
		return true
	}

	return false
}

// SetSmsProvider gets a reference to the given SmsProvider and assigns it to the SmsProvider field.
func (o *Operator) SetSmsProvider(v SmsProvider) {
	o.SmsProvider = &v
}

// GetUseNumericSender returns the UseNumericSender field value if set, zero value otherwise.
func (o *Operator) GetUseNumericSender() bool {
	if o == nil || o.UseNumericSender == nil {
		var ret bool
		return ret
	}
	return *o.UseNumericSender
}

// GetUseNumericSenderOk returns a tuple with the UseNumericSender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetUseNumericSenderOk() (*bool, bool) {
	if o == nil || o.UseNumericSender == nil {
		return nil, false
	}
	return o.UseNumericSender, true
}

// HasUseNumericSender returns a boolean if a field has been set.
func (o *Operator) HasUseNumericSender() bool {
	if o != nil && o.UseNumericSender != nil {
		return true
	}

	return false
}

// SetUseNumericSender gets a reference to the given bool and assigns it to the UseNumericSender field.
func (o *Operator) SetUseNumericSender(v bool) {
	o.UseNumericSender = &v
}

// GetUseNumericSenderSpecified returns the UseNumericSenderSpecified field value if set, zero value otherwise.
func (o *Operator) GetUseNumericSenderSpecified() bool {
	if o == nil || o.UseNumericSenderSpecified == nil {
		var ret bool
		return ret
	}
	return *o.UseNumericSenderSpecified
}

// GetUseNumericSenderSpecifiedOk returns a tuple with the UseNumericSenderSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetUseNumericSenderSpecifiedOk() (*bool, bool) {
	if o == nil || o.UseNumericSenderSpecified == nil {
		return nil, false
	}
	return o.UseNumericSenderSpecified, true
}

// HasUseNumericSenderSpecified returns a boolean if a field has been set.
func (o *Operator) HasUseNumericSenderSpecified() bool {
	if o != nil && o.UseNumericSenderSpecified != nil {
		return true
	}

	return false
}

// SetUseNumericSenderSpecified gets a reference to the given bool and assigns it to the UseNumericSenderSpecified field.
func (o *Operator) SetUseNumericSenderSpecified(v bool) {
	o.UseNumericSenderSpecified = &v
}

// GetAllowNativeLogin returns the AllowNativeLogin field value if set, zero value otherwise.
func (o *Operator) GetAllowNativeLogin() bool {
	if o == nil || o.AllowNativeLogin == nil {
		var ret bool
		return ret
	}
	return *o.AllowNativeLogin
}

// GetAllowNativeLoginOk returns a tuple with the AllowNativeLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetAllowNativeLoginOk() (*bool, bool) {
	if o == nil || o.AllowNativeLogin == nil {
		return nil, false
	}
	return o.AllowNativeLogin, true
}

// HasAllowNativeLogin returns a boolean if a field has been set.
func (o *Operator) HasAllowNativeLogin() bool {
	if o != nil && o.AllowNativeLogin != nil {
		return true
	}

	return false
}

// SetAllowNativeLogin gets a reference to the given bool and assigns it to the AllowNativeLogin field.
func (o *Operator) SetAllowNativeLogin(v bool) {
	o.AllowNativeLogin = &v
}

// GetAllowNativeLoginSpecified returns the AllowNativeLoginSpecified field value if set, zero value otherwise.
func (o *Operator) GetAllowNativeLoginSpecified() bool {
	if o == nil || o.AllowNativeLoginSpecified == nil {
		var ret bool
		return ret
	}
	return *o.AllowNativeLoginSpecified
}

// GetAllowNativeLoginSpecifiedOk returns a tuple with the AllowNativeLoginSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetAllowNativeLoginSpecifiedOk() (*bool, bool) {
	if o == nil || o.AllowNativeLoginSpecified == nil {
		return nil, false
	}
	return o.AllowNativeLoginSpecified, true
}

// HasAllowNativeLoginSpecified returns a boolean if a field has been set.
func (o *Operator) HasAllowNativeLoginSpecified() bool {
	if o != nil && o.AllowNativeLoginSpecified != nil {
		return true
	}

	return false
}

// SetAllowNativeLoginSpecified gets a reference to the given bool and assigns it to the AllowNativeLoginSpecified field.
func (o *Operator) SetAllowNativeLoginSpecified(v bool) {
	o.AllowNativeLoginSpecified = &v
}

// GetAllowSingleSignon returns the AllowSingleSignon field value if set, zero value otherwise.
func (o *Operator) GetAllowSingleSignon() bool {
	if o == nil || o.AllowSingleSignon == nil {
		var ret bool
		return ret
	}
	return *o.AllowSingleSignon
}

// GetAllowSingleSignonOk returns a tuple with the AllowSingleSignon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetAllowSingleSignonOk() (*bool, bool) {
	if o == nil || o.AllowSingleSignon == nil {
		return nil, false
	}
	return o.AllowSingleSignon, true
}

// HasAllowSingleSignon returns a boolean if a field has been set.
func (o *Operator) HasAllowSingleSignon() bool {
	if o != nil && o.AllowSingleSignon != nil {
		return true
	}

	return false
}

// SetAllowSingleSignon gets a reference to the given bool and assigns it to the AllowSingleSignon field.
func (o *Operator) SetAllowSingleSignon(v bool) {
	o.AllowSingleSignon = &v
}

// GetAllowSingleSignonSpecified returns the AllowSingleSignonSpecified field value if set, zero value otherwise.
func (o *Operator) GetAllowSingleSignonSpecified() bool {
	if o == nil || o.AllowSingleSignonSpecified == nil {
		var ret bool
		return ret
	}
	return *o.AllowSingleSignonSpecified
}

// GetAllowSingleSignonSpecifiedOk returns a tuple with the AllowSingleSignonSpecified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetAllowSingleSignonSpecifiedOk() (*bool, bool) {
	if o == nil || o.AllowSingleSignonSpecified == nil {
		return nil, false
	}
	return o.AllowSingleSignonSpecified, true
}

// HasAllowSingleSignonSpecified returns a boolean if a field has been set.
func (o *Operator) HasAllowSingleSignonSpecified() bool {
	if o != nil && o.AllowSingleSignonSpecified != nil {
		return true
	}

	return false
}

// SetAllowSingleSignonSpecified gets a reference to the given bool and assigns it to the AllowSingleSignonSpecified field.
func (o *Operator) SetAllowSingleSignonSpecified(v bool) {
	o.AllowSingleSignonSpecified = &v
}

// GetDefaultDashboard returns the DefaultDashboard field value if set, zero value otherwise.
func (o *Operator) GetDefaultDashboard() string {
	if o == nil || o.DefaultDashboard == nil {
		var ret string
		return ret
	}
	return *o.DefaultDashboard
}

// GetDefaultDashboardOk returns a tuple with the DefaultDashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetDefaultDashboardOk() (*string, bool) {
	if o == nil || o.DefaultDashboard == nil {
		return nil, false
	}
	return o.DefaultDashboard, true
}

// HasDefaultDashboard returns a boolean if a field has been set.
func (o *Operator) HasDefaultDashboard() bool {
	if o != nil && o.DefaultDashboard != nil {
		return true
	}

	return false
}

// SetDefaultDashboard gets a reference to the given string and assigns it to the DefaultDashboard field.
func (o *Operator) SetDefaultDashboard(v string) {
	o.DefaultDashboard = &v
}

// GetSetupMode returns the SetupMode field value if set, zero value otherwise.
func (o *Operator) GetSetupMode() OperatorSetupMode {
	if o == nil || o.SetupMode == nil {
		var ret OperatorSetupMode
		return ret
	}
	return *o.SetupMode
}

// GetSetupModeOk returns a tuple with the SetupMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetSetupModeOk() (*OperatorSetupMode, bool) {
	if o == nil || o.SetupMode == nil {
		return nil, false
	}
	return o.SetupMode, true
}

// HasSetupMode returns a boolean if a field has been set.
func (o *Operator) HasSetupMode() bool {
	if o != nil && o.SetupMode != nil {
		return true
	}

	return false
}

// SetSetupMode gets a reference to the given OperatorSetupMode and assigns it to the SetupMode field.
func (o *Operator) SetSetupMode(v OperatorSetupMode) {
	o.SetupMode = &v
}

func (o Operator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OperatorGuid != nil {
		toSerialize["OperatorGuid"] = o.OperatorGuid
	}
	if o.Hash != nil {
		toSerialize["Hash"] = o.Hash
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.FullName != nil {
		toSerialize["FullName"] = o.FullName
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	if o.MobilePhone != nil {
		toSerialize["MobilePhone"] = o.MobilePhone
	}
	if o.OutgoingPhoneNumberId != nil {
		toSerialize["OutgoingPhoneNumberId"] = o.OutgoingPhoneNumberId
	}
	if o.OutgoingPhoneNumberIdSpecified != nil {
		toSerialize["OutgoingPhoneNumberIdSpecified"] = o.OutgoingPhoneNumberIdSpecified
	}
	if o.IsAccountAdministrator != nil {
		toSerialize["IsAccountAdministrator"] = o.IsAccountAdministrator
	}
	if o.BackupEmail != nil {
		toSerialize["BackupEmail"] = o.BackupEmail
	}
	if o.IsOnDuty != nil {
		toSerialize["IsOnDuty"] = o.IsOnDuty
	}
	if o.CultureName != nil {
		toSerialize["CultureName"] = o.CultureName
	}
	if o.CultureNameSpecified != nil {
		toSerialize["CultureNameSpecified"] = o.CultureNameSpecified
	}
	if o.TimeZoneId != nil {
		toSerialize["TimeZoneId"] = o.TimeZoneId
	}
	if o.TimeZoneIdSpecified != nil {
		toSerialize["TimeZoneIdSpecified"] = o.TimeZoneIdSpecified
	}
	if o.SmsProvider != nil {
		toSerialize["SmsProvider"] = o.SmsProvider
	}
	if o.UseNumericSender != nil {
		toSerialize["UseNumericSender"] = o.UseNumericSender
	}
	if o.UseNumericSenderSpecified != nil {
		toSerialize["UseNumericSenderSpecified"] = o.UseNumericSenderSpecified
	}
	if o.AllowNativeLogin != nil {
		toSerialize["AllowNativeLogin"] = o.AllowNativeLogin
	}
	if o.AllowNativeLoginSpecified != nil {
		toSerialize["AllowNativeLoginSpecified"] = o.AllowNativeLoginSpecified
	}
	if o.AllowSingleSignon != nil {
		toSerialize["AllowSingleSignon"] = o.AllowSingleSignon
	}
	if o.AllowSingleSignonSpecified != nil {
		toSerialize["AllowSingleSignonSpecified"] = o.AllowSingleSignonSpecified
	}
	if o.DefaultDashboard != nil {
		toSerialize["DefaultDashboard"] = o.DefaultDashboard
	}
	if o.SetupMode != nil {
		toSerialize["SetupMode"] = o.SetupMode
	}
	return json.Marshal(toSerialize)
}

type NullableOperator struct {
	value *Operator
	isSet bool
}

func (v NullableOperator) Get() *Operator {
	return v.value
}

func (v *NullableOperator) Set(val *Operator) {
	v.value = val
	v.isSet = true
}

func (v NullableOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperator(val *Operator) *NullableOperator {
	return &NullableOperator{value: val, isSet: true}
}

func (v NullableOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


