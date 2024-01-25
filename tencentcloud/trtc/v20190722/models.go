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

package v20190722

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type AbnormalEvent struct {
	// The error event ID. For details, see https://www.tencentcloud.com/document/product/647/37906?has_map=1
	AbnormalEventId *uint64 `json:"AbnormalEventId,omitnil" name:"AbnormalEventId"`

	// The remote user ID. If this parameter is empty, it indicates that the error event is not associated with a remote user.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PeerId *string `json:"PeerId,omitnil" name:"PeerId"`
}

type AbnormalExperience struct {
	// The user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The abnormal experience ID.
	ExperienceId *uint64 `json:"ExperienceId,omitnil" name:"ExperienceId"`

	// The room ID (string).
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The possible error events.
	AbnormalEventList []*AbnormalEvent `json:"AbnormalEventList,omitnil" name:"AbnormalEventList"`

	// The report time.
	EventTime *uint64 `json:"EventTime,omitnil" name:"EventTime"`
}

type AgentParams struct {
	// The [user ID](https://intl.cloud.tencent.com/document/product/647/37714) of the relaying robot in the TRTC room, which cannot be the same as a user ID already in use. We recommend you include the room ID in this user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The signature (similar to a login password) required for the relaying robot to enter the room. For information on how to calculate the signature, see [What is UserSig?](https://intl.cloud.tencent.com/document/product/647/38104). |
	UserSig *string `json:"UserSig,omitnil" name:"UserSig"`

	// The timeout period (seconds) for relaying to stop automatically after all the users whose streams are mixed leave the room. The value cannot be smaller than 5 or larger than 86400 (24 hours). Default value: 30.
	MaxIdleTime *uint64 `json:"MaxIdleTime,omitnil" name:"MaxIdleTime"`
}

type AudioEncode struct {
	// The audio sample rate (Hz). Valid values: 48000, 44100, 32000, 24000, 16000, 8000.
	SampleRate *uint64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// The number of sound channels. Valid values: 1 (mono), 2 (dual).
	Channel *uint64 `json:"Channel,omitnil" name:"Channel"`

	// The audio bitrate (Kbps). Value range: 8-500.
	BitRate *uint64 `json:"BitRate,omitnil" name:"BitRate"`

	// The audio codec. Valid values: 0 (LC-AAC), 1 (HE-AAC), 2 (HE-AACv2). The default value is 0. If this parameter is set to 2, `Channel` must be 2. If it is set to 1 or 2, `SampleRate` can only be 48000, 44100, 32000, 24000, or 16000.
	Codec *uint64 `json:"Codec,omitnil" name:"Codec"`
}

type AudioEncodeParams struct {
	// Audio Sample rate, Value range [48000, 44100], unit is Hz.
	SampleRate *uint64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// Audio Channel number, Value range [1,2], 1 means Audio is Mono-channel, 2 means Audio is Dual-channel.
	Channel *uint64 `json:"Channel,omitnil" name:"Channel"`

	// Audio Bitrate, Value range [8,500], unit is kbps.
	BitRate *uint64 `json:"BitRate,omitnil" name:"BitRate"`
}

type AudioParams struct {
	// The audio sample rate.
	// 1: 48000 Hz (default)
	// 2: 44100 Hz
	// 3: 16000 Hz
	SampleRate *uint64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// The number of sound channels.
	// 1: Mono-channel
	// 2: Dual-channel (default)
	Channel *uint64 `json:"Channel,omitnil" name:"Channel"`

	// The audio bitrate (bps). Value range: [32000, 128000]. Default: 64000.
	BitRate *uint64 `json:"BitRate,omitnil" name:"BitRate"`
}

type CloudStorage struct {
	// The cloud storage provider.
	// `0`: Tencent Cloud COS; `1`: AWS storage. Other vendors are not supported currently.
	Vendor *uint64 `json:"Vendor,omitnil" name:"Vendor"`

	// The region of cloud storage.
	Region *string `json:"Region,omitnil" name:"Region"`

	// The storage bucket.
	Bucket *string `json:"Bucket,omitnil" name:"Bucket"`

	// The access_key of the cloud storage account.
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// The secret_key of the cloud storage account.
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// The bucket to save data, which is an array of strings that can contain letters (a-z and A-Z), numbers (0-9), underscores (_), and hyphens (-). For example, if the value of this parameter is `["prefix1", "prefix2"]`, the recording file `xxx.m3u8` will be saved as `prefix1/prefix2/TaskId/xxx.m3u8`.
	FileNamePrefix []*string `json:"FileNamePrefix,omitnil" name:"FileNamePrefix"`
}

type CloudVod struct {
	// The Tencent Cloud VOD parameters.
	TencentVod *TencentVod `json:"TencentVod,omitnil" name:"TencentVod"`
}

