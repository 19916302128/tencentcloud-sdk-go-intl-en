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

package v20210111

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-01-11"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddSmsSignRequest() (request *AddSmsSignRequest) {
    request = &AddSmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "AddSmsSign")
    return
}

func NewAddSmsSignResponse() (response *AddSmsSignResponse) {
    response = &AddSmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddSmsSign
// 1. This API is used to add an SMS signature. Please read the [Tencent Cloud SMS Signature Review Standards](https://intl.cloud.tencent.com/document/product/382/39022?from_cn_redirect=1) before starting an application.
//
// 2. ⚠️ Note: individual users cannot use this API to apply for SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, you can log in to the console to apply for SMS signatures.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_PARAMETERSOTHERERROR = "FailedOperation.ParametersOtherError"
//  FAILEDOPERATION_SIGNNUMBERLIMIT = "FailedOperation.SignNumberLimit"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDUSEDMETHOD = "InvalidParameterValue.InvalidUsedMethod"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  INVALIDPARAMETERVALUE_SIGNALREADYPASSEDCHECK = "InvalidParameterValue.SignAlreadyPassedCheck"
//  INVALIDPARAMETERVALUE_SIGNEXISTANDUNAPPROVED = "InvalidParameterValue.SignExistAndUnapproved"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) AddSmsSign(request *AddSmsSignRequest) (response *AddSmsSignResponse, err error) {
    if request == nil {
        request = NewAddSmsSignRequest()
    }
    response = NewAddSmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewAddSmsTemplateRequest() (request *AddSmsTemplateRequest) {
    request = &AddSmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "AddSmsTemplate")
    return
}

func NewAddSmsTemplateResponse() (response *AddSmsTemplateResponse) {
    response = &AddSmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddSmsTemplate
// 1. This API is used to add an SMS template. Please read the [Tencent Cloud SMS Body Template Review Standards](https://intl.cloud.tencent.com/document/product/382/39023?from_cn_redirect=1) before starting an application.
//
// 2. ⚠️ Note: individual users cannot use this API to apply for SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, you can log in to the [console](https://console.cloud.tencent.com/smsv2) to apply for SMS body templates.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURE = "FailedOperation.MissingSignature"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_DIRTYWORDFOUND = "InvalidParameter.DirtyWordFound"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) AddSmsTemplate(request *AddSmsTemplateRequest) (response *AddSmsTemplateResponse, err error) {
    if request == nil {
        request = NewAddSmsTemplateRequest()
    }
    response = NewAddSmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCallbackStatusStatisticsRequest() (request *CallbackStatusStatisticsRequest) {
    request = &CallbackStatusStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "CallbackStatusStatistics")
    return
}

func NewCallbackStatusStatisticsResponse() (response *CallbackStatusStatisticsResponse) {
    response = &CallbackStatusStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CallbackStatusStatistics
// This API is used to collect statistics on user receipts.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_OFFSETVERIFYFAIL = "InvalidParameterValue.OffsetVerifyFail"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) CallbackStatusStatistics(request *CallbackStatusStatisticsRequest) (response *CallbackStatusStatisticsResponse, err error) {
    if request == nil {
        request = NewCallbackStatusStatisticsRequest()
    }
    response = NewCallbackStatusStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmsSignRequest() (request *DeleteSmsSignRequest) {
    request = &DeleteSmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "DeleteSmsSign")
    return
}

func NewDeleteSmsSignResponse() (response *DeleteSmsSignResponse) {
    response = &DeleteSmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSmsSign
// ⚠️ Note: individual users cannot use this API to delete SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). Please log in to the [console](https://console.cloud.tencent.com/smsv2) to delete SMS signatures.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DeleteSmsSign(request *DeleteSmsSignRequest) (response *DeleteSmsSignResponse, err error) {
    if request == nil {
        request = NewDeleteSmsSignRequest()
    }
    response = NewDeleteSmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSmsTemplateRequest() (request *DeleteSmsTemplateRequest) {
    request = &DeleteSmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "DeleteSmsTemplate")
    return
}

