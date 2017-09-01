package main

import (
	"fmt"
	"go-jwt/jwt"
	"go-jwt/test"
)

func main() {
	// type Data struct {
	// 	Name string
	// }
	// data := Data{Name: "murphy"}

	// payload := jwt.Payload{
	// 	Sub:    "test",
	// 	Exp:    time.Now().Unix(),
	// 	Public: data,
	// }
	// token := jwt.Encode(payload, "test")
	fmt.Println(test.Avg(1, 2, 3, 4))
	fmt.Println(jwt.CompareHmac("eyJhbGciOkiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0IiwiZXhwIjoxNTAxMTU0NDI0LCJwdWJsaWMiOnsiTmFtZSI6Im11cnBoeSJ9fQ==", "X/Rn3YdCfgl1obKmh2oY/V5XNOhPEhsvQ42QnZOvDTc=", "test"))
}
