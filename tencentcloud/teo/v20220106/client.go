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

package v20220106

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2022-01-06"

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


func NewCheckCertificateRequest() (request *CheckCertificateRequest) {
    request = &CheckCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "CheckCertificate")
    
    
    return
}

func NewCheckCertificateResponse() (response *CheckCertificateResponse) {
    response = &CheckCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckCertificate
// This API is used to verify a certificate.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CheckCertificate(request *CheckCertificateRequest) (response *CheckCertificateResponse, err error) {
    return c.CheckCertificateWithContext(context.Background(), request)
}

// CheckCertificate
// This API is used to verify a certificate.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CheckCertificateWithContext(ctx context.Context, request *CheckCertificateRequest) (response *CheckCertificateResponse, err error) {
    if request == nil {
        request = NewCheckCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRequest() (request *CreateApplicationProxyRequest) {
    request = &CreateApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxy")
    
    
    return
}

func NewCreateApplicationProxyResponse() (response *CreateApplicationProxyResponse) {
    response = &CreateApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationProxy
// This API is used to create an application proxy.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateApplicationProxy(request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    return c.CreateApplicationProxyWithContext(context.Background(), request)
}

// CreateApplicationProxy
// This API is used to create an application proxy.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateApplicationProxyWithContext(ctx context.Context, request *CreateApplicationProxyRequest) (response *CreateApplicationProxyResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRuleRequest() (request *CreateApplicationProxyRuleRequest) {
    request = &CreateApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxyRule")
    
    
    return
}

func NewCreateApplicationProxyRuleResponse() (response *CreateApplicationProxyRuleResponse) {
    response = &CreateApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationProxyRule
// This API is used to create an application proxy rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) CreateApplicationProxyRule(request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    return c.CreateApplicationProxyRuleWithContext(context.Background(), request)
}

// CreateApplicationProxyRule
// This API is used to create an application proxy rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) CreateApplicationProxyRuleWithContext(ctx context.Context, request *CreateApplicationProxyRuleRequest) (response *CreateApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationProxyRulesRequest() (request *CreateApplicationProxyRulesRequest) {
    request = &CreateApplicationProxyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "CreateApplicationProxyRules")
    
    
    return
}

func NewCreateApplicationProxyRulesResponse() (response *CreateApplicationProxyRulesResponse) {
    response = &CreateApplicationProxyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApplicationProxyRules
// This API is used to batch create application proxy rules.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) CreateApplicationProxyRules(request *CreateApplicationProxyRulesRequest) (response *CreateApplicationProxyRulesResponse, err error) {
    return c.CreateApplicationProxyRulesWithContext(context.Background(), request)
}

// CreateApplicationProxyRules
// This API is used to batch create application proxy rules.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) CreateApplicationProxyRulesWithContext(ctx context.Context, request *CreateApplicationProxyRulesRequest) (response *CreateApplicationProxyRulesResponse, err error) {
    if request == nil {
        request = NewCreateApplicationProxyRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplicationProxyRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationProxyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDnsRecordRequest() (request *CreateDnsRecordRequest) {
    request = &CreateDnsRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "CreateDnsRecord")
    
    
    return
}

func NewCreateDnsRecordResponse() (response *CreateDnsRecordResponse) {
    response = &CreateDnsRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDnsRecord
// This API is used to create a DNS record.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHLBRECORD = "InvalidParameterValue.ConflictWithLBRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDnsRecord(request *CreateDnsRecordRequest) (response *CreateDnsRecordResponse, err error) {
    return c.CreateDnsRecordWithContext(context.Background(), request)
}

// CreateDnsRecord
// This API is used to create a DNS record.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHLBRECORD = "InvalidParameterValue.ConflictWithLBRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDnsRecordWithContext(ctx context.Context, request *CreateDnsRecordRequest) (response *CreateDnsRecordResponse, err error) {
    if request == nil {
        request = NewCreateDnsRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDnsRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDnsRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadBalancingRequest() (request *CreateLoadBalancingRequest) {
    request = &CreateLoadBalancingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "CreateLoadBalancing")
    
    
    return
}

func NewCreateLoadBalancingResponse() (response *CreateLoadBalancingResponse) {
    response = &CreateLoadBalancingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLoadBalancing
// This API is used to create a CLB instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHLBRECORD = "InvalidParameterValue.ConflictWithLBRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateLoadBalancing(request *CreateLoadBalancingRequest) (response *CreateLoadBalancingResponse, err error) {
    return c.CreateLoadBalancingWithContext(context.Background(), request)
}

