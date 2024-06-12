package apis_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"lets_go_gym_backend/apis"
	"lets_go_gym_backend/middleware"
	MockRepo "lets_go_gym_backend/repositories_test"
	RequestHelper "lets_go_gym_backend/test"
)

// TODO look into Auth middleware
func TestUserBookmarkWithData(t *testing.T) {
	mockUserBookmarkRepositoryWithSuccessResult := MockRepo.NewMockUserBookmarkRepositoryWithSuccessResult()
	sessionTokenRepo := MockRepo.NewMockSessionTokenRepository()
	userBookmarkHandler := apis.NewUserBookmarkHandler(mockUserBookmarkRepositoryWithSuccessResult)

	// TODO wip
	t.Run("GetUserBookmarks_SuccessWithEmptyList", func(t *testing.T) {
		statusCode, responseBody := RequestHelper.ServeHTTPRequestWithMiddleware(http.MethodGet, "/", nil, userBookmarkHandler.RegisterRoutes, middleware.AuthRequired(sessionTokenRepo.FindByValue))

		if statusCode != http.StatusOK {
			t.Errorf("expected %d but received %d", http.StatusOK, statusCode)
		}

		var data map[string]interface{}
		err := json.Unmarshal(responseBody, &data)
		if err != nil {
			t.Error(err)
			return
		}
	})
}

func TestUserBookmarkWithoutData(t *testing.T) {

}
