package gedis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"
)

type writer struct {
	lists []string
}

func (w *writer) cmd(cmd string) {
	w.lists = append(w.lists, cmd)
}

func (w *writer) add(v any) error {
	s, err := toString(v)
	if err != nil {
		return err
	}

	w.lists = append(w.lists, s)
	return nil
}

func (w *writer) addString(s string) error {
	w.lists = append(w.lists, s)
	return nil
}

func (w *writer) toBytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString("*")
	buffer.WriteString(strconv.Itoa(len(w.lists)))
	for _, list := range w.lists {
		buffer.WriteString("\r\n")
		buffer.WriteString("$")
		buffer.WriteString(strconv.Itoa(utf8.RuneCountInString(list)))
		buffer.WriteString("\r\n")
		buffer.WriteString(list)
	}
	buffer.WriteString("\r\n")
	return buffer.Bytes()
}

func toString(v any) (string, error) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		return v.(string), nil
	case reflect.Struct:
		b, err := json.Marshal(v)
		if err != nil {
			return "", err
		}

		return string(b), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}
