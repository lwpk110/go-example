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

// BrowserLocationChangeRequiredResponse struct for BrowserLocationChangeRequiredResponse
type BrowserLocationChangeRequiredResponse struct {
	Code *Code `json:"code,omitempty"`
	// Detail contains further information on the nature of the error.
	Detail *string `json:"detail,omitempty"`
	// Message is the error message.
	Msg *string `json:"msg,omitempty"`
	// Since when the flow has expired
	RedirectBrowserTo *string `json:"redirect_browser_to,omitempty"`
	// TraceId is the identifier for a trace. It is globally unique.
	TraceId *string `json:"traceId,omitempty"`
	// Type A URI reference that identifies the error type.
	Type *string `json:"type,omitempty"`
}

// NewBrowserLocationChangeRequiredResponse instantiates a new BrowserLocationChangeRequiredResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserLocationChangeRequiredResponse() *BrowserLocationChangeRequiredResponse {
	this := BrowserLocationChangeRequiredResponse{}
	return &this
}

// NewBrowserLocationChangeRequiredResponseWithDefaults instantiates a new BrowserLocationChangeRequiredResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserLocationChangeRequiredResponseWithDefaults() *BrowserLocationChangeRequiredResponse {
	this := BrowserLocationChangeRequiredResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BrowserLocationChangeRequiredResponse) GetCode() Code {
	if o == nil || o.Code == nil {
		var ret Code
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserLocationChangeRequiredResponse) GetCodeOk() (*Code, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BrowserLocationChangeRequiredResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given Code and assigns it to the Code field.
func (o *BrowserLocationChangeRequiredResponse) SetCode(v Code) {
	o.Code = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *BrowserLocationChangeRequiredResponse) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserLocationChangeRequiredResponse) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *BrowserLocationChangeRequiredResponse) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *BrowserLocationChangeRequiredResponse) SetDetail(v string) {
	o.Detail = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *BrowserLocationChangeRequiredResponse) GetMsg() string {
	if o == nil || o.Msg == nil {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserLocationChangeRequiredResponse) GetMsgOk() (*string, bool) {
	if o == nil || o.Msg == nil {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *BrowserLocationChangeRequiredResponse) HasMsg() bool {
	if o != nil && o.Msg != nil {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *BrowserLocationChangeRequiredResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetRedirectBrowserTo returns the RedirectBrowserTo field value if set, zero value otherwise.
func (o *BrowserLocationChangeRequiredResponse) GetRedirectBrowserTo() string {
	if o == nil || o.RedirectBrowserTo == nil {
		var ret string
		return ret
	}
	return *o.RedirectBrowserTo
}

// GetRedirectBrowserToOk returns a tuple with the RedirectBrowserTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserLocationChangeRequiredResponse) GetRedirectBrowserToOk() (*string, bool) {
	if o == nil || o.RedirectBrowserTo == nil {
		return nil, false
	}
	return o.RedirectBrowserTo, true
}

// HasRedirectBrowserTo returns a boolean if a field has been set.
func (o *BrowserLocationChangeRequiredResponse) HasRedirectBrowserTo() bool {
	if o != nil && o.RedirectBrowserTo != nil {
		return true
	}

	return false
}

// SetRedirectBrowserTo gets a reference to the given string and assigns it to the RedirectBrowserTo field.
func (o *BrowserLocationChangeRequiredResponse) SetRedirectBrowserTo(v string) {
	o.RedirectBrowserTo = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *BrowserLocationChangeRequiredResponse) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserLocationChangeRequiredResponse) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *BrowserLocationChangeRequiredResponse) HasTraceId() bool {
	if o != nil && o.TraceId != nil {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *BrowserLocationChangeRequiredResponse) SetTraceId(v string) {
	o.TraceId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BrowserLocationChangeRequiredResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserLocationChangeRequiredResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BrowserLocationChangeRequiredResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BrowserLocationChangeRequiredResponse) SetType(v string) {
	o.Type = &v
}

func (o BrowserLocationChangeRequiredResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Msg != nil {
		toSerialize["msg"] = o.Msg
	}
	if o.RedirectBrowserTo != nil {
		toSerialize["redirect_browser_to"] = o.RedirectBrowserTo
	}
	if o.TraceId != nil {
		toSerialize["traceId"] = o.TraceId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBrowserLocationChangeRequiredResponse struct {
	value *BrowserLocationChangeRequiredResponse
	isSet bool
}

func (v NullableBrowserLocationChangeRequiredResponse) Get() *BrowserLocationChangeRequiredResponse {
	return v.value
}

func (v *NullableBrowserLocationChangeRequiredResponse) Set(val *BrowserLocationChangeRequiredResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserLocationChangeRequiredResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserLocationChangeRequiredResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserLocationChangeRequiredResponse(val *BrowserLocationChangeRequiredResponse) *NullableBrowserLocationChangeRequiredResponse {
	return &NullableBrowserLocationChangeRequiredResponse{value: val, isSet: true}
}

func (v NullableBrowserLocationChangeRequiredResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserLocationChangeRequiredResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


