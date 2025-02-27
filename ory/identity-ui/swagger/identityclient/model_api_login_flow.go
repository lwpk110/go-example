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

// APILoginFlow struct for APILoginFlow
type APILoginFlow struct {
	Session Session `json:"session"`
	// The Session Token  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization Header:  Authorization: bearer ${session-token}  The session token is only issued for API flows, not for Browser flows!
	SessionToken *string `json:"session_token,omitempty"`
}

// NewAPILoginFlow instantiates a new APILoginFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPILoginFlow(session Session) *APILoginFlow {
	this := APILoginFlow{}
	this.Session = session
	return &this
}

// NewAPILoginFlowWithDefaults instantiates a new APILoginFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPILoginFlowWithDefaults() *APILoginFlow {
	this := APILoginFlow{}
	return &this
}

// GetSession returns the Session field value
func (o *APILoginFlow) GetSession() Session {
	if o == nil {
		var ret Session
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *APILoginFlow) GetSessionOk() (*Session, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *APILoginFlow) SetSession(v Session) {
	o.Session = v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *APILoginFlow) GetSessionToken() string {
	if o == nil || o.SessionToken == nil {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APILoginFlow) GetSessionTokenOk() (*string, bool) {
	if o == nil || o.SessionToken == nil {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *APILoginFlow) HasSessionToken() bool {
	if o != nil && o.SessionToken != nil {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *APILoginFlow) SetSessionToken(v string) {
	o.SessionToken = &v
}

func (o APILoginFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["session"] = o.Session
	}
	if o.SessionToken != nil {
		toSerialize["session_token"] = o.SessionToken
	}
	return json.Marshal(toSerialize)
}

type NullableAPILoginFlow struct {
	value *APILoginFlow
	isSet bool
}

func (v NullableAPILoginFlow) Get() *APILoginFlow {
	return v.value
}

func (v *NullableAPILoginFlow) Set(val *APILoginFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableAPILoginFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableAPILoginFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPILoginFlow(val *APILoginFlow) *NullableAPILoginFlow {
	return &NullableAPILoginFlow{value: val, isSet: true}
}

func (v NullableAPILoginFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPILoginFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


