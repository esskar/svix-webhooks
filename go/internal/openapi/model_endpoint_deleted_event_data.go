/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EndpointDeletedEventData struct for EndpointDeletedEventData
type EndpointDeletedEventData struct {
	AppId string `json:"appId"`
	// Optional unique identifier for the application
	AppUid *string `json:"appUid,omitempty"`
	EndpointId string `json:"endpointId"`
}

// NewEndpointDeletedEventData instantiates a new EndpointDeletedEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointDeletedEventData(appId string, endpointId string) *EndpointDeletedEventData {
	this := EndpointDeletedEventData{}
	this.AppId = appId
	this.EndpointId = endpointId
	return &this
}

// NewEndpointDeletedEventDataWithDefaults instantiates a new EndpointDeletedEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointDeletedEventDataWithDefaults() *EndpointDeletedEventData {
	this := EndpointDeletedEventData{}
	return &this
}

// GetAppId returns the AppId field value
func (o *EndpointDeletedEventData) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *EndpointDeletedEventData) GetAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *EndpointDeletedEventData) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise.
func (o *EndpointDeletedEventData) GetAppUid() string {
	if o == nil || o.AppUid == nil {
		var ret string
		return ret
	}
	return *o.AppUid
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointDeletedEventData) GetAppUidOk() (*string, bool) {
	if o == nil || o.AppUid == nil {
		return nil, false
	}
	return o.AppUid, true
}

// HasAppUid returns a boolean if a field has been set.
func (o *EndpointDeletedEventData) HasAppUid() bool {
	if o != nil && o.AppUid != nil {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given string and assigns it to the AppUid field.
func (o *EndpointDeletedEventData) SetAppUid(v string) {
	o.AppUid = &v
}

// GetEndpointId returns the EndpointId field value
func (o *EndpointDeletedEventData) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *EndpointDeletedEventData) GetEndpointIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *EndpointDeletedEventData) SetEndpointId(v string) {
	o.EndpointId = v
}

func (o EndpointDeletedEventData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appId"] = o.AppId
	}
	if o.AppUid != nil {
		toSerialize["appUid"] = o.AppUid
	}
	if true {
		toSerialize["endpointId"] = o.EndpointId
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointDeletedEventData struct {
	value *EndpointDeletedEventData
	isSet bool
}

func (v NullableEndpointDeletedEventData) Get() *EndpointDeletedEventData {
	return v.value
}

func (v *NullableEndpointDeletedEventData) Set(val *EndpointDeletedEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointDeletedEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointDeletedEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointDeletedEventData(val *EndpointDeletedEventData) *NullableEndpointDeletedEventData {
	return &NullableEndpointDeletedEventData{value: val, isSet: true}
}

func (v NullableEndpointDeletedEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointDeletedEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


