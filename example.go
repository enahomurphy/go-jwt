package main

import (
	"fmt"
	"go-jwt/jwt"
)

func main() {
	fmt.Println(jwt.CompareHmac("eyJhbGciOkiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0IiwiZXhwIjoxNTAxMTU0NDI0LCJwdWJsaWMiOnsiTmFtZSI6Im11cnBoeSJ9fQ==", "X/Rn3YdCfgl1obKmh2oY/V5XNOhPEhsvQ42QnZOvDTc=", "test"))
}
