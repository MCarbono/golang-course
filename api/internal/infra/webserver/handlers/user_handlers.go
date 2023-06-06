package handlers

import (
	"api/internal/dto"
	"api/internal/entity"
	"api/internal/infra/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.UserInterface
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, JwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
		JwtExpiresIn: JwtExpiresIn,
	}
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTINPUT  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	var jwtInput dto.GetJWTINPUT
	err := json.NewDecoder(r.Body).Decode(&jwtInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.UserDB.FindByEmail(jwtInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		messageErr := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(messageErr)
		return
	}
	if !user.ValidatePassword(jwtInput.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		messageErr := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(messageErr)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      400         {object}  Error
// @Failure      500         {object}  Error
// @Router       /users [post]
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := entity.NewUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		messageErr := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(messageErr)
	}
	err = u.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		messageErr := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(messageErr)
	}
	w.WriteHeader(http.StatusCreated)
}
