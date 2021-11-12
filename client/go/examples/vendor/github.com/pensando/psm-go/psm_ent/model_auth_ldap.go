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

// AuthLdap struct for AuthLdap
type AuthLdap struct {
	Domains *[]AuthLdapDomain `json:"domains,omitempty"`
}

// NewAuthLdap instantiates a new AuthLdap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthLdap() *AuthLdap {
	this := AuthLdap{}
	return &this
}

// NewAuthLdapWithDefaults instantiates a new AuthLdap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthLdapWithDefaults() *AuthLdap {
	this := AuthLdap{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *AuthLdap) GetDomains() []AuthLdapDomain {
	if o == nil || o.Domains == nil {
		var ret []AuthLdapDomain
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdap) GetDomainsOk() (*[]AuthLdapDomain, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *AuthLdap) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []AuthLdapDomain and assigns it to the Domains field.
func (o *AuthLdap) SetDomains(v []AuthLdapDomain) {
	o.Domains = &v
}

func (o AuthLdap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	return json.Marshal(toSerialize)
}

type NullableAuthLdap struct {
	value *AuthLdap
	isSet bool
}

func (v NullableAuthLdap) Get() *AuthLdap {
	return v.value
}

func (v *NullableAuthLdap) Set(val *AuthLdap) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthLdap) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthLdap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthLdap(val *AuthLdap) *NullableAuthLdap {
	return &NullableAuthLdap{value: val, isSet: true}
}

func (v NullableAuthLdap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthLdap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


