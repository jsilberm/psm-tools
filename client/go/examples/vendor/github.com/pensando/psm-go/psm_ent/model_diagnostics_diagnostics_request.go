/*
 * Diagnostics API reference
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

// DiagnosticsDiagnosticsRequest DiagnosticsRequest sends a diagnostics request to the service.
type DiagnosticsDiagnosticsRequest struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	// Parameters to be passed for a diagnostic query.
	Parameters *map[string]string `json:"parameters,omitempty"`
	// Query is type of diagnostic information requested like Profile, Log. This string is service specific and meaning is assigned by the service.
	Query *string `json:"query,omitempty"`
	ServicePort *DiagnosticsServicePort `json:"service-port,omitempty"`
}

// NewDiagnosticsDiagnosticsRequest instantiates a new DiagnosticsDiagnosticsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticsDiagnosticsRequest() *DiagnosticsDiagnosticsRequest {
	this := DiagnosticsDiagnosticsRequest{}
	return &this
}

// NewDiagnosticsDiagnosticsRequestWithDefaults instantiates a new DiagnosticsDiagnosticsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticsDiagnosticsRequestWithDefaults() *DiagnosticsDiagnosticsRequest {
	this := DiagnosticsDiagnosticsRequest{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *DiagnosticsDiagnosticsRequest) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsDiagnosticsRequest) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *DiagnosticsDiagnosticsRequest) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *DiagnosticsDiagnosticsRequest) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *DiagnosticsDiagnosticsRequest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsDiagnosticsRequest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *DiagnosticsDiagnosticsRequest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *DiagnosticsDiagnosticsRequest) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DiagnosticsDiagnosticsRequest) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsDiagnosticsRequest) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DiagnosticsDiagnosticsRequest) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *DiagnosticsDiagnosticsRequest) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *DiagnosticsDiagnosticsRequest) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsDiagnosticsRequest) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *DiagnosticsDiagnosticsRequest) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *DiagnosticsDiagnosticsRequest) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *DiagnosticsDiagnosticsRequest) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsDiagnosticsRequest) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *DiagnosticsDiagnosticsRequest) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *DiagnosticsDiagnosticsRequest) SetQuery(v string) {
	o.Query = &v
}

// GetServicePort returns the ServicePort field value if set, zero value otherwise.
func (o *DiagnosticsDiagnosticsRequest) GetServicePort() DiagnosticsServicePort {
	if o == nil || o.ServicePort == nil {
		var ret DiagnosticsServicePort
		return ret
	}
	return *o.ServicePort
}

// GetServicePortOk returns a tuple with the ServicePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsDiagnosticsRequest) GetServicePortOk() (*DiagnosticsServicePort, bool) {
	if o == nil || o.ServicePort == nil {
		return nil, false
	}
	return o.ServicePort, true
}

// HasServicePort returns a boolean if a field has been set.
func (o *DiagnosticsDiagnosticsRequest) HasServicePort() bool {
	if o != nil && o.ServicePort != nil {
		return true
	}

	return false
}

// SetServicePort gets a reference to the given DiagnosticsServicePort and assigns it to the ServicePort field.
func (o *DiagnosticsDiagnosticsRequest) SetServicePort(v DiagnosticsServicePort) {
	o.ServicePort = &v
}

func (o DiagnosticsDiagnosticsRequest) MarshalJSON() ([]byte, error) {
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
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.ServicePort != nil {
		toSerialize["service-port"] = o.ServicePort
	}
	return json.Marshal(toSerialize)
}

type NullableDiagnosticsDiagnosticsRequest struct {
	value *DiagnosticsDiagnosticsRequest
	isSet bool
}

func (v NullableDiagnosticsDiagnosticsRequest) Get() *DiagnosticsDiagnosticsRequest {
	return v.value
}

func (v *NullableDiagnosticsDiagnosticsRequest) Set(val *DiagnosticsDiagnosticsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticsDiagnosticsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticsDiagnosticsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticsDiagnosticsRequest(val *DiagnosticsDiagnosticsRequest) *NullableDiagnosticsDiagnosticsRequest {
	return &NullableDiagnosticsDiagnosticsRequest{value: val, isSet: true}
}

func (v NullableDiagnosticsDiagnosticsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticsDiagnosticsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


