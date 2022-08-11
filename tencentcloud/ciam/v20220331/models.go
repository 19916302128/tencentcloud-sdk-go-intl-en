// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20220331

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

// Predefined struct for user
type CreateApiImportUserJobRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Imported user data
	DataFlowUserCreateList []*ImportUser `json:"DataFlowUserCreateList,omitempty" name:"DataFlowUserCreateList"`
}

type CreateApiImportUserJobRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Imported user data
	DataFlowUserCreateList []*ImportUser `json:"DataFlowUserCreateList,omitempty" name:"DataFlowUserCreateList"`
}

func (r *CreateApiImportUserJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiImportUserJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "DataFlowUserCreateList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiImportUserJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiImportUserJobResponseParams struct {
	// Data flow task
	Job *Job `json:"Job,omitempty" name:"Job"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateApiImportUserJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiImportUserJobResponseParams `json:"Response"`
}

func (r *CreateApiImportUserJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiImportUserJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileExportUserJobRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Exported data type
	// 
	// <li> **JSON** </li>  JSON
	// <li> **NDJSON** </li>  New-line Delimited JSON
	// <li> **CSV** </li>  Comma-Separated Values
	Format *string `json:"Format,omitempty" name:"Format"`

	// Valid values of `Key`: `condition`, `userGroupId`.
	// 
	// <li> **condition** </li>	Values = Query condition, which can be user ID, username, mobile number, or email address.
	// <li> **userGroupId** </li>	Values = User group ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Attributes and mapping names included in the exported user information. If this parameter is empty, all attributes will be included.
	ExportPropertyMaps []*ExportPropertyMap `json:"ExportPropertyMaps,omitempty" name:"ExportPropertyMaps"`
}

type CreateFileExportUserJobRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Exported data type
	// 
	// <li> **JSON** </li>  JSON
	// <li> **NDJSON** </li>  New-line Delimited JSON
	// <li> **CSV** </li>  Comma-Separated Values
	Format *string `json:"Format,omitempty" name:"Format"`

	// Valid values of `Key`: `condition`, `userGroupId`.
	// 
	// <li> **condition** </li>	Values = Query condition, which can be user ID, username, mobile number, or email address.
	// <li> **userGroupId** </li>	Values = User group ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Attributes and mapping names included in the exported user information. If this parameter is empty, all attributes will be included.
	ExportPropertyMaps []*ExportPropertyMap `json:"ExportPropertyMaps,omitempty" name:"ExportPropertyMaps"`
}

func (r *CreateFileExportUserJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileExportUserJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Format")
	delete(f, "Filters")
	delete(f, "ExportPropertyMaps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileExportUserJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileExportUserJobResponseParams struct {
	// Data flow task
	Job *Job `json:"Job,omitempty" name:"Job"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFileExportUserJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileExportUserJobResponseParams `json:"Response"`
}

func (r *CreateFileExportUserJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileExportUserJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Mobile number
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// Email address
	Email *string `json:"Email,omitempty" name:"Email"`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Nickname
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// Address
	Address *string `json:"Address,omitempty" name:"Address"`

	// User group ID
	UserGroup []*string `json:"UserGroup,omitempty" name:"UserGroup"`

	// Date of birth
	Birthdate *int64 `json:"Birthdate,omitempty" name:"Birthdate"`

	// Custom attribute
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitempty" name:"CustomizationAttributes"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Mobile number
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// Email address
	Email *string `json:"Email,omitempty" name:"Email"`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Nickname
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// Address
	Address *string `json:"Address,omitempty" name:"Address"`

	// User group ID
	UserGroup []*string `json:"UserGroup,omitempty" name:"UserGroup"`

	// Date of birth
	Birthdate *int64 `json:"Birthdate,omitempty" name:"Birthdate"`

	// Custom attribute
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitempty" name:"CustomizationAttributes"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "PhoneNumber")
	delete(f, "Email")
	delete(f, "Password")
	delete(f, "UserName")
	delete(f, "Nickname")
	delete(f, "Address")
	delete(f, "UserGroup")
	delete(f, "Birthdate")
	delete(f, "CustomizationAttributes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// Information of the created user
	// Note: This field may return null, indicating that no valid values can be obtained.
	User *User `json:"User,omitempty" name:"User"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Array of user IDs
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Array of user IDs
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *DeleteUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteUsersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsersResponseParams `json:"Response"`
}

func (r *DeleteUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserByIdRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`


	Original *bool `json:"Original,omitempty" name:"Original"`
}

type DescribeUserByIdRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	Original *bool `json:"Original,omitempty" name:"Original"`
}

