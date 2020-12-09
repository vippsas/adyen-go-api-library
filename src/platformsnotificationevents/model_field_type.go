/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
)

// FieldType struct for FieldType
type FieldType struct {
	// The full name of the property.
	Field *string `json:"field,omitempty"`
	// The type of the field.
	FieldName *string `json:"fieldName,omitempty"`
	// The code of the shareholder that the field belongs to. If empty, the field belongs to an account holder.
	ShareholderCode *string `json:"shareholderCode,omitempty"`
}

// NewFieldType instantiates a new FieldType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldType() *FieldType {
	this := FieldType{}
	return &this
}

// NewFieldTypeWithDefaults instantiates a new FieldType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldTypeWithDefaults() *FieldType {
	this := FieldType{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *FieldType) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldType) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *FieldType) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *FieldType) SetField(v string) {
	o.Field = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *FieldType) GetFieldName() string {
	if o == nil || o.FieldName == nil {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldType) GetFieldNameOk() (*string, bool) {
	if o == nil || o.FieldName == nil {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *FieldType) HasFieldName() bool {
	if o != nil && o.FieldName != nil {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *FieldType) SetFieldName(v string) {
	o.FieldName = &v
}

// GetShareholderCode returns the ShareholderCode field value if set, zero value otherwise.
func (o *FieldType) GetShareholderCode() string {
	if o == nil || o.ShareholderCode == nil {
		var ret string
		return ret
	}
	return *o.ShareholderCode
}

// GetShareholderCodeOk returns a tuple with the ShareholderCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldType) GetShareholderCodeOk() (*string, bool) {
	if o == nil || o.ShareholderCode == nil {
		return nil, false
	}
	return o.ShareholderCode, true
}

// HasShareholderCode returns a boolean if a field has been set.
func (o *FieldType) HasShareholderCode() bool {
	if o != nil && o.ShareholderCode != nil {
		return true
	}

	return false
}

// SetShareholderCode gets a reference to the given string and assigns it to the ShareholderCode field.
func (o *FieldType) SetShareholderCode(v string) {
	o.ShareholderCode = &v
}

func (o FieldType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.FieldName != nil {
		toSerialize["fieldName"] = o.FieldName
	}
	if o.ShareholderCode != nil {
		toSerialize["shareholderCode"] = o.ShareholderCode
	}
	return json.Marshal(toSerialize)
}

type NullableFieldType struct {
	value *FieldType
	isSet bool
}

func (v NullableFieldType) Get() *FieldType {
	return v.value
}

func (v *NullableFieldType) Set(val *FieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldType(val *FieldType) *NullableFieldType {
	return &NullableFieldType{value: val, isSet: true}
}

func (v NullableFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

