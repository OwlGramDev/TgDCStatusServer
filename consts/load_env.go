package consts

import (
	"github.com/joho/godotenv"
	"os"
)

var SecretDCKey string
var IPInfoKey string

func LoadEnv() {
	err := godotenv.Load(EnvFilePath)
	if err != nil {
		panic("Error loading .env file")
	}
	IPInfoKey = os.Getenv("IP_INFO_KEY")
	SecretDCKey = os.Getenv("SECRET_DC_KEY")
}
