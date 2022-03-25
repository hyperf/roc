package formatter

type JsonRPCRequest[T any, T2 any] struct {
	Id      string `json:"id"`
	Path    string `json:"path"`
	Data    T      `json:"data"`
	Context T2     `json:"context"`
}

type JsonRPCRoute struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

type JsonRPCResponse[T any, T2 any] struct {
	Id      string `json:"id"`
	Result  T      `json:"result"`
	Context T2     `json:"context"`
}

type JsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JsonRPCErrorResponse[T2 any] struct {
	Id      string        `json:"id"`
	Error   *JsonRPCError `json:"error"`
	Context T2            `json:"context"`
}
