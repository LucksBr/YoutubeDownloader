package router

import "net/http"

type RouterInterface interface {
	RegisterRoute(methodHttp MethodHttp, path string, handlerFunc http.HandlerFunc)
	GetServerHttp() http.Handler
}

