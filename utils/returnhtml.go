package utils

import (
	"os"
)

type HTML struct {
	Path string
}

func (h *HTML) ReturnHTML() (string, error) {
	file, err := os.Open(h.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return "", err
	}
	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		return "", err
	}
	returnedData := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n" + string(data)
	return returnedData, nil
}
