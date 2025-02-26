package middleware

import (
	models2 "douyin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJWTMiddleWare(t *testing.T) {
	router := gin.Default()
	router.Use(JWTMiddleWare())
	router.GET("/test", func(c *gin.Context) {
		userId, exists := c.Get("user_id")
		if exists {
			c.JSON(http.StatusOK, gin.H{"user_id": userId})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "user_id not set"})
		}
	})

	// 测试token有效的情况
	user := models2.UserLogin{
		Id:         1,
		UserInfoId: 123,
		Username:   "douyin",
		Password:   "douyin",
	}
	tokenString, _ := ReleaseToken(user)
	req, _ := http.NewRequest("GET", "/test?token="+tokenString, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"user_id\":123")

}
