package controllers

import (
	"KayaKuy/config"
	"KayaKuy/models"
	"KayaKuy/services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

func (b *userHandler) RedirectHandler(c *gin.Context) {
	provider := c.Param("provider")

	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("AUTH Failed load file environment")
	} else {
		fmt.Println("AUTH success read file environment")
	}

	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     os.Getenv("CLIENT_ID_GITHUB"),
			"clientSecret": os.Getenv("CLIENT_SECRET_GITHUB"),
			"redirectURL":  "https://kayakuy-production.up.railway.app/api/v1/auth/github/callback",
		},
	}

	providerScopes := map[string][]string{
		"github": []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := config.Gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// Handle callback of provider
func (b *userHandler) CallbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	// Handle callback and check for errors
	user, _, err := config.Gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	var newUser = b.userService.GetOrRegisterUser(provider, user)
	var jwtToken = createToken(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"data":    newUser,
		"token":   jwtToken,
		"message": "Login Success",
	})
}

func createToken(user *models.User) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().AddDate(0, 0, 7).Unix(),
		"iat":       time.Now().Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func (b *userHandler) Register(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		panic(err)
	}

	err = b.userService.Register(user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to register user",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Register",
	})
}

func (b *userHandler) Login(c *gin.Context) {
	var inputUser models.User
	var user models.User

	err := c.ShouldBindJSON(&inputUser)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result":  "Failed to login",
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	err = b.userService.Login(inputUser, &user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": err.Error(),
		})
		c.Abort()
	} else {
		jwtToken := createToken(&user)
		c.JSON(http.StatusOK, gin.H{
			"data":    user,
			"token":   jwtToken,
			"message": "Login Success",
		})
	}
}

func (a *userHandler) UpdateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	user.ID = int64(c.MustGet("jwt_user_id").(float64))

	ct, err := a.userService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result": "Failed to Update Account",
		})

		c.Abort()
		return
	}

	if ct.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "success update Account",
		})
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"result": "Failed to Update Account",
		})

		c.Abort()
		return
	}

}
