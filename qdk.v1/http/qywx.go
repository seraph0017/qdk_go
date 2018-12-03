package http

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/url"
)

func ReceiveMessage(c *gin.Context) {

	u, _ := url.ParseQuery(c.Request.URL.String())
	fmt.Println(u)

	c.JSON(200, gin.H{
		"message": "test",
	})
}

func msg_signature(token, timestamp, nonce, msg_encrypt string) {
	h := sha1.New()
	for _, s := range []string{token, timestamp, nonce, msg_encrypt} {
		io.WriteString(h, s)
	}
	return
}
