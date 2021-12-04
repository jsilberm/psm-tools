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

// AuthRadiusDomain struct for AuthRadiusDomain
type AuthRadiusDomain struct {
	// NasID is a string identifying the NAS(API Gw) originating the Access-Request.
	NasId *string `json:"nas-id,omitempty"`
	Servers *[]AuthRadiusServer `json:"servers,omitempty"`
	// Tag to group domains for authentication.
	Tag *string `json:"tag,omitempty"`
}

// NewAuthRadiusDomain instantiates a new AuthRadiusDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRadiusDomain() *AuthRadiusDomain {
	this := AuthRadiusDomain{}
	return &this
}

// NewAuthRadiusDomainWithDefaults instantiates a new AuthRadiusDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRadiusDomainWithDefaults() *AuthRadiusDomain {
	this := AuthRadiusDomain{}
	return &this
}

// GetNasId returns the NasId field value if set, zero value otherwise.
func (o *AuthRadiusDomain) GetNasId() string {
	if o == nil || o.NasId == nil {
		var ret string
		return ret
	}
	return *o.NasId
}

// GetNasIdOk returns a tuple with the NasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRadiusDomain) GetNasIdOk() (*string, bool) {
	if o == nil || o.NasId == nil {
		return nil, false
	}
	return o.NasId, true
}

// HasNasId returns a boolean if a field has been set.
func (o *AuthRadiusDomain) HasNasId() bool {
	if o != nil && o.NasId != nil {
		return true
	}

	return false
}

// SetNasId gets a reference to the given string and assigns it to the NasId field.
func (o *AuthRadiusDomain) SetNasId(v string) {
	o.NasId = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *AuthRadiusDomain) GetServers() []AuthRadiusServer {
	if o == nil || o.Servers == nil {
		var ret []AuthRadiusServer
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRadiusDomain) GetServersOk() (*[]AuthRadiusServer, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *AuthRadiusDomain) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []AuthRadiusServer and assigns it to the Servers field.
func (o *AuthRadiusDomain) SetServers(v []AuthRadiusServer) {
	o.Servers = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AuthRadiusDomain) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRadiusDomain) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AuthRadiusDomain) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AuthRadiusDomain) SetTag(v string) {
	o.Tag = &v
}

func (o AuthRadiusDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NasId != nil {
		toSerialize["nas-id"] = o.NasId
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableAuthRadiusDomain struct {
	value *AuthRadiusDomain
	isSet bool
}

func (v NullableAuthRadiusDomain) Get() *AuthRadiusDomain {
	return v.value
}

func (v *NullableAuthRadiusDomain) Set(val *AuthRadiusDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRadiusDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRadiusDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRadiusDomain(val *AuthRadiusDomain) *NullableAuthRadiusDomain {
	return &NullableAuthRadiusDomain{value: val, isSet: true}
}

func (v NullableAuthRadiusDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRadiusDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

