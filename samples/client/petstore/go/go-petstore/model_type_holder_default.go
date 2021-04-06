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

// TypeHolderDefault struct for TypeHolderDefault
type TypeHolderDefault struct {
	StringItem string `json:"string_item"`
	NumberItem float32 `json:"number_item"`
	IntegerItem int32 `json:"integer_item"`
	BoolItem bool `json:"bool_item"`
	ArrayItem []int32 `json:"array_item"`
}

// NewTypeHolderDefault instantiates a new TypeHolderDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypeHolderDefault(stringItem string, numberItem float32, integerItem int32, boolItem bool, arrayItem []int32) *TypeHolderDefault {
	this := TypeHolderDefault{}
	this.StringItem = stringItem
	this.NumberItem = numberItem
	this.IntegerItem = integerItem
	this.BoolItem = boolItem
	this.ArrayItem = arrayItem
	return &this
}

// NewTypeHolderDefaultWithDefaults instantiates a new TypeHolderDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypeHolderDefaultWithDefaults() *TypeHolderDefault {
	this := TypeHolderDefault{}
	var stringItem string = "what"
	this.StringItem = stringItem
	var boolItem bool = true
	this.BoolItem = boolItem
	return &this
}

// GetStringItem returns the StringItem field value
func (o *TypeHolderDefault) GetStringItem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StringItem
}

// GetStringItemOk returns a tuple with the StringItem field value
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetStringItemOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StringItem, true
}

// SetStringItem sets field value
func (o *TypeHolderDefault) SetStringItem(v string) {
	o.StringItem = v
}

// GetNumberItem returns the NumberItem field value
func (o *TypeHolderDefault) GetNumberItem() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NumberItem
}

// GetNumberItemOk returns a tuple with the NumberItem field value
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetNumberItemOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumberItem, true
}

// SetNumberItem sets field value
func (o *TypeHolderDefault) SetNumberItem(v float32) {
	o.NumberItem = v
}

// GetIntegerItem returns the IntegerItem field value
func (o *TypeHolderDefault) GetIntegerItem() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntegerItem
}

// GetIntegerItemOk returns a tuple with the IntegerItem field value
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetIntegerItemOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IntegerItem, true
}

// SetIntegerItem sets field value
func (o *TypeHolderDefault) SetIntegerItem(v int32) {
	o.IntegerItem = v
}

// GetBoolItem returns the BoolItem field value
func (o *TypeHolderDefault) GetBoolItem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BoolItem
}

// GetBoolItemOk returns a tuple with the BoolItem field value
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetBoolItemOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BoolItem, true
}

// SetBoolItem sets field value
func (o *TypeHolderDefault) SetBoolItem(v bool) {
	o.BoolItem = v
}

// GetArrayItem returns the ArrayItem field value
func (o *TypeHolderDefault) GetArrayItem() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ArrayItem
}

// GetArrayItemOk returns a tuple with the ArrayItem field value
// and a boolean to check if the value has been set.
func (o *TypeHolderDefault) GetArrayItemOk() ([]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ArrayItem, true
}

// SetArrayItem sets field value
func (o *TypeHolderDefault) SetArrayItem(v []int32) {
	o.ArrayItem = v
}

func (o TypeHolderDefault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["string_item"] = o.StringItem
	}
	if true {
		toSerialize["number_item"] = o.NumberItem
	}
	if true {
		toSerialize["integer_item"] = o.IntegerItem
	}
	if true {
		toSerialize["bool_item"] = o.BoolItem
	}
	if true {
		toSerialize["array_item"] = o.ArrayItem
	}
	return json.Marshal(toSerialize)
}

type NullableTypeHolderDefault struct {
	value *TypeHolderDefault
	isSet bool
}

func (v NullableTypeHolderDefault) Get() *TypeHolderDefault {
	return v.value
}

func (v *NullableTypeHolderDefault) Set(val *TypeHolderDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeHolderDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeHolderDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeHolderDefault(val *TypeHolderDefault) *NullableTypeHolderDefault {
	return &NullableTypeHolderDefault{value: val, isSet: true}
}

func (v NullableTypeHolderDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeHolderDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


