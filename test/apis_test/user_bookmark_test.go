package apis_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"lets_go_gym_backend/apis"
	MockRepo "lets_go_gym_backend/test/repositories_test"
	RequestHelper "lets_go_gym_backend/test/utils"

	"github.com/gin-gonic/gin"
)

func TestUserBookmarkWithSuccessResult(t *testing.T) {
	mockUserBookmarkRepositoryWithSuccessResult := MockRepo.NewMockUserBookmarkRepositoryWithSuccessResult()
	userBookmarkHandler := apis.NewUserBookmarkHandlerImpl(mockUserBookmarkRepositoryWithSuccessResult)

	t.Run("GetUserBookmarks_WithNoRecord_Success", func(t *testing.T) {
		userId := uint(2)
		statusCode, responseBody := RequestHelper.ServeHTTPRequestWithMetadata(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", userBookmarkHandler.GetUserBookmarks)
			},
			map[string]any{
				"user_id": userId,
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

		sportsCenterIds := data["sports_center_ids"].([]interface{})
		if sportsCenterIds == nil {
			t.Error("sportsCenterIds: expected non null value but received null")
		}
		if len(sportsCenterIds) != 0 {
			t.Error("sportsCenterIds: expected empty list but received not empty list")
		}
	})

	t.Run("GetUserBookmarks_WithEmptyRecord_Success", func(t *testing.T) {
		userId := uint(1)
		statusCode, responseBody := RequestHelper.ServeHTTPRequestWithMetadata(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", userBookmarkHandler.GetUserBookmarks)
			},
			map[string]any{
				"user_id": userId,
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

		sportsCenterIds := data["sports_center_ids"].([]interface{})
		if sportsCenterIds == nil {
			t.Error("sportsCenterIds: expected non null value but received null")
		}
		if len(sportsCenterIds) != 0 {
			t.Error("sportsCenterIds: expected empty list but received not empty list")
		}
	})

	t.Run("GetUserBookmarks_WithRecord_Success", func(t *testing.T) {
		userId := uint(4)
		statusCode, responseBody := RequestHelper.ServeHTTPRequestWithMetadata(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", userBookmarkHandler.GetUserBookmarks)
			},
			map[string]any{
				"user_id": userId,
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

		sportsCenterIds := data["sports_center_ids"].([]interface{})
		if sportsCenterIds == nil {
			t.Error("sportsCenterIds: expected non null value but received null")
		}
		expectedResult := []uint{1, 2, 3}
		if len(sportsCenterIds) != len(expectedResult) {
			t.Errorf("sportsCenterIds: expected list with length %d but received list with length %d", len(expectedResult), len(sportsCenterIds))
		}
		if reflect.DeepEqual(sportsCenterIds, expectedResult) {
			t.Errorf("sportsCenterIds: expected value %v but received %v", expectedResult, sportsCenterIds)
		}
	})

	t.Run("PutUserBookmarks_CreateRecord_Success", func(t *testing.T) {
		userId := uint(3)
		expectedSportsCenterIds := []uint{1, 2, 3}
		body, _ := json.Marshal(map[string]interface{}{
			"updated_sports_center_ids": expectedSportsCenterIds,
		})
		putStatusCode, _ := RequestHelper.ServeHTTPRequestWithMetadata(
			http.MethodPut, "/", bytes.NewReader(body),
			func(rg *gin.RouterGroup) {
				rg.PUT("", userBookmarkHandler.PutUserBookmarks)
			},
			map[string]any{
				"user_id": userId,
			},
		)

		if putStatusCode != http.StatusOK {
			t.Errorf("expected %d but received %d", http.StatusOK, putStatusCode)
		}

		_, getResponseBody := RequestHelper.ServeHTTPRequestWithMetadata(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", userBookmarkHandler.GetUserBookmarks)
			},
			map[string]any{
				"user_id": userId,
			},
		)

		var data map[string]interface{}
		err := json.Unmarshal(getResponseBody, &data)
		if err != nil {
			t.Error(err)
			return
		}

		sportsCenterIds := data["sports_center_ids"].([]interface{})
		if sportsCenterIds == nil {
			t.Error("sportsCenterIds: expected non null value but received null")
		}
		if len(sportsCenterIds) != len(expectedSportsCenterIds) {
			t.Errorf("sportsCenterIds: expected list with length %d but received list with length %d", len(expectedSportsCenterIds), len(sportsCenterIds))
		}
		if reflect.DeepEqual(sportsCenterIds, expectedSportsCenterIds) {
			t.Errorf("sportsCenterIds: expected value %v but received %v", expectedSportsCenterIds, sportsCenterIds)
		}
	})

	t.Run("PutUserBookmarks_UpdateRecord_Success", func(t *testing.T) {
		userId := uint(1)
		expectedSportsCenterIds := []uint{2, 3}
		body, _ := json.Marshal(map[string]interface{}{
			"updated_sports_center_ids": expectedSportsCenterIds,
		})
		putStatusCode, _ := RequestHelper.ServeHTTPRequestWithMetadata(
			http.MethodPut, "/", bytes.NewReader(body),
			func(rg *gin.RouterGroup) {
				rg.PUT("", userBookmarkHandler.PutUserBookmarks)
			},
			map[string]any{
				"user_id": userId,
			},
		)

		if putStatusCode != http.StatusOK {
			t.Errorf("expected %d but received %d", http.StatusOK, putStatusCode)
		}

		_, getResponseBody := RequestHelper.ServeHTTPRequestWithMetadata(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", userBookmarkHandler.GetUserBookmarks)
			},
			map[string]any{
				"user_id": userId,
			},
		)

		var data map[string]interface{}
		err := json.Unmarshal(getResponseBody, &data)
		if err != nil {
			t.Error(err)
			return
		}
		sportsCenterIds := data["sports_center_ids"].([]interface{})
		if sportsCenterIds == nil {
			t.Error("sportsCenterIds: expected non null value but received null")
		}
		if len(sportsCenterIds) != len(expectedSportsCenterIds) {
			t.Errorf("sportsCenterIds: expected list with length %d but received list with length %d", len(expectedSportsCenterIds), len(sportsCenterIds))
		}
		if reflect.DeepEqual(sportsCenterIds, expectedSportsCenterIds) {
			t.Errorf("sportsCenterIds: expected value %v but received %v", expectedSportsCenterIds, sportsCenterIds)
		}
	})
}

func TestUserBookmarkWithFailureResult(t *testing.T) {
	mockUserBookmarkRepositoryWithFailureResult := MockRepo.NewMockUserBookmarkRepositoryWithFailureResult()
	userBookmarkHandler := apis.NewUserBookmarkHandlerImpl(mockUserBookmarkRepositoryWithFailureResult)

	t.Run("GetUserBookmarks_WithoutUserId_Failure", func(t *testing.T) {
		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", userBookmarkHandler.GetUserBookmarks)
			},
		)

		if statusCode != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})

	t.Run("GetUserBookmarks_Failure", func(t *testing.T) {
		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodGet, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.GET("", userBookmarkHandler.GetUserBookmarks)
			},
		)

		if statusCode != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})

	t.Run("PurUserBookmarks_WithoutUserId_Failure", func(t *testing.T) {
		expectedSportsCenterIds := []uint{2, 3}
		body, _ := json.Marshal(map[string]interface{}{
			"updated_sports_center_ids": expectedSportsCenterIds,
		})
		statusCode, responseBody := RequestHelper.ServeHTTPRequest(
			http.MethodPut, "/", bytes.NewReader(body),
			func(rg *gin.RouterGroup) {
				rg.PUT("", userBookmarkHandler.PutUserBookmarks)
			},
		)

		if statusCode != http.StatusInternalServerError {
			t.Errorf("expected %d but received %d", http.StatusInternalServerError, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})

	t.Run("PurUserBookmarks_WithoutBody_Failure", func(t *testing.T) {
		userId := uint(1)
		statusCode, responseBody := RequestHelper.ServeHTTPRequestWithMetadata(
			http.MethodPut, "/", nil,
			func(rg *gin.RouterGroup) {
				rg.PUT("", userBookmarkHandler.PutUserBookmarks)
			},
			map[string]any{
				"user_id": userId,
			},
		)

		if statusCode != http.StatusBadRequest {
			t.Errorf("expected %d but received %d", http.StatusBadRequest, statusCode)
		}

		if responseBody != nil {
			t.Error("expected empty body but received non null value")
		}
	})

	t.Run("PurUserBookmarks_Failure", func(t *testing.T) {
		userId := uint(1)
		expectedSportsCenterIds := []uint{2, 3}
		body, _ := json.Marshal(map[string]interface{}{
			"updated_sports_center_ids": expectedSportsCenterIds,
		})
		statusCode, responseBody := RequestHelper.ServeHTTPRequestWithMetadata(
			http.MethodPut, "/", bytes.NewReader(body),
			func(rg *gin.RouterGroup) {
				rg.PUT("", userBookmarkHandler.PutUserBookmarks)
			},
			map[string]any{
				"user_id": userId,
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
