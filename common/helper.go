package common

import (
	"crypto/rand"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TableName(str string) string {
	str1, err := beego.AppConfig.String("dbprefix")
	if err != nil {
		return str
	}
	return fmt.Sprintf("%s%s", str1, str)
}

func GenValidateCode(width int) string {
	if width <= 0 {
		width = 6
	}
	const digits = "0123456789"
	b := make([]byte, width)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			panic(err)
		}
		b[i] = digits[n.Int64()]
	}
	return string(b)
}
