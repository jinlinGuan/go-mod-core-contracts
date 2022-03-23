//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	//"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var emptyString = " "

var testDeviceCommand = AddDeviceCommandRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	ProfileName: TestDeviceProfileName,
	DeviceCommand: dtos.DeviceCommand{
		Name:      TestDeviceCommandName,
		ReadWrite: common.ReadWrite_RW,
		ResourceOperations: []dtos.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
		}},
	},
}

var testUpdateDeviceCommand = UpdateDeviceCommandRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	ProfileName:   TestDeviceProfileName,
	DeviceCommand: mockDeviceCommandDTO(),
}

func mockDeviceCommandDTO() dtos.UpdateDeviceCommand {
	testName := TestDeviceCommandName
	testReadWrite := common.ReadWrite_RW
	testResourceOperations := []dtos.ResourceOperation{{
		DeviceResource: TestDeviceResourceName,
	}}
	dc := dtos.UpdateDeviceCommand{}
	dc.Name = &testName
	dc.ReadWrite = &testReadWrite
	dc.ResourceOperations = testResourceOperations
	return dc
}

func TestAddDeviceCommandRequest_Validate(t *testing.T) {
	valid := testDeviceCommand
	noProfileName := testDeviceCommand
	noProfileName.ProfileName = emptyString
	noDeviceCommandName := testDeviceCommand
	noDeviceCommandName.DeviceCommand.Name = emptyString
	invalidReadWrite := testDeviceCommand
	invalidReadWrite.DeviceCommand.ReadWrite = "invalid"
	noResourceOperations := testDeviceCommand
	noResourceOperations.DeviceCommand.ResourceOperations = nil

	tests := []struct {
		name        string
		request     AddDeviceCommandRequest
		expectedErr bool
	}{
		{"valid AddDeviceCommandRequest", valid, false},
		{"invalid AddDeviceCommandRequest, no ProfileName", noProfileName, true},
		{"invalid AddDeviceCommandRequest, no DeviceCommand Name", noDeviceCommandName, true},
		{"invalid AddDeviceCommandRequest, invalid ReadWrite", invalidReadWrite, true},
		{"invalid AddDeviceCommandRequest, no ResourceOperations", noResourceOperations, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAddDeviceCommandRequest_UnmarshalJSON(t *testing.T) {
	valid := testDeviceCommand
	resultTestBytes, _ := json.Marshal(testDeviceCommand)
	type args struct {
		data []byte
	}

	tests := []struct {
		name        string
		request     AddDeviceCommandRequest
		args        args
		expectedErr bool
	}{
		{"unmarshal AddDeviceCommandRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddDeviceCommandRequest, empty data", AddDeviceCommandRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddDeviceCommandRequest, string data", AddDeviceCommandRequest{}, args{[]byte("Invalid AddDeviceCommandRequest")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.request
			err := tt.request.UnmarshalJSON(tt.args.data)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.request, "Unmarshal did not result in expected AddDeviceCommandRequest.")
			}
		})
	}
}

func TestUpdateDeviceCommandRequest_Validate(t *testing.T) {
	valid := testUpdateDeviceCommand
	noProfileName := testUpdateDeviceCommand
	noProfileName.ProfileName = emptyString
	noDeviceCommandName := testUpdateDeviceCommand
	noDeviceCommandName.DeviceCommand.Name = &emptyString

	// ReadWrite
	validNilReadWrite := valid
	validNilReadWrite.DeviceCommand.ReadWrite = nil
	// ResourceOperations
	validNilResourceOperations := valid
	validNilResourceOperations.DeviceCommand.ResourceOperations = nil

	tests := []struct {
		name        string
		request     UpdateDeviceCommandRequest
		expectedErr bool
	}{
		{"valid", valid, false},
		{"invalid - ", noProfileName, true},
		{"invalid - ", noDeviceCommandName, true},
		{"valid -", validNilReadWrite, false},
		{"valid - ", validNilResourceOperations, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateDeviceCommandRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
	    "apiVersion" : "v2",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
        "profileName": "TestProfile",
		"deviceCommand":{"name":"TestCommand"}
	}`
	var req UpdateDeviceCommandRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.DeviceCommand.ReadWrite)
	assert.Nil(t, req.DeviceCommand.IsHidden)
	assert.Nil(t, req.DeviceCommand.ResourceOperations)
}

/*
func TestUpdateDeviceCommandRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
	    "apiVersion" : "v2",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
        "profileName": "TestProfile",
		"deviceCommand":{"name":"TestCommand", "resourceOperations":[]}
	}`
    var req UpdateDeviceCommandRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.DeviceCommand.ResourceOperations)
}
*/
