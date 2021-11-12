/*
 * Cluster API reference
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

// ClusterFault Captures any fault at DSS end.
type ClusterFault struct {
	// Description of the fault occured.
	Description *string `json:"description,omitempty"`
	// Time at which the fault occured.
	LastOccuredTime *string `json:"last-occured-time,omitempty"`
	// Mitigation action,if any possible.
	Mitigation *string `json:"mitigation,omitempty"`
	// Severity of fault occured at DSS end.
	Severity *string `json:"severity,omitempty"`
}

// NewClusterFault instantiates a new ClusterFault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFault() *ClusterFault {
	this := ClusterFault{}
	var mitigation string = "system-reboot"
	this.Mitigation = &mitigation
	var severity string = "info"
	this.Severity = &severity
	return &this
}

// NewClusterFaultWithDefaults instantiates a new ClusterFault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFaultWithDefaults() *ClusterFault {
	this := ClusterFault{}
	var mitigation string = "system-reboot"
	this.Mitigation = &mitigation
	var severity string = "info"
	this.Severity = &severity
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterFault) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFault) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterFault) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterFault) SetDescription(v string) {
	o.Description = &v
}

// GetLastOccuredTime returns the LastOccuredTime field value if set, zero value otherwise.
func (o *ClusterFault) GetLastOccuredTime() string {
	if o == nil || o.LastOccuredTime == nil {
		var ret string
		return ret
	}
	return *o.LastOccuredTime
}

// GetLastOccuredTimeOk returns a tuple with the LastOccuredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFault) GetLastOccuredTimeOk() (*string, bool) {
	if o == nil || o.LastOccuredTime == nil {
		return nil, false
	}
	return o.LastOccuredTime, true
}

// HasLastOccuredTime returns a boolean if a field has been set.
func (o *ClusterFault) HasLastOccuredTime() bool {
	if o != nil && o.LastOccuredTime != nil {
		return true
	}

	return false
}

// SetLastOccuredTime gets a reference to the given string and assigns it to the LastOccuredTime field.
func (o *ClusterFault) SetLastOccuredTime(v string) {
	o.LastOccuredTime = &v
}

// GetMitigation returns the Mitigation field value if set, zero value otherwise.
func (o *ClusterFault) GetMitigation() string {
	if o == nil || o.Mitigation == nil {
		var ret string
		return ret
	}
	return *o.Mitigation
}

// GetMitigationOk returns a tuple with the Mitigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFault) GetMitigationOk() (*string, bool) {
	if o == nil || o.Mitigation == nil {
		return nil, false
	}
	return o.Mitigation, true
}

// HasMitigation returns a boolean if a field has been set.
func (o *ClusterFault) HasMitigation() bool {
	if o != nil && o.Mitigation != nil {
		return true
	}

	return false
}

// SetMitigation gets a reference to the given string and assigns it to the Mitigation field.
func (o *ClusterFault) SetMitigation(v string) {
	o.Mitigation = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ClusterFault) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFault) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ClusterFault) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ClusterFault) SetSeverity(v string) {
	o.Severity = &v
}

func (o ClusterFault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.LastOccuredTime != nil {
		toSerialize["last-occured-time"] = o.LastOccuredTime
	}
	if o.Mitigation != nil {
		toSerialize["mitigation"] = o.Mitigation
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	return json.Marshal(toSerialize)
}

type NullableClusterFault struct {
	value *ClusterFault
	isSet bool
}

func (v NullableClusterFault) Get() *ClusterFault {
	return v.value
}

func (v *NullableClusterFault) Set(val *ClusterFault) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFault) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFault(val *ClusterFault) *NullableClusterFault {
	return &NullableClusterFault{value: val, isSet: true}
}

func (v NullableClusterFault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


