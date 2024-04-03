package gedis

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

type response struct {
	n       int
	err     error
	payload []byte
}

func read(ch chan (*response), reader io.Reader) {
	b := make([]byte, 1024)
	for {
		n, err := reader.Read(b)
		ch <- &response{
			n:       n,
			err:     err,
			payload: b,
		}
		return
	}
}

func makeResponsChAndRead(pool *pool) chan (*response) {
	ch := make(chan *response)
	go read(ch, pool.conn)
	return ch
}

func readStrings(resp *response) ([]string, error) {
	var buffer bytes.Buffer
	var ls []string

	_, err := buffer.Write(resp.payload[:resp.n])
	if err != nil {
		return nil, err
	}

	scan := bufio.NewScanner(&buffer)
	for scan.Scan() {
		s := scan.Text()
		if strings.HasPrefix(s, string(BulkStrings)) || strings.HasPrefix(s, string(Arrays)) {
			continue
		}

		ls = append(ls, s)
	}

	return ls, nil
}

func readIntegers(resp *response) (int, error) {
	var buffer bytes.Buffer

	_, err := buffer.Write(resp.payload[:resp.n])
	if err != nil {
		return 0, err
	}

	scan := bufio.NewScanner(&buffer)
	for scan.Scan() {
		s := scan.Text()
		if strings.HasPrefix(s, string(BulkStrings)) || strings.HasPrefix(s, string(Arrays)) {
			continue
		}

		if strings.HasPrefix(s, string(Integers)) {
			s = strings.Replace(s, ":", "", -1)
			return strconv.Atoi(s)
		}
	}

	return 0, nil
}
