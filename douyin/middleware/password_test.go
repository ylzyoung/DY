package middleware

import "testing"

func TestSHA1(t *testing.T) {
	println(SHA1("haha123456"))
	print(SHA1("haha23456"))
}
