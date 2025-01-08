package middleware

import "os"

var JwtSecretKey = []byte(os.Getenv("SECRETKEY"))

