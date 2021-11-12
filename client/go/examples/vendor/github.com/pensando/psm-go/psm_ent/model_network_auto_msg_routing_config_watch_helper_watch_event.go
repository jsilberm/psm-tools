/*
 * Network API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// NetworkAutoMsgRoutingConfigWatchHelperWatchEvent struct for NetworkAutoMsgRoutingConfigWatchHelperWatchEvent
type NetworkAutoMsgRoutingConfigWatchHelperWatchEvent struct {
	Object *NetworkRoutingConfig `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewNetworkAutoMsgRoutingConfigWatchHelperWatchEvent instantiates a new NetworkAutoMsgRoutingConfigWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAutoMsgRoutingConfigWatchHelperWatchEvent() *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent {
	this := NetworkAutoMsgRoutingConfigWatchHelperWatchEvent{}
	return &this
}

// NewNetworkAutoMsgRoutingConfigWatchHelperWatchEventWithDefaults instantiates a new NetworkAutoMsgRoutingConfigWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAutoMsgRoutingConfigWatchHelperWatchEventWithDefaults() *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent {
	this := NetworkAutoMsgRoutingConfigWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) GetObject() NetworkRoutingConfig {
	if o == nil || o.Object == nil {
		var ret NetworkRoutingConfig
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) GetObjectOk() (*NetworkRoutingConfig, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given NetworkRoutingConfig and assigns it to the Object field.
func (o *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) SetObject(v NetworkRoutingConfig) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent struct {
	value *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent
	isSet bool
}

func (v NullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent) Get() *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent {
	return v.value
}

func (v *NullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent) Set(val *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent(val *NetworkAutoMsgRoutingConfigWatchHelperWatchEvent) *NullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent {
	return &NullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAutoMsgRoutingConfigWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


