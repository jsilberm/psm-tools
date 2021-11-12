/*
 * Monitoring API reference
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

// MonitoringAutoMsgFlowExportPolicyWatchHelper AutoMsgFlowExportPolicyWatchHelper is a wrapper object for watch events for FlowExportPolicy objects.
type MonitoringAutoMsgFlowExportPolicyWatchHelper struct {
	Events *[]MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewMonitoringAutoMsgFlowExportPolicyWatchHelper instantiates a new MonitoringAutoMsgFlowExportPolicyWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAutoMsgFlowExportPolicyWatchHelper() *MonitoringAutoMsgFlowExportPolicyWatchHelper {
	this := MonitoringAutoMsgFlowExportPolicyWatchHelper{}
	return &this
}

// NewMonitoringAutoMsgFlowExportPolicyWatchHelperWithDefaults instantiates a new MonitoringAutoMsgFlowExportPolicyWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAutoMsgFlowExportPolicyWatchHelperWithDefaults() *MonitoringAutoMsgFlowExportPolicyWatchHelper {
	this := MonitoringAutoMsgFlowExportPolicyWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *MonitoringAutoMsgFlowExportPolicyWatchHelper) GetEvents() []MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAutoMsgFlowExportPolicyWatchHelper) GetEventsOk() (*[]MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *MonitoringAutoMsgFlowExportPolicyWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent and assigns it to the Events field.
func (o *MonitoringAutoMsgFlowExportPolicyWatchHelper) SetEvents(v []MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent) {
	o.Events = &v
}

func (o MonitoringAutoMsgFlowExportPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringAutoMsgFlowExportPolicyWatchHelper struct {
	value *MonitoringAutoMsgFlowExportPolicyWatchHelper
	isSet bool
}

func (v NullableMonitoringAutoMsgFlowExportPolicyWatchHelper) Get() *MonitoringAutoMsgFlowExportPolicyWatchHelper {
	return v.value
}

func (v *NullableMonitoringAutoMsgFlowExportPolicyWatchHelper) Set(val *MonitoringAutoMsgFlowExportPolicyWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAutoMsgFlowExportPolicyWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAutoMsgFlowExportPolicyWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAutoMsgFlowExportPolicyWatchHelper(val *MonitoringAutoMsgFlowExportPolicyWatchHelper) *NullableMonitoringAutoMsgFlowExportPolicyWatchHelper {
	return &NullableMonitoringAutoMsgFlowExportPolicyWatchHelper{value: val, isSet: true}
}

func (v NullableMonitoringAutoMsgFlowExportPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAutoMsgFlowExportPolicyWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


