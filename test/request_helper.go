package test

import (
	"io"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func ServeHTTPRequest(method string, target string, body io.Reader, registerRoutesFunc func(*gin.RouterGroup)) (int, []byte) {
	req := httptest.NewRequest(method, target, body)
	rr := httptest.NewRecorder()
	router := gin.Default()

	registerRoutesFunc(&router.RouterGroup)

	router.ServeHTTP(rr, req)

	return rr.Code, rr.Body.Bytes()
}
