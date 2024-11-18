/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListResponseStreamSinkOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseStreamSinkOut{}

// ListResponseStreamSinkOut struct for ListResponseStreamSinkOut
type ListResponseStreamSinkOut struct {
	Data []StreamSinkOut `json:"data"`
	Done bool `json:"done"`
	Iterator NullableString `json:"iterator"`
	PrevIterator NullableString `json:"prevIterator,omitempty"`
}

type _ListResponseStreamSinkOut ListResponseStreamSinkOut

// NewListResponseStreamSinkOut instantiates a new ListResponseStreamSinkOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseStreamSinkOut(data []StreamSinkOut, done bool, iterator NullableString) *ListResponseStreamSinkOut {
	this := ListResponseStreamSinkOut{}
	this.Data = data
	this.Done = done
	this.Iterator = iterator
	return &this
}

// NewListResponseStreamSinkOutWithDefaults instantiates a new ListResponseStreamSinkOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseStreamSinkOutWithDefaults() *ListResponseStreamSinkOut {
	this := ListResponseStreamSinkOut{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseStreamSinkOut) GetData() []StreamSinkOut {
	if o == nil {
		var ret []StreamSinkOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseStreamSinkOut) GetDataOk() ([]StreamSinkOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseStreamSinkOut) SetData(v []StreamSinkOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *ListResponseStreamSinkOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseStreamSinkOut) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseStreamSinkOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListResponseStreamSinkOut) GetIterator() string {
	if o == nil || o.Iterator.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iterator.Get()
}

// GetIteratorOk returns a tuple with the Iterator field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseStreamSinkOut) GetIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iterator.Get(), o.Iterator.IsSet()
}

// SetIterator sets field value
func (o *ListResponseStreamSinkOut) SetIterator(v string) {
	o.Iterator.Set(&v)
}

// GetPrevIterator returns the PrevIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseStreamSinkOut) GetPrevIterator() string {
	if o == nil || IsNil(o.PrevIterator.Get()) {
		var ret string
		return ret
	}
	return *o.PrevIterator.Get()
}

// GetPrevIteratorOk returns a tuple with the PrevIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseStreamSinkOut) GetPrevIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevIterator.Get(), o.PrevIterator.IsSet()
}

// HasPrevIterator returns a boolean if a field has been set.
func (o *ListResponseStreamSinkOut) HasPrevIterator() bool {
	if o != nil && o.PrevIterator.IsSet() {
		return true
	}

	return false
}

// SetPrevIterator gets a reference to the given NullableString and assigns it to the PrevIterator field.
func (o *ListResponseStreamSinkOut) SetPrevIterator(v string) {
	o.PrevIterator.Set(&v)
}
// SetPrevIteratorNil sets the value for PrevIterator to be an explicit nil
func (o *ListResponseStreamSinkOut) SetPrevIteratorNil() {
	o.PrevIterator.Set(nil)
}

// UnsetPrevIterator ensures that no value is present for PrevIterator, not even an explicit nil
func (o *ListResponseStreamSinkOut) UnsetPrevIterator() {
	o.PrevIterator.Unset()
}

func (o ListResponseStreamSinkOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseStreamSinkOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["done"] = o.Done
	toSerialize["iterator"] = o.Iterator.Get()
	if o.PrevIterator.IsSet() {
		toSerialize["prevIterator"] = o.PrevIterator.Get()
	}
	return toSerialize, nil
}

func (o *ListResponseStreamSinkOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"done",
		"iterator",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListResponseStreamSinkOut := _ListResponseStreamSinkOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponseStreamSinkOut)

	if err != nil {
		return err
	}

	*o = ListResponseStreamSinkOut(varListResponseStreamSinkOut)

	return err
}

type NullableListResponseStreamSinkOut struct {
	value *ListResponseStreamSinkOut
	isSet bool
}

func (v NullableListResponseStreamSinkOut) Get() *ListResponseStreamSinkOut {
	return v.value
}

func (v *NullableListResponseStreamSinkOut) Set(val *ListResponseStreamSinkOut) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseStreamSinkOut) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseStreamSinkOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseStreamSinkOut(val *ListResponseStreamSinkOut) *NullableListResponseStreamSinkOut {
	return &NullableListResponseStreamSinkOut{value: val, isSet: true}
}

func (v NullableListResponseStreamSinkOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseStreamSinkOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


