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

// OuterComposite struct for OuterComposite
type OuterComposite struct {
	MyNumber *float32 `json:"my_number,omitempty"`
	MyString *string `json:"my_string,omitempty"`
	MyBoolean *bool `json:"my_boolean,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OuterComposite OuterComposite

// NewOuterComposite instantiates a new OuterComposite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOuterComposite() *OuterComposite {
	this := OuterComposite{}
	return &this
}

// NewOuterCompositeWithDefaults instantiates a new OuterComposite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOuterCompositeWithDefaults() *OuterComposite {
	this := OuterComposite{}
	return &this
}

// GetMyNumber returns the MyNumber field value if set, zero value otherwise.
func (o *OuterComposite) GetMyNumber() *float32 {
	if o == nil || o.MyNumber == nil {
		var ret *float32
		return ret
	}
	return o.MyNumber
}

// GetMyNumberOk returns a tuple with the MyNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OuterComposite) GetMyNumberOk() (*float32, bool) {
	if o == nil || o.MyNumber == nil {
		return nil, false
	}
	return o.MyNumber, true
}

// HasMyNumber returns a boolean if a field has been set.
func (o *OuterComposite) HasMyNumber() bool {
	if o != nil && o.MyNumber != nil {
		return true
	}

	return false
}

// SetMyNumber gets a reference to the given float32 and assigns it to the MyNumber field.
func (o *OuterComposite) SetMyNumber(v float32) {
	o.MyNumber = &v
}

// GetMyString returns the MyString field value if set, zero value otherwise.
func (o *OuterComposite) GetMyString() *string {
	if o == nil || o.MyString == nil {
		var ret *string
		return ret
	}
	return o.MyString
}

// GetMyStringOk returns a tuple with the MyString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OuterComposite) GetMyStringOk() (*string, bool) {
	if o == nil || o.MyString == nil {
		return nil, false
	}
	return o.MyString, true
}

// HasMyString returns a boolean if a field has been set.
func (o *OuterComposite) HasMyString() bool {
	if o != nil && o.MyString != nil {
		return true
	}

	return false
}

// SetMyString gets a reference to the given string and assigns it to the MyString field.
func (o *OuterComposite) SetMyString(v string) {
	o.MyString = &v
}

// GetMyBoolean returns the MyBoolean field value if set, zero value otherwise.
func (o *OuterComposite) GetMyBoolean() *bool {
	if o == nil || o.MyBoolean == nil {
		var ret *bool
		return ret
	}
	return o.MyBoolean
}

// GetMyBooleanOk returns a tuple with the MyBoolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OuterComposite) GetMyBooleanOk() (*bool, bool) {
	if o == nil || o.MyBoolean == nil {
		return nil, false
	}
	return o.MyBoolean, true
}

// HasMyBoolean returns a boolean if a field has been set.
func (o *OuterComposite) HasMyBoolean() bool {
	if o != nil && o.MyBoolean != nil {
		return true
	}

	return false
}

// SetMyBoolean gets a reference to the given bool and assigns it to the MyBoolean field.
func (o *OuterComposite) SetMyBoolean(v bool) {
	o.MyBoolean = &v
}

func (o OuterComposite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MyNumber != nil {
		toSerialize["my_number"] = o.MyNumber
	}
	if o.MyString != nil {
		toSerialize["my_string"] = o.MyString
	}
	if o.MyBoolean != nil {
		toSerialize["my_boolean"] = o.MyBoolean
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OuterComposite) UnmarshalJSON(bytes []byte) (err error) {
	varOuterComposite := _OuterComposite{}

	if err = json.Unmarshal(bytes, &varOuterComposite); err == nil {
		*o = OuterComposite(varOuterComposite)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "my_number")
		delete(additionalProperties, "my_string")
		delete(additionalProperties, "my_boolean")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOuterComposite struct {
	value *OuterComposite
	isSet bool
}

func (v NullableOuterComposite) Get() *OuterComposite {
	return v.value
}

func (v *NullableOuterComposite) Set(val *OuterComposite) {
	v.value = val
	v.isSet = true
}

func (v NullableOuterComposite) IsSet() bool {
	return v.isSet
}

func (v *NullableOuterComposite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOuterComposite(val *OuterComposite) *NullableOuterComposite {
	return &NullableOuterComposite{value: val, isSet: true}
}

func (v NullableOuterComposite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOuterComposite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


