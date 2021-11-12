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

// NetworkPolicerProfile PolicerProfile.
type NetworkPolicerProfile struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *NetworkPolicerProfileSpec `json:"spec,omitempty"`
	Status *NetworkPolicerProfileStatus `json:"status,omitempty"`
}

// NewNetworkPolicerProfile instantiates a new NetworkPolicerProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPolicerProfile() *NetworkPolicerProfile {
	this := NetworkPolicerProfile{}
	return &this
}

// NewNetworkPolicerProfileWithDefaults instantiates a new NetworkPolicerProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPolicerProfileWithDefaults() *NetworkPolicerProfile {
	this := NetworkPolicerProfile{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkPolicerProfile) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPolicerProfile) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkPolicerProfile) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkPolicerProfile) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkPolicerProfile) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPolicerProfile) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkPolicerProfile) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkPolicerProfile) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NetworkPolicerProfile) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPolicerProfile) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NetworkPolicerProfile) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *NetworkPolicerProfile) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *NetworkPolicerProfile) GetSpec() NetworkPolicerProfileSpec {
	if o == nil || o.Spec == nil {
		var ret NetworkPolicerProfileSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPolicerProfile) GetSpecOk() (*NetworkPolicerProfileSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *NetworkPolicerProfile) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given NetworkPolicerProfileSpec and assigns it to the Spec field.
func (o *NetworkPolicerProfile) SetSpec(v NetworkPolicerProfileSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkPolicerProfile) GetStatus() NetworkPolicerProfileStatus {
	if o == nil || o.Status == nil {
		var ret NetworkPolicerProfileStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPolicerProfile) GetStatusOk() (*NetworkPolicerProfileStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkPolicerProfile) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NetworkPolicerProfileStatus and assigns it to the Status field.
func (o *NetworkPolicerProfile) SetStatus(v NetworkPolicerProfileStatus) {
	o.Status = &v
}

func (o NetworkPolicerProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkPolicerProfile struct {
	value *NetworkPolicerProfile
	isSet bool
}

func (v NullableNetworkPolicerProfile) Get() *NetworkPolicerProfile {
	return v.value
}

func (v *NullableNetworkPolicerProfile) Set(val *NetworkPolicerProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPolicerProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPolicerProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPolicerProfile(val *NetworkPolicerProfile) *NullableNetworkPolicerProfile {
	return &NullableNetworkPolicerProfile{value: val, isSet: true}
}

func (v NullableNetworkPolicerProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPolicerProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


