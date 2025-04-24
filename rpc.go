package nerve

import "encoding/json"

const JSONRPCVer = "2.0"

type RPCEvent struct {
	JSONRPC string `json:"jsonrpc"`
	Event   string `json:"event"`
	Data    any    `json:"data,omitempty"`
}

type RPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  any             `json:"result,omitempty"`
	Error   *RPCError       `json:"error,omitempty"`
}

func makeErrorResponse(id json.RawMessage, err *RPCError) *RPCResponse {
	return &RPCResponse{
		JSONRPC: JSONRPCVer,
		ID:      id,
		Error:   err,
	}
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewRPCError(code int, msg string) *RPCError {
	return &RPCError{Code: code, Message: msg}
}

type RPCErrorable interface {
	ToRPCError() *RPCError
}

var (
	ErrParse          = NewRPCError(-32700, "Parse error")
	ErrInvalidReq     = NewRPCError(-32600, "Invalid request")
	ErrMethodNotFound = NewRPCError(-32601, "Method not found")
	ErrInvalidParams  = NewRPCError(-32602, "Invalid params")
)

func NewInternalError(msg string) *RPCError {
	return &RPCError{
		Code:    -32603,
		Message: msg,
	}
}
