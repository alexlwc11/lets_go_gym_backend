package apis_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"lets_go_gym_backend/apis"
	MockRepo "lets_go_gym_backend/repositories_test"
	RequestHelper "lets_go_gym_backend/test"
)

func TestAppInfo(t *testing.T) {
	t.Run("GetAppInfo_Success", func(t *testing.T) {
		mockAppVersionRepoWithSuccessResult := MockRepo.NewMockAppVersionRepositoryWithSuccessResult()
		mockDataInfoRepoWithSuccessResult := MockRepo.NewMockDataInfoRepositoryWithSuccessResult()
		appInfoHandler := apis.NewAppInfoHandler(
			mockAppVersionRepoWithSuccessResult, mockDataInfoRepoWithSuccessResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(http.MethodGet, "/app_info", nil, appInfoHandler.RegisterRoutes)

		if statusCode != http.StatusOK {
			t.Errorf("expected %d but received %d", http.StatusOK, statusCode)
		}

		var data map[string]interface{}
		err := json.Unmarshal(responseBody, &data)
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

	t.Run("GetAppInfo_FailWithAppVersionFailureResult", func(t *testing.T) {
		mockAppVersionRepoWithFailureResult := MockRepo.NewMockAppVersionRepositoryWithFailureResult()
		mockDataInfoRepoWithSuccessResult := MockRepo.NewMockDataInfoRepositoryWithSuccessResult()
		appInfoHandler := apis.NewAppInfoHandler(
			mockAppVersionRepoWithFailureResult, mockDataInfoRepoWithSuccessResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(http.MethodGet, "/app_info", nil, appInfoHandler.RegisterRoutes)

		if statusCode != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})

	t.Run("GetAppInfo_FailWithDataInfoFailureResult", func(t *testing.T) {
		mockAppVersionRepoWithSuccessResult := MockRepo.NewMockAppVersionRepositoryWithSuccessResult()
		mockDataInfoRepoWithFailureResult := MockRepo.NewMockDataInfoRepositoryWithFailureResult()
		appInfoHandler := apis.NewAppInfoHandler(
			mockAppVersionRepoWithSuccessResult, mockDataInfoRepoWithFailureResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(http.MethodGet, "/app_info", nil, appInfoHandler.RegisterRoutes)

		if statusCode != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})
}
