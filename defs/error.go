package defs

type ErrResponse struct {
	HttpSC int
	Error  Err
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}
	ErrorDBError                = ErrResponse{HttpSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults         = ErrResponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
