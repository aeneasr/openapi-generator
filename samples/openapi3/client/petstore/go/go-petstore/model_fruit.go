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

// Fruit - struct for Fruit
type Fruit struct {
	Apple  *Apple
	Banana *Banana
}

// AppleAsFruit is a convenience function that returns Apple wrapped in Fruit
func AppleAsFruit(v *Apple) Fruit {
	return Fruit{
		Apple: v,
	}
}

// BananaAsFruit is a convenience function that returns Banana wrapped in Fruit
func BananaAsFruit(v *Banana) Fruit {
	return Fruit{
		Banana: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Fruit) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Apple
	err = newStrictDecoder(data).Decode(&dst.Apple)
	if err == nil {
		jsonApple, _ := json.Marshal(dst.Apple)
		if string(jsonApple) == "{}" { // empty struct
			dst.Apple = nil
		} else {
			match++
		}
	} else {
		dst.Apple = nil
	}

	// try to unmarshal data into Banana
	err = newStrictDecoder(data).Decode(&dst.Banana)
	if err == nil {
		jsonBanana, _ := json.Marshal(dst.Banana)
		if string(jsonBanana) == "{}" { // empty struct
			dst.Banana = nil
		} else {
			match++
		}
	} else {
		dst.Banana = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Apple = nil
		dst.Banana = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(Fruit)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(Fruit)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Fruit) MarshalJSON() ([]byte, error) {
	if src.Apple != nil {
		return json.Marshal(&src.Apple)
	}

	if src.Banana != nil {
		return json.Marshal(&src.Banana)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Fruit) GetActualInstance() interface{} {
	if obj.Apple != nil {
		return obj.Apple
	}

	if obj.Banana != nil {
		return obj.Banana
	}

	// all schemas are nil
	return nil
}

type NullableFruit struct {
	value *Fruit
	isSet bool
}

func (v NullableFruit) Get() *Fruit {
	return v.value
}

func (v *NullableFruit) Set(val *Fruit) {
	v.value = val
	v.isSet = true
}

func (v NullableFruit) IsSet() bool {
	return v.isSet
}

func (v *NullableFruit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFruit(val *Fruit) *NullableFruit {
	return &NullableFruit{value: val, isSet: true}
}

func (v NullableFruit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFruit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
