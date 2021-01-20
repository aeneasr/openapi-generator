/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// AdditionalPropertiesBoolean struct for AdditionalPropertiesBoolean
type AdditionalPropertiesBoolean struct {
	Name *string `json:"name,omitempty"`
}

// NewAdditionalPropertiesBoolean instantiates a new AdditionalPropertiesBoolean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalPropertiesBoolean() *AdditionalPropertiesBoolean {
	this := AdditionalPropertiesBoolean{}
	return &this
}

// NewAdditionalPropertiesBooleanWithDefaults instantiates a new AdditionalPropertiesBoolean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalPropertiesBooleanWithDefaults() *AdditionalPropertiesBoolean {
	this := AdditionalPropertiesBoolean{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdditionalPropertiesBoolean) GetName() *string {
	if o == nil || o.Name == nil {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalPropertiesBoolean) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdditionalPropertiesBoolean) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdditionalPropertiesBoolean) SetName(v string) {
	o.Name = &v
}

func (o AdditionalPropertiesBoolean) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalPropertiesBoolean struct {
	value *AdditionalPropertiesBoolean
	isSet bool
}

func (v NullableAdditionalPropertiesBoolean) Get() *AdditionalPropertiesBoolean {
	return v.value
}

func (v *NullableAdditionalPropertiesBoolean) Set(val *AdditionalPropertiesBoolean) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalPropertiesBoolean) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalPropertiesBoolean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalPropertiesBoolean(val *AdditionalPropertiesBoolean) *NullableAdditionalPropertiesBoolean {
	return &NullableAdditionalPropertiesBoolean{value: val, isSet: true}
}

func (v NullableAdditionalPropertiesBoolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalPropertiesBoolean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


