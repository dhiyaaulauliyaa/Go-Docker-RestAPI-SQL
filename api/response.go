package api

type Response struct {
	Message string `json:"message" default:"Success"`
	Error   error  `json:"error"`
	Data    any    `json:"data"`
}

func errorResponse(err error, message string) Response {
	return Response{
		Message: message,
		Error:   err,
	}
}

func successResponse(Data any) Response {
	return Response{
		Message: "Success",
		Data:    Data,
	}
}
