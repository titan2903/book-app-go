package middleware

import (
	"book-app/auth"
	"book-app/service"
	"book-app/transport"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService auth.AuthService, userService service.ServiceUser) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := transport.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //! menghentikan ke proses selanjutnya jika error
			return
		}

		//! Bearer tokentokentoken
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ") //! memisahkan token dan kata bearer
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := transport.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //! menghentikan ke proses selanjutnya jika error
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims) //! mengambil data token dalam claims / payload

		if !ok || !token.Valid {
			response := transport.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //! menghentikan ke proses selanjutnya jika error
			return
		}

		userID := claim["user_id"].(float64) //! mengubah dari string ke integer

		user, err := userService.GetUserByID(int(userID))
		if err != nil {
			response := transport.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) //! menghentikan ke proses selanjutnya jika error
			return
		}

		c.Set("currentUser", user)
	}
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}