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

// AccountCursorAllOf struct for AccountCursorAllOf
type AccountCursorAllOf struct {
	Data []Account `json:"data"`
}

// NewAccountCursorAllOf instantiates a new AccountCursorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCursorAllOf(data []Account) *AccountCursorAllOf {
	this := AccountCursorAllOf{}
	this.Data = data
	return &this
}

// NewAccountCursorAllOfWithDefaults instantiates a new AccountCursorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCursorAllOfWithDefaults() *AccountCursorAllOf {
	this := AccountCursorAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *AccountCursorAllOf) GetData() []Account {
	if o == nil {
		var ret []Account
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccountCursorAllOf) GetDataOk() (*[]Account, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AccountCursorAllOf) SetData(v []Account) {
	o.Data = v
}

func (o AccountCursorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCursorAllOf struct {
	value *AccountCursorAllOf
	isSet bool
}

func (v NullableAccountCursorAllOf) Get() *AccountCursorAllOf {
	return v.value
}

func (v *NullableAccountCursorAllOf) Set(val *AccountCursorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCursorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCursorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCursorAllOf(val *AccountCursorAllOf) *NullableAccountCursorAllOf {
	return &NullableAccountCursorAllOf{value: val, isSet: true}
}

func (v NullableAccountCursorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCursorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