func NewDeleteSmsTemplateResponse() (response *DeleteSmsTemplateResponse) {
    response = &DeleteSmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSmsTemplate
// ⚠️ Note: individual users cannot use this API to delete SMS body templates. Please log in to the [console](https://console.cloud.tencent.com/smsv2) to do so. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1).
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DeleteSmsTemplate(request *DeleteSmsTemplateRequest) (response *DeleteSmsTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteSmsTemplateRequest()
    }
    response = NewDeleteSmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsSignListRequest() (request *DescribeSmsSignListRequest) {
    request = &DescribeSmsSignListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "DescribeSmsSignList")
    return
}

func NewDescribeSmsSignListResponse() (response *DescribeSmsSignListResponse) {
    response = &DescribeSmsSignListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmsSignList
// ⚠️ Note: individual users cannot use this API to query SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, you can log in to the [console](https://console.cloud.tencent.com/smsv2) to query SMS signatures.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DescribeSmsSignList(request *DescribeSmsSignListRequest) (response *DescribeSmsSignListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsSignListRequest()
    }
    response = NewDescribeSmsSignListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsTemplateListRequest() (request *DescribeSmsTemplateListRequest) {
    request = &DescribeSmsTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "DescribeSmsTemplateList")
    return
}

func NewDescribeSmsTemplateListResponse() (response *DescribeSmsTemplateListResponse) {
    response = &DescribeSmsTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmsTemplateList
// ⚠️ Note: individual users cannot use this API to query SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1).
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"
//  INVALIDPARAMETERVALUE_TEMPLATEWITHDIRTYWORDS = "InvalidParameterValue.TemplateWithDirtyWords"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) DescribeSmsTemplateList(request *DescribeSmsTemplateListRequest) (response *DescribeSmsTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsTemplateListRequest()
    }
    response = NewDescribeSmsTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmsSignRequest() (request *ModifySmsSignRequest) {
    request = &ModifySmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "ModifySmsSign")
    return
}

func NewModifySmsSignResponse() (response *ModifySmsSignResponse) {
    response = &ModifySmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySmsSign
// 1. This API is used to modify an SMS signature. Please read the [Tencent Cloud SMS Signature Review Standards](https://intl.cloud.tencent.com/document/product/382/39022?from_cn_redirect=1) before making a modification.
//
// 2. ⚠️ Note: individual users cannot use this API to modify SMS signatures. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, you can log in to the [console](https://console.cloud.tencent.com/smsv2) to modify SMS signatures.
//
// 3. Modifications can be made only if the signature status is **Pending Review** or **Rejected**. **Approved** signatures cannot be modified.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MISSINGSIGNATURETOMODIFY = "FailedOperation.MissingSignatureToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_SIGNNUMBERLIMIT = "FailedOperation.SignNumberLimit"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPIDANDBIZID = "InvalidParameter.AppidAndBizId"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_INVALIDDOCUMENTTYPE = "InvalidParameterValue.InvalidDocumentType"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDSIGNPURPOSE = "InvalidParameterValue.InvalidSignPurpose"
//  INVALIDPARAMETERVALUE_MISSINGSIGNATURELIST = "InvalidParameterValue.MissingSignatureList"
//  INVALIDPARAMETERVALUE_SIGNALREADYPASSEDCHECK = "InvalidParameterValue.SignAlreadyPassedCheck"
//  INVALIDPARAMETERVALUE_SIGNEXISTANDUNAPPROVED = "InvalidParameterValue.SignExistAndUnapproved"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) ModifySmsSign(request *ModifySmsSignRequest) (response *ModifySmsSignResponse, err error) {
    if request == nil {
        request = NewModifySmsSignRequest()
    }
    response = NewModifySmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmsTemplateRequest() (request *ModifySmsTemplateRequest) {
    request = &ModifySmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "ModifySmsTemplate")
    return
}

