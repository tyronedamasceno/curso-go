package security

import "golang.org/x/crypto/bcrypt"

func HashPwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

func CheckPwd(hashPwd, strPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(strPwd))
}
