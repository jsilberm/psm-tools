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

// NetworkIPAMPoolInfo struct for NetworkIPAMPoolInfo
type NetworkIPAMPoolInfo struct {
	// ending IP address of the address pool range.
	IpAddressEnd *string `json:"ip-address-end,omitempty"`
	// starting IP address of the address pool range.
	IpAddressStart *string `json:"ip-address-start,omitempty"`
}

// NewNetworkIPAMPoolInfo instantiates a new NetworkIPAMPoolInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkIPAMPoolInfo() *NetworkIPAMPoolInfo {
	this := NetworkIPAMPoolInfo{}
	return &this
}

// NewNetworkIPAMPoolInfoWithDefaults instantiates a new NetworkIPAMPoolInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkIPAMPoolInfoWithDefaults() *NetworkIPAMPoolInfo {
	this := NetworkIPAMPoolInfo{}
	return &this
}

// GetIpAddressEnd returns the IpAddressEnd field value if set, zero value otherwise.
func (o *NetworkIPAMPoolInfo) GetIpAddressEnd() string {
	if o == nil || o.IpAddressEnd == nil {
		var ret string
		return ret
	}
	return *o.IpAddressEnd
}

// GetIpAddressEndOk returns a tuple with the IpAddressEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMPoolInfo) GetIpAddressEndOk() (*string, bool) {
	if o == nil || o.IpAddressEnd == nil {
		return nil, false
	}
	return o.IpAddressEnd, true
}

// HasIpAddressEnd returns a boolean if a field has been set.
func (o *NetworkIPAMPoolInfo) HasIpAddressEnd() bool {
	if o != nil && o.IpAddressEnd != nil {
		return true
	}

	return false
}

// SetIpAddressEnd gets a reference to the given string and assigns it to the IpAddressEnd field.
func (o *NetworkIPAMPoolInfo) SetIpAddressEnd(v string) {
	o.IpAddressEnd = &v
}

// GetIpAddressStart returns the IpAddressStart field value if set, zero value otherwise.
func (o *NetworkIPAMPoolInfo) GetIpAddressStart() string {
	if o == nil || o.IpAddressStart == nil {
		var ret string
		return ret
	}
	return *o.IpAddressStart
}

// GetIpAddressStartOk returns a tuple with the IpAddressStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMPoolInfo) GetIpAddressStartOk() (*string, bool) {
	if o == nil || o.IpAddressStart == nil {
		return nil, false
	}
	return o.IpAddressStart, true
}

// HasIpAddressStart returns a boolean if a field has been set.
func (o *NetworkIPAMPoolInfo) HasIpAddressStart() bool {
	if o != nil && o.IpAddressStart != nil {
		return true
	}

	return false
}

// SetIpAddressStart gets a reference to the given string and assigns it to the IpAddressStart field.
func (o *NetworkIPAMPoolInfo) SetIpAddressStart(v string) {
	o.IpAddressStart = &v
}

func (o NetworkIPAMPoolInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddressEnd != nil {
		toSerialize["ip-address-end"] = o.IpAddressEnd
	}
	if o.IpAddressStart != nil {
		toSerialize["ip-address-start"] = o.IpAddressStart
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkIPAMPoolInfo struct {
	value *NetworkIPAMPoolInfo
	isSet bool
}

func (v NullableNetworkIPAMPoolInfo) Get() *NetworkIPAMPoolInfo {
	return v.value
}

func (v *NullableNetworkIPAMPoolInfo) Set(val *NetworkIPAMPoolInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkIPAMPoolInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkIPAMPoolInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkIPAMPoolInfo(val *NetworkIPAMPoolInfo) *NullableNetworkIPAMPoolInfo {
	return &NullableNetworkIPAMPoolInfo{value: val, isSet: true}
}

func (v NullableNetworkIPAMPoolInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkIPAMPoolInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


