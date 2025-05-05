package middlewares

import (
	"backend/internal/logger"
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogRequestBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		// JSON などの POST/PUT/PATCH リクエストだけログ出力対象
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			// Body 読み取り
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Log.Errorf("Failed to read request body: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "failed to read request"})
				return
			}

			// ログ出力
			logger.Log.Infof("Request Body: %s", string(bodyBytes))

			// 読み取り後、再セット（バインディングのため）
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// 次のハンドラーへ
		c.Next()
	}
}
