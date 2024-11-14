package bcrypt_utils

import "golang.org/x/crypto/bcrypt"

func Compare(input_string string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input_string))
	return err == nil
}
