package response

import (
	"net/http"
)

type (
	Message map[string]interface{}

	Response struct {
		Status  int            `json:"status"`
		Message Message        `json:"message"`
		Errors  []CaptureError `json:"errors,omitempty"`
		Data    interface{}    `json:"data,omitempty"`
		Meta    interface{}    `json:"meta,omitempty"`
		Header  http.Header    `json:"header,omitempty"`
		Body    interface{}    `json:"body,omitempty"`
	}

	CaptureError struct {
		Details string `json:"details"`
		Message string `json:"message"`
	}
)

var (
	Text = http.StatusText

	MsgSuccess = map[string]interface{}{"status": true, "en": "Success", "id": "Sukses"}
	MsgFailed  = map[string]interface{}{"status": false, "en": "Failed", "id": "Gagal"}
)

func ResponsSuccess(statusCode int, message Message, data interface{}) Response {
	return Response{
		Status:  statusCode,
		Message: MsgSuccess,
		Data:    data,
	}
}

func ResponseError(statusCode int, messageStatus Message, details string) Response {
	return Response{
		Status:  statusCode,
		Message: messageStatus,
		Errors: []CaptureError{
			{
				Message: Text(statusCode),
				Details: details,
			},
		},
	}
}
