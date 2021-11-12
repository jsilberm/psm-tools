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

// NetworkNetworkInterfaceHostStatus NetworkInterfaceHostStatus is populated for PF and VF.
type NetworkNetworkInterfaceHostStatus struct {
	// PCIE Device ID.
	DeviceId *string `json:"device-id,omitempty"`
	// interface name seen by the host driver.
	HostIfname *string `json:"host-ifname,omitempty"`
	// mac address of the interface.
	MacAddress *string `json:"mac-address,omitempty"`
}

// NewNetworkNetworkInterfaceHostStatus instantiates a new NetworkNetworkInterfaceHostStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkNetworkInterfaceHostStatus() *NetworkNetworkInterfaceHostStatus {
	this := NetworkNetworkInterfaceHostStatus{}
	return &this
}

// NewNetworkNetworkInterfaceHostStatusWithDefaults instantiates a new NetworkNetworkInterfaceHostStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkNetworkInterfaceHostStatusWithDefaults() *NetworkNetworkInterfaceHostStatus {
	this := NetworkNetworkInterfaceHostStatus{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *NetworkNetworkInterfaceHostStatus) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkNetworkInterfaceHostStatus) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *NetworkNetworkInterfaceHostStatus) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *NetworkNetworkInterfaceHostStatus) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetHostIfname returns the HostIfname field value if set, zero value otherwise.
func (o *NetworkNetworkInterfaceHostStatus) GetHostIfname() string {
	if o == nil || o.HostIfname == nil {
		var ret string
		return ret
	}
	return *o.HostIfname
}

// GetHostIfnameOk returns a tuple with the HostIfname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkNetworkInterfaceHostStatus) GetHostIfnameOk() (*string, bool) {
	if o == nil || o.HostIfname == nil {
		return nil, false
	}
	return o.HostIfname, true
}

// HasHostIfname returns a boolean if a field has been set.
func (o *NetworkNetworkInterfaceHostStatus) HasHostIfname() bool {
	if o != nil && o.HostIfname != nil {
		return true
	}

	return false
}

// SetHostIfname gets a reference to the given string and assigns it to the HostIfname field.
func (o *NetworkNetworkInterfaceHostStatus) SetHostIfname(v string) {
	o.HostIfname = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *NetworkNetworkInterfaceHostStatus) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkNetworkInterfaceHostStatus) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *NetworkNetworkInterfaceHostStatus) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *NetworkNetworkInterfaceHostStatus) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o NetworkNetworkInterfaceHostStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["device-id"] = o.DeviceId
	}
	if o.HostIfname != nil {
		toSerialize["host-ifname"] = o.HostIfname
	}
	if o.MacAddress != nil {
		toSerialize["mac-address"] = o.MacAddress
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkNetworkInterfaceHostStatus struct {
	value *NetworkNetworkInterfaceHostStatus
	isSet bool
}

func (v NullableNetworkNetworkInterfaceHostStatus) Get() *NetworkNetworkInterfaceHostStatus {
	return v.value
}

func (v *NullableNetworkNetworkInterfaceHostStatus) Set(val *NetworkNetworkInterfaceHostStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkNetworkInterfaceHostStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkNetworkInterfaceHostStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkNetworkInterfaceHostStatus(val *NetworkNetworkInterfaceHostStatus) *NullableNetworkNetworkInterfaceHostStatus {
	return &NullableNetworkNetworkInterfaceHostStatus{value: val, isSet: true}
}

func (v NullableNetworkNetworkInterfaceHostStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkNetworkInterfaceHostStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


