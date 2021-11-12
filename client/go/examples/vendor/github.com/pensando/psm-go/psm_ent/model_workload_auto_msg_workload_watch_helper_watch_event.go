/*
 * Workload API reference
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

// WorkloadAutoMsgWorkloadWatchHelperWatchEvent struct for WorkloadAutoMsgWorkloadWatchHelperWatchEvent
type WorkloadAutoMsgWorkloadWatchHelperWatchEvent struct {
	Object *WorkloadWorkload `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewWorkloadAutoMsgWorkloadWatchHelperWatchEvent instantiates a new WorkloadAutoMsgWorkloadWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadAutoMsgWorkloadWatchHelperWatchEvent() *WorkloadAutoMsgWorkloadWatchHelperWatchEvent {
	this := WorkloadAutoMsgWorkloadWatchHelperWatchEvent{}
	return &this
}

// NewWorkloadAutoMsgWorkloadWatchHelperWatchEventWithDefaults instantiates a new WorkloadAutoMsgWorkloadWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadAutoMsgWorkloadWatchHelperWatchEventWithDefaults() *WorkloadAutoMsgWorkloadWatchHelperWatchEvent {
	this := WorkloadAutoMsgWorkloadWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) GetObject() WorkloadWorkload {
	if o == nil || o.Object == nil {
		var ret WorkloadWorkload
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) GetObjectOk() (*WorkloadWorkload, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given WorkloadWorkload and assigns it to the Object field.
func (o *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) SetObject(v WorkloadWorkload) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o WorkloadAutoMsgWorkloadWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent struct {
	value *WorkloadAutoMsgWorkloadWatchHelperWatchEvent
	isSet bool
}

func (v NullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent) Get() *WorkloadAutoMsgWorkloadWatchHelperWatchEvent {
	return v.value
}

func (v *NullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent) Set(val *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent(val *WorkloadAutoMsgWorkloadWatchHelperWatchEvent) *NullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent {
	return &NullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadAutoMsgWorkloadWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


