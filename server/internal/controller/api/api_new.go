// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package api

import (
	"hotgo/api/api"
)

type ControllerMember struct{}

func NewMember() api.IApiMember {
	return &ControllerMember{}
}

type ControllerPay struct{}

func NewPay() api.IApiPay {
	return &ControllerPay{}
}

type ControllerPcclient struct{}

func NewPcclient() api.IApiPcclient {
	return &ControllerPcclient{}
}

type ControllerTest struct{}

func NewTest() api.IApiTest {
	return &ControllerTest{}
}

type ControllerUser struct{}

func NewUser() api.IApiUser {
	return &ControllerUser{}
}

