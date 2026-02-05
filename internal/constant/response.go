package servicesconstant

type SuccessMessage string

const (
	SuccessMessage_Success       SuccessMessage = "successfully"
	SuccessMessage_Created       SuccessMessage = "successfully created"
	SuccessMessage_Updated       SuccessMessage = "successfully updated"
	SuccessMessage_Deleted       SuccessMessage = "successfully deleted"
	SuccessMessage_ResetPassword SuccessMessage = "A password reset email has already been sent. Please check your inbox and follow the instructions to reset your password."
	SuccessMessage_Logout        SuccessMessage = "Logout Successfully"

	SuccessCode               = "0000"
	SuccessCode_Created       = "0001"
	SuccessCode_Updated       = "0002"
	SuccessCode_Deleted       = "0003"
	SuccessCode_ResetPassword = "0004"
	SuccessCode_Logout        = "0005"
)
