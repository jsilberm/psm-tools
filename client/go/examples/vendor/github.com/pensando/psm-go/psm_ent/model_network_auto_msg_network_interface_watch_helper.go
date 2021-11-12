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

// NetworkAutoMsgNetworkInterfaceWatchHelper AutoMsgNetworkInterfaceWatchHelper is a wrapper object for watch events for NetworkInterface objects.
type NetworkAutoMsgNetworkInterfaceWatchHelper struct {
	Events *[]NetworkAutoMsgNetworkInterfaceWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewNetworkAutoMsgNetworkInterfaceWatchHelper instantiates a new NetworkAutoMsgNetworkInterfaceWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAutoMsgNetworkInterfaceWatchHelper() *NetworkAutoMsgNetworkInterfaceWatchHelper {
	this := NetworkAutoMsgNetworkInterfaceWatchHelper{}
	return &this
}

// NewNetworkAutoMsgNetworkInterfaceWatchHelperWithDefaults instantiates a new NetworkAutoMsgNetworkInterfaceWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAutoMsgNetworkInterfaceWatchHelperWithDefaults() *NetworkAutoMsgNetworkInterfaceWatchHelper {
	this := NetworkAutoMsgNetworkInterfaceWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *NetworkAutoMsgNetworkInterfaceWatchHelper) GetEvents() []NetworkAutoMsgNetworkInterfaceWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []NetworkAutoMsgNetworkInterfaceWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgNetworkInterfaceWatchHelper) GetEventsOk() (*[]NetworkAutoMsgNetworkInterfaceWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *NetworkAutoMsgNetworkInterfaceWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []NetworkAutoMsgNetworkInterfaceWatchHelperWatchEvent and assigns it to the Events field.
func (o *NetworkAutoMsgNetworkInterfaceWatchHelper) SetEvents(v []NetworkAutoMsgNetworkInterfaceWatchHelperWatchEvent) {
	o.Events = &v
}

func (o NetworkAutoMsgNetworkInterfaceWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAutoMsgNetworkInterfaceWatchHelper struct {
	value *NetworkAutoMsgNetworkInterfaceWatchHelper
	isSet bool
}

func (v NullableNetworkAutoMsgNetworkInterfaceWatchHelper) Get() *NetworkAutoMsgNetworkInterfaceWatchHelper {
	return v.value
}

func (v *NullableNetworkAutoMsgNetworkInterfaceWatchHelper) Set(val *NetworkAutoMsgNetworkInterfaceWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAutoMsgNetworkInterfaceWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAutoMsgNetworkInterfaceWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAutoMsgNetworkInterfaceWatchHelper(val *NetworkAutoMsgNetworkInterfaceWatchHelper) *NullableNetworkAutoMsgNetworkInterfaceWatchHelper {
	return &NullableNetworkAutoMsgNetworkInterfaceWatchHelper{value: val, isSet: true}
}

func (v NullableNetworkAutoMsgNetworkInterfaceWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAutoMsgNetworkInterfaceWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


