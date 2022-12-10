//
// Copyright (C) 2021 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

type TransmissionClient struct {
	baseUrl      string
	authInjector interfaces.AuthenticationInjector
}

// NewTransmissionClient creates an instance of TransmissionClient
func NewTransmissionClient(baseUrl string, authInjector interfaces.AuthenticationInjector) interfaces.TransmissionClient {
	return &TransmissionClient{
		baseUrl:      baseUrl,
		authInjector: authInjector,
	}
}

// TransmissionById query transmission by id.
func (client *TransmissionClient) TransmissionById(ctx context.Context, id string) (res responses.TransmissionResponse, err errors.EdgeX) {
	path := path.Join(common.ApiTransmissionRoute, common.Id, id)
	err = utils.GetRequest(ctx, &res, client.baseUrl, path, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// TransmissionsByTimeRange query transmissions with time range, offset and limit
func (client *TransmissionClient) TransmissionsByTimeRange(ctx context.Context, start int, end int, offset int, limit int) (res responses.MultiTransmissionsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiTransmissionRoute, common.Start, strconv.Itoa(start), common.End, strconv.Itoa(end))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllTransmissions query transmissions with offset and limit
func (client *TransmissionClient) AllTransmissions(ctx context.Context, offset int, limit int) (res responses.MultiTransmissionsResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, common.ApiAllTransmissionRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// TransmissionsByStatus queries transmissions with status, offset and limit
func (client *TransmissionClient) TransmissionsByStatus(ctx context.Context, status string, offset int, limit int) (res responses.MultiTransmissionsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiTransmissionRoute, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteProcessedTransmissionsByAge deletes the processed transmissions if the current timestamp minus their created timestamp is less than the age parameter.
func (client *TransmissionClient) DeleteProcessedTransmissionsByAge(ctx context.Context, age int) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := path.Join(common.ApiTransmissionRoute, common.Age, strconv.Itoa(age))
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, path, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// TransmissionsBySubscriptionName query transmissions with subscriptionName, offset and limit
func (client *TransmissionClient) TransmissionsBySubscriptionName(ctx context.Context, subscriptionName string, offset int, limit int) (res responses.MultiTransmissionsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiTransmissionRoute, common.Subscription, common.Name, subscriptionName)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// TransmissionsByNotificationId query transmissions with notification id, offset and limit
func (client *TransmissionClient) TransmissionsByNotificationId(ctx context.Context, id string, offset int, limit int) (res responses.MultiTransmissionsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiTransmissionRoute, common.Notification, common.Id, id)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}
