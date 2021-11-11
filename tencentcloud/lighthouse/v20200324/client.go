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

package v20200324

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-03-24"

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


func NewApplyInstanceSnapshotRequest() (request *ApplyInstanceSnapshotRequest) {
    request = &ApplyInstanceSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ApplyInstanceSnapshot")
    
    return
}

func NewApplyInstanceSnapshotResponse() (response *ApplyInstanceSnapshotResponse) {
    response = &ApplyInstanceSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyInstanceSnapshot
// This API is used to roll back the system disk snapshot of the specified instance.
//
// <li>Only rollback to the original system disk is supported.</li>
//
// <li>Only snapshots in `NORMAL` status can be used for rollback. To query the status of a snapshot, you can call the `DescribeSnapshots` API and see the `SnapshotState` field in the response.</li>
//
// <li>When a snapshot is rolled back, the status of the instance must be `STOPPED` or `RUNNING`. You can call the `DescribeInstances` API to query the instance status. Instances in `RUNNING` status will be forcibly shut down before snapshot rollback.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_DISKSIZENOTMATCH = "InvalidParameterValue.DiskSizeNotMatch"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
func (c *Client) ApplyInstanceSnapshot(request *ApplyInstanceSnapshotRequest) (response *ApplyInstanceSnapshotResponse, err error) {
    if request == nil {
        request = NewApplyInstanceSnapshotRequest()
    }
    response = NewApplyInstanceSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateInstancesKeyPairsRequest() (request *AssociateInstancesKeyPairsRequest) {
    request = &AssociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "AssociateInstancesKeyPairs")
    
    return
}

func NewAssociateInstancesKeyPairsResponse() (response *AssociateInstancesKeyPairsResponse) {
    response = &AssociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateInstancesKeyPairs
// This API is used to bind a user-specified key pair to an instance.
//
// * Only instances on LINUX_UNIX in [RUNNING, STOPPED] status are supported. Instances in `RUNNING` status will be forcibly shut down before binding.
//
// * If the public key of a key pair is written to the SSH configuration of the instance, you will be able to log in to the instance with the private key of the key pair.
//
// * If the instance is already associated with a key, the old key will become invalid.
//
// * If you currently use a password to log in, you will no longer be able to do so after you associate the instance with a key.
//
// * Batch operations are supported. The maximum number of instances in each request is 100. If instances not available for the operation are selected, you will get an error code.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDDUPLICATE = "UnsupportedOperation.KeyPairBindDuplicate"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_WINDOWSNOTALLOWTOASSOCIATEKEYPAIR = "UnsupportedOperation.WindowsNotAllowToAssociateKeyPair"
func (c *Client) AssociateInstancesKeyPairs(request *AssociateInstancesKeyPairsRequest) (response *AssociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewAssociateInstancesKeyPairsRequest()
    }
    response = NewAssociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewAttachCcnRequest() (request *AttachCcnRequest) {
    request = &AttachCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "AttachCcn")
    
    return
}

func NewAttachCcnResponse() (response *AttachCcnResponse) {
    response = &AttachCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachCcn
// This API is used to associate a CCN instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ATTACHCCNCONDITIONUNSATISFIED = "UnsupportedOperation.AttachCcnConditionUnsatisfied"
//  UNSUPPORTEDOPERATION_ATTACHCCNFAILED = "UnsupportedOperation.AttachCcnFailed"
//  UNSUPPORTEDOPERATION_CCNALREADYATTACHED = "UnsupportedOperation.CcnAlreadyAttached"
func (c *Client) AttachCcn(request *AttachCcnRequest) (response *AttachCcnResponse, err error) {
    if request == nil {
        request = NewAttachCcnRequest()
    }
    response = NewAttachCcnResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBlueprintRequest() (request *CreateBlueprintRequest) {
    request = &CreateBlueprintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateBlueprint")
    
    return
}

func NewCreateBlueprintResponse() (response *CreateBlueprintResponse) {
    response = &CreateBlueprintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBlueprint
// This API is used to create an image.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEBLUEPRINTFAILED = "FailedOperation.CreateBlueprintFailed"
//  FAILEDOPERATION_UNABLETOCREATEBLUEPRINT = "FailedOperation.UnableToCreateBlueprint"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateBlueprint(request *CreateBlueprintRequest) (response *CreateBlueprintResponse, err error) {
    if request == nil {
        request = NewCreateBlueprintRequest()
    }
    response = NewCreateBlueprintResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFirewallRulesRequest() (request *CreateFirewallRulesRequest) {
    request = &CreateFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateFirewallRules")
    
    return
}

func NewCreateFirewallRulesResponse() (response *CreateFirewallRulesResponse) {
    response = &CreateFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFirewallRules
// This API is used to add a firewall rule on an instance.
//
// 
//
// 
//
// * `FirewallVersion` is the firewall version number. Every time you update the firewall rule version, it will be automatically increased by 1 to prevent the updated rule from expiring. If it is left empty, conflicts will not be considered.
//
// 
//
// In the `FirewallRules` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETER_FIREWALLRULESEXIST = "InvalidParameter.FirewallRulesExist"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  LIMITEXCEEDED_FIREWALLRULESLIMITEXCEEDED = "LimitExceeded.FirewallRulesLimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) CreateFirewallRules(request *CreateFirewallRulesRequest) (response *CreateFirewallRulesResponse, err error) {
    if request == nil {
        request = NewCreateFirewallRulesRequest()
    }
    response = NewCreateFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceSnapshotRequest() (request *CreateInstanceSnapshotRequest) {
    request = &CreateInstanceSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateInstanceSnapshot")
    
    return
}

func NewCreateInstanceSnapshotResponse() (response *CreateInstanceSnapshotResponse) {
    response = &CreateInstanceSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstanceSnapshot
// This API is used to create a system disk snapshot of the specified instance.
//
// error code that may be returned:
//  INTERNALERROR_GETSNAPSHOTALLOCQUOTALOCKERROR = "InternalError.GetSnapshotAllocQuotaLockError"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  LIMITEXCEEDED_SNAPSHOTQUOTALIMITEXCEEDED = "LimitExceeded.SnapshotQuotaLimitExceeded"
//  OPERATIONDENIED_OPERATIONDENIEDCREATESNAPSHOTFORSTORAGEBUNDLE = "OperationDenied.OperationDeniedCreateSnapshotForStorageBundle"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"
//  UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) CreateInstanceSnapshot(request *CreateInstanceSnapshotRequest) (response *CreateInstanceSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateInstanceSnapshotRequest()
    }
    response = NewCreateInstanceSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKeyPairRequest() (request *CreateKeyPairRequest) {
    request = &CreateKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateKeyPair")
    
    return
}

func NewCreateKeyPairResponse() (response *CreateKeyPairResponse) {
    response = &CreateKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateKeyPair
// This API is used to create a key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidKeyPairNameIncludeIllegalChar"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
func (c *Client) CreateKeyPair(request *CreateKeyPairRequest) (response *CreateKeyPairResponse, err error) {
    if request == nil {
        request = NewCreateKeyPairRequest()
    }
    response = NewCreateKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlueprintsRequest() (request *DeleteBlueprintsRequest) {
    request = &DeleteBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteBlueprints")
    
    return
}

func NewDeleteBlueprintsResponse() (response *DeleteBlueprintsResponse) {
    response = &DeleteBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBlueprints
// This API is used to delete an image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"
//  UNSUPPORTEDOPERATION_BLUEPRINTOCCUPIED = "UnsupportedOperation.BlueprintOccupied"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
//  UNSUPPORTEDOPERATION_POSTDESTROYRESOURCEFAILED = "UnsupportedOperation.PostDestroyResourceFailed"
func (c *Client) DeleteBlueprints(request *DeleteBlueprintsRequest) (response *DeleteBlueprintsResponse, err error) {
    if request == nil {
        request = NewDeleteBlueprintsRequest()
    }
    response = NewDeleteBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFirewallRulesRequest() (request *DeleteFirewallRulesRequest) {
    request = &DeleteFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteFirewallRules")
    
    return
}

func NewDeleteFirewallRulesResponse() (response *DeleteFirewallRulesResponse) {
    response = &DeleteFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFirewallRules
// This API is used to delete a firewall rule of an instance.
//
// 
//
// * `FirewallVersion` is used to specify the version of the firewall to be manipulated. If the `FirewallVersion` value passed in is not equal to the current latest version of the firewall, a failure will be returned. If `FirewallVersion` is not passed in, the specified rule will be deleted directly.
//
// 
//
// In the `FirewallRules` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) DeleteFirewallRules(request *DeleteFirewallRulesRequest) (response *DeleteFirewallRulesResponse, err error) {
    if request == nil {
        request = NewDeleteFirewallRulesRequest()
    }
    response = NewDeleteFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKeyPairsRequest() (request *DeleteKeyPairsRequest) {
    request = &DeleteKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteKeyPairs")
    
    return
}

func NewDeleteKeyPairsResponse() (response *DeleteKeyPairsResponse) {
    response = &DeleteKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteKeyPairs
// This API is used to delete a key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  RESOURCEINUSE_KEYPAIRINUSE = "ResourceInUse.KeyPairInUse"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNSUPPORTEDOPERATION_KEYPAIRBINDTOBLUEPRINTS = "UnsupportedOperation.KeyPairBindToBlueprints"
func (c *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (response *DeleteKeyPairsResponse, err error) {
    if request == nil {
        request = NewDeleteKeyPairsRequest()
    }
    response = NewDeleteKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteSnapshots")
    
    return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
    response = &DeleteSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSnapshots
// This API is used to delete a snapshot.
//
// The snapshot must be in `NORMAL` status. To query the status of a snapshot, you can call the `DescribeSnapshots` API and see the `SnapshotState` field in the response.
//
// error code that may be returned:
//  FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
//  RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"
//  UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"
//  UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotsRequest()
    }
    response = NewDeleteSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlueprintInstancesRequest() (request *DescribeBlueprintInstancesRequest) {
    request = &DescribeBlueprintInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBlueprintInstances")
    
    return
}

func NewDescribeBlueprintInstancesResponse() (response *DescribeBlueprintInstancesResponse) {
    response = &DescribeBlueprintInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlueprintInstances
// This API is used to query the information of an image instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBlueprintInstances(request *DescribeBlueprintInstancesRequest) (response *DescribeBlueprintInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeBlueprintInstancesRequest()
    }
    response = NewDescribeBlueprintInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlueprintsRequest() (request *DescribeBlueprintsRequest) {
    request = &DescribeBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBlueprints")
    
    return
}

func NewDescribeBlueprintsResponse() (response *DescribeBlueprintsResponse) {
    response = &DescribeBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBlueprints
// This API is used to query the information of an image.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
func (c *Client) DescribeBlueprints(request *DescribeBlueprintsRequest) (response *DescribeBlueprintsResponse, err error) {
    if request == nil {
        request = NewDescribeBlueprintsRequest()
    }
    response = NewDescribeBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBundlesRequest() (request *DescribeBundlesRequest) {
    request = &DescribeBundlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBundles")
    
    return
}

func NewDescribeBundlesResponse() (response *DescribeBundlesResponse) {
    response = &DescribeBundlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBundles
// This API is used to query the information of a package.
//
// error code that may be returned:
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INTERNALERROR_REQUESTERROR = "InternalError.RequestError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"
func (c *Client) DescribeBundles(request *DescribeBundlesRequest) (response *DescribeBundlesResponse, err error) {
    if request == nil {
        request = NewDescribeBundlesRequest()
    }
    response = NewDescribeBundlesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnAttachedInstancesRequest() (request *DescribeCcnAttachedInstancesRequest) {
    request = &DescribeCcnAttachedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeCcnAttachedInstances")
    
    return
}

func NewDescribeCcnAttachedInstancesResponse() (response *DescribeCcnAttachedInstancesResponse) {
    response = &DescribeCcnAttachedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcnAttachedInstances
// This API is used to query the information of instances associated with CCN.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_DESCRIBECCNATTACHEDINSTANCESFAILED = "UnsupportedOperation.DescribeCcnAttachedInstancesFailed"
func (c *Client) DescribeCcnAttachedInstances(request *DescribeCcnAttachedInstancesRequest) (response *DescribeCcnAttachedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnAttachedInstancesRequest()
    }
    response = NewDescribeCcnAttachedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirewallRulesRequest() (request *DescribeFirewallRulesRequest) {
    request = &DescribeFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeFirewallRules")
    
    return
}

func NewDescribeFirewallRulesResponse() (response *DescribeFirewallRulesResponse) {
    response = &DescribeFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirewallRules
// This API is used to query the firewall rules of an instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeFirewallRules(request *DescribeFirewallRulesRequest) (response *DescribeFirewallRulesResponse, err error) {
    if request == nil {
        request = NewDescribeFirewallRulesRequest()
    }
    response = NewDescribeFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirewallRulesTemplateRequest() (request *DescribeFirewallRulesTemplateRequest) {
    request = &DescribeFirewallRulesTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeFirewallRulesTemplate")
    
    return
}

func NewDescribeFirewallRulesTemplateResponse() (response *DescribeFirewallRulesTemplateResponse) {
    response = &DescribeFirewallRulesTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFirewallRulesTemplate
// This API is used to query a firewall rule template.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeFirewallRulesTemplate(request *DescribeFirewallRulesTemplateRequest) (response *DescribeFirewallRulesTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeFirewallRulesTemplateRequest()
    }
    response = NewDescribeFirewallRulesTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGeneralResourceQuotasRequest() (request *DescribeGeneralResourceQuotasRequest) {
    request = &DescribeGeneralResourceQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeGeneralResourceQuotas")
    
    return
}

func NewDescribeGeneralResourceQuotasResponse() (response *DescribeGeneralResourceQuotasResponse) {
    response = &DescribeGeneralResourceQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGeneralResourceQuotas
// This API is used to query the quota information of general resources.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INVALIDRESOURCEQUOTARESOURCENAME = "InvalidParameterValue.InvalidResourceQuotaResourceName"
func (c *Client) DescribeGeneralResourceQuotas(request *DescribeGeneralResourceQuotasRequest) (response *DescribeGeneralResourceQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeGeneralResourceQuotasRequest()
    }
    response = NewDescribeGeneralResourceQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLoginKeyPairAttributeRequest() (request *DescribeInstanceLoginKeyPairAttributeRequest) {
    request = &DescribeInstanceLoginKeyPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstanceLoginKeyPairAttribute")
    
    return
}

func NewDescribeInstanceLoginKeyPairAttributeResponse() (response *DescribeInstanceLoginKeyPairAttributeResponse) {
    response = &DescribeInstanceLoginKeyPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceLoginKeyPairAttribute
// This API is used to query the attributes of the default login key of an instance.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeInstanceLoginKeyPairAttribute(request *DescribeInstanceLoginKeyPairAttributeRequest) (response *DescribeInstanceLoginKeyPairAttributeResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLoginKeyPairAttributeRequest()
    }
    response = NewDescribeInstanceLoginKeyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceVncUrlRequest() (request *DescribeInstanceVncUrlRequest) {
    request = &DescribeInstanceVncUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstanceVncUrl")
    
    return
}

func NewDescribeInstanceVncUrlResponse() (response *DescribeInstanceVncUrlResponse) {
    response = &DescribeInstanceVncUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceVncUrl
// This API is used to query the VNC URL of an instance, and the obtained address can be used for VNC login to the instance.
//
// 
//
// * This feature is available to instances in `RUNNING` status.
//
// * A VNC URL is only valid for 15 seconds. If you do not access the URL within 15 seconds, it will become invalid, and you will have to query another one.
//
// * Once the VNC URL is accessed, it will become invalid, and you will have to query another one if needed.
//
// * If the connection is interrupted, up to 30 reconnection requests per minute are allowed.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (response *DescribeInstanceVncUrlResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceVncUrlRequest()
    }
    response = NewDescribeInstanceVncUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstances")
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// This API is used to query the details of one or multiple instances.
//
// 
//
// * You can query the details of an instance according to its ID, name, or private IP.
//
// * For more information on filters, please see [Filters](https://intl.cloud.tencent.com/document/product/1207/47576?from_cn_redirect=1#Filter).
//
// * If no parameter is defined, the status of a certain number of instances under the current account will be returned. The number is specified by `Limit` and is 20 by default.
//
// * The latest operation (LatestOperation) and the latest operation status (LatestOperationState) of the instance can be queried.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDeniedActionsRequest() (request *DescribeInstancesDeniedActionsRequest) {
    request = &DescribeInstancesDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesDeniedActions")
    
    return
}

func NewDescribeInstancesDeniedActionsResponse() (response *DescribeInstancesDeniedActionsResponse) {
    response = &DescribeInstancesDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDeniedActions
// This API is used to query the list of operation limits of one or more instances.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeInstancesDeniedActions(request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDeniedActionsRequest()
    }
    response = NewDescribeInstancesDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesReturnableRequest() (request *DescribeInstancesReturnableRequest) {
    request = &DescribeInstancesReturnableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesReturnable")
    
    return
}

func NewDescribeInstancesReturnableResponse() (response *DescribeInstancesReturnableResponse) {
    response = &DescribeInstancesReturnableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesReturnable
// This API is used to query whether the specified instance can be returned.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESRETURNABLEERROR = "InternalError.DescribeInstancesReturnableError"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeInstancesReturnable(request *DescribeInstancesReturnableRequest) (response *DescribeInstancesReturnableResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesReturnableRequest()
    }
    response = NewDescribeInstancesReturnableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesTrafficPackagesRequest() (request *DescribeInstancesTrafficPackagesRequest) {
    request = &DescribeInstancesTrafficPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesTrafficPackages")
    
    return
}

func NewDescribeInstancesTrafficPackagesResponse() (response *DescribeInstancesTrafficPackagesResponse) {
    response = &DescribeInstancesTrafficPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesTrafficPackages
// This API is used to query the traffic package details of one or more instances.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTRAFFICPACKAGESFAILED = "InternalError.DescribeInstancesTrafficPackagesFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
func (c *Client) DescribeInstancesTrafficPackages(request *DescribeInstancesTrafficPackagesRequest) (response *DescribeInstancesTrafficPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesTrafficPackagesRequest()
    }
    response = NewDescribeInstancesTrafficPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeyPairsRequest() (request *DescribeKeyPairsRequest) {
    request = &DescribeKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeKeyPairs")
    
    return
}

func NewDescribeKeyPairsResponse() (response *DescribeKeyPairsResponse) {
    response = &DescribeKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKeyPairs
// This API is used to query key pairs.
//
// error code that may be returned:
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
func (c *Client) DescribeKeyPairs(request *DescribeKeyPairsRequest) (response *DescribeKeyPairsResponse, err error) {
    if request == nil {
        request = NewDescribeKeyPairsRequest()
    }
    response = NewDescribeKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModifyInstanceBundlesRequest() (request *DescribeModifyInstanceBundlesRequest) {
    request = &DescribeModifyInstanceBundlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeModifyInstanceBundles")
    
    return
}

func NewDescribeModifyInstanceBundlesResponse() (response *DescribeModifyInstanceBundlesResponse) {
    response = &DescribeModifyInstanceBundlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModifyInstanceBundles
// This API is used to query the list of package options of an instance.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATION = "InternalError.DescribeInstancesModification"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATIONERROR = "InternalError.DescribeInstancesModificationError"
//  INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeModifyInstanceBundles(request *DescribeModifyInstanceBundlesRequest) (response *DescribeModifyInstanceBundlesResponse, err error) {
    if request == nil {
        request = NewDescribeModifyInstanceBundlesRequest()
    }
    response = NewDescribeModifyInstanceBundlesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeRegions")
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// This API is used to query the information of regions.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATION = "InternalError.DescribeInstancesModification"
//  INTERNALERROR_DESCRIBEINSTANCESMODIFICATIONERROR = "InternalError.DescribeInstancesModificationError"
//  INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResetInstanceBlueprintsRequest() (request *DescribeResetInstanceBlueprintsRequest) {
    request = &DescribeResetInstanceBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeResetInstanceBlueprints")
    
    return
}

func NewDescribeResetInstanceBlueprintsResponse() (response *DescribeResetInstanceBlueprintsResponse) {
    response = &DescribeResetInstanceBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResetInstanceBlueprints
// This API is used to query the image information of a reset instance.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) DescribeResetInstanceBlueprints(request *DescribeResetInstanceBlueprintsRequest) (response *DescribeResetInstanceBlueprintsResponse, err error) {
    if request == nil {
        request = NewDescribeResetInstanceBlueprintsRequest()
    }
    response = NewDescribeResetInstanceBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeSnapshots")
    
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshots
// This API is used to query the list of snapshots.
//
// error code that may be returned:
//  INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"
//  INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"
//  INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"
//  INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"
//  INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"
//  INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsDeniedActionsRequest() (request *DescribeSnapshotsDeniedActionsRequest) {
    request = &DescribeSnapshotsDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeSnapshotsDeniedActions")
    
    return
}

func NewDescribeSnapshotsDeniedActionsResponse() (response *DescribeSnapshotsDeniedActionsResponse) {
    response = &DescribeSnapshotsDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotsDeniedActions
// This API is used to query the list of operation limits of one or more snapshots.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
func (c *Client) DescribeSnapshotsDeniedActions(request *DescribeSnapshotsDeniedActionsRequest) (response *DescribeSnapshotsDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsDeniedActionsRequest()
    }
    response = NewDescribeSnapshotsDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeZones")
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZones
// This API is used to query the list of AZs in a region.
//
// error code that may be returned:
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDetachCcnRequest() (request *DetachCcnRequest) {
    request = &DetachCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DetachCcn")
    
    return
}

func NewDetachCcnResponse() (response *DetachCcnResponse) {
    response = &DetachCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachCcn
// This API is used to unassociate a CCN instance.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_DETACHCCNFAILED = "UnsupportedOperation.DetachCcnFailed"
func (c *Client) DetachCcn(request *DetachCcnRequest) (response *DetachCcnResponse, err error) {
    if request == nil {
        request = NewDetachCcnRequest()
    }
    response = NewDetachCcnResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateInstancesKeyPairsRequest() (request *DisassociateInstancesKeyPairsRequest) {
    request = &DisassociateInstancesKeyPairsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DisassociateInstancesKeyPairs")
    
    return
}

func NewDisassociateInstancesKeyPairsResponse() (response *DisassociateInstancesKeyPairsResponse) {
    response = &DisassociateInstancesKeyPairsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateInstancesKeyPairs
// This API is used to unbind an instance from the specified key pair.
//
// 
//
// * Only instances on LINUX_UNIX in [RUNNING, STOPPED] status are supported. Instances in `RUNNING` status will be forcibly shut down before unbinding.
//
// * After a key pair is unassociated from an instance, you can log in to the instance with password.
//
// * If no password was set, you cannot log in to the instance with SSH after unbinding. You can call the ResetInstancesPassword API to set a login password.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_KEYPAIRNOTBOUNDTOINSTANCE = "UnsupportedOperation.KeyPairNotBoundToInstance"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) DisassociateInstancesKeyPairs(request *DisassociateInstancesKeyPairsRequest) (response *DisassociateInstancesKeyPairsResponse, err error) {
    if request == nil {
        request = NewDisassociateInstancesKeyPairsRequest()
    }
    response = NewDisassociateInstancesKeyPairsResponse()
    err = c.Send(request, response)
    return
}

func NewImportKeyPairRequest() (request *ImportKeyPairRequest) {
    request = &ImportKeyPairRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ImportKeyPair")
    
    return
}

func NewImportKeyPairResponse() (response *ImportKeyPairResponse) {
    response = &ImportKeyPairResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportKeyPair
// This API is used to import the specified key pair.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"
//  FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"
//  FAILEDOPERATION_IMPORTKEYPAIRFAILED = "FailedOperation.ImportKeyPairFailed"
//  INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"
//  INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYDUPLICATED = "InvalidParameterValue.KeyPairPublicKeyDuplicated"
//  INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYMALFORMED = "InvalidParameterValue.KeyPairPublicKeyMalformed"
//  LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"
func (c *Client) ImportKeyPair(request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
    if request == nil {
        request = NewImportKeyPairRequest()
    }
    response = NewImportKeyPairResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateBlueprintRequest() (request *InquirePriceCreateBlueprintRequest) {
    request = &InquirePriceCreateBlueprintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceCreateBlueprint")
    
    return
}

func NewInquirePriceCreateBlueprintResponse() (response *InquirePriceCreateBlueprintResponse) {
    response = &InquirePriceCreateBlueprintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceCreateBlueprint
// This API is used to query the price of a created image.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
func (c *Client) InquirePriceCreateBlueprint(request *InquirePriceCreateBlueprintRequest) (response *InquirePriceCreateBlueprintResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateBlueprintRequest()
    }
    response = NewInquirePriceCreateBlueprintResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateInstancesRequest() (request *InquirePriceCreateInstancesRequest) {
    request = &InquirePriceCreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceCreateInstances")
    
    return
}

func NewInquirePriceCreateInstancesResponse() (response *InquirePriceCreateInstancesResponse) {
    response = &InquirePriceCreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceCreateInstances
// This API is used to query the price of a created instance.
//
// error code that may be returned:
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
func (c *Client) InquirePriceCreateInstances(request *InquirePriceCreateInstancesRequest) (response *InquirePriceCreateInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateInstancesRequest()
    }
    response = NewInquirePriceCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewInstancesRequest() (request *InquirePriceRenewInstancesRequest) {
    request = &InquirePriceRenewInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "InquirePriceRenewInstances")
    
    return
}

func NewInquirePriceRenewInstancesResponse() (response *InquirePriceRenewInstancesResponse) {
    response = &InquirePriceRenewInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquirePriceRenewInstances
// This API is used to query the price of renewed instance.
//
// error code that may be returned:
//  INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
func (c *Client) InquirePriceRenewInstances(request *InquirePriceRenewInstancesRequest) (response *InquirePriceRenewInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewInstancesRequest()
    }
    response = NewInquirePriceRenewInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlueprintAttributeRequest() (request *ModifyBlueprintAttributeRequest) {
    request = &ModifyBlueprintAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyBlueprintAttribute")
    
    return
}

func NewModifyBlueprintAttributeResponse() (response *ModifyBlueprintAttributeResponse) {
    response = &ModifyBlueprintAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBlueprintAttribute
// This API is used to modify an image attribute.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"
func (c *Client) ModifyBlueprintAttribute(request *ModifyBlueprintAttributeRequest) (response *ModifyBlueprintAttributeResponse, err error) {
    if request == nil {
        request = NewModifyBlueprintAttributeRequest()
    }
    response = NewModifyBlueprintAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFirewallRuleDescriptionRequest() (request *ModifyFirewallRuleDescriptionRequest) {
    request = &ModifyFirewallRuleDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyFirewallRuleDescription")
    
    return
}

func NewModifyFirewallRuleDescriptionResponse() (response *ModifyFirewallRuleDescriptionResponse) {
    response = &ModifyFirewallRuleDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFirewallRuleDescription
// This API is used to modify the description of a single firewall rule.
//
// 
//
// * `FirewallVersion` is used to specify the version of the firewall to be manipulated. If the `FirewallVersion` value passed in is not equal to the current latest version of the firewall, a failure will be returned. If `FirewallVersion` is not passed in, the description of the specified rule will be modified directly.
//
// 
//
// In the `FirewallRule` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRuleDescription(request *ModifyFirewallRuleDescriptionRequest) (response *ModifyFirewallRuleDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyFirewallRuleDescriptionRequest()
    }
    response = NewModifyFirewallRuleDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFirewallRulesRequest() (request *ModifyFirewallRulesRequest) {
    request = &ModifyFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyFirewallRules")
    
    return
}

func NewModifyFirewallRulesResponse() (response *ModifyFirewallRulesResponse) {
    response = &ModifyFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyFirewallRules
// This API is used to reset the firewall rules of an instance.
//
// 
//
// This API deletes all firewall rules of the current instance first and then adds new rules.
//
// 
//
// * `FirewallVersion` is used to specify the version of the firewall to be manipulated. If the `FirewallVersion` value passed in is not equal to the current latest version of the firewall, a failure will be returned. If `FirewallVersion` is not passed in, the specified rule will be reset directly.
//
// 
//
// In the `FirewallRules` parameter:
//
// * Valid values of the `Protocol` field include `TCP`, `UDP`, `ICMP`, and `ALL`.
//
// * For the `Port` field, you can enter only `ALL`, a single port number, several port numbers separated by commas, or a port range indicated by two port numbers separated by a minus sign. If `Port` is a range, the port number on the left of the minus sign must be smaller than the one on the right. If `Protocol` is not `TCP` or `UDP`, `Port` can only be empty or `ALL`. The length of the `Port` field cannot exceed 64 characters.
//
// * For the `CidrBlock` field, you can enter any string that conforms to the CIDR format standard. Multi-Tenant network isolation rules take precedence over private network rules in the firewall.
//
// * For the `Action` field, you can enter only `ACCEPT` or `DROP`.
//
// * The length of the `FirewallRuleDescription` field cannot exceed 64 characters.
//
// error code that may be returned:
//  FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"
//  INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"
//  UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"
func (c *Client) ModifyFirewallRules(request *ModifyFirewallRulesRequest) (response *ModifyFirewallRulesResponse, err error) {
    if request == nil {
        request = NewModifyFirewallRulesRequest()
    }
    response = NewModifyFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesAttribute")
    
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesAttribute
// This API is used to modify the attributes of instances.
//
// * The instance name is used only for users’ convenience.
//
// * Batch operations are supported. Each request can contain up to 100 instances at a time.
//
// * This API is async. A successful request will return a `RequestId`, it does not mean the operation is completed. You can call the `DescribeInstances` API to query the operation result. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETER_ONLYALLOWMODIFYONEATTRIBUTE = "InvalidParameter.OnlyAllowModifyOneAttribute"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesLoginKeyPairAttributeRequest() (request *ModifyInstancesLoginKeyPairAttributeRequest) {
    request = &ModifyInstancesLoginKeyPairAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesLoginKeyPairAttribute")
    
    return
}

func NewModifyInstancesLoginKeyPairAttributeResponse() (response *ModifyInstancesLoginKeyPairAttributeResponse) {
    response = &ModifyInstancesLoginKeyPairAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesLoginKeyPairAttribute
// This API is used to set the attributes of the default login key pair of an instance.
//
// 
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCELOGINKEYPAIRPERMITLOGIN = "InvalidParameterValue.InvalidInstanceLoginKeyPairPermitLogin"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesLoginKeyPairAttribute(request *ModifyInstancesLoginKeyPairAttributeRequest) (response *ModifyInstancesLoginKeyPairAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesLoginKeyPairAttributeRequest()
    }
    response = NewModifyInstancesLoginKeyPairAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesRenewFlagRequest() (request *ModifyInstancesRenewFlagRequest) {
    request = &ModifyInstancesRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyInstancesRenewFlag")
    
    return
}

func NewModifyInstancesRenewFlagResponse() (response *ModifyInstancesRenewFlagResponse) {
    response = &ModifyInstancesRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancesRenewFlag
// This API is used to modify the renewal flags of monthly subscribed instances.
//
// 
//
// * Instances marked with "auto-renewal" will be automatically renewed for one month when they expire.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ModifyInstancesRenewFlag(request *ModifyInstancesRenewFlagRequest) (response *ModifyInstancesRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyInstancesRenewFlagRequest()
    }
    response = NewModifyInstancesRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotAttributeRequest() (request *ModifySnapshotAttributeRequest) {
    request = &ModifySnapshotAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifySnapshotAttribute")
    
    return
}

func NewModifySnapshotAttributeResponse() (response *ModifySnapshotAttributeResponse) {
    response = &ModifySnapshotAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySnapshotAttribute
// This API is used to modify the attributes of a snapshot.
//
// <li>The snapshot name is used only for users’ convenience.</li>
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"
//  INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"
//  RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"
func (c *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
    if request == nil {
        request = NewModifySnapshotAttributeRequest()
    }
    response = NewModifySnapshotAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "RebootInstances")
    
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RebootInstances
// This API is used to restart instances.
//
// 
//
// * You can only perform this operation on instances whose status is `RUNNING`.
//
// * The instance status will become `REBOOTING` when the API is called successfully and will become `RUNNING` when the instance is successfully restarted.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetAttachCcnRequest() (request *ResetAttachCcnRequest) {
    request = &ResetAttachCcnRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetAttachCcn")
    
    return
}

func NewResetAttachCcnResponse() (response *ResetAttachCcnResponse) {
    response = &ResetAttachCcnResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAttachCcn
// This API is used to apply for association again after a CCN instance association application expires.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"
//  UNSUPPORTEDOPERATION_RESETATTACHCCNFAILED = "UnsupportedOperation.ResetAttachCcnFailed"
func (c *Client) ResetAttachCcn(request *ResetAttachCcnRequest) (response *ResetAttachCcnResponse, err error) {
    if request == nil {
        request = NewResetAttachCcnRequest()
    }
    response = NewResetAttachCcnResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
    request = &ResetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetInstance")
    
    return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
    response = &ResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstance
// This API is used to reinstall the image on the specified instance.
//
// 
//
// * If you specify a `BlueprintId`, the specified image is used; otherwise, the image used by the current instance is used.
//
// * The system disk will be formatted and reset. Therefore, make sure that no important files are stored on the system disk.
//
// * Currently, this API does not support switching the operating system between `LINUX_UNIX` and `WINDOWS` for instances.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_BLUEPRINTCONFIGNOTMATCH = "InvalidParameterValue.BlueprintConfigNotMatch"
//  INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"
//  INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    if request == nil {
        request = NewResetInstanceRequest()
    }
    response = NewResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesPasswordRequest() (request *ResetInstancesPasswordRequest) {
    request = &ResetInstancesPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetInstancesPassword")
    
    return
}

func NewResetInstancesPasswordResponse() (response *ResetInstancesPasswordResponse) {
    response = &ResetInstancesPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetInstancesPassword
// This API is used to reset the password of the instance OS to a user-specified password.
//
// * You can only use this API to modify the password of the admin account. The name of the admin account varies by OS (on Windows, it is `Administrator`; on Ubuntu, it is `ubuntu`; on other systems, it is `root`).
//
// * Batch operations are supported. You can reset the passwords of multiple instances to the same one. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) ResetInstancesPassword(request *ResetInstancesPasswordRequest) (response *ResetInstancesPasswordResponse, err error) {
    if request == nil {
        request = NewResetInstancesPasswordRequest()
    }
    response = NewResetInstancesPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "StartInstances")
    
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartInstances
// This API is used to start one or more instances.
//
// 
//
// * You can only perform this operation on instances whose status is `STOPPED`.
//
// * The instance status will become `STARTING` when the API is called successfully and will become `RUNNING` when the instance is successfully started.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
    request = &StopInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "StopInstances")
    
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopInstances
// This API is used to shut down one or more instances.
//
// * You can only perform this operation on instances whose status is `RUNNING`.
//
// * The instance status will become `STOPPING` when the API is called successfully and will become `STOPPED` when the instance is successfully shut down.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"
//  INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"
//  INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "TerminateInstances")
    
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateInstances
// This API is used to return instances.
//
// 
//
// * Instances in `SHUTDOWN` status can be terminated through this API and cannot be recovered.
//
// * Batch operations are supported. The maximum number of instances in each request is 100.
//
// * This API is async. After the request is sent successfully, a `RequestId` will be returned. At this time, the operation is not completed immediately. The result of the instance operation can be queried by calling the `DescribeInstances` API. If the latest operation status (LatestOperationState) of the instance is `SUCCESS`, the operation is successful.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"
//  OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"
//  RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"
//  UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}