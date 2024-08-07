package middlewares

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"strconv"
	"wallet-api/internal/lib/types"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": "invalid body"},
			)
		}
		// снова записываем тело запроса для дальнейшей обработки, 
		// так как ранее мы его исчерпали чтением
		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		
		// зашифровываем тело запроса с помощью алгоритма шифрования SHA1 и секретного ключа
		hash := hmac.New(sha1.New, []byte(os.Getenv("SECRET_KEY")))
		hash.Write(body)
		expectedMac := hex.EncodeToString(hash.Sum(nil))
		// получаем из заголовка запроса ожидаемую хэш-сумму
		actualMac := c.Request.Header.Get("X-Digest")

		// сравниваем обе хэш-суммы - они должны совпадать, тогда
		// пользователь считается аутентифицированным
		if actualMac != expectedMac {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// достаем из заголовка так же идентификатор пользователя и кладем его в контекст
		userId, err := strconv.Atoi(c.Request.Header.Get("X-UserId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": types.ErrInvalidUserId.Error(),
			})
		}
		c.Set(types.KeyUserId, userId)
	}
}
