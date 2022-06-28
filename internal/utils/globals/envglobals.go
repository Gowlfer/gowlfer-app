package globals

import "os"

var SecretKey = os.Getenv("SECRET_KEY")
