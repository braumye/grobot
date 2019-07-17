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

func newError(errType string, items ...interface{}) error {
	errType = stringBuilder(errType, ":")

	for _, result := range items {
		switch result.(type) {
		case string:
			errType = stringBuilder(errType, result.(string))
		case int:
			errType = stringBuilder(errType, strconv.Itoa(result.(int)))
		case error:
			errType = stringBuilder(errType, result.(error).Error())
		}
	}

	return errors.New(errType)
}
