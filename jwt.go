package charm

import (
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

const jwtPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEAvqBYpgl0hjxDgtaMLs+V
rXcOMLCgg7CMbjuuIyfQaL+KysyPqf0/O8xIwMo7R11DjRVWFhKYUFmSf7e7/S5B
9OzzGeTHwxk4nKEhbRRj94Lp0EuZZy6CpZYr5ScwphfsSO8gCWnQftmlOwG21ynM
EnnaEWGxl4cXd+oIagdMsP6PJEmPAocc4R5Y4jf37TVa0/VmgETfCwv1FPWPxu/k
tOWw3YWrGL9GrxKq4AudpiEp7S9o6Ln76Cq23mkZWOV3cKwenYzZLMWHQR2IbLSu
UOQgkcCuqHXbA9kjqyi47/faokeK93dBknUFOb12cEiExqRfywfxHbPg/IYDzrvo
TcLLfLPB1CEXXObNjjbDhdGf5Dr6mAFLuT8Is29Nqnn6kldmj+dUinOszIjpP9+B
UCQWDF1yPZY/K4XDj0at5gSnkvBn2NI7IP6Ps5aXaP8zuCjA9Lhj8JWlaGTKsZB+
4doKSp/wMaMXyj34fMI26pmPdepmQqBXeGD9r94glOCVAgMBAAE=
-----END PUBLIC KEY-----`

// Auth is the authenticated users's charm id and jwt returned from the ssh server
type Auth struct {
	CharmID string `json:"charm_id"`
	JWT     string `json:"jwt"`
	claims  *jwt.StandardClaims
}

func (cc *Client) setJWTKey() error {
	var bk []byte
	var err error
	if cc.config.JWTKey != "" {
		bk, err = ioutil.ReadFile(cc.config.JWTKey)
		if err != nil {
			return err
		}
	} else {
		bk = []byte(jwtPublicKey)
	}
	pk, err := jwt.ParseRSAPublicKeyFromPEM(bk)
	if err != nil {
		return err
	}
	cc.jwtPublicKey = pk
	return nil
}
