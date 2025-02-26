package handler

import (
	"log"
	"music_service/internal/config"
	"music_service/internal/context"
	"music_service/internal/user"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var cfg *config.Config
var types map[string]bool

var uHdl *user.Handler
var urlMap map[string]map[string]reflect.Value

func init() {
	cfg = config.Get()

	types = make(map[string]bool)
	types[".css"] = true
	types[".js"] = true
	types[".ico"] = true
	types[".jpg"] = true
	types[".png"] = true

	urlMap = make(map[string]map[string]reflect.Value)
	urlMap["POST"] = make(map[string]reflect.Value)
	urlMap["GET"] = make(map[string]reflect.Value)
	urlMap["PUT"] = make(map[string]reflect.Value)
	urlMap["DELETE"] = make(map[string]reflect.Value)

	mapsHdl := cfg.Handlers

	uHdl = &user.Handler{}
	structHdl := reflect.TypeOf(uHdl)

	for methodNum := 0; methodNum < structHdl.NumMethod(); methodNum++ {
		method := structHdl.Method(methodNum)
		val, ok := mapsHdl[method.Name]
		if !ok {
			log.Println("Error reading handler ", method.Name, methodNum)
			continue
		}

		urlMap[val.Method][val.Url] = reflect.ValueOf(uHdl).MethodByName(method.Name)
	}
	log.Println("urlMap has been read")
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	ctx := &context.Context{
		Response: w,
		Request:  r,
	}

	path := r.URL.Path[1:]
	log.Println("Path:", r.URL.Path)

	//if ok := static(path); ok {
	//	user.SetHeaders(ctx)
	//	http.ServeFile(ctx.Response, ctx.Request, "./internal/static/"+path)
	//	return
	//}

	methodMap, ok := urlMap[ctx.Request.Method]
	if !ok {
		http.Error(ctx.Response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	method, ok := methodMap[path]
	if !ok {
		errorPath(ctx, 404)
		return
	}

	log.Println("method: ", method)
	method.Call([]reflect.Value{reflect.ValueOf(ctx)})
}

func errorPath(ctx *context.Context, statusCode int) {
	//user.SetHeaders(ctx)
	//http.ServeFile(ctx.Response, ctx.Request, "./internal/static/html/errors/"+strconv.Itoa(statusCode)+".html")
	//ctx.Response.WriteHeader(statusCode)
	_, err := ctx.Response.Write([]byte(strconv.Itoa(statusCode)))
	if err != nil {
		log.Println(err)
	}
}

func static(path string) bool {
	splitPath := strings.Split(path, "/")
	fileName := splitPath[len(splitPath)-1]
	splitName := strings.Split(fileName, ".")
	fileExt := "." + splitName[len(splitName)-1]
	if types[fileExt] {
		return true
	}
	return false
}
