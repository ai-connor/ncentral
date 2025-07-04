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

// checks if the Filter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Filter{}

// Filter Response for list of device filters.
type Filter struct {
	FilterId *string `json:"filterId,omitempty"`
	FilterName *string `json:"filterName,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewFilter instantiates a new Filter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilter() *Filter {
	this := Filter{}
	return &this
}

// NewFilterWithDefaults instantiates a new Filter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterWithDefaults() *Filter {
	this := Filter{}
	return &this
}

// GetFilterId returns the FilterId field value if set, zero value otherwise.
func (o *Filter) GetFilterId() string {
	if o == nil || IsNil(o.FilterId) {
		var ret string
		return ret
	}
	return *o.FilterId
}

// GetFilterIdOk returns a tuple with the FilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetFilterIdOk() (*string, bool) {
	if o == nil || IsNil(o.FilterId) {
		return nil, false
	}
	return o.FilterId, true
}

// HasFilterId returns a boolean if a field has been set.
func (o *Filter) HasFilterId() bool {
	if o != nil && !IsNil(o.FilterId) {
		return true
	}

	return false
}

// SetFilterId gets a reference to the given string and assigns it to the FilterId field.
func (o *Filter) SetFilterId(v string) {
	o.FilterId = &v
}

// GetFilterName returns the FilterName field value if set, zero value otherwise.
func (o *Filter) GetFilterName() string {
	if o == nil || IsNil(o.FilterName) {
		var ret string
		return ret
	}
	return *o.FilterName
}

// GetFilterNameOk returns a tuple with the FilterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetFilterNameOk() (*string, bool) {
	if o == nil || IsNil(o.FilterName) {
		return nil, false
	}
	return o.FilterName, true
}

// HasFilterName returns a boolean if a field has been set.
func (o *Filter) HasFilterName() bool {
	if o != nil && !IsNil(o.FilterName) {
		return true
	}

	return false
}

// SetFilterName gets a reference to the given string and assigns it to the FilterName field.
func (o *Filter) SetFilterName(v string) {
	o.FilterName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Filter) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Filter) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Filter) SetDescription(v string) {
	o.Description = &v
}

func (o Filter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilterId) {
		toSerialize["filterId"] = o.FilterId
	}
	if !IsNil(o.FilterName) {
		toSerialize["filterName"] = o.FilterName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableFilter struct {
	value *Filter
	isSet bool
}

func (v NullableFilter) Get() *Filter {
	return v.value
}

func (v *NullableFilter) Set(val *Filter) {
	v.value = val
	v.isSet = true
}

func (v NullableFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilter(val *Filter) *NullableFilter {
	return &NullableFilter{value: val, isSet: true}
}

func (v NullableFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


