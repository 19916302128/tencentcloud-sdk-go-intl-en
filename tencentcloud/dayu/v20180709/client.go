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

package v20180709

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-07-09"

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


func NewCreateBasicDDoSAlarmThresholdRequest() (request *CreateBasicDDoSAlarmThresholdRequest) {
    request = &CreateBasicDDoSAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateBasicDDoSAlarmThreshold")
    return
}

func NewCreateBasicDDoSAlarmThresholdResponse() (response *CreateBasicDDoSAlarmThresholdResponse) {
    response = &CreateBasicDDoSAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to set the DDoS alarm threshold for Anti-DDoS Basic, which is only supported for Anti-DDoS Basic.
func (c *Client) CreateBasicDDoSAlarmThreshold(request *CreateBasicDDoSAlarmThresholdRequest) (response *CreateBasicDDoSAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewCreateBasicDDoSAlarmThresholdRequest()
    }
    response = NewCreateBasicDDoSAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBoundIPRequest() (request *CreateBoundIPRequest) {
    request = &CreateBoundIPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateBoundIP")
    return
}

func NewCreateBoundIPResponse() (response *CreateBoundIPResponse) {
    response = &CreateBoundIPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to bind an IP to an Anti-DDoS Pro instance, which supports both single IP instances and multi-IP instances. It should be noted that this API is async; therefore, if a binding/unbinding operation is in progress, new binding/unbinding operations cannot be initiated.
func (c *Client) CreateBoundIP(request *CreateBoundIPRequest) (response *CreateBoundIPResponse, err error) {
    if request == nil {
        request = NewCreateBoundIPRequest()
    }
    response = NewCreateBoundIPResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCCFrequencyRulesRequest() (request *CreateCCFrequencyRulesRequest) {
    request = &CreateCCFrequencyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateCCFrequencyRules")
    return
}

func NewCreateCCFrequencyRulesResponse() (response *CreateCCFrequencyRulesResponse) {
    response = &CreateCCFrequencyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add an access frequency control rule for CC protection.
func (c *Client) CreateCCFrequencyRules(request *CreateCCFrequencyRulesRequest) (response *CreateCCFrequencyRulesResponse, err error) {
    if request == nil {
        request = NewCreateCCFrequencyRulesRequest()
    }
    response = NewCreateCCFrequencyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCCSelfDefinePolicyRequest() (request *CreateCCSelfDefinePolicyRequest) {
    request = &CreateCCSelfDefinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateCCSelfDefinePolicy")
    return
}

func NewCreateCCSelfDefinePolicyResponse() (response *CreateCCSelfDefinePolicyResponse) {
    response = &CreateCCSelfDefinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a custom CC policy.
func (c *Client) CreateCCSelfDefinePolicy(request *CreateCCSelfDefinePolicyRequest) (response *CreateCCSelfDefinePolicyResponse, err error) {
    if request == nil {
        request = NewCreateCCSelfDefinePolicyRequest()
    }
    response = NewCreateCCSelfDefinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSPolicyRequest() (request *CreateDDoSPolicyRequest) {
    request = &CreateDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateDDoSPolicy")
    return
}

func NewCreateDDoSPolicyResponse() (response *CreateDDoSPolicyResponse) {
    response = &CreateDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add an advanced DDoS policy.
func (c *Client) CreateDDoSPolicy(request *CreateDDoSPolicyRequest) (response *CreateDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewCreateDDoSPolicyRequest()
    }
    response = NewCreateDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDDoSPolicyCaseRequest() (request *CreateDDoSPolicyCaseRequest) {
    request = &CreateDDoSPolicyCaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateDDoSPolicyCase")
    return
}

func NewCreateDDoSPolicyCaseResponse() (response *CreateDDoSPolicyCaseResponse) {
    response = &CreateDDoSPolicyCaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add a policy scenario.
func (c *Client) CreateDDoSPolicyCase(request *CreateDDoSPolicyCaseRequest) (response *CreateDDoSPolicyCaseResponse, err error) {
    if request == nil {
        request = NewCreateDDoSPolicyCaseRequest()
    }
    response = NewCreateDDoSPolicyCaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceNameRequest() (request *CreateInstanceNameRequest) {
    request = &CreateInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateInstanceName")
    return
}

func NewCreateInstanceNameResponse() (response *CreateInstanceNameResponse) {
    response = &CreateInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to rename a resource instance, which supports single IP instances, multi-IP instances, Anti-DDoS Advanced, and Anti-DDoS Ultimate.
func (c *Client) CreateInstanceName(request *CreateInstanceNameRequest) (response *CreateInstanceNameResponse, err error) {
    if request == nil {
        request = NewCreateInstanceNameRequest()
    }
    response = NewCreateInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL4HealthConfigRequest() (request *CreateL4HealthConfigRequest) {
    request = &CreateL4HealthConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL4HealthConfig")
    return
}

func NewCreateL4HealthConfigResponse() (response *CreateL4HealthConfigResponse) {
    response = &CreateL4HealthConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to upload layer-4 health check configuration.
func (c *Client) CreateL4HealthConfig(request *CreateL4HealthConfigRequest) (response *CreateL4HealthConfigResponse, err error) {
    if request == nil {
        request = NewCreateL4HealthConfigRequest()
    }
    response = NewCreateL4HealthConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL4RulesRequest() (request *CreateL4RulesRequest) {
    request = &CreateL4RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL4Rules")
    return
}

func NewCreateL4RulesResponse() (response *CreateL4RulesResponse) {
    response = &CreateL4RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add a layer-4 forwarding rule.
func (c *Client) CreateL4Rules(request *CreateL4RulesRequest) (response *CreateL4RulesResponse, err error) {
    if request == nil {
        request = NewCreateL4RulesRequest()
    }
    response = NewCreateL4RulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7CCRuleRequest() (request *CreateL7CCRuleRequest) {
    request = &CreateL7CCRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7CCRule")
    return
}

func NewCreateL7CCRuleResponse() (response *CreateL7CCRuleResponse) {
    response = &CreateL7CCRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add a custom frequency control rule for layer-7 CC access (it works in the IP + Host dimension and does not support specific URIs). As it has been disused, please call the new `CreateCCFrequencyRules` API, which supports both IP + Host and URI.
func (c *Client) CreateL7CCRule(request *CreateL7CCRuleRequest) (response *CreateL7CCRuleResponse, err error) {
    if request == nil {
        request = NewCreateL7CCRuleRequest()
    }
    response = NewCreateL7CCRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7HealthConfigRequest() (request *CreateL7HealthConfigRequest) {
    request = &CreateL7HealthConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7HealthConfig")
    return
}

func NewCreateL7HealthConfigResponse() (response *CreateL7HealthConfigResponse) {
    response = &CreateL7HealthConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to upload layer-7 health check configuration.
func (c *Client) CreateL7HealthConfig(request *CreateL7HealthConfigRequest) (response *CreateL7HealthConfigResponse, err error) {
    if request == nil {
        request = NewCreateL7HealthConfigRequest()
    }
    response = NewCreateL7HealthConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RuleCertRequest() (request *CreateL7RuleCertRequest) {
    request = &CreateL7RuleCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7RuleCert")
    return
}

func NewCreateL7RuleCertResponse() (response *CreateL7RuleCertResponse) {
    response = &CreateL7RuleCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to configure a certificate for a layer-7 forwarding rule.
func (c *Client) CreateL7RuleCert(request *CreateL7RuleCertRequest) (response *CreateL7RuleCertResponse, err error) {
    if request == nil {
        request = NewCreateL7RuleCertRequest()
    }
    response = NewCreateL7RuleCertResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RulesRequest() (request *CreateL7RulesRequest) {
    request = &CreateL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7Rules")
    return
}

func NewCreateL7RulesResponse() (response *CreateL7RulesResponse) {
    response = &CreateL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add a layer-7 (website) forwarding rule.
func (c *Client) CreateL7Rules(request *CreateL7RulesRequest) (response *CreateL7RulesResponse, err error) {
    if request == nil {
        request = NewCreateL7RulesRequest()
    }
    response = NewCreateL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateL7RulesUploadRequest() (request *CreateL7RulesUploadRequest) {
    request = &CreateL7RulesUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateL7RulesUpload")
    return
}

func NewCreateL7RulesUploadResponse() (response *CreateL7RulesUploadResponse) {
    response = &CreateL7RulesUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to upload layer-7 forwarding rules in batches.
func (c *Client) CreateL7RulesUpload(request *CreateL7RulesUploadRequest) (response *CreateL7RulesUploadResponse, err error) {
    if request == nil {
        request = NewCreateL7RulesUploadRequest()
    }
    response = NewCreateL7RulesUploadResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNetReturnRequest() (request *CreateNetReturnRequest) {
    request = &CreateNetReturnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateNetReturn")
    return
}

func NewCreateNetReturnResponse() (response *CreateNetReturnResponse) {
    response = &CreateNetReturnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to switch to the real server in Anti-DDoS Ultimate.
func (c *Client) CreateNetReturn(request *CreateNetReturnRequest) (response *CreateNetReturnResponse, err error) {
    if request == nil {
        request = NewCreateNetReturnRequest()
    }
    response = NewCreateNetReturnResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNewL7RulesUploadRequest() (request *CreateNewL7RulesUploadRequest) {
    request = &CreateNewL7RulesUploadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateNewL7RulesUpload")
    return
}

func NewCreateNewL7RulesUploadResponse() (response *CreateNewL7RulesUploadResponse) {
    response = &CreateNewL7RulesUploadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to batch upload Layer-7 forwarding rules.
func (c *Client) CreateNewL7RulesUpload(request *CreateNewL7RulesUploadRequest) (response *CreateNewL7RulesUploadResponse, err error) {
    if request == nil {
        request = NewCreateNewL7RulesUploadRequest()
    }
    response = NewCreateNewL7RulesUploadResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUnblockIpRequest() (request *CreateUnblockIpRequest) {
    request = &CreateUnblockIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "CreateUnblockIp")
    return
}

func NewCreateUnblockIpResponse() (response *CreateUnblockIpResponse) {
    response = &CreateUnblockIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to unblock an IP.
func (c *Client) CreateUnblockIp(request *CreateUnblockIpRequest) (response *CreateUnblockIpResponse, err error) {
    if request == nil {
        request = NewCreateUnblockIpRequest()
    }
    response = NewCreateUnblockIpResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCFrequencyRulesRequest() (request *DeleteCCFrequencyRulesRequest) {
    request = &DeleteCCFrequencyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteCCFrequencyRules")
    return
}

func NewDeleteCCFrequencyRulesResponse() (response *DeleteCCFrequencyRulesResponse) {
    response = &DeleteCCFrequencyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete an access frequency control rule for CC protection.
func (c *Client) DeleteCCFrequencyRules(request *DeleteCCFrequencyRulesRequest) (response *DeleteCCFrequencyRulesResponse, err error) {
    if request == nil {
        request = NewDeleteCCFrequencyRulesRequest()
    }
    response = NewDeleteCCFrequencyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCCSelfDefinePolicyRequest() (request *DeleteCCSelfDefinePolicyRequest) {
    request = &DeleteCCSelfDefinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteCCSelfDefinePolicy")
    return
}

func NewDeleteCCSelfDefinePolicyResponse() (response *DeleteCCSelfDefinePolicyResponse) {
    response = &DeleteCCSelfDefinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a custom CC policy.
func (c *Client) DeleteCCSelfDefinePolicy(request *DeleteCCSelfDefinePolicyRequest) (response *DeleteCCSelfDefinePolicyResponse, err error) {
    if request == nil {
        request = NewDeleteCCSelfDefinePolicyRequest()
    }
    response = NewDeleteCCSelfDefinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSPolicyRequest() (request *DeleteDDoSPolicyRequest) {
    request = &DeleteDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteDDoSPolicy")
    return
}

func NewDeleteDDoSPolicyResponse() (response *DeleteDDoSPolicyResponse) {
    response = &DeleteDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete an advanced DDoS protection policy.
func (c *Client) DeleteDDoSPolicy(request *DeleteDDoSPolicyRequest) (response *DeleteDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSPolicyRequest()
    }
    response = NewDeleteDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDDoSPolicyCaseRequest() (request *DeleteDDoSPolicyCaseRequest) {
    request = &DeleteDDoSPolicyCaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteDDoSPolicyCase")
    return
}

func NewDeleteDDoSPolicyCaseResponse() (response *DeleteDDoSPolicyCaseResponse) {
    response = &DeleteDDoSPolicyCaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a policy scenario.
func (c *Client) DeleteDDoSPolicyCase(request *DeleteDDoSPolicyCaseRequest) (response *DeleteDDoSPolicyCaseResponse, err error) {
    if request == nil {
        request = NewDeleteDDoSPolicyCaseRequest()
    }
    response = NewDeleteDDoSPolicyCaseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL4RulesRequest() (request *DeleteL4RulesRequest) {
    request = &DeleteL4RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteL4Rules")
    return
}

func NewDeleteL4RulesResponse() (response *DeleteL4RulesResponse) {
    response = &DeleteL4RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete one or more layer-4 forwarding rules.
func (c *Client) DeleteL4Rules(request *DeleteL4RulesRequest) (response *DeleteL4RulesResponse, err error) {
    if request == nil {
        request = NewDeleteL4RulesRequest()
    }
    response = NewDeleteL4RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteL7RulesRequest() (request *DeleteL7RulesRequest) {
    request = &DeleteL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DeleteL7Rules")
    return
}

func NewDeleteL7RulesResponse() (response *DeleteL7RulesResponse) {
    response = &DeleteL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete one or more layer-7 forwarding rules.
func (c *Client) DeleteL7Rules(request *DeleteL7RulesRequest) (response *DeleteL7RulesResponse, err error) {
    if request == nil {
        request = NewDeleteL7RulesRequest()
    }
    response = NewDeleteL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActionLogRequest() (request *DescribeActionLogRequest) {
    request = &DescribeActionLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeActionLog")
    return
}

func NewDescribeActionLogResponse() (response *DescribeActionLogResponse) {
    response = &DescribeActionLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get operation logs.
func (c *Client) DescribeActionLog(request *DescribeActionLogRequest) (response *DescribeActionLogResponse, err error) {
    if request == nil {
        request = NewDescribeActionLogRequest()
    }
    response = NewDescribeActionLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBGPIPL7RuleMaxCntRequest() (request *DescribeBGPIPL7RuleMaxCntRequest) {
    request = &DescribeBGPIPL7RuleMaxCntRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBGPIPL7RuleMaxCnt")
    return
}

func NewDescribeBGPIPL7RuleMaxCntResponse() (response *DescribeBGPIPL7RuleMaxCntResponse) {
    response = &DescribeBGPIPL7RuleMaxCntResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the maximum number of layer-7 rules that can be added for Anti-DDoS Advanced.
func (c *Client) DescribeBGPIPL7RuleMaxCnt(request *DescribeBGPIPL7RuleMaxCntRequest) (response *DescribeBGPIPL7RuleMaxCntResponse, err error) {
    if request == nil {
        request = NewDescribeBGPIPL7RuleMaxCntRequest()
    }
    response = NewDescribeBGPIPL7RuleMaxCntResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaradDataRequest() (request *DescribeBaradDataRequest) {
    request = &DescribeBaradDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBaradData")
    return
}

func NewDescribeBaradDataResponse() (response *DescribeBaradDataResponse) {
    response = &DescribeBaradDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to provide business forwarding metric data of Anti-DDoS services.
func (c *Client) DescribeBaradData(request *DescribeBaradDataRequest) (response *DescribeBaradDataResponse, err error) {
    if request == nil {
        request = NewDescribeBaradDataRequest()
    }
    response = NewDescribeBaradDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicCCThresholdRequest() (request *DescribeBasicCCThresholdRequest) {
    request = &DescribeBasicCCThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBasicCCThreshold")
    return
}

func NewDescribeBasicCCThresholdResponse() (response *DescribeBasicCCThresholdResponse) {
    response = &DescribeBasicCCThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the CC protection threshold of Anti-DDoS Basic.
func (c *Client) DescribeBasicCCThreshold(request *DescribeBasicCCThresholdRequest) (response *DescribeBasicCCThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeBasicCCThresholdRequest()
    }
    response = NewDescribeBasicCCThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicDeviceThresholdRequest() (request *DescribeBasicDeviceThresholdRequest) {
    request = &DescribeBasicDeviceThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBasicDeviceThreshold")
    return
}

func NewDescribeBasicDeviceThresholdResponse() (response *DescribeBasicDeviceThresholdResponse) {
    response = &DescribeBasicDeviceThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the blackhole threshold of Anti-DDoS Basic.
func (c *Client) DescribeBasicDeviceThreshold(request *DescribeBasicDeviceThresholdRequest) (response *DescribeBasicDeviceThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeBasicDeviceThresholdRequest()
    }
    response = NewDescribeBasicDeviceThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBizHttpStatusRequest() (request *DescribeBizHttpStatusRequest) {
    request = &DescribeBizHttpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeBizHttpStatus")
    return
}

func NewDescribeBizHttpStatusResponse() (response *DescribeBizHttpStatusResponse) {
    response = &DescribeBizHttpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the statistics on the status codes of business traffic.
func (c *Client) DescribeBizHttpStatus(request *DescribeBizHttpStatusRequest) (response *DescribeBizHttpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBizHttpStatusRequest()
    }
    response = NewDescribeBizHttpStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCAlarmThresholdRequest() (request *DescribeCCAlarmThresholdRequest) {
    request = &DescribeCCAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCAlarmThreshold")
    return
}

func NewDescribeCCAlarmThresholdResponse() (response *DescribeCCAlarmThresholdResponse) {
    response = &DescribeCCAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the alarm notification threshold set for CC attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
func (c *Client) DescribeCCAlarmThreshold(request *DescribeCCAlarmThresholdRequest) (response *DescribeCCAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeCCAlarmThresholdRequest()
    }
    response = NewDescribeCCAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCEvListRequest() (request *DescribeCCEvListRequest) {
    request = &DescribeCCEvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCEvList")
    return
}

func NewDescribeCCEvListResponse() (response *DescribeCCEvListResponse) {
    response = &DescribeCCEvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the CC attack event list.
func (c *Client) DescribeCCEvList(request *DescribeCCEvListRequest) (response *DescribeCCEvListResponse, err error) {
    if request == nil {
        request = NewDescribeCCEvListRequest()
    }
    response = NewDescribeCCEvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCFrequencyRulesRequest() (request *DescribeCCFrequencyRulesRequest) {
    request = &DescribeCCFrequencyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCFrequencyRules")
    return
}

func NewDescribeCCFrequencyRulesResponse() (response *DescribeCCFrequencyRulesResponse) {
    response = &DescribeCCFrequencyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get an access frequency control rule for CC protection.
func (c *Client) DescribeCCFrequencyRules(request *DescribeCCFrequencyRulesRequest) (response *DescribeCCFrequencyRulesResponse, err error) {
    if request == nil {
        request = NewDescribeCCFrequencyRulesRequest()
    }
    response = NewDescribeCCFrequencyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCIpAllowDenyRequest() (request *DescribeCCIpAllowDenyRequest) {
    request = &DescribeCCIpAllowDenyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCIpAllowDeny")
    return
}

func NewDescribeCCIpAllowDenyResponse() (response *DescribeCCIpAllowDenyResponse) {
    response = &DescribeCCIpAllowDenyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the CC IP blocklist/allowlist.
func (c *Client) DescribeCCIpAllowDeny(request *DescribeCCIpAllowDenyRequest) (response *DescribeCCIpAllowDenyResponse, err error) {
    if request == nil {
        request = NewDescribeCCIpAllowDenyRequest()
    }
    response = NewDescribeCCIpAllowDenyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCSelfDefinePolicyRequest() (request *DescribeCCSelfDefinePolicyRequest) {
    request = &DescribeCCSelfDefinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCSelfDefinePolicy")
    return
}

func NewDescribeCCSelfDefinePolicyResponse() (response *DescribeCCSelfDefinePolicyResponse) {
    response = &DescribeCCSelfDefinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get a custom CC policy.
func (c *Client) DescribeCCSelfDefinePolicy(request *DescribeCCSelfDefinePolicyRequest) (response *DescribeCCSelfDefinePolicyResponse, err error) {
    if request == nil {
        request = NewDescribeCCSelfDefinePolicyRequest()
    }
    response = NewDescribeCCSelfDefinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCTrendRequest() (request *DescribeCCTrendRequest) {
    request = &DescribeCCTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCTrend")
    return
}

func NewDescribeCCTrendResponse() (response *DescribeCCTrendResponse) {
    response = &DescribeCCTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get CC attack metric data, including total requests peak (QPS) and attack requests (QPS).
func (c *Client) DescribeCCTrend(request *DescribeCCTrendRequest) (response *DescribeCCTrendResponse, err error) {
    if request == nil {
        request = NewDescribeCCTrendRequest()
    }
    response = NewDescribeCCTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCUrlAllowRequest() (request *DescribeCCUrlAllowRequest) {
    request = &DescribeCCUrlAllowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeCCUrlAllow")
    return
}

func NewDescribeCCUrlAllowResponse() (response *DescribeCCUrlAllowResponse) {
    response = &DescribeCCUrlAllowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the CC URL allowlist.
func (c *Client) DescribeCCUrlAllow(request *DescribeCCUrlAllowRequest) (response *DescribeCCUrlAllowResponse, err error) {
    if request == nil {
        request = NewDescribeCCUrlAllowRequest()
    }
    response = NewDescribeCCUrlAllowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAlarmThresholdRequest() (request *DescribeDDoSAlarmThresholdRequest) {
    request = &DescribeDDoSAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSAlarmThreshold")
    return
}

func NewDescribeDDoSAlarmThresholdResponse() (response *DescribeDDoSAlarmThresholdResponse) {
    response = &DescribeDDoSAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the alarm notification threshold set for DDoS attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
func (c *Client) DescribeDDoSAlarmThreshold(request *DescribeDDoSAlarmThresholdRequest) (response *DescribeDDoSAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAlarmThresholdRequest()
    }
    response = NewDescribeDDoSAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackIPRegionMapRequest() (request *DescribeDDoSAttackIPRegionMapRequest) {
    request = &DescribeDDoSAttackIPRegionMapRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSAttackIPRegionMap")
    return
}

func NewDescribeDDoSAttackIPRegionMapResponse() (response *DescribeDDoSAttackIPRegionMapResponse) {
    response = &DescribeDDoSAttackIPRegionMapResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the geographical distribution map of DDoS attack source IPs. It supports display by global regions and Chinese provinces.
func (c *Client) DescribeDDoSAttackIPRegionMap(request *DescribeDDoSAttackIPRegionMapRequest) (response *DescribeDDoSAttackIPRegionMapResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackIPRegionMapRequest()
    }
    response = NewDescribeDDoSAttackIPRegionMapResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSAttackSourceRequest() (request *DescribeDDoSAttackSourceRequest) {
    request = &DescribeDDoSAttackSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSAttackSource")
    return
}

func NewDescribeDDoSAttackSourceResponse() (response *DescribeDDoSAttackSourceResponse) {
    response = &DescribeDDoSAttackSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the DDoS attack source list.
func (c *Client) DescribeDDoSAttackSource(request *DescribeDDoSAttackSourceRequest) (response *DescribeDDoSAttackSourceResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSAttackSourceRequest()
    }
    response = NewDescribeDDoSAttackSourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSCountRequest() (request *DescribeDDoSCountRequest) {
    request = &DescribeDDoSCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSCount")
    return
}

func NewDescribeDDoSCountResponse() (response *DescribeDDoSCountResponse) {
    response = &DescribeDDoSCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the DDoS attack proportion analysis.
func (c *Client) DescribeDDoSCount(request *DescribeDDoSCountRequest) (response *DescribeDDoSCountResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSCountRequest()
    }
    response = NewDescribeDDoSCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSDefendStatusRequest() (request *DescribeDDoSDefendStatusRequest) {
    request = &DescribeDDoSDefendStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSDefendStatus")
    return
}

func NewDescribeDDoSDefendStatusResponse() (response *DescribeDDoSDefendStatusResponse) {
    response = &DescribeDDoSDefendStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the DDoS protection status (temporarily disabled status). It is supported for Anti-DDoS Basic, single IP instance, multi-IP instance, Anti-DDoS Advanced, and Anti-DDoS Ultimate. It is used to query whether DDoS protection is temporarily disabled, and if yes, return parameters such as temporary disablement duration.
func (c *Client) DescribeDDoSDefendStatus(request *DescribeDDoSDefendStatusRequest) (response *DescribeDDoSDefendStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSDefendStatusRequest()
    }
    response = NewDescribeDDoSDefendStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSEvInfoRequest() (request *DescribeDDoSEvInfoRequest) {
    request = &DescribeDDoSEvInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSEvInfo")
    return
}

func NewDescribeDDoSEvInfoResponse() (response *DescribeDDoSEvInfoResponse) {
    response = &DescribeDDoSEvInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get details of a specific DDoS attack.
func (c *Client) DescribeDDoSEvInfo(request *DescribeDDoSEvInfoRequest) (response *DescribeDDoSEvInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSEvInfoRequest()
    }
    response = NewDescribeDDoSEvInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSEvListRequest() (request *DescribeDDoSEvListRequest) {
    request = &DescribeDDoSEvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSEvList")
    return
}

func NewDescribeDDoSEvListResponse() (response *DescribeDDoSEvListResponse) {
    response = &DescribeDDoSEvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the DDoS attack event list.
func (c *Client) DescribeDDoSEvList(request *DescribeDDoSEvListRequest) (response *DescribeDDoSEvListResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSEvListRequest()
    }
    response = NewDescribeDDoSEvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSIpLogRequest() (request *DescribeDDoSIpLogRequest) {
    request = &DescribeDDoSIpLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSIpLog")
    return
}

func NewDescribeDDoSIpLogResponse() (response *DescribeDDoSIpLogResponse) {
    response = &DescribeDDoSIpLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get a DDoS IP attack log.
func (c *Client) DescribeDDoSIpLog(request *DescribeDDoSIpLogRequest) (response *DescribeDDoSIpLogResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSIpLogRequest()
    }
    response = NewDescribeDDoSIpLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetCountRequest() (request *DescribeDDoSNetCountRequest) {
    request = &DescribeDDoSNetCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetCount")
    return
}

func NewDescribeDDoSNetCountResponse() (response *DescribeDDoSNetCountResponse) {
    response = &DescribeDDoSNetCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the DDoS attack proportion analysis for an Anti-DDoS Ultimate resource.
func (c *Client) DescribeDDoSNetCount(request *DescribeDDoSNetCountRequest) (response *DescribeDDoSNetCountResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetCountRequest()
    }
    response = NewDescribeDDoSNetCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetEvInfoRequest() (request *DescribeDDoSNetEvInfoRequest) {
    request = &DescribeDDoSNetEvInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetEvInfo")
    return
}

func NewDescribeDDoSNetEvInfoResponse() (response *DescribeDDoSNetEvInfoResponse) {
    response = &DescribeDDoSNetEvInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the DDoS attack event details of an Anti-DDoS Ultimate resource.
func (c *Client) DescribeDDoSNetEvInfo(request *DescribeDDoSNetEvInfoRequest) (response *DescribeDDoSNetEvInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetEvInfoRequest()
    }
    response = NewDescribeDDoSNetEvInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetEvListRequest() (request *DescribeDDoSNetEvListRequest) {
    request = &DescribeDDoSNetEvListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetEvList")
    return
}

func NewDescribeDDoSNetEvListResponse() (response *DescribeDDoSNetEvListResponse) {
    response = &DescribeDDoSNetEvListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the DDoS attack event list of an Anti-DDoS Ultimate resource.
func (c *Client) DescribeDDoSNetEvList(request *DescribeDDoSNetEvListRequest) (response *DescribeDDoSNetEvListResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetEvListRequest()
    }
    response = NewDescribeDDoSNetEvListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetIpLogRequest() (request *DescribeDDoSNetIpLogRequest) {
    request = &DescribeDDoSNetIpLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetIpLog")
    return
}

func NewDescribeDDoSNetIpLogResponse() (response *DescribeDDoSNetIpLogResponse) {
    response = &DescribeDDoSNetIpLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the DDoS IP attack logs of an Anti-DDoS Ultimate resource.
func (c *Client) DescribeDDoSNetIpLog(request *DescribeDDoSNetIpLogRequest) (response *DescribeDDoSNetIpLogResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetIpLogRequest()
    }
    response = NewDescribeDDoSNetIpLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSNetTrendRequest() (request *DescribeDDoSNetTrendRequest) {
    request = &DescribeDDoSNetTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSNetTrend")
    return
}

func NewDescribeDDoSNetTrendResponse() (response *DescribeDDoSNetTrendResponse) {
    response = &DescribeDDoSNetTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the DDoS attack metric data of an Anti-DDoS Ultimate resource.
func (c *Client) DescribeDDoSNetTrend(request *DescribeDDoSNetTrendRequest) (response *DescribeDDoSNetTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSNetTrendRequest()
    }
    response = NewDescribeDDoSNetTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSPolicyRequest() (request *DescribeDDoSPolicyRequest) {
    request = &DescribeDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSPolicy")
    return
}

func NewDescribeDDoSPolicyResponse() (response *DescribeDDoSPolicyResponse) {
    response = &DescribeDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get an advanced DDoS policy.
func (c *Client) DescribeDDoSPolicy(request *DescribeDDoSPolicyRequest) (response *DescribeDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSPolicyRequest()
    }
    response = NewDescribeDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSTrendRequest() (request *DescribeDDoSTrendRequest) {
    request = &DescribeDDoSTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSTrend")
    return
}

func NewDescribeDDoSTrendResponse() (response *DescribeDDoSTrendResponse) {
    response = &DescribeDDoSTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the data of DDoS attack traffic bandwidth and attack packet rate.
func (c *Client) DescribeDDoSTrend(request *DescribeDDoSTrendRequest) (response *DescribeDDoSTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSTrendRequest()
    }
    response = NewDescribeDDoSTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDDoSUsedStatisRequest() (request *DescribeDDoSUsedStatisRequest) {
    request = &DescribeDDoSUsedStatisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeDDoSUsedStatis")
    return
}

func NewDescribeDDoSUsedStatisResponse() (response *DescribeDDoSUsedStatisResponse) {
    response = &DescribeDDoSUsedStatisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to count the number of days of Anti-DDoS resource use and number of DDoS attacks defended against.
func (c *Client) DescribeDDoSUsedStatis(request *DescribeDDoSUsedStatisRequest) (response *DescribeDDoSUsedStatisResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSUsedStatisRequest()
    }
    response = NewDescribeDDoSUsedStatisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIPProductInfoRequest() (request *DescribeIPProductInfoRequest) {
    request = &DescribeIPProductInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeIPProductInfo")
    return
}

func NewDescribeIPProductInfoResponse() (response *DescribeIPProductInfoResponse) {
    response = &DescribeIPProductInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the Tencent Cloud asset information corresponding to an IP of a single IP instance or multi-IP instance, which is supported only for IPs of single IP instances and multi-IP instances.
func (c *Client) DescribeIPProductInfo(request *DescribeIPProductInfoRequest) (response *DescribeIPProductInfoResponse, err error) {
    if request == nil {
        request = NewDescribeIPProductInfoRequest()
    }
    response = NewDescribeIPProductInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInsurePacksRequest() (request *DescribeInsurePacksRequest) {
    request = &DescribeInsurePacksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeInsurePacks")
    return
}

func NewDescribeInsurePacksResponse() (response *DescribeInsurePacksResponse) {
    response = &DescribeInsurePacksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the guarantee package list.
func (c *Client) DescribeInsurePacks(request *DescribeInsurePacksRequest) (response *DescribeInsurePacksResponse, err error) {
    if request == nil {
        request = NewDescribeInsurePacksRequest()
    }
    response = NewDescribeInsurePacksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpBlockListRequest() (request *DescribeIpBlockListRequest) {
    request = &DescribeIpBlockListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeIpBlockList")
    return
}

func NewDescribeIpBlockListResponse() (response *DescribeIpBlockListResponse) {
    response = &DescribeIpBlockListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the blocked IP list.
func (c *Client) DescribeIpBlockList(request *DescribeIpBlockListRequest) (response *DescribeIpBlockListResponse, err error) {
    if request == nil {
        request = NewDescribeIpBlockListRequest()
    }
    response = NewDescribeIpBlockListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpUnBlockListRequest() (request *DescribeIpUnBlockListRequest) {
    request = &DescribeIpUnBlockListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeIpUnBlockList")
    return
}

func NewDescribeIpUnBlockListResponse() (response *DescribeIpUnBlockListResponse) {
    response = &DescribeIpUnBlockListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the IP unblocking records.
func (c *Client) DescribeIpUnBlockList(request *DescribeIpUnBlockListRequest) (response *DescribeIpUnBlockListResponse, err error) {
    if request == nil {
        request = NewDescribeIpUnBlockListRequest()
    }
    response = NewDescribeIpUnBlockListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4HealthConfigRequest() (request *DescribeL4HealthConfigRequest) {
    request = &DescribeL4HealthConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeL4HealthConfig")
    return
}

func NewDescribeL4HealthConfigResponse() (response *DescribeL4HealthConfigResponse) {
    response = &DescribeL4HealthConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to export the layer-4 health check configuration.
func (c *Client) DescribeL4HealthConfig(request *DescribeL4HealthConfigRequest) (response *DescribeL4HealthConfigResponse, err error) {
    if request == nil {
        request = NewDescribeL4HealthConfigRequest()
    }
    response = NewDescribeL4HealthConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL4RulesErrHealthRequest() (request *DescribeL4RulesErrHealthRequest) {
    request = &DescribeL4RulesErrHealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeL4RulesErrHealth")
    return
}

func NewDescribeL4RulesErrHealthResponse() (response *DescribeL4RulesErrHealthResponse) {
    response = &DescribeL4RulesErrHealthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the exception result of a layer-4 forwarding rule health check.
func (c *Client) DescribeL4RulesErrHealth(request *DescribeL4RulesErrHealthRequest) (response *DescribeL4RulesErrHealthResponse, err error) {
    if request == nil {
        request = NewDescribeL4RulesErrHealthRequest()
    }
    response = NewDescribeL4RulesErrHealthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeL7HealthConfigRequest() (request *DescribeL7HealthConfigRequest) {
    request = &DescribeL7HealthConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeL7HealthConfig")
    return
}

func NewDescribeL7HealthConfigResponse() (response *DescribeL7HealthConfigResponse) {
    response = &DescribeL7HealthConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to export the layer-7 health check configuration.
func (c *Client) DescribeL7HealthConfig(request *DescribeL7HealthConfigRequest) (response *DescribeL7HealthConfigResponse, err error) {
    if request == nil {
        request = NewDescribeL7HealthConfigRequest()
    }
    response = NewDescribeL7HealthConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackIndexRequest() (request *DescribePackIndexRequest) {
    request = &DescribePackIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribePackIndex")
    return
}

func NewDescribePackIndexResponse() (response *DescribePackIndexResponse) {
    response = &DescribePackIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the product overview statistics. It is supported for Anti-DDoS Pro, Anti-DDoS Advanced, and Anti-DDoS Ultimate.
func (c *Client) DescribePackIndex(request *DescribePackIndexRequest) (response *DescribePackIndexResponse, err error) {
    if request == nil {
        request = NewDescribePackIndexRequest()
    }
    response = NewDescribePackIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePcapRequest() (request *DescribePcapRequest) {
    request = &DescribePcapRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribePcap")
    return
}

func NewDescribePcapResponse() (response *DescribePcapResponse) {
    response = &DescribePcapResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to download the PCAP packet of an attack event.
func (c *Client) DescribePcap(request *DescribePcapRequest) (response *DescribePcapResponse, err error) {
    if request == nil {
        request = NewDescribePcapRequest()
    }
    response = NewDescribePcapResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyCaseRequest() (request *DescribePolicyCaseRequest) {
    request = &DescribePolicyCaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribePolicyCase")
    return
}

func NewDescribePolicyCaseResponse() (response *DescribePolicyCaseResponse) {
    response = &DescribePolicyCaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the policy scenario.
func (c *Client) DescribePolicyCase(request *DescribePolicyCaseRequest) (response *DescribePolicyCaseResponse, err error) {
    if request == nil {
        request = NewDescribePolicyCaseRequest()
    }
    response = NewDescribePolicyCaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResIpListRequest() (request *DescribeResIpListRequest) {
    request = &DescribeResIpListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeResIpList")
    return
}

func NewDescribeResIpListResponse() (response *DescribeResIpListResponse) {
    response = &DescribeResIpListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the IP list of a resource.
func (c *Client) DescribeResIpList(request *DescribeResIpListRequest) (response *DescribeResIpListResponse, err error) {
    if request == nil {
        request = NewDescribeResIpListRequest()
    }
    response = NewDescribeResIpListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceListRequest() (request *DescribeResourceListRequest) {
    request = &DescribeResourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeResourceList")
    return
}

func NewDescribeResourceListResponse() (response *DescribeResourceListResponse) {
    response = &DescribeResourceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the resource list.
func (c *Client) DescribeResourceList(request *DescribeResourceListRequest) (response *DescribeResourceListResponse, err error) {
    if request == nil {
        request = NewDescribeResourceListRequest()
    }
    response = NewDescribeResourceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuleSetsRequest() (request *DescribeRuleSetsRequest) {
    request = &DescribeRuleSetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeRuleSets")
    return
}

func NewDescribeRuleSetsResponse() (response *DescribeRuleSetsResponse) {
    response = &DescribeRuleSetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the number of rules of a resource.
func (c *Client) DescribeRuleSets(request *DescribeRuleSetsRequest) (response *DescribeRuleSetsResponse, err error) {
    if request == nil {
        request = NewDescribeRuleSetsRequest()
    }
    response = NewDescribeRuleSetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSchedulingDomainListRequest() (request *DescribeSchedulingDomainListRequest) {
    request = &DescribeSchedulingDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeSchedulingDomainList")
    return
}

func NewDescribeSchedulingDomainListResponse() (response *DescribeSchedulingDomainListResponse) {
    response = &DescribeSchedulingDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Get scheduling domain name list
func (c *Client) DescribeSchedulingDomainList(request *DescribeSchedulingDomainListRequest) (response *DescribeSchedulingDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeSchedulingDomainListRequest()
    }
    response = NewDescribeSchedulingDomainListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecIndexRequest() (request *DescribeSecIndexRequest) {
    request = &DescribeSecIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeSecIndex")
    return
}

func NewDescribeSecIndexResponse() (response *DescribeSecIndexResponse) {
    response = &DescribeSecIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the security statistics of the current month.
func (c *Client) DescribeSecIndex(request *DescribeSecIndexRequest) (response *DescribeSecIndexResponse, err error) {
    if request == nil {
        request = NewDescribeSecIndexRequest()
    }
    response = NewDescribeSecIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSourceIpSegmentRequest() (request *DescribeSourceIpSegmentRequest) {
    request = &DescribeSourceIpSegmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeSourceIpSegment")
    return
}

func NewDescribeSourceIpSegmentResponse() (response *DescribeSourceIpSegmentResponse) {
    response = &DescribeSourceIpSegmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the intermediate IP range. It is supported for Anti-DDoS Advanced and Anti-DDoS Ultimate.
func (c *Client) DescribeSourceIpSegment(request *DescribeSourceIpSegmentRequest) (response *DescribeSourceIpSegmentResponse, err error) {
    if request == nil {
        request = NewDescribeSourceIpSegmentRequest()
    }
    response = NewDescribeSourceIpSegmentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTransmitStatisRequest() (request *DescribeTransmitStatisRequest) {
    request = &DescribeTransmitStatisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeTransmitStatis")
    return
}

func NewDescribeTransmitStatisResponse() (response *DescribeTransmitStatisResponse) {
    response = &DescribeTransmitStatisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the business forwarding statistics, including forwarded traffic and packet forwarding rate.
func (c *Client) DescribeTransmitStatis(request *DescribeTransmitStatisRequest) (response *DescribeTransmitStatisResponse, err error) {
    if request == nil {
        request = NewDescribeTransmitStatisRequest()
    }
    response = NewDescribeTransmitStatisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnBlockStatisRequest() (request *DescribeUnBlockStatisRequest) {
    request = &DescribeUnBlockStatisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribeUnBlockStatis")
    return
}

func NewDescribeUnBlockStatisResponse() (response *DescribeUnBlockStatisResponse) {
    response = &DescribeUnBlockStatisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the number of blackhole unblockings.
func (c *Client) DescribeUnBlockStatis(request *DescribeUnBlockStatisRequest) (response *DescribeUnBlockStatisResponse, err error) {
    if request == nil {
        request = NewDescribeUnBlockStatisRequest()
    }
    response = NewDescribeUnBlockStatisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribleL4RulesRequest() (request *DescribleL4RulesRequest) {
    request = &DescribleL4RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribleL4Rules")
    return
}

func NewDescribleL4RulesResponse() (response *DescribleL4RulesResponse) {
    response = &DescribleL4RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get a layer-4 forwarding rule.
func (c *Client) DescribleL4Rules(request *DescribleL4RulesRequest) (response *DescribleL4RulesResponse, err error) {
    if request == nil {
        request = NewDescribleL4RulesRequest()
    }
    response = NewDescribleL4RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribleL7RulesRequest() (request *DescribleL7RulesRequest) {
    request = &DescribleL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribleL7Rules")
    return
}

func NewDescribleL7RulesResponse() (response *DescribleL7RulesResponse) {
    response = &DescribleL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get a layer-7 forwarding rule.
func (c *Client) DescribleL7Rules(request *DescribleL7RulesRequest) (response *DescribleL7RulesResponse, err error) {
    if request == nil {
        request = NewDescribleL7RulesRequest()
    }
    response = NewDescribleL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribleRegionCountRequest() (request *DescribleRegionCountRequest) {
    request = &DescribleRegionCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "DescribleRegionCount")
    return
}

func NewDescribleRegionCountResponse() (response *DescribleRegionCountResponse) {
    response = &DescribleRegionCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the number of resource instances in a region.
func (c *Client) DescribleRegionCount(request *DescribleRegionCountRequest) (response *DescribleRegionCountResponse, err error) {
    if request == nil {
        request = NewDescribleRegionCountRequest()
    }
    response = NewDescribleRegionCountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCAlarmThresholdRequest() (request *ModifyCCAlarmThresholdRequest) {
    request = &ModifyCCAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCAlarmThreshold")
    return
}

func NewModifyCCAlarmThresholdResponse() (response *ModifyCCAlarmThresholdResponse) {
    response = &ModifyCCAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to set the alarm notification threshold for CC attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
func (c *Client) ModifyCCAlarmThreshold(request *ModifyCCAlarmThresholdRequest) (response *ModifyCCAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewModifyCCAlarmThresholdRequest()
    }
    response = NewModifyCCAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCFrequencyRulesRequest() (request *ModifyCCFrequencyRulesRequest) {
    request = &ModifyCCFrequencyRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCFrequencyRules")
    return
}

func NewModifyCCFrequencyRulesResponse() (response *ModifyCCFrequencyRulesResponse) {
    response = &ModifyCCFrequencyRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify an access frequency control rule for CC protection.
func (c *Client) ModifyCCFrequencyRules(request *ModifyCCFrequencyRulesRequest) (response *ModifyCCFrequencyRulesResponse, err error) {
    if request == nil {
        request = NewModifyCCFrequencyRulesRequest()
    }
    response = NewModifyCCFrequencyRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCFrequencyRulesStatusRequest() (request *ModifyCCFrequencyRulesStatusRequest) {
    request = &ModifyCCFrequencyRulesStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCFrequencyRulesStatus")
    return
}

func NewModifyCCFrequencyRulesStatusResponse() (response *ModifyCCFrequencyRulesStatusResponse) {
    response = &ModifyCCFrequencyRulesStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable or disable an access frequency control rule for CC protection.
func (c *Client) ModifyCCFrequencyRulesStatus(request *ModifyCCFrequencyRulesStatusRequest) (response *ModifyCCFrequencyRulesStatusResponse, err error) {
    if request == nil {
        request = NewModifyCCFrequencyRulesStatusRequest()
    }
    response = NewModifyCCFrequencyRulesStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCHostProtectionRequest() (request *ModifyCCHostProtectionRequest) {
    request = &ModifyCCHostProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCHostProtection")
    return
}

func NewModifyCCHostProtectionResponse() (response *ModifyCCHostProtectionResponse) {
    response = &ModifyCCHostProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable or disable CC domain name protection.
func (c *Client) ModifyCCHostProtection(request *ModifyCCHostProtectionRequest) (response *ModifyCCHostProtectionResponse, err error) {
    if request == nil {
        request = NewModifyCCHostProtectionRequest()
    }
    response = NewModifyCCHostProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCIpAllowDenyRequest() (request *ModifyCCIpAllowDenyRequest) {
    request = &ModifyCCIpAllowDenyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCIpAllowDeny")
    return
}

func NewModifyCCIpAllowDenyResponse() (response *ModifyCCIpAllowDenyResponse) {
    response = &ModifyCCIpAllowDenyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add/remove a CC IP to/from the blocklist/allowlist.
func (c *Client) ModifyCCIpAllowDeny(request *ModifyCCIpAllowDenyRequest) (response *ModifyCCIpAllowDenyResponse, err error) {
    if request == nil {
        request = NewModifyCCIpAllowDenyRequest()
    }
    response = NewModifyCCIpAllowDenyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCLevelRequest() (request *ModifyCCLevelRequest) {
    request = &ModifyCCLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCLevel")
    return
}

func NewModifyCCLevelResponse() (response *ModifyCCLevelResponse) {
    response = &ModifyCCLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify CC protection level.
func (c *Client) ModifyCCLevel(request *ModifyCCLevelRequest) (response *ModifyCCLevelResponse, err error) {
    if request == nil {
        request = NewModifyCCLevelRequest()
    }
    response = NewModifyCCLevelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCPolicySwitchRequest() (request *ModifyCCPolicySwitchRequest) {
    request = &ModifyCCPolicySwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCPolicySwitch")
    return
}

func NewModifyCCPolicySwitchResponse() (response *ModifyCCPolicySwitchResponse) {
    response = &ModifyCCPolicySwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable or disable a custom CC policy.
func (c *Client) ModifyCCPolicySwitch(request *ModifyCCPolicySwitchRequest) (response *ModifyCCPolicySwitchResponse, err error) {
    if request == nil {
        request = NewModifyCCPolicySwitchRequest()
    }
    response = NewModifyCCPolicySwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCSelfDefinePolicyRequest() (request *ModifyCCSelfDefinePolicyRequest) {
    request = &ModifyCCSelfDefinePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCSelfDefinePolicy")
    return
}

func NewModifyCCSelfDefinePolicyResponse() (response *ModifyCCSelfDefinePolicyResponse) {
    response = &ModifyCCSelfDefinePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify a custom CC policy.
func (c *Client) ModifyCCSelfDefinePolicy(request *ModifyCCSelfDefinePolicyRequest) (response *ModifyCCSelfDefinePolicyResponse, err error) {
    if request == nil {
        request = NewModifyCCSelfDefinePolicyRequest()
    }
    response = NewModifyCCSelfDefinePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCThresholdRequest() (request *ModifyCCThresholdRequest) {
    request = &ModifyCCThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCThreshold")
    return
}

func NewModifyCCThresholdResponse() (response *ModifyCCThresholdResponse) {
    response = &ModifyCCThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the CC protection threshold.
func (c *Client) ModifyCCThreshold(request *ModifyCCThresholdRequest) (response *ModifyCCThresholdResponse, err error) {
    if request == nil {
        request = NewModifyCCThresholdRequest()
    }
    response = NewModifyCCThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCCUrlAllowRequest() (request *ModifyCCUrlAllowRequest) {
    request = &ModifyCCUrlAllowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyCCUrlAllow")
    return
}

func NewModifyCCUrlAllowResponse() (response *ModifyCCUrlAllowResponse) {
    response = &ModifyCCUrlAllowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add/remove a CC URL to/from the allowlist.
func (c *Client) ModifyCCUrlAllow(request *ModifyCCUrlAllowRequest) (response *ModifyCCUrlAllowResponse, err error) {
    if request == nil {
        request = NewModifyCCUrlAllowRequest()
    }
    response = NewModifyCCUrlAllowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSAIStatusRequest() (request *ModifyDDoSAIStatusRequest) {
    request = &ModifyDDoSAIStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSAIStatus")
    return
}

func NewModifyDDoSAIStatusResponse() (response *ModifyDDoSAIStatusResponse) {
    response = &ModifyDDoSAIStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to read or modify DDoS AI protection status.
func (c *Client) ModifyDDoSAIStatus(request *ModifyDDoSAIStatusRequest) (response *ModifyDDoSAIStatusResponse, err error) {
    if request == nil {
        request = NewModifyDDoSAIStatusRequest()
    }
    response = NewModifyDDoSAIStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSAlarmThresholdRequest() (request *ModifyDDoSAlarmThresholdRequest) {
    request = &ModifyDDoSAlarmThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSAlarmThreshold")
    return
}

func NewModifyDDoSAlarmThresholdResponse() (response *ModifyDDoSAlarmThresholdResponse) {
    response = &ModifyDDoSAlarmThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to set the alarm notification threshold for DDoS attacks in Anti-DDoS Pro, Anti-DDoS Advanced, Anti-DDoS Ultimate, and Chess Shield.
func (c *Client) ModifyDDoSAlarmThreshold(request *ModifyDDoSAlarmThresholdRequest) (response *ModifyDDoSAlarmThresholdResponse, err error) {
    if request == nil {
        request = NewModifyDDoSAlarmThresholdRequest()
    }
    response = NewModifyDDoSAlarmThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSDefendStatusRequest() (request *ModifyDDoSDefendStatusRequest) {
    request = &ModifyDDoSDefendStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSDefendStatus")
    return
}

func NewModifyDDoSDefendStatusResponse() (response *ModifyDDoSDefendStatusResponse) {
    response = &ModifyDDoSDefendStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable or disable DDoS. It can disable DDoS protection for a period of time, which will be automatically enabled after the period of time elapses.
func (c *Client) ModifyDDoSDefendStatus(request *ModifyDDoSDefendStatusRequest) (response *ModifyDDoSDefendStatusResponse, err error) {
    if request == nil {
        request = NewModifyDDoSDefendStatusRequest()
    }
    response = NewModifyDDoSDefendStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSLevelRequest() (request *ModifyDDoSLevelRequest) {
    request = &ModifyDDoSLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSLevel")
    return
}

func NewModifyDDoSLevelResponse() (response *ModifyDDoSLevelResponse) {
    response = &ModifyDDoSLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to read or modify DDoS protection level.
func (c *Client) ModifyDDoSLevel(request *ModifyDDoSLevelRequest) (response *ModifyDDoSLevelResponse, err error) {
    if request == nil {
        request = NewModifyDDoSLevelRequest()
    }
    response = NewModifyDDoSLevelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSPolicyRequest() (request *ModifyDDoSPolicyRequest) {
    request = &ModifyDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSPolicy")
    return
}

func NewModifyDDoSPolicyResponse() (response *ModifyDDoSPolicyResponse) {
    response = &ModifyDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify an advanced DDoS policy.
func (c *Client) ModifyDDoSPolicy(request *ModifyDDoSPolicyRequest) (response *ModifyDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewModifyDDoSPolicyRequest()
    }
    response = NewModifyDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSPolicyCaseRequest() (request *ModifyDDoSPolicyCaseRequest) {
    request = &ModifyDDoSPolicyCaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSPolicyCase")
    return
}

func NewModifyDDoSPolicyCaseResponse() (response *ModifyDDoSPolicyCaseResponse) {
    response = &ModifyDDoSPolicyCaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify a policy scenario.
func (c *Client) ModifyDDoSPolicyCase(request *ModifyDDoSPolicyCaseRequest) (response *ModifyDDoSPolicyCaseResponse, err error) {
    if request == nil {
        request = NewModifyDDoSPolicyCaseRequest()
    }
    response = NewModifyDDoSPolicyCaseResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSPolicyNameRequest() (request *ModifyDDoSPolicyNameRequest) {
    request = &ModifyDDoSPolicyNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSPolicyName")
    return
}

func NewModifyDDoSPolicyNameResponse() (response *ModifyDDoSPolicyNameResponse) {
    response = &ModifyDDoSPolicyNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to rename an advanced DDoS policy.
func (c *Client) ModifyDDoSPolicyName(request *ModifyDDoSPolicyNameRequest) (response *ModifyDDoSPolicyNameResponse, err error) {
    if request == nil {
        request = NewModifyDDoSPolicyNameRequest()
    }
    response = NewModifyDDoSPolicyNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSSwitchRequest() (request *ModifyDDoSSwitchRequest) {
    request = &ModifyDDoSSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSSwitch")
    return
}

func NewModifyDDoSSwitchResponse() (response *ModifyDDoSSwitchResponse) {
    response = &ModifyDDoSSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable or disable DDoS protection, which is only supported for Anti-DDoS Basic.
func (c *Client) ModifyDDoSSwitch(request *ModifyDDoSSwitchRequest) (response *ModifyDDoSSwitchResponse, err error) {
    if request == nil {
        request = NewModifyDDoSSwitchRequest()
    }
    response = NewModifyDDoSSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSThresholdRequest() (request *ModifyDDoSThresholdRequest) {
    request = &ModifyDDoSThresholdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSThreshold")
    return
}

func NewModifyDDoSThresholdResponse() (response *ModifyDDoSThresholdResponse) {
    response = &ModifyDDoSThresholdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the DDoS cleansing threshold.
func (c *Client) ModifyDDoSThreshold(request *ModifyDDoSThresholdRequest) (response *ModifyDDoSThresholdResponse, err error) {
    if request == nil {
        request = NewModifyDDoSThresholdRequest()
    }
    response = NewModifyDDoSThresholdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDDoSWaterKeyRequest() (request *ModifyDDoSWaterKeyRequest) {
    request = &ModifyDDoSWaterKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyDDoSWaterKey")
    return
}

func NewModifyDDoSWaterKeyResponse() (response *ModifyDDoSWaterKeyResponse) {
    response = &ModifyDDoSWaterKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to add, delete, enable, or disable a watermark key.
func (c *Client) ModifyDDoSWaterKey(request *ModifyDDoSWaterKeyRequest) (response *ModifyDDoSWaterKeyResponse, err error) {
    if request == nil {
        request = NewModifyDDoSWaterKeyRequest()
    }
    response = NewModifyDDoSWaterKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyElasticLimitRequest() (request *ModifyElasticLimitRequest) {
    request = &ModifyElasticLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyElasticLimit")
    return
}

func NewModifyElasticLimitResponse() (response *ModifyElasticLimitResponse) {
    response = &ModifyElasticLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the elastic protection threshold.
func (c *Client) ModifyElasticLimit(request *ModifyElasticLimitRequest) (response *ModifyElasticLimitResponse, err error) {
    if request == nil {
        request = NewModifyElasticLimitRequest()
    }
    response = NewModifyElasticLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4HealthRequest() (request *ModifyL4HealthRequest) {
    request = &ModifyL4HealthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyL4Health")
    return
}

func NewModifyL4HealthResponse() (response *ModifyL4HealthResponse) {
    response = &ModifyL4HealthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the health check parameters of a layer-4 forwarding rule. It is supported for Anti-DDoS Advanced and Anti-DDoS Ultimate.
func (c *Client) ModifyL4Health(request *ModifyL4HealthRequest) (response *ModifyL4HealthResponse, err error) {
    if request == nil {
        request = NewModifyL4HealthRequest()
    }
    response = NewModifyL4HealthResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4KeepTimeRequest() (request *ModifyL4KeepTimeRequest) {
    request = &ModifyL4KeepTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyL4KeepTime")
    return
}

func NewModifyL4KeepTimeResponse() (response *ModifyL4KeepTimeResponse) {
    response = &ModifyL4KeepTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the session persistence of a layer-4 forwarding rule. It is supported for Anti-DDoS Advanced and Anti-DDoS Ultimate.
func (c *Client) ModifyL4KeepTime(request *ModifyL4KeepTimeRequest) (response *ModifyL4KeepTimeResponse, err error) {
    if request == nil {
        request = NewModifyL4KeepTimeRequest()
    }
    response = NewModifyL4KeepTimeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL4RulesRequest() (request *ModifyL4RulesRequest) {
    request = &ModifyL4RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyL4Rules")
    return
}

func NewModifyL4RulesResponse() (response *ModifyL4RulesResponse) {
    response = &ModifyL4RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify a layer-4 forwarding rule.
func (c *Client) ModifyL4Rules(request *ModifyL4RulesRequest) (response *ModifyL4RulesResponse, err error) {
    if request == nil {
        request = NewModifyL4RulesRequest()
    }
    response = NewModifyL4RulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyL7RulesRequest() (request *ModifyL7RulesRequest) {
    request = &ModifyL7RulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyL7Rules")
    return
}

func NewModifyL7RulesResponse() (response *ModifyL7RulesResponse) {
    response = &ModifyL7RulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify the layer-7 forwarding rules.
func (c *Client) ModifyL7Rules(request *ModifyL7RulesRequest) (response *ModifyL7RulesResponse, err error) {
    if request == nil {
        request = NewModifyL7RulesRequest()
    }
    response = NewModifyL7RulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetReturnSwitchRequest() (request *ModifyNetReturnSwitchRequest) {
    request = &ModifyNetReturnSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyNetReturnSwitch")
    return
}

func NewModifyNetReturnSwitchResponse() (response *ModifyNetReturnSwitchResponse) {
    response = &ModifyNetReturnSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to switch a client to the real server and set the switch duration when the client is under attack or blocked.
func (c *Client) ModifyNetReturnSwitch(request *ModifyNetReturnSwitchRequest) (response *ModifyNetReturnSwitchResponse, err error) {
    if request == nil {
        request = NewModifyNetReturnSwitchRequest()
    }
    response = NewModifyNetReturnSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNewDomainRulesRequest() (request *ModifyNewDomainRulesRequest) {
    request = &ModifyNewDomainRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyNewDomainRules")
    return
}

func NewModifyNewDomainRulesResponse() (response *ModifyNewDomainRulesResponse) {
    response = &ModifyNewDomainRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify layer-7 forwarding rules.
func (c *Client) ModifyNewDomainRules(request *ModifyNewDomainRulesRequest) (response *ModifyNewDomainRulesResponse, err error) {
    if request == nil {
        request = NewModifyNewDomainRulesRequest()
    }
    response = NewModifyNewDomainRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNewL4RuleRequest() (request *ModifyNewL4RuleRequest) {
    request = &ModifyNewL4RuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyNewL4Rule")
    return
}

func NewModifyNewL4RuleResponse() (response *ModifyNewL4RuleResponse) {
    response = &ModifyNewL4RuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to modify layer-4 forwarding rules.
func (c *Client) ModifyNewL4Rule(request *ModifyNewL4RuleRequest) (response *ModifyNewL4RuleResponse, err error) {
    if request == nil {
        request = NewModifyNewL4RuleRequest()
    }
    response = NewModifyNewL4RuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResBindDDoSPolicyRequest() (request *ModifyResBindDDoSPolicyRequest) {
    request = &ModifyResBindDDoSPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyResBindDDoSPolicy")
    return
}

func NewModifyResBindDDoSPolicyResponse() (response *ModifyResBindDDoSPolicyResponse) {
    response = &ModifyResBindDDoSPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to bind an advanced DDoS policy to an instance.
func (c *Client) ModifyResBindDDoSPolicy(request *ModifyResBindDDoSPolicyRequest) (response *ModifyResBindDDoSPolicyResponse, err error) {
    if request == nil {
        request = NewModifyResBindDDoSPolicyRequest()
    }
    response = NewModifyResBindDDoSPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceRenewFlagRequest() (request *ModifyResourceRenewFlagRequest) {
    request = &ModifyResourceRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dayu", APIVersion, "ModifyResourceRenewFlag")
    return
}

func NewModifyResourceRenewFlagResponse() (response *ModifyResourceRenewFlagResponse) {
    response = &ModifyResourceRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to enable or disable auto-renewal for a resource.
func (c *Client) ModifyResourceRenewFlag(request *ModifyResourceRenewFlagRequest) (response *ModifyResourceRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyResourceRenewFlagRequest()
    }
    response = NewModifyResourceRenewFlagResponse()
    err = c.Send(request, response)
    return
}
