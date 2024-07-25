package common

const (
	RESULT_OK = iota
	RESULT_ERROR
	RESULT_PARAM_ERROR
	RESULT_NO_AUTHORITY_ERROR
)

type Response struct {
	Status  bool
	Message interface{}
	Data    interface{}
	Code    int
}

func ResponseOK(data interface{}) Response {
	return Response{Status: true, Data: data, Code: RESULT_OK}
}

func ResponseErr(code int, message interface{}) Response {
	return Response{Status: false, Message: message, Code: code}
}
