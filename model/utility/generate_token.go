package utility

import (
	"crypto/md5"
	"encoding/hex"
	//config "github.com/pardev/cantik-mart/config"
	"time"
)

func GenerateToken(email string) string {

	newstring := email + GetTime()
	return GetMD5Hash(newstring)
}

func GetTime() string {

	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func GetMD5Hash(text string) string {

	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
