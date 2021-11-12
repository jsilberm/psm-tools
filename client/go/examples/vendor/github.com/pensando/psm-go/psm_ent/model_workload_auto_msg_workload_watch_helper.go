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

// WorkloadAutoMsgWorkloadWatchHelper AutoMsgWorkloadWatchHelper is a wrapper object for watch events for Workload objects.
type WorkloadAutoMsgWorkloadWatchHelper struct {
	Events *[]WorkloadAutoMsgWorkloadWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewWorkloadAutoMsgWorkloadWatchHelper instantiates a new WorkloadAutoMsgWorkloadWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadAutoMsgWorkloadWatchHelper() *WorkloadAutoMsgWorkloadWatchHelper {
	this := WorkloadAutoMsgWorkloadWatchHelper{}
	return &this
}

// NewWorkloadAutoMsgWorkloadWatchHelperWithDefaults instantiates a new WorkloadAutoMsgWorkloadWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadAutoMsgWorkloadWatchHelperWithDefaults() *WorkloadAutoMsgWorkloadWatchHelper {
	this := WorkloadAutoMsgWorkloadWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *WorkloadAutoMsgWorkloadWatchHelper) GetEvents() []WorkloadAutoMsgWorkloadWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []WorkloadAutoMsgWorkloadWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadAutoMsgWorkloadWatchHelper) GetEventsOk() (*[]WorkloadAutoMsgWorkloadWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *WorkloadAutoMsgWorkloadWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []WorkloadAutoMsgWorkloadWatchHelperWatchEvent and assigns it to the Events field.
func (o *WorkloadAutoMsgWorkloadWatchHelper) SetEvents(v []WorkloadAutoMsgWorkloadWatchHelperWatchEvent) {
	o.Events = &v
}

func (o WorkloadAutoMsgWorkloadWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableWorkloadAutoMsgWorkloadWatchHelper struct {
	value *WorkloadAutoMsgWorkloadWatchHelper
	isSet bool
}

func (v NullableWorkloadAutoMsgWorkloadWatchHelper) Get() *WorkloadAutoMsgWorkloadWatchHelper {
	return v.value
}

func (v *NullableWorkloadAutoMsgWorkloadWatchHelper) Set(val *WorkloadAutoMsgWorkloadWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadAutoMsgWorkloadWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadAutoMsgWorkloadWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadAutoMsgWorkloadWatchHelper(val *WorkloadAutoMsgWorkloadWatchHelper) *NullableWorkloadAutoMsgWorkloadWatchHelper {
	return &NullableWorkloadAutoMsgWorkloadWatchHelper{value: val, isSet: true}
}

func (v NullableWorkloadAutoMsgWorkloadWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadAutoMsgWorkloadWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


