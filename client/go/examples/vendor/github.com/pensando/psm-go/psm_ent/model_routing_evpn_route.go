/*
 * Routing API reference
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

// RoutingEVPNRoute struct for RoutingEVPNRoute
type RoutingEVPNRoute struct {
	Type *int64 `json:"type,omitempty"`
	Type2Prefix *EVPNRouteEVPNType2Route `json:"type2-prefix,omitempty"`
	Type5Prefix *EVPNRouteEVPNType5Route `json:"type5-prefix,omitempty"`
}

// NewRoutingEVPNRoute instantiates a new RoutingEVPNRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingEVPNRoute() *RoutingEVPNRoute {
	this := RoutingEVPNRoute{}
	return &this
}

// NewRoutingEVPNRouteWithDefaults instantiates a new RoutingEVPNRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingEVPNRouteWithDefaults() *RoutingEVPNRoute {
	this := RoutingEVPNRoute{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoutingEVPNRoute) GetType() int64 {
	if o == nil || o.Type == nil {
		var ret int64
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingEVPNRoute) GetTypeOk() (*int64, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoutingEVPNRoute) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int64 and assigns it to the Type field.
func (o *RoutingEVPNRoute) SetType(v int64) {
	o.Type = &v
}

// GetType2Prefix returns the Type2Prefix field value if set, zero value otherwise.
func (o *RoutingEVPNRoute) GetType2Prefix() EVPNRouteEVPNType2Route {
	if o == nil || o.Type2Prefix == nil {
		var ret EVPNRouteEVPNType2Route
		return ret
	}
	return *o.Type2Prefix
}

// GetType2PrefixOk returns a tuple with the Type2Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingEVPNRoute) GetType2PrefixOk() (*EVPNRouteEVPNType2Route, bool) {
	if o == nil || o.Type2Prefix == nil {
		return nil, false
	}
	return o.Type2Prefix, true
}

// HasType2Prefix returns a boolean if a field has been set.
func (o *RoutingEVPNRoute) HasType2Prefix() bool {
	if o != nil && o.Type2Prefix != nil {
		return true
	}

	return false
}

// SetType2Prefix gets a reference to the given EVPNRouteEVPNType2Route and assigns it to the Type2Prefix field.
func (o *RoutingEVPNRoute) SetType2Prefix(v EVPNRouteEVPNType2Route) {
	o.Type2Prefix = &v
}

// GetType5Prefix returns the Type5Prefix field value if set, zero value otherwise.
func (o *RoutingEVPNRoute) GetType5Prefix() EVPNRouteEVPNType5Route {
	if o == nil || o.Type5Prefix == nil {
		var ret EVPNRouteEVPNType5Route
		return ret
	}
	return *o.Type5Prefix
}

// GetType5PrefixOk returns a tuple with the Type5Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingEVPNRoute) GetType5PrefixOk() (*EVPNRouteEVPNType5Route, bool) {
	if o == nil || o.Type5Prefix == nil {
		return nil, false
	}
	return o.Type5Prefix, true
}

// HasType5Prefix returns a boolean if a field has been set.
func (o *RoutingEVPNRoute) HasType5Prefix() bool {
	if o != nil && o.Type5Prefix != nil {
		return true
	}

	return false
}

// SetType5Prefix gets a reference to the given EVPNRouteEVPNType5Route and assigns it to the Type5Prefix field.
func (o *RoutingEVPNRoute) SetType5Prefix(v EVPNRouteEVPNType5Route) {
	o.Type5Prefix = &v
}

func (o RoutingEVPNRoute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Type2Prefix != nil {
		toSerialize["type2-prefix"] = o.Type2Prefix
	}
	if o.Type5Prefix != nil {
		toSerialize["type5-prefix"] = o.Type5Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableRoutingEVPNRoute struct {
	value *RoutingEVPNRoute
	isSet bool
}

func (v NullableRoutingEVPNRoute) Get() *RoutingEVPNRoute {
	return v.value
}

func (v *NullableRoutingEVPNRoute) Set(val *RoutingEVPNRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingEVPNRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingEVPNRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingEVPNRoute(val *RoutingEVPNRoute) *NullableRoutingEVPNRoute {
	return &NullableRoutingEVPNRoute{value: val, isSet: true}
}

func (v NullableRoutingEVPNRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingEVPNRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


