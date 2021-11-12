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

// AuthUserPreferenceSpec Holds a byte stream of the user preference. The typing of the user preference will be in the UI code.
type AuthUserPreferenceSpec struct {
	Options *string `json:"options,omitempty"`
}

// NewAuthUserPreferenceSpec instantiates a new AuthUserPreferenceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthUserPreferenceSpec() *AuthUserPreferenceSpec {
	this := AuthUserPreferenceSpec{}
	return &this
}

// NewAuthUserPreferenceSpecWithDefaults instantiates a new AuthUserPreferenceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthUserPreferenceSpecWithDefaults() *AuthUserPreferenceSpec {
	this := AuthUserPreferenceSpec{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AuthUserPreferenceSpec) GetOptions() string {
	if o == nil || o.Options == nil {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthUserPreferenceSpec) GetOptionsOk() (*string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AuthUserPreferenceSpec) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *AuthUserPreferenceSpec) SetOptions(v string) {
	o.Options = &v
}

func (o AuthUserPreferenceSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableAuthUserPreferenceSpec struct {
	value *AuthUserPreferenceSpec
	isSet bool
}

func (v NullableAuthUserPreferenceSpec) Get() *AuthUserPreferenceSpec {
	return v.value
}

func (v *NullableAuthUserPreferenceSpec) Set(val *AuthUserPreferenceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthUserPreferenceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthUserPreferenceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthUserPreferenceSpec(val *AuthUserPreferenceSpec) *NullableAuthUserPreferenceSpec {
	return &NullableAuthUserPreferenceSpec{value: val, isSet: true}
}

func (v NullableAuthUserPreferenceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthUserPreferenceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


