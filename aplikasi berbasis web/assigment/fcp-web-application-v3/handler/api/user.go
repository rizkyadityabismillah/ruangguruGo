package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	var user model.User
    
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid decode json"})
        return
    }

    if user.Email == "" || user.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid decode json"})
        return
    }

    // Panggil userService untuk login user
    userID, err := u.userService.Login(&model.User{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
        Subject:   strconv.Itoa(user.ID),
    })

    tokenString, err := token.SignedString([]byte(model.JwtKey))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
        return
    }

    // Set JWT token as cookie
    c.SetCookie("session_token", tokenString, 3600, "/", "", false, true)

    c.JSON(http.StatusOK, gin.H{
        "user_id": userID,
        "message": "login success",
    })
}

// GetUserTaskCategory adalah handler untuk endpoint yang mengambil daftar tugas pengguna dengan kategori terkait
func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	_, err := c.Cookie("session_token")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tasks, err := u.userService.GetUserTaskCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error internal server"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
