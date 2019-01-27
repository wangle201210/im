package helper

import (
	"math/rand"
	"strconv"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


func Interface2int64(inter interface{}) (i int64) {
	switch inter.(type) {

	case string:
		i, _ = strconv.ParseInt(inter.(string), 10, 64)
		break
	case int64:
		i = inter.(int64)
		break
	case float64:
		i = int64(inter.(float64))
		break
	}
	return
}

func Interface2string(inter interface{}) (s string) {
	switch inter.(type) {

	case string:
		s = inter.(string)
		break
	case int64:
		s =strconv.FormatInt(inter.(int64),10)
		break
	case float64:
		s = strconv.FormatInt(int64(inter.(float64)),10)
		break
	}
	return
}

