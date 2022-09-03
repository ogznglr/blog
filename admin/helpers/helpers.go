package helpers

import (
	"errors"
	"fmt"
	"strconv"
)

func ToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("Can't convert")
	}
	return num, nil
}

func DeleteFirstSlender(str string) string {
	var str2 string
	var counter bool
	for _, value := range str {
		if value == '/' && counter == false {
			counter = true
			continue
		}
		str2 = fmt.Sprintf(str2+"%c", value)
	}
	return str2
}
