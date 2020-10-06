package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// HasPrefix 判断字符串s是否有prefix前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// Contains 判断字符串s是否有substr子串
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// Comma 把字符串s每三位用逗号分隔
func Comma(s string) string {
	length := len(s)
	if length <= 3 {
		return s
	}
	return Comma(s[:length-3]) + "," + s[length-3:]
}

// Ints2string 把int[]转化成字符串
func Ints2string(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// Sum 对任意数量的int求和
func Sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

// CountDown 秒倒计时
func CountDown(delay int) {
	tick := time.Tick(1 * time.Second)
	for i := delay; i >= 0; i-- {
		fmt.Println(i)
		<-tick
	}
}

// ReadFile 读文件
func ReadFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// HashPassword 把明文密码转化成hash值
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash 检查密码hash值
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
