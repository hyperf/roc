package response

type JsonRPCResponse[T any, T2 any] struct {
	Id      string `json:"id"`
	Path    string `json:"path"`
	Data    T      `json:"data"`
	Context T2     `json:"context"`
}
