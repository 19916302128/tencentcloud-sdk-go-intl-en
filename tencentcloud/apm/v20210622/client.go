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

package v20210622

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-06-22"

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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateApmInstanceRequest() (request *CreateApmInstanceRequest) {
    request = &CreateApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "CreateApmInstance")
    
    
    return
}

func NewCreateApmInstanceResponse() (response *CreateApmInstanceResponse) {
    response = &CreateApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApmInstance
// This API is used to create a business purchase in the APM business system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_REGIONNOTSUPPORT = "FailedOperation.RegionNotSupport"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) CreateApmInstance(request *CreateApmInstanceRequest) (response *CreateApmInstanceResponse, err error) {
    return c.CreateApmInstanceWithContext(context.Background(), request)
}

// CreateApmInstance
// This API is used to create a business purchase in the APM business system.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_REGIONNOTSUPPORT = "FailedOperation.RegionNotSupport"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) CreateApmInstanceWithContext(ctx context.Context, request *CreateApmInstanceRequest) (response *CreateApmInstanceResponse, err error) {
    if request == nil {
        request = NewCreateApmInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmAgentRequest() (request *DescribeApmAgentRequest) {
    request = &DescribeApmAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmAgent")
    
    
    return
}

func NewDescribeApmAgentResponse() (response *DescribeApmAgentResponse) {
    response = &DescribeApmAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmAgent
// Obtaining APM Access Point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_NOTINNERVPC = "FailedOperation.NotInnerVPC"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApmAgent(request *DescribeApmAgentRequest) (response *DescribeApmAgentResponse, err error) {
    return c.DescribeApmAgentWithContext(context.Background(), request)
}

// DescribeApmAgent
// Obtaining APM Access Point.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_NOTINNERVPC = "FailedOperation.NotInnerVPC"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeApmAgentWithContext(ctx context.Context, request *DescribeApmAgentRequest) (response *DescribeApmAgentResponse, err error) {
    if request == nil {
        request = NewDescribeApmAgentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApmInstancesRequest() (request *DescribeApmInstancesRequest) {
    request = &DescribeApmInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeApmInstances")
    
    
    return
}

func NewDescribeApmInstancesResponse() (response *DescribeApmInstancesResponse) {
    response = &DescribeApmInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApmInstances
// This API is used to obtain the list of APM business systems.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmInstances(request *DescribeApmInstancesRequest) (response *DescribeApmInstancesResponse, err error) {
    return c.DescribeApmInstancesWithContext(context.Background(), request)
}

// DescribeApmInstances
// This API is used to obtain the list of APM business systems.
//
// error code that may be returned:
//  AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"
//  AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
func (c *Client) DescribeApmInstancesWithContext(ctx context.Context, request *DescribeApmInstancesRequest) (response *DescribeApmInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeApmInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApmInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApmInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralApmApplicationConfigRequest() (request *DescribeGeneralApmApplicationConfigRequest) {
    request = &DescribeGeneralApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralApmApplicationConfig")
    
    
    return
}

func NewDescribeGeneralApmApplicationConfigResponse() (response *DescribeGeneralApmApplicationConfigResponse) {
    response = &DescribeGeneralApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralApmApplicationConfig
// This API is used to query the application configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeGeneralApmApplicationConfig(request *DescribeGeneralApmApplicationConfigRequest) (response *DescribeGeneralApmApplicationConfigResponse, err error) {
    return c.DescribeGeneralApmApplicationConfigWithContext(context.Background(), request)
}

// DescribeGeneralApmApplicationConfig
// This API is used to query the application configuration information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"
func (c *Client) DescribeGeneralApmApplicationConfigWithContext(ctx context.Context, request *DescribeGeneralApmApplicationConfigRequest) (response *DescribeGeneralApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralApmApplicationConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralMetricDataRequest() (request *DescribeGeneralMetricDataRequest) {
    request = &DescribeGeneralMetricDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralMetricData")
    
    
    return
}

func NewDescribeGeneralMetricDataResponse() (response *DescribeGeneralMetricDataResponse) {
    response = &DescribeGeneralMetricDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralMetricData
// This API is a general API used to obtain metric data. Users submit request parameters as needed and receive the corresponding metric data.
//
// The API call frequency is limited to 20 requests per second and 1200 requests per minute. The number of data points per request is limited to 1440.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_METRICFILTERSLACKPARAMS = "FailedOperation.MetricFiltersLackParams"
//  FAILEDOPERATION_QUERYTIMEINTERVALISNOTSUPPORTED = "FailedOperation.QueryTimeIntervalIsNotSupported"
//  FAILEDOPERATION_VIEWNAMENOTEXISTORILLEGAL = "FailedOperation.ViewNameNotExistOrIllegal"
//  INVALIDPARAMETER_FILTERSFIELDSNOTEXISTORILLEGAL = "InvalidParameter.FiltersFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_GROUPBYFIELDSNOTEXISTORILLEGAL = "InvalidParameter.GroupByFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_METRICFILTERSLACKPARAMS = "InvalidParameter.MetricFiltersLackParams"
//  INVALIDPARAMETER_METRICSFIELDNOTEXISTORILLEGAL = "InvalidParameter.MetricsFieldNotExistOrIllegal"
//  INVALIDPARAMETER_METRICSFIELDSNOTALLOWEMPTY = "InvalidParameter.MetricsFieldsNotAllowEmpty"
//  INVALIDPARAMETER_PERIODISILLEGAL = "InvalidParameter.PeriodIsIllegal"
//  INVALIDPARAMETER_QUERYTIMEINTERVALISNOTSUPPORTED = "InvalidParameter.QueryTimeIntervalIsNotSupported"
//  INVALIDPARAMETER_VIEWNAMENOTEXISTORILLEGAL = "InvalidParameter.ViewNameNotExistOrIllegal"
func (c *Client) DescribeGeneralMetricData(request *DescribeGeneralMetricDataRequest) (response *DescribeGeneralMetricDataResponse, err error) {
    return c.DescribeGeneralMetricDataWithContext(context.Background(), request)
}

// DescribeGeneralMetricData
// This API is a general API used to obtain metric data. Users submit request parameters as needed and receive the corresponding metric data.
//
// The API call frequency is limited to 20 requests per second and 1200 requests per minute. The number of data points per request is limited to 1440.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_METRICFILTERSLACKPARAMS = "FailedOperation.MetricFiltersLackParams"
//  FAILEDOPERATION_QUERYTIMEINTERVALISNOTSUPPORTED = "FailedOperation.QueryTimeIntervalIsNotSupported"
//  FAILEDOPERATION_VIEWNAMENOTEXISTORILLEGAL = "FailedOperation.ViewNameNotExistOrIllegal"
//  INVALIDPARAMETER_FILTERSFIELDSNOTEXISTORILLEGAL = "InvalidParameter.FiltersFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_GROUPBYFIELDSNOTEXISTORILLEGAL = "InvalidParameter.GroupByFieldsNotExistOrIllegal"
//  INVALIDPARAMETER_METRICFILTERSLACKPARAMS = "InvalidParameter.MetricFiltersLackParams"
//  INVALIDPARAMETER_METRICSFIELDNOTEXISTORILLEGAL = "InvalidParameter.MetricsFieldNotExistOrIllegal"
//  INVALIDPARAMETER_METRICSFIELDSNOTALLOWEMPTY = "InvalidParameter.MetricsFieldsNotAllowEmpty"
//  INVALIDPARAMETER_PERIODISILLEGAL = "InvalidParameter.PeriodIsIllegal"
//  INVALIDPARAMETER_QUERYTIMEINTERVALISNOTSUPPORTED = "InvalidParameter.QueryTimeIntervalIsNotSupported"
//  INVALIDPARAMETER_VIEWNAMENOTEXISTORILLEGAL = "InvalidParameter.ViewNameNotExistOrIllegal"
func (c *Client) DescribeGeneralMetricDataWithContext(ctx context.Context, request *DescribeGeneralMetricDataRequest) (response *DescribeGeneralMetricDataResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralMetricDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralMetricData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralMetricDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralOTSpanListRequest() (request *DescribeGeneralOTSpanListRequest) {
    request = &DescribeGeneralOTSpanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralOTSpanList")
    
    
    return
}

func NewDescribeGeneralOTSpanListResponse() (response *DescribeGeneralOTSpanListResponse) {
    response = &DescribeGeneralOTSpanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralOTSpanList
// General Query OpenTelemetry Call Chain List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralOTSpanList(request *DescribeGeneralOTSpanListRequest) (response *DescribeGeneralOTSpanListResponse, err error) {
    return c.DescribeGeneralOTSpanListWithContext(context.Background(), request)
}

// DescribeGeneralOTSpanList
// General Query OpenTelemetry Call Chain List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralOTSpanListWithContext(ctx context.Context, request *DescribeGeneralOTSpanListRequest) (response *DescribeGeneralOTSpanListResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralOTSpanListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralOTSpanList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralOTSpanListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralSpanListRequest() (request *DescribeGeneralSpanListRequest) {
    request = &DescribeGeneralSpanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralSpanList")
    
    
    return
}

func NewDescribeGeneralSpanListResponse() (response *DescribeGeneralSpanListResponse) {
    response = &DescribeGeneralSpanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGeneralSpanList
// General Query Call Chain List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralSpanList(request *DescribeGeneralSpanListRequest) (response *DescribeGeneralSpanListResponse, err error) {
    return c.DescribeGeneralSpanListWithContext(context.Background(), request)
}

// DescribeGeneralSpanList
// General Query Call Chain List.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGeneralSpanListWithContext(ctx context.Context, request *DescribeGeneralSpanListRequest) (response *DescribeGeneralSpanListResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralSpanListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGeneralSpanList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGeneralSpanListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricRecordsRequest() (request *DescribeMetricRecordsRequest) {
    request = &DescribeMetricRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeMetricRecords")
    
    
    return
}

func NewDescribeMetricRecordsResponse() (response *DescribeMetricRecordsResponse) {
    response = &DescribeMetricRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricRecords
// This API is used to query metric list. To query metrics, it is recommended to use the DescribeGeneralMetricData API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricRecords(request *DescribeMetricRecordsRequest) (response *DescribeMetricRecordsResponse, err error) {
    return c.DescribeMetricRecordsWithContext(context.Background(), request)
}

// DescribeMetricRecords
// This API is used to query metric list. To query metrics, it is recommended to use the DescribeGeneralMetricData API.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeMetricRecordsWithContext(ctx context.Context, request *DescribeMetricRecordsRequest) (response *DescribeMetricRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeMetricRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceOverviewRequest() (request *DescribeServiceOverviewRequest) {
    request = &DescribeServiceOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeServiceOverview")
    
    
    return
}

func NewDescribeServiceOverviewResponse() (response *DescribeServiceOverviewResponse) {
    response = &DescribeServiceOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServiceOverview
// This API is used to pull application overview data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServiceOverview(request *DescribeServiceOverviewRequest) (response *DescribeServiceOverviewResponse, err error) {
    return c.DescribeServiceOverviewWithContext(context.Background(), request)
}

// DescribeServiceOverview
// This API is used to pull application overview data.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServiceOverviewWithContext(ctx context.Context, request *DescribeServiceOverviewRequest) (response *DescribeServiceOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeServiceOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagValuesRequest() (request *DescribeTagValuesRequest) {
    request = &DescribeTagValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "DescribeTagValues")
    
    
    return
}

func NewDescribeTagValuesResponse() (response *DescribeTagValuesResponse) {
    response = &DescribeTagValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagValues
// This API is used to query dimensional data by dimension name and filter condition.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DescribeTagValues(request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
    return c.DescribeTagValuesWithContext(context.Background(), request)
}

// DescribeTagValues
// This API is used to query dimensional data by dimension name and filter condition.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
func (c *Client) DescribeTagValuesWithContext(ctx context.Context, request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
    if request == nil {
        request = NewDescribeTagValuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagValuesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApmInstanceRequest() (request *ModifyApmInstanceRequest) {
    request = &ModifyApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyApmInstance")
    
    
    return
}

func NewModifyApmInstanceResponse() (response *ModifyApmInstanceResponse) {
    response = &ModifyApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApmInstance
// This API is used to modify the APM business system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTMODIFY = "FailedOperation.InstanceCannotModify"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyApmInstance(request *ModifyApmInstanceRequest) (response *ModifyApmInstanceResponse, err error) {
    return c.ModifyApmInstanceWithContext(context.Background(), request)
}

// ModifyApmInstance
// This API is used to modify the APM business system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTMODIFY = "FailedOperation.InstanceCannotModify"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"
//  FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"
//  FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyApmInstanceWithContext(ctx context.Context, request *ModifyApmInstanceRequest) (response *ModifyApmInstanceResponse, err error) {
    if request == nil {
        request = NewModifyApmInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApmInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGeneralApmApplicationConfigRequest() (request *ModifyGeneralApmApplicationConfigRequest) {
    request = &ModifyGeneralApmApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "ModifyGeneralApmApplicationConfig")
    
    
    return
}

func NewModifyGeneralApmApplicationConfigResponse() (response *ModifyGeneralApmApplicationConfigResponse) {
    response = &ModifyGeneralApmApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGeneralApmApplicationConfig
// OpenAPI available for external use. Customers can flexibly specify the fields to be modified, and then add the list of services to be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APMCREDENTIALNOTEXIST = "FailedOperation.ApmCredentialNotExist"
//  FAILEDOPERATION_DUPLICATESERVICE = "FailedOperation.DuplicateService"
//  FAILEDOPERATION_DUPLICATETAGFIELD = "FailedOperation.DuplicateTagField"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
//  FAILEDOPERATION_INVALIDTAGFIELD = "FailedOperation.InvalidTagField"
//  FAILEDOPERATION_INVALIDTOKEN = "FailedOperation.InvalidToken"
//  FAILEDOPERATION_SERVICELISTEXCEEDINGLIMITNUMBER = "FailedOperation.ServiceListExceedingLimitNumber"
//  FAILEDOPERATION_SERVICELISTNULL = "FailedOperation.ServiceListNull"
func (c *Client) ModifyGeneralApmApplicationConfig(request *ModifyGeneralApmApplicationConfigRequest) (response *ModifyGeneralApmApplicationConfigResponse, err error) {
    return c.ModifyGeneralApmApplicationConfigWithContext(context.Background(), request)
}

// ModifyGeneralApmApplicationConfig
// OpenAPI available for external use. Customers can flexibly specify the fields to be modified, and then add the list of services to be modified.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APMCREDENTIALNOTEXIST = "FailedOperation.ApmCredentialNotExist"
//  FAILEDOPERATION_DUPLICATESERVICE = "FailedOperation.DuplicateService"
//  FAILEDOPERATION_DUPLICATETAGFIELD = "FailedOperation.DuplicateTagField"
//  FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"
//  FAILEDOPERATION_INVALIDTAGFIELD = "FailedOperation.InvalidTagField"
//  FAILEDOPERATION_INVALIDTOKEN = "FailedOperation.InvalidToken"
//  FAILEDOPERATION_SERVICELISTEXCEEDINGLIMITNUMBER = "FailedOperation.ServiceListExceedingLimitNumber"
//  FAILEDOPERATION_SERVICELISTNULL = "FailedOperation.ServiceListNull"
func (c *Client) ModifyGeneralApmApplicationConfigWithContext(ctx context.Context, request *ModifyGeneralApmApplicationConfigRequest) (response *ModifyGeneralApmApplicationConfigResponse, err error) {
    if request == nil {
        request = NewModifyGeneralApmApplicationConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGeneralApmApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGeneralApmApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateApmInstanceRequest() (request *TerminateApmInstanceRequest) {
    request = &TerminateApmInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apm", APIVersion, "TerminateApmInstance")
    
    
    return
}

func NewTerminateApmInstanceResponse() (response *TerminateApmInstanceResponse) {
    response = &TerminateApmInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateApmInstance
// Termination of APM business system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTTERMINATE = "FailedOperation.InstanceCannotTerminate"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TerminateApmInstance(request *TerminateApmInstanceRequest) (response *TerminateApmInstanceResponse, err error) {
    return c.TerminateApmInstanceWithContext(context.Background(), request)
}

// TerminateApmInstance
// Termination of APM business system.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"
//  FAILEDOPERATION_INSTANCECANNOTTERMINATE = "FailedOperation.InstanceCannotTerminate"
//  FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"
//  FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) TerminateApmInstanceWithContext(ctx context.Context, request *TerminateApmInstanceRequest) (response *TerminateApmInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateApmInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateApmInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateApmInstanceResponse()
    err = c.Send(request, response)
    return
}
