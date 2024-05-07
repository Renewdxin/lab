package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// encrypt函数使用 Caesar 密码对输入的字符串进行加密。
// 参数 data 是需要加密的字符串，shift 是字符移位的数量。
// 返回值是加密后的字符串。
func encrypt(data string, shift int) string {
	encrypted := ""
	for _, c := range data {
		if c >= 'a' && c <= 'z' {
			// 对小写字母进行加密
			encrypted += string((c-'a'+int32(shift))%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			// 对大写字母进行加密
			encrypted += string((c-'A'+int32(shift))%26 + 'A')
		} else {
			// 非字母字符不加密，直接添加到结果中
			encrypted += string(c)
		}
	}
	return encrypted
}

func main() {
	reader := os.Stdin
	fmt.Print("请输入加密数据: ")
	var input string
	fmt.Fscanln(reader, &input)

	// 使用 encrypt 函数加密输入的数据，假设移位量为3
	encryptedData := encrypt(input, 3)

	// 创建一个 JSON 数据包，包含加密后的数据
	jsonData := map[string]string{
		"data": encryptedData,
	}
	jsonValue, _ := json.Marshal(jsonData)

	// 发送 POST 请求，将加密数据发送到指定的 URL
	response, err := http.Post("http://localhost:8080/crack/login", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer response.Body.Close()

	// 打印信息，确认数据发送成功
	fmt.Println("成功发送数据")
}
