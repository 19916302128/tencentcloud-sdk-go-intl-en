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

package v20210331

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-03-31"

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


func NewAddOrganizationNodeRequest() (request *AddOrganizationNodeRequest) {
    request = &AddOrganizationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddOrganizationNode")
    
    
    return
}

func NewAddOrganizationNodeResponse() (response *AddOrganizationNodeResponse) {
    response = &AddOrganizationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddOrganizationNode
// This API is used to add an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) AddOrganizationNode(request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
    return c.AddOrganizationNodeWithContext(context.Background(), request)
}

// AddOrganizationNode
// This API is used to add an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) AddOrganizationNodeWithContext(ctx context.Context, request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
    if request == nil {
        request = NewAddOrganizationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddOrganizationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddOrganizationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewBindOrganizationMemberAuthAccountRequest() (request *BindOrganizationMemberAuthAccountRequest) {
    request = &BindOrganizationMemberAuthAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationMemberAuthAccount")
    
    
    return
}

func NewBindOrganizationMemberAuthAccountResponse() (response *BindOrganizationMemberAuthAccountResponse) {
    response = &BindOrganizationMemberAuthAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindOrganizationMemberAuthAccount
// This API is used to bind an organization member to a sub-account of the organization admin.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTIDENTITYEXIST = "FailedOperation.SubAccountIdentityExist"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindOrganizationMemberAuthAccount(request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
    return c.BindOrganizationMemberAuthAccountWithContext(context.Background(), request)
}

// BindOrganizationMemberAuthAccount
// This API is used to bind an organization member to a sub-account of the organization admin.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTIDENTITYEXIST = "FailedOperation.SubAccountIdentityExist"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindOrganizationMemberAuthAccountWithContext(ctx context.Context, request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
    if request == nil {
        request = NewBindOrganizationMemberAuthAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindOrganizationMemberAuthAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCancelOrganizationMemberAuthAccountRequest() (request *CancelOrganizationMemberAuthAccountRequest) {
    request = &CancelOrganizationMemberAuthAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationMemberAuthAccount")
    
    
    return
}

func NewCancelOrganizationMemberAuthAccountResponse() (response *CancelOrganizationMemberAuthAccountResponse) {
    response = &CancelOrganizationMemberAuthAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelOrganizationMemberAuthAccount
// This API is used to unbind an organization member from a sub-account of the organization admin.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CancelOrganizationMemberAuthAccount(request *CancelOrganizationMemberAuthAccountRequest) (response *CancelOrganizationMemberAuthAccountResponse, err error) {
    return c.CancelOrganizationMemberAuthAccountWithContext(context.Background(), request)
}

// CancelOrganizationMemberAuthAccount
// This API is used to unbind an organization member from a sub-account of the organization admin.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CancelOrganizationMemberAuthAccountWithContext(ctx context.Context, request *CancelOrganizationMemberAuthAccountRequest) (response *CancelOrganizationMemberAuthAccountResponse, err error) {
    if request == nil {
        request = NewCancelOrganizationMemberAuthAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelOrganizationMemberAuthAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrgServiceAssignRequest() (request *CreateOrgServiceAssignRequest) {
    request = &CreateOrgServiceAssignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrgServiceAssign")
    
    
    return
}

func NewCreateOrgServiceAssignResponse() (response *CreateOrgServiceAssignResponse) {
    response = &CreateOrgServiceAssignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrgServiceAssign
// This API is used to add a delegated admin of the organization service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEORGSERVICEASSIGNOVERLIMIT = "LimitExceeded.CreateOrgServiceAssignOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrgServiceAssign(request *CreateOrgServiceAssignRequest) (response *CreateOrgServiceAssignResponse, err error) {
    return c.CreateOrgServiceAssignWithContext(context.Background(), request)
}

// CreateOrgServiceAssign
// This API is used to add a delegated admin of the organization service.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEORGSERVICEASSIGNOVERLIMIT = "LimitExceeded.CreateOrgServiceAssignOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrgServiceAssignWithContext(ctx context.Context, request *CreateOrgServiceAssignRequest) (response *CreateOrgServiceAssignResponse, err error) {
    if request == nil {
        request = NewCreateOrgServiceAssignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrgServiceAssign require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrgServiceAssignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMemberRequest() (request *CreateOrganizationMemberRequest) {
    request = &CreateOrganizationMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMember")
    
    
    return
}

func NewCreateOrganizationMemberResponse() (response *CreateOrganizationMemberResponse) {
    response = &CreateOrganizationMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationMember
// This API is used to create an organization member.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//  FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//  FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
func (c *Client) CreateOrganizationMember(request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
    return c.CreateOrganizationMemberWithContext(context.Background(), request)
}

// CreateOrganizationMember
// This API is used to create an organization member.
//
// error code that may be returned:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//  FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//  FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
func (c *Client) CreateOrganizationMemberWithContext(ctx context.Context, request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMemberPolicyRequest() (request *CreateOrganizationMemberPolicyRequest) {
    request = &CreateOrganizationMemberPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMemberPolicy")
    
    
    return
}

func NewCreateOrganizationMemberPolicyResponse() (response *CreateOrganizationMemberPolicyResponse) {
    response = &CreateOrganizationMemberPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationMemberPolicy
// This API is used to create an organization member access policy.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CreateOrganizationMemberPolicy(request *CreateOrganizationMemberPolicyRequest) (response *CreateOrganizationMemberPolicyResponse, err error) {
    return c.CreateOrganizationMemberPolicyWithContext(context.Background(), request)
}

// CreateOrganizationMemberPolicy
// This API is used to create an organization member access policy.
//
// error code that may be returned:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CreateOrganizationMemberPolicyWithContext(ctx context.Context, request *CreateOrganizationMemberPolicyRequest) (response *CreateOrganizationMemberPolicyResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMemberPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrgServiceAssignRequest() (request *DeleteOrgServiceAssignRequest) {
    request = &DeleteOrgServiceAssignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrgServiceAssign")
    
    
    return
}

func NewDeleteOrgServiceAssignResponse() (response *DeleteOrgServiceAssignResponse) {
    response = &DeleteOrgServiceAssignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrgServiceAssign
// This API is used to delete a delegated admin of the organization service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  FAILEDOPERATION_ORGANIZATIONSERVICEASSIGNISUSE = "FailedOperation.OrganizationServiceAssignIsUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DeleteOrgServiceAssign(request *DeleteOrgServiceAssignRequest) (response *DeleteOrgServiceAssignResponse, err error) {
    return c.DeleteOrgServiceAssignWithContext(context.Background(), request)
}

// DeleteOrgServiceAssign
// This API is used to delete a delegated admin of the organization service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  FAILEDOPERATION_ORGANIZATIONSERVICEASSIGNISUSE = "FailedOperation.OrganizationServiceAssignIsUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DeleteOrgServiceAssignWithContext(ctx context.Context, request *DeleteOrgServiceAssignRequest) (response *DeleteOrgServiceAssignResponse, err error) {
    if request == nil {
        request = NewDeleteOrgServiceAssignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrgServiceAssign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrgServiceAssignResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationMembersRequest() (request *DeleteOrganizationMembersRequest) {
    request = &DeleteOrganizationMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMembers")
    
    
    return
}

func NewDeleteOrganizationMembersResponse() (response *DeleteOrganizationMembersResponse) {
    response = &DeleteOrganizationMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationMembers
// This API is used to batch delete organization members.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
func (c *Client) DeleteOrganizationMembers(request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
    return c.DeleteOrganizationMembersWithContext(context.Background(), request)
}

// DeleteOrganizationMembers
// This API is used to batch delete organization members.
//
// error code that may be returned:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
func (c *Client) DeleteOrganizationMembersWithContext(ctx context.Context, request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationNodesRequest() (request *DeleteOrganizationNodesRequest) {
    request = &DeleteOrganizationNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationNodes")
    
    
    return
}

func NewDeleteOrganizationNodesResponse() (response *DeleteOrganizationNodesResponse) {
    response = &DeleteOrganizationNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationNodes
// This API is used to batch delete organization nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONNODEDELETEOVERLIMIT = "FailedOperation.OrganizationNodeDeleteOverLimit"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEMPTY = "FailedOperation.OrganizationNodeNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationNodes(request *DeleteOrganizationNodesRequest) (response *DeleteOrganizationNodesResponse, err error) {
    return c.DeleteOrganizationNodesWithContext(context.Background(), request)
}

// DeleteOrganizationNodes
// This API is used to batch delete organization nodes.
//
// error code that may be returned:
//  FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONNODEDELETEOVERLIMIT = "FailedOperation.OrganizationNodeDeleteOverLimit"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEMPTY = "FailedOperation.OrganizationNodeNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationNodesWithContext(ctx context.Context, request *DeleteOrganizationNodesRequest) (response *DeleteOrganizationNodesResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationRequest() (request *DescribeOrganizationRequest) {
    request = &DescribeOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganization")
    
    
    return
}

func NewDescribeOrganizationResponse() (response *DescribeOrganizationResponse) {
    response = &DescribeOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganization
// This API is used to get the organization information.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganization(request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    return c.DescribeOrganizationWithContext(context.Background(), request)
}

// DescribeOrganization
// This API is used to get the organization information.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationWithContext(ctx context.Context, request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberAuthAccountsRequest() (request *DescribeOrganizationMemberAuthAccountsRequest) {
    request = &DescribeOrganizationMemberAuthAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberAuthAccounts")
    
    
    return
}

func NewDescribeOrganizationMemberAuthAccountsResponse() (response *DescribeOrganizationMemberAuthAccountsResponse) {
    response = &DescribeOrganizationMemberAuthAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberAuthAccounts
// This API is used to get the list of sub-accounts bound to an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthAccounts(request *DescribeOrganizationMemberAuthAccountsRequest) (response *DescribeOrganizationMemberAuthAccountsResponse, err error) {
    return c.DescribeOrganizationMemberAuthAccountsWithContext(context.Background(), request)
}

// DescribeOrganizationMemberAuthAccounts
// This API is used to get the list of sub-accounts bound to an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthAccountsWithContext(ctx context.Context, request *DescribeOrganizationMemberAuthAccountsRequest) (response *DescribeOrganizationMemberAuthAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberAuthAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberAuthAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberAuthAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberAuthIdentitiesRequest() (request *DescribeOrganizationMemberAuthIdentitiesRequest) {
    request = &DescribeOrganizationMemberAuthIdentitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberAuthIdentities")
    
    
    return
}

func NewDescribeOrganizationMemberAuthIdentitiesResponse() (response *DescribeOrganizationMemberAuthIdentitiesResponse) {
    response = &DescribeOrganizationMemberAuthIdentitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberAuthIdentities
// This API is used to get the list of manageable identities of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthIdentities(request *DescribeOrganizationMemberAuthIdentitiesRequest) (response *DescribeOrganizationMemberAuthIdentitiesResponse, err error) {
    return c.DescribeOrganizationMemberAuthIdentitiesWithContext(context.Background(), request)
}

// DescribeOrganizationMemberAuthIdentities
// This API is used to get the list of manageable identities of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthIdentitiesWithContext(ctx context.Context, request *DescribeOrganizationMemberAuthIdentitiesRequest) (response *DescribeOrganizationMemberAuthIdentitiesResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberAuthIdentitiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberAuthIdentities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberAuthIdentitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberPoliciesRequest() (request *DescribeOrganizationMemberPoliciesRequest) {
    request = &DescribeOrganizationMemberPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberPolicies")
    
    
    return
}

func NewDescribeOrganizationMemberPoliciesResponse() (response *DescribeOrganizationMemberPoliciesResponse) {
    response = &DescribeOrganizationMemberPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberPolicies
// This API is used to get the list of authorization policies of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberPolicies(request *DescribeOrganizationMemberPoliciesRequest) (response *DescribeOrganizationMemberPoliciesResponse, err error) {
    return c.DescribeOrganizationMemberPoliciesWithContext(context.Background(), request)
}

// DescribeOrganizationMemberPolicies
// This API is used to get the list of authorization policies of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberPoliciesWithContext(ctx context.Context, request *DescribeOrganizationMemberPoliciesRequest) (response *DescribeOrganizationMemberPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMembersRequest() (request *DescribeOrganizationMembersRequest) {
    request = &DescribeOrganizationMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembers")
    
    
    return
}

func NewDescribeOrganizationMembersResponse() (response *DescribeOrganizationMembersResponse) {
    response = &DescribeOrganizationMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMembers
// This API is used to get the list of organization members.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DescribeOrganizationMembers(request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
    return c.DescribeOrganizationMembersWithContext(context.Background(), request)
}

// DescribeOrganizationMembers
// This API is used to get the list of organization members.
//
// error code that may be returned:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DescribeOrganizationMembersWithContext(ctx context.Context, request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationNodesRequest() (request *DescribeOrganizationNodesRequest) {
    request = &DescribeOrganizationNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodes")
    
    
    return
}

func NewDescribeOrganizationNodesResponse() (response *DescribeOrganizationNodesResponse) {
    response = &DescribeOrganizationNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationNodes
// This API is used to get the list of organization nodes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationNodes(request *DescribeOrganizationNodesRequest) (response *DescribeOrganizationNodesResponse, err error) {
    return c.DescribeOrganizationNodesWithContext(context.Background(), request)
}

// DescribeOrganizationNodes
// This API is used to get the list of organization nodes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationNodesWithContext(ctx context.Context, request *DescribeOrganizationNodesRequest) (response *DescribeOrganizationNodesResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationNodesResponse()
    err = c.Send(request, response)
    return
}

func NewListOrgServiceAssignMemberRequest() (request *ListOrgServiceAssignMemberRequest) {
    request = &ListOrgServiceAssignMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrgServiceAssignMember")
    
    
    return
}

func NewListOrgServiceAssignMemberResponse() (response *ListOrgServiceAssignMemberResponse) {
    response = &ListOrgServiceAssignMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrgServiceAssignMember
// This API is used to obtain the list of delegated admins of the organization service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrgServiceAssignMember(request *ListOrgServiceAssignMemberRequest) (response *ListOrgServiceAssignMemberResponse, err error) {
    return c.ListOrgServiceAssignMemberWithContext(context.Background(), request)
}

// ListOrgServiceAssignMember
// This API is used to obtain the list of delegated admins of the organization service.
//
// error code that may be returned:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrgServiceAssignMemberWithContext(ctx context.Context, request *ListOrgServiceAssignMemberRequest) (response *ListOrgServiceAssignMemberResponse, err error) {
    if request == nil {
        request = NewListOrgServiceAssignMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrgServiceAssignMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrgServiceAssignMemberResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationIdentityRequest() (request *ListOrganizationIdentityRequest) {
    request = &ListOrganizationIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationIdentity")
    
    
    return
}

func NewListOrganizationIdentityResponse() (response *ListOrganizationIdentityResponse) {
    response = &ListOrganizationIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationIdentity
// This API is used to get the list of access identities of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationIdentity(request *ListOrganizationIdentityRequest) (response *ListOrganizationIdentityResponse, err error) {
    return c.ListOrganizationIdentityWithContext(context.Background(), request)
}

// ListOrganizationIdentity
// This API is used to get the list of access identities of an organization member.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationIdentityWithContext(ctx context.Context, request *ListOrganizationIdentityRequest) (response *ListOrganizationIdentityResponse, err error) {
    if request == nil {
        request = NewListOrganizationIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationServiceRequest() (request *ListOrganizationServiceRequest) {
    request = &ListOrganizationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationService")
    
    
    return
}

func NewListOrganizationServiceResponse() (response *ListOrganizationServiceResponse) {
    response = &ListOrganizationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationService
// This API is used to obtain the list of organization service settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationService(request *ListOrganizationServiceRequest) (response *ListOrganizationServiceResponse, err error) {
    return c.ListOrganizationServiceWithContext(context.Background(), request)
}

// ListOrganizationService
// This API is used to obtain the list of organization service settings.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationServiceWithContext(ctx context.Context, request *ListOrganizationServiceRequest) (response *ListOrganizationServiceResponse, err error) {
    if request == nil {
        request = NewListOrganizationServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationServiceResponse()
    err = c.Send(request, response)
    return
}

func NewMoveOrganizationNodeMembersRequest() (request *MoveOrganizationNodeMembersRequest) {
    request = &MoveOrganizationNodeMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "MoveOrganizationNodeMembers")
    
    
    return
}

func NewMoveOrganizationNodeMembersResponse() (response *MoveOrganizationNodeMembersResponse) {
    response = &MoveOrganizationNodeMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MoveOrganizationNodeMembers
// This API is used to move a member to the specified organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) MoveOrganizationNodeMembers(request *MoveOrganizationNodeMembersRequest) (response *MoveOrganizationNodeMembersResponse, err error) {
    return c.MoveOrganizationNodeMembersWithContext(context.Background(), request)
}

// MoveOrganizationNodeMembers
// This API is used to move a member to the specified organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) MoveOrganizationNodeMembersWithContext(ctx context.Context, request *MoveOrganizationNodeMembersRequest) (response *MoveOrganizationNodeMembersResponse, err error) {
    if request == nil {
        request = NewMoveOrganizationNodeMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MoveOrganizationNodeMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewMoveOrganizationNodeMembersResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationNodeRequest() (request *UpdateOrganizationNodeRequest) {
    request = &UpdateOrganizationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationNode")
    
    
    return
}

func NewUpdateOrganizationNodeResponse() (response *UpdateOrganizationNodeResponse) {
    response = &UpdateOrganizationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationNode
// This API is used to update an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) UpdateOrganizationNode(request *UpdateOrganizationNodeRequest) (response *UpdateOrganizationNodeResponse, err error) {
    return c.UpdateOrganizationNodeWithContext(context.Background(), request)
}

// UpdateOrganizationNode
// This API is used to update an organization node.
//
// error code that may be returned:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) UpdateOrganizationNodeWithContext(ctx context.Context, request *UpdateOrganizationNodeRequest) (response *UpdateOrganizationNodeResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationNodeResponse()
    err = c.Send(request, response)
    return
}
