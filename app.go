package main

import (
	"bufio"
	"os"
	"strings"
)

// crackPassword 尝试使用提供的密码与字典文件中的密码进行匹配
func crackPassword(providedPassword string) bool {
	// 打开字典文件
	file, err := os.Open("dictionary.txt")
	if err != nil {
		panic(err) // 实际使用时应更优雅地处理错误
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// 逐行读取字典文件
	for scanner.Scan() {
		// 获取字典中的密码（假设每行一个密码）
		dictionaryPassword := strings.TrimSpace(scanner.Text())
		// 如果提供的密码与字典中的密码匹配，返回true
		if dictionaryPassword == providedPassword {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err) // 实际使用时应更优雅地处理错误
	}

	// 如果没有找到匹配的密码，返回false
	return false
}

var correctPattern = []string{"1,1", "2,2", "3,3", "4,4"}

func checkPattern(input []string) bool {
	if len(input) != len(correctPattern) {
		return false
	}

	for i, point := range input {
		if point != correctPattern[i] {
			return false
		}
	}
	return true
}