// CreateLoadBalancing
// This API is used to create a CLB instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_CONFLICTRECORD = "InvalidParameterValue.ConflictRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHLBRECORD = "InvalidParameterValue.ConflictWithLBRecord"
//  INVALIDPARAMETERVALUE_CONFLICTWITHNSRECORD = "InvalidParameterValue.ConflictWithNSRecord"
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_INVALIDDNSNAME = "InvalidParameterValue.InvalidDNSName"
//  INVALIDPARAMETERVALUE_INVALIDPROXYORIGIN = "InvalidParameterValue.InvalidProxyOrigin"
//  INVALIDPARAMETERVALUE_RECORDALREADYEXISTS = "InvalidParameterValue.RecordAlreadyExists"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateLoadBalancingWithContext(ctx context.Context, request *CreateLoadBalancingRequest) (response *CreateLoadBalancingResponse, err error) {
    if request == nil {
        request = NewCreateLoadBalancingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoadBalancing require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoadBalancingResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrefetchTaskRequest() (request *CreatePrefetchTaskRequest) {
    request = &CreatePrefetchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "CreatePrefetchTask")
    
    
    return
}

func NewCreatePrefetchTaskResponse() (response *CreatePrefetchTaskResponse) {
    response = &CreatePrefetchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePrefetchTask
// This API is used to create a pre-warming task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
func (c *Client) CreatePrefetchTask(request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    return c.CreatePrefetchTaskWithContext(context.Background(), request)
}

// CreatePrefetchTask
// This API is used to create a pre-warming task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
func (c *Client) CreatePrefetchTaskWithContext(ctx context.Context, request *CreatePrefetchTaskRequest) (response *CreatePrefetchTaskResponse, err error) {
    if request == nil {
        request = NewCreatePrefetchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrefetchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrefetchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePurgeTaskRequest() (request *CreatePurgeTaskRequest) {
    request = &CreatePurgeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "CreatePurgeTask")
    
    
    return
}

func NewCreatePurgeTaskResponse() (response *CreatePurgeTaskResponse) {
    response = &CreatePurgeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePurgeTask
// This API is used to create a cache purging task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
func (c *Client) CreatePurgeTask(request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    return c.CreatePurgeTaskWithContext(context.Background(), request)
}

// CreatePurgeTask
// This API is used to create a cache purging task.
//
// error code that may be returned:
//  INTERNALERROR_BACKENDERROR = "InternalError.BackendError"
//  INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"
//  INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"
//  INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  INVALIDPARAMETER_TARGET = "InvalidParameter.Target"
//  INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"
//  INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"
//  LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"
//  LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"
func (c *Client) CreatePurgeTaskWithContext(ctx context.Context, request *CreatePurgeTaskRequest) (response *CreatePurgeTaskResponse, err error) {
    if request == nil {
        request = NewCreatePurgeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePurgeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePurgeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneRequest() (request *CreateZoneRequest) {
    request = &CreateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "CreateZone")
    
    
    return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
    response = &CreateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateZone
// This API is used to access a new site.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    return c.CreateZoneWithContext(context.Background(), request)
}

// CreateZone
// This API is used to access a new site.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
func (c *Client) CreateZoneWithContext(ctx context.Context, request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    if request == nil {
        request = NewCreateZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRequest() (request *DeleteApplicationProxyRequest) {
    request = &DeleteApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxy")
    
    
    return
}

func NewDeleteApplicationProxyResponse() (response *DeleteApplicationProxyResponse) {
    response = &DeleteApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplicationProxy
// This API is used to delete an application proxy.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
func (c *Client) DeleteApplicationProxy(request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    return c.DeleteApplicationProxyWithContext(context.Background(), request)
}

// DeleteApplicationProxy
// This API is used to delete an application proxy.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
func (c *Client) DeleteApplicationProxyWithContext(ctx context.Context, request *DeleteApplicationProxyRequest) (response *DeleteApplicationProxyResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApplicationProxyRuleRequest() (request *DeleteApplicationProxyRuleRequest) {
    request = &DeleteApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DeleteApplicationProxyRule")
    
    
    return
}

func NewDeleteApplicationProxyRuleResponse() (response *DeleteApplicationProxyRuleResponse) {
    response = &DeleteApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteApplicationProxyRule
// This API is used to delete an application proxy rule.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
func (c *Client) DeleteApplicationProxyRule(request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    return c.DeleteApplicationProxyRuleWithContext(context.Background(), request)
}

// DeleteApplicationProxyRule
// This API is used to delete an application proxy rule.
//
// error code that may be returned:
//  DRYRUNOPERATION = "DryRunOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINUSE_OTHERS = "ResourceInUse.Others"
func (c *Client) DeleteApplicationProxyRuleWithContext(ctx context.Context, request *DeleteApplicationProxyRuleRequest) (response *DeleteApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewDeleteApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDnsRecordsRequest() (request *DeleteDnsRecordsRequest) {
    request = &DeleteDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DeleteDnsRecords")
    
    
    return
}

func NewDeleteDnsRecordsResponse() (response *DeleteDnsRecordsResponse) {
    response = &DeleteDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDnsRecords
// This API is used to batch delete DNS records.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDnsRecords(request *DeleteDnsRecordsRequest) (response *DeleteDnsRecordsResponse, err error) {
    return c.DeleteDnsRecordsWithContext(context.Background(), request)
}

// DeleteDnsRecords
// This API is used to batch delete DNS records.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDnsRecordsWithContext(ctx context.Context, request *DeleteDnsRecordsRequest) (response *DeleteDnsRecordsResponse, err error) {
    if request == nil {
        request = NewDeleteDnsRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDnsRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLoadBalancingRequest() (request *DeleteLoadBalancingRequest) {
    request = &DeleteLoadBalancingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DeleteLoadBalancing")
    
    
    return
}

func NewDeleteLoadBalancingResponse() (response *DeleteLoadBalancingResponse) {
    response = &DeleteLoadBalancingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLoadBalancing
// This API is used to delete a CLB instance.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLoadBalancing(request *DeleteLoadBalancingRequest) (response *DeleteLoadBalancingResponse, err error) {
    return c.DeleteLoadBalancingWithContext(context.Background(), request)
}

// DeleteLoadBalancing
// This API is used to delete a CLB instance.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteLoadBalancingWithContext(ctx context.Context, request *DeleteLoadBalancingRequest) (response *DeleteLoadBalancingResponse, err error) {
    if request == nil {
        request = NewDeleteLoadBalancingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLoadBalancing require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLoadBalancingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
    request = &DeleteZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DeleteZone")
    
    
    return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
    response = &DeleteZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteZone
// This API is used to delete a site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    return c.DeleteZoneWithContext(context.Background(), request)
}

// DeleteZone
// This API is used to delete a site.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteZoneWithContext(ctx context.Context, request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    if request == nil {
        request = NewDeleteZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationProxyRequest() (request *DescribeApplicationProxyRequest) {
    request = &DescribeApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeApplicationProxy")
    
    
    return
}

func NewDescribeApplicationProxyResponse() (response *DescribeApplicationProxyResponse) {
    response = &DescribeApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationProxy
// This API is used to obtain a list of application proxies.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeApplicationProxy(request *DescribeApplicationProxyRequest) (response *DescribeApplicationProxyResponse, err error) {
    return c.DescribeApplicationProxyWithContext(context.Background(), request)
}

// DescribeApplicationProxy
// This API is used to obtain a list of application proxies.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeApplicationProxyWithContext(ctx context.Context, request *DescribeApplicationProxyRequest) (response *DescribeApplicationProxyResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationProxyDetailRequest() (request *DescribeApplicationProxyDetailRequest) {
    request = &DescribeApplicationProxyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeApplicationProxyDetail")
    
    
    return
}

func NewDescribeApplicationProxyDetailResponse() (response *DescribeApplicationProxyDetailResponse) {
    response = &DescribeApplicationProxyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationProxyDetail
// This API is used to obtain the details of an application proxy.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeApplicationProxyDetail(request *DescribeApplicationProxyDetailRequest) (response *DescribeApplicationProxyDetailResponse, err error) {
    return c.DescribeApplicationProxyDetailWithContext(context.Background(), request)
}

// DescribeApplicationProxyDetail
// This API is used to obtain the details of an application proxy.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeApplicationProxyDetailWithContext(ctx context.Context, request *DescribeApplicationProxyDetailRequest) (response *DescribeApplicationProxyDetailResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationProxyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationProxyDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationProxyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCnameStatusRequest() (request *DescribeCnameStatusRequest) {
    request = &DescribeCnameStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeCnameStatus")
    
    
    return
}

func NewDescribeCnameStatusResponse() (response *DescribeCnameStatusResponse) {
    response = &DescribeCnameStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCnameStatus
// This API is used to query the CNAME status of a domain name.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCnameStatus(request *DescribeCnameStatusRequest) (response *DescribeCnameStatusResponse, err error) {
    return c.DescribeCnameStatusWithContext(context.Background(), request)
}

// DescribeCnameStatus
// This API is used to query the CNAME status of a domain name.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCnameStatusWithContext(ctx context.Context, request *DescribeCnameStatusRequest) (response *DescribeCnameStatusResponse, err error) {
    if request == nil {
        request = NewDescribeCnameStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCnameStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCnameStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDefaultCertificatesRequest() (request *DescribeDefaultCertificatesRequest) {
    request = &DescribeDefaultCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDefaultCertificates")
    
    
    return
}

func NewDescribeDefaultCertificatesResponse() (response *DescribeDefaultCertificatesResponse) {
    response = &DescribeDefaultCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDefaultCertificates
// This API is used to query a list of default certificates.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DescribeDefaultCertificates(request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    return c.DescribeDefaultCertificatesWithContext(context.Background(), request)
}

// DescribeDefaultCertificates
// This API is used to query a list of default certificates.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_ZONENOTFOUND = "ResourceUnavailable.ZoneNotFound"
func (c *Client) DescribeDefaultCertificatesWithContext(ctx context.Context, request *DescribeDefaultCertificatesRequest) (response *DescribeDefaultCertificatesResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultCertificatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDnsDataRequest() (request *DescribeDnsDataRequest) {
    request = &DescribeDnsDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDnsData")
    
    
    return
}

func NewDescribeDnsDataResponse() (response *DescribeDnsDataResponse) {
    response = &DescribeDnsDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDnsData
// This API is used to obtain collected DNS requests.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDnsData(request *DescribeDnsDataRequest) (response *DescribeDnsDataResponse, err error) {
    return c.DescribeDnsDataWithContext(context.Background(), request)
}

// DescribeDnsData
// This API is used to obtain collected DNS requests.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDnsDataWithContext(ctx context.Context, request *DescribeDnsDataRequest) (response *DescribeDnsDataResponse, err error) {
    if request == nil {
        request = NewDescribeDnsDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDnsData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDnsDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDnsRecordsRequest() (request *DescribeDnsRecordsRequest) {
    request = &DescribeDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDnsRecords")
    
    
    return
}

func NewDescribeDnsRecordsResponse() (response *DescribeDnsRecordsResponse) {
    response = &DescribeDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDnsRecords
// This API is used to query DNS records. Paging, sorting and filtering are supported.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnsRecords(request *DescribeDnsRecordsRequest) (response *DescribeDnsRecordsResponse, err error) {
    return c.DescribeDnsRecordsWithContext(context.Background(), request)
}

// DescribeDnsRecords
// This API is used to query DNS records. Paging, sorting and filtering are supported.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnsRecordsWithContext(ctx context.Context, request *DescribeDnsRecordsRequest) (response *DescribeDnsRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDnsRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDnsRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDnssecRequest() (request *DescribeDnssecRequest) {
    request = &DescribeDnssecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeDnssec")
    
    
    return
}

func NewDescribeDnssecResponse() (response *DescribeDnssecResponse) {
    response = &DescribeDnssecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDnssec
// This API is used to query DNSSEC information.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnssec(request *DescribeDnssecRequest) (response *DescribeDnssecResponse, err error) {
    return c.DescribeDnssecWithContext(context.Background(), request)
}

// DescribeDnssec
// This API is used to query DNSSEC information.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeDnssecWithContext(ctx context.Context, request *DescribeDnssecRequest) (response *DescribeDnssecResponse, err error) {
    if request == nil {
        request = NewDescribeDnssecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDnssec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDnssecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsCertificateRequest() (request *DescribeHostsCertificateRequest) {
    request = &DescribeHostsCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeHostsCertificate")
    
    
    return
}

func NewDescribeHostsCertificateResponse() (response *DescribeHostsCertificateResponse) {
    response = &DescribeHostsCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHostsCertificate
// This API is used to query certificates of domain names. Paging, sorting and filtering are supported.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeHostsCertificate(request *DescribeHostsCertificateRequest) (response *DescribeHostsCertificateResponse, err error) {
    return c.DescribeHostsCertificateWithContext(context.Background(), request)
}

// DescribeHostsCertificate
// This API is used to query certificates of domain names. Paging, sorting and filtering are supported.
//
// error code that may be returned:
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeHostsCertificateWithContext(ctx context.Context, request *DescribeHostsCertificateRequest) (response *DescribeHostsCertificateResponse, err error) {
    if request == nil {
        request = NewDescribeHostsCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostsCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostsCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostsSettingRequest() (request *DescribeHostsSettingRequest) {
    request = &DescribeHostsSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeHostsSetting")
    
    
    return
}

func NewDescribeHostsSettingResponse() (response *DescribeHostsSettingResponse) {
    response = &DescribeHostsSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHostsSetting
// This API is used to query detailed domain name configuration.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeHostsSetting(request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    return c.DescribeHostsSettingWithContext(context.Background(), request)
}

// DescribeHostsSetting
// This API is used to query detailed domain name configuration.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeHostsSettingWithContext(ctx context.Context, request *DescribeHostsSettingRequest) (response *DescribeHostsSettingResponse, err error) {
    if request == nil {
        request = NewDescribeHostsSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostsSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostsSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdentificationRequest() (request *DescribeIdentificationRequest) {
    request = &DescribeIdentificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeIdentification")
    
    
    return
}

func NewDescribeIdentificationResponse() (response *DescribeIdentificationResponse) {
    response = &DescribeIdentificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIdentification
// This API is used to query verification results.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIdentification(request *DescribeIdentificationRequest) (response *DescribeIdentificationResponse, err error) {
    return c.DescribeIdentificationWithContext(context.Background(), request)
}

// DescribeIdentification
// This API is used to query verification results.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeIdentificationWithContext(ctx context.Context, request *DescribeIdentificationRequest) (response *DescribeIdentificationResponse, err error) {
    if request == nil {
        request = NewDescribeIdentificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdentification require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdentificationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancingRequest() (request *DescribeLoadBalancingRequest) {
    request = &DescribeLoadBalancingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLoadBalancing")
    
    
    return
}

func NewDescribeLoadBalancingResponse() (response *DescribeLoadBalancingResponse) {
    response = &DescribeLoadBalancingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalancing
// This API is used to obtain a list of CLB instances.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeLoadBalancing(request *DescribeLoadBalancingRequest) (response *DescribeLoadBalancingResponse, err error) {
    return c.DescribeLoadBalancingWithContext(context.Background(), request)
}

// DescribeLoadBalancing
// This API is used to obtain a list of CLB instances.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeLoadBalancingWithContext(ctx context.Context, request *DescribeLoadBalancingRequest) (response *DescribeLoadBalancingResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancing require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadBalancingDetailRequest() (request *DescribeLoadBalancingDetailRequest) {
    request = &DescribeLoadBalancingDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeLoadBalancingDetail")
    
    
    return
}

func NewDescribeLoadBalancingDetailResponse() (response *DescribeLoadBalancingDetailResponse) {
    response = &DescribeLoadBalancingDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLoadBalancingDetail
// This API is used to query the details of a CLB instance.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLoadBalancingDetail(request *DescribeLoadBalancingDetailRequest) (response *DescribeLoadBalancingDetailResponse, err error) {
    return c.DescribeLoadBalancingDetailWithContext(context.Background(), request)
}

// DescribeLoadBalancingDetail
// This API is used to query the details of a CLB instance.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLoadBalancingDetailWithContext(ctx context.Context, request *DescribeLoadBalancingDetailRequest) (response *DescribeLoadBalancingDetailResponse, err error) {
    if request == nil {
        request = NewDescribeLoadBalancingDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadBalancingDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadBalancingDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrefetchTasksRequest() (request *DescribePrefetchTasksRequest) {
    request = &DescribePrefetchTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribePrefetchTasks")
    
    
    return
}

func NewDescribePrefetchTasksResponse() (response *DescribePrefetchTasksResponse) {
    response = &DescribePrefetchTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePrefetchTasks
// This API is used to query the pre-warming task status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasks(request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    return c.DescribePrefetchTasksWithContext(context.Background(), request)
}

// DescribePrefetchTasks
// This API is used to query the pre-warming task status.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
//  UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"
func (c *Client) DescribePrefetchTasksWithContext(ctx context.Context, request *DescribePrefetchTasksRequest) (response *DescribePrefetchTasksResponse, err error) {
    if request == nil {
        request = NewDescribePrefetchTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePrefetchTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePrefetchTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribePurgeTasks")
    
    
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeTasks
// This API is used to query the cache purging history.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// This API is used to query the cache purging history.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) DescribePurgeTasksWithContext(ctx context.Context, request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneDetailsRequest() (request *DescribeZoneDetailsRequest) {
    request = &DescribeZoneDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneDetails")
    
    
    return
}

func NewDescribeZoneDetailsResponse() (response *DescribeZoneDetailsResponse) {
    response = &DescribeZoneDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneDetails
// This API is used to query the details of a site by site ID.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeZoneDetails(request *DescribeZoneDetailsRequest) (response *DescribeZoneDetailsResponse, err error) {
    return c.DescribeZoneDetailsWithContext(context.Background(), request)
}

// DescribeZoneDetails
// This API is used to query the details of a site by site ID.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) DescribeZoneDetailsWithContext(ctx context.Context, request *DescribeZoneDetailsRequest) (response *DescribeZoneDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeZoneDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZoneSettingRequest() (request *DescribeZoneSettingRequest) {
    request = &DescribeZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZoneSetting")
    
    
    return
}

func NewDescribeZoneSettingResponse() (response *DescribeZoneSettingResponse) {
    response = &DescribeZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZoneSetting
// This API is used to query the site configuration.
//
// error code that may be returned:
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeZoneSetting(request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    return c.DescribeZoneSettingWithContext(context.Background(), request)
}

// DescribeZoneSetting
// This API is used to query the site configuration.
//
// error code that may be returned:
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeZoneSettingWithContext(ctx context.Context, request *DescribeZoneSettingRequest) (response *DescribeZoneSettingResponse, err error) {
    if request == nil {
        request = NewDescribeZoneSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZones
// This API is used to query the list of user sites.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// This API is used to query the list of user sites.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadL7LogsRequest() (request *DownloadL7LogsRequest) {
    request = &DownloadL7LogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "DownloadL7Logs")
    
    
    return
}

func NewDownloadL7LogsResponse() (response *DownloadL7LogsResponse) {
    response = &DownloadL7LogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadL7Logs
// This API is used to query layer-7 logs.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DownloadL7Logs(request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    return c.DownloadL7LogsWithContext(context.Background(), request)
}

// DownloadL7Logs
// This API is used to query layer-7 logs.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DownloadL7LogsWithContext(ctx context.Context, request *DownloadL7LogsRequest) (response *DownloadL7LogsResponse, err error) {
    if request == nil {
        request = NewDownloadL7LogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadL7Logs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadL7LogsResponse()
    err = c.Send(request, response)
    return
}

func NewIdentifyZoneRequest() (request *IdentifyZoneRequest) {
    request = &IdentifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "IdentifyZone")
    
    
    return
}

func NewIdentifyZoneResponse() (response *IdentifyZoneResponse) {
    response = &IdentifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IdentifyZone
// This API is used to verify ownership of the site.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZone(request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    return c.IdentifyZoneWithContext(context.Background(), request)
}

// IdentifyZone
// This API is used to verify ownership of the site.
//
// error code that may be returned:
//  UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"
func (c *Client) IdentifyZoneWithContext(ctx context.Context, request *IdentifyZoneRequest) (response *IdentifyZoneResponse, err error) {
    if request == nil {
        request = NewIdentifyZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IdentifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewIdentifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewImportDnsRecordsRequest() (request *ImportDnsRecordsRequest) {
    request = &ImportDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ImportDnsRecords")
    
    
    return
}

func NewImportDnsRecordsResponse() (response *ImportDnsRecordsResponse) {
    response = &ImportDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportDnsRecords
// This API is used to import DNS records.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ImportDnsRecords(request *ImportDnsRecordsRequest) (response *ImportDnsRecordsResponse, err error) {
    return c.ImportDnsRecordsWithContext(context.Background(), request)
}

// ImportDnsRecords
// This API is used to import DNS records.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ImportDnsRecordsWithContext(ctx context.Context, request *ImportDnsRecordsRequest) (response *ImportDnsRecordsResponse, err error) {
    if request == nil {
        request = NewImportDnsRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportDnsRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRequest() (request *ModifyApplicationProxyRequest) {
    request = &ModifyApplicationProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxy")
    
    
    return
}

func NewModifyApplicationProxyResponse() (response *ModifyApplicationProxyResponse) {
    response = &ModifyApplicationProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxy
// This API is used to modify an application proxy.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyApplicationProxy(request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    return c.ModifyApplicationProxyWithContext(context.Background(), request)
}

// ModifyApplicationProxy
// This API is used to modify an application proxy.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyApplicationProxyWithContext(ctx context.Context, request *ModifyApplicationProxyRequest) (response *ModifyApplicationProxyResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleRequest() (request *ModifyApplicationProxyRuleRequest) {
    request = &ModifyApplicationProxyRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRule")
    
    
    return
}

func NewModifyApplicationProxyRuleResponse() (response *ModifyApplicationProxyRuleResponse) {
    response = &ModifyApplicationProxyRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxyRule
// This API is used to modify an application proxy rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyRule(request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    return c.ModifyApplicationProxyRuleWithContext(context.Background(), request)
}

// ModifyApplicationProxyRule
// This API is used to modify an application proxy rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyRuleWithContext(ctx context.Context, request *ModifyApplicationProxyRuleRequest) (response *ModifyApplicationProxyRuleResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyRuleStatusRequest() (request *ModifyApplicationProxyRuleStatusRequest) {
    request = &ModifyApplicationProxyRuleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyRuleStatus")
    
    
    return
}

func NewModifyApplicationProxyRuleStatusResponse() (response *ModifyApplicationProxyRuleStatusResponse) {
    response = &ModifyApplicationProxyRuleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxyRuleStatus
// This API is used to modify the status of an application proxy rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyRuleStatus(request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    return c.ModifyApplicationProxyRuleStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyRuleStatus
// This API is used to modify the status of an application proxy rule.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyRuleStatusWithContext(ctx context.Context, request *ModifyApplicationProxyRuleStatusRequest) (response *ModifyApplicationProxyRuleStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyRuleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyRuleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyRuleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationProxyStatusRequest() (request *ModifyApplicationProxyStatusRequest) {
    request = &ModifyApplicationProxyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyApplicationProxyStatus")
    
    
    return
}

func NewModifyApplicationProxyStatusResponse() (response *ModifyApplicationProxyStatusResponse) {
    response = &ModifyApplicationProxyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyApplicationProxyStatus
// This API is used to modify the status of an application proxy.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyStatus(request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    return c.ModifyApplicationProxyStatusWithContext(context.Background(), request)
}

// ModifyApplicationProxyStatus
// This API is used to modify the status of an application proxy.
//
// error code that may be returned:
//  INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"
func (c *Client) ModifyApplicationProxyStatusWithContext(ctx context.Context, request *ModifyApplicationProxyStatusRequest) (response *ModifyApplicationProxyStatusResponse, err error) {
    if request == nil {
        request = NewModifyApplicationProxyStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationProxyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationProxyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDefaultCertificateRequest() (request *ModifyDefaultCertificateRequest) {
    request = &ModifyDefaultCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDefaultCertificate")
    
    
    return
}

func NewModifyDefaultCertificateResponse() (response *ModifyDefaultCertificateResponse) {
    response = &ModifyDefaultCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDefaultCertificate
// This API is used to modify the status of a default certificate.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
func (c *Client) ModifyDefaultCertificate(request *ModifyDefaultCertificateRequest) (response *ModifyDefaultCertificateResponse, err error) {
    return c.ModifyDefaultCertificateWithContext(context.Background(), request)
}

// ModifyDefaultCertificate
// This API is used to modify the status of a default certificate.
//
// error code that may be returned:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_CERTNOTFOUND = "ResourceUnavailable.CertNotFound"
func (c *Client) ModifyDefaultCertificateWithContext(ctx context.Context, request *ModifyDefaultCertificateRequest) (response *ModifyDefaultCertificateResponse, err error) {
    if request == nil {
        request = NewModifyDefaultCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDefaultCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDefaultCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDnsRecordRequest() (request *ModifyDnsRecordRequest) {
    request = &ModifyDnsRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDnsRecord")
    
    
    return
}

func NewModifyDnsRecordResponse() (response *ModifyDnsRecordResponse) {
    response = &ModifyDnsRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDnsRecord
// This API is used to modify DNS records.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDnsRecord(request *ModifyDnsRecordRequest) (response *ModifyDnsRecordResponse, err error) {
    return c.ModifyDnsRecordWithContext(context.Background(), request)
}

// ModifyDnsRecord
// This API is used to modify DNS records.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDDNSCONTENT = "InvalidParameterValue.InvalidDNSContent"
//  INVALIDPARAMETERVALUE_RECORDNOTALLOWED = "InvalidParameterValue.RecordNotAllowed"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDnsRecordWithContext(ctx context.Context, request *ModifyDnsRecordRequest) (response *ModifyDnsRecordResponse, err error) {
    if request == nil {
        request = NewModifyDnsRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDnsRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDnsRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDnssecRequest() (request *ModifyDnssecRequest) {
    request = &ModifyDnssecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyDnssec")
    
    
    return
}

func NewModifyDnssecResponse() (response *ModifyDnssecResponse) {
    response = &ModifyDnssecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDnssec
// This API is used to modify the DNSSEC status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDnssec(request *ModifyDnssecRequest) (response *ModifyDnssecResponse, err error) {
    return c.ModifyDnssecWithContext(context.Background(), request)
}

// ModifyDnssec
// This API is used to modify the DNSSEC status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDnssecWithContext(ctx context.Context, request *ModifyDnssecRequest) (response *ModifyDnssecResponse, err error) {
    if request == nil {
        request = NewModifyDnssecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDnssec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDnssecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyHostsCertificateRequest() (request *ModifyHostsCertificateRequest) {
    request = &ModifyHostsCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyHostsCertificate")
    
    
    return
}

func NewModifyHostsCertificateResponse() (response *ModifyHostsCertificateResponse) {
    response = &ModifyHostsCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyHostsCertificate
// This API is used to modify the certificate of a domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
func (c *Client) ModifyHostsCertificate(request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    return c.ModifyHostsCertificateWithContext(context.Background(), request)
}

// ModifyHostsCertificate
// This API is used to modify the certificate of a domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
func (c *Client) ModifyHostsCertificateWithContext(ctx context.Context, request *ModifyHostsCertificateRequest) (response *ModifyHostsCertificateResponse, err error) {
    if request == nil {
        request = NewModifyHostsCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyHostsCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyHostsCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancingRequest() (request *ModifyLoadBalancingRequest) {
    request = &ModifyLoadBalancingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyLoadBalancing")
    
    
    return
}

func NewModifyLoadBalancingResponse() (response *ModifyLoadBalancingResponse) {
    response = &ModifyLoadBalancingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoadBalancing
// This API is used to modify a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
func (c *Client) ModifyLoadBalancing(request *ModifyLoadBalancingRequest) (response *ModifyLoadBalancingResponse, err error) {
    return c.ModifyLoadBalancingWithContext(context.Background(), request)
}

// ModifyLoadBalancing
// This API is used to modify a CLB instance.
//
// error code that may be returned:
//  FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDCERTINFO = "InvalidParameter.InvalidCertInfo"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE_HOSTNOTFOUND = "ResourceUnavailable.HostNotFound"
func (c *Client) ModifyLoadBalancingWithContext(ctx context.Context, request *ModifyLoadBalancingRequest) (response *ModifyLoadBalancingResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancing require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancingStatusRequest() (request *ModifyLoadBalancingStatusRequest) {
    request = &ModifyLoadBalancingStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyLoadBalancingStatus")
    
    
    return
}

func NewModifyLoadBalancingStatusResponse() (response *ModifyLoadBalancingStatusResponse) {
    response = &ModifyLoadBalancingStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLoadBalancingStatus
// This API is used to modify the status of a CLB instance.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyLoadBalancingStatus(request *ModifyLoadBalancingStatusRequest) (response *ModifyLoadBalancingStatusResponse, err error) {
    return c.ModifyLoadBalancingStatusWithContext(context.Background(), request)
}

// ModifyLoadBalancingStatus
// This API is used to modify the status of a CLB instance.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyLoadBalancingStatusWithContext(ctx context.Context, request *ModifyLoadBalancingStatusRequest) (response *ModifyLoadBalancingStatusResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancingStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLoadBalancingStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLoadBalancingStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneRequest() (request *ModifyZoneRequest) {
    request = &ModifyZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZone")
    
    
    return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
    response = &ModifyZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZone
// This API is used to modify the site information.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyZone(request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    return c.ModifyZoneWithContext(context.Background(), request)
}

// ModifyZone
// This API is used to modify the site information.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyZoneWithContext(ctx context.Context, request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
    if request == nil {
        request = NewModifyZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneCnameSpeedUpRequest() (request *ModifyZoneCnameSpeedUpRequest) {
    request = &ModifyZoneCnameSpeedUpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneCnameSpeedUp")
    
    
    return
}

func NewModifyZoneCnameSpeedUpResponse() (response *ModifyZoneCnameSpeedUpResponse) {
    response = &ModifyZoneCnameSpeedUpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZoneCnameSpeedUp
// This API is used to modify the CNAME acceleration status.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyZoneCnameSpeedUp(request *ModifyZoneCnameSpeedUpRequest) (response *ModifyZoneCnameSpeedUpResponse, err error) {
    return c.ModifyZoneCnameSpeedUpWithContext(context.Background(), request)
}

// ModifyZoneCnameSpeedUp
// This API is used to modify the CNAME acceleration status.
//
// error code that may be returned:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyZoneCnameSpeedUpWithContext(ctx context.Context, request *ModifyZoneCnameSpeedUpRequest) (response *ModifyZoneCnameSpeedUpResponse, err error) {
    if request == nil {
        request = NewModifyZoneCnameSpeedUpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneCnameSpeedUp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneCnameSpeedUpResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneSettingRequest() (request *ModifyZoneSettingRequest) {
    request = &ModifyZoneSettingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneSetting")
    
    
    return
}

func NewModifyZoneSettingResponse() (response *ModifyZoneSettingResponse) {
    response = &ModifyZoneSettingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZoneSetting
// This API is used to modify the site configuration.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSetting(request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    return c.ModifyZoneSettingWithContext(context.Background(), request)
}

// ModifyZoneSetting
// This API is used to modify the site configuration.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDREQUESTHEADERNAME = "InvalidParameter.InvalidRequestHeaderName"
//  INVALIDPARAMETER_SETTINGINVALIDPARAM = "InvalidParameter.SettingInvalidParam"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyZoneSettingWithContext(ctx context.Context, request *ModifyZoneSettingRequest) (response *ModifyZoneSettingResponse, err error) {
    if request == nil {
        request = NewModifyZoneSettingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneSetting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneSettingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyZoneStatusRequest() (request *ModifyZoneStatusRequest) {
    request = &ModifyZoneStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ModifyZoneStatus")
    
    
    return
}

func NewModifyZoneStatusResponse() (response *ModifyZoneStatusResponse) {
    response = &ModifyZoneStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyZoneStatus
// This API is used to change the site status.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyZoneStatus(request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    return c.ModifyZoneStatusWithContext(context.Background(), request)
}

// ModifyZoneStatus
// This API is used to change the site status.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyZoneStatusWithContext(ctx context.Context, request *ModifyZoneStatusRequest) (response *ModifyZoneStatusResponse, err error) {
    if request == nil {
        request = NewModifyZoneStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyZoneStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyZoneStatusResponse()
    err = c.Send(request, response)
    return
}

func NewReclaimZoneRequest() (request *ReclaimZoneRequest) {
    request = &ReclaimZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ReclaimZone")
    
    
    return
}

func NewReclaimZoneResponse() (response *ReclaimZoneResponse) {
    response = &ReclaimZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReclaimZone
// This API is used to reclaim a site from other users after its ownership is verified.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReclaimZone(request *ReclaimZoneRequest) (response *ReclaimZoneResponse, err error) {
    return c.ReclaimZoneWithContext(context.Background(), request)
}

// ReclaimZone
// This API is used to reclaim a site from other users after its ownership is verified.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReclaimZoneWithContext(ctx context.Context, request *ReclaimZoneRequest) (response *ReclaimZoneResponse, err error) {
    if request == nil {
        request = NewReclaimZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReclaimZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewReclaimZoneResponse()
    err = c.Send(request, response)
    return
}

func NewScanDnsRecordsRequest() (request *ScanDnsRecordsRequest) {
    request = &ScanDnsRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("teo", APIVersion, "ScanDnsRecords")
    
    
    return
}

func NewScanDnsRecordsResponse() (response *ScanDnsRecordsResponse) {
    response = &ScanDnsRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanDnsRecords
// This API is used to scan resolution records.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ScanDnsRecords(request *ScanDnsRecordsRequest) (response *ScanDnsRecordsResponse, err error) {
    return c.ScanDnsRecordsWithContext(context.Background(), request)
}

// ScanDnsRecords
// This API is used to scan resolution records.
//
// error code that may be returned:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ScanDnsRecordsWithContext(ctx context.Context, request *ScanDnsRecordsRequest) (response *ScanDnsRecordsResponse, err error) {
    if request == nil {
        request = NewScanDnsRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScanDnsRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewScanDnsRecordsResponse()
    err = c.Send(request, response)
    return
}
