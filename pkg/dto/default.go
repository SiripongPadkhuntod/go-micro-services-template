package dto

import servicesconstant "go-micro-services-template/internal/constant"

type EmptyStruct struct{}
type PingRequest struct {
	BodyField  string `json:"body_field" `          // Comes from JSON body
	ParamField string `json:"-" uri:"param-field" ` // Comes from URI params
	QueryField string `json:"-" form:"query-field"` // Comes from query string
}
type SuccessResponse struct {
	Code    string                          `json:"code" default:"0000"`
	Message servicesconstant.SuccessMessage `json:"description"`
}

// var (
// 	SuccessResponse_Success       = SuccessResponse{Message: servicesconstant.SuccessMessage_Success, Code: servicesconstant.SuccessCode}
// 	SuccessResponse_Created       = SuccessResponse{Message: servicesconstant.SuccessMessage_Created, Code: servicesconstant.SuccessCode_Created}
// 	SuccessResponse_Updated       = SuccessResponse{Message: servicesconstant.SuccessMessage_Updated, Code: servicesconstant.SuccessCode_Updated}
// 	SuccessResponse_Deleted       = SuccessResponse{Message: servicesconstant.SuccessMessage_Deleted, Code: servicesconstant.SuccessCode_Deleted}
// 	SuccessResponse_ResetPassword = SuccessResponse{Message: servicesconstant.SuccessMessage_ResetPassword, Code: servicesconstant.SuccessCode_ResetPassword}
// 	SuccessResponse_Logout        = SuccessResponse{Message: servicesconstant.SuccessMessage_Logout, Code: servicesconstant.SuccessCode_Logout}
// )

// type ErrorResponse struct {
// 	Code    string                        `json:"code" default:"9999"`
// 	Message servicesconstant.ErrorMessage `json:"description"`
// }

// var (
// 	ErrorResponse_RefreshTokenNotFound = ErrorResponse{Message: servicesconstant.ErrorMessage_RefreshTokenNotFound, Code: servicesconstant.ErrorCodeRefreshTokenNotFound}
// 	ErrorResponse_NoContent            = ErrorResponse{Message: servicesconstant.ErrorMessage_NoContent, Code: servicesconstant.ErrorCodeNoContent}
// 	ErrorResponse_BuildLogoutRequest   = ErrorResponse{Message: servicesconstant.ErrorMessage_BuildLogoutRequest, Code: servicesconstant.ErrorCodeBuildLogoutRequest}
// 	ErrorResponse_ConnectionIDP        = ErrorResponse{Message: servicesconstant.ErrorMessage_ConnectionIDP, Code: servicesconstant.ErrorCodeConnectionIDP}
// 	ErrorResponse_ReturnFromIDP        = ErrorResponse{Message: servicesconstant.ErrorMessage_ReturnFromIDP, Code: servicesconstant.ErrorCodeReturnFromIDP}
// 	ErrorResponse_BuildRefreshRequest  = ErrorResponse{Message: servicesconstant.ErrorMessage_BuildRefreshRequest, Code: servicesconstant.ErrorCodeBuildRefreshRequest}
// 	ErrorResponse_RefreshTokenFailed   = ErrorResponse{Message: servicesconstant.ErrorMessage_RefreshTokenFailed, Code: servicesconstant.ErrorCodeRefreshTokenFailed}
// 	ErrorResponse_RefreshDecodeFailed  = ErrorResponse{Message: servicesconstant.ErrorMessage_RefreshDecodeFailed, Code: servicesconstant.ErrorCodeRefreshDecodeFailed}
// )

// type BaseResponse struct {
// 	Code string `json:"code" default:"0000"`
// 	// SuccessMessage servicesconstant.SuccessMessage `json:"description_success" omitempty`
// 	// ErrorMessage   servicesconstant.ErrorMessage   `json:"description_error" omitempty`
// 	Message interface{} `json:"description,omitempty"`
// }

// var (
// 	BaseResponse_Created       = BaseResponse{Message: servicesconstant.SuccessMessage_Created, Code: servicesconstant.SuccessCode}
// 	BaseResponse_Updated       = BaseResponse{Message: servicesconstant.SuccessMessage_Updated, Code: servicesconstant.SuccessCode}
// 	BaseResponse_Deleted       = BaseResponse{Message: servicesconstant.SuccessMessage_Deleted, Code: servicesconstant.SuccessCode}
// 	BaseResponse_ResetPassword = BaseResponse{Message: servicesconstant.SuccessMessage_ResetPassword, Code: servicesconstant.SuccessCode}
// 	BaseResponse_Logout        = BaseResponse{Message: servicesconstant.SuccessMessage_Logout, Code: servicesconstant.SuccessCode}

// 	BaseResponse_RefreshTokenNotFound = BaseResponse{Message: servicesconstant.ErrorMessage_RefreshTokenNotFound, Code: servicesconstant.ErrorCodeRefreshTokenNotFound}
// 	BaseResponse_NoContent            = BaseResponse{Message: servicesconstant.ErrorMessage_NoContent, Code: servicesconstant.ErrorCodeNoContent}
// 	BaseResponse_BuildLogoutRequest   = BaseResponse{Message: servicesconstant.ErrorMessage_BuildLogoutRequest, Code: servicesconstant.ErrorCodeBuildLogoutRequest}
// 	BaseResponse_ConnectionIDP        = BaseResponse{Message: servicesconstant.ErrorMessage_ConnectionIDP, Code: servicesconstant.ErrorCodeConnectionIDP}
// 	BaseResponse_ReturnFromIDP        = BaseResponse{Message: servicesconstant.ErrorMessage_ReturnFromIDP, Code: servicesconstant.ErrorCodeReturnFromIDP}
// 	BaseResponse_BuildRefreshRequest  = BaseResponse{Message: servicesconstant.ErrorMessage_BuildRefreshRequest, Code: servicesconstant.ErrorCodeBuildRefreshRequest}
// 	BaseResponse_RefreshTokenFailed   = BaseResponse{Message: servicesconstant.ErrorMessage_RefreshTokenFailed, Code: servicesconstant.ErrorCodeRefreshTokenFailed}
// 	BaseResponse_RefreshDecodeFailed  = BaseResponse{Message: servicesconstant.ErrorMessage_RefreshDecodeFailed, Code: servicesconstant.ErrorCodeRefreshDecodeFailed}
// 	BaseResponse_InternalServerError  = BaseResponse{Message: servicesconstant.ErrorMessage_EmptyIDPUrl, Code: servicesconstant.ErrorCodeInternalServerError}
// )
