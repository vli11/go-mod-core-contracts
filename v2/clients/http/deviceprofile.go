//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

type DeviceProfileClient struct {
	baseUrl        string
	resourcesCache map[string]responses.DeviceResourceResponse
	mux            sync.RWMutex
}

// NewDeviceProfileClient creates an instance of DeviceProfileClient
func NewDeviceProfileClient(baseUrl string) interfaces.DeviceProfileClient {
	return &DeviceProfileClient{
		baseUrl:        baseUrl,
		resourcesCache: make(map[string]responses.DeviceResourceResponse),
	}
}

// Add adds new device profile
func (client *DeviceProfileClient) Add(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	var responses []common.BaseWithIdResponse
	err := utils.PostRequestWithRawData(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileRoute, reqs)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// Update updates device profile
func (client *DeviceProfileClient) Update(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseResponse, errors.EdgeX) {
	var responses []common.BaseResponse
	err := utils.PutRequest(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileRoute, reqs)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// AddByYaml adds new device profile by uploading a yaml file
func (client *DeviceProfileClient) AddByYaml(ctx context.Context, yamlFilePath string) (common.BaseWithIdResponse, errors.EdgeX) {
	var responses common.BaseWithIdResponse
	err := utils.PostByFileRequest(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileUploadFileRoute, yamlFilePath)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// UpdateByYaml updates device profile by uploading a yaml file
func (client *DeviceProfileClient) UpdateByYaml(ctx context.Context, yamlFilePath string) (common.BaseResponse, errors.EdgeX) {
	var responses common.BaseResponse
	err := utils.PutByFileRequest(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileUploadFileRoute, yamlFilePath)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// DeleteByName deletes the device profile by name
func (client *DeviceProfileClient) DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	requestPath := path.Join(v2.ApiDeviceProfileRoute, v2.Name, url.QueryEscape(name))
	err := utils.DeleteRequest(ctx, &response, client.baseUrl, requestPath)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

// DeviceProfileByName queries the device profile by name
func (client *DeviceProfileClient) DeviceProfileByName(ctx context.Context, name string) (res responses.DeviceProfileResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(v2.ApiDeviceProfileRoute, v2.Name, url.QueryEscape(name))
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllDeviceProfiles queries the device profiles with offset, and limit
func (client *DeviceProfileClient) AllDeviceProfiles(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(v2.Labels, strings.Join(labels, v2.CommaSeparator))
	}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err := utils.GetRequest(ctx, &res, client.baseUrl, v2.ApiAllDeviceProfileRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeviceProfilesByModel queries the device profiles with offset, limit and model
func (client *DeviceProfileClient) DeviceProfilesByModel(ctx context.Context, model string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(v2.ApiDeviceProfileRoute, v2.Model, url.QueryEscape(model))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeviceProfilesByManufacturer queries the device profiles with offset, limit and manufacturer
func (client *DeviceProfileClient) DeviceProfilesByManufacturer(ctx context.Context, manufacturer string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(v2.ApiDeviceProfileRoute, v2.Manufacturer, url.QueryEscape(manufacturer))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeviceProfilesByManufacturerAndModel queries the device profiles with offset, limit, manufacturer and model
func (client *DeviceProfileClient) DeviceProfilesByManufacturerAndModel(ctx context.Context, manufacturer string, model string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(v2.ApiDeviceProfileRoute, v2.Manufacturer, url.QueryEscape(manufacturer), v2.Model, url.QueryEscape(model))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeviceResourceByProfileNameAndResourceName queries the device resource by profileName and resourceName
func (client *DeviceProfileClient) DeviceResourceByProfileNameAndResourceName(ctx context.Context, profileName string, resourceName string) (res responses.DeviceResourceResponse, edgexError errors.EdgeX) {
	resourceMapKey := fmt.Sprintf("%s:%s", profileName, resourceName)
	res, exists := client.resourceByMapKey(resourceMapKey)
	if exists {
		return res, nil
	}
	requestPath := path.Join(v2.ApiDeviceResourceRoute, v2.Profile, url.QueryEscape(profileName), v2.Resource, url.QueryEscape(resourceName))
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	client.setResourceWithMapKey(res, resourceMapKey)
	return res, nil
}

func (client *DeviceProfileClient) resourceByMapKey(key string) (res responses.DeviceResourceResponse, exists bool) {
	client.mux.RLock()
	defer client.mux.RUnlock()
	res, exists = client.resourcesCache[key]
	return
}

func (client *DeviceProfileClient) setResourceWithMapKey(res responses.DeviceResourceResponse, key string) {
	client.mux.Lock()
	defer client.mux.Unlock()
	client.resourcesCache[key] = res
}

func (client *DeviceProfileClient) CleanResourcesCache() {
	client.mux.Lock()
	defer client.mux.Unlock()
	client.resourcesCache = make(map[string]responses.DeviceResourceResponse)
}
