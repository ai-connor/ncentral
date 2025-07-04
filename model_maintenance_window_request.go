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

// checks if the MaintenanceWindowRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaintenanceWindowRequest{}

// MaintenanceWindowRequest Represents a single maintenance Window.
type MaintenanceWindowRequest struct {
	// Actions Applicable To Window
	ApplicableAction []ApplicableAction `json:"applicableAction"`
	// Schedule represented as Cron
	Cron string `json:"cron"`
	// Maintenance Window should last for
	Duration int32 `json:"duration"`
	// Maintenance Window should last for
	Enabled bool `json:"enabled"`
	// Name of Maintenance Window
	Name string `json:"name"`
	// Type of Maintenance Window (only allowed type is currently 'action')
	Type string `json:"type"`
	// Place Device in Downtime During Reboot (Reboot Window)
	DowntimeOnAction *bool `json:"downtimeOnAction,omitempty"`
	// Force Device out of Downtime After (Reboot Window) - used when 'downtimeOnAction' is true
	MaxDowntime *int32 `json:"maxDowntime,omitempty"`
	// Reboot Method (Reboot Window), must be one of: ['allowUserToPostpone', 'forceUserToReboot', 'forceRebootWithoutNotification', 'onlyAcceptedReboot']
	RebootMethod *string `json:"rebootMethod,omitempty"`
	// Minutes before continuing with reboot. (Reboot Window)
	RebootDelay *int32 `json:"rebootDelay,omitempty"`
	// Display Custom Message To User (Reboot Window)
	UserMessageEnabled *bool `json:"userMessageEnabled,omitempty"`
	// Custom message to display to user when 'userMessageEnabled' is true (Reboot Window)
	UserMessage *string `json:"userMessage,omitempty"`
	// Enable Custom Message Sender (Reboot Window)
	MessageSenderEnabled *bool `json:"messageSenderEnabled,omitempty"`
	// Message from when 'messageSenderEnabled' is true (Reboot Window)
	MessageSender *string `json:"messageSender,omitempty"`
	// Preserve State of Device During Reboot (/g flag) (Reboot Window)
	PreserveStateEnabled *bool `json:"preserveStateEnabled,omitempty"`
}

type _MaintenanceWindowRequest MaintenanceWindowRequest

// NewMaintenanceWindowRequest instantiates a new MaintenanceWindowRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindowRequest(applicableAction []ApplicableAction, cron string, duration int32, enabled bool, name string, type_ string) *MaintenanceWindowRequest {
	this := MaintenanceWindowRequest{}
	this.ApplicableAction = applicableAction
	this.Cron = cron
	this.Duration = duration
	this.Enabled = enabled
	this.Name = name
	this.Type = type_
	return &this
}

// NewMaintenanceWindowRequestWithDefaults instantiates a new MaintenanceWindowRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowRequestWithDefaults() *MaintenanceWindowRequest {
	this := MaintenanceWindowRequest{}
	return &this
}

// GetApplicableAction returns the ApplicableAction field value
func (o *MaintenanceWindowRequest) GetApplicableAction() []ApplicableAction {
	if o == nil {
		var ret []ApplicableAction
		return ret
	}

	return o.ApplicableAction
}

// GetApplicableActionOk returns a tuple with the ApplicableAction field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetApplicableActionOk() ([]ApplicableAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicableAction, true
}

// SetApplicableAction sets field value
func (o *MaintenanceWindowRequest) SetApplicableAction(v []ApplicableAction) {
	o.ApplicableAction = v
}

// GetCron returns the Cron field value
func (o *MaintenanceWindowRequest) GetCron() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cron
}

// GetCronOk returns a tuple with the Cron field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetCronOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cron, true
}

// SetCron sets field value
func (o *MaintenanceWindowRequest) SetCron(v string) {
	o.Cron = v
}

// GetDuration returns the Duration field value
func (o *MaintenanceWindowRequest) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *MaintenanceWindowRequest) SetDuration(v int32) {
	o.Duration = v
}

// GetEnabled returns the Enabled field value
func (o *MaintenanceWindowRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MaintenanceWindowRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetName returns the Name field value
func (o *MaintenanceWindowRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MaintenanceWindowRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *MaintenanceWindowRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MaintenanceWindowRequest) SetType(v string) {
	o.Type = v
}

// GetDowntimeOnAction returns the DowntimeOnAction field value if set, zero value otherwise.
func (o *MaintenanceWindowRequest) GetDowntimeOnAction() bool {
	if o == nil || IsNil(o.DowntimeOnAction) {
		var ret bool
		return ret
	}
	return *o.DowntimeOnAction
}

// GetDowntimeOnActionOk returns a tuple with the DowntimeOnAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetDowntimeOnActionOk() (*bool, bool) {
	if o == nil || IsNil(o.DowntimeOnAction) {
		return nil, false
	}
	return o.DowntimeOnAction, true
}

// HasDowntimeOnAction returns a boolean if a field has been set.
func (o *MaintenanceWindowRequest) HasDowntimeOnAction() bool {
	if o != nil && !IsNil(o.DowntimeOnAction) {
		return true
	}

	return false
}

