/*
OpenAPI definition

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProjectName.Client

import (
	"encoding/json"
)

// VersionResponse struct for VersionResponse
type VersionResponse struct {
	Version *string `json:"version,omitempty"`
	Stable *bool `json:"stable,omitempty"`
}

// NewVersionResponse instantiates a new VersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionResponse() *VersionResponse {
	this := VersionResponse{}
	return &this
}

// NewVersionResponseWithDefaults instantiates a new VersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionResponseWithDefaults() *VersionResponse {
	this := VersionResponse{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VersionResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VersionResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VersionResponse) SetVersion(v string) {
	o.Version = &v
}

// GetStable returns the Stable field value if set, zero value otherwise.
func (o *VersionResponse) GetStable() bool {
	if o == nil || o.Stable == nil {
		var ret bool
		return ret
	}
	return *o.Stable
}

// GetStableOk returns a tuple with the Stable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionResponse) GetStableOk() (*bool, bool) {
	if o == nil || o.Stable == nil {
		return nil, false
	}
	return o.Stable, true
}

// HasStable returns a boolean if a field has been set.
func (o *VersionResponse) HasStable() bool {
	if o != nil && o.Stable != nil {
		return true
	}

	return false
}

// SetStable gets a reference to the given bool and assigns it to the Stable field.
func (o *VersionResponse) SetStable(v bool) {
	o.Stable = &v
}

func (o VersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Stable != nil {
		toSerialize["stable"] = o.Stable
	}
	return json.Marshal(toSerialize)
}

type NullableVersionResponse struct {
	value *VersionResponse
	isSet bool
}

func (v NullableVersionResponse) Get() *VersionResponse {
	return v.value
}

func (v *NullableVersionResponse) Set(val *VersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionResponse(val *VersionResponse) *NullableVersionResponse {
	return &NullableVersionResponse{value: val, isSet: true}
}

func (v NullableVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


