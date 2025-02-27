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
	"fmt"
)

// SubmitRegistrationFlowBody - struct for SubmitRegistrationFlowBody
type SubmitRegistrationFlowBody struct {
	SubmitRegistrationFlowBodyWithOidc *SubmitRegistrationFlowBodyWithOidc
	SubmitRegistrationFlowBodyWithPassword *SubmitRegistrationFlowBodyWithPassword
}

// SubmitRegistrationFlowBodyWithOidcAsSubmitRegistrationFlowBody is a convenience function that returns SubmitRegistrationFlowBodyWithOidc wrapped in SubmitRegistrationFlowBody
func SubmitRegistrationFlowBodyWithOidcAsSubmitRegistrationFlowBody(v *SubmitRegistrationFlowBodyWithOidc) SubmitRegistrationFlowBody {
	return SubmitRegistrationFlowBody{
		SubmitRegistrationFlowBodyWithOidc: v,
	}
}

// SubmitRegistrationFlowBodyWithPasswordAsSubmitRegistrationFlowBody is a convenience function that returns SubmitRegistrationFlowBodyWithPassword wrapped in SubmitRegistrationFlowBody
func SubmitRegistrationFlowBodyWithPasswordAsSubmitRegistrationFlowBody(v *SubmitRegistrationFlowBodyWithPassword) SubmitRegistrationFlowBody {
	return SubmitRegistrationFlowBody{
		SubmitRegistrationFlowBodyWithPassword: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubmitRegistrationFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SubmitRegistrationFlowBodyWithOidc
	err = newStrictDecoder(data).Decode(&dst.SubmitRegistrationFlowBodyWithOidc)
	if err == nil {
		jsonSubmitRegistrationFlowBodyWithOidc, _ := json.Marshal(dst.SubmitRegistrationFlowBodyWithOidc)
		if string(jsonSubmitRegistrationFlowBodyWithOidc) == "{}" { // empty struct
			dst.SubmitRegistrationFlowBodyWithOidc = nil
		} else {
			match++
		}
	} else {
		dst.SubmitRegistrationFlowBodyWithOidc = nil
	}

	// try to unmarshal data into SubmitRegistrationFlowBodyWithPassword
	err = newStrictDecoder(data).Decode(&dst.SubmitRegistrationFlowBodyWithPassword)
	if err == nil {
		jsonSubmitRegistrationFlowBodyWithPassword, _ := json.Marshal(dst.SubmitRegistrationFlowBodyWithPassword)
		if string(jsonSubmitRegistrationFlowBodyWithPassword) == "{}" { // empty struct
			dst.SubmitRegistrationFlowBodyWithPassword = nil
		} else {
			match++
		}
	} else {
		dst.SubmitRegistrationFlowBodyWithPassword = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SubmitRegistrationFlowBodyWithOidc = nil
		dst.SubmitRegistrationFlowBodyWithPassword = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SubmitRegistrationFlowBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SubmitRegistrationFlowBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubmitRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	if src.SubmitRegistrationFlowBodyWithOidc != nil {
		return json.Marshal(&src.SubmitRegistrationFlowBodyWithOidc)
	}

	if src.SubmitRegistrationFlowBodyWithPassword != nil {
		return json.Marshal(&src.SubmitRegistrationFlowBodyWithPassword)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubmitRegistrationFlowBody) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SubmitRegistrationFlowBodyWithOidc != nil {
		return obj.SubmitRegistrationFlowBodyWithOidc
	}

	if obj.SubmitRegistrationFlowBodyWithPassword != nil {
		return obj.SubmitRegistrationFlowBodyWithPassword
	}

	// all schemas are nil
	return nil
}

type NullableSubmitRegistrationFlowBody struct {
	value *SubmitRegistrationFlowBody
	isSet bool
}

func (v NullableSubmitRegistrationFlowBody) Get() *SubmitRegistrationFlowBody {
	return v.value
}

func (v *NullableSubmitRegistrationFlowBody) Set(val *SubmitRegistrationFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitRegistrationFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitRegistrationFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitRegistrationFlowBody(val *SubmitRegistrationFlowBody) *NullableSubmitRegistrationFlowBody {
	return &NullableSubmitRegistrationFlowBody{value: val, isSet: true}
}

func (v NullableSubmitRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitRegistrationFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


