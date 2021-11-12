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

// NetworkRouteDistinguisher struct for NetworkRouteDistinguisher
type NetworkRouteDistinguisher struct {
	// Administrator subfield of Value. Length depends on Type. swagger-tags: type=any.
	AdminValue *map[string]interface{} `json:"admin-value,omitempty"`
	// Assigned subfield of Value. Length depends on Type.
	AssignedValue *int64 `json:"assigned-value,omitempty"`
	// RD Type as in rfc4364.
	Type *string `json:"type,omitempty"`
}

// NewNetworkRouteDistinguisher instantiates a new NetworkRouteDistinguisher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRouteDistinguisher() *NetworkRouteDistinguisher {
	this := NetworkRouteDistinguisher{}
	var type_ string = "type0"
	this.Type = &type_
	return &this
}

// NewNetworkRouteDistinguisherWithDefaults instantiates a new NetworkRouteDistinguisher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRouteDistinguisherWithDefaults() *NetworkRouteDistinguisher {
	this := NetworkRouteDistinguisher{}
	var type_ string = "type0"
	this.Type = &type_
	return &this
}

// GetAdminValue returns the AdminValue field value if set, zero value otherwise.
func (o *NetworkRouteDistinguisher) GetAdminValue() map[string]interface{} {
	if o == nil || o.AdminValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdminValue
}

// GetAdminValueOk returns a tuple with the AdminValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouteDistinguisher) GetAdminValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdminValue == nil {
		return nil, false
	}
	return o.AdminValue, true
}

// HasAdminValue returns a boolean if a field has been set.
func (o *NetworkRouteDistinguisher) HasAdminValue() bool {
	if o != nil && o.AdminValue != nil {
		return true
	}

	return false
}

// SetAdminValue gets a reference to the given map[string]interface{} and assigns it to the AdminValue field.
func (o *NetworkRouteDistinguisher) SetAdminValue(v map[string]interface{}) {
	o.AdminValue = &v
}

// GetAssignedValue returns the AssignedValue field value if set, zero value otherwise.
func (o *NetworkRouteDistinguisher) GetAssignedValue() int64 {
	if o == nil || o.AssignedValue == nil {
		var ret int64
		return ret
	}
	return *o.AssignedValue
}

// GetAssignedValueOk returns a tuple with the AssignedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouteDistinguisher) GetAssignedValueOk() (*int64, bool) {
	if o == nil || o.AssignedValue == nil {
		return nil, false
	}
	return o.AssignedValue, true
}

// HasAssignedValue returns a boolean if a field has been set.
func (o *NetworkRouteDistinguisher) HasAssignedValue() bool {
	if o != nil && o.AssignedValue != nil {
		return true
	}

	return false
}

// SetAssignedValue gets a reference to the given int64 and assigns it to the AssignedValue field.
func (o *NetworkRouteDistinguisher) SetAssignedValue(v int64) {
	o.AssignedValue = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkRouteDistinguisher) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouteDistinguisher) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkRouteDistinguisher) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkRouteDistinguisher) SetType(v string) {
	o.Type = &v
}

func (o NetworkRouteDistinguisher) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminValue != nil {
		toSerialize["admin-value"] = o.AdminValue
	}
	if o.AssignedValue != nil {
		toSerialize["assigned-value"] = o.AssignedValue
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRouteDistinguisher struct {
	value *NetworkRouteDistinguisher
	isSet bool
}

func (v NullableNetworkRouteDistinguisher) Get() *NetworkRouteDistinguisher {
	return v.value
}

func (v *NullableNetworkRouteDistinguisher) Set(val *NetworkRouteDistinguisher) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRouteDistinguisher) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRouteDistinguisher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRouteDistinguisher(val *NetworkRouteDistinguisher) *NullableNetworkRouteDistinguisher {
	return &NullableNetworkRouteDistinguisher{value: val, isSet: true}
}

func (v NullableNetworkRouteDistinguisher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRouteDistinguisher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


