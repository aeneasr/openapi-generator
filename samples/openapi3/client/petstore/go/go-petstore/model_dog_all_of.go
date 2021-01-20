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

// DogAllOf struct for DogAllOf
type DogAllOf struct {
	Breed *string `json:"breed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DogAllOf DogAllOf

// NewDogAllOf instantiates a new DogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDogAllOf() *DogAllOf {
	this := DogAllOf{}
	return &this
}

// NewDogAllOfWithDefaults instantiates a new DogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDogAllOfWithDefaults() *DogAllOf {
	this := DogAllOf{}
	return &this
}

// GetBreed returns the Breed field value if set, zero value otherwise.
func (o *DogAllOf) GetBreed() *string {
	if o == nil || o.Breed == nil {
		var ret *string
		return ret
	}
	return o.Breed
}

// GetBreedOk returns a tuple with the Breed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DogAllOf) GetBreedOk() (*string, bool) {
	if o == nil || o.Breed == nil {
		return nil, false
	}
	return o.Breed, true
}

// HasBreed returns a boolean if a field has been set.
func (o *DogAllOf) HasBreed() bool {
	if o != nil && o.Breed != nil {
		return true
	}

	return false
}

// SetBreed gets a reference to the given string and assigns it to the Breed field.
func (o *DogAllOf) SetBreed(v string) {
	o.Breed = &v
}

func (o DogAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Breed != nil {
		toSerialize["breed"] = o.Breed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DogAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDogAllOf := _DogAllOf{}

	if err = json.Unmarshal(bytes, &varDogAllOf); err == nil {
		*o = DogAllOf(varDogAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "breed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDogAllOf struct {
	value *DogAllOf
	isSet bool
}

func (v NullableDogAllOf) Get() *DogAllOf {
	return v.value
}

func (v *NullableDogAllOf) Set(val *DogAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDogAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDogAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDogAllOf(val *DogAllOf) *NullableDogAllOf {
	return &NullableDogAllOf{value: val, isSet: true}
}

func (v NullableDogAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDogAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


