package apis_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"lets_go_gym_backend/apis"
	MockRepo "lets_go_gym_backend/test/repositories_test"
	RequestHelper "lets_go_gym_backend/test/utils"

	"github.com/gin-gonic/gin"
)

func TestSportsCenter_WithoutAuth(t *testing.T) {
	t.Run("GetSportsCenters_Success", func(t *testing.T) {
		mockSportsCenterRepositoryWithSuccessResult := MockRepo.NewMockSportsCenterRepositoryWithSuccessResult()
		sportsCenterHandler := apis.NewSportsCenterHandlerImpl(mockSportsCenterRepositoryWithSuccessResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", sportsCenterHandler.GetAllSportsCenters)
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

		sportsCenters := data["sports_centers"].([]interface{})
		if sportsCenters == nil {
			t.Error("sportsCenters: expected non null value but received null")
		} else if len(sportsCenters) == 0 {
			t.Error("sportsCenters: expected non empty array but received empty array")
			return
		}

		sportsCenter := sportsCenters[0].(map[string]interface{})
		if sportsCenter["id"] == nil {
			t.Error("id: expected non null value but received null")
		}
		if sportsCenter["external_id"] == nil {
			t.Error("external_id: expected non null value but received null")
		}
		if sportsCenter["district_id"] == nil {
			t.Error("district_id: expected non null value but received null")
		}
		if sportsCenter["name_en"] == nil {
			t.Error("name_en: expected non null value but received null")
		}
		if sportsCenter["name_zh"] == nil {
			t.Error("name_zh: expected non null value but received null")
		}
		if sportsCenter["address_en"] == nil {
			t.Error("address_en: expected non null value but received null")
		}
		if sportsCenter["address_zh"] == nil {
			t.Error("address_zh: expected non null value but received null")
		}
		if sportsCenter["phone_numbers"] == nil {
			t.Error("phone_numbers: expected non null value but received null")
		}
		if sportsCenter["latitude_dms"] == nil {
			t.Error("latitude_dms: expected non null value but received null")
		}
		if sportsCenter["longitude_dms"] == nil {
			t.Error("longitude_dms: expected non null value but received null")
		}
		if sportsCenter["latitude_dd"] == nil {
			t.Error("latitude_dd: expected non null value but received null")
		}
		if sportsCenter["longitude_dd"] == nil {
			t.Error("longitude_dd: expected non null value but received null")
		}
	})

	t.Run("GetSportsCenters_Failure", func(t *testing.T) {
		mockSportsCenterRepositoryWithFailureResult := MockRepo.NewMockSportsCenterRepositoryWithFailureResult()
		sportsCenterHandler := apis.NewSportsCenterHandlerImpl(mockSportsCenterRepositoryWithFailureResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", sportsCenterHandler.GetAllSportsCenters)
			},
		)

		if statusCode != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})

	t.Run("GetDetailsUrl_Success", func(t *testing.T) {
		mockSportsCenterRepositoryWithSuccessResult := MockRepo.NewMockSportsCenterRepositoryWithSuccessResult()
		sportsCenterHandler := apis.NewSportsCenterHandlerImpl(mockSportsCenterRepositoryWithSuccessResult)

		sportsCenterId := 1
		target := fmt.Sprintf("/%d/details_url", sportsCenterId)
		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, target, nil, func(rg *gin.RouterGroup) {
				rg.GET("/:id/details_url", sportsCenterHandler.GetDetailsUrl)
			},
		)

		if statusCode != http.StatusOK {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, statusCode)
		}

		var data map[string]interface{}
		err := json.Unmarshal(responseBody, &data)
		if err != nil {
			t.Error(err)
			return
		}

		url := data["url"]
		expectedUrl := "https://www.lcsd.gov.hk/clpss/tc/webApp/FitnessRoomDetails.do?id=123"
		if url != expectedUrl {
			t.Errorf("expected %s but received %s", expectedUrl, url)
		}
	})

	t.Run("GetDetailsUrl_Failure_WithValidIdProvided", func(t *testing.T) {
		mockSportsCenterRepositoryWithFailureResult := MockRepo.NewMockSportsCenterRepositoryWithFailureResult()
		sportsCenterHandler := apis.NewSportsCenterHandlerImpl(mockSportsCenterRepositoryWithFailureResult)

		sportsCenterId := 1
		target := fmt.Sprintf("/%d/details_url", sportsCenterId)
		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, target, nil,
			func(rg *gin.RouterGroup) {
				rg.GET("/:id/details_url", sportsCenterHandler.GetDetailsUrl)
			},
		)

		if statusCode != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})

	t.Run("GetDetailsUrl_Failure_WithInvalidIdProvided", func(t *testing.T) {
		mockSportsCenterRepositoryWithFailureResult := MockRepo.NewMockSportsCenterRepositoryWithFailureResult()
		sportsCenterHandler := apis.NewSportsCenterHandlerImpl(mockSportsCenterRepositoryWithFailureResult)

		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, "/a/details_url", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("/:id/details_url", sportsCenterHandler.GetDetailsUrl)
			},
		)

		if statusCode != http.StatusNotFound {
			t.Errorf("expected %d but received %d", http.StatusNotFound, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})
}
