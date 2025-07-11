/*
N-central API-Service

<h3>API Access</h3> <p>In order to use the API-Service endpoints, ensure the following prerequisites are met:</p> <ol>     <li>User is created in N-central with appropriate permissions and configuration         (roles, access groups, MFA disabled, 2FA disabled).</li>     <li>API access is set up in N-central by having a JWT         (Json Web Token, referred to as \"N-central User-API Token\") generated,         acting as the permanent secret solely used for fetching access and refresh tokens.</li> </ol> <p>To access the API-Service endpoints, the JWT must first be exchanged with access and refresh tokens:</p> <ul>     <li>         <p><b>Using Swagger UI:</b></p>         <ol>             <li>Click on the <b>\"Authorize\"</b> button and enter the JWT into the value field                 under <b>\"N-central User-API Token (JWT)\"</b>, then navigate to                 the <code>/api/auth/authenticate</code> endpoint, click <b>\"Try it out\"</b>                 and then <b>\"Execute\"</b>.                 <br/>The new “Access Token” and “Refresh Token” fields will be available in the “Server Response”                 section below. Note that the expiry is 3600s ( 1h ). Copy the value of the access token.</li>             <li>Click on the <b>lock icon</b> to the right of any endpoint                 (or the <b>\"Authorize\"</b> button at the top), enter the access token in the                 <b>\"API-Access Token\"</b> field and click <b>\"Authorize\"</b>.</li>             <li>To call an API-Service endpoint, navigate to it, click <b>\"Try it out\"</b> button and             then <b>\"Execute\"</b>. If the steps above were successful, the access token is included in requests             automatically.</li>         </ol>     </li>     <li>         <p><b>Using a different HTTP client:</b></p>         <ol>             <li>Call the <code>/api/auth/authenticate</code> endpoint.                 The JWT token must be specified under the <code>Authorization</code> header,                 in the form <code>Bearer &lt;YOUR_JWT&gt;\"</code>.                 <br/>The access and refresh tokens are present in the response.</li>             <li>When calling API-Service endpoints, make sure to specify the <b>access token</b> in             the <code>Authorization</code> header as <code>Bearer &lt;ACCESS_TOKEN&gt;</code>.</li>         </ol>     </li> </ul> <h3>API Pagination & Sorting</h3> <p>Certain API-Service query endpoints support pagination and sorting through the use of query parameters.</p> <p><b>Pagination query parameters:</b></p> <ul>     <li>pageSize: number between 1 and 1000 specifying how many items to return for each page (if available).         If unspecified, the default is 50.</li>     <li>pageNumber: number specifying what page of data to retrieve, starting from 1 as the first page.         If unspecified, the default is 1.</li> </ul> <p>A paginated response contains several related fields, such as pageSize, pageNumber, itemCount, totalItems,     totalPages, _links (first, last, previous and next pages) and _warning (containing any warning messages,     if present).</p> <p><b>Sorting query parameters:</b></p> <ul>     <li>sortBy: the name of the field to sort the results by. If unspecified, no sorting is applied.</li>     <li>sortOrder: the case insensitive sorting direction supporting ASC/ASCENDING/NATURAL and         DESC/DESCENDING/REVERSE. If unspecified, the default is ASC.</li> </ul> <h3>API Rate Limiting</h3> <p>The API-Service endpoints are rate limited to ensure the stability, availability and performance of     the overall system.</p> <p>Upon reaching such a rate limit, the endpoints return HTTP Status 429 - Too Many Requests.</p> <p>The system will accept further requests once existing in-flight requests are completed.</p> 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ncentral

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DeviceAccessGroupCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAccessGroupCreateRequest{}

// DeviceAccessGroupCreateRequest Request for creating a device type access group.
type DeviceAccessGroupCreateRequest struct {
	// Name of the access group.
	GroupName string `json:"groupName"`
	// Description of the access group.
	GroupDescription string `json:"groupDescription"`
	// List of device IDs to attach to the access group.
	DeviceIds []string `json:"deviceIds,omitempty"`
	// List of user IDs to be associated with the access group.
	UserIds []string `json:"userIds,omitempty"`
}

type _DeviceAccessGroupCreateRequest DeviceAccessGroupCreateRequest

// NewDeviceAccessGroupCreateRequest instantiates a new DeviceAccessGroupCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAccessGroupCreateRequest(groupName string, groupDescription string) *DeviceAccessGroupCreateRequest {
	this := DeviceAccessGroupCreateRequest{}
	this.GroupName = groupName
	this.GroupDescription = groupDescription
	return &this
}

// NewDeviceAccessGroupCreateRequestWithDefaults instantiates a new DeviceAccessGroupCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAccessGroupCreateRequestWithDefaults() *DeviceAccessGroupCreateRequest {
	this := DeviceAccessGroupCreateRequest{}
	return &this
}

// GetGroupName returns the GroupName field value
func (o *DeviceAccessGroupCreateRequest) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *DeviceAccessGroupCreateRequest) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *DeviceAccessGroupCreateRequest) SetGroupName(v string) {
	o.GroupName = v
}

// GetGroupDescription returns the GroupDescription field value
func (o *DeviceAccessGroupCreateRequest) GetGroupDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupDescription
}

// GetGroupDescriptionOk returns a tuple with the GroupDescription field value
// and a boolean to check if the value has been set.
func (o *DeviceAccessGroupCreateRequest) GetGroupDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupDescription, true
}

// SetGroupDescription sets field value
func (o *DeviceAccessGroupCreateRequest) SetGroupDescription(v string) {
	o.GroupDescription = v
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *DeviceAccessGroupCreateRequest) GetDeviceIds() []string {
	if o == nil || IsNil(o.DeviceIds) {
		var ret []string
		return ret
	}
	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessGroupCreateRequest) GetDeviceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceIds) {
		return nil, false
	}
	return o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *DeviceAccessGroupCreateRequest) HasDeviceIds() bool {
	if o != nil && !IsNil(o.DeviceIds) {
		return true
	}

	return false
}

// SetDeviceIds gets a reference to the given []string and assigns it to the DeviceIds field.
func (o *DeviceAccessGroupCreateRequest) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *DeviceAccessGroupCreateRequest) GetUserIds() []string {
	if o == nil || IsNil(o.UserIds) {
		var ret []string
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessGroupCreateRequest) GetUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *DeviceAccessGroupCreateRequest) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []string and assigns it to the UserIds field.
func (o *DeviceAccessGroupCreateRequest) SetUserIds(v []string) {
	o.UserIds = v
}

func (o DeviceAccessGroupCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAccessGroupCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupName"] = o.GroupName
	toSerialize["groupDescription"] = o.GroupDescription
	if !IsNil(o.DeviceIds) {
		toSerialize["deviceIds"] = o.DeviceIds
	}
	if !IsNil(o.UserIds) {
		toSerialize["userIds"] = o.UserIds
	}
	return toSerialize, nil
}

func (o *DeviceAccessGroupCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupName",
		"groupDescription",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeviceAccessGroupCreateRequest := _DeviceAccessGroupCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceAccessGroupCreateRequest)

	if err != nil {
		return err
	}

	*o = DeviceAccessGroupCreateRequest(varDeviceAccessGroupCreateRequest)

	return err
}

type NullableDeviceAccessGroupCreateRequest struct {
	value *DeviceAccessGroupCreateRequest
	isSet bool
}

func (v NullableDeviceAccessGroupCreateRequest) Get() *DeviceAccessGroupCreateRequest {
	return v.value
}

func (v *NullableDeviceAccessGroupCreateRequest) Set(val *DeviceAccessGroupCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAccessGroupCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAccessGroupCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAccessGroupCreateRequest(val *DeviceAccessGroupCreateRequest) *NullableDeviceAccessGroupCreateRequest {
	return &NullableDeviceAccessGroupCreateRequest{value: val, isSet: true}
}

func (v NullableDeviceAccessGroupCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAccessGroupCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


