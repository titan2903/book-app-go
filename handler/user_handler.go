package handler

import (
	"book-app/auth"
	"book-app/entity"
	"book-app/formatter"
	"book-app/service"
	"book-app/transport"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.ServiceUser
	authService auth.AuthService
}

func NewUserHandler(userService service.ServiceUser, authService auth.AuthService) *userHandler {
	return &userHandler{userService, authService}
}

func(h *userHandler) RegisterUser(c *gin.Context) {
	var input transport.RegisterUserInput
	
	err := c.ShouldBindJSON(&input) //! validasi di lakukan di sini

	if err != nil {
		errors := transport.FormatValidationError(err)
		fmt.Println("masuk1")
		errorMessage := gin.H{"errors": errors}
		response := transport.ApiResponse("Register account failed", http.StatusBadRequest, "error", errorMessage) //! entity tidak bisa di proses 

		c.JSON(http.StatusBadRequest, response)
		return
	}

	isUsernameAvailable, err := h.userService.IsUserNameAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := transport.ApiResponse("Check Username failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isUsernameAvailable,
	}

	if !isUsernameAvailable {
		response := transport.ApiResponse("Username has been registered", http.StatusConflict, "error", data)
		c.JSON(http.StatusConflict, response)
		return
	}

	newUser, _ := h.userService.RegisterUser(input)
	if err != nil {
		fmt.Println("masuk2")
		response := transport.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil) //! Invalid input
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		fmt.Println("masuk3")
		response := transport.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.FormatUser(newUser, token)

	response := transport.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func(h *userHandler) LoginUser(c *gin.Context) {
	var input transport.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := transport.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := transport.ApiResponse("Login account failed", http.StatusUnprocessableEntity, "error", errorMessage) //! entity tidak bisa di proses 
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := transport.ApiResponse("Login account failed", http.StatusUnprocessableEntity, "error", errorMessage) //! entity tidak bisa di proses, karena username salah, id dan username tidak di temukan dan password salah
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := transport.ApiResponse("Login account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.FormatUser(loggedinUser, token)

	response := transport.ApiResponse("Login has been Successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// func(h *userHandler) CheckUsernameHasBeenRegister(c *gin.Context) {
// 	var input transport.CheckUsernameInput

// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		errors := transport.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}

// 		response := transport.ApiResponse("Check Username failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	isUsernameAvailable, err := h.userService.IsUserNameAvailable(input)
// 	if err != nil {
// 		errorMessage := gin.H{"errors": "Server Error"}
// 		response := transport.ApiResponse("Check Username failed", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	data := gin.H{
// 		"is_available": isUsernameAvailable,
// 	}

// 	metaMessage := "Username has been registered"

// 	if isUsernameAvailable {
// 		metaMessage = "Username is available"
// 	}

// 	response := transport.ApiResponse(metaMessage, http.StatusOK, "success", data)
// 	c.JSON(http.StatusOK, response)
// }

func(h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(entity.User)
	formatter := formatter.FormatFetchUser(currentUser)

	response := transport.ApiResponse("Fetch user data successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func(h *userHandler) UpdateUser(c *gin.Context) {
	var input transport.FormUpdateUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := transport.FormatValidationError(err)
		response := transport.ApiResponse("Failed  Update User", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(entity.User)
	input.User = currentUser

	updateUser, err := h.userService.UpdateUser(input)
	if err != nil {
		response := transport.ApiResponse("Failed Create Book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	response := transport.ApiResponse("Success Create Book", http.StatusOK, "success", formatter.FormatFetchUser(updateUser))

	c.JSON(http.StatusOK, response)
}