func (r *DescribeUserByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserId")
	delete(f, "Original")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserByIdResponseParams struct {
	// User information
	// Note: This field may return null, indicating that no valid values can be obtained.
	User *User `json:"User,omitempty" name:"User"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserByIdResponseParams `json:"Response"`
}

func (r *DescribeUserByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorDetails struct {
	// User information
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Failure cause
	Error *string `json:"Error,omitempty" name:"Error"`
}

type ExportPropertyMap struct {
	// User attribute code
	UserPropertyCode *string `json:"UserPropertyCode,omitempty" name:"UserPropertyCode"`

	// User attribute mapping name
	ColumnName *string `json:"ColumnName,omitempty" name:"ColumnName"`
}

type FailedUsers struct {
	// ID of the failed user
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailedUserIdentification *string `json:"FailedUserIdentification,omitempty" name:"FailedUserIdentification"`

	// Failure cause for user import
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailedReason *string `json:"FailedReason,omitempty" name:"FailedReason"`
}

type Filter struct {
	// Key value
	Key *string `json:"Key,omitempty" name:"Key"`

	// Value
	Values []*string `json:"Values,omitempty" name:"Values"`

	// Logical value
	Logic *bool `json:"Logic,omitempty" name:"Logic"`
}

type ImportUser struct {
	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Mobile number
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// Email address
	Email *string `json:"Email,omitempty" name:"Email"`

	// ID card number
	ResidentIdentityCard *string `json:"ResidentIdentityCard,omitempty" name:"ResidentIdentityCard"`

	// Nickname
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// Address
	Address *string `json:"Address,omitempty" name:"Address"`

	// User group ID
	UserGroup []*string `json:"UserGroup,omitempty" name:"UserGroup"`

	// `qqOpenId` on QQ
	QqOpenId *string `json:"QqOpenId,omitempty" name:"QqOpenId"`

	// `qqUnionId` on QQ
	QqUnionId *string `json:"QqUnionId,omitempty" name:"QqUnionId"`

	// `wechatOpenId` on WeChat
	WechatOpenId *string `json:"WechatOpenId,omitempty" name:"WechatOpenId"`

	// `wechatUnionId` on WeChat
	WechatUnionId *string `json:"WechatUnionId,omitempty" name:"WechatUnionId"`

	// `alipayUserId` on Alipay
	AlipayUserId *string `json:"AlipayUserId,omitempty" name:"AlipayUserId"`

	// Description
	Description *string `json:"Description,omitempty" name:"Description"`

	// Date of birth
	Birthdate *string `json:"Birthdate,omitempty" name:"Birthdate"`

	// Name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Coordinate
	Locale *string `json:"Locale,omitempty" name:"Locale"`

	// Gender. Valid values: `MALE`, `FEMALE`, `UNKNOWN`.
	Gender *string `json:"Gender,omitempty" name:"Gender"`

	// Identity verification method
	IdentityVerificationMethod *string `json:"IdentityVerificationMethod,omitempty" name:"IdentityVerificationMethod"`

	// Whether the identity is verified
	IdentityVerified *bool `json:"IdentityVerified,omitempty" name:"IdentityVerified"`

	// Job
	Job *string `json:"Job,omitempty" name:"Job"`

	// Country/Region
	Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

	// Time zone
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Password ciphertext
	Password *string `json:"Password,omitempty" name:"Password"`

	// Custom attribute
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitempty" name:"CustomizationAttributes"`

	// Password salt
	Salt *Salt `json:"Salt,omitempty" name:"Salt"`

	// Password encryption method. Valid values: `SHA1`, `BCRYPT`.
	PasswordEncryptTypeEnum *string `json:"PasswordEncryptTypeEnum,omitempty" name:"PasswordEncryptTypeEnum"`
}

type Job struct {
	// Task ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// Task status
	// 
	// <li> **PENDING** </li>  Pending
	// <li> **PROCESSING** </li>  Executing
	// <li> **COMPLETED** </li>  Completed
	// <li> **FAILED** </li>  Failed
	Status *string `json:"Status,omitempty" name:"Status"`

	// Task type
	// 
	// <li> **IMPORT_USER** </li>  User import
	// <li> **EXPORT_USER** </li>  User export
	Type *string `json:"Type,omitempty" name:"Type"`

	// Task creation time
	CreatedDate *int64 `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// Data type of the task
	// 
	// <li> **JSON** </li>  JSON
	// <li> **NDJSON** </li>  New-line Delimited JSON
	// <li> **CSV** </li>  Comma-Separated Values
	// Note: This field may return null, indicating that no valid values can be obtained.
	Format *string `json:"Format,omitempty" name:"Format"`

	// Task result download address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Location *string `json:"Location,omitempty" name:"Location"`

	// Failure details
	// Note: This field may return null, indicating that no valid values can be obtained.
	ErrorDetails []*ErrorDetails `json:"ErrorDetails,omitempty" name:"ErrorDetails"`

	// Failed user
	// Note: This field may return null, indicating that no valid values can be obtained.
	FailedUsers []*FailedUsers `json:"FailedUsers,omitempty" name:"FailedUsers"`
}

// Predefined struct for user
type LinkAccountRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Primary user ID
	PrimaryUserId *string `json:"PrimaryUserId,omitempty" name:"PrimaryUserId"`

	// Secondary user ID
	SecondaryUserId *string `json:"SecondaryUserId,omitempty" name:"SecondaryUserId"`

	// Fusion attribute
	// 
	// <li> **PHONENUMBER** </li>	  Mobile number
	// <li> **EMAIL** </li>  Email
	UserLinkedOnAttribute *string `json:"UserLinkedOnAttribute,omitempty" name:"UserLinkedOnAttribute"`
}

type LinkAccountRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Primary user ID
	PrimaryUserId *string `json:"PrimaryUserId,omitempty" name:"PrimaryUserId"`

	// Secondary user ID
	SecondaryUserId *string `json:"SecondaryUserId,omitempty" name:"SecondaryUserId"`

	// Fusion attribute
	// 
	// <li> **PHONENUMBER** </li>	  Mobile number
	// <li> **EMAIL** </li>  Email
	UserLinkedOnAttribute *string `json:"UserLinkedOnAttribute,omitempty" name:"UserLinkedOnAttribute"`
}

func (r *LinkAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "PrimaryUserId")
	delete(f, "SecondaryUserId")
	delete(f, "UserLinkedOnAttribute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LinkAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LinkAccountResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LinkAccountResponse struct {
	*tchttp.BaseResponse
	Response *LinkAccountResponseParams `json:"Response"`
}

func (r *LinkAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListJobsRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// List of task IDs. If this parameter is empty, all tasks will be returned.
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`
}

type ListJobsRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// List of task IDs. If this parameter is empty, all tasks will be returned.
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`
}

func (r *ListJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "JobIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListJobsResponseParams struct {
	// List of tasks
	// Note: This field may return null, indicating that no valid values can be obtained.
	JobSet []*Job `json:"JobSet,omitempty" name:"JobSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListJobsResponse struct {
	*tchttp.BaseResponse
	Response *ListJobsResponseParams `json:"Response"`
}

