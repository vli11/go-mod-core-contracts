//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
)

func TestNewEventCountResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedCount := uint32(1000)
	expectedDeviceName := "device1"
	actual := NewEventCountResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedCount, expectedDeviceName)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedCount, actual.Count)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
}

func TestNewEventCountResponseNoMessage(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedCount := uint32(1000)
	expectedDeviceName := "device1"
	actual := NewEventCountResponseNoMessage(expectedRequestId, expectedStatusCode, expectedCount, expectedDeviceName)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedCount, actual.Count)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
}

func TestNewEventResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedEvent := dtos.Event{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"}
	actual := NewEventResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedEvent)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedEvent, actual.Event)
}

func TestNewEventResponseNoMessage(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedEvent := dtos.Event{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"}
	actual := NewEventResponseNoMessage(expectedRequestId, expectedStatusCode, expectedEvent)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedEvent, actual.Event)
}

func TestNewMultiEventsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedEvents := []dtos.Event{
		{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"},
		{Id: "11111111-2222-3333-4444-555555555555"},
	}
	actual := NewMultiEventsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedEvents)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedEvents, actual.Events)
}

func TestNewUpdateEventPushedByIdResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedId := "11111111-2222-3333-4444-555555555555"
	actual := NewUpdateEventPushedByIdResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedId)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedId, actual.Id)
}

func TestNewUpdateEventPushedByIdResponseNoMessage(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedId := "11111111-2222-3333-4444-555555555555"
	actual := NewUpdateEventPushedByIdResponseNoMessage(expectedRequestId, expectedStatusCode, expectedId)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedId, actual.Id)
}