// SetDowntimeOnAction gets a reference to the given bool and assigns it to the DowntimeOnAction field.
func (o *MaintenanceWindowRequest) SetDowntimeOnAction(v bool) {
	o.DowntimeOnAction = &v
}

// GetMaxDowntime returns the MaxDowntime field value if set, zero value otherwise.
func (o *MaintenanceWindowRequest) GetMaxDowntime() int32 {
	if o == nil || IsNil(o.MaxDowntime) {
		var ret int32
		return ret
	}
	return *o.MaxDowntime
}

// GetMaxDowntimeOk returns a tuple with the MaxDowntime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetMaxDowntimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDowntime) {
		return nil, false
	}
	return o.MaxDowntime, true
}

// HasMaxDowntime returns a boolean if a field has been set.
func (o *MaintenanceWindowRequest) HasMaxDowntime() bool {
	if o != nil && !IsNil(o.MaxDowntime) {
		return true
	}

	return false
}

// SetMaxDowntime gets a reference to the given int32 and assigns it to the MaxDowntime field.
func (o *MaintenanceWindowRequest) SetMaxDowntime(v int32) {
	o.MaxDowntime = &v
}

// GetRebootMethod returns the RebootMethod field value if set, zero value otherwise.
func (o *MaintenanceWindowRequest) GetRebootMethod() string {
	if o == nil || IsNil(o.RebootMethod) {
		var ret string
		return ret
	}
	return *o.RebootMethod
}

// GetRebootMethodOk returns a tuple with the RebootMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetRebootMethodOk() (*string, bool) {
	if o == nil || IsNil(o.RebootMethod) {
		return nil, false
	}
	return o.RebootMethod, true
}

// HasRebootMethod returns a boolean if a field has been set.
func (o *MaintenanceWindowRequest) HasRebootMethod() bool {
	if o != nil && !IsNil(o.RebootMethod) {
		return true
	}

	return false
}

// SetRebootMethod gets a reference to the given string and assigns it to the RebootMethod field.
func (o *MaintenanceWindowRequest) SetRebootMethod(v string) {
	o.RebootMethod = &v
}

// GetRebootDelay returns the RebootDelay field value if set, zero value otherwise.
func (o *MaintenanceWindowRequest) GetRebootDelay() int32 {
	if o == nil || IsNil(o.RebootDelay) {
		var ret int32
		return ret
	}
	return *o.RebootDelay
}

// GetRebootDelayOk returns a tuple with the RebootDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetRebootDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.RebootDelay) {
		return nil, false
	}
	return o.RebootDelay, true
}

// HasRebootDelay returns a boolean if a field has been set.
func (o *MaintenanceWindowRequest) HasRebootDelay() bool {
	if o != nil && !IsNil(o.RebootDelay) {
		return true
	}

	return false
}

// SetRebootDelay gets a reference to the given int32 and assigns it to the RebootDelay field.
func (o *MaintenanceWindowRequest) SetRebootDelay(v int32) {
	o.RebootDelay = &v
}

// GetUserMessageEnabled returns the UserMessageEnabled field value if set, zero value otherwise.
func (o *MaintenanceWindowRequest) GetUserMessageEnabled() bool {
	if o == nil || IsNil(o.UserMessageEnabled) {
		var ret bool
		return ret
	}
	return *o.UserMessageEnabled
}

// GetUserMessageEnabledOk returns a tuple with the UserMessageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetUserMessageEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UserMessageEnabled) {
		return nil, false
	}
	return o.UserMessageEnabled, true
}

// HasUserMessageEnabled returns a boolean if a field has been set.
func (o *MaintenanceWindowRequest) HasUserMessageEnabled() bool {
	if o != nil && !IsNil(o.UserMessageEnabled) {
		return true
	}

	return false
}

// SetUserMessageEnabled gets a reference to the given bool and assigns it to the UserMessageEnabled field.
func (o *MaintenanceWindowRequest) SetUserMessageEnabled(v bool) {
	o.UserMessageEnabled = &v
}

// GetUserMessage returns the UserMessage field value if set, zero value otherwise.
func (o *MaintenanceWindowRequest) GetUserMessage() string {
	if o == nil || IsNil(o.UserMessage) {
		var ret string
		return ret
	}
	return *o.UserMessage
}

// GetUserMessageOk returns a tuple with the UserMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetUserMessageOk() (*string, bool) {
	if o == nil || IsNil(o.UserMessage) {
		return nil, false
	}
	return o.UserMessage, true
}

// HasUserMessage returns a boolean if a field has been set.
func (o *MaintenanceWindowRequest) HasUserMessage() bool {
	if o != nil && !IsNil(o.UserMessage) {
		return true
	}

	return false
}

// SetUserMessage gets a reference to the given string and assigns it to the UserMessage field.
func (o *MaintenanceWindowRequest) SetUserMessage(v string) {
	o.UserMessage = &v
}

// GetMessageSenderEnabled returns the MessageSenderEnabled field value if set, zero value otherwise.
func (o *MaintenanceWindowRequest) GetMessageSenderEnabled() bool {
	if o == nil || IsNil(o.MessageSenderEnabled) {
		var ret bool
		return ret
	}
	return *o.MessageSenderEnabled
}

