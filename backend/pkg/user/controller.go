package user

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"pockethealth/internchallenge/pkg/router"
	"regexp"
	"strings"
)

// An UserApiController binds http requests to an api service and writes the service results to the http response
type UserApiController struct {
	service UserApiService
}

// NewUserApiController creates a new api controller
func NewUserApiController(s UserApiService) router.Router {
	return UserApiController{service: s}
}

// Routes returns all of the api route for the UserApiController
func (c UserApiController) Routes() router.Routes {
	return router.Routes{
		{
			Name:        "PostRegister",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/register",
			HandlerFunc: c.PostRegister,
		},
		{
			Name:        "GetAllUsers",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/users",
			HandlerFunc: c.GetAllUsers,
		},
	}
}

type PostRegisterBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Color string `json:"Color"`
}

type PostRegisterResponse struct {
	UserId string `json:"user_id"`
}

// PostRegister - Register a New User
func (c UserApiController) PostRegister(w http.ResponseWriter, r *http.Request) {
	// read request body
	data := &PostRegisterBody{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Input Checkings
	if _, err := mail.ParseAddress(data.Email); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	match, _ := regexp.MatchString("^[a-zA-Z0-9]*$", data.Name)
	if !match {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	match, _ = regexp.MatchString("^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$", data.Color)
	if !match {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// call service
	userId, err := c.service.PostRegister(r.Context(), data.Name, data.Email, data.Color)
	if err != nil {
		panic(err)
	}

	// create and send response
	resp := PostRegisterResponse{
		UserId: userId,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		panic(err)
	}
}

// GetAllUsers - To check backend color implementations
func (c UserApiController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// call service
	users, err := c.service.GetAllUsers(r.Context())
	if err != nil {
		panic(err)
	}

	// create and send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		panic(err)
	}
}
