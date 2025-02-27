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
	"time"
)

// RecoveryFlow This request is used when an identity wants to recover their account.  We recommend reading the [Account Recovery Documentation](../self-service/flows/password-reset-account-recovery)
type RecoveryFlow struct {
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ExpiresAt is the time (UTC) when the request expires. If the user still wishes to update the setting, a new request has to be initiated.
	ExpiresAt time.Time `json:"expires_at"`
	// ID represents the request's unique ID. When performing the recovery flow, this represents the id in the recovery ui's query parameter: http://<selfservice.flows.recovery.ui_url>?request=<id>
	Id string `json:"id"`
	// IssuedAt is the time (UTC) when the request occurred.
	IssuedAt time.Time `json:"issued_at"`
	Method *RecoveryMethod `json:"method,omitempty"`
	// RequestURL is the initial URL that was requested from Identity. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl string `json:"request_url"`
	// ReturnTo contains the requested return_to URL.
	ReturnTo *string `json:"return_to,omitempty"`
	State RecoveryState `json:"state"`
	Type FlowType `json:"type"`
	Ui UiContainer `json:"ui"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewRecoveryFlow instantiates a new RecoveryFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, state RecoveryState, type_ FlowType, ui UiContainer) *RecoveryFlow {
	this := RecoveryFlow{}
	this.ExpiresAt = expiresAt
	this.Id = id
	this.IssuedAt = issuedAt
	this.RequestUrl = requestUrl
	this.State = state
	this.Type = type_
	this.Ui = ui
	return &this
}

// NewRecoveryFlowWithDefaults instantiates a new RecoveryFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryFlowWithDefaults() *RecoveryFlow {
	this := RecoveryFlow{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RecoveryFlow) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RecoveryFlow) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RecoveryFlow) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *RecoveryFlow) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *RecoveryFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *RecoveryFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RecoveryFlow) SetId(v string) {
	o.Id = v
}

// GetIssuedAt returns the IssuedAt field value
func (o *RecoveryFlow) GetIssuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *RecoveryFlow) SetIssuedAt(v time.Time) {
	o.IssuedAt = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RecoveryFlow) GetMethod() RecoveryMethod {
	if o == nil || o.Method == nil {
		var ret RecoveryMethod
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetMethodOk() (*RecoveryMethod, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RecoveryFlow) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given RecoveryMethod and assigns it to the Method field.
func (o *RecoveryFlow) SetMethod(v RecoveryMethod) {
	o.Method = &v
}

// GetRequestUrl returns the RequestUrl field value
func (o *RecoveryFlow) GetRequestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetRequestUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestUrl, true
}

// SetRequestUrl sets field value
func (o *RecoveryFlow) SetRequestUrl(v string) {
	o.RequestUrl = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *RecoveryFlow) GetReturnTo() string {
	if o == nil || o.ReturnTo == nil {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetReturnToOk() (*string, bool) {
	if o == nil || o.ReturnTo == nil {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *RecoveryFlow) HasReturnTo() bool {
	if o != nil && o.ReturnTo != nil {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *RecoveryFlow) SetReturnTo(v string) {
	o.ReturnTo = &v
}

// GetState returns the State field value
func (o *RecoveryFlow) GetState() RecoveryState {
	if o == nil {
		var ret RecoveryState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetStateOk() (*RecoveryState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *RecoveryFlow) SetState(v RecoveryState) {
	o.State = v
}

// GetType returns the Type field value
func (o *RecoveryFlow) GetType() FlowType {
	if o == nil {
		var ret FlowType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetTypeOk() (*FlowType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RecoveryFlow) SetType(v FlowType) {
	o.Type = v
}

// GetUi returns the Ui field value
func (o *RecoveryFlow) GetUi() UiContainer {
	if o == nil {
		var ret UiContainer
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetUiOk() (*UiContainer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ui, true
}

// SetUi sets field value
func (o *RecoveryFlow) SetUi(v UiContainer) {
	o.Ui = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RecoveryFlow) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RecoveryFlow) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RecoveryFlow) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o RecoveryFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["issued_at"] = o.IssuedAt
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["request_url"] = o.RequestUrl
	}
	if o.ReturnTo != nil {
		toSerialize["return_to"] = o.ReturnTo
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["ui"] = o.Ui
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRecoveryFlow struct {
	value *RecoveryFlow
	isSet bool
}

func (v NullableRecoveryFlow) Get() *RecoveryFlow {
	return v.value
}

func (v *NullableRecoveryFlow) Set(val *RecoveryFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryFlow(val *RecoveryFlow) *NullableRecoveryFlow {
	return &NullableRecoveryFlow{value: val, isSet: true}
}

func (v NullableRecoveryFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


