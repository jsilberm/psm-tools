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

// MonitoringPropagationStatus struct for MonitoringPropagationStatus
type MonitoringPropagationStatus struct {
	// list of DSCs status.
	DscStatus *[]MonitoringDSCStatus `json:"dsc-status,omitempty"`
	// The Generation ID this status is for.
	GenerationId *string `json:"generation-id,omitempty"`
	// The Version running on the slowest DSC.
	MinVersion *string `json:"min-version,omitempty"`
	// Number of Naples pending. If this is 0 it can be assumed that everything is up to date.
	Pending *int32 `json:"pending,omitempty"`
	// list of DSCs where propagation did not complete.
	PendingDscs *[]string `json:"pending-dscs,omitempty"`
	// Textual description of propagation status.
	Status *string `json:"status,omitempty"`
	// The number of Naples that this version has already been pushed to.
	Updated *int32 `json:"updated,omitempty"`
}

// NewMonitoringPropagationStatus instantiates a new MonitoringPropagationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringPropagationStatus() *MonitoringPropagationStatus {
	this := MonitoringPropagationStatus{}
	return &this
}

// NewMonitoringPropagationStatusWithDefaults instantiates a new MonitoringPropagationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringPropagationStatusWithDefaults() *MonitoringPropagationStatus {
	this := MonitoringPropagationStatus{}
	return &this
}

// GetDscStatus returns the DscStatus field value if set, zero value otherwise.
func (o *MonitoringPropagationStatus) GetDscStatus() []MonitoringDSCStatus {
	if o == nil || o.DscStatus == nil {
		var ret []MonitoringDSCStatus
		return ret
	}
	return *o.DscStatus
}

// GetDscStatusOk returns a tuple with the DscStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringPropagationStatus) GetDscStatusOk() (*[]MonitoringDSCStatus, bool) {
	if o == nil || o.DscStatus == nil {
		return nil, false
	}
	return o.DscStatus, true
}

// HasDscStatus returns a boolean if a field has been set.
func (o *MonitoringPropagationStatus) HasDscStatus() bool {
	if o != nil && o.DscStatus != nil {
		return true
	}

	return false
}

// SetDscStatus gets a reference to the given []MonitoringDSCStatus and assigns it to the DscStatus field.
func (o *MonitoringPropagationStatus) SetDscStatus(v []MonitoringDSCStatus) {
	o.DscStatus = &v
}

// GetGenerationId returns the GenerationId field value if set, zero value otherwise.
func (o *MonitoringPropagationStatus) GetGenerationId() string {
	if o == nil || o.GenerationId == nil {
		var ret string
		return ret
	}
	return *o.GenerationId
}

// GetGenerationIdOk returns a tuple with the GenerationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringPropagationStatus) GetGenerationIdOk() (*string, bool) {
	if o == nil || o.GenerationId == nil {
		return nil, false
	}
	return o.GenerationId, true
}

// HasGenerationId returns a boolean if a field has been set.
func (o *MonitoringPropagationStatus) HasGenerationId() bool {
	if o != nil && o.GenerationId != nil {
		return true
	}

	return false
}

// SetGenerationId gets a reference to the given string and assigns it to the GenerationId field.
func (o *MonitoringPropagationStatus) SetGenerationId(v string) {
	o.GenerationId = &v
}

// GetMinVersion returns the MinVersion field value if set, zero value otherwise.
func (o *MonitoringPropagationStatus) GetMinVersion() string {
	if o == nil || o.MinVersion == nil {
		var ret string
		return ret
	}
	return *o.MinVersion
}

// GetMinVersionOk returns a tuple with the MinVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringPropagationStatus) GetMinVersionOk() (*string, bool) {
	if o == nil || o.MinVersion == nil {
		return nil, false
	}
	return o.MinVersion, true
}

// HasMinVersion returns a boolean if a field has been set.
func (o *MonitoringPropagationStatus) HasMinVersion() bool {
	if o != nil && o.MinVersion != nil {
		return true
	}

	return false
}

// SetMinVersion gets a reference to the given string and assigns it to the MinVersion field.
func (o *MonitoringPropagationStatus) SetMinVersion(v string) {
	o.MinVersion = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *MonitoringPropagationStatus) GetPending() int32 {
	if o == nil || o.Pending == nil {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringPropagationStatus) GetPendingOk() (*int32, bool) {
	if o == nil || o.Pending == nil {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *MonitoringPropagationStatus) HasPending() bool {
	if o != nil && o.Pending != nil {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *MonitoringPropagationStatus) SetPending(v int32) {
	o.Pending = &v
}

// GetPendingDscs returns the PendingDscs field value if set, zero value otherwise.
func (o *MonitoringPropagationStatus) GetPendingDscs() []string {
	if o == nil || o.PendingDscs == nil {
		var ret []string
		return ret
	}
	return *o.PendingDscs
}

// GetPendingDscsOk returns a tuple with the PendingDscs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringPropagationStatus) GetPendingDscsOk() (*[]string, bool) {
	if o == nil || o.PendingDscs == nil {
		return nil, false
	}
	return o.PendingDscs, true
}

// HasPendingDscs returns a boolean if a field has been set.
func (o *MonitoringPropagationStatus) HasPendingDscs() bool {
	if o != nil && o.PendingDscs != nil {
		return true
	}

	return false
}

// SetPendingDscs gets a reference to the given []string and assigns it to the PendingDscs field.
func (o *MonitoringPropagationStatus) SetPendingDscs(v []string) {
	o.PendingDscs = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitoringPropagationStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringPropagationStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitoringPropagationStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MonitoringPropagationStatus) SetStatus(v string) {
	o.Status = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *MonitoringPropagationStatus) GetUpdated() int32 {
	if o == nil || o.Updated == nil {
		var ret int32
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringPropagationStatus) GetUpdatedOk() (*int32, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *MonitoringPropagationStatus) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given int32 and assigns it to the Updated field.
func (o *MonitoringPropagationStatus) SetUpdated(v int32) {
	o.Updated = &v
}

func (o MonitoringPropagationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DscStatus != nil {
		toSerialize["dsc-status"] = o.DscStatus
	}
	if o.GenerationId != nil {
		toSerialize["generation-id"] = o.GenerationId
	}
	if o.MinVersion != nil {
		toSerialize["min-version"] = o.MinVersion
	}
	if o.Pending != nil {
		toSerialize["pending"] = o.Pending
	}
	if o.PendingDscs != nil {
		toSerialize["pending-dscs"] = o.PendingDscs
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringPropagationStatus struct {
	value *MonitoringPropagationStatus
	isSet bool
}

func (v NullableMonitoringPropagationStatus) Get() *MonitoringPropagationStatus {
	return v.value
}

func (v *NullableMonitoringPropagationStatus) Set(val *MonitoringPropagationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringPropagationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringPropagationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringPropagationStatus(val *MonitoringPropagationStatus) *NullableMonitoringPropagationStatus {
	return &NullableMonitoringPropagationStatus{value: val, isSet: true}
}

func (v NullableMonitoringPropagationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringPropagationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


