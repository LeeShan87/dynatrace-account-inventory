/*
Dynatrace Configuration API

Documentation of the Dynatrace Configuration API. To read about use-cases and examples, see [Dynatrace Documentation](https://dt-url.net/4u43kxw).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configv1

import (
	"encoding/json"
)

// checks if the MarkdownTile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarkdownTile{}

// MarkdownTile Configuration of the Markdown tile.
type MarkdownTile struct {
	// The markdown-formatted content of the tile.
	Markdown string `json:"markdown"`
}

// NewMarkdownTile instantiates a new MarkdownTile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkdownTile(markdown string, bounds TileBounds, name string, tileType string) *MarkdownTile {
	this := MarkdownTile{}
	this.Bounds = bounds
	this.Name = name
	this.TileType = tileType
	this.Markdown = markdown
	return &this
}

// NewMarkdownTileWithDefaults instantiates a new MarkdownTile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkdownTileWithDefaults() *MarkdownTile {
	this := MarkdownTile{}
	return &this
}

// GetMarkdown returns the Markdown field value
func (o *MarkdownTile) GetMarkdown() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Markdown
}

// GetMarkdownOk returns a tuple with the Markdown field value
// and a boolean to check if the value has been set.
func (o *MarkdownTile) GetMarkdownOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Markdown, true
}

// SetMarkdown sets field value
func (o *MarkdownTile) SetMarkdown(v string) {
	o.Markdown = v
}

func (o MarkdownTile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkdownTile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["markdown"] = o.Markdown
	return toSerialize, nil
}

type NullableMarkdownTile struct {
	value *MarkdownTile
	isSet bool
}

func (v NullableMarkdownTile) Get() *MarkdownTile {
	return v.value
}

func (v *NullableMarkdownTile) Set(val *MarkdownTile) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkdownTile) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkdownTile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkdownTile(val *MarkdownTile) *NullableMarkdownTile {
	return &NullableMarkdownTile{value: val, isSet: true}
}

func (v NullableMarkdownTile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkdownTile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


