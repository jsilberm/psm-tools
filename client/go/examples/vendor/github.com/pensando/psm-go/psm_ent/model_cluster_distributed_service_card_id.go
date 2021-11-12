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

// ClusterDistributedServiceCardID DistributedServiceCardID contains the identifying information of a DistributedServiceCard.
type ClusterDistributedServiceCardID struct {
	// Name contains the name of the DistributedServiceCard on a host.
	Id *string `json:"id,omitempty"`
	// MACAddress contains the primary MAC address of a DistributedServiceCard. Should be a valid MAC address.
	MacAddress *string `json:"mac-address,omitempty"`
}

// NewClusterDistributedServiceCardID instantiates a new ClusterDistributedServiceCardID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDistributedServiceCardID() *ClusterDistributedServiceCardID {
	this := ClusterDistributedServiceCardID{}
	return &this
}

// NewClusterDistributedServiceCardIDWithDefaults instantiates a new ClusterDistributedServiceCardID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDistributedServiceCardIDWithDefaults() *ClusterDistributedServiceCardID {
	this := ClusterDistributedServiceCardID{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCardID) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCardID) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCardID) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterDistributedServiceCardID) SetId(v string) {
	o.Id = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCardID) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCardID) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCardID) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *ClusterDistributedServiceCardID) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o ClusterDistributedServiceCardID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MacAddress != nil {
		toSerialize["mac-address"] = o.MacAddress
	}
	return json.Marshal(toSerialize)
}

type NullableClusterDistributedServiceCardID struct {
	value *ClusterDistributedServiceCardID
	isSet bool
}

func (v NullableClusterDistributedServiceCardID) Get() *ClusterDistributedServiceCardID {
	return v.value
}

func (v *NullableClusterDistributedServiceCardID) Set(val *ClusterDistributedServiceCardID) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDistributedServiceCardID) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDistributedServiceCardID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDistributedServiceCardID(val *ClusterDistributedServiceCardID) *NullableClusterDistributedServiceCardID {
	return &NullableClusterDistributedServiceCardID{value: val, isSet: true}
}

func (v NullableClusterDistributedServiceCardID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDistributedServiceCardID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


