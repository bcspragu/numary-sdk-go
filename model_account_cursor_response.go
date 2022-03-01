/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// AccountCursorResponse struct for AccountCursorResponse
type AccountCursorResponse struct {
	Cursor AccountCursor `json:"cursor"`
}

// NewAccountCursorResponse instantiates a new AccountCursorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCursorResponse(cursor AccountCursor) *AccountCursorResponse {
	this := AccountCursorResponse{}
	this.Cursor = cursor
	return &this
}

// NewAccountCursorResponseWithDefaults instantiates a new AccountCursorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCursorResponseWithDefaults() *AccountCursorResponse {
	this := AccountCursorResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *AccountCursorResponse) GetCursor() AccountCursor {
	if o == nil {
		var ret AccountCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *AccountCursorResponse) GetCursorOk() (*AccountCursor, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *AccountCursorResponse) SetCursor(v AccountCursor) {
	o.Cursor = v
}

func (o AccountCursorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCursorResponse struct {
	value *AccountCursorResponse
	isSet bool
}

func (v NullableAccountCursorResponse) Get() *AccountCursorResponse {
	return v.value
}

func (v *NullableAccountCursorResponse) Set(val *AccountCursorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCursorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCursorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCursorResponse(val *AccountCursorResponse) *NullableAccountCursorResponse {
	return &NullableAccountCursorResponse{value: val, isSet: true}
}

func (v NullableAccountCursorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCursorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


