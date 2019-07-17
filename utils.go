package grobot

import (
	"errors"
	"strconv"
	"strings"
)

func stringBuilder(strs ...string) string {
	stringBuilder := strings.Builder{}

	for _, s := range strs {
		stringBuilder.WriteString(s)
	}

	return stringBuilder.String()
}

func newError(errType string, result interface{}) error {
	var rs string

	switch result.(type) {
	case string:
		rs = stringBuilder(errType, ":", result.(string))
	case int:
		rs = stringBuilder(errType, ":", strconv.Itoa(result.(int)))
	case error:
		rs = stringBuilder(errType, ":", result.(error).Error())
	default:
		rs = errType
	}

	return errors.New(rs)
}
