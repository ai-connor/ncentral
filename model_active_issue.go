/*
N-central API-Service

<h3>API Access</h3> <p>In order to use the API-Service endpoints, ensure the following prerequisites are met:</p> <ol>     <li>User is created in N-central with appropriate permissions and configuration         (roles, access groups, MFA disabled, 2FA disabled).</li>     <li>API access is set up in N-central by having a JWT         (Json Web Token, referred to as \"N-central User-API Token\") generated,         acting as the permanent secret solely used for fetching access and refresh tokens.</li> </ol> <p>To access the API-Service endpoints, the JWT must first be exchanged with access and refresh tokens:</p> <ul>     <li>         <p><b>Using Swagger UI:</b></p>         <ol>             <li>Click on the <b>\"Authorize\"</b> button and enter the JWT into the value field                 under <b>\"N-central User-API Token (JWT)\"</b>, then navigate to                 the <code>/api/auth/authenticate</code> endpoint, click <b>\"Try it out\"</b>                 and then <b>\"Execute\"</b>.                 <br/>The new “Access Token” and “Refresh Token” fields will be available in the “Server Response”                 section below. Note that the expiry is 3600s ( 1h ). Copy the value of the access token.</li>             <li>Click on the <b>lock icon</b> to the right of any endpoint                 (or the <b>\"Authorize\"</b> button at the top), enter the access token in the                 <b>\"API-Access Token\"</b> field and click <b>\"Authorize\"</b>.</li>             <li>To call an API-Service endpoint, navigate to it, click <b>\"Try it out\"</b> button and             then <b>\"Execute\"</b>. If the steps above were successful, the access token is included in requests             automatically.</li>         </ol>     </li>     <li>         <p><b>Using a different HTTP client:</b></p>         <ol>             <li>Call the <code>/api/auth/authenticate</code> endpoint.                 The JWT token must be specified under the <code>Authorization</code> header,                 in the form <code>Bearer &lt;YOUR_JWT&gt;\"</code>.                 <br/>The access and refresh tokens are present in the response.</li>             <li>When calling API-Service endpoints, make sure to specify the <b>access token</b> in             the <code>Authorization</code> header as <code>Bearer &lt;ACCESS_TOKEN&gt;</code>.</li>         </ol>     </li> </ul> <h3>API Pagination & Sorting</h3> <p>Certain API-Service query endpoints support pagination and sorting through the use of query parameters.</p> <p><b>Pagination query parameters:</b></p> <ul>     <li>pageSize: number between 1 and 1000 specifying how many items to return for each page (if available).         If unspecified, the default is 50.</li>     <li>pageNumber: number specifying what page of data to retrieve, starting from 1 as the first page.         If unspecified, the default is 1.</li> </ul> <p>A paginated response contains several related fields, such as pageSize, pageNumber, itemCount, totalItems,     totalPages, _links (first, last, previous and next pages) and _warning (containing any warning messages,     if present).</p> <p><b>Sorting query parameters:</b></p> <ul>     <li>sortBy: the name of the field to sort the results by. If unspecified, no sorting is applied.</li>     <li>sortOrder: the case insensitive sorting direction supporting ASC/ASCENDING/NATURAL and         DESC/DESCENDING/REVERSE. If unspecified, the default is ASC.</li> </ul> <h3>API Rate Limiting</h3> <p>The API-Service endpoints are rate limited to ensure the stability, availability and performance of     the overall system.</p> <p>Upon reaching such a rate limit, the endpoints return HTTP Status 429 - Too Many Requests.</p> <p>The system will accept further requests once existing in-flight requests are completed.</p> 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ncentral

import (
	"encoding/json"
)

// checks if the ActiveIssue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveIssue{}

// ActiveIssue Represents a response containing an active issue.
type ActiveIssue struct {
	// The organization unit id.
	OrgUnitId *int32 `json:"orgUnitId,omitempty"`
	// The device id.
	DeviceId *int32 `json:"deviceId,omitempty"`
	// The notification state.
	NotificationState *int32 `json:"notificationState,omitempty"`
	// The service id.
	ServiceId *int32 `json:"serviceId,omitempty"`
	// The service name.
	ServiceName *string `json:"serviceName,omitempty"`
	// The service type.
	ServiceType *string `json:"serviceType,omitempty"`
	// The task id.
	TaskId *int32 `json:"taskId,omitempty"`
	// The service item id.
	ServiceItemId *int32 `json:"serviceItemId,omitempty"`
	// The _extra information.
	Extra map[string]map[string]interface{} `json:"_extra,omitempty"`
}

// NewActiveIssue instantiates a new ActiveIssue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveIssue() *ActiveIssue {
	this := ActiveIssue{}
	return &this
}

// NewActiveIssueWithDefaults instantiates a new ActiveIssue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveIssueWithDefaults() *ActiveIssue {
	this := ActiveIssue{}
	return &this
}

// GetOrgUnitId returns the OrgUnitId field value if set, zero value otherwise.
func (o *ActiveIssue) GetOrgUnitId() int32 {
	if o == nil || IsNil(o.OrgUnitId) {
		var ret int32
		return ret
	}
	return *o.OrgUnitId
}

// GetOrgUnitIdOk returns a tuple with the OrgUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveIssue) GetOrgUnitIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgUnitId) {
		return nil, false
	}
	return o.OrgUnitId, true
}

// HasOrgUnitId returns a boolean if a field has been set.
func (o *ActiveIssue) HasOrgUnitId() bool {
	if o != nil && !IsNil(o.OrgUnitId) {
		return true
	}

	return false
}

// SetOrgUnitId gets a reference to the given int32 and assigns it to the OrgUnitId field.
func (o *ActiveIssue) SetOrgUnitId(v int32) {
	o.OrgUnitId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ActiveIssue) GetDeviceId() int32 {
	if o == nil || IsNil(o.DeviceId) {
		var ret int32
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveIssue) GetDeviceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ActiveIssue) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int32 and assigns it to the DeviceId field.
func (o *ActiveIssue) SetDeviceId(v int32) {
	o.DeviceId = &v
}

// GetNotificationState returns the NotificationState field value if set, zero value otherwise.
func (o *ActiveIssue) GetNotificationState() int32 {
	if o == nil || IsNil(o.NotificationState) {
		var ret int32
		return ret
	}
	return *o.NotificationState
}

// GetNotificationStateOk returns a tuple with the NotificationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveIssue) GetNotificationStateOk() (*int32, bool) {
	if o == nil || IsNil(o.NotificationState) {
		return nil, false
	}
	return o.NotificationState, true
}

// HasNotificationState returns a boolean if a field has been set.
func (o *ActiveIssue) HasNotificationState() bool {
	if o != nil && !IsNil(o.NotificationState) {
		return true
	}

	return false
}

// SetNotificationState gets a reference to the given int32 and assigns it to the NotificationState field.
func (o *ActiveIssue) SetNotificationState(v int32) {
	o.NotificationState = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ActiveIssue) GetServiceId() int32 {
	if o == nil || IsNil(o.ServiceId) {
		var ret int32
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveIssue) GetServiceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ActiveIssue) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given int32 and assigns it to the ServiceId field.
func (o *ActiveIssue) SetServiceId(v int32) {
	o.ServiceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ActiveIssue) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveIssue) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ActiveIssue) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ActiveIssue) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *ActiveIssue) GetServiceType() string {
	if o == nil || IsNil(o.ServiceType) {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveIssue) GetServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *ActiveIssue) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *ActiveIssue) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *ActiveIssue) GetTaskId() int32 {
	if o == nil || IsNil(o.TaskId) {
		var ret int32
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveIssue) GetTaskIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *ActiveIssue) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given int32 and assigns it to the TaskId field.
func (o *ActiveIssue) SetTaskId(v int32) {
	o.TaskId = &v
}

// GetServiceItemId returns the ServiceItemId field value if set, zero value otherwise.
func (o *ActiveIssue) GetServiceItemId() int32 {
	if o == nil || IsNil(o.ServiceItemId) {
		var ret int32
		return ret
	}
	return *o.ServiceItemId
}

// GetServiceItemIdOk returns a tuple with the ServiceItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveIssue) GetServiceItemIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceItemId) {
		return nil, false
	}
	return o.ServiceItemId, true
}

// HasServiceItemId returns a boolean if a field has been set.
func (o *ActiveIssue) HasServiceItemId() bool {
	if o != nil && !IsNil(o.ServiceItemId) {
		return true
	}

	return false
}

// SetServiceItemId gets a reference to the given int32 and assigns it to the ServiceItemId field.
func (o *ActiveIssue) SetServiceItemId(v int32) {
	o.ServiceItemId = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *ActiveIssue) GetExtra() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Extra) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveIssue) GetExtraOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Extra) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *ActiveIssue) HasExtra() bool {
	if o != nil && !IsNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]map[string]interface{} and assigns it to the Extra field.
func (o *ActiveIssue) SetExtra(v map[string]map[string]interface{}) {
	o.Extra = v
}

func (o ActiveIssue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveIssue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgUnitId) {
		toSerialize["orgUnitId"] = o.OrgUnitId
	}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.NotificationState) {
		toSerialize["notificationState"] = o.NotificationState
	}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	if !IsNil(o.ServiceType) {
		toSerialize["serviceType"] = o.ServiceType
	}
	if !IsNil(o.TaskId) {
		toSerialize["taskId"] = o.TaskId
	}
	if !IsNil(o.ServiceItemId) {
		toSerialize["serviceItemId"] = o.ServiceItemId
	}
	if !IsNil(o.Extra) {
		toSerialize["_extra"] = o.Extra
	}
	return toSerialize, nil
}

type NullableActiveIssue struct {
	value *ActiveIssue
	isSet bool
}

func (v NullableActiveIssue) Get() *ActiveIssue {
	return v.value
}

func (v *NullableActiveIssue) Set(val *ActiveIssue) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveIssue(val *ActiveIssue) *NullableActiveIssue {
	return &NullableActiveIssue{value: val, isSet: true}
}

func (v NullableActiveIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


