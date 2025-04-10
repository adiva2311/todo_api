package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		panic(err)
	}
	return string(hashed), nil
}

func CheckPasswordHash(password string, hash_password string) bool {
	// err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password))
	// if err != nil {
	// 	return false
	// }
	// return true
	return bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password)) == nil
}