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

// UiNodeImageAttributes struct for UiNodeImageAttributes
type UiNodeImageAttributes struct {
	// Height of the image
	Height int64 `json:"height"`
	// A unique identifier
	Id string `json:"id"`
	// NodeType represents this node's types. It is a mirror of `node.type` and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \"img\".
	NodeType string `json:"node_type"`
	// The image's source URL.  format: uri
	Src string `json:"src"`
	// Width of the image
	Width int64 `json:"width"`
}

// NewUiNodeImageAttributes instantiates a new UiNodeImageAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiNodeImageAttributes(height int64, id string, nodeType string, src string, width int64) *UiNodeImageAttributes {
	this := UiNodeImageAttributes{}
	this.Height = height
	this.Id = id
	this.NodeType = nodeType
	this.Src = src
	this.Width = width
	return &this
}

// NewUiNodeImageAttributesWithDefaults instantiates a new UiNodeImageAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiNodeImageAttributesWithDefaults() *UiNodeImageAttributes {
	this := UiNodeImageAttributes{}
	return &this
}

// GetHeight returns the Height field value
func (o *UiNodeImageAttributes) GetHeight() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *UiNodeImageAttributes) GetHeightOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *UiNodeImageAttributes) SetHeight(v int64) {
	o.Height = v
}

// GetId returns the Id field value
func (o *UiNodeImageAttributes) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UiNodeImageAttributes) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UiNodeImageAttributes) SetId(v string) {
	o.Id = v
}

// GetNodeType returns the NodeType field value
func (o *UiNodeImageAttributes) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *UiNodeImageAttributes) GetNodeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *UiNodeImageAttributes) SetNodeType(v string) {
	o.NodeType = v
}

// GetSrc returns the Src field value
func (o *UiNodeImageAttributes) GetSrc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Src
}

// GetSrcOk returns a tuple with the Src field value
// and a boolean to check if the value has been set.
func (o *UiNodeImageAttributes) GetSrcOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Src, true
}

// SetSrc sets field value
func (o *UiNodeImageAttributes) SetSrc(v string) {
	o.Src = v
}

// GetWidth returns the Width field value
func (o *UiNodeImageAttributes) GetWidth() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *UiNodeImageAttributes) GetWidthOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *UiNodeImageAttributes) SetWidth(v int64) {
	o.Width = v
}

func (o UiNodeImageAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_type"] = o.NodeType
	}
	if true {
		toSerialize["src"] = o.Src
	}
	if true {
		toSerialize["width"] = o.Width
	}
	return json.Marshal(toSerialize)
}

type NullableUiNodeImageAttributes struct {
	value *UiNodeImageAttributes
	isSet bool
}

func (v NullableUiNodeImageAttributes) Get() *UiNodeImageAttributes {
	return v.value
}

func (v *NullableUiNodeImageAttributes) Set(val *UiNodeImageAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeImageAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeImageAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeImageAttributes(val *UiNodeImageAttributes) *NullableUiNodeImageAttributes {
	return &NullableUiNodeImageAttributes{value: val, isSet: true}
}

func (v NullableUiNodeImageAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeImageAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

