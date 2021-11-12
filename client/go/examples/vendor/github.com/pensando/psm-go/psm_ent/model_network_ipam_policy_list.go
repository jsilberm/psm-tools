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

// NetworkIPAMPolicyList IPAMPolicyList is a container object for list of IPAMPolicy objects.
type NetworkIPAMPolicyList struct {
	ApiVersion *string `json:"api-version,omitempty"`
	// List of IPAMPolicy objects.
	Items *[]NetworkIPAMPolicy `json:"items,omitempty"`
	Kind *string `json:"kind,omitempty"`
	ListMeta *ApiListMeta `json:"list-meta,omitempty"`
}

// NewNetworkIPAMPolicyList instantiates a new NetworkIPAMPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkIPAMPolicyList() *NetworkIPAMPolicyList {
	this := NetworkIPAMPolicyList{}
	return &this
}

// NewNetworkIPAMPolicyListWithDefaults instantiates a new NetworkIPAMPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkIPAMPolicyListWithDefaults() *NetworkIPAMPolicyList {
	this := NetworkIPAMPolicyList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkIPAMPolicyList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMPolicyList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkIPAMPolicyList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkIPAMPolicyList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *NetworkIPAMPolicyList) GetItems() []NetworkIPAMPolicy {
	if o == nil || o.Items == nil {
		var ret []NetworkIPAMPolicy
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMPolicyList) GetItemsOk() (*[]NetworkIPAMPolicy, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *NetworkIPAMPolicyList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []NetworkIPAMPolicy and assigns it to the Items field.
func (o *NetworkIPAMPolicyList) SetItems(v []NetworkIPAMPolicy) {
	o.Items = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkIPAMPolicyList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMPolicyList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkIPAMPolicyList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkIPAMPolicyList) SetKind(v string) {
	o.Kind = &v
}

// GetListMeta returns the ListMeta field value if set, zero value otherwise.
func (o *NetworkIPAMPolicyList) GetListMeta() ApiListMeta {
	if o == nil || o.ListMeta == nil {
		var ret ApiListMeta
		return ret
	}
	return *o.ListMeta
}

// GetListMetaOk returns a tuple with the ListMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMPolicyList) GetListMetaOk() (*ApiListMeta, bool) {
	if o == nil || o.ListMeta == nil {
		return nil, false
	}
	return o.ListMeta, true
}

// HasListMeta returns a boolean if a field has been set.
func (o *NetworkIPAMPolicyList) HasListMeta() bool {
	if o != nil && o.ListMeta != nil {
		return true
	}

	return false
}

// SetListMeta gets a reference to the given ApiListMeta and assigns it to the ListMeta field.
func (o *NetworkIPAMPolicyList) SetListMeta(v ApiListMeta) {
	o.ListMeta = &v
}

func (o NetworkIPAMPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.ListMeta != nil {
		toSerialize["list-meta"] = o.ListMeta
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkIPAMPolicyList struct {
	value *NetworkIPAMPolicyList
	isSet bool
}

func (v NullableNetworkIPAMPolicyList) Get() *NetworkIPAMPolicyList {
	return v.value
}

func (v *NullableNetworkIPAMPolicyList) Set(val *NetworkIPAMPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkIPAMPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkIPAMPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkIPAMPolicyList(val *NetworkIPAMPolicyList) *NullableNetworkIPAMPolicyList {
	return &NullableNetworkIPAMPolicyList{value: val, isSet: true}
}

func (v NullableNetworkIPAMPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkIPAMPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