func NewModifySmsTemplateResponse() (response *ModifySmsTemplateResponse) {
    response = &ModifySmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySmsTemplate
// 1. This API is used to modify an SMS body template. Please read the [Tencent Cloud SMS Body Template Review Standards](https://intl.cloud.tencent.com/document/product/382/39023?from_cn_redirect=1) before making a modification.
//
// 2. ⚠️ Note: individual users cannot use this API to modify SMS body templates. For more information, please see [Identity Verification Overview](https://intl.cloud.tencent.com/document/product/378/3629?from_cn_redirect=1). If your account identity is individual, you can log in to the [console](https://console.cloud.tencent.com/smsv2) to modify SMS body templates.
//
// 3. Modifications can be made only if the body template status is **Pending Review** or **Rejected**. **Approved** body templates cannot be modified.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2019-07-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_FORBIDADDMARKETINGTEMPLATES = "FailedOperation.ForbidAddMarketingTemplates"
//  FAILEDOPERATION_MISSINGTEMPLATETOMODIFY = "FailedOperation.MissingTemplateToModify"
//  FAILEDOPERATION_NOTENTERPRISECERTIFICATION = "FailedOperation.NotEnterpriseCertification"
//  FAILEDOPERATION_OTHERERROR = "FailedOperation.OtherError"
//  FAILEDOPERATION_TEMPLATEALREADYPASSEDCHECK = "FailedOperation.TemplateAlreadyPassedCheck"
//  FAILEDOPERATION_TEMPLATENUMBERLIMIT = "FailedOperation.TemplateNumberLimit"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_DIRTYWORDFOUND = "InvalidParameter.DirtyWordFound"
//  INVALIDPARAMETER_INVALIDPARAMETERS = "InvalidParameter.InvalidParameters"
//  INVALIDPARAMETERVALUE_INVALIDINTERNATIONAL = "InvalidParameterValue.InvalidInternational"
//  INVALIDPARAMETERVALUE_INVALIDTEMPLATEFORMAT = "InvalidParameterValue.InvalidTemplateFormat"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) ModifySmsTemplate(request *ModifySmsTemplateRequest) (response *ModifySmsTemplateResponse, err error) {
    if request == nil {
        request = NewModifySmsTemplateRequest()
    }
    response = NewModifySmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsReplyStatusRequest() (request *PullSmsReplyStatusRequest) {
    request = &PullSmsReplyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsReplyStatus")
    return
}

func NewPullSmsReplyStatusResponse() (response *PullSmsReplyStatusResponse) {
    response = &PullSmsReplyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PullSmsReplyStatus
// This API is used to pull SMS reply status.
//
// Currently, you can also [configure the reply callback](https://intl.cloud.tencent.com/document/product/382/42907?from_cn_redirect=1) to get replies.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsReplyStatus(request *PullSmsReplyStatusRequest) (response *PullSmsReplyStatusResponse, err error) {
    if request == nil {
        request = NewPullSmsReplyStatusRequest()
    }
    response = NewPullSmsReplyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsReplyStatusByPhoneNumberRequest() (request *PullSmsReplyStatusByPhoneNumberRequest) {
    request = &PullSmsReplyStatusByPhoneNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsReplyStatusByPhoneNumber")
    return
}

func NewPullSmsReplyStatusByPhoneNumberResponse() (response *PullSmsReplyStatusByPhoneNumberResponse) {
    response = &PullSmsReplyStatusByPhoneNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PullSmsReplyStatusByPhoneNumber
// This API is used to pull SMS reply status for one single number.
//
// Currently, you can also [configure the reply callback](https://intl.cloud.tencent.com/document/product/382/42907?from_cn_redirect=1) to get replies.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsReplyStatusByPhoneNumber(request *PullSmsReplyStatusByPhoneNumberRequest) (response *PullSmsReplyStatusByPhoneNumberResponse, err error) {
    if request == nil {
        request = NewPullSmsReplyStatusByPhoneNumberRequest()
    }
    response = NewPullSmsReplyStatusByPhoneNumberResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsSendStatusRequest() (request *PullSmsSendStatusRequest) {
    request = &PullSmsSendStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsSendStatus")
    return
}

func NewPullSmsSendStatusResponse() (response *PullSmsSendStatusResponse) {
    response = &PullSmsSendStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PullSmsSendStatus
// This API is used to pull SMS delivery status.
//
// Currently, you can also [configure the callback](https://intl.cloud.tencent.com/document/product/382/37809?from_cn_redirect=1#.E8.AE.BE.E7.BD.AE.E4.BA.8B.E4.BB.B6.E5.9B.9E.E8.B0.83.E9.85.8D.E7.BD.AE) to get the delivery status.
//
// >- Note: you need to contact [SMS Helper](https://intl.cloud.tencent.com/document/product/382/3773?from_cn_redirect=1#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81) to activate this API.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsSendStatus(request *PullSmsSendStatusRequest) (response *PullSmsSendStatusResponse, err error) {
    if request == nil {
        request = NewPullSmsSendStatusRequest()
    }
    response = NewPullSmsSendStatusResponse()
    err = c.Send(request, response)
    return
}

func NewPullSmsSendStatusByPhoneNumberRequest() (request *PullSmsSendStatusByPhoneNumberRequest) {
    request = &PullSmsSendStatusByPhoneNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "PullSmsSendStatusByPhoneNumber")
    return
}

func NewPullSmsSendStatusByPhoneNumberResponse() (response *PullSmsSendStatusByPhoneNumberResponse) {
    response = &PullSmsSendStatusByPhoneNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PullSmsSendStatusByPhoneNumber
// This API is used to pull SMS delivery status for one single number.
//
// Currently, you can also [configure the callback](https://intl.cloud.tencent.com/document/product/382/37809?from_cn_redirect=1#.E8.AE.BE.E7.BD.AE.E4.BA.8B.E4.BB.B6.E5.9B.9E.E8.B0.83.E9.85.8D.E7.BD.AE) to get the delivery status.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11
//
// &Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERONBLACKLIST = "FailedOperation.PhoneNumberOnBlacklist"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_INVALIDSTARTTIME = "InvalidParameterValue.InvalidStartTime"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) PullSmsSendStatusByPhoneNumber(request *PullSmsSendStatusByPhoneNumberRequest) (response *PullSmsSendStatusByPhoneNumberResponse, err error) {
    if request == nil {
        request = NewPullSmsSendStatusByPhoneNumberRequest()
    }
    response = NewPullSmsSendStatusByPhoneNumberResponse()
    err = c.Send(request, response)
    return
}

func NewSendSmsRequest() (request *SendSmsRequest) {
    request = &SendSmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "SendSms")
    return
}

func NewSendSmsResponse() (response *SendSmsResponse) {
    response = &SendSmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendSms
// This API is used to send SMS verification codes, notification, or marketing messages to users.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_CONTAINSENSITIVEWORD = "FailedOperation.ContainSensitiveWord"
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_INSUFFICIENTBALANCEINSMSPACKAGE = "FailedOperation.InsufficientBalanceInSmsPackage"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_MARKETINGSENDTIMECONSTRAINT = "FailedOperation.MarketingSendTimeConstraint"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_SIGNATUREINCORRECTORUNAPPROVED = "FailedOperation.SignatureIncorrectOrUnapproved"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CONTENTLENGTHLIMIT = "InvalidParameterValue.ContentLengthLimit"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_PROHIBITEDUSEURLINTEMPLATEPARAMETER = "InvalidParameterValue.ProhibitedUseUrlInTemplateParameter"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERFORMATERROR = "InvalidParameterValue.TemplateParameterFormatError"
//  INVALIDPARAMETERVALUE_TEMPLATEPARAMETERLENGTHLIMIT = "InvalidParameterValue.TemplateParameterLengthLimit"
//  LIMITEXCEEDED_APPDAILYLIMIT = "LimitExceeded.AppDailyLimit"
//  LIMITEXCEEDED_DAILYLIMIT = "LimitExceeded.DailyLimit"
//  LIMITEXCEEDED_DELIVERYFREQUENCYLIMIT = "LimitExceeded.DeliveryFrequencyLimit"
//  LIMITEXCEEDED_PHONENUMBERCOUNTLIMIT = "LimitExceeded.PhoneNumberCountLimit"
//  LIMITEXCEEDED_PHONENUMBERDAILYLIMIT = "LimitExceeded.PhoneNumberDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERONEHOURLIMIT = "LimitExceeded.PhoneNumberOneHourLimit"
//  LIMITEXCEEDED_PHONENUMBERSAMECONTENTDAILYLIMIT = "LimitExceeded.PhoneNumberSameContentDailyLimit"
//  LIMITEXCEEDED_PHONENUMBERTHIRTYSECONDLIMIT = "LimitExceeded.PhoneNumberThirtySecondLimit"
//  MISSINGPARAMETER_EMPTYPHONENUMBERSET = "MissingParameter.EmptyPhoneNumberSet"
//  UNAUTHORIZEDOPERATION_INDIVIDUALUSERMARKETINGSMSPERMISSIONDENY = "UnauthorizedOperation.IndividualUserMarketingSmsPermissionDeny"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
//  UNSUPPORTEDOPERATION_CONTAINDOMESTICANDINTERNATIONALPHONENUMBER = "UnsupportedOperation.ContainDomesticAndInternationalPhoneNumber"
//  UNSUPPORTEDOPERATION_UNSUPORTEDREGION = "UnsupportedOperation.UnsuportedRegion"
func (c *Client) SendSms(request *SendSmsRequest) (response *SendSmsResponse, err error) {
    if request == nil {
        request = NewSendSmsRequest()
    }
    response = NewSendSmsResponse()
    err = c.Send(request, response)
    return
}

func NewSendStatusStatisticsRequest() (request *SendStatusStatisticsRequest) {
    request = &SendStatusStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sms", APIVersion, "SendStatusStatistics")
    return
}

func NewSendStatusStatisticsResponse() (response *SendStatusStatisticsResponse) {
    response = &SendStatusStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendStatusStatistics
// This API is used to collect statistics on SMS messages sent by users.
//
// >- Note: because of the improved security of **TencentCloud API 3.0**, **API authentication** is more complicated. We recommend you use the Tencent Cloud SMS service with the SDK.
//
// >- You can run this API directly in [API 3.0 Explorer](https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2021-01-11&Action=SendSms), which eliminates the signature calculation steps. After it is executed successfully, API Explorer can **automatically generate** SDK code samples.
//
// error code that may be returned:
//  FAILEDOPERATION_FAILRESOLVEPACKET = "FailedOperation.FailResolvePacket"
//  FAILEDOPERATION_JSONPARSEFAIL = "FailedOperation.JsonParseFail"
//  FAILEDOPERATION_PHONENUMBERINBLACKLIST = "FailedOperation.PhoneNumberInBlacklist"
//  FAILEDOPERATION_TEMPLATEINCORRECTORUNAPPROVED = "FailedOperation.TemplateIncorrectOrUnapproved"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_REQUESTTIMEEXCEPTION = "InternalError.RequestTimeException"
//  INTERNALERROR_RESTAPIINTERFACENOTEXIST = "InternalError.RestApiInterfaceNotExist"
//  INTERNALERROR_SENDANDRECVFAIL = "InternalError.SendAndRecvFail"
//  INTERNALERROR_SIGFIELDMISSING = "InternalError.SigFieldMissing"
//  INTERNALERROR_SIGVERIFICATIONFAIL = "InternalError.SigVerificationFail"
//  INTERNALERROR_TIMEOUT = "InternalError.Timeout"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BEGINTIMEVERIFYFAIL = "InvalidParameterValue.BeginTimeVerifyFail"
//  INVALIDPARAMETERVALUE_ENDTIMEVERIFYFAIL = "InvalidParameterValue.EndTimeVerifyFail"
//  INVALIDPARAMETERVALUE_INCORRECTPHONENUMBER = "InvalidParameterValue.IncorrectPhoneNumber"
//  INVALIDPARAMETERVALUE_LIMITVERIFYFAIL = "InvalidParameterValue.LimitVerifyFail"
//  INVALIDPARAMETERVALUE_OFFSETVERIFYFAIL = "InvalidParameterValue.OffsetVerifyFail"
//  INVALIDPARAMETERVALUE_SDKAPPIDNOTEXIST = "InvalidParameterValue.SdkAppIdNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_REQUESTIPNOTINWHITELIST = "UnauthorizedOperation.RequestIpNotInWhitelist"
//  UNAUTHORIZEDOPERATION_REQUESTPERMISSIONDENY = "UnauthorizedOperation.RequestPermissionDeny"
//  UNAUTHORIZEDOPERATION_SDKAPPIDISDISABLED = "UnauthorizedOperation.SdkAppIdIsDisabled"
//  UNAUTHORIZEDOPERATION_SERIVCESUSPENDDUETOARREARS = "UnauthorizedOperation.SerivceSuspendDueToArrears"
//  UNAUTHORIZEDOPERATION_SMSSDKAPPIDVERIFYFAIL = "UnauthorizedOperation.SmsSdkAppIdVerifyFail"
//  UNSUPPORTEDOPERATION_ = "UnsupportedOperation."
func (c *Client) SendStatusStatistics(request *SendStatusStatisticsRequest) (response *SendStatusStatisticsResponse, err error) {
    if request == nil {
        request = NewSendStatusStatisticsRequest()
    }
    response = NewSendStatusStatisticsResponse()
    err = c.Send(request, response)
    return
}