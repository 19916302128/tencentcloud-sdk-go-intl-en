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

package v20211228

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2021-12-28"

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


func NewCreateInstanceNewRequest() (request *CreateInstanceNewRequest) {
    request = &CreateInstanceNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "CreateInstanceNew")
    
    
    return
}

func NewCreateInstanceNewResponse() (response *CreateInstanceNewResponse) {
    response = &CreateInstanceNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstanceNew
// This API is used to create clusters.
func (c *Client) CreateInstanceNew(request *CreateInstanceNewRequest) (response *CreateInstanceNewResponse, err error) {
    return c.CreateInstanceNewWithContext(context.Background(), request)
}

// CreateInstanceNew
// This API is used to create clusters.
func (c *Client) CreateInstanceNewWithContext(ctx context.Context, request *CreateInstanceNewRequest) (response *CreateInstanceNewResponse, err error) {
    if request == nil {
        request = NewCreateInstanceNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterConfigsRequest() (request *DescribeClusterConfigsRequest) {
    request = &DescribeClusterConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeClusterConfigs")
    
    
    return
}

func NewDescribeClusterConfigsResponse() (response *DescribeClusterConfigsResponse) {
    response = &DescribeClusterConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterConfigs
// This API is used to get the contents of the latest configuration files (config.xml, metrika.xml, and user.xml) of the cluster and display them to the user.
func (c *Client) DescribeClusterConfigs(request *DescribeClusterConfigsRequest) (response *DescribeClusterConfigsResponse, err error) {
    return c.DescribeClusterConfigsWithContext(context.Background(), request)
}

// DescribeClusterConfigs
// This API is used to get the contents of the latest configuration files (config.xml, metrika.xml, and user.xml) of the cluster and display them to the user.
func (c *Client) DescribeClusterConfigsWithContext(ctx context.Context, request *DescribeClusterConfigsRequest) (response *DescribeClusterConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseAuditDownloadRequest() (request *DescribeDatabaseAuditDownloadRequest) {
    request = &DescribeDatabaseAuditDownloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeDatabaseAuditDownload")
    
    
    return
}

func NewDescribeDatabaseAuditDownloadResponse() (response *DescribeDatabaseAuditDownloadResponse) {
    response = &DescribeDatabaseAuditDownloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseAuditDownload
// This API is used to download database audit logs.
func (c *Client) DescribeDatabaseAuditDownload(request *DescribeDatabaseAuditDownloadRequest) (response *DescribeDatabaseAuditDownloadResponse, err error) {
    return c.DescribeDatabaseAuditDownloadWithContext(context.Background(), request)
}

// DescribeDatabaseAuditDownload
// This API is used to download database audit logs.
func (c *Client) DescribeDatabaseAuditDownloadWithContext(ctx context.Context, request *DescribeDatabaseAuditDownloadRequest) (response *DescribeDatabaseAuditDownloadResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseAuditDownloadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseAuditDownload require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseAuditDownloadResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseAuditRecordsRequest() (request *DescribeDatabaseAuditRecordsRequest) {
    request = &DescribeDatabaseAuditRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeDatabaseAuditRecords")
    
    
    return
}

func NewDescribeDatabaseAuditRecordsResponse() (response *DescribeDatabaseAuditRecordsResponse) {
    response = &DescribeDatabaseAuditRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseAuditRecords
// This API is used to get database audit records.
func (c *Client) DescribeDatabaseAuditRecords(request *DescribeDatabaseAuditRecordsRequest) (response *DescribeDatabaseAuditRecordsResponse, err error) {
    return c.DescribeDatabaseAuditRecordsWithContext(context.Background(), request)
}

// DescribeDatabaseAuditRecords
// This API is used to get database audit records.
func (c *Client) DescribeDatabaseAuditRecordsWithContext(ctx context.Context, request *DescribeDatabaseAuditRecordsRequest) (response *DescribeDatabaseAuditRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseAuditRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseAuditRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseAuditRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstance")
    
    
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstance
// This API is used to query the specific information of a cluster based on the cluster ID.
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

// DescribeInstance
// This API is used to query the specific information of a cluster based on the cluster ID.
func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodesRequest() (request *DescribeInstanceNodesRequest) {
    request = &DescribeInstanceNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstanceNodes")
    
    
    return
}

func NewDescribeInstanceNodesResponse() (response *DescribeInstanceNodesResponse) {
    response = &DescribeInstanceNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNodes
// This API is used to get the list of cluster node information.
func (c *Client) DescribeInstanceNodes(request *DescribeInstanceNodesRequest) (response *DescribeInstanceNodesResponse, err error) {
    return c.DescribeInstanceNodesWithContext(context.Background(), request)
}

// DescribeInstanceNodes
// This API is used to get the list of cluster node information.
func (c *Client) DescribeInstanceNodesWithContext(ctx context.Context, request *DescribeInstanceNodesRequest) (response *DescribeInstanceNodesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodesInfoRequest() (request *DescribeInstanceNodesInfoRequest) {
    request = &DescribeInstanceNodesInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstanceNodesInfo")
    
    
    return
}

func NewDescribeInstanceNodesInfoResponse() (response *DescribeInstanceNodesInfoResponse) {
    response = &DescribeInstanceNodesInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNodesInfo
// This API is used to get the BE/FE node roles.
func (c *Client) DescribeInstanceNodesInfo(request *DescribeInstanceNodesInfoRequest) (response *DescribeInstanceNodesInfoResponse, err error) {
    return c.DescribeInstanceNodesInfoWithContext(context.Background(), request)
}

// DescribeInstanceNodesInfo
// This API is used to get the BE/FE node roles.
func (c *Client) DescribeInstanceNodesInfoWithContext(ctx context.Context, request *DescribeInstanceNodesInfoRequest) (response *DescribeInstanceNodesInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodesInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodesInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodesInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceStateRequest() (request *DescribeInstanceStateRequest) {
    request = &DescribeInstanceStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstanceState")
    
    
    return
}

func NewDescribeInstanceStateResponse() (response *DescribeInstanceStateResponse) {
    response = &DescribeInstanceStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceState
// This API is used to display cluster status, process progress, etc. in the cluster details page.
func (c *Client) DescribeInstanceState(request *DescribeInstanceStateRequest) (response *DescribeInstanceStateResponse, err error) {
    return c.DescribeInstanceStateWithContext(context.Background(), request)
}

// DescribeInstanceState
// This API is used to display cluster status, process progress, etc. in the cluster details page.
func (c *Client) DescribeInstanceStateWithContext(ctx context.Context, request *DescribeInstanceStateRequest) (response *DescribeInstanceStateResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// This API is used to get the list of clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// This API is used to get the list of clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowQueryRecordsRequest() (request *DescribeSlowQueryRecordsRequest) {
    request = &DescribeSlowQueryRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeSlowQueryRecords")
    
    
    return
}

func NewDescribeSlowQueryRecordsResponse() (response *DescribeSlowQueryRecordsResponse) {
    response = &DescribeSlowQueryRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowQueryRecords
// This API is used to get the slow log list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSlowQueryRecords(request *DescribeSlowQueryRecordsRequest) (response *DescribeSlowQueryRecordsResponse, err error) {
    return c.DescribeSlowQueryRecordsWithContext(context.Background(), request)
}

// DescribeSlowQueryRecords
// This API is used to get the slow log list.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSlowQueryRecordsWithContext(ctx context.Context, request *DescribeSlowQueryRecordsRequest) (response *DescribeSlowQueryRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowQueryRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowQueryRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowQueryRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowQueryRecordsDownloadRequest() (request *DescribeSlowQueryRecordsDownloadRequest) {
    request = &DescribeSlowQueryRecordsDownloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeSlowQueryRecordsDownload")
    
    
    return
}

func NewDescribeSlowQueryRecordsDownloadResponse() (response *DescribeSlowQueryRecordsDownloadResponse) {
    response = &DescribeSlowQueryRecordsDownloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowQueryRecordsDownload
// This API is used to download slow log files.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSlowQueryRecordsDownload(request *DescribeSlowQueryRecordsDownloadRequest) (response *DescribeSlowQueryRecordsDownloadResponse, err error) {
    return c.DescribeSlowQueryRecordsDownloadWithContext(context.Background(), request)
}

// DescribeSlowQueryRecordsDownload
// This API is used to download slow log files.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSlowQueryRecordsDownloadWithContext(ctx context.Context, request *DescribeSlowQueryRecordsDownloadRequest) (response *DescribeSlowQueryRecordsDownloadResponse, err error) {
    if request == nil {
        request = NewDescribeSlowQueryRecordsDownloadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowQueryRecordsDownload require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowQueryRecordsDownloadResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyInstanceRequest() (request *DestroyInstanceRequest) {
    request = &DestroyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DestroyInstance")
    
    
    return
}

func NewDestroyInstanceResponse() (response *DestroyInstanceResponse) {
    response = &DestroyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyInstance
// This API is used to terminate clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DestroyInstance(request *DestroyInstanceRequest) (response *DestroyInstanceResponse, err error) {
    return c.DestroyInstanceWithContext(context.Background(), request)
}

// DestroyInstance
// This API is used to terminate clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) DestroyInstanceWithContext(ctx context.Context, request *DestroyInstanceRequest) (response *DestroyInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstance
// This API is used to modify the cluster's name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// This API is used to modify the cluster's name.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
    request = &ResizeDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ResizeDisk")
    
    
    return
}

func NewResizeDiskResponse() (response *ResizeDiskResponse) {
    response = &ResizeDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResizeDisk
// This API is used to expand cloud disks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
    return c.ResizeDiskWithContext(context.Background(), request)
}

// ResizeDisk
// This API is used to expand cloud disks.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ResizeDiskWithContext(ctx context.Context, request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
    if request == nil {
        request = NewResizeDiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeDiskResponse()
    err = c.Send(request, response)
    return
}

func NewRestartClusterForNodeRequest() (request *RestartClusterForNodeRequest) {
    request = &RestartClusterForNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "RestartClusterForNode")
    
    
    return
}

func NewRestartClusterForNodeResponse() (response *RestartClusterForNodeResponse) {
    response = &RestartClusterForNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartClusterForNode
// This API is used to indicate the rolling restart of the clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) RestartClusterForNode(request *RestartClusterForNodeRequest) (response *RestartClusterForNodeResponse, err error) {
    return c.RestartClusterForNodeWithContext(context.Background(), request)
}

// RestartClusterForNode
// This API is used to indicate the rolling restart of the clusters.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) RestartClusterForNodeWithContext(ctx context.Context, request *RestartClusterForNodeRequest) (response *RestartClusterForNodeResponse, err error) {
    if request == nil {
        request = NewRestartClusterForNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartClusterForNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartClusterForNodeResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutInstanceRequest() (request *ScaleOutInstanceRequest) {
    request = &ScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ScaleOutInstance")
    
    
    return
}

func NewScaleOutInstanceResponse() (response *ScaleOutInstanceResponse) {
    response = &ScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleOutInstance
// This API is used to horizontally scale out nodes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleOutInstance(request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    return c.ScaleOutInstanceWithContext(context.Background(), request)
}

// ScaleOutInstance
// This API is used to horizontally scale out nodes.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleOutInstanceWithContext(ctx context.Context, request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    if request == nil {
        request = NewScaleOutInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleOutInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleOutInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewScaleUpInstanceRequest() (request *ScaleUpInstanceRequest) {
    request = &ScaleUpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ScaleUpInstance")
    
    
    return
}

func NewScaleUpInstanceResponse() (response *ScaleUpInstanceResponse) {
    response = &ScaleUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleUpInstance
// This API is used to scale up/down computing resources.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleUpInstance(request *ScaleUpInstanceRequest) (response *ScaleUpInstanceResponse, err error) {
    return c.ScaleUpInstanceWithContext(context.Background(), request)
}

// ScaleUpInstance
// This API is used to scale up/down computing resources.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleUpInstanceWithContext(ctx context.Context, request *ScaleUpInstanceRequest) (response *ScaleUpInstanceResponse, err error) {
    if request == nil {
        request = NewScaleUpInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleUpInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleUpInstanceResponse()
    err = c.Send(request, response)
    return
}
