package utils

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

func ServeHTTPRequestWithMetadata(method string, target string, body io.Reader, registerRoutesFunc func(*gin.RouterGroup), metadata map[string]any) (int, []byte) {
	req := httptest.NewRequest(method, target, body)
	rr := httptest.NewRecorder()
	router := gin.Default()

	if metadata != nil && (len(metadata)) > 0 {
		router.Use(func(c *gin.Context) {
			for key, value := range metadata {
				c.Set(key, value)
			}
			c.Next()
		})
	}

	registerRoutesFunc(&router.RouterGroup)

	router.ServeHTTP(rr, req)

	return rr.Code, rr.Body.Bytes()
}
