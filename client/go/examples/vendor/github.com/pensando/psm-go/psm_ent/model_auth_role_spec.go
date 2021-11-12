/*
 * Auth API reference
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

// AuthRoleSpec spec part of role object.
type AuthRoleSpec struct {
	// Permissions define actions allowed on resources. A resource can be an API Server object or an arbitrary API endpoint.
	Permissions *[]AuthPermission `json:"permissions,omitempty"`
}

// NewAuthRoleSpec instantiates a new AuthRoleSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRoleSpec() *AuthRoleSpec {
	this := AuthRoleSpec{}
	return &this
}

// NewAuthRoleSpecWithDefaults instantiates a new AuthRoleSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRoleSpecWithDefaults() *AuthRoleSpec {
	this := AuthRoleSpec{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *AuthRoleSpec) GetPermissions() []AuthPermission {
	if o == nil || o.Permissions == nil {
		var ret []AuthPermission
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRoleSpec) GetPermissionsOk() (*[]AuthPermission, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *AuthRoleSpec) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []AuthPermission and assigns it to the Permissions field.
func (o *AuthRoleSpec) SetPermissions(v []AuthPermission) {
	o.Permissions = &v
}

func (o AuthRoleSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableAuthRoleSpec struct {
	value *AuthRoleSpec
	isSet bool
}

func (v NullableAuthRoleSpec) Get() *AuthRoleSpec {
	return v.value
}

func (v *NullableAuthRoleSpec) Set(val *AuthRoleSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRoleSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRoleSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRoleSpec(val *AuthRoleSpec) *NullableAuthRoleSpec {
	return &NullableAuthRoleSpec{value: val, isSet: true}
}

func (v NullableAuthRoleSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRoleSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


