package router

type MethodHttp string

const (
	GET    MethodHttp = "GET"
	POST   MethodHttp = "POST"
	PUT    MethodHttp = "PUT"
	DELETE MethodHttp = "DELETE"
	PATCH  MethodHttp = "PATCH"
)