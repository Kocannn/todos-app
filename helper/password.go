package helper

import "golang.org/x/crypto/bcrypt"

func HassPassword(password string) (string, error) {
	pw := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) (bool,error){
  err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
  if err != nil {
    return false ,err
  }
  return true ,nil
}