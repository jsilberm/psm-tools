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

// MonitoringEventPolicySpec EventPolicySpec is the specification of an Event Policy.
type MonitoringEventPolicySpec struct {
	Config *MonitoringSyslogExportConfig `json:"config,omitempty"`
	// event export format, SYSLOG_BSD default.
	Format *string `json:"format,omitempty"`
	Selector *FieldsSelector `json:"selector,omitempty"`
	// export target ip/port/protocol.
	Targets *[]MonitoringExportConfig `json:"targets,omitempty"`
}

// NewMonitoringEventPolicySpec instantiates a new MonitoringEventPolicySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringEventPolicySpec() *MonitoringEventPolicySpec {
	this := MonitoringEventPolicySpec{}
	var format string = "syslog-bsd"
	this.Format = &format
	return &this
}

// NewMonitoringEventPolicySpecWithDefaults instantiates a new MonitoringEventPolicySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringEventPolicySpecWithDefaults() *MonitoringEventPolicySpec {
	this := MonitoringEventPolicySpec{}
	var format string = "syslog-bsd"
	this.Format = &format
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *MonitoringEventPolicySpec) GetConfig() MonitoringSyslogExportConfig {
	if o == nil || o.Config == nil {
		var ret MonitoringSyslogExportConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventPolicySpec) GetConfigOk() (*MonitoringSyslogExportConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *MonitoringEventPolicySpec) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given MonitoringSyslogExportConfig and assigns it to the Config field.
func (o *MonitoringEventPolicySpec) SetConfig(v MonitoringSyslogExportConfig) {
	o.Config = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *MonitoringEventPolicySpec) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventPolicySpec) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MonitoringEventPolicySpec) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *MonitoringEventPolicySpec) SetFormat(v string) {
	o.Format = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *MonitoringEventPolicySpec) GetSelector() FieldsSelector {
	if o == nil || o.Selector == nil {
		var ret FieldsSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventPolicySpec) GetSelectorOk() (*FieldsSelector, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *MonitoringEventPolicySpec) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given FieldsSelector and assigns it to the Selector field.
func (o *MonitoringEventPolicySpec) SetSelector(v FieldsSelector) {
	o.Selector = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *MonitoringEventPolicySpec) GetTargets() []MonitoringExportConfig {
	if o == nil || o.Targets == nil {
		var ret []MonitoringExportConfig
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEventPolicySpec) GetTargetsOk() (*[]MonitoringExportConfig, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *MonitoringEventPolicySpec) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []MonitoringExportConfig and assigns it to the Targets field.
func (o *MonitoringEventPolicySpec) SetTargets(v []MonitoringExportConfig) {
	o.Targets = &v
}

func (o MonitoringEventPolicySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringEventPolicySpec struct {
	value *MonitoringEventPolicySpec
	isSet bool
}

func (v NullableMonitoringEventPolicySpec) Get() *MonitoringEventPolicySpec {
	return v.value
}

func (v *NullableMonitoringEventPolicySpec) Set(val *MonitoringEventPolicySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringEventPolicySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringEventPolicySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringEventPolicySpec(val *MonitoringEventPolicySpec) *NullableMonitoringEventPolicySpec {
	return &NullableMonitoringEventPolicySpec{value: val, isSet: true}
}

func (v NullableMonitoringEventPolicySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringEventPolicySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