// GetMessageSenderEnabledOk returns a tuple with the MessageSenderEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetMessageSenderEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MessageSenderEnabled) {
		return nil, false
	}
	return o.MessageSenderEnabled, true
}

// HasMessageSenderEnabled returns a boolean if a field has been set.
func (o *MaintenanceWindowRequest) HasMessageSenderEnabled() bool {
	if o != nil && !IsNil(o.MessageSenderEnabled) {
		return true
	}

	return false
}

// SetMessageSenderEnabled gets a reference to the given bool and assigns it to the MessageSenderEnabled field.
func (o *MaintenanceWindowRequest) SetMessageSenderEnabled(v bool) {
	o.MessageSenderEnabled = &v
}

// GetMessageSender returns the MessageSender field value if set, zero value otherwise.
func (o *MaintenanceWindowRequest) GetMessageSender() string {
	if o == nil || IsNil(o.MessageSender) {
		var ret string
		return ret
	}
	return *o.MessageSender
}

// GetMessageSenderOk returns a tuple with the MessageSender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetMessageSenderOk() (*string, bool) {
	if o == nil || IsNil(o.MessageSender) {
		return nil, false
	}
	return o.MessageSender, true
}

// HasMessageSender returns a boolean if a field has been set.
func (o *MaintenanceWindowRequest) HasMessageSender() bool {
	if o != nil && !IsNil(o.MessageSender) {
		return true
	}

	return false
}

// SetMessageSender gets a reference to the given string and assigns it to the MessageSender field.
func (o *MaintenanceWindowRequest) SetMessageSender(v string) {
	o.MessageSender = &v
}

// GetPreserveStateEnabled returns the PreserveStateEnabled field value if set, zero value otherwise.
func (o *MaintenanceWindowRequest) GetPreserveStateEnabled() bool {
	if o == nil || IsNil(o.PreserveStateEnabled) {
		var ret bool
		return ret
	}
	return *o.PreserveStateEnabled
}

// GetPreserveStateEnabledOk returns a tuple with the PreserveStateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowRequest) GetPreserveStateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PreserveStateEnabled) {
		return nil, false
	}
	return o.PreserveStateEnabled, true
}

// HasPreserveStateEnabled returns a boolean if a field has been set.
func (o *MaintenanceWindowRequest) HasPreserveStateEnabled() bool {
	if o != nil && !IsNil(o.PreserveStateEnabled) {
		return true
	}

	return false
}

// SetPreserveStateEnabled gets a reference to the given bool and assigns it to the PreserveStateEnabled field.
func (o *MaintenanceWindowRequest) SetPreserveStateEnabled(v bool) {
	o.PreserveStateEnabled = &v
}

func (o MaintenanceWindowRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaintenanceWindowRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicableAction"] = o.ApplicableAction
	toSerialize["cron"] = o.Cron
	toSerialize["duration"] = o.Duration
	toSerialize["enabled"] = o.Enabled
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.DowntimeOnAction) {
		toSerialize["downtimeOnAction"] = o.DowntimeOnAction
	}
	if !IsNil(o.MaxDowntime) {
		toSerialize["maxDowntime"] = o.MaxDowntime
	}
	if !IsNil(o.RebootMethod) {
		toSerialize["rebootMethod"] = o.RebootMethod
	}
	if !IsNil(o.RebootDelay) {
		toSerialize["rebootDelay"] = o.RebootDelay
	}
	if !IsNil(o.UserMessageEnabled) {
		toSerialize["userMessageEnabled"] = o.UserMessageEnabled
	}
	if !IsNil(o.UserMessage) {
		toSerialize["userMessage"] = o.UserMessage
	}
	if !IsNil(o.MessageSenderEnabled) {
		toSerialize["messageSenderEnabled"] = o.MessageSenderEnabled
	}
	if !IsNil(o.MessageSender) {
		toSerialize["messageSender"] = o.MessageSender
	}
	if !IsNil(o.PreserveStateEnabled) {
		toSerialize["preserveStateEnabled"] = o.PreserveStateEnabled
	}
	return toSerialize, nil
}

func (o *MaintenanceWindowRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicableAction",
		"cron",
		"duration",
		"enabled",
		"name",
		"type",
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

	varMaintenanceWindowRequest := _MaintenanceWindowRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMaintenanceWindowRequest)

	if err != nil {
		return err
	}

	*o = MaintenanceWindowRequest(varMaintenanceWindowRequest)

	return err
}

type NullableMaintenanceWindowRequest struct {
	value *MaintenanceWindowRequest
	isSet bool
}

func (v NullableMaintenanceWindowRequest) Get() *MaintenanceWindowRequest {
	return v.value
}

func (v *NullableMaintenanceWindowRequest) Set(val *MaintenanceWindowRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindowRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindowRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindowRequest(val *MaintenanceWindowRequest) *NullableMaintenanceWindowRequest {
	return &NullableMaintenanceWindowRequest{value: val, isSet: true}
}

func (v NullableMaintenanceWindowRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindowRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


