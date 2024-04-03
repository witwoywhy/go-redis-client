package gedis

import (
	"io"
)

func authen(password string, readWriter io.ReadWriter) error {
	ch := make(chan *response)
	go read(ch, readWriter)

	w := &writer{lists: make([]string, 0)}
	w.cmd(AUTH)
	w.addString(password)

	_, err := readWriter.Write(w.toBytes())
	if err != nil {
		return err
	}

	err = isSimpleError(<-ch)
	if err != nil {
		return err
	}

	return nil
}
