// package apis_test

// import (
// 	"encoding/json"
// 	"net/http"
// 	"testing"

// 	"lets_go_gym_backend/apis"
// 	MockRepo "lets_go_gym_backend/repositories_test"
// 	RequestHelper "lets_go_gym_backend/test"
// )

// // TODO look into Auth middleware
// func TestUserBookmarkWithData(t *testing.T) {
// 	mockUserBookmarkRepositoryWithSuccessResult := MockRepo.NewMockUserBookmarkRepositoryWithSuccessResult()
// 	userBookmarkHandler := apis.NewUserBookmarkHandler(mockUserBookmarkRepositoryWithSuccessResult)

// 	// TODO wip
// 	t.Run("GetUserBookmarks_SuccessWithEmptyList", func(t *testing.T) {
// 		statusCode, responseBody := RequestHelper.ServeHTTPRequest(http.MethodGet, "/", nil, userBookmarkHandler.RegisterRoutes)

// 		if statusCode != http.StatusOK {
// 			t.Errorf("expected %d but received %d", http.StatusOK, statusCode)
// 		}

// 		var data map[string]interface{}
// 		err := json.Unmarshal(responseBody, &data)
// 		if err != nil {
// 			t.Error(err)
// 			return
// 		}
// 	})
// }

// func TestUserBookmarkWithoutData(t *testing.T) {

// }
