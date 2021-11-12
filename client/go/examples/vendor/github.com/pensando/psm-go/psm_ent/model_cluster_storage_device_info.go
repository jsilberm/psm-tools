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

// ClusterStorageDeviceInfo Storage device information.
type ClusterStorageDeviceInfo struct {
	// Capacity in bytes.
	Capacity *string `json:"capacity,omitempty"`
	// Used life in percentage.
	PercentLifeUsedA *int32 `json:"percent-life-used-A,omitempty"`
	PercentLifeUsedB *int32 `json:"percent-life-used-B,omitempty"`
	// Serial Number.
	SerialNum *string `json:"serial-num,omitempty"`
	// Storage Type (TBD for Naples) Eg: SATA, SCSI, NVMe  or HDD, SSD, NVMe.
	Type *string `json:"type,omitempty"`
	// Vendor info.
	Vendor *string `json:"vendor,omitempty"`
}

// NewClusterStorageDeviceInfo instantiates a new ClusterStorageDeviceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStorageDeviceInfo() *ClusterStorageDeviceInfo {
	this := ClusterStorageDeviceInfo{}
	return &this
}

// NewClusterStorageDeviceInfoWithDefaults instantiates a new ClusterStorageDeviceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStorageDeviceInfoWithDefaults() *ClusterStorageDeviceInfo {
	this := ClusterStorageDeviceInfo{}
	return &this
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *ClusterStorageDeviceInfo) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageDeviceInfo) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *ClusterStorageDeviceInfo) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *ClusterStorageDeviceInfo) SetCapacity(v string) {
	o.Capacity = &v
}

// GetPercentLifeUsedA returns the PercentLifeUsedA field value if set, zero value otherwise.
func (o *ClusterStorageDeviceInfo) GetPercentLifeUsedA() int32 {
	if o == nil || o.PercentLifeUsedA == nil {
		var ret int32
		return ret
	}
	return *o.PercentLifeUsedA
}

// GetPercentLifeUsedAOk returns a tuple with the PercentLifeUsedA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageDeviceInfo) GetPercentLifeUsedAOk() (*int32, bool) {
	if o == nil || o.PercentLifeUsedA == nil {
		return nil, false
	}
	return o.PercentLifeUsedA, true
}

// HasPercentLifeUsedA returns a boolean if a field has been set.
func (o *ClusterStorageDeviceInfo) HasPercentLifeUsedA() bool {
	if o != nil && o.PercentLifeUsedA != nil {
		return true
	}

	return false
}

// SetPercentLifeUsedA gets a reference to the given int32 and assigns it to the PercentLifeUsedA field.
func (o *ClusterStorageDeviceInfo) SetPercentLifeUsedA(v int32) {
	o.PercentLifeUsedA = &v
}

// GetPercentLifeUsedB returns the PercentLifeUsedB field value if set, zero value otherwise.
func (o *ClusterStorageDeviceInfo) GetPercentLifeUsedB() int32 {
	if o == nil || o.PercentLifeUsedB == nil {
		var ret int32
		return ret
	}
	return *o.PercentLifeUsedB
}

// GetPercentLifeUsedBOk returns a tuple with the PercentLifeUsedB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageDeviceInfo) GetPercentLifeUsedBOk() (*int32, bool) {
	if o == nil || o.PercentLifeUsedB == nil {
		return nil, false
	}
	return o.PercentLifeUsedB, true
}

// HasPercentLifeUsedB returns a boolean if a field has been set.
func (o *ClusterStorageDeviceInfo) HasPercentLifeUsedB() bool {
	if o != nil && o.PercentLifeUsedB != nil {
		return true
	}

	return false
}

// SetPercentLifeUsedB gets a reference to the given int32 and assigns it to the PercentLifeUsedB field.
func (o *ClusterStorageDeviceInfo) SetPercentLifeUsedB(v int32) {
	o.PercentLifeUsedB = &v
}

// GetSerialNum returns the SerialNum field value if set, zero value otherwise.
func (o *ClusterStorageDeviceInfo) GetSerialNum() string {
	if o == nil || o.SerialNum == nil {
		var ret string
		return ret
	}
	return *o.SerialNum
}

// GetSerialNumOk returns a tuple with the SerialNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageDeviceInfo) GetSerialNumOk() (*string, bool) {
	if o == nil || o.SerialNum == nil {
		return nil, false
	}
	return o.SerialNum, true
}

// HasSerialNum returns a boolean if a field has been set.
func (o *ClusterStorageDeviceInfo) HasSerialNum() bool {
	if o != nil && o.SerialNum != nil {
		return true
	}

	return false
}

// SetSerialNum gets a reference to the given string and assigns it to the SerialNum field.
func (o *ClusterStorageDeviceInfo) SetSerialNum(v string) {
	o.SerialNum = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterStorageDeviceInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageDeviceInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterStorageDeviceInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterStorageDeviceInfo) SetType(v string) {
	o.Type = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ClusterStorageDeviceInfo) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageDeviceInfo) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ClusterStorageDeviceInfo) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *ClusterStorageDeviceInfo) SetVendor(v string) {
	o.Vendor = &v
}

func (o ClusterStorageDeviceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	if o.PercentLifeUsedA != nil {
		toSerialize["percent-life-used-A"] = o.PercentLifeUsedA
	}
	if o.PercentLifeUsedB != nil {
		toSerialize["percent-life-used-B"] = o.PercentLifeUsedB
	}
	if o.SerialNum != nil {
		toSerialize["serial-num"] = o.SerialNum
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableClusterStorageDeviceInfo struct {
	value *ClusterStorageDeviceInfo
	isSet bool
}

func (v NullableClusterStorageDeviceInfo) Get() *ClusterStorageDeviceInfo {
	return v.value
}

func (v *NullableClusterStorageDeviceInfo) Set(val *ClusterStorageDeviceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStorageDeviceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStorageDeviceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStorageDeviceInfo(val *ClusterStorageDeviceInfo) *NullableClusterStorageDeviceInfo {
	return &NullableClusterStorageDeviceInfo{value: val, isSet: true}
}

func (v NullableClusterStorageDeviceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStorageDeviceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


