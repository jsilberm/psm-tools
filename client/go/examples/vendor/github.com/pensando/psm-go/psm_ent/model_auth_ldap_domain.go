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

// AuthLdapDomain struct for AuthLdapDomain
type AuthLdapDomain struct {
	AttributeMapping *AuthLdapAttributeMapping `json:"attribute-mapping,omitempty"`
	// The LDAP base DN to be used in a user search.
	BaseDn *string `json:"base-dn,omitempty"`
	// The bind DN is the string that Venice uses to log in to the LDAP server. Venice uses this account to validate the remote user attempting to log in. The base DN is the container name and path in the LDAPserver where Venice searches for the remote user account. This is where the password is validated. This contains the user authorization and assigned RBAC roles for use on Venice. Venice requests the attribute from theLDAP server.
	BindDn *string `json:"bind-dn,omitempty"`
	// The password for the LDAP database account specified in the Root DN field.
	BindPassword *string `json:"bind-password,omitempty"`
	// Servers is a list that lets you configure multiple LDAP servers for high availability.
	Servers *[]AuthLdapServer `json:"servers,omitempty"`
	// Tag to group domains for authentication.
	Tag *string `json:"tag,omitempty"`
}

// NewAuthLdapDomain instantiates a new AuthLdapDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthLdapDomain() *AuthLdapDomain {
	this := AuthLdapDomain{}
	return &this
}

// NewAuthLdapDomainWithDefaults instantiates a new AuthLdapDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthLdapDomainWithDefaults() *AuthLdapDomain {
	this := AuthLdapDomain{}
	return &this
}

// GetAttributeMapping returns the AttributeMapping field value if set, zero value otherwise.
func (o *AuthLdapDomain) GetAttributeMapping() AuthLdapAttributeMapping {
	if o == nil || o.AttributeMapping == nil {
		var ret AuthLdapAttributeMapping
		return ret
	}
	return *o.AttributeMapping
}

// GetAttributeMappingOk returns a tuple with the AttributeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapDomain) GetAttributeMappingOk() (*AuthLdapAttributeMapping, bool) {
	if o == nil || o.AttributeMapping == nil {
		return nil, false
	}
	return o.AttributeMapping, true
}

// HasAttributeMapping returns a boolean if a field has been set.
func (o *AuthLdapDomain) HasAttributeMapping() bool {
	if o != nil && o.AttributeMapping != nil {
		return true
	}

	return false
}

// SetAttributeMapping gets a reference to the given AuthLdapAttributeMapping and assigns it to the AttributeMapping field.
func (o *AuthLdapDomain) SetAttributeMapping(v AuthLdapAttributeMapping) {
	o.AttributeMapping = &v
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *AuthLdapDomain) GetBaseDn() string {
	if o == nil || o.BaseDn == nil {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapDomain) GetBaseDnOk() (*string, bool) {
	if o == nil || o.BaseDn == nil {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *AuthLdapDomain) HasBaseDn() bool {
	if o != nil && o.BaseDn != nil {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *AuthLdapDomain) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetBindDn returns the BindDn field value if set, zero value otherwise.
func (o *AuthLdapDomain) GetBindDn() string {
	if o == nil || o.BindDn == nil {
		var ret string
		return ret
	}
	return *o.BindDn
}

// GetBindDnOk returns a tuple with the BindDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapDomain) GetBindDnOk() (*string, bool) {
	if o == nil || o.BindDn == nil {
		return nil, false
	}
	return o.BindDn, true
}

// HasBindDn returns a boolean if a field has been set.
func (o *AuthLdapDomain) HasBindDn() bool {
	if o != nil && o.BindDn != nil {
		return true
	}

	return false
}

// SetBindDn gets a reference to the given string and assigns it to the BindDn field.
func (o *AuthLdapDomain) SetBindDn(v string) {
	o.BindDn = &v
}

// GetBindPassword returns the BindPassword field value if set, zero value otherwise.
func (o *AuthLdapDomain) GetBindPassword() string {
	if o == nil || o.BindPassword == nil {
		var ret string
		return ret
	}
	return *o.BindPassword
}

// GetBindPasswordOk returns a tuple with the BindPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapDomain) GetBindPasswordOk() (*string, bool) {
	if o == nil || o.BindPassword == nil {
		return nil, false
	}
	return o.BindPassword, true
}

// HasBindPassword returns a boolean if a field has been set.
func (o *AuthLdapDomain) HasBindPassword() bool {
	if o != nil && o.BindPassword != nil {
		return true
	}

	return false
}

// SetBindPassword gets a reference to the given string and assigns it to the BindPassword field.
func (o *AuthLdapDomain) SetBindPassword(v string) {
	o.BindPassword = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *AuthLdapDomain) GetServers() []AuthLdapServer {
	if o == nil || o.Servers == nil {
		var ret []AuthLdapServer
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapDomain) GetServersOk() (*[]AuthLdapServer, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *AuthLdapDomain) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []AuthLdapServer and assigns it to the Servers field.
func (o *AuthLdapDomain) SetServers(v []AuthLdapServer) {
	o.Servers = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AuthLdapDomain) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapDomain) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AuthLdapDomain) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AuthLdapDomain) SetTag(v string) {
	o.Tag = &v
}

func (o AuthLdapDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttributeMapping != nil {
		toSerialize["attribute-mapping"] = o.AttributeMapping
	}
	if o.BaseDn != nil {
		toSerialize["base-dn"] = o.BaseDn
	}
	if o.BindDn != nil {
		toSerialize["bind-dn"] = o.BindDn
	}
	if o.BindPassword != nil {
		toSerialize["bind-password"] = o.BindPassword
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableAuthLdapDomain struct {
	value *AuthLdapDomain
	isSet bool
}

func (v NullableAuthLdapDomain) Get() *AuthLdapDomain {
	return v.value
}

func (v *NullableAuthLdapDomain) Set(val *AuthLdapDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthLdapDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthLdapDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthLdapDomain(val *AuthLdapDomain) *NullableAuthLdapDomain {
	return &NullableAuthLdapDomain{value: val, isSet: true}
}

func (v NullableAuthLdapDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthLdapDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


