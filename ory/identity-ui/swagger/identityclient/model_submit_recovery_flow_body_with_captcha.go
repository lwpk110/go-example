/*
 * Identity
 *
 * Welcome to the Identity HTTP API documentation! You will find documentation for all HTTP APIs here.
 *
 * API version: latest
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package identityclient

import (
	"encoding/json"
)

// SubmitRecoveryFlowBodyWithCaptcha struct for SubmitRecoveryFlowBodyWithCaptcha
type SubmitRecoveryFlowBodyWithCaptcha struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method supports `captcha` only right now.
	Method string `json:"method"`
	// Recovery Token
	Token *string `json:"token,omitempty"`
	Traits RecoveryTraits `json:"traits"`
}

// NewSubmitRecoveryFlowBodyWithCaptcha instantiates a new SubmitRecoveryFlowBodyWithCaptcha object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitRecoveryFlowBodyWithCaptcha(method string, traits RecoveryTraits) *SubmitRecoveryFlowBodyWithCaptcha {
	this := SubmitRecoveryFlowBodyWithCaptcha{}
	this.Method = method
	this.Traits = traits
	return &this
}

// NewSubmitRecoveryFlowBodyWithCaptchaWithDefaults instantiates a new SubmitRecoveryFlowBodyWithCaptcha object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitRecoveryFlowBodyWithCaptchaWithDefaults() *SubmitRecoveryFlowBodyWithCaptcha {
	this := SubmitRecoveryFlowBodyWithCaptcha{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitRecoveryFlowBodyWithCaptcha) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRecoveryFlowBodyWithCaptcha) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitRecoveryFlowBodyWithCaptcha) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitRecoveryFlowBodyWithCaptcha) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitRecoveryFlowBodyWithCaptcha) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitRecoveryFlowBodyWithCaptcha) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitRecoveryFlowBodyWithCaptcha) SetMethod(v string) {
	o.Method = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SubmitRecoveryFlowBodyWithCaptcha) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitRecoveryFlowBodyWithCaptcha) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SubmitRecoveryFlowBodyWithCaptcha) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SubmitRecoveryFlowBodyWithCaptcha) SetToken(v string) {
	o.Token = &v
}

// GetTraits returns the Traits field value
func (o *SubmitRecoveryFlowBodyWithCaptcha) GetTraits() RecoveryTraits {
	if o == nil {
		var ret RecoveryTraits
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *SubmitRecoveryFlowBodyWithCaptcha) GetTraitsOk() (*RecoveryTraits, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Traits, true
}

// SetTraits sets field value
func (o *SubmitRecoveryFlowBodyWithCaptcha) SetTraits(v RecoveryTraits) {
	o.Traits = v
}

func (o SubmitRecoveryFlowBodyWithCaptcha) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["traits"] = o.Traits
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitRecoveryFlowBodyWithCaptcha struct {
	value *SubmitRecoveryFlowBodyWithCaptcha
	isSet bool
}

func (v NullableSubmitRecoveryFlowBodyWithCaptcha) Get() *SubmitRecoveryFlowBodyWithCaptcha {
	return v.value
}

func (v *NullableSubmitRecoveryFlowBodyWithCaptcha) Set(val *SubmitRecoveryFlowBodyWithCaptcha) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitRecoveryFlowBodyWithCaptcha) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitRecoveryFlowBodyWithCaptcha) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitRecoveryFlowBodyWithCaptcha(val *SubmitRecoveryFlowBodyWithCaptcha) *NullableSubmitRecoveryFlowBodyWithCaptcha {
	return &NullableSubmitRecoveryFlowBodyWithCaptcha{value: val, isSet: true}
}

func (v NullableSubmitRecoveryFlowBodyWithCaptcha) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitRecoveryFlowBodyWithCaptcha) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

