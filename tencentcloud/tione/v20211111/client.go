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

package v20211111

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-11-11"

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


func NewCreateDatasetRequest() (request *CreateDatasetRequest) {
    request = &CreateDatasetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "CreateDataset")
    
    
    return
}

func NewCreateDatasetResponse() (response *CreateDatasetResponse) {
    response = &CreateDatasetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDataset
// 创建数据集
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateDataset(request *CreateDatasetRequest) (response *CreateDatasetResponse, err error) {
    return c.CreateDatasetWithContext(context.Background(), request)
}

// CreateDataset
// 创建数据集
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateDatasetWithContext(ctx context.Context, request *CreateDatasetRequest) (response *CreateDatasetResponse, err error) {
    if request == nil {
        request = NewCreateDatasetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatasetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTrainingModelRequest() (request *CreateTrainingModelRequest) {
    request = &CreateTrainingModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "CreateTrainingModel")
    
    
    return
}

func NewCreateTrainingModelResponse() (response *CreateTrainingModelResponse) {
    response = &CreateTrainingModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTrainingModel
// 导入模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_MOVEMODELDIRFAILED = "FailedOperation.MoveModelDirFailed"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingModel(request *CreateTrainingModelRequest) (response *CreateTrainingModelResponse, err error) {
    return c.CreateTrainingModelWithContext(context.Background(), request)
}

// CreateTrainingModel
// 导入模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_MOVEMODELDIRFAILED = "FailedOperation.MoveModelDirFailed"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingModelWithContext(ctx context.Context, request *CreateTrainingModelRequest) (response *CreateTrainingModelResponse, err error) {
    if request == nil {
        request = NewCreateTrainingModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTrainingModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTrainingModelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTrainingTaskRequest() (request *CreateTrainingTaskRequest) {
    request = &CreateTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "CreateTrainingTask")
    
    
    return
}

func NewCreateTrainingTaskResponse() (response *CreateTrainingTaskResponse) {
    response = &CreateTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTrainingTask
// 创建模型训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATASETNUMLIMITEXCEEDED = "InvalidParameterValue.DatasetNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingTask(request *CreateTrainingTaskRequest) (response *CreateTrainingTaskResponse, err error) {
    return c.CreateTrainingTaskWithContext(context.Background(), request)
}

// CreateTrainingTask
// 创建模型训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATASETNUMLIMITEXCEEDED = "InvalidParameterValue.DatasetNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingTaskWithContext(ctx context.Context, request *CreateTrainingTaskRequest) (response *CreateTrainingTaskResponse, err error) {
    if request == nil {
        request = NewCreateTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDatasetRequest() (request *DeleteDatasetRequest) {
    request = &DeleteDatasetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DeleteDataset")
    
    
    return
}

func NewDeleteDatasetResponse() (response *DeleteDatasetResponse) {
    response = &DeleteDatasetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDataset
// 删除数据集
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteDataset(request *DeleteDatasetRequest) (response *DeleteDatasetResponse, err error) {
    return c.DeleteDatasetWithContext(context.Background(), request)
}

// DeleteDataset
// 删除数据集
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteDatasetWithContext(ctx context.Context, request *DeleteDatasetRequest) (response *DeleteDatasetResponse, err error) {
    if request == nil {
        request = NewDeleteDatasetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDatasetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTrainingModelRequest() (request *DeleteTrainingModelRequest) {
    request = &DeleteTrainingModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DeleteTrainingModel")
    
    
    return
}

func NewDeleteTrainingModelResponse() (response *DeleteTrainingModelResponse) {
    response = &DeleteTrainingModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTrainingModel
// 删除模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_NOMODEL = "ResourceNotFound.NoModel"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingModel(request *DeleteTrainingModelRequest) (response *DeleteTrainingModelResponse, err error) {
    return c.DeleteTrainingModelWithContext(context.Background(), request)
}

// DeleteTrainingModel
// 删除模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_NOMODEL = "ResourceNotFound.NoModel"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingModelWithContext(ctx context.Context, request *DeleteTrainingModelRequest) (response *DeleteTrainingModelResponse, err error) {
    if request == nil {
        request = NewDeleteTrainingModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTrainingModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTrainingModelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTrainingModelVersionRequest() (request *DeleteTrainingModelVersionRequest) {
    request = &DeleteTrainingModelVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DeleteTrainingModelVersion")
    
    
    return
}

func NewDeleteTrainingModelVersionResponse() (response *DeleteTrainingModelVersionResponse) {
    response = &DeleteTrainingModelVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTrainingModelVersion
// 删除模型版本
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingModelVersion(request *DeleteTrainingModelVersionRequest) (response *DeleteTrainingModelVersionResponse, err error) {
    return c.DeleteTrainingModelVersionWithContext(context.Background(), request)
}

// DeleteTrainingModelVersion
// 删除模型版本
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingModelVersionWithContext(ctx context.Context, request *DeleteTrainingModelVersionRequest) (response *DeleteTrainingModelVersionResponse, err error) {
    if request == nil {
        request = NewDeleteTrainingModelVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTrainingModelVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTrainingModelVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTrainingTaskRequest() (request *DeleteTrainingTaskRequest) {
    request = &DeleteTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DeleteTrainingTask")
    
    
    return
}

func NewDeleteTrainingTaskResponse() (response *DeleteTrainingTaskResponse) {
    response = &DeleteTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTrainingTask
// 删除训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingTask(request *DeleteTrainingTaskRequest) (response *DeleteTrainingTaskResponse, err error) {
    return c.DeleteTrainingTaskWithContext(context.Background(), request)
}

// DeleteTrainingTask
// 删除训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingTaskWithContext(ctx context.Context, request *DeleteTrainingTaskRequest) (response *DeleteTrainingTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingResourceGroupsRequest() (request *DescribeBillingResourceGroupsRequest) {
    request = &DescribeBillingResourceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBillingResourceGroups")
    
    
    return
}

func NewDescribeBillingResourceGroupsResponse() (response *DescribeBillingResourceGroupsResponse) {
    response = &DescribeBillingResourceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillingResourceGroups
// 查询资源组详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VALIDATEERROR = "InvalidParameter.ValidateError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillingResourceGroups(request *DescribeBillingResourceGroupsRequest) (response *DescribeBillingResourceGroupsResponse, err error) {
    return c.DescribeBillingResourceGroupsWithContext(context.Background(), request)
}

// DescribeBillingResourceGroups
// 查询资源组详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VALIDATEERROR = "InvalidParameter.ValidateError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillingResourceGroupsWithContext(ctx context.Context, request *DescribeBillingResourceGroupsRequest) (response *DescribeBillingResourceGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeBillingResourceGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingResourceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingResourceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingSpecsPriceRequest() (request *DescribeBillingSpecsPriceRequest) {
    request = &DescribeBillingSpecsPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBillingSpecsPrice")
    
    
    return
}

func NewDescribeBillingSpecsPriceResponse() (response *DescribeBillingSpecsPriceResponse) {
    response = &DescribeBillingSpecsPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillingSpecsPrice
// 本接口(DescribeBillingSpecsPrice)用于查询计费项价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillingSpecsPrice(request *DescribeBillingSpecsPriceRequest) (response *DescribeBillingSpecsPriceResponse, err error) {
    return c.DescribeBillingSpecsPriceWithContext(context.Background(), request)
}

// DescribeBillingSpecsPrice
// 本接口(DescribeBillingSpecsPrice)用于查询计费项价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillingSpecsPriceWithContext(ctx context.Context, request *DescribeBillingSpecsPriceRequest) (response *DescribeBillingSpecsPriceResponse, err error) {
    if request == nil {
        request = NewDescribeBillingSpecsPriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingSpecsPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingSpecsPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatasetDetailStructuredRequest() (request *DescribeDatasetDetailStructuredRequest) {
    request = &DescribeDatasetDetailStructuredRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeDatasetDetailStructured")
    
    
    return
}

func NewDescribeDatasetDetailStructuredResponse() (response *DescribeDatasetDetailStructuredResponse) {
    response = &DescribeDatasetDetailStructuredResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatasetDetailStructured
// 查询结构化数据集详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetDetailStructured(request *DescribeDatasetDetailStructuredRequest) (response *DescribeDatasetDetailStructuredResponse, err error) {
    return c.DescribeDatasetDetailStructuredWithContext(context.Background(), request)
}

// DescribeDatasetDetailStructured
// 查询结构化数据集详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetDetailStructuredWithContext(ctx context.Context, request *DescribeDatasetDetailStructuredRequest) (response *DescribeDatasetDetailStructuredResponse, err error) {
    if request == nil {
        request = NewDescribeDatasetDetailStructuredRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasetDetailStructured require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasetDetailStructuredResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatasetDetailUnstructuredRequest() (request *DescribeDatasetDetailUnstructuredRequest) {
    request = &DescribeDatasetDetailUnstructuredRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeDatasetDetailUnstructured")
    
    
    return
}

func NewDescribeDatasetDetailUnstructuredResponse() (response *DescribeDatasetDetailUnstructuredResponse) {
    response = &DescribeDatasetDetailUnstructuredResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatasetDetailUnstructured
// 查询非结构化数据集详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetDetailUnstructured(request *DescribeDatasetDetailUnstructuredRequest) (response *DescribeDatasetDetailUnstructuredResponse, err error) {
    return c.DescribeDatasetDetailUnstructuredWithContext(context.Background(), request)
}

// DescribeDatasetDetailUnstructured
// 查询非结构化数据集详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetDetailUnstructuredWithContext(ctx context.Context, request *DescribeDatasetDetailUnstructuredRequest) (response *DescribeDatasetDetailUnstructuredResponse, err error) {
    if request == nil {
        request = NewDescribeDatasetDetailUnstructuredRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasetDetailUnstructured require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasetDetailUnstructuredResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatasetsRequest() (request *DescribeDatasetsRequest) {
    request = &DescribeDatasetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeDatasets")
    
    
    return
}

func NewDescribeDatasetsResponse() (response *DescribeDatasetsResponse) {
    response = &DescribeDatasetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatasets
// 查询数据集列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasets(request *DescribeDatasetsRequest) (response *DescribeDatasetsResponse, err error) {
    return c.DescribeDatasetsWithContext(context.Background(), request)
}

// DescribeDatasets
// 查询数据集列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetsWithContext(ctx context.Context, request *DescribeDatasetsRequest) (response *DescribeDatasetsResponse, err error) {
    if request == nil {
        request = NewDescribeDatasetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInferTemplatesRequest() (request *DescribeInferTemplatesRequest) {
    request = &DescribeInferTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeInferTemplates")
    
    
    return
}

func NewDescribeInferTemplatesResponse() (response *DescribeInferTemplatesResponse) {
    response = &DescribeInferTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInferTemplates
// 查询推理镜像模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInferTemplates(request *DescribeInferTemplatesRequest) (response *DescribeInferTemplatesResponse, err error) {
    return c.DescribeInferTemplatesWithContext(context.Background(), request)
}

// DescribeInferTemplates
// 查询推理镜像模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInferTemplatesWithContext(ctx context.Context, request *DescribeInferTemplatesRequest) (response *DescribeInferTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeInferTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLatestTrainingMetricsRequest() (request *DescribeLatestTrainingMetricsRequest) {
    request = &DescribeLatestTrainingMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeLatestTrainingMetrics")
    
    
    return
}

func NewDescribeLatestTrainingMetricsResponse() (response *DescribeLatestTrainingMetricsResponse) {
    response = &DescribeLatestTrainingMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLatestTrainingMetrics
// 查询最近上报的训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_CAMFAILURE = "FailedOperation.CAMFailure"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CREATEJOBINSTANCEFAILED = "FailedOperation.CreateJobInstanceFailed"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FREEZEBILLFAILED = "InternalError.FreezeBillFailed"
//  INTERNALERROR_INSUFFICIENTBALANCE = "InternalError.InsufficientBalance"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLatestTrainingMetrics(request *DescribeLatestTrainingMetricsRequest) (response *DescribeLatestTrainingMetricsResponse, err error) {
    return c.DescribeLatestTrainingMetricsWithContext(context.Background(), request)
}

// DescribeLatestTrainingMetrics
// 查询最近上报的训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_CAMFAILURE = "FailedOperation.CAMFailure"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CREATEJOBINSTANCEFAILED = "FailedOperation.CreateJobInstanceFailed"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FREEZEBILLFAILED = "InternalError.FreezeBillFailed"
//  INTERNALERROR_INSUFFICIENTBALANCE = "InternalError.InsufficientBalance"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLatestTrainingMetricsWithContext(ctx context.Context, request *DescribeLatestTrainingMetricsRequest) (response *DescribeLatestTrainingMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeLatestTrainingMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLatestTrainingMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLatestTrainingMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogsRequest() (request *DescribeLogsRequest) {
    request = &DescribeLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeLogs")
    
    
    return
}

func NewDescribeLogsResponse() (response *DescribeLogsResponse) {
    response = &DescribeLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogs
// 获取训练、推理、Notebook服务的日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeLogs(request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    return c.DescribeLogsWithContext(context.Background(), request)
}

// DescribeLogs
// 获取训练、推理、Notebook服务的日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
func (c *Client) DescribeLogsWithContext(ctx context.Context, request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    if request == nil {
        request = NewDescribeLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingFrameworksRequest() (request *DescribeTrainingFrameworksRequest) {
    request = &DescribeTrainingFrameworksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingFrameworks")
    
    
    return
}

func NewDescribeTrainingFrameworksResponse() (response *DescribeTrainingFrameworksResponse) {
    response = &DescribeTrainingFrameworksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingFrameworks
// 训练框架列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingFrameworks(request *DescribeTrainingFrameworksRequest) (response *DescribeTrainingFrameworksResponse, err error) {
    return c.DescribeTrainingFrameworksWithContext(context.Background(), request)
}

// DescribeTrainingFrameworks
// 训练框架列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingFrameworksWithContext(ctx context.Context, request *DescribeTrainingFrameworksRequest) (response *DescribeTrainingFrameworksResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingFrameworksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingFrameworks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingFrameworksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingMetricsRequest() (request *DescribeTrainingMetricsRequest) {
    request = &DescribeTrainingMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingMetrics")
    
    
    return
}

func NewDescribeTrainingMetricsResponse() (response *DescribeTrainingMetricsResponse) {
    response = &DescribeTrainingMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingMetrics
// 查询训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTrainingMetrics(request *DescribeTrainingMetricsRequest) (response *DescribeTrainingMetricsResponse, err error) {
    return c.DescribeTrainingMetricsWithContext(context.Background(), request)
}

// DescribeTrainingMetrics
// 查询训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTrainingMetricsWithContext(ctx context.Context, request *DescribeTrainingMetricsRequest) (response *DescribeTrainingMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingModelVersionRequest() (request *DescribeTrainingModelVersionRequest) {
    request = &DescribeTrainingModelVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingModelVersion")
    
    
    return
}

func NewDescribeTrainingModelVersionResponse() (response *DescribeTrainingModelVersionResponse) {
    response = &DescribeTrainingModelVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingModelVersion
// 查询模型版本
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelVersion(request *DescribeTrainingModelVersionRequest) (response *DescribeTrainingModelVersionResponse, err error) {
    return c.DescribeTrainingModelVersionWithContext(context.Background(), request)
}

// DescribeTrainingModelVersion
// 查询模型版本
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelVersionWithContext(ctx context.Context, request *DescribeTrainingModelVersionRequest) (response *DescribeTrainingModelVersionResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingModelVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingModelVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingModelVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingModelVersionsRequest() (request *DescribeTrainingModelVersionsRequest) {
    request = &DescribeTrainingModelVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingModelVersions")
    
    
    return
}

func NewDescribeTrainingModelVersionsResponse() (response *DescribeTrainingModelVersionsResponse) {
    response = &DescribeTrainingModelVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingModelVersions
// 模型版本列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOMODEL = "ResourceNotFound.NoModel"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelVersions(request *DescribeTrainingModelVersionsRequest) (response *DescribeTrainingModelVersionsResponse, err error) {
    return c.DescribeTrainingModelVersionsWithContext(context.Background(), request)
}

// DescribeTrainingModelVersions
// 模型版本列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOMODEL = "ResourceNotFound.NoModel"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelVersionsWithContext(ctx context.Context, request *DescribeTrainingModelVersionsRequest) (response *DescribeTrainingModelVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingModelVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingModelVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingModelVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingModelsRequest() (request *DescribeTrainingModelsRequest) {
    request = &DescribeTrainingModelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingModels")
    
    
    return
}

func NewDescribeTrainingModelsResponse() (response *DescribeTrainingModelsResponse) {
    response = &DescribeTrainingModelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingModels
// 模型列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYBINDINGTAGSFAILED = "FailedOperation.QueryBindingTagsFailed"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYMODELSBYTAGSFAILED = "FailedOperation.QueryModelsByTagsFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModels(request *DescribeTrainingModelsRequest) (response *DescribeTrainingModelsResponse, err error) {
    return c.DescribeTrainingModelsWithContext(context.Background(), request)
}

// DescribeTrainingModels
// 模型列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYBINDINGTAGSFAILED = "FailedOperation.QueryBindingTagsFailed"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYMODELSBYTAGSFAILED = "FailedOperation.QueryModelsByTagsFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelsWithContext(ctx context.Context, request *DescribeTrainingModelsRequest) (response *DescribeTrainingModelsResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingModelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingModels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingModelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingTaskRequest() (request *DescribeTrainingTaskRequest) {
    request = &DescribeTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingTask")
    
    
    return
}

func NewDescribeTrainingTaskResponse() (response *DescribeTrainingTaskResponse) {
    response = &DescribeTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingTask
// 训练任务详情
//
// 可能返回的错误码:
//  INTERNALERROR_NOPERMISSION = "InternalError.NoPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTrainingTask(request *DescribeTrainingTaskRequest) (response *DescribeTrainingTaskResponse, err error) {
    return c.DescribeTrainingTaskWithContext(context.Background(), request)
}

// DescribeTrainingTask
// 训练任务详情
//
// 可能返回的错误码:
//  INTERNALERROR_NOPERMISSION = "InternalError.NoPermission"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTrainingTaskWithContext(ctx context.Context, request *DescribeTrainingTaskRequest) (response *DescribeTrainingTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingTaskPodsRequest() (request *DescribeTrainingTaskPodsRequest) {
    request = &DescribeTrainingTaskPodsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingTaskPods")
    
    
    return
}

func NewDescribeTrainingTaskPodsResponse() (response *DescribeTrainingTaskPodsResponse) {
    response = &DescribeTrainingTaskPodsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingTaskPods
// 训练任务pod列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingTaskPods(request *DescribeTrainingTaskPodsRequest) (response *DescribeTrainingTaskPodsResponse, err error) {
    return c.DescribeTrainingTaskPodsWithContext(context.Background(), request)
}

// DescribeTrainingTaskPods
// 训练任务pod列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingTaskPodsWithContext(ctx context.Context, request *DescribeTrainingTaskPodsRequest) (response *DescribeTrainingTaskPodsResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingTaskPodsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingTaskPods require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingTaskPodsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingTasksRequest() (request *DescribeTrainingTasksRequest) {
    request = &DescribeTrainingTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingTasks")
    
    
    return
}

func NewDescribeTrainingTasksResponse() (response *DescribeTrainingTasksResponse) {
    response = &DescribeTrainingTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingTasks
// 训练任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTrainingTasks(request *DescribeTrainingTasksRequest) (response *DescribeTrainingTasksResponse, err error) {
    return c.DescribeTrainingTasksWithContext(context.Background(), request)
}

// DescribeTrainingTasks
// 训练任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTrainingTasksWithContext(ctx context.Context, request *DescribeTrainingTasksRequest) (response *DescribeTrainingTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingTasksResponse()
    err = c.Send(request, response)
    return
}

func NewPushTrainingMetricsRequest() (request *PushTrainingMetricsRequest) {
    request = &PushTrainingMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "PushTrainingMetrics")
    
    
    return
}

func NewPushTrainingMetricsResponse() (response *PushTrainingMetricsResponse) {
    response = &PushTrainingMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PushTrainingMetrics
// 上报训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) PushTrainingMetrics(request *PushTrainingMetricsRequest) (response *PushTrainingMetricsResponse, err error) {
    return c.PushTrainingMetricsWithContext(context.Background(), request)
}

// PushTrainingMetrics
// 上报训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) PushTrainingMetricsWithContext(ctx context.Context, request *PushTrainingMetricsRequest) (response *PushTrainingMetricsResponse, err error) {
    if request == nil {
        request = NewPushTrainingMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PushTrainingMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewPushTrainingMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewStartTrainingTaskRequest() (request *StartTrainingTaskRequest) {
    request = &StartTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "StartTrainingTask")
    
    
    return
}

func NewStartTrainingTaskResponse() (response *StartTrainingTaskResponse) {
    response = &StartTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartTrainingTask
// 启动模型训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StartTrainingTask(request *StartTrainingTaskRequest) (response *StartTrainingTaskResponse, err error) {
    return c.StartTrainingTaskWithContext(context.Background(), request)
}

// StartTrainingTask
// 启动模型训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StartTrainingTaskWithContext(ctx context.Context, request *StartTrainingTaskRequest) (response *StartTrainingTaskResponse, err error) {
    if request == nil {
        request = NewStartTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopTrainingTaskRequest() (request *StopTrainingTaskRequest) {
    request = &StopTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tione", APIVersion, "StopTrainingTask")
    
    
    return
}

func NewStopTrainingTaskResponse() (response *StopTrainingTaskResponse) {
    response = &StopTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopTrainingTask
// 停止模型训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StopTrainingTask(request *StopTrainingTaskRequest) (response *StopTrainingTaskResponse, err error) {
    return c.StopTrainingTaskWithContext(context.Background(), request)
}

// StopTrainingTask
// 停止模型训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StopTrainingTaskWithContext(ctx context.Context, request *StopTrainingTaskRequest) (response *StopTrainingTaskResponse, err error) {
    if request == nil {
        request = NewStopTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopTrainingTaskResponse()
    err = c.Send(request, response)
    return
}