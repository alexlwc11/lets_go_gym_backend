package app_info_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"lets_go_gym_backend/apis"
)

func TestGetAppInfo(t *testing.T) {
	t.Run("GetAppInfoSuccess", func(t *testing.T) {
		mockAppVersionRepoWithSuccessResult := NewMockAppVersionRepositoryWithSuccessResult()
		mockDataInfoRepoWithSuccessResult := NewMockDataInfoRepositoryWithSuccessResult()
		appInfoHandler := apis.NewAppInfoHandler(
			mockAppVersionRepoWithSuccessResult, mockDataInfoRepoWithSuccessResult)

		req := httptest.NewRequest(http.MethodGet, "/app_info", nil)
		rr := httptest.NewRecorder()
		router := gin.Default()

		appInfoHandler.RegisterRoutes(&router.RouterGroup)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected %d but received %d", http.StatusOK, rr.Code)
		}

		body := rr.Body.Bytes()

		var data map[string]interface{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			t.Error(err)
			return
		}

		if data["latest_build_version"] == nil {
			t.Error("latest_build_version: expected non null value but received null")
		}

		if data["minimum_build_version"] == nil {
			t.Error("minimum_build_version: expected non null value but received null")
		}

		dataInfo := data["data_info"].(map[string]interface{})
		if dataInfo == nil {
			t.Error("data_info: expected non null value but received null")
		}
		if dataInfo["region_data_last_updated_at"] == nil {
			t.Error("region_data_last_updated_at: expected non null value but received null")
		}
		if dataInfo["district_data_last_updated_at"] == nil {
			t.Error("district_data_last_updated_at: expected non null value but received null")
		}
		if dataInfo["sports_center_data_last_updated_at"] == nil {
			t.Error("sports_center_data_last_updated_at: expected non null value but received null")
		}
	})

	t.Run("GetAppInfoFailWithAppVersionFailureResult", func(t *testing.T) {
		mockAppVersionRepoWithFailureResult := NewMockAppVersionRepositoryWithFailureResult()
		mockDataInfoRepoWithSuccessResult := NewMockDataInfoRepositoryWithSuccessResult()
		appInfoHandler := apis.NewAppInfoHandler(
			mockAppVersionRepoWithFailureResult, mockDataInfoRepoWithSuccessResult)

		req := httptest.NewRequest(http.MethodGet, "/app_info", nil)
		rr := httptest.NewRecorder()
		router := gin.Default()

		appInfoHandler.RegisterRoutes(&router.RouterGroup)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, rr.Code)
		}

		body := rr.Body.Bytes()
		if body != nil {
			t.Error("expected empty body but received non null value")
		}
	})

	t.Run("GetAppInfoFailWithDataInfoFailureResult", func(t *testing.T) {
		mockAppVersionRepoWithSuccessResult := NewMockAppVersionRepositoryWithSuccessResult()
		mockDataInfoRepoWithFailureResult := NewMockDataInfoRepositoryWithFailureResult()
		appInfoHandler := apis.NewAppInfoHandler(
			mockAppVersionRepoWithSuccessResult, mockDataInfoRepoWithFailureResult)

		req := httptest.NewRequest(http.MethodGet, "/app_info", nil)
		rr := httptest.NewRecorder()
		router := gin.Default()

		appInfoHandler.RegisterRoutes(&router.RouterGroup)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, rr.Code)
		}

		body := rr.Body.Bytes()
		if body != nil {
			t.Error("expected empty body but received non null value")
		}
	})
}
