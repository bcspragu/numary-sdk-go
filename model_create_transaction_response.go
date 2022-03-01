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

// CreateTransactionResponse struct for CreateTransactionResponse
type CreateTransactionResponse struct {
	Data []Transaction `json:"data"`
}

// NewCreateTransactionResponse instantiates a new CreateTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransactionResponse(data []Transaction) *CreateTransactionResponse {
	this := CreateTransactionResponse{}
	this.Data = data
	return &this
}

// NewCreateTransactionResponseWithDefaults instantiates a new CreateTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransactionResponseWithDefaults() *CreateTransactionResponse {
	this := CreateTransactionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateTransactionResponse) GetData() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateTransactionResponse) GetDataOk() (*[]Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateTransactionResponse) SetData(v []Transaction) {
	o.Data = v
}

func (o CreateTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTransactionResponse struct {
	value *CreateTransactionResponse
	isSet bool
}

func (v NullableCreateTransactionResponse) Get() *CreateTransactionResponse {
	return v.value
}

func (v *NullableCreateTransactionResponse) Set(val *CreateTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransactionResponse(val *CreateTransactionResponse) *NullableCreateTransactionResponse {
	return &NullableCreateTransactionResponse{value: val, isSet: true}
}

func (v NullableCreateTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


