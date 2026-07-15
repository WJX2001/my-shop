package common

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
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
