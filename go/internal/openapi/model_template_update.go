/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TemplateUpdate struct for TemplateUpdate
type TemplateUpdate struct {
	Description *string `json:"description,omitempty"`
	FilterTypes []string `json:"filterTypes,omitempty"`
	Instructions *string `json:"instructions,omitempty"`
	InstructionsLink NullableString `json:"instructionsLink,omitempty"`
	Logo string `json:"logo"`
	Name *string `json:"name,omitempty"`
	Transformation string `json:"transformation"`
}

// NewTemplateUpdate instantiates a new TemplateUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateUpdate(logo string, transformation string) *TemplateUpdate {
	this := TemplateUpdate{}
	var description string = ""
	this.Description = &description
	var instructions string = ""
	this.Instructions = &instructions
	this.Logo = logo
	var name string = ""
	this.Name = &name
	this.Transformation = transformation
	return &this
}

// NewTemplateUpdateWithDefaults instantiates a new TemplateUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateUpdateWithDefaults() *TemplateUpdate {
	this := TemplateUpdate{}
	var description string = ""
	this.Description = &description
	var instructions string = ""
	this.Instructions = &instructions
	var name string = ""
	this.Name = &name
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TemplateUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TemplateUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetFilterTypes returns the FilterTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateUpdate) GetFilterTypes() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.FilterTypes
}

// GetFilterTypesOk returns a tuple with the FilterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateUpdate) GetFilterTypesOk() (*[]string, bool) {
	if o == nil || o.FilterTypes == nil {
		return nil, false
	}
	return &o.FilterTypes, true
}

// HasFilterTypes returns a boolean if a field has been set.
func (o *TemplateUpdate) HasFilterTypes() bool {
	if o != nil && o.FilterTypes != nil {
		return true
	}

	return false
}

// SetFilterTypes gets a reference to the given []string and assigns it to the FilterTypes field.
func (o *TemplateUpdate) SetFilterTypes(v []string) {
	o.FilterTypes = v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *TemplateUpdate) GetInstructions() string {
	if o == nil || o.Instructions == nil {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetInstructionsOk() (*string, bool) {
	if o == nil || o.Instructions == nil {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *TemplateUpdate) HasInstructions() bool {
	if o != nil && o.Instructions != nil {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *TemplateUpdate) SetInstructions(v string) {
	o.Instructions = &v
}

// GetInstructionsLink returns the InstructionsLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateUpdate) GetInstructionsLink() string {
	if o == nil || o.InstructionsLink.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstructionsLink.Get()
}

// GetInstructionsLinkOk returns a tuple with the InstructionsLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateUpdate) GetInstructionsLinkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstructionsLink.Get(), o.InstructionsLink.IsSet()
}

// HasInstructionsLink returns a boolean if a field has been set.
func (o *TemplateUpdate) HasInstructionsLink() bool {
	if o != nil && o.InstructionsLink.IsSet() {
		return true
	}

	return false
}

// SetInstructionsLink gets a reference to the given NullableString and assigns it to the InstructionsLink field.
func (o *TemplateUpdate) SetInstructionsLink(v string) {
	o.InstructionsLink.Set(&v)
}
// SetInstructionsLinkNil sets the value for InstructionsLink to be an explicit nil
func (o *TemplateUpdate) SetInstructionsLinkNil() {
	o.InstructionsLink.Set(nil)
}

// UnsetInstructionsLink ensures that no value is present for InstructionsLink, not even an explicit nil
func (o *TemplateUpdate) UnsetInstructionsLink() {
	o.InstructionsLink.Unset()
}

// GetLogo returns the Logo field value
func (o *TemplateUpdate) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetLogoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *TemplateUpdate) SetLogo(v string) {
	o.Logo = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateUpdate) SetName(v string) {
	o.Name = &v
}

// GetTransformation returns the Transformation field value
func (o *TemplateUpdate) GetTransformation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transformation
}

// GetTransformationOk returns a tuple with the Transformation field value
// and a boolean to check if the value has been set.
func (o *TemplateUpdate) GetTransformationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transformation, true
}

// SetTransformation sets field value
func (o *TemplateUpdate) SetTransformation(v string) {
	o.Transformation = v
}

func (o TemplateUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FilterTypes != nil {
		toSerialize["filterTypes"] = o.FilterTypes
	}
	if o.Instructions != nil {
		toSerialize["instructions"] = o.Instructions
	}
	if o.InstructionsLink.IsSet() {
		toSerialize["instructionsLink"] = o.InstructionsLink.Get()
	}
	if true {
		toSerialize["logo"] = o.Logo
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["transformation"] = o.Transformation
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateUpdate struct {
	value *TemplateUpdate
	isSet bool
}

func (v NullableTemplateUpdate) Get() *TemplateUpdate {
	return v.value
}

func (v *NullableTemplateUpdate) Set(val *TemplateUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateUpdate(val *TemplateUpdate) *NullableTemplateUpdate {
	return &NullableTemplateUpdate{value: val, isSet: true}
}

func (v NullableTemplateUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


