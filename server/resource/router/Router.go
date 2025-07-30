package router

import "net/http"

type Router struct {
	mux *http.ServeMux
	routes map[string]map[MethodHttp]http.HandlerFunc
}

func CreateRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
		routes: make(map[string]map[MethodHttp]http.HandlerFunc),
	}
}

// TO-DO: CREATE VALIDATION IF NAME PATH ("/") AND METHODPATH('GET','POST'...) ALREADY EXISTS
func (router *Router) RegisterRoute(methodHttp MethodHttp, path string, handlerFunc http.HandlerFunc) {
	if router.routes[path] == nil {
		router.routes[path] = make(map[MethodHttp]http.HandlerFunc)
	}

	router.routes[path][methodHttp] = handlerFunc

	router.mux.HandleFunc(path, func(response http.ResponseWriter, request *http.Request){
		if handler, ok := router.routes[path][MethodHttp(request.Method)]; ok {
			handler(response, request)
		}else {
			http.Error(response, "Métódo não permitido", http.StatusMethodNotAllowed)
		}
	})
}

func (router *Router) GetServerHttp() http.Handler {
	return router.mux
}