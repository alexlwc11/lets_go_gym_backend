package apis

/*
	This file is just for generating the swagger documents.
	Please refer to [github.com/alexlwc11/simple_auth_go] module for details
*/

// Register godoc
//
//	@Summary		Register
//	@Description	New user registration
//	@Tags			Auth
//	@Accept			json
//	@Param			user_info	body	userInfoInDto	true	"User info for registration"
//	@Produce		json
//	@Success		200	{object}	sessionTokenOutDto
//	@Failure		500
//	@Router			/register [post]
func Register() {}

// SignIn godoc
//
//	@Summary		Sign in
//	@Description	Existing user sign in
//	@Tags			Auth
//	@Accept			json
//	@Param			user_info	body	userInfoInDto	true	"User info for signing in"
//	@Produce		json
//	@Success		200	{object}	sessionTokenOutDto
//	@Failure		500
//	@Router			/sign_in [post]
func SignIn() {}

// Refresh godoc
//
//	@Summary		Refresh
//	@Description	Get new set of tokens with refresh token
//	@Tags			Auth
//	@Accept			json
//	@Param			refresh_token	body	refreshInDto	true	"Refresh token"
//	@Produce		json
//	@Success		200	{object}	sessionTokenOutDto
//	@Failure		500
//	@Router			/refresh [post]
func Refresh() {}
