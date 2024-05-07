package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// decrypt函数用于解密经过移位加密的字符串
// 参数data为待解密的字符串，shift为移位量（假设为正数）
// 返回值为解密后的字符串
func decrypt(data string, shift int) string {
	decrypted := ""
	for _, c := range data {
		// 对字母进行解密操作，其他字符保持不变
		if c >= 'a' && c <= 'z' {
			decrypted += string((c-'a'+26-int32(shift))%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			decrypted += string((c-'A'+26-int32(shift))%26 + 'A')
		} else {
			decrypted += string(c)
		}
	}
	return decrypted
}

func main() {
	router := gin.Default()
	// 定义了一个POST请求的处理函数，请求路径为/crack/login。
	//该处理函数首先解析客户端发送的JSON数据，提取出data字段，然后调用decrypt函数对该字段进行解密（假设移位量为3），并将解密结果返回给客户端
	//返回的数据包括原始加密字符串和解密后的字符串。
	router.POST("/crack/login", func(c *gin.Context) {
		var json struct {
			Data string `json:"data"`
		}
		if c.BindJSON(&json) == nil {
			original := decrypt(json.Data, 3) // 假设移位量为3
			c.JSON(http.StatusOK, gin.H{
				"encrypted": json.Data,
				"decrypted": original,
			})
		}
	})
	router.Run(":8080")
}
