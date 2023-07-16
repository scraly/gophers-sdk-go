/*
gophers-api

HTTP server that handle cute Gophers.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gophers-api

import (
	"encoding/json"
)

// checks if the Gopher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gopher{}

// Gopher struct for Gopher
type Gopher struct {
	Name *string `json:"name,omitempty"`
	Displayname *string `json:"displayname,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewGopher instantiates a new Gopher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGopher() *Gopher {
	this := Gopher{}
	return &this
}

// NewGopherWithDefaults instantiates a new Gopher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGopherWithDefaults() *Gopher {
	this := Gopher{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Gopher) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gopher) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Gopher) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Gopher) SetName(v string) {
	o.Name = &v
}

// GetDisplayname returns the Displayname field value if set, zero value otherwise.
func (o *Gopher) GetDisplayname() string {
	if o == nil || IsNil(o.Displayname) {
		var ret string
		return ret
	}
	return *o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gopher) GetDisplaynameOk() (*string, bool) {
	if o == nil || IsNil(o.Displayname) {
		return nil, false
	}
	return o.Displayname, true
}

// HasDisplayname returns a boolean if a field has been set.
func (o *Gopher) HasDisplayname() bool {
	if o != nil && !IsNil(o.Displayname) {
		return true
	}

	return false
}

// SetDisplayname gets a reference to the given string and assigns it to the Displayname field.
func (o *Gopher) SetDisplayname(v string) {
	o.Displayname = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Gopher) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gopher) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Gopher) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Gopher) SetUrl(v string) {
	o.Url = &v
}

func (o Gopher) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gopher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Displayname) {
		toSerialize["displayname"] = o.Displayname
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableGopher struct {
	value *Gopher
	isSet bool
}

func (v NullableGopher) Get() *Gopher {
	return v.value
}

func (v *NullableGopher) Set(val *Gopher) {
	v.value = val
	v.isSet = true
}

func (v NullableGopher) IsSet() bool {
	return v.isSet
}

func (v *NullableGopher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGopher(val *Gopher) *NullableGopher {
	return &NullableGopher{value: val, isSet: true}
}

func (v NullableGopher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGopher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


