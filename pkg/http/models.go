package http

type ContextKey string

const (
	QUERY = ContextKey("query")
	BODY  = ContextKey("body")
)

type Resp struct {
	Code int         `json:"-"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type ListPetsQuery struct {
	Category string `json:"category"`
}
