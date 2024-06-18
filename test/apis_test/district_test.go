package apis_test

import (
	"encoding/json"

	"net/http"
	"testing"

	"lets_go_gym_backend/apis"
	MockRepo "lets_go_gym_backend/test/repositories_test"
	RequestHelper "lets_go_gym_backend/test/utils"

	"github.com/gin-gonic/gin"
)

func TestDistricts_WithoutAuth(t *testing.T) {
	t.Run("GetDistricts_Success", func(t *testing.T) {
		mockDistrictRepositoryWithSuccessResult := MockRepo.NewMockDistrictRepositoryWithSuccessResult()
		districtHandler := apis.NewDistrictHandlerImpl(mockDistrictRepositoryWithSuccessResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", districtHandler.GetAllDistricts)
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

		districts := data["districts"].([]interface{})
		if districts == nil {
			t.Error("districts: expected non null value but received null")
		} else if len(districts) == 0 {
			t.Error("districts: expected non empty array but received empty array")
			return
		}

		district := districts[0].(map[string]interface{})
		if district["id"] == nil {
			t.Error("id: expected non null value but received null")
		}
		if district["region_id"] == nil {
			t.Error("region_id: expected non null value but received null")
		}
		if district["name_en"] == nil {
			t.Error("name_en: expected non null value but received null")
		}
		if district["name_zh"] == nil {
			t.Error("name_zh: expected non null value but received null")
		}
	})

	t.Run("GetDistricts_Failure", func(t *testing.T) {
		mockDistrictRepositoryWithFailureResult := MockRepo.NewMockDistrictRepositoryWithFailureResult()
		districtHandler := apis.NewDistrictHandlerImpl(mockDistrictRepositoryWithFailureResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", districtHandler.GetAllDistricts)
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
