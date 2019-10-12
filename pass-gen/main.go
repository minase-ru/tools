package main

import (
	"golang.org/x/crypto/bcrypt"
	"crypto/md5"
	"fmt"
)

func cmd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}


func generatePassword(p string) (string, error) {
	s, err := bcrypt.GenerateFromPassword([]byte(cmd5(p)), bcrypt.DefaultCost)
	return string(s), err
}

func main()  {
	s, _ := generatePassword("123123")
	fmt.Print(s)
}