// Predefined struct for user
type CreateCloudRecordingRequestParams struct {
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The [room ID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The [user ID](https://www.tencentcloud.com/document/product/647/37714#userid) of the recording robot in the TRTC room, which cannot be identical to the user IDs of anchors in the room or other recording robots. To distinguish this user ID from others, we recommend you include the room ID in the user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The signature (similar to a login password) required for the recording robot to enter the room. Each user ID corresponds to a signature. For information on how to calculate the signature, see [What is UserSig?](https://intl.cloud.tencent.com/document/product/647/38104).
	UserSig *string `json:"UserSig,omitnil" name:"UserSig"`

	// The on-cloud recording parameters.
	RecordParams *RecordParams `json:"RecordParams,omitnil" name:"RecordParams"`

	// The storage information of the recording file. Currently, you can save recording files to Tencent Cloud VOD or COS.
	StorageParams *StorageParams `json:"StorageParams,omitnil" name:"StorageParams"`

	// The type of the TRTC room ID, which must be the same as the ID type of the room whose streams are recorded.
	// 0: String
	// 1: 32-bit integer (default)
	RoomIdType *uint64 `json:"RoomIdType,omitnil" name:"RoomIdType"`

	// The stream mixing parameters, which are valid if the mixed-stream recording mode is used.
	MixTranscodeParams *MixTranscodeParams `json:"MixTranscodeParams,omitnil" name:"MixTranscodeParams"`

	// The layout parameters, which are valid if the mixed-stream recording mode is used.
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitnil" name:"MixLayoutParams"`

	// The amount of time (in hours) during which API requests can be made after recording starts. Calculation starts when a recording task is started (when the recording task ID is returned). Once the period elapses, the query, modification, and stop recording APIs can no longer be called, but the recording task will continue. The default value is `72` (three days), and the maximum and minimum values allowed are `720` (30 days) and `6` respectively. If you do not set this parameter, the query, modification, and stop recording APIs can be called within 72 hours after recording starts.
	ResourceExpiredHour *uint64 `json:"ResourceExpiredHour,omitnil" name:"ResourceExpiredHour"`

	// The permission ticket for a TRTC room. This parameter is required if advanced permission control is enabled in the console, in which case the TRTC backend will verify users’ [PrivateMapKey](https://intl.cloud.tencent.com/document/product/647/32240?from_cn_redirect=1), which include an encrypted room ID and permission bit list. A user providing only `UserSig` and not `PrivateMapKey` will be unable to enter the room.
	PrivateMapKey *string `json:"PrivateMapKey,omitnil" name:"PrivateMapKey"`
}

type CreateCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The [room ID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are recorded.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The [user ID](https://www.tencentcloud.com/document/product/647/37714#userid) of the recording robot in the TRTC room, which cannot be identical to the user IDs of anchors in the room or other recording robots. To distinguish this user ID from others, we recommend you include the room ID in the user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The signature (similar to a login password) required for the recording robot to enter the room. Each user ID corresponds to a signature. For information on how to calculate the signature, see [What is UserSig?](https://intl.cloud.tencent.com/document/product/647/38104).
	UserSig *string `json:"UserSig,omitnil" name:"UserSig"`

	// The on-cloud recording parameters.
	RecordParams *RecordParams `json:"RecordParams,omitnil" name:"RecordParams"`

	// The storage information of the recording file. Currently, you can save recording files to Tencent Cloud VOD or COS.
	StorageParams *StorageParams `json:"StorageParams,omitnil" name:"StorageParams"`

	// The type of the TRTC room ID, which must be the same as the ID type of the room whose streams are recorded.
	// 0: String
	// 1: 32-bit integer (default)
	RoomIdType *uint64 `json:"RoomIdType,omitnil" name:"RoomIdType"`

	// The stream mixing parameters, which are valid if the mixed-stream recording mode is used.
	MixTranscodeParams *MixTranscodeParams `json:"MixTranscodeParams,omitnil" name:"MixTranscodeParams"`

	// The layout parameters, which are valid if the mixed-stream recording mode is used.
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitnil" name:"MixLayoutParams"`

	// The amount of time (in hours) during which API requests can be made after recording starts. Calculation starts when a recording task is started (when the recording task ID is returned). Once the period elapses, the query, modification, and stop recording APIs can no longer be called, but the recording task will continue. The default value is `72` (three days), and the maximum and minimum values allowed are `720` (30 days) and `6` respectively. If you do not set this parameter, the query, modification, and stop recording APIs can be called within 72 hours after recording starts.
	ResourceExpiredHour *uint64 `json:"ResourceExpiredHour,omitnil" name:"ResourceExpiredHour"`

	// The permission ticket for a TRTC room. This parameter is required if advanced permission control is enabled in the console, in which case the TRTC backend will verify users’ [PrivateMapKey](https://intl.cloud.tencent.com/document/product/647/32240?from_cn_redirect=1), which include an encrypted room ID and permission bit list. A user providing only `UserSig` and not `PrivateMapKey` will be unable to enter the room.
	PrivateMapKey *string `json:"PrivateMapKey,omitnil" name:"PrivateMapKey"`
}

func (r *CreateCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserId")
	delete(f, "UserSig")
	delete(f, "RecordParams")
	delete(f, "StorageParams")
	delete(f, "RoomIdType")
	delete(f, "MixTranscodeParams")
	delete(f, "MixLayoutParams")
	delete(f, "ResourceExpiredHour")
	delete(f, "PrivateMapKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudRecordingResponseParams struct {
	// The task ID assigned by the recording service, which uniquely identifies a recording process and becomes invalid after a recording task ends. After a recording task starts, if you want to perform other actions on the task, you need to specify the task ID when making API requests.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudRecordingResponseParams `json:"Response"`
}

func (r *CreateCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudRecordingRequestParams struct {
	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DeleteCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DeleteCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudRecordingResponseParams struct {
	// The task ID assigned by the recording service, which uniquely identifies a recording process and becomes invalid after a recording task ends.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudRecordingResponseParams `json:"Response"`
}

func (r *DeleteCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallDetailInfoRequestParams struct {
	// The unique ID of a call, whose format is `SdkAppId_CreateTime`, such as `1400xxxxxx_218695_1590065777`. `createTime` is the UNIX timestamp (seconds) when the room was created. Its value can be obtained using the [DescribeRoomInfo](https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1) API.
	CommId *string `json:"CommId,omitnil" name:"CommId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`.
	// Note: If `DataType` is not null, the end time and start time cannot be more than one hour apart; if `DataType` is null, the end time and start time cannot be more than four hours apart.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The users to query. If you do not specify this, the data of six users will be returned.
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// The metrics to query. If you do not specify this, only the user list will be returned. If you pass in `all`, all metrics will be returned.
	// `appCpu`: The CPU utilization of the application.
	// `sysCpu`: The CPU utilization of the system.
	// `aBit`: The upstream/downstream audio bitrate (bps).
	// `aBlock`: The audio stutter duration (ms).
	// `bigvBit`: The upstream/downstream video bitrate (bps).
	// `bigvCapFps`: The frame rate for capturing videos.
	// `bigvEncFps`: The frame rate for sending videos.
	// `bigvDecFps`: The rendering frame rate.
	// `bigvBlock`: The video stutter duration (ms).
	// `aLoss`: The upstream/downstream audio packet loss.
	// `bigvLoss`: The upstream/downstream video packet loss.
	// `bigvWidth`: The upstream/downstream resolution (width).
	// `bigvHeight`: The upstream/downstream resolution (height).
	DataType []*string `json:"DataType,omitnil" name:"DataType"`

	// The page number. The default is 0.
	// Note: If `PageNumber` or `PageSize` is not specified, six records will be returned.
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of records per page. The default is `6`.
	// Value range: 1-100.
	// Note: If `DataType` is not null, the length of the array `UserIds` and the value of `PageSize` cannot exceed `6`.
	// If `DataType` is null, the length of the array `UserIds` and the value of `PageSize` cannot exceed `100`.
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeCallDetailInfoRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a call, whose format is `SdkAppId_CreateTime`, such as `1400xxxxxx_218695_1590065777`. `createTime` is the UNIX timestamp (seconds) when the room was created. Its value can be obtained using the [DescribeRoomInfo](https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1) API.
	CommId *string `json:"CommId,omitnil" name:"CommId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`.
	// Note: If `DataType` is not null, the end time and start time cannot be more than one hour apart; if `DataType` is null, the end time and start time cannot be more than four hours apart.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The users to query. If you do not specify this, the data of six users will be returned.
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// The metrics to query. If you do not specify this, only the user list will be returned. If you pass in `all`, all metrics will be returned.
	// `appCpu`: The CPU utilization of the application.
	// `sysCpu`: The CPU utilization of the system.
	// `aBit`: The upstream/downstream audio bitrate (bps).
	// `aBlock`: The audio stutter duration (ms).
	// `bigvBit`: The upstream/downstream video bitrate (bps).
	// `bigvCapFps`: The frame rate for capturing videos.
	// `bigvEncFps`: The frame rate for sending videos.
	// `bigvDecFps`: The rendering frame rate.
	// `bigvBlock`: The video stutter duration (ms).
	// `aLoss`: The upstream/downstream audio packet loss.
	// `bigvLoss`: The upstream/downstream video packet loss.
	// `bigvWidth`: The upstream/downstream resolution (width).
	// `bigvHeight`: The upstream/downstream resolution (height).
	DataType []*string `json:"DataType,omitnil" name:"DataType"`

	// The page number. The default is 0.
	// Note: If `PageNumber` or `PageSize` is not specified, six records will be returned.
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of records per page. The default is `6`.
	// Value range: 1-100.
	// Note: If `DataType` is not null, the length of the array `UserIds` and the value of `PageSize` cannot exceed `6`.
	// If `DataType` is null, the length of the array `UserIds` and the value of `PageSize` cannot exceed `100`.
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeCallDetailInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallDetailInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "UserIds")
	delete(f, "DataType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallDetailInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallDetailInfoResponseParams struct {
	// The number of records returned.
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// The user information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserList []*UserInformation `json:"UserList,omitnil" name:"UserList"`

	// The call quality data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Data []*QualityData `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCallDetailInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallDetailInfoResponseParams `json:"Response"`
}

func (r *DescribeCallDetailInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallDetailInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRecordingRequestParams struct {
	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudRecordingResponseParams struct {
	// The unique ID of the recording task.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The status of the on-cloud recording task.
	// Idle: The task is idle.
	// InProgress: The task is in progress.
	// Exited: The task is being ended.
	Status *string `json:"Status,omitnil" name:"Status"`

	// The information of the recording files.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	StorageFileList []*StorageFile `json:"StorageFileList,omitnil" name:"StorageFileList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudRecordingResponseParams `json:"Response"`
}

func (r *DescribeCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMixTranscodingUsageRequestParams struct {
	// The start date in the format of YYYY-MM-DD.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date in the format of YYYY-MM-DD.
	// The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The `SDKAppID` of the TRTC application to which the target room belongs. If you do not specify this parameter, the usage statistics of all TRTC applications under the current account will be returned.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type DescribeMixTranscodingUsageRequest struct {
	*tchttp.BaseRequest
	
	// The start date in the format of YYYY-MM-DD.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date in the format of YYYY-MM-DD.
	// The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The `SDKAppID` of the TRTC application to which the target room belongs. If you do not specify this parameter, the usage statistics of all TRTC applications under the current account will be returned.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

func (r *DescribeMixTranscodingUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMixTranscodingUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMixTranscodingUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMixTranscodingUsageResponseParams struct {
	// The usage type. Each element of this parameter corresponds to an element of `UsageValue` in the order they are listed.
	UsageKey []*string `json:"UsageKey,omitnil" name:"UsageKey"`

	// The usage data in each time unit.
	UsageList []*TrtcUsage `json:"UsageList,omitnil" name:"UsageList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMixTranscodingUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMixTranscodingUsageResponseParams `json:"Response"`
}

func (r *DescribeMixTranscodingUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMixTranscodingUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingUsageRequestParams struct {
	// The start date in the format of YYYY-MM-DD.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date in the format of YYYY-MM-DD.
	// The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Whether to query single-stream or mixed-stream recording. Valid values: `single`, `multi`.
	MixType *string `json:"MixType,omitnil" name:"MixType"`

	// The `SDKAppID` of the TRTC application to which the target room belongs. If you do not specify this parameter, the usage statistics of all TRTC applications under the current account will be returned.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type DescribeRecordingUsageRequest struct {
	*tchttp.BaseRequest
	
	// The start date in the format of YYYY-MM-DD.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date in the format of YYYY-MM-DD.
	// The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Whether to query single-stream or mixed-stream recording. Valid values: `single`, `multi`.
	MixType *string `json:"MixType,omitnil" name:"MixType"`

	// The `SDKAppID` of the TRTC application to which the target room belongs. If you do not specify this parameter, the usage statistics of all TRTC applications under the current account will be returned.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

func (r *DescribeRecordingUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MixType")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRecordingUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRecordingUsageResponseParams struct {
	// The usage type. Each element of this parameter corresponds to an element of `UsageValue` in the order they are listed.
	UsageKey []*string `json:"UsageKey,omitnil" name:"UsageKey"`

	// The usage data in each time unit.
	UsageList []*TrtcUsage `json:"UsageList,omitnil" name:"UsageList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRecordingUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRecordingUsageResponseParams `json:"Response"`
}

func (r *DescribeRecordingUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRecordingUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelayUsageRequestParams struct {
	// The start date in the format of YYYY-MM-DD.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date in the format of YYYY-MM-DD.
	// The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The `SDKAppID` of the TRTC application to which the target room belongs. If you do not specify this parameter, the usage statistics of all TRTC applications under the current account will be returned.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type DescribeRelayUsageRequest struct {
	*tchttp.BaseRequest
	
	// The start date in the format of YYYY-MM-DD.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date in the format of YYYY-MM-DD.
	// The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The `SDKAppID` of the TRTC application to which the target room belongs. If you do not specify this parameter, the usage statistics of all TRTC applications under the current account will be returned.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

func (r *DescribeRelayUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelayUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRelayUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelayUsageResponseParams struct {
	// The usage type. Each element of this parameter corresponds to an element of `UsageValue` in the order they are listed.
	UsageKey []*string `json:"UsageKey,omitnil" name:"UsageKey"`

	// The usage data in each time unit.
	UsageList []*TrtcUsage `json:"UsageList,omitnil" name:"UsageList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRelayUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRelayUsageResponseParams `json:"Response"`
}

func (r *DescribeRelayUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelayUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomInfoRequestParams struct {
	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`.
	// Note: The end and start time cannot be more than 24 hours apart.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The room ID, such as `223`.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The page number. The default is 0.
	// Note: If `PageNumber` or `PageSize` is not specified, 10 records will be returned.
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of records per page. The default is `10`.
	// Value range: 1-100.
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeRoomInfoRequest struct {
	*tchttp.BaseRequest
	
	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`.
	// Note: The end and start time cannot be more than 24 hours apart.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The room ID, such as `223`.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The page number. The default is 0.
	// Note: If `PageNumber` or `PageSize` is not specified, 10 records will be returned.
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of records per page. The default is `10`.
	// Value range: 1-100.
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeRoomInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomInfoResponseParams struct {
	// The number of records returned.
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// The room information.
	RoomList []*RoomState `json:"RoomList,omitnil" name:"RoomList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRoomInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomInfoResponseParams `json:"Response"`
}

func (r *DescribeRoomInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScaleInfoRequestParams struct {
	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`. The end time and start time should preferably be more than 24 hours apart.
	// Note: Data is collected on a daily basis. To query the data of a day, make sure the end time is later than 00:00 on that day. Otherwise, no data will be returned. For example, to query the data on the 20th, the end time must be later than 00:00 on the 20th.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeScaleInfoRequest struct {
	*tchttp.BaseRequest
	
	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`. The end time and start time should preferably be more than 24 hours apart.
	// Note: Data is collected on a daily basis. To query the data of a day, make sure the end time is later than 00:00 on that day. Otherwise, no data will be returned. For example, to query the data on the 20th, the end time must be later than 00:00 on the 20th.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeScaleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScaleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScaleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScaleInfoResponseParams struct {
	// The number of records returned.
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// The returned data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScaleList []*ScaleInfomation `json:"ScaleList,omitnil" name:"ScaleList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScaleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScaleInfoResponseParams `json:"Response"`
}

func (r *DescribeScaleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScaleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamIngestRequestParams struct {
	// The SDKAppId of TRTC should be the same as the SDKAppId corresponding to the task room.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique Id of the task, will return after successfully starting the task.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeStreamIngestRequest struct {
	*tchttp.BaseRequest
	
	// The SDKAppId of TRTC should be the same as the SDKAppId corresponding to the task room.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique Id of the task, will return after successfully starting the task.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeStreamIngestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamIngestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamIngestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamIngestResponseParams struct {
	// Task status information. InProgress: Indicates that the current task is in progress. NotExist: Indicates that the current task does not exist. Example value: InProgress
	Status *string `json:"Status,omitnil" name:"Status"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeStreamIngestResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamIngestResponseParams `json:"Response"`
}

func (r *DescribeStreamIngestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamIngestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketQualityDataRequestParams struct {
	// User SDKAppId (e.g., 1400xxxxxx)
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Query start time, format is YYYY-MM-DD. (The query time range depends on the monitoring dashboard function version, the premium edition can query up to 30 days)
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query end time, format is YYYY-MM-DD.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The granularity of the returned data, which can be set to the following values:
	// d: by day. This returns data for the entire UTC day of the query time range. 
	// h: by hour. This returns data for the entire UTC hour of the query time range.
	Period *string `json:"Period,omitnil" name:"Period"`
}

type DescribeTRTCMarketQualityDataRequest struct {
	*tchttp.BaseRequest
	
	// User SDKAppId (e.g., 1400xxxxxx)
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Query start time, format is YYYY-MM-DD. (The query time range depends on the monitoring dashboard function version, the premium edition can query up to 30 days)
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query end time, format is YYYY-MM-DD.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The granularity of the returned data, which can be set to the following values:
	// d: by day. This returns data for the entire UTC day of the query time range. 
	// h: by hour. This returns data for the entire UTC hour of the query time range.
	Period *string `json:"Period,omitnil" name:"Period"`
}

func (r *DescribeTRTCMarketQualityDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketQualityDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCMarketQualityDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketQualityDataResponseParams struct {
	// TRTC Data Dashboard output parameters
	Data *TRTCDataResult `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTRTCMarketQualityDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCMarketQualityDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCMarketQualityDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketQualityDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketScaleDataRequestParams struct {
	// User SDKAppId
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Query start time, format is YYYY-MM-DD. (The query time range depends on the monitoring dashboard function version, the premium edition can query up to 60 days)
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query end time, format is YYYY-MM-DD.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The granularity of the returned data, which can be set to the following values:
	//  d: by day. This returns data for the entire UTC day of the query time range.
	//  h: by hour. This returns data for the entire UTC hour of the query time range.
	Period *string `json:"Period,omitnil" name:"Period"`
}

type DescribeTRTCMarketScaleDataRequest struct {
	*tchttp.BaseRequest
	
	// User SDKAppId
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Query start time, format is YYYY-MM-DD. (The query time range depends on the monitoring dashboard function version, the premium edition can query up to 60 days)
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// Query end time, format is YYYY-MM-DD.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The granularity of the returned data, which can be set to the following values:
	//  d: by day. This returns data for the entire UTC day of the query time range.
	//  h: by hour. This returns data for the entire UTC hour of the query time range.
	Period *string `json:"Period,omitnil" name:"Period"`
}

func (r *DescribeTRTCMarketScaleDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketScaleDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCMarketScaleDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCMarketScaleDataResponseParams struct {
	// TRTC Data Dashboard output parameters
	Data *TRTCDataResult `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTRTCMarketScaleDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCMarketScaleDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCMarketScaleDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCMarketScaleDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeQualityDataRequestParams struct {
	// User SDKAppId (e.g., 1400xxxxxx)
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Start time, unix timestamp, Unit: seconds (Query time range depends on the monitoring dashboard function version, standard edition can query the last 3 hours, premium edition can query the last 12 hours)
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// End time, unix timestamp, Unit: seconds
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// Room ID
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

type DescribeTRTCRealTimeQualityDataRequest struct {
	*tchttp.BaseRequest
	
	// User SDKAppId (e.g., 1400xxxxxx)
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Start time, unix timestamp, Unit: seconds (Query time range depends on the monitoring dashboard function version, standard edition can query the last 3 hours, premium edition can query the last 12 hours)
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// End time, unix timestamp, Unit: seconds
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// Room ID
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

func (r *DescribeTRTCRealTimeQualityDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeQualityDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCRealTimeQualityDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeQualityDataResponseParams struct {
	// TRTC Real- Time Monitoring output parameters
	Data *TRTCDataResult `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTRTCRealTimeQualityDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCRealTimeQualityDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCRealTimeQualityDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeQualityDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeScaleDataRequestParams struct {
	// User SDKAppId (e.g., 1400xxxxxx)
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Start time, unix timestamp, Unit: seconds (Query time range depends on the function version of the monitoring dashboard, premium edition can query up to 1 hours)
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// End time, unix timestamp, Unit: seconds
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// Room ID
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

type DescribeTRTCRealTimeScaleDataRequest struct {
	*tchttp.BaseRequest
	
	// User SDKAppId (e.g., 1400xxxxxx)
	SdkAppId *string `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Start time, unix timestamp, Unit: seconds (Query time range depends on the function version of the monitoring dashboard, premium edition can query up to 1 hours)
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// End time, unix timestamp, Unit: seconds
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// Room ID
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

func (r *DescribeTRTCRealTimeScaleDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeScaleDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTRTCRealTimeScaleDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTRTCRealTimeScaleDataResponseParams struct {
	// TRTC Real- Time Monitoring
	//  output parameter
	Data *TRTCDataResult `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTRTCRealTimeScaleDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTRTCRealTimeScaleDataResponseParams `json:"Response"`
}

func (r *DescribeTRTCRealTimeScaleDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTRTCRealTimeScaleDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcRoomUsageRequestParams struct {
	// The `SDKAppID` of the room.
	SdkAppid *uint64 `json:"SdkAppid,omitnil" name:"SdkAppid"`

	// The start time in the format of `YYYY-MM-DD HH:MM` (accurate to the minute).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end time in the format of `YYYY-MM-DD HH:MM`. The start and end time cannot be more than 24 hours apart.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeTrtcRoomUsageRequest struct {
	*tchttp.BaseRequest
	
	// The `SDKAppID` of the room.
	SdkAppid *uint64 `json:"SdkAppid,omitnil" name:"SdkAppid"`

	// The start time in the format of `YYYY-MM-DD HH:MM` (accurate to the minute).
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end time in the format of `YYYY-MM-DD HH:MM`. The start and end time cannot be more than 24 hours apart.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeTrtcRoomUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcRoomUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppid")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrtcRoomUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcRoomUsageResponseParams struct {
	// The usage data grouped by room, in CSV format.
	Data *string `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTrtcRoomUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrtcRoomUsageResponseParams `json:"Response"`
}

func (r *DescribeTrtcRoomUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcRoomUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcUsageRequestParams struct {
	// The start date in the format of YYYY-MM-DD.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date in the format of YYYY-MM-DD.
	// The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The `SDKAppID` of the TRTC application to which the target room belongs. If you do not specify this parameter, the usage statistics of all TRTC applications under the current account will be returned.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type DescribeTrtcUsageRequest struct {
	*tchttp.BaseRequest
	
	// The start date in the format of YYYY-MM-DD.
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// The end date in the format of YYYY-MM-DD.
	// The period queried per request cannot be longer than 31 days.
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// The `SDKAppID` of the TRTC application to which the target room belongs. If you do not specify this parameter, the usage statistics of all TRTC applications under the current account will be returned.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

func (r *DescribeTrtcUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrtcUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrtcUsageResponseParams struct {
	// The usage type. Each element of this parameter corresponds to an element of `UsageValue` in the order they are listed.
	UsageKey []*string `json:"UsageKey,omitnil" name:"UsageKey"`

	// The usage data in each time unit.
	UsageList []*TrtcUsage `json:"UsageList,omitnil" name:"UsageList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTrtcUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrtcUsageResponseParams `json:"Response"`
}

func (r *DescribeTrtcUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrtcUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnusualEventRequestParams struct {
	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`. The end time and start time cannot be more than one hour apart.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The room ID. Up to 20 random abnormal user experiences of the specified room will be returned.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

type DescribeUnusualEventRequest struct {
	*tchttp.BaseRequest
	
	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`. The end time and start time cannot be more than one hour apart.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The room ID. Up to 20 random abnormal user experiences of the specified room will be returned.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

func (r *DescribeUnusualEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnusualEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnusualEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnusualEventResponseParams struct {
	// The number of records returned.
	// Value range: 0-20.
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// The information of the abnormal user experiences.
	AbnormalExperienceList []*AbnormalExperience `json:"AbnormalExperienceList,omitnil" name:"AbnormalExperienceList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUnusualEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnusualEventResponseParams `json:"Response"`
}

func (r *DescribeUnusualEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnusualEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserEventRequestParams struct {
	// The unique ID of a call, whose format is `SdkAppId_CreateTime`, such as `1400xxxxxx_218695_1590065777`. `createTime` is the UNIX timestamp (seconds) when the room was created. Its value can be obtained using the [DescribeRoomInfo](https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1) API.
	CommId *string `json:"CommId,omitnil" name:"CommId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`.
	// Note: If you pass in an end time later than the room end time, the room end time will be used.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The room ID, such as `223`.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

type DescribeUserEventRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a call, whose format is `SdkAppId_CreateTime`, such as `1400xxxxxx_218695_1590065777`. `createTime` is the UNIX timestamp (seconds) when the room was created. Its value can be obtained using the [DescribeRoomInfo](https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1) API.
	CommId *string `json:"CommId,omitnil" name:"CommId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`.
	// Note: If you pass in an end time later than the room end time, the room end time will be used.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The room ID, such as `223`.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`
}

func (r *DescribeUserEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserId")
	delete(f, "RoomId")
	delete(f, "SdkAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserEventResponseParams struct {
	// The event list. An empty array will be returned if no data is obtained.
	Data []*EventList `json:"Data,omitnil" name:"Data"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserEventResponseParams `json:"Response"`
}

func (r *DescribeUserEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoRequestParams struct {
	// The unique ID of a call, whose format is `SdkAppId_CreateTime`, such as `1400xxxxxx_218695_1590065777`. `createTime` is the UNIX timestamp (seconds) when the room was created. Its value can be obtained using the [DescribeRoomInfo](https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1) API.
	CommId *string `json:"CommId,omitnil" name:"CommId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`.
	// Note: The end and start time cannot be more than four hours apart.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The users to query. If you do not specify this, the information of six users will be returned.
	// Array length: 1-100.
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// The page number. The default is 0.
	// Note: If `PageNumber` or `PageSize` is not specified, six records will be returned.
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of records per page. The default is `6`.
	// Array length: 1-100.
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// The unique ID of a call, whose format is `SdkAppId_CreateTime`, such as `1400xxxxxx_218695_1590065777`. `createTime` is the UNIX timestamp (seconds) when the room was created. Its value can be obtained using the [DescribeRoomInfo](https://intl.cloud.tencent.com/document/product/647/44050?from_cn_redirect=1) API.
	CommId *string `json:"CommId,omitnil" name:"CommId"`

	// The start time, which is a Unix timestamp (seconds) in local time, such as `1590065777`.
	// Note: Only data in the last 14 days can be queried.
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// The end time, which is a Unix timestamp (seconds) in local time, such as `1590065877`.
	// Note: The end and start time cannot be more than four hours apart.
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// The application ID, such as `1400xxxxxx`.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The users to query. If you do not specify this, the information of six users will be returned.
	// Array length: 1-100.
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`

	// The page number. The default is 0.
	// Note: If `PageNumber` or `PageSize` is not specified, six records will be returned.
	PageNumber *uint64 `json:"PageNumber,omitnil" name:"PageNumber"`

	// The number of records per page. The default is `6`.
	// Array length: 1-100.
	PageSize *uint64 `json:"PageSize,omitnil" name:"PageSize"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CommId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SdkAppId")
	delete(f, "UserIds")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoResponseParams struct {
	// The number of records returned.
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// The user information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserList []*UserInformation `json:"UserList,omitnil" name:"UserList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserInfoResponseParams `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomByStrRoomIdRequestParams struct {
	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Room ID
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

type DismissRoomByStrRoomIdRequest struct {
	*tchttp.BaseRequest
	
	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Room ID
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`
}

func (r *DismissRoomByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DismissRoomByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomByStrRoomIdResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DismissRoomByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *DismissRoomByStrRoomIdResponseParams `json:"Response"`
}

func (r *DismissRoomByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomRequestParams struct {
	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Room number.
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

type DismissRoomRequest struct {
	*tchttp.BaseRequest
	
	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Room number.
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`
}

func (r *DismissRoomRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DismissRoomRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismissRoomResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DismissRoomResponse struct {
	*tchttp.BaseResponse
	Response *DismissRoomResponseParams `json:"Response"`
}

func (r *DismissRoomResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismissRoomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventList struct {
	// The event information.
	Content []*EventMessage `json:"Content,omitnil" name:"Content"`

	// The user ID of the sender.
	PeerId *string `json:"PeerId,omitnil" name:"PeerId"`
}

type EventMessage struct {
	// The video stream type. Valid values:
	// `0`: A non-video event
	// `2`: The big video
	// `3`: The small video
	// `7`: A relayed video
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// The event reporting time in the format of UNIX timestamp (milliseconds), such as `1589891188801`.
	Time *uint64 `json:"Time,omitnil" name:"Time"`

	// The event ID. Events are classified into SDK events and WebRTC events. For more information, see https://www.tencentcloud.com/document/product/647/37906?has_map=1
	EventId *uint64 `json:"EventId,omitnil" name:"EventId"`

	// The first event parameter, such as the video width.
	ParamOne *int64 `json:"ParamOne,omitnil" name:"ParamOne"`

	// The second event parameter, such as the video height.
	ParamTwo *int64 `json:"ParamTwo,omitnil" name:"ParamTwo"`
}

type MaxVideoUser struct {
	// The stream information.
	UserMediaStream *UserMediaStream `json:"UserMediaStream,omitnil" name:"UserMediaStream"`
}

type McuAudioParams struct {
	// The audio encoding parameters.
	AudioEncode *AudioEncode `json:"AudioEncode,omitnil" name:"AudioEncode"`

	// The audio mix allowlist. For the `StartPublishCdnStream` API, if you do not pass this parameter or leave it empty, the audios of all anchors will be mixed. For the `UpdatePublishCdnStream` API, if you do not pass this parameter, no changes will be made to the current allowlist; if you pass in an empty string, the audios of all anchors will be mixed.
	// In cases where `SubscribeAudioList` and `UnSubscribeAudioList` are used at the same time, you need to specify both parameters. If you pass neither `SubscribeAudioList` nor `UnSubscribeAudioList`, no changes will be made. If a user is included in both parameters, the user’s audio will not be mixed.
	SubscribeAudioList []*McuUserInfoParams `json:"SubscribeAudioList,omitnil" name:"SubscribeAudioList"`

	// The audio mix blocklist. If you do not pass this parameter or leave it empty, there won’t be a blocklist. For the `UpdatePublishCdnStream` API, if you do not pass this parameter, no changes will be made to the current blocklist; if you pass in an empty string, the blocklist will be reset.
	// In cases where `SubscribeAudioList` and `UnSubscribeAudioList` are used at the same time, you need to specify both parameters. If you pass neither `SubscribeAudioList` nor `UnSubscribeAudioList`, no changes will be made. If a user is included in both parameters, the user’s audio will not be mixed.
	UnSubscribeAudioList []*McuUserInfoParams `json:"UnSubscribeAudioList,omitnil" name:"UnSubscribeAudioList"`
}

type McuCustomCrop struct {
	// The horizontal offset (pixels) of the starting point for cropping. This parameter must be greater than 0.
	LocationX *uint64 `json:"LocationX,omitnil" name:"LocationX"`

	// The vertical offset (pixels) of the starting point for cropping. This parameter must be greater than 0.
	LocationY *uint64 `json:"LocationY,omitnil" name:"LocationY"`

	// The video width (pixels) after cropping. The sum of this parameter and `LocationX` cannot be greater than 10000.
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// The video height (pixels) after cropping. The sum of this parameter and `LocationY` cannot be greater than 10000.
	Height *uint64 `json:"Height,omitnil" name:"Height"`
}

type McuFeedBackRoomParams struct {
	// The room ID.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The ID type of the room to which streams are relayed. `0` indicates integer, and `1` indicates string.
	RoomIdType *uint64 `json:"RoomIdType,omitnil" name:"RoomIdType"`

	// The [user ID](https://www.tencentcloud.com/document/product/647/37714) of the relaying robot in the TRTC room, which cannot be the same as a user ID already in use. We recommend you include the room ID in this user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The signature (similar to login password) required for the relaying robot to enter the room. For information on how to calculate the signature, see [What is UserSig?](https://www.tencentcloud.com/document/product/647/38104).
	UserSig *string `json:"UserSig,omitnil" name:"UserSig"`
}

type McuLayout struct {
	// The information of the stream that is displayed. If you do not pass this parameter, TRTC will display the videos of anchors in the room according to their room entry sequence.
	UserMediaStream *UserMediaStream `json:"UserMediaStream,omitnil" name:"UserMediaStream"`

	// The video width (pixels). If you do not pass this parameter, 0 will be used.
	ImageWidth *uint64 `json:"ImageWidth,omitnil" name:"ImageWidth"`

	// The video height (pixels). If you do not pass this parameter, 0 will be used.
	ImageHeight *uint64 `json:"ImageHeight,omitnil" name:"ImageHeight"`

	// The horizontal offset (pixels) of the video. The sum of `LocationX` and `ImageWidth` cannot exceed the width of the canvas. If you do not pass this parameter, 0 will be used.
	LocationX *uint64 `json:"LocationX,omitnil" name:"LocationX"`

	// The vertical offset of the video. The sum of `LocationY` and `ImageHeight` cannot exceed the height of the canvas. If you do not pass this parameter, 0 will be used.
	LocationY *uint64 `json:"LocationY,omitnil" name:"LocationY"`

	// The image layer of the video. If you do not pass this parameter, 0 will be used.
	ZOrder *uint64 `json:"ZOrder,omitnil" name:"ZOrder"`

	// The rendering mode of the video. 0 (the video is scaled and the excess parts are cropped), 1 (the video is scaled), 2 (the video is scaled and the blank spaces are filled with black bars). If you do not pass this parameter, 0 will be used.
	RenderMode *uint64 `json:"RenderMode,omitnil" name:"RenderMode"`

	// (Not supported yet) The background color of a video. Below are the values for some commonly used colors:
	// Red: `0xcc0033`
	// Yellow: `0xcc9900`
	// Green: `0xcccc33`
	// Blue: `0x99CCFF`
	// Black: `0x000000`
	// White: `0xFFFFFF`
	// Grey: `0x999999`
	BackGroundColor *string `json:"BackGroundColor,omitnil" name:"BackGroundColor"`

	// The URL of the background image for the video. This parameter allows you to specify an image to display when the user’s camera is turned off or before the user enters the room. If the dimensions of the image specified are different from those of the video window, the image will be stretched to fit the space. This parameter has a higher priority than `BackGroundColor`.
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil" name:"BackgroundImageUrl"`

	// Custom cropping.
	CustomCrop *McuCustomCrop `json:"CustomCrop,omitnil" name:"CustomCrop"`

	// The display mode of the sub-background image during output: 0 for cropping, 1 for scaling and displaying the background, 2 for scaling and displaying the black background, 3 for proportional scaling. If not filled in, the default is 3.
	BackgroundRenderMode *uint64 `json:"BackgroundRenderMode,omitnil" name:"BackgroundRenderMode"`
}

type McuLayoutParams struct {
	// The layout mode. Valid values: 1 (floating), 2 (screen sharing), 3 (grid), 4 (custom). Floating, screen sharing, and grid are dynamic layouts. Custom layouts are static layouts.
	MixLayoutMode *uint64 `json:"MixLayoutMode,omitnil" name:"MixLayoutMode"`

	// Whether to display users who publish only audio. 0: No; 1: Yes. This parameter is valid only if a dynamic layout is used. If you do not pass this parameter, 0 will be used.
	PureAudioHoldPlaceMode *uint64 `json:"PureAudioHoldPlaceMode,omitnil" name:"PureAudioHoldPlaceMode"`

	// The details of a custom layout.
	MixLayoutList []*McuLayout `json:"MixLayoutList,omitnil" name:"MixLayoutList"`

	// The information of the large video in screen sharing or floating layout mode.
	MaxVideoUser *MaxVideoUser `json:"MaxVideoUser,omitnil" name:"MaxVideoUser"`

	// The image fill mode. This parameter is valid if the layout mode is screen sharing, floating, or grid. `0`: The image will be cropped. `1`: The image will be scaled. `2`: The image will be scaled and there may be black bars.
	RenderMode *uint64 `json:"RenderMode,omitnil" name:"RenderMode"`
}

type McuLayoutVolume struct {
	// The application data, which will be embedded in the `app_data` field of the custom SEI. It must be shorter than 4,096 characters.
	AppData *string `json:"AppData,omitnil" name:"AppData"`

	// The payload type of the SEI message. The default is 100. Value range: 100-254 (244 is used internally by Tencent Cloud for timestamps).
	PayloadType *uint64 `json:"PayloadType,omitnil" name:"PayloadType"`

	// The SEI sending interval (milliseconds). The default value is 1000.
	Interval *uint64 `json:"Interval,omitnil" name:"Interval"`

	// Valid values: `1`: SEI is guaranteed when keyframes are sent; `0` (default): SEI is not guaranteed when keyframes are sent.
	FollowIdr *uint64 `json:"FollowIdr,omitnil" name:"FollowIdr"`
}

type McuPassThrough struct {
	// The payload of the pass-through SEI.
	PayloadContent *string `json:"PayloadContent,omitnil" name:"PayloadContent"`

	// The payload type of the SEI message. Value range: 5 and 100-254 (244 is used internally by Tencent Cloud for timestamps).
	PayloadType *uint64 `json:"PayloadType,omitnil" name:"PayloadType"`

	// This parameter is required only if `PayloadType` is 5. It must be a 32-character hexadecimal string. If `PayloadType` is not 5, this parameter will be ignored.
	PayloadUuid *string `json:"PayloadUuid,omitnil" name:"PayloadUuid"`

	// The SEI sending interval (milliseconds). The default value is 1000.
	Interval *uint64 `json:"Interval,omitnil" name:"Interval"`

	// Valid values: `1`: SEI is guaranteed when keyframes are sent; `0` (default): SEI is not guaranteed when keyframes are sent.
	FollowIdr *uint64 `json:"FollowIdr,omitnil" name:"FollowIdr"`
}

type McuPublishCdnParam struct {
	// The URLs of the CDNs to relay to.
	PublishCdnUrl *string `json:"PublishCdnUrl,omitnil" name:"PublishCdnUrl"`

	// Whether to relay to Tencent Cloud’s CDN. `0`: Third-party CDN; `1` (default): Tencent Cloud’s CDN. Relaying to a third-party CDN will incur fees. To avoid unexpected charges, we recommend you pass in a specific value. For details, see the API document.
	IsTencentCdn *uint64 `json:"IsTencentCdn,omitnil" name:"IsTencentCdn"`
}

type McuSeiParams struct {
	// The audio volume layout SEI.
	LayoutVolume *McuLayoutVolume `json:"LayoutVolume,omitnil" name:"LayoutVolume"`

	// The pass-through SEI.
	PassThrough *McuPassThrough `json:"PassThrough,omitnil" name:"PassThrough"`
}

type McuUserInfoParams struct {
	// The user information.
	UserInfo *MixUserInfo `json:"UserInfo,omitnil" name:"UserInfo"`
}

type McuVideoParams struct {
	// The video encoding parameters.
	VideoEncode *VideoEncode `json:"VideoEncode,omitnil" name:"VideoEncode"`

	// The layout parameters.
	LayoutParams *McuLayoutParams `json:"LayoutParams,omitnil" name:"LayoutParams"`

	// The canvas color. Below are the values for some common colors:
	// Red: 0xcc0033
	// Yellow: 0xcc9900
	// Green: 0xcccc33
	// Blue: 0x99CCFF
	// Black: 0x000000
	// White: 0xFFFFFF
	// Grey: 0x999999
	BackGroundColor *string `json:"BackGroundColor,omitnil" name:"BackGroundColor"`

	// The URL of the background image for the canvas. This parameter has a higher priority than `BackGroundColor`.
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil" name:"BackgroundImageUrl"`

	// The watermark information for the mixed stream.
	WaterMarkList []*McuWaterMarkParams `json:"WaterMarkList,omitnil" name:"WaterMarkList"`

	// Background image display mode during output: 0 for crop, 1 for scale and display with black background, 2 for proportional scaling. The backend default is proportional scaling.
	BackgroundRenderMode *uint64 `json:"BackgroundRenderMode,omitnil" name:"BackgroundRenderMode"`
}

type McuWaterMarkImage struct {
	// The URL of the watermark image, which must be in PNG, JPG, or JPEG format and cannot exceed 5 MB.
	WaterMarkUrl *string `json:"WaterMarkUrl,omitnil" name:"WaterMarkUrl"`

	// The watermark width (pixels).
	WaterMarkWidth *uint64 `json:"WaterMarkWidth,omitnil" name:"WaterMarkWidth"`

	// The watermark height (pixels).
	WaterMarkHeight *uint64 `json:"WaterMarkHeight,omitnil" name:"WaterMarkHeight"`

	// The horizontal offset (pixels) of the watermark.
	LocationX *uint64 `json:"LocationX,omitnil" name:"LocationX"`

	// The vertical offset (pixels) of the watermark.
	LocationY *uint64 `json:"LocationY,omitnil" name:"LocationY"`

	// The image layer of the watermark. If you do not pass this parameter, 0 will be used.
	ZOrder *uint64 `json:"ZOrder,omitnil" name:"ZOrder"`
}

type McuWaterMarkParams struct {
	// The watermark type. Valid values: `0` (default): Image; `1`: Text.
	WaterMarkType *uint64 `json:"WaterMarkType,omitnil" name:"WaterMarkType"`

	// The watermark image information. This parameter is required if `WaterMarkType` is 0.
	WaterMarkImage *McuWaterMarkImage `json:"WaterMarkImage,omitnil" name:"WaterMarkImage"`

	// The text watermark configuration. This parameter is required if `WaterMarkType` is `1`.
	WaterMarkText *McuWaterMarkText `json:"WaterMarkText,omitnil" name:"WaterMarkText"`
}

type McuWaterMarkText struct {
	// The text.
	Text *string `json:"Text,omitnil" name:"Text"`

	// The watermark width (pixels).
	WaterMarkWidth *uint64 `json:"WaterMarkWidth,omitnil" name:"WaterMarkWidth"`

	// The watermark height (pixels).
	WaterMarkHeight *uint64 `json:"WaterMarkHeight,omitnil" name:"WaterMarkHeight"`

	// The horizontal offset (pixels) of the watermark.
	LocationX *uint64 `json:"LocationX,omitnil" name:"LocationX"`

	// The vertical offset (pixels) of the watermark.
	LocationY *uint64 `json:"LocationY,omitnil" name:"LocationY"`

	// The font size.
	FontSize *uint64 `json:"FontSize,omitnil" name:"FontSize"`

	// The text color. The default color is white. Values for some commonly used colors: Red: `0xcc0033`; yellow: `0xcc9900`; green: `0xcccc33`; blue: `0x99CCFF`; black: `0x000000`; white: `0xFFFFFF`; gray: `0x999999`.	
	FontColor *string `json:"FontColor,omitnil" name:"FontColor"`

	// The text fill color. If you do not specify this parameter, the fill color will be transparent. Values for some commonly used colors: Red: `0xcc0033`; yellow: `0xcc9900`; green: `0xcccc33`; blue: `0x99CCFF`; black: `0x000000`; white: `0xFFFFFF`; gray: `0x999999`.	
	BackGroundColor *string `json:"BackGroundColor,omitnil" name:"BackGroundColor"`
}

type MixLayout struct {
	// The Y axis of the window’s top-left corner. Value range: [0, 1920]. The value cannot be larger than the canvas height.
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// The X axis of the window’s top-left corner. Value range: [0, 1920]. The value cannot be larger than the canvas width.
	Left *uint64 `json:"Left,omitnil" name:"Left"`

	// The relative width of the window. Value range: [0, 1920]. The sum of the values of this parameter and `Left` cannot exceed the canvas width.
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// The relative height of the window. Value range: [0, 1920]. The sum of the values of this parameter and `Top` cannot exceed the canvas height.
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// The user ID (string) of the anchor whose video is shown in the window. If you do not set this parameter, anchors’ videos will be shown in their room entry sequence.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The degree of transparency of the canvas. Value range: [0, 255]. 0 means fully opaque, and 255 means fully transparent.
	Alpha *uint64 `json:"Alpha,omitnil" name:"Alpha"`

	// 0: Stretch. In this mode, the image is stretched to fill the space available. The whole image is visible after scaling. However, if the original aspect ratio is different from the target, the image may be distorted.
	// 
	// 1: Crop (default). In this mode, if the original aspect ratio is different from the target, the image will be cropped according to the target before being stretched to fill the space available. The image will not be distorted.
	// 
	// 2: Blank. This mode stretches the image while keeping its original aspect ratio. If the original aspect ratio is different from the target, there may be blank spaces to the top and bottom or to the left and right of the window.
	// 
	// 3: Smart stretch. This mode is similar to the crop mode, except that it restricts cropping to 20% of the image’s width or height at most.
	RenderMode *uint64 `json:"RenderMode,omitnil" name:"RenderMode"`

	// The type of the stream subscribed to.
	// 0: Primary stream (default)
	// 1: Substream
	MediaId *uint64 `json:"MediaId,omitnil" name:"MediaId"`

	// The image layer. 0 is the default value and means the bottommost layer.
	ImageLayer *uint64 `json:"ImageLayer,omitnil" name:"ImageLayer"`

	// The URL of the background image for a window. The image must be in JPG or PNG format and cannot be larger than 5 MB. If the image’s aspect ratio is different from that of the window, the image will be rendered according to the value of `RenderMode`.
	SubBackgroundImage *string `json:"SubBackgroundImage,omitnil" name:"SubBackgroundImage"`
}

type MixLayoutParams struct {
	// Layout mode:
	// 1: Floating
	// 2: Screen sharing
	// 3: Grid (default)
	// 4: Custom
	// 
	// Floating: By default, the video of the first anchor (you can also specify an anchor) who enters the room is scaled to fill the screen. When other anchors enter the room, their videos appear smaller and are superimposed over the large video from left to right starting from the bottom of the canvas according to their room entry sequence. If the total number of videos is 17 or less, there will be four windows in each row (4 x 4); if it is greater than 17, there will be five windows in each row (5 x 5). Up to 25 videos can be displayed. A user who publishes only audio will still be displayed in one window.
	// 
	// Screen sharing: The video of a specified anchor occupies a larger part of the canvas on the left side (if you do not specify an anchor, the left window will display the canvas background). The videos of other anchors are smaller and are positioned on the right side. If the total number of videos is 17 or less, the small videos are positioned from top to bottom in up to two columns on the right side, with eight videos per column at most. If there are more than 17 videos, the additional videos are positioned at the bottom of the canvas from left to right. Up to 25 videos can be displayed. A user who publishes only audio will still be displayed in one window.
	// 
	// Grid: The videos of anchors are scaled and positioned automatically according to the total number of anchors in a room. Each video has the same size. Up to 25 videos can be displayed.
	// 
	// Custom: Specify the layout of videos by using the `MixLayoutList` parameter.
	MixLayoutMode *uint64 `json:"MixLayoutMode,omitnil" name:"MixLayoutMode"`

	// The custom layout details. This parameter is valid if `MixLayoutMode` is set to `4`. Up to 25 videos can be displayed.
	MixLayoutList []*MixLayout `json:"MixLayoutList,omitnil" name:"MixLayoutList"`

	// The background color, which is a hexadecimal value (starting with "#", followed by the color value) converted from an 8-bit RGB value. For example, the RGB value of orange is `R:255 G:165 B:0`, and its hexadecimal value is `#FFA500`. The default color is black.
	BackGroundColor *string `json:"BackGroundColor,omitnil" name:"BackGroundColor"`

	// The user whose video is displayed in the big window. This parameter is valid if `MixLayoutMode` is set to `1` (floating) or `2` (screen sharing). If it is left empty, the first anchor entering the room is displayed in the big window in the floating mode and the canvas background is displayed in the screen sharing mode.
	MaxResolutionUserId *string `json:"MaxResolutionUserId,omitnil" name:"MaxResolutionUserId"`

	// The stream type.
	// 0: Primary stream (default)
	// 1: Substream (screen sharing stream)
	// This parameter specifies the type of the stream displayed in the big window. If it appears in `MixLayoutList`, it indicates the type of the stream of a specified user.
	MediaId *uint64 `json:"MediaId,omitnil" name:"MediaId"`

	// The URL of the background image, which cannot contain Chinese characters. The image must be in JPG or PNG format and cannot be larger than 5 MB.
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil" name:"BackgroundImageUrl"`

	// `1` means to use placeholders, and `0` (default) means to not use placeholders. If this parameter is set to `1`, when a user is not publishing video, a placeholder image will be displayed in the window reserved for the user.
	PlaceHolderMode *uint64 `json:"PlaceHolderMode,omitnil" name:"PlaceHolderMode"`

	// The render mode to use when the aspect ratio of a video is different from that of the window. This parameter is defined the same as `RenderMode` in `MixLayoufList`.
	BackgroundImageRenderMode *uint64 `json:"BackgroundImageRenderMode,omitnil" name:"BackgroundImageRenderMode"`

	// The URL of the background image for a window. The image must be in JPG or PNG format and cannot be larger than 5 MB. If the image’s aspect ratio is different from that of the window, the image will be rendered according to the value of `RenderMode`.
	DefaultSubBackgroundImage *string `json:"DefaultSubBackgroundImage,omitnil" name:"DefaultSubBackgroundImage"`

	// The watermark layout. Up to 25 watermarks are supported.
	WaterMarkList []*WaterMark `json:"WaterMarkList,omitnil" name:"WaterMarkList"`

	// The render mode to use when the aspect ratio of a video is different from that of the window. This parameter is invalid if a custom layout is used. It is defined the same as `RenderMode` in `MixLayoufList`.
	RenderMode *uint64 `json:"RenderMode,omitnil" name:"RenderMode"`

	// This parameter is valid only if the screen sharing layout is used. If you set it to `1`, the large video window will appear on the right and the small window on the left. The default value is `0`.
	MaxResolutionUserAlign *uint64 `json:"MaxResolutionUserAlign,omitnil" name:"MaxResolutionUserAlign"`
}

type MixTranscodeParams struct {
	// The video transcoding parameters for recording. If you set this parameter, you must specify all its fields. If you do not set it, the default will be used.
	VideoParams *VideoParams `json:"VideoParams,omitnil" name:"VideoParams"`

	// The audio transcoding parameters for recording. If you set this parameter, you must specify all its fields. If you do not set it, the default will be used.
	AudioParams *AudioParams `json:"AudioParams,omitnil" name:"AudioParams"`
}

type MixUserInfo struct {
	// User ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// If a dynamic layout is used, the value of this parameter should be the ID of the main room. If a custom layout is used, the value of this parameter should be the same as the room ID in `MixLayoutList`.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The type of the `RoomId` parameter. 0: integer; 1: string.
	RoomIdType *uint64 `json:"RoomIdType,omitnil" name:"RoomIdType"`
}

// Predefined struct for user
type ModifyCloudRecordingRequestParams struct {
	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The new stream mixing layout to use.
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitnil" name:"MixLayoutParams"`

	// The allowlist/blocklist for stream subscription.
	SubscribeStreamUserIds *SubscribeStreamUserIds `json:"SubscribeStreamUserIds,omitnil" name:"SubscribeStreamUserIds"`
}

type ModifyCloudRecordingRequest struct {
	*tchttp.BaseRequest
	
	// The `SDKAppID` of the room whose streams are recorded.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique ID of the recording task, which is returned after recording starts successfully.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The new stream mixing layout to use.
	MixLayoutParams *MixLayoutParams `json:"MixLayoutParams,omitnil" name:"MixLayoutParams"`

	// The allowlist/blocklist for stream subscription.
	SubscribeStreamUserIds *SubscribeStreamUserIds `json:"SubscribeStreamUserIds,omitnil" name:"SubscribeStreamUserIds"`
}

func (r *ModifyCloudRecordingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudRecordingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	delete(f, "MixLayoutParams")
	delete(f, "SubscribeStreamUserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudRecordingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudRecordingResponseParams struct {
	// The task ID assigned by the recording service, which uniquely identifies a recording process and becomes invalid after a recording task ends.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCloudRecordingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudRecordingResponseParams `json:"Response"`
}

func (r *ModifyCloudRecordingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudRecordingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QualityData struct {
	// The quality data.
	Content []*TimeValue `json:"Content,omitnil" name:"Content"`

	// The user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The remote user ID. An empty string indicates that the data is upstream data.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PeerId *string `json:"PeerId,omitnil" name:"PeerId"`

	// The data type.
	DataType *string `json:"DataType,omitnil" name:"DataType"`
}

type RecordParams struct {
	// The recording mode.
	// 1: Single-stream recording. Records the audio and video of each subscribed user (`UserId`) in a room and saves the recording files to the cloud.
	// 2: Mixed-stream recording. Mixes the audios and videos of subscribed users (`UserId`) in a room, records the mixed stream, and saves the recording files to the cloud.
	RecordMode *uint64 `json:"RecordMode,omitnil" name:"RecordMode"`

	// The time period (seconds) to wait to automatically stop recording after there are no anchors (users who publish streams) in a room. Value range: 5-86400 (max 24 hours). Default value: 30.
	MaxIdleTime *uint64 `json:"MaxIdleTime,omitnil" name:"MaxIdleTime"`

	// The media type of the streams to record.
	// 0: Audio and video streams (default)
	// 1: Audio streams only
	// 2: Video streams only
	StreamType *uint64 `json:"StreamType,omitnil" name:"StreamType"`

	// The allowlist/blocklist for stream subscription.
	SubscribeStreamUserIds *SubscribeStreamUserIds `json:"SubscribeStreamUserIds,omitnil" name:"SubscribeStreamUserIds"`

	// The output format. `0` (default): HLS; `1`: HLS + MP4; `2`: HLS + AAC;  `3` : MP4,  `4` : AAC. This parameter is invalid if you save recording files to VOD. To specify the format of files saved to VOD, use `MediaType` of `TencentVod`.
	OutputFormat *uint64 `json:"OutputFormat,omitnil" name:"OutputFormat"`

	// Whether to merge the audio and video of a user in the single-stream recording mode. 0 (default): Do not mix the audio and video; 1: Mix the audio and video into one TS file. You don’t need to specify this parameter for mixed-stream recording, which merges audios and videos by default.
	AvMerge *uint64 `json:"AvMerge,omitnil" name:"AvMerge"`

	// The maximum file duration allowed (minutes). If the output format is AAC or MP4, and the maximum file duration is exceeded, the file will be segmented. Value range: 1-1440. Default value: 1440 (24 hours). The maximum file size allowed is 2 GB. If the file size exceeds 2 GB, or the file duration exceeds 24 hours, the file will also be segmented.
	// This parameter is invalid if the output format is HLS.
	MaxMediaFileDuration *uint64 `json:"MaxMediaFileDuration,omitnil" name:"MaxMediaFileDuration"`

	// The type of stream to record. `0` (default): The primary stream and substream; `1`: The primary stream; `2`: The substream.
	MediaId *uint64 `json:"MediaId,omitnil" name:"MediaId"`
}

// Predefined struct for user
type RemoveUserByStrRoomIdRequestParams struct {
	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Room ID
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// List of up to 10 users to be removed
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`
}

type RemoveUserByStrRoomIdRequest struct {
	*tchttp.BaseRequest
	
	// `SDKAppId` of TRTC
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Room ID
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// List of up to 10 users to be removed
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`
}

func (r *RemoveUserByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserByStrRoomIdResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RemoveUserByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserByStrRoomIdResponseParams `json:"Response"`
}

func (r *RemoveUserByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserRequestParams struct {
	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Room number.
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// List of up to 10 users to be removed.
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`
}

type RemoveUserRequest struct {
	*tchttp.BaseRequest
	
	// `SDKAppId` of TRTC.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// Room number.
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// List of up to 10 users to be removed.
	UserIds []*string `json:"UserIds,omitnil" name:"UserIds"`
}

func (r *RemoveUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RemoveUserResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserResponseParams `json:"Response"`
}

func (r *RemoveUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoomState struct {
	// The call ID, which uniquely identifies a call.
	CommId *string `json:"CommId,omitnil" name:"CommId"`

	// The room ID.
	RoomString *string `json:"RoomString,omitnil" name:"RoomString"`

	// The room creation time.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// The room termination time.
	DestroyTime *uint64 `json:"DestroyTime,omitnil" name:"DestroyTime"`

	// Whether the room is terminated.
	IsFinished *bool `json:"IsFinished,omitnil" name:"IsFinished"`

	// The user ID of the room creator.
	UserId *string `json:"UserId,omitnil" name:"UserId"`
}

type RowValues struct {
	// Data value
	RowValue []*int64 `json:"RowValue,omitnil" name:"RowValue"`
}

type ScaleInfomation struct {
	// Start time for each day
	Time *uint64 `json:"Time,omitnil" name:"Time"`

	// The number of users. If a user enters a room multiple times, it will be counted as one user.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserNumber *uint64 `json:"UserNumber,omitnil" name:"UserNumber"`

	// The number of room entries. Every time a user enters a room, it will be counted as one room entry.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UserCount *uint64 `json:"UserCount,omitnil" name:"UserCount"`

	// The total number of rooms of the application on a day.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RoomNumbers *uint64 `json:"RoomNumbers,omitnil" name:"RoomNumbers"`
}

type SeriesInfos struct {
	// Data columns
	Columns []*string `json:"Columns,omitnil" name:"Columns"`

	// Data values
	Values []*RowValues `json:"Values,omitnil" name:"Values"`
}

// Predefined struct for user
type SetUserBlockedByStrRoomIdRequestParams struct {
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The room ID (string).
	StrRoomId *string `json:"StrRoomId,omitnil" name:"StrRoomId"`

	// The user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// Whether to disable the user’s audio and video. 0: Enable; 1: Disable.
	IsMute *uint64 `json:"IsMute,omitnil" name:"IsMute"`
}

type SetUserBlockedByStrRoomIdRequest struct {
	*tchttp.BaseRequest
	
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The room ID (string).
	StrRoomId *string `json:"StrRoomId,omitnil" name:"StrRoomId"`

	// The user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// Whether to disable the user’s audio and video. 0: Enable; 1: Disable.
	IsMute *uint64 `json:"IsMute,omitnil" name:"IsMute"`
}

func (r *SetUserBlockedByStrRoomIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserBlockedByStrRoomIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "StrRoomId")
	delete(f, "UserId")
	delete(f, "IsMute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetUserBlockedByStrRoomIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetUserBlockedByStrRoomIdResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetUserBlockedByStrRoomIdResponse struct {
	*tchttp.BaseResponse
	Response *SetUserBlockedByStrRoomIdResponseParams `json:"Response"`
}

func (r *SetUserBlockedByStrRoomIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserBlockedByStrRoomIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetUserBlockedRequestParams struct {
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The room ID (number).
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// The user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// Whether to disable the user’s audio and video. 0: Enable; 1: Disable.
	IsMute *uint64 `json:"IsMute,omitnil" name:"IsMute"`
}

type SetUserBlockedRequest struct {
	*tchttp.BaseRequest
	
	// The application ID.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The room ID (number).
	RoomId *uint64 `json:"RoomId,omitnil" name:"RoomId"`

	// The user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// Whether to disable the user’s audio and video. 0: Enable; 1: Disable.
	IsMute *uint64 `json:"IsMute,omitnil" name:"IsMute"`
}

func (r *SetUserBlockedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserBlockedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "UserId")
	delete(f, "IsMute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetUserBlockedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetUserBlockedResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetUserBlockedResponse struct {
	*tchttp.BaseResponse
	Response *SetUserBlockedResponseParams `json:"Response"`
}

func (r *SetUserBlockedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserBlockedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SingleSubscribeParams struct {
	// The stream information.
	UserMediaStream *UserMediaStream `json:"UserMediaStream,omitnil" name:"UserMediaStream"`
}

// Predefined struct for user
type StartPublishCdnStreamRequestParams struct {
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The ID of the room whose streams are relayed (the main room).
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The type of the `RoomId` parameter, which must be the same as the ID type of the room whose streams are relayed. 0: integer; 1: string.
	RoomIdType *uint64 `json:"RoomIdType,omitnil" name:"RoomIdType"`

	// The information of the relaying robot in the room.
	AgentParams *AgentParams `json:"AgentParams,omitnil" name:"AgentParams"`

	// Whether to transcode the streams. `0`: No. `1`: Yes. This parameter determines whether transcoding fees are charged. If it is `0`, streams will only be relayed, and no transcoding fees will be incurred. If it is `1`, streams will be transcoded before being relayed, and transcoding fees will be incurred.
	WithTranscoding *uint64 `json:"WithTranscoding,omitnil" name:"WithTranscoding"`

	// The audio encoding parameters. Because audio is always transcoded (no fees are incurred), this parameter is required when you start a relay task.
	AudioParams *McuAudioParams `json:"AudioParams,omitnil" name:"AudioParams"`

	// The video encoding parameters for relaying. If you do not pass this parameter, only audio will be relayed.
	VideoParams *McuVideoParams `json:"VideoParams,omitnil" name:"VideoParams"`

	// The information of a single stream relayed. When you relay a single stream, set `WithTranscoding` to 0.
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitnil" name:"SingleSubscribeParams"`

	// The information of the CDNs to relay to. You need to specify at least one between this parameter and `FeedBackRoomParams.N`.
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil" name:"PublishCdnParams"`

	// The stream mixing SEI parameters.
	SeiParams *McuSeiParams `json:"SeiParams,omitnil" name:"SeiParams"`

	// The information of the room to which streams are relayed. Between this parameter and `PublishCdnParams`, you must specify at least one. Please note that relaying to a TRTC room is only supported in some SDK versions. For details, please contact technical support.
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitnil" name:"FeedBackRoomParams"`
}

type StartPublishCdnStreamRequest struct {
	*tchttp.BaseRequest
	
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The ID of the room whose streams are relayed (the main room).
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// The type of the `RoomId` parameter, which must be the same as the ID type of the room whose streams are relayed. 0: integer; 1: string.
	RoomIdType *uint64 `json:"RoomIdType,omitnil" name:"RoomIdType"`

	// The information of the relaying robot in the room.
	AgentParams *AgentParams `json:"AgentParams,omitnil" name:"AgentParams"`

	// Whether to transcode the streams. `0`: No. `1`: Yes. This parameter determines whether transcoding fees are charged. If it is `0`, streams will only be relayed, and no transcoding fees will be incurred. If it is `1`, streams will be transcoded before being relayed, and transcoding fees will be incurred.
	WithTranscoding *uint64 `json:"WithTranscoding,omitnil" name:"WithTranscoding"`

	// The audio encoding parameters. Because audio is always transcoded (no fees are incurred), this parameter is required when you start a relay task.
	AudioParams *McuAudioParams `json:"AudioParams,omitnil" name:"AudioParams"`

	// The video encoding parameters for relaying. If you do not pass this parameter, only audio will be relayed.
	VideoParams *McuVideoParams `json:"VideoParams,omitnil" name:"VideoParams"`

	// The information of a single stream relayed. When you relay a single stream, set `WithTranscoding` to 0.
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitnil" name:"SingleSubscribeParams"`

	// The information of the CDNs to relay to. You need to specify at least one between this parameter and `FeedBackRoomParams.N`.
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil" name:"PublishCdnParams"`

	// The stream mixing SEI parameters.
	SeiParams *McuSeiParams `json:"SeiParams,omitnil" name:"SeiParams"`

	// The information of the room to which streams are relayed. Between this parameter and `PublishCdnParams`, you must specify at least one. Please note that relaying to a TRTC room is only supported in some SDK versions. For details, please contact technical support.
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitnil" name:"FeedBackRoomParams"`
}

func (r *StartPublishCdnStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishCdnStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "RoomIdType")
	delete(f, "AgentParams")
	delete(f, "WithTranscoding")
	delete(f, "AudioParams")
	delete(f, "VideoParams")
	delete(f, "SingleSubscribeParams")
	delete(f, "PublishCdnParams")
	delete(f, "SeiParams")
	delete(f, "FeedBackRoomParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishCdnStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishCdnStreamResponseParams struct {
	// The task ID, which is generated by the Tencent Cloud server. You need to pass in the task ID when making a request to update or stop a relaying task.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartPublishCdnStreamResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishCdnStreamResponseParams `json:"Response"`
}

func (r *StartPublishCdnStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishCdnStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStreamIngestRequestParams struct {
	// TRTC's [SdkAppId](https://intl.cloud.tencent.com/document/product/647/46351?from_cn_redirect=1#sdkappid), the same as the SdkAppId corresponding to the Record room.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// TRTC's [RoomId](https://intl.cloud.tencent.com/document/product/647/46351?from_cn_redirect=1#roomid), the RoomId corresponding to the Record TRTC room.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// Type of TRTC RoomId. 【*Note】Must be the same as the RoomId type corresponding to the Record room: 0: String type RoomId 1: 32-bit Integer type RoomId (default)
	RoomIdType *uint64 `json:"RoomIdType,omitnil" name:"RoomIdType"`

	// UserId of the Pull stream Relay Robot, used to enter the room and initiate the Pull stream Relay Task.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// UserSig corresponding to the Pull stream Relay Robot UserId, i.e., UserId and UserSig are equivalent to the Robot's Login password for entering the room. For the specific Calculation method, please refer to the TRTC [UserSig](https://intl.cloud.tencent.com/document/product/647/45910?from_cn_redirect=1#UserSig) Scheme.
	UserSig *string `json:"UserSig,omitnil" name:"UserSig"`

	// 	
	// Source URL. Example value: https://a.b/test.mp4
	SourceUrl []*string `json:"SourceUrl,omitnil" name:"SourceUrl"`

	// TRTC room permission Encryption ticket, only needed when advanced permission control is enabled in the Console. After enabling advanced permission control in the TRTC Console, TRTC's backend service system will verify a so-called [PrivateMapKey] 'Permission ticket', which contains an encrypted RoomId and an encrypted 'Permission bit list'. Since PrivateMapKey contains RoomId, providing only UserSig without PrivateMapKey does not allow entry into the specified room.
	PrivateMapKey *string `json:"PrivateMapKey,omitnil" name:"PrivateMapKey"`

	// Video Codec Parameters. Optional, if not filled, Keep original stream Parameters.
	VideoEncodeParams *VideoEncodeParams `json:"VideoEncodeParams,omitnil" name:"VideoEncodeParams"`

	// Audio Codec Parameters. Optional, if not filled, Keep original stream Parameters.
	AudioEncodeParams *AudioEncodeParams `json:"AudioEncodeParams,omitnil" name:"AudioEncodeParams"`
}

type StartStreamIngestRequest struct {
	*tchttp.BaseRequest
	
	// TRTC's [SdkAppId](https://intl.cloud.tencent.com/document/product/647/46351?from_cn_redirect=1#sdkappid), the same as the SdkAppId corresponding to the Record room.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// TRTC's [RoomId](https://intl.cloud.tencent.com/document/product/647/46351?from_cn_redirect=1#roomid), the RoomId corresponding to the Record TRTC room.
	RoomId *string `json:"RoomId,omitnil" name:"RoomId"`

	// Type of TRTC RoomId. 【*Note】Must be the same as the RoomId type corresponding to the Record room: 0: String type RoomId 1: 32-bit Integer type RoomId (default)
	RoomIdType *uint64 `json:"RoomIdType,omitnil" name:"RoomIdType"`

	// UserId of the Pull stream Relay Robot, used to enter the room and initiate the Pull stream Relay Task.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// UserSig corresponding to the Pull stream Relay Robot UserId, i.e., UserId and UserSig are equivalent to the Robot's Login password for entering the room. For the specific Calculation method, please refer to the TRTC [UserSig](https://intl.cloud.tencent.com/document/product/647/45910?from_cn_redirect=1#UserSig) Scheme.
	UserSig *string `json:"UserSig,omitnil" name:"UserSig"`

	// 	
	// Source URL. Example value: https://a.b/test.mp4
	SourceUrl []*string `json:"SourceUrl,omitnil" name:"SourceUrl"`

	// TRTC room permission Encryption ticket, only needed when advanced permission control is enabled in the Console. After enabling advanced permission control in the TRTC Console, TRTC's backend service system will verify a so-called [PrivateMapKey] 'Permission ticket', which contains an encrypted RoomId and an encrypted 'Permission bit list'. Since PrivateMapKey contains RoomId, providing only UserSig without PrivateMapKey does not allow entry into the specified room.
	PrivateMapKey *string `json:"PrivateMapKey,omitnil" name:"PrivateMapKey"`

	// Video Codec Parameters. Optional, if not filled, Keep original stream Parameters.
	VideoEncodeParams *VideoEncodeParams `json:"VideoEncodeParams,omitnil" name:"VideoEncodeParams"`

	// Audio Codec Parameters. Optional, if not filled, Keep original stream Parameters.
	AudioEncodeParams *AudioEncodeParams `json:"AudioEncodeParams,omitnil" name:"AudioEncodeParams"`
}

func (r *StartStreamIngestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamIngestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomId")
	delete(f, "RoomIdType")
	delete(f, "UserId")
	delete(f, "UserSig")
	delete(f, "SourceUrl")
	delete(f, "PrivateMapKey")
	delete(f, "VideoEncodeParams")
	delete(f, "AudioEncodeParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartStreamIngestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartStreamIngestResponseParams struct {
	// The Task ID of the Pull stream Relay. The Task ID is a unique identifier for a Pull stream Relay lifecycle process, and it loses its meaning when the task ends. The Task ID needs to be saved by the business as a parameter for the next operation on this task.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartStreamIngestResponse struct {
	*tchttp.BaseResponse
	Response *StartStreamIngestResponseParams `json:"Response"`
}

func (r *StartStreamIngestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartStreamIngestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishCdnStreamRequestParams struct {
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The task ID.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type StopPublishCdnStreamRequest struct {
	*tchttp.BaseRequest
	
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The task ID.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *StopPublishCdnStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishCdnStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopPublishCdnStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishCdnStreamResponseParams struct {
	// The task ID.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopPublishCdnStreamResponse struct {
	*tchttp.BaseResponse
	Response *StopPublishCdnStreamResponseParams `json:"Response"`
}

func (r *StopPublishCdnStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishCdnStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamIngestRequestParams struct {
	// The SDKAppId of TRTC, which is the same as the SDKAppId corresponding to the task's room.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique Task ID, which will be returned after the task is successfully started.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type StopStreamIngestRequest struct {
	*tchttp.BaseRequest
	
	// The SDKAppId of TRTC, which is the same as the SDKAppId corresponding to the task's room.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The unique Task ID, which will be returned after the task is successfully started.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *StopStreamIngestRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamIngestRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopStreamIngestRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopStreamIngestResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopStreamIngestResponse struct {
	*tchttp.BaseResponse
	Response *StopStreamIngestResponseParams `json:"Response"`
}

func (r *StopStreamIngestResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopStreamIngestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageFile struct {
	// The user whose stream is recorded into the file. In the mixed-stream recording mode, this parameter will be empty.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The filename.
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// The type of the media recorded.
	// video
	// audio
	// audio_video
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	TrackType *string `json:"TrackType,omitnil" name:"TrackType"`

	// The start time (Unix timestamp) of the recording file.
	BeginTimeStamp *uint64 `json:"BeginTimeStamp,omitnil" name:"BeginTimeStamp"`
}

type StorageParams struct {
	// The account information for third-party storage. Please note that if you save files to COS, a recording-to-COS fee will be incurred. For details, see the document "Billing of On-Cloud Recording". If you save files to VOD, there won't be such a fee.
	CloudStorage *CloudStorage `json:"CloudStorage,omitnil" name:"CloudStorage"`

	// The account information for VOD storage.
	CloudVod *CloudVod `json:"CloudVod,omitnil" name:"CloudVod"`
}

type SubscribeStreamUserIds struct {
	// The allowlist for audio subscription. For example, `["1", "2", "3"]` means to only subscribe to the audios of users 1, 2, and 3, and ["1.*$"] means to only subscribe to the audios of users whose ID prefix is `1`. If this parameter is left empty, the audios of all anchors in the room will be received. The array can contain at most 32 elements.
	SubscribeAudioUserIds []*string `json:"SubscribeAudioUserIds,omitnil" name:"SubscribeAudioUserIds"`

	// The blocklist for audio subscription. For example, `["1", "2", "3"]` means to not subscribe to the audios of users 1, 2, and 3, and `["1.*$"]` means to not subscribe to users whose ID prefix is `1`. If this parameter is left empty, the audios of all anchors in the room will be received. The array can contain at most 32 elements.
	UnSubscribeAudioUserIds []*string `json:"UnSubscribeAudioUserIds,omitnil" name:"UnSubscribeAudioUserIds"`

	// The allowlist for video subscription. For example, `["1", "2", "3"]` means to only subscribe to the videos of users 1, 2, and 3, and `["1.*$"]` means to only subscribe to the videos of users whose ID prefix is `1`. If this parameter is left empty, the videos of all anchors in the room will be received. The array can contain at most 32 elements.
	SubscribeVideoUserIds []*string `json:"SubscribeVideoUserIds,omitnil" name:"SubscribeVideoUserIds"`

	// The blocklist for video subscription. For example, `["1", "2", "3"]` means to not subscribe to the videos of users 1, 2, and 3, and `["1.*$"]` means to not subscribe to the videos of users whose ID prefix is `1`. If this parameter is left empty, the videos of all anchors in the room will be received. The array can contain at most 32 elements.
	UnSubscribeVideoUserIds []*string `json:"UnSubscribeVideoUserIds,omitnil" name:"UnSubscribeVideoUserIds"`
}

type TRTCDataResult struct {
	// StatementID value, fixed at 0 for Monitoring Dashboard.
	StatementID *int64 `json:"StatementID,omitnil" name:"StatementID"`

	// Query result data, returned in Columns-Values format.
	Series []*SeriesInfos `json:"Series,omitnil" name:"Series"`

	// Total value, fixed at 1 for Monitoring Dashboard.
	Total *int64 `json:"Total,omitnil" name:"Total"`
}

type TencentVod struct {
	// The operation to perform on the media uploaded. The value of this parameter is the name of a task flow template. You can create a custom task flow template in Tencent Cloud VOD.
	Procedure *string `json:"Procedure,omitnil" name:"Procedure"`

	// The expiration time of the media file, which is a time period (seconds) from the current time. For example, `86400` means to save the media file for one day. To save the file permanently, set this parameter to `0`.
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// The storage region. Set this parameter if you have special requirements on the storage region.
	StorageRegion *string `json:"StorageRegion,omitnil" name:"StorageRegion"`

	// The category ID, which is returned after you create a category by calling an API. You can use categories to manage media files.
	// The default value is `0`, which means others.
	ClassId *uint64 `json:"ClassId,omitnil" name:"ClassId"`

	// The VOD subapplication ID. If you need to access a resource in a subapplication, set this parameter to the subapplication ID; otherwise, leave it empty.
	SubAppId *uint64 `json:"SubAppId,omitnil" name:"SubAppId"`

	// The task flow context, which is passed through after the task is completed.
	SessionContext *string `json:"SessionContext,omitnil" name:"SessionContext"`

	// The upload context, which is passed through after upload is completed.
	SourceContext *string `json:"SourceContext,omitnil" name:"SourceContext"`

	// The format of recording files uploaded to VOD. `0` (default): MP4; `1`: HLS; `2`: AAC (valid only if `StreamType` is `1`); `3`: HLS+MP4; `4`: HLS+AAC.
	MediaType *uint64 `json:"MediaType,omitnil" name:"MediaType"`

	// The custom prefix of recording files. This parameter is valid only if recording files are uploaded to VOD. It can contain letters, numbers, underscores, and hyphens and cannot exceed 64 bytes. This prefix and the automatically generated filename are connected with `__UserId_u_`.
	UserDefineRecordId *string `json:"UserDefineRecordId,omitnil" name:"UserDefineRecordId"`
}

type TimeValue struct {
	// The UNIX timestamp (seconds), such as `1590065877`.
	Time *uint64 `json:"Time,omitnil" name:"Time"`

	// The metric value. For example, if the video capturing frame rate (`bigvCapFps`) at the time `1590065877` is `0`, the value of this parameter will be `0`.
	Value *float64 `json:"Value,omitnil" name:"Value"`
}

type TrtcUsage struct {
	// The time point in the format of `YYYY-MM-DD HH:mm:ss`. If more than one day is queried, `HH:mm:ss` is `00:00:00`.
	TimeKey *string `json:"TimeKey,omitnil" name:"TimeKey"`

	// The usage (minutes). Each element of this parameter corresponds to an element of `UsageKey` in the order they are listed.
	UsageValue []*float64 `json:"UsageValue,omitnil" name:"UsageValue"`
}

// Predefined struct for user
type UpdatePublishCdnStreamRequestParams struct {
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The task ID.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The sequence of a request. This parameter ensures the requests to change the parameters of the same relaying task are in the correct order. It increases each time a new request is made.
	SequenceNumber *uint64 `json:"SequenceNumber,omitnil" name:"SequenceNumber"`

	// Whether to transcode the streams. 0: No; 1: Yes.
	WithTranscoding *uint64 `json:"WithTranscoding,omitnil" name:"WithTranscoding"`

	// Pass this parameter to change the users whose audios are mixed. If you do not pass this parameter, no changes will be made.
	AudioParams *McuAudioParams `json:"AudioParams,omitnil" name:"AudioParams"`

	// Pass this parameter to change video parameters other than the codec, including the video layout, background image, background color, and watermark information. This parameter is valid only if streams are transcoded. If you do not pass it, no changes will be made.
	VideoParams *McuVideoParams `json:"VideoParams,omitnil" name:"VideoParams"`

	// Pass this parameter to change the single stream that is relayed. This parameter is valid only if streams are not transcoded. If you do not pass this parameter, no changes will be made.
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitnil" name:"SingleSubscribeParams"`

	// Pass this parameter to change the CDNs to relay to. If you do not pass this parameter, no changes will be made.
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil" name:"PublishCdnParams"`

	// The stream mixing SEI parameters.
	SeiParams *McuSeiParams `json:"SeiParams,omitnil" name:"SeiParams"`

	// The information of the room to which streams are relayed.
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitnil" name:"FeedBackRoomParams"`
}

type UpdatePublishCdnStreamRequest struct {
	*tchttp.BaseRequest
	
	// The [SDKAppID](https://intl.cloud.tencent.com/document/product/647/37714) of the TRTC room whose streams are relayed.
	SdkAppId *uint64 `json:"SdkAppId,omitnil" name:"SdkAppId"`

	// The task ID.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The sequence of a request. This parameter ensures the requests to change the parameters of the same relaying task are in the correct order. It increases each time a new request is made.
	SequenceNumber *uint64 `json:"SequenceNumber,omitnil" name:"SequenceNumber"`

	// Whether to transcode the streams. 0: No; 1: Yes.
	WithTranscoding *uint64 `json:"WithTranscoding,omitnil" name:"WithTranscoding"`

	// Pass this parameter to change the users whose audios are mixed. If you do not pass this parameter, no changes will be made.
	AudioParams *McuAudioParams `json:"AudioParams,omitnil" name:"AudioParams"`

	// Pass this parameter to change video parameters other than the codec, including the video layout, background image, background color, and watermark information. This parameter is valid only if streams are transcoded. If you do not pass it, no changes will be made.
	VideoParams *McuVideoParams `json:"VideoParams,omitnil" name:"VideoParams"`

	// Pass this parameter to change the single stream that is relayed. This parameter is valid only if streams are not transcoded. If you do not pass this parameter, no changes will be made.
	SingleSubscribeParams *SingleSubscribeParams `json:"SingleSubscribeParams,omitnil" name:"SingleSubscribeParams"`

	// Pass this parameter to change the CDNs to relay to. If you do not pass this parameter, no changes will be made.
	PublishCdnParams []*McuPublishCdnParam `json:"PublishCdnParams,omitnil" name:"PublishCdnParams"`

	// The stream mixing SEI parameters.
	SeiParams *McuSeiParams `json:"SeiParams,omitnil" name:"SeiParams"`

	// The information of the room to which streams are relayed.
	FeedBackRoomParams []*McuFeedBackRoomParams `json:"FeedBackRoomParams,omitnil" name:"FeedBackRoomParams"`
}

func (r *UpdatePublishCdnStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePublishCdnStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "TaskId")
	delete(f, "SequenceNumber")
	delete(f, "WithTranscoding")
	delete(f, "AudioParams")
	delete(f, "VideoParams")
	delete(f, "SingleSubscribeParams")
	delete(f, "PublishCdnParams")
	delete(f, "SeiParams")
	delete(f, "FeedBackRoomParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePublishCdnStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePublishCdnStreamResponseParams struct {
	// The task ID.
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdatePublishCdnStreamResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePublishCdnStreamResponseParams `json:"Response"`
}

func (r *UpdatePublishCdnStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePublishCdnStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInformation struct {
	// The room ID.
	RoomStr *string `json:"RoomStr,omitnil" name:"RoomStr"`

	// The user ID.
	UserId *string `json:"UserId,omitnil" name:"UserId"`

	// The time when the user entered the room.
	JoinTs *uint64 `json:"JoinTs,omitnil" name:"JoinTs"`

	// The time when the user left the room. If the user is still in the room, the current time will be returned.
	LeaveTs *uint64 `json:"LeaveTs,omitnil" name:"LeaveTs"`

	// The device type.
	DeviceType *string `json:"DeviceType,omitnil" name:"DeviceType"`

	// The SDK version number.
	SdkVersion *string `json:"SdkVersion,omitnil" name:"SdkVersion"`

	// The client IP address.
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`

	// Whether a user has left the room.
	Finished *bool `json:"Finished,omitnil" name:"Finished"`
}

type UserMediaStream struct {
	// The user information.
	UserInfo *MixUserInfo `json:"UserInfo,omitnil" name:"UserInfo"`

	// The stream type. 0: Camera; 1: Screen sharing. If you do not pass this parameter, 0 will be used.
	StreamType *uint64 `json:"StreamType,omitnil" name:"StreamType"`
}

type VideoEncode struct {
	// The width of the output stream (pixels). This parameter is required if audio and video are relayed. Value range: [0, 1920].
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// The height of the output stream (pixels). This parameter is required if audio and video are relayed. Value range: [0, 1080].
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// The frame rate (fps) of the output stream. This parameter is required if audio and video are relayed. Value range: [0, 60].
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// The bitrate (Kbps) of the output stream. This parameter is required if audio and video are relayed. Value range: [0, 10000].
	BitRate *uint64 `json:"BitRate,omitnil" name:"BitRate"`

	// The GOP (seconds) of the output stream. This parameter is required if audio and video are relayed. Value range: [1, 5].
	Gop *uint64 `json:"Gop,omitnil" name:"Gop"`
}

type VideoEncodeParams struct {
	// Width. Value range [0,1920], unit is pixel value.
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// Height. Value range [0,1080], unit is pixel value.
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// Frame Rate. Value range [1,60], indicating that the frame rate can be selected from 1 to 60fps.
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// Bitrate. Value range [1,10000], unit is kbps.
	BitRate *uint64 `json:"BitRate,omitnil" name:"BitRate"`

	// Gop. Value range [1,2], unit is second.
	Gop *uint64 `json:"Gop,omitnil" name:"Gop"`
}

type VideoParams struct {
	// The video width in pixels. The value of this parameter cannot be larger than 1920, and the result of multiplying `Width` and `Height` cannot exceed 1920 x 1080. The default value is `360`.
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// The video height in pixels. The value of this parameter cannot be larger than 1920, and the result of multiplying `Width` and `Height` cannot exceed 1920 x 1080. The default value is `640`.
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// The video frame rate. Value range: [1, 60]. Default: 15.
	Fps *uint64 `json:"Fps,omitnil" name:"Fps"`

	// The video bitrate (bps). Value range: [64000, 8192000]. Default: 550000.
	BitRate *uint64 `json:"BitRate,omitnil" name:"BitRate"`

	// The keyframe interval (seconds). Default value: 10.
	Gop *uint64 `json:"Gop,omitnil" name:"Gop"`
}

type WaterMark struct {
	// The watermark type. 0 (default): image; 1: text (not supported yet).
	WaterMarkType *uint64 `json:"WaterMarkType,omitnil" name:"WaterMarkType"`

	// The information of watermark images. This parameter is required if the watermark type is image.
	WaterMarkImage *WaterMarkImage `json:"WaterMarkImage,omitnil" name:"WaterMarkImage"`

	// The information of the text watermark. This parameter is required if `WaterMarkType` is `1`.
	WaterMarkChar *WaterMarkChar `json:"WaterMarkChar,omitnil" name:"WaterMarkChar"`

	// The information of the timestamp watermark. This parameter is required if `WaterMarkType` is `2`.
	WaterMarkTimestamp *WaterMarkTimestamp `json:"WaterMarkTimestamp,omitnil" name:"WaterMarkTimestamp"`
}

type WaterMarkChar struct {
	// The Y coordinate of the text watermark from the top left.
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// The X coordinate of the text watermark from the top left.
	Left *uint64 `json:"Left,omitnil" name:"Left"`

	// The watermark width (pixels).
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// The watermark height (pixels).
	Height *uint64 `json:"Height,omitnil" name:"Height"`

	// The text.
	Chars *string `json:"Chars,omitnil" name:"Chars"`

	// The font size (pixels). The default value is `14`.
	FontSize *uint64 `json:"FontSize,omitnil" name:"FontSize"`

	// The text color. The default color is white.
	FontColor *string `json:"FontColor,omitnil" name:"FontColor"`

	// The background color. If this parameter is empty, the background will be transparent (default).
	BackGroundColor *string `json:"BackGroundColor,omitnil" name:"BackGroundColor"`
}

type WaterMarkImage struct {
	// The download URLs of the watermark images, which must be in JPG or PNG format and cannot be larger than 5 MB.
	WaterMarkUrl *string `json:"WaterMarkUrl,omitnil" name:"WaterMarkUrl"`

	// The Y axis of the image's top-left corner. Value range: [0, 2560]. The value cannot be larger than the canvas height.
	Top *uint64 `json:"Top,omitnil" name:"Top"`

	// The X axis of the image’s top-left corner. Value range: [0, 2560]. The value cannot be larger than the canvas width.
	Left *uint64 `json:"Left,omitnil" name:"Left"`

	// The relative width of the image. Value range: [0, 2560]. The sum of the values of this parameter and `Left` cannot exceed the canvas width.
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// The relative height of the image. Value range: [0, 2560]. The sum of the values of this parameter and `Top` cannot exceed the canvas height.
	Height *uint64 `json:"Height,omitnil" name:"Height"`
}

type WaterMarkTimestamp struct {
	// The position of the timestamp watermark. Valid values: `0` (top left), `1` (top right), `2` (bottom left), `3` (bottom right), `4` (top center), `5` (bottom center), `6` (center).
	Pos *uint64 `json:"Pos,omitnil" name:"Pos"`

	// The time zone. The default is UTC+8.
	TimeZone *uint64 `json:"TimeZone,omitnil" name:"TimeZone"`
}