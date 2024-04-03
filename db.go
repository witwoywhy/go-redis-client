package gedis

import (
	"io"
	"strconv"
)

func selectDB(db int, readWriter io.ReadWriter) error {
	ch := make(chan *response)
	go read(ch, readWriter)

	w := &writer{lists: make([]string, 0)}
	w.cmd(SELECT)
	w.addString(strconv.Itoa(db))

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
