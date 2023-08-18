/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
	"fmt"
)

// PropertyDefinitionType - The type of the property's value.
type PropertyDefinitionType struct {
	RefPointer *RefPointer
	String *string
}

// RefPointerAsPropertyDefinitionType is a convenience function that returns RefPointer wrapped in PropertyDefinitionType
func RefPointerAsPropertyDefinitionType(v *RefPointer) PropertyDefinitionType {
	return PropertyDefinitionType{
		RefPointer: v,
	}
}

// stringAsPropertyDefinitionType is a convenience function that returns string wrapped in PropertyDefinitionType
func StringAsPropertyDefinitionType(v *string) PropertyDefinitionType {
	return PropertyDefinitionType{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PropertyDefinitionType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RefPointer
	err = newStrictDecoder(data).Decode(&dst.RefPointer)
	if err == nil {
		jsonRefPointer, _ := json.Marshal(dst.RefPointer)
		if string(jsonRefPointer) == "{}" { // empty struct
			dst.RefPointer = nil
		} else {
			match++
		}
	} else {
		dst.RefPointer = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RefPointer = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PropertyDefinitionType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PropertyDefinitionType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PropertyDefinitionType) MarshalJSON() ([]byte, error) {
	if src.RefPointer != nil {
		return json.Marshal(&src.RefPointer)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PropertyDefinitionType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RefPointer != nil {
		return obj.RefPointer
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullablePropertyDefinitionType struct {
	value *PropertyDefinitionType
	isSet bool
}

func (v NullablePropertyDefinitionType) Get() *PropertyDefinitionType {
	return v.value
}

func (v *NullablePropertyDefinitionType) Set(val *PropertyDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyDefinitionType(val *PropertyDefinitionType) *NullablePropertyDefinitionType {
	return &NullablePropertyDefinitionType{value: val, isSet: true}
}

func (v NullablePropertyDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

