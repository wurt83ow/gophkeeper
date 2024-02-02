// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// PostAddDataTableUserIDJSONBody defines parameters for PostAddDataTableUserID.
type PostAddDataTableUserIDJSONBody map[string]string

// PostLoginJSONBody defines parameters for PostLogin.
type PostLoginJSONBody struct {
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

// PostRegisterJSONBody defines parameters for PostRegister.
type PostRegisterJSONBody struct {
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

// PutUpdateDataTableUserIDIdJSONBody defines parameters for PutUpdateDataTableUserIDId.
type PutUpdateDataTableUserIDIdJSONBody map[string]string

// PostAddDataTableUserIDJSONRequestBody defines body for PostAddDataTableUserID for application/json ContentType.
type PostAddDataTableUserIDJSONRequestBody PostAddDataTableUserIDJSONBody

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody PostLoginJSONBody

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody PostRegisterJSONBody

// PutUpdateDataTableUserIDIdJSONRequestBody defines body for PutUpdateDataTableUserIDId for application/json ContentType.
type PutUpdateDataTableUserIDIdJSONRequestBody PutUpdateDataTableUserIDIdJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /addData/{table}/{userID})
	PostAddDataTableUserID(w http.ResponseWriter, r *http.Request, table string, userID int)

	// (DELETE /clearData/{table}/{userID})
	DeleteClearDataTableUserID(w http.ResponseWriter, r *http.Request, table string, userID int)

	// (DELETE /deleteData/{table}/{userID}/{id})
	DeleteDeleteDataTableUserIDId(w http.ResponseWriter, r *http.Request, table string, userID int, id string)

	// (GET /getAllData/{table}/{userID})
	GetGetAllDataTableUserID(w http.ResponseWriter, r *http.Request, table string, userID int)

	// (GET /getData/{table}/{userID})
	GetGetDataTableUserID(w http.ResponseWriter, r *http.Request, table string, userID int)

	// (GET /getPassword/{username})
	GetGetPasswordUsername(w http.ResponseWriter, r *http.Request, username string)

	// (GET /getUserID/{username})
	GetGetUserIDUsername(w http.ResponseWriter, r *http.Request, username string)

	// (POST /login)
	PostLogin(w http.ResponseWriter, r *http.Request)

	// (POST /register)
	PostRegister(w http.ResponseWriter, r *http.Request)

	// (POST /sendFile/{userID})
	PostSendFileUserID(w http.ResponseWriter, r *http.Request, userID int)

	// (PUT /updateData/{table}/{userID}/{id})
	PutUpdateDataTableUserIDId(w http.ResponseWriter, r *http.Request, table string, userID int, id int)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (POST /addData/{table}/{userID})
func (_ Unimplemented) PostAddDataTableUserID(w http.ResponseWriter, r *http.Request, table string, userID int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (DELETE /clearData/{table}/{userID})
func (_ Unimplemented) DeleteClearDataTableUserID(w http.ResponseWriter, r *http.Request, table string, userID int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (DELETE /deleteData/{table}/{userID}/{id})
func (_ Unimplemented) DeleteDeleteDataTableUserIDId(w http.ResponseWriter, r *http.Request, table string, userID int, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /getAllData/{table}/{userID})
func (_ Unimplemented) GetGetAllDataTableUserID(w http.ResponseWriter, r *http.Request, table string, userID int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /getData/{table}/{userID})
func (_ Unimplemented) GetGetDataTableUserID(w http.ResponseWriter, r *http.Request, table string, userID int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /getPassword/{username})
func (_ Unimplemented) GetGetPasswordUsername(w http.ResponseWriter, r *http.Request, username string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /getUserID/{username})
func (_ Unimplemented) GetGetUserIDUsername(w http.ResponseWriter, r *http.Request, username string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /login)
func (_ Unimplemented) PostLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /register)
func (_ Unimplemented) PostRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /sendFile/{userID})
func (_ Unimplemented) PostSendFileUserID(w http.ResponseWriter, r *http.Request, userID int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PUT /updateData/{table}/{userID}/{id})
func (_ Unimplemented) PutUpdateDataTableUserIDId(w http.ResponseWriter, r *http.Request, table string, userID int, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostAddDataTableUserID operation middleware
func (siw *ServerInterfaceWrapper) PostAddDataTableUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "table" -------------
	var table string

	err = runtime.BindStyledParameterWithOptions("simple", "table", chi.URLParam(r, "table"), &table, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "table", Err: err})
		return
	}

	// ------------- Path parameter "userID" -------------
	var userID int

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostAddDataTableUserID(w, r, table, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteClearDataTableUserID operation middleware
func (siw *ServerInterfaceWrapper) DeleteClearDataTableUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "table" -------------
	var table string

	err = runtime.BindStyledParameterWithOptions("simple", "table", chi.URLParam(r, "table"), &table, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "table", Err: err})
		return
	}

	// ------------- Path parameter "userID" -------------
	var userID int

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteClearDataTableUserID(w, r, table, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteDeleteDataTableUserIDId operation middleware
func (siw *ServerInterfaceWrapper) DeleteDeleteDataTableUserIDId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "table" -------------
	var table string

	err = runtime.BindStyledParameterWithOptions("simple", "table", chi.URLParam(r, "table"), &table, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "table", Err: err})
		return
	}

	// ------------- Path parameter "userID" -------------
	var userID int

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteDeleteDataTableUserIDId(w, r, table, userID, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetGetAllDataTableUserID operation middleware
func (siw *ServerInterfaceWrapper) GetGetAllDataTableUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "table" -------------
	var table string

	err = runtime.BindStyledParameterWithOptions("simple", "table", chi.URLParam(r, "table"), &table, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "table", Err: err})
		return
	}

	// ------------- Path parameter "userID" -------------
	var userID int

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGetAllDataTableUserID(w, r, table, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetGetDataTableUserID operation middleware
func (siw *ServerInterfaceWrapper) GetGetDataTableUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "table" -------------
	var table string

	err = runtime.BindStyledParameterWithOptions("simple", "table", chi.URLParam(r, "table"), &table, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "table", Err: err})
		return
	}

	// ------------- Path parameter "userID" -------------
	var userID int

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGetDataTableUserID(w, r, table, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetGetPasswordUsername operation middleware
func (siw *ServerInterfaceWrapper) GetGetPasswordUsername(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithOptions("simple", "username", chi.URLParam(r, "username"), &username, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "username", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGetPasswordUsername(w, r, username)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetGetUserIDUsername operation middleware
func (siw *ServerInterfaceWrapper) GetGetUserIDUsername(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithOptions("simple", "username", chi.URLParam(r, "username"), &username, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "username", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGetUserIDUsername(w, r, username)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostLogin operation middleware
func (siw *ServerInterfaceWrapper) PostLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostLogin(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostRegister operation middleware
func (siw *ServerInterfaceWrapper) PostRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostRegister(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostSendFileUserID operation middleware
func (siw *ServerInterfaceWrapper) PostSendFileUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userID" -------------
	var userID int

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostSendFileUserID(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutUpdateDataTableUserIDId operation middleware
func (siw *ServerInterfaceWrapper) PutUpdateDataTableUserIDId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "table" -------------
	var table string

	err = runtime.BindStyledParameterWithOptions("simple", "table", chi.URLParam(r, "table"), &table, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "table", Err: err})
		return
	}

	// ------------- Path parameter "userID" -------------
	var userID int

	err = runtime.BindStyledParameterWithOptions("simple", "userID", chi.URLParam(r, "userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUpdateDataTableUserIDId(w, r, table, userID, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/addData/{table}/{userID}", wrapper.PostAddDataTableUserID)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/clearData/{table}/{userID}", wrapper.DeleteClearDataTableUserID)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/deleteData/{table}/{userID}/{id}", wrapper.DeleteDeleteDataTableUserIDId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/getAllData/{table}/{userID}", wrapper.GetGetAllDataTableUserID)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/getData/{table}/{userID}", wrapper.GetGetDataTableUserID)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/getPassword/{username}", wrapper.GetGetPasswordUsername)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/getUserID/{username}", wrapper.GetGetUserIDUsername)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/login", wrapper.PostLogin)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/register", wrapper.PostRegister)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/sendFile/{userID}", wrapper.PostSendFileUserID)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/updateData/{table}/{userID}/{id}", wrapper.PutUpdateDataTableUserIDId)
	})

	return r
}
