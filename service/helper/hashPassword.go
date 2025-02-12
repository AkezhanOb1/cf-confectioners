package helper

import "golang.org/x/crypto/bcrypt"

//HashPassword is a
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}


func comparePassword(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}