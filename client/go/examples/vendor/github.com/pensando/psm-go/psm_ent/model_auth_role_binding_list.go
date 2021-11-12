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

// AuthRoleBindingList RoleBindingList is a container object for list of RoleBinding objects.
type AuthRoleBindingList struct {
	ApiVersion *string `json:"api-version,omitempty"`
	// List of RoleBinding objects.
	Items *[]AuthRoleBinding `json:"items,omitempty"`
	Kind *string `json:"kind,omitempty"`
	ListMeta *ApiListMeta `json:"list-meta,omitempty"`
}

// NewAuthRoleBindingList instantiates a new AuthRoleBindingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRoleBindingList() *AuthRoleBindingList {
	this := AuthRoleBindingList{}
	return &this
}

// NewAuthRoleBindingListWithDefaults instantiates a new AuthRoleBindingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRoleBindingListWithDefaults() *AuthRoleBindingList {
	this := AuthRoleBindingList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *AuthRoleBindingList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRoleBindingList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *AuthRoleBindingList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *AuthRoleBindingList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AuthRoleBindingList) GetItems() []AuthRoleBinding {
	if o == nil || o.Items == nil {
		var ret []AuthRoleBinding
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRoleBindingList) GetItemsOk() (*[]AuthRoleBinding, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AuthRoleBindingList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AuthRoleBinding and assigns it to the Items field.
func (o *AuthRoleBindingList) SetItems(v []AuthRoleBinding) {
	o.Items = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AuthRoleBindingList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRoleBindingList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AuthRoleBindingList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *AuthRoleBindingList) SetKind(v string) {
	o.Kind = &v
}

// GetListMeta returns the ListMeta field value if set, zero value otherwise.
func (o *AuthRoleBindingList) GetListMeta() ApiListMeta {
	if o == nil || o.ListMeta == nil {
		var ret ApiListMeta
		return ret
	}
	return *o.ListMeta
}

// GetListMetaOk returns a tuple with the ListMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRoleBindingList) GetListMetaOk() (*ApiListMeta, bool) {
	if o == nil || o.ListMeta == nil {
		return nil, false
	}
	return o.ListMeta, true
}

// HasListMeta returns a boolean if a field has been set.
func (o *AuthRoleBindingList) HasListMeta() bool {
	if o != nil && o.ListMeta != nil {
		return true
	}

	return false
}

// SetListMeta gets a reference to the given ApiListMeta and assigns it to the ListMeta field.
func (o *AuthRoleBindingList) SetListMeta(v ApiListMeta) {
	o.ListMeta = &v
}

func (o AuthRoleBindingList) MarshalJSON() ([]byte, error) {
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

type NullableAuthRoleBindingList struct {
	value *AuthRoleBindingList
	isSet bool
}

func (v NullableAuthRoleBindingList) Get() *AuthRoleBindingList {
	return v.value
}

func (v *NullableAuthRoleBindingList) Set(val *AuthRoleBindingList) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRoleBindingList) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRoleBindingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRoleBindingList(val *AuthRoleBindingList) *NullableAuthRoleBindingList {
	return &NullableAuthRoleBindingList{value: val, isSet: true}
}

func (v NullableAuthRoleBindingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRoleBindingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


