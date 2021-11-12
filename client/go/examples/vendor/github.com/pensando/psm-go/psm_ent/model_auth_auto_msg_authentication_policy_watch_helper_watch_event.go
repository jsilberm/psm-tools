/*
 * Auth API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent struct for AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent
type AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent struct {
	Object *AuthAuthenticationPolicy `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent instantiates a new AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent() *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent {
	this := AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent{}
	return &this
}

// NewAuthAutoMsgAuthenticationPolicyWatchHelperWatchEventWithDefaults instantiates a new AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAutoMsgAuthenticationPolicyWatchHelperWatchEventWithDefaults() *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent {
	this := AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) GetObject() AuthAuthenticationPolicy {
	if o == nil || o.Object == nil {
		var ret AuthAuthenticationPolicy
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) GetObjectOk() (*AuthAuthenticationPolicy, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given AuthAuthenticationPolicy and assigns it to the Object field.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) SetObject(v AuthAuthenticationPolicy) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent struct {
	value *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent
	isSet bool
}

func (v NullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) Get() *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent {
	return v.value
}

func (v *NullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) Set(val *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent(val *AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) *NullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent {
	return &NullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


