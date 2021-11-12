/*
 * Browser API reference
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

// BrowserObject Object is a node in the dependency tree representing a config object with links to related objects.
type BrowserObject struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	// Links points to the relations of the object. The key for the map is the path to the filed which is causing the relation.
	Links *map[string]ObjectURIs `json:"links,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	// QueryType specifies the direction of the relations in Links.
	QueryType *string `json:"query-type,omitempty"`
	// Reverse is the view from the object looking back in the reverse direction of the dependency tree.
	Reverse *string `json:"reverse,omitempty"`
	// URI is the Browser URI for this object.
	Uri *string `json:"uri,omitempty"`
}

// NewBrowserObject instantiates a new BrowserObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserObject() *BrowserObject {
	this := BrowserObject{}
	var queryType string = "dependencies"
	this.QueryType = &queryType
	return &this
}

// NewBrowserObjectWithDefaults instantiates a new BrowserObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserObjectWithDefaults() *BrowserObject {
	this := BrowserObject{}
	var queryType string = "dependencies"
	this.QueryType = &queryType
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *BrowserObject) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserObject) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *BrowserObject) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *BrowserObject) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *BrowserObject) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserObject) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *BrowserObject) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *BrowserObject) SetKind(v string) {
	o.Kind = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BrowserObject) GetLinks() map[string]ObjectURIs {
	if o == nil || o.Links == nil {
		var ret map[string]ObjectURIs
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserObject) GetLinksOk() (*map[string]ObjectURIs, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BrowserObject) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]ObjectURIs and assigns it to the Links field.
func (o *BrowserObject) SetLinks(v map[string]ObjectURIs) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BrowserObject) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserObject) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BrowserObject) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *BrowserObject) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *BrowserObject) GetQueryType() string {
	if o == nil || o.QueryType == nil {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserObject) GetQueryTypeOk() (*string, bool) {
	if o == nil || o.QueryType == nil {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *BrowserObject) HasQueryType() bool {
	if o != nil && o.QueryType != nil {
		return true
	}

	return false
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *BrowserObject) SetQueryType(v string) {
	o.QueryType = &v
}

// GetReverse returns the Reverse field value if set, zero value otherwise.
func (o *BrowserObject) GetReverse() string {
	if o == nil || o.Reverse == nil {
		var ret string
		return ret
	}
	return *o.Reverse
}

// GetReverseOk returns a tuple with the Reverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserObject) GetReverseOk() (*string, bool) {
	if o == nil || o.Reverse == nil {
		return nil, false
	}
	return o.Reverse, true
}

// HasReverse returns a boolean if a field has been set.
func (o *BrowserObject) HasReverse() bool {
	if o != nil && o.Reverse != nil {
		return true
	}

	return false
}

// SetReverse gets a reference to the given string and assigns it to the Reverse field.
func (o *BrowserObject) SetReverse(v string) {
	o.Reverse = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BrowserObject) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserObject) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BrowserObject) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BrowserObject) SetUri(v string) {
	o.Uri = &v
}

func (o BrowserObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.QueryType != nil {
		toSerialize["query-type"] = o.QueryType
	}
	if o.Reverse != nil {
		toSerialize["reverse"] = o.Reverse
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableBrowserObject struct {
	value *BrowserObject
	isSet bool
}

func (v NullableBrowserObject) Get() *BrowserObject {
	return v.value
}

func (v *NullableBrowserObject) Set(val *BrowserObject) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserObject) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserObject(val *BrowserObject) *NullableBrowserObject {
	return &NullableBrowserObject{value: val, isSet: true}
}

func (v NullableBrowserObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


