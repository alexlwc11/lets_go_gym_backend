package apis_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"lets_go_gym_backend/apis"
	MockRepo "lets_go_gym_backend/repositories_test"
	RequestHelper "lets_go_gym_backend/test"

	"github.com/gin-gonic/gin"
)

func TestRegions_WithoutAuth(t *testing.T) {
	t.Run("GetRegions_Success", func(t *testing.T) {
		mockRegionRepositoryWithSuccessResult := MockRepo.NewMockRegionRepositoryWithSuccessResult()
		regionHandler := apis.NewRegionHandlerImpl(mockRegionRepositoryWithSuccessResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, "/", nil, func(rg *gin.RouterGroup) {
				rg.GET("", regionHandler.GetAllRegions)
			},
		)

		if statusCode != http.StatusOK {
			t.Errorf("expected %d but received %d", http.StatusOK, statusCode)
		}

		var data map[string]interface{}
		err := json.Unmarshal(responseBody, &data)
		if err != nil {
			t.Error(err)
			return
		}

		regions := data["regions"].([]interface{})
		if regions == nil {
			t.Error("regions: expected non null value but received null")
		} else if len(regions) == 0 {
			t.Error("regions: expected non empty array but received empty array")
			return
		}

		region := regions[0].(map[string]interface{})
		if region["id"] == nil {
			t.Error("id: expected non null value but received null")
		}
		if region["code"] == nil {
			t.Error("code: expected non null value but received null")
		}
		if region["name_en"] == nil {
			t.Error("name_en: expected non null value but received null")
		}
		if region["name_zh"] == nil {
			t.Error("name_zh: expected non null value but received null")
		}
	})

	t.Run("GetRegions_Failure", func(t *testing.T) {
		mockRegionRepositoryWithFailureResult := MockRepo.NewMockRegionRepositoryWithFailureResult()
		regionHandler := apis.NewRegionHandlerImpl(mockRegionRepositoryWithFailureResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, "/", nil, func(rg *gin.RouterGroup) {
				rg.GET("", regionHandler.GetAllRegions)
			},
		)

		if statusCode != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})
}
