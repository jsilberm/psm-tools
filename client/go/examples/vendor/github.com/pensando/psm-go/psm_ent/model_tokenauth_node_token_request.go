/*
 * Tokenauth API reference
 *
 * APIs to generate node auth tokens  
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
	"time"
)

// TokenauthNodeTokenRequest NodeTokenRequest is a message that allows user to retrieve an auth token to connect to a cluster node and perform privileged operations.
type TokenauthNodeTokenRequest struct {
	// Audience represents a list of nodes the token is valid for. \"*\" indicates all nodes.
	Audience *[]string `json:"audience,omitempty"`
	// ValidityEnd indicates the time at which the token becomes invalid.
	ValidityEnd *time.Time `json:"validity-end,omitempty"`
	// ValidityStart indicates the time at which the token becomes valid.
	ValidityStart *time.Time `json:"validity-start,omitempty"`
}

// NewTokenauthNodeTokenRequest instantiates a new TokenauthNodeTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenauthNodeTokenRequest() *TokenauthNodeTokenRequest {
	this := TokenauthNodeTokenRequest{}
	return &this
}

// NewTokenauthNodeTokenRequestWithDefaults instantiates a new TokenauthNodeTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenauthNodeTokenRequestWithDefaults() *TokenauthNodeTokenRequest {
	this := TokenauthNodeTokenRequest{}
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *TokenauthNodeTokenRequest) GetAudience() []string {
	if o == nil || o.Audience == nil {
		var ret []string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenauthNodeTokenRequest) GetAudienceOk() (*[]string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *TokenauthNodeTokenRequest) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given []string and assigns it to the Audience field.
func (o *TokenauthNodeTokenRequest) SetAudience(v []string) {
	o.Audience = &v
}

// GetValidityEnd returns the ValidityEnd field value if set, zero value otherwise.
func (o *TokenauthNodeTokenRequest) GetValidityEnd() time.Time {
	if o == nil || o.ValidityEnd == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidityEnd
}

// GetValidityEndOk returns a tuple with the ValidityEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenauthNodeTokenRequest) GetValidityEndOk() (*time.Time, bool) {
	if o == nil || o.ValidityEnd == nil {
		return nil, false
	}
	return o.ValidityEnd, true
}

// HasValidityEnd returns a boolean if a field has been set.
func (o *TokenauthNodeTokenRequest) HasValidityEnd() bool {
	if o != nil && o.ValidityEnd != nil {
		return true
	}

	return false
}

// SetValidityEnd gets a reference to the given time.Time and assigns it to the ValidityEnd field.
func (o *TokenauthNodeTokenRequest) SetValidityEnd(v time.Time) {
	o.ValidityEnd = &v
}

// GetValidityStart returns the ValidityStart field value if set, zero value otherwise.
func (o *TokenauthNodeTokenRequest) GetValidityStart() time.Time {
	if o == nil || o.ValidityStart == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidityStart
}

// GetValidityStartOk returns a tuple with the ValidityStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenauthNodeTokenRequest) GetValidityStartOk() (*time.Time, bool) {
	if o == nil || o.ValidityStart == nil {
		return nil, false
	}
	return o.ValidityStart, true
}

// HasValidityStart returns a boolean if a field has been set.
func (o *TokenauthNodeTokenRequest) HasValidityStart() bool {
	if o != nil && o.ValidityStart != nil {
		return true
	}

	return false
}

// SetValidityStart gets a reference to the given time.Time and assigns it to the ValidityStart field.
func (o *TokenauthNodeTokenRequest) SetValidityStart(v time.Time) {
	o.ValidityStart = &v
}

func (o TokenauthNodeTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.ValidityEnd != nil {
		toSerialize["validity-end"] = o.ValidityEnd
	}
	if o.ValidityStart != nil {
		toSerialize["validity-start"] = o.ValidityStart
	}
	return json.Marshal(toSerialize)
}

type NullableTokenauthNodeTokenRequest struct {
	value *TokenauthNodeTokenRequest
	isSet bool
}

func (v NullableTokenauthNodeTokenRequest) Get() *TokenauthNodeTokenRequest {
	return v.value
}

func (v *NullableTokenauthNodeTokenRequest) Set(val *TokenauthNodeTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenauthNodeTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenauthNodeTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenauthNodeTokenRequest(val *TokenauthNodeTokenRequest) *NullableTokenauthNodeTokenRequest {
	return &NullableTokenauthNodeTokenRequest{value: val, isSet: true}
}

func (v NullableTokenauthNodeTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenauthNodeTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


