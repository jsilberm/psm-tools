/*
 * Security API reference
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

// SecurityAutoMsgAppWatchHelperWatchEvent struct for SecurityAutoMsgAppWatchHelperWatchEvent
type SecurityAutoMsgAppWatchHelperWatchEvent struct {
	Object *SecurityApp `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewSecurityAutoMsgAppWatchHelperWatchEvent instantiates a new SecurityAutoMsgAppWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAutoMsgAppWatchHelperWatchEvent() *SecurityAutoMsgAppWatchHelperWatchEvent {
	this := SecurityAutoMsgAppWatchHelperWatchEvent{}
	return &this
}

// NewSecurityAutoMsgAppWatchHelperWatchEventWithDefaults instantiates a new SecurityAutoMsgAppWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAutoMsgAppWatchHelperWatchEventWithDefaults() *SecurityAutoMsgAppWatchHelperWatchEvent {
	this := SecurityAutoMsgAppWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *SecurityAutoMsgAppWatchHelperWatchEvent) GetObject() SecurityApp {
	if o == nil || o.Object == nil {
		var ret SecurityApp
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutoMsgAppWatchHelperWatchEvent) GetObjectOk() (*SecurityApp, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *SecurityAutoMsgAppWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given SecurityApp and assigns it to the Object field.
func (o *SecurityAutoMsgAppWatchHelperWatchEvent) SetObject(v SecurityApp) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityAutoMsgAppWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutoMsgAppWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityAutoMsgAppWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecurityAutoMsgAppWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o SecurityAutoMsgAppWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityAutoMsgAppWatchHelperWatchEvent struct {
	value *SecurityAutoMsgAppWatchHelperWatchEvent
	isSet bool
}

func (v NullableSecurityAutoMsgAppWatchHelperWatchEvent) Get() *SecurityAutoMsgAppWatchHelperWatchEvent {
	return v.value
}

func (v *NullableSecurityAutoMsgAppWatchHelperWatchEvent) Set(val *SecurityAutoMsgAppWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAutoMsgAppWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAutoMsgAppWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAutoMsgAppWatchHelperWatchEvent(val *SecurityAutoMsgAppWatchHelperWatchEvent) *NullableSecurityAutoMsgAppWatchHelperWatchEvent {
	return &NullableSecurityAutoMsgAppWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableSecurityAutoMsgAppWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAutoMsgAppWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