func (r *ListJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLogMessageByConditionRequestParams struct {
	// User pool ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Pagination data
	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`

	// Start timestamp accurate to the millisecond
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Valid values of `Key`: `events`.
	// 
	// <li> **events** </li>	Values can be one or multiple items in ["SIGNUP", "USER_UPDATE", "USER_DELETE", "USER_CREATE", "ACCOUNT_LINKING"].
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type ListLogMessageByConditionRequest struct {
	*tchttp.BaseRequest
	
	// User pool ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Pagination data
	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`

	// Start timestamp accurate to the millisecond
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// Valid values of `Key`: `events`.
	// 
	// <li> **events** </li>	Values can be one or multiple items in ["SIGNUP", "USER_UPDATE", "USER_DELETE", "USER_CREATE", "ACCOUNT_LINKING"].
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *ListLogMessageByConditionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogMessageByConditionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Pageable")
	delete(f, "StartTime")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListLogMessageByConditionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLogMessageByConditionResponseParams struct {
	// Total number
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Pagination object
	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`

	// List of logs
	// Note: This field may return null, indicating that no valid values can be obtained.
	Content []*LogMessage `json:"Content,omitempty" name:"Content"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListLogMessageByConditionResponse struct {
	*tchttp.BaseResponse
	Response *ListLogMessageByConditionResponseParams `json:"Response"`
}

func (r *ListLogMessageByConditionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogMessageByConditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserByPropertyRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Query attribute
	// 
	// <li> **phoneNumber** </li>	  Mobile number
	// <li> **email** </li>  Email
	PropertyCode *string `json:"PropertyCode,omitempty" name:"PropertyCode"`

	// Attribute value
	PropertyValue *string `json:"PropertyValue,omitempty" name:"PropertyValue"`


	Original *bool `json:"Original,omitempty" name:"Original"`
}

type ListUserByPropertyRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Query attribute
	// 
	// <li> **phoneNumber** </li>	  Mobile number
	// <li> **email** </li>  Email
	PropertyCode *string `json:"PropertyCode,omitempty" name:"PropertyCode"`

	// Attribute value
	PropertyValue *string `json:"PropertyValue,omitempty" name:"PropertyValue"`

	Original *bool `json:"Original,omitempty" name:"Original"`
}

func (r *ListUserByPropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserByPropertyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "PropertyCode")
	delete(f, "PropertyValue")
	delete(f, "Original")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserByPropertyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserByPropertyResponseParams struct {
	// List of users
	// Note: This field may return null, indicating that no valid values can be obtained.
	Users []*User `json:"Users,omitempty" name:"Users"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListUserByPropertyResponse struct {
	*tchttp.BaseResponse
	Response *ListUserByPropertyResponseParams `json:"Response"`
}

func (r *ListUserByPropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserByPropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Pagination data
	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`

	// Valid values of `Key`: `condition`, `userGroupId`.
	// 
	// <li> **condition** </li>	Values = Query condition, which can be user ID, username, mobile number, or email address.
	// <li> **userGroupId** </li>	Values = User group ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`


	Original *bool `json:"Original,omitempty" name:"Original"`
}

type ListUserRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Pagination data
	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`

	// Valid values of `Key`: `condition`, `userGroupId`.
	// 
	// <li> **condition** </li>	Values = Query condition, which can be user ID, username, mobile number, or email address.
	// <li> **userGroupId** </li>	Values = User group ID
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	Original *bool `json:"Original,omitempty" name:"Original"`
}

func (r *ListUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Pageable")
	delete(f, "Filters")
	delete(f, "Original")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserResponseParams struct {
	// Total number
	// Note: This field may return null, indicating that no valid values can be obtained.
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// Pagination object
	// Note: This field may return null, indicating that no valid values can be obtained.
	Pageable *Pageable `json:"Pageable,omitempty" name:"Pageable"`

	// List of users
	// Note: This field may return null, indicating that no valid values can be obtained.
	Content []*User `json:"Content,omitempty" name:"Content"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListUserResponse struct {
	*tchttp.BaseResponse
	Response *ListUserResponseParams `json:"Response"`
}

func (r *ListUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogMessage struct {
	// Log ID
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// Tenant ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// User pool ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Event code
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventCode *string `json:"EventCode,omitempty" name:"EventCode"`

	// Event timestamp in milliseconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	EventDate *int64 `json:"EventDate,omitempty" name:"EventDate"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Event participant
	// 
	// <li> **TENANT** </li>  Tenant
	// <li> **USER** </li>  User
	// Note: This field may return null, indicating that no valid values can be obtained.
	Participant *string `json:"Participant,omitempty" name:"Participant"`

	// Application `clientId`
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationClientId *string `json:"ApplicationClientId,omitempty" name:"ApplicationClientId"`

	// Application name
	// Note: This field may return null, indicating that no valid values can be obtained.
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// Authentication source ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuthSourceId *string `json:"AuthSourceId,omitempty" name:"AuthSourceId"`

	// Authentication source name
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuthSourceName *string `json:"AuthSourceName,omitempty" name:"AuthSourceName"`

	// Authentication source type
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuthSourceType *string `json:"AuthSourceType,omitempty" name:"AuthSourceType"`

	// Authentication source category
	// Note: This field may return null, indicating that no valid values can be obtained.
	AuthSourceCategory *string `json:"AuthSourceCategory,omitempty" name:"AuthSourceCategory"`

	// IP address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// User agent
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// User ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Details
	// Note: This field may return null, indicating that no valid values can be obtained.
	Detail *string `json:"Detail,omitempty" name:"Detail"`
}

type MemberMap struct {
	// Key
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value
	Value *string `json:"Value,omitempty" name:"Value"`

	// Type
	// Note: This field may return null, indicating that no valid values can be obtained.
	Type *string `json:"Type,omitempty" name:"Type"`
}

type Pageable struct {
	// Number of entries per page
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// Current page number
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
}

// Predefined struct for user
type ResetPasswordRequestParams struct {
	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserStoreId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetPasswordResponseParams struct {
	// User password after reset
	Password *string `json:"Password,omitempty" name:"Password"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetPasswordResponseParams `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Salt struct {
	// Salt value
	SaltValue *string `json:"SaltValue,omitempty" name:"SaltValue"`

	// Salt value location
	SaltLocation *SaltLocation `json:"SaltLocation,omitempty" name:"SaltLocation"`
}

type SaltLocation struct {
	// Password salt type. Valid values: `HEAD`, `TAIL`, `OTHER`.
	SaltLocationTypeEnum *string `json:"SaltLocationTypeEnum,omitempty" name:"SaltLocationTypeEnum"`

	// Salting rule
	SaltLocationRule *SaltLocationRule `json:"SaltLocationRule,omitempty" name:"SaltLocationRule"`
}

type SaltLocationRule struct {
	// Expression
	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

// Predefined struct for user
type SetPasswordRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`
}

type SetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *SetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserId")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPasswordResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *SetPasswordResponseParams `json:"Response"`
}

func (r *SetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserRequestParams struct {
	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Mobile number
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// Email address
	Email *string `json:"Email,omitempty" name:"Email"`

	// Nickname
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// Address
	Address *string `json:"Address,omitempty" name:"Address"`

	// User group
	UserGroup []*string `json:"UserGroup,omitempty" name:"UserGroup"`

	// Date of birth
	Birthdate *int64 `json:"Birthdate,omitempty" name:"Birthdate"`

	// Custom attribute
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitempty" name:"CustomizationAttributes"`
}

type UpdateUserRequest struct {
	*tchttp.BaseRequest
	
	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Username
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Mobile number
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// Email address
	Email *string `json:"Email,omitempty" name:"Email"`

	// Nickname
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// Address
	Address *string `json:"Address,omitempty" name:"Address"`

	// User group
	UserGroup []*string `json:"UserGroup,omitempty" name:"UserGroup"`

	// Date of birth
	Birthdate *int64 `json:"Birthdate,omitempty" name:"Birthdate"`

	// Custom attribute
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitempty" name:"CustomizationAttributes"`
}

func (r *UpdateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserStoreId")
	delete(f, "UserName")
	delete(f, "PhoneNumber")
	delete(f, "Email")
	delete(f, "Nickname")
	delete(f, "Address")
	delete(f, "UserGroup")
	delete(f, "Birthdate")
	delete(f, "CustomizationAttributes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserResponseParams struct {
	// User information after update
	// Note: This field may return null, indicating that no valid values can be obtained.
	User *User `json:"User,omitempty" name:"User"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateUserResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserResponseParams `json:"Response"`
}

func (r *UpdateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserStatusRequestParams struct {
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// User status
	// 
	// <li> **NORMAL** </li>	  Normal
	// <li> **LOCK** </li>  Locked
	// <li> **FREEZE** </li>	  Frozen
	Status *string `json:"Status,omitempty" name:"Status"`
}

type UpdateUserStatusRequest struct {
	*tchttp.BaseRequest
	
	// User directory ID
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// User status
	// 
	// <li> **NORMAL** </li>	  Normal
	// <li> **LOCK** </li>  Locked
	// <li> **FREEZE** </li>	  Frozen
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateUserStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserStatusResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateUserStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserStatusResponseParams `json:"Response"`
}

func (r *UpdateUserStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {
	// User ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// Username
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// Mobile number
	// Note: This field may return null, indicating that no valid values can be obtained.
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// Email address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Email *string `json:"Email,omitempty" name:"Email"`

	// Last login time
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastSignOn *int64 `json:"LastSignOn,omitempty" name:"LastSignOn"`

	// Creation time
	CreatedDate *int64 `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// Status
	Status *string `json:"Status,omitempty" name:"Status"`

	// User source
	UserDataSourceEnum *string `json:"UserDataSourceEnum,omitempty" name:"UserDataSourceEnum"`

	// Nickname
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// Address
	// Note: This field may return null, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitempty" name:"Address"`

	// Date of birth
	// Note: This field may return null, indicating that no valid values can be obtained.
	Birthdate *int64 `json:"Birthdate,omitempty" name:"Birthdate"`

	// User group ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserGroups []*string `json:"UserGroups,omitempty" name:"UserGroups"`

	// Last modified time
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastModifiedDate *int64 `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// Custom attribute
	// Note: This field may return null, indicating that no valid values can be obtained.
	CustomAttributes []*MemberMap `json:"CustomAttributes,omitempty" name:"CustomAttributes"`

	// ID card number
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResidentIdentityCard *string `json:"ResidentIdentityCard,omitempty" name:"ResidentIdentityCard"`

	// `OpenId` on QQ
	// Note: This field may return null, indicating that no valid values can be obtained.
	QqOpenId *string `json:"QqOpenId,omitempty" name:"QqOpenId"`

	// `UnionId` on QQ
	// Note: This field may return null, indicating that no valid values can be obtained.
	QqUnionId *string `json:"QqUnionId,omitempty" name:"QqUnionId"`

	// `WechatOpenId` on WeChat
	// Note: This field may return null, indicating that no valid values can be obtained.
	WechatOpenId *string `json:"WechatOpenId,omitempty" name:"WechatOpenId"`

	// `WechatUnionId` on WeChat
	// Note: This field may return null, indicating that no valid values can be obtained.
	WechatUnionId *string `json:"WechatUnionId,omitempty" name:"WechatUnionId"`

	// `AlipayUserId` on Alipay
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlipayUserId *string `json:"AlipayUserId,omitempty" name:"AlipayUserId"`

	// Description
	// Note: This field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitempty" name:"Description"`

	// Name
	// Note: This field may return null, indicating that no valid values can be obtained.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Coordinate
	// Note: This field may return null, indicating that no valid values can be obtained.
	Locale *string `json:"Locale,omitempty" name:"Locale"`

	// Gender
	// Note: This field may return null, indicating that no valid values can be obtained.
	Gender *string `json:"Gender,omitempty" name:"Gender"`

	// Identity verification method
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityVerificationMethod *string `json:"IdentityVerificationMethod,omitempty" name:"IdentityVerificationMethod"`

	// Whether the identity is verified
	// Note: This field may return null, indicating that no valid values can be obtained.
	IdentityVerified *bool `json:"IdentityVerified,omitempty" name:"IdentityVerified"`

	// Job
	// Note: This field may return null, indicating that no valid values can be obtained.
	Job *string `json:"Job,omitempty" name:"Job"`

	// Country/Region
	// Note: This field may return null, indicating that no valid values can be obtained.
	Nationality *string `json:"Nationality,omitempty" name:"Nationality"`

	// Whether the account is the primary account (after account merge, this parameter is `true` for primary accounts and `false` for secondary account).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Primary *bool `json:"Primary,omitempty" name:"Primary"`

	// Time zone
	// Note: This field may return null, indicating that no valid values can be obtained.
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// Whether the account has ever logged in.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AlreadyFirstLogin *bool `json:"AlreadyFirstLogin,omitempty" name:"AlreadyFirstLogin"`

	// Tenant ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// User directory ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserStoreId *string `json:"UserStoreId,omitempty" name:"UserStoreId"`

	// Version
	// Note: This field may return null, indicating that no valid values can be obtained.
	Version *int64 `json:"Version,omitempty" name:"Version"`

	// Lock type (locked by admin or locked by login policy)
	// Note: This field may return null, indicating that no valid values can be obtained.
	LockType *string `json:"LockType,omitempty" name:"LockType"`

	// Lock time
	// Note: This field may return null, indicating that no valid values can be obtained.
	LockTime *int64 `json:"LockTime,omitempty" name:"LockTime"`
}