//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	//"encoding/json"

	//"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	//"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	//"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

//
type AddDeviceResourceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string              `json:"profileName" validate:"required"`
	Resource              dtos.DeviceResource `json:"resource"`
}

//
type UpdateDeviceResourceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string                    `json:"profileName" validate:"required"`
	Resource              dtos.UpdateDeviceResource `json:"resource"`
}

/*
func NewDeviceResourceRequest(dto dtos.DeviceResource) DeviceResourceRequest {
	return DeviceResourceRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Resource:     dto,
	}
}
*/
