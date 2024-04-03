package gedis

import (
	"bytes"
	"errors"
	"strings"
)

func getMessageError(response string) string {
	return strings.Join(strings.Split(response, " ")[1:], " ")
}

func isSimpleError(response *response) error {
	if response.err != nil {
		return response.err
	}

	if bytes.HasPrefix(response.payload, []byte{45}) {
		return errors.New(getMessageError(string(response.payload)))
	}

	return nil
}
