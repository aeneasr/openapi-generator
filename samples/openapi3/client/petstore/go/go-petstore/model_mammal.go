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
	"fmt"
)

// Mammal - struct for Mammal
type Mammal struct {
	Whale *Whale
	Zebra *Zebra
}

// WhaleAsMammal is a convenience function that returns Whale wrapped in Mammal
func WhaleAsMammal(v *Whale) Mammal {
	return Mammal{
		Whale: v,
	}
}

// ZebraAsMammal is a convenience function that returns Zebra wrapped in Mammal
func ZebraAsMammal(v *Zebra) Mammal {
	return Mammal{
		Zebra: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Mammal) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Whale
	err = json.Unmarshal(data, &dst.Whale)
	if err == nil {
		jsonWhale, _ := json.Marshal(dst.Whale)
		if string(jsonWhale) == "{}" { // empty struct
			dst.Whale = nil
		} else {
			match++
		}
	} else {
		dst.Whale = nil
	}

	// try to unmarshal data into Zebra
	err = json.Unmarshal(data, &dst.Zebra)
	if err == nil {
		jsonZebra, _ := json.Marshal(dst.Zebra)
		if string(jsonZebra) == "{}" { // empty struct
			dst.Zebra = nil
		} else {
			match++
		}
	} else {
		dst.Zebra = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Whale = nil
		dst.Zebra = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(Mammal)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(Mammal)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Mammal) MarshalJSON() ([]byte, error) {
	if src.Whale != nil {
		return json.Marshal(&src.Whale)
	}

	if src.Zebra != nil {
		return json.Marshal(&src.Zebra)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Mammal) GetActualInstance() (interface{}) {
	if obj.Whale != nil {
		return obj.Whale
	}

	if obj.Zebra != nil {
		return obj.Zebra
	}

	// all schemas are nil
	return nil
}

type NullableMammal struct {
	value *Mammal
	isSet bool
}

func (v NullableMammal) Get() *Mammal {
	return v.value
}

func (v *NullableMammal) Set(val *Mammal) {
	v.value = val
	v.isSet = true
}

func (v NullableMammal) IsSet() bool {
	return v.isSet
}

func (v *NullableMammal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMammal(val *Mammal) *NullableMammal {
	return &NullableMammal{value: val, isSet: true}
}

func (v NullableMammal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMammal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


