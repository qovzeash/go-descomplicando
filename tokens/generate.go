package tokens

import (
	"fmt"
	"qovzeash/go-descomplicando/utils"

	"github.com/cristalhq/jwt/v5"
)

func Generate() {
	key := []byte(`secret`)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	utils.CheckErr(err)

	claims := &jwt.RegisteredClaims{
		Audience: []string{"admin"},
		ID:       "random-unique-string",
	}

	builder := jwt.NewBuilder(signer)

	token, err := builder.Build(claims)
	utils.CheckErr(err)

	var newToken string = token.String()

	fmt.Println(newToken)

}
