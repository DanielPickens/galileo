/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package handlers

import (
	"github.com/google/uuid"
)

func GetUUIDParam(param string) (uuid.UUID, error) {
	return uuid.Parse(param)
}

type ApiResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Payload interface{} `json:"payload"`
}

func BuildResponse(code string, message string, errors []string, payload interface{}) *ApiResponse {
	r := &ApiResponse{}
	r.Code = code
	r.Message = message
	r.Errors = errors
	r.Payload = payload
	return r
}

type FieldValidationError struct {
	Namespace string      `json:"namespace"`
	Field     string      `json:"field"`
	Error     string      `json:"error"`
	Kind      string      `json:"kind"`
	Value     interface{} `json:"value"`
}
