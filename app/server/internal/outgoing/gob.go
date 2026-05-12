package outgoing

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func EncodeGob[T any](data T) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return []byte{}, fmt.Errorf("EncodingGob error: %w", err)
	}

	return buffer.Bytes(), nil
}

func DecodeGob[T any](data []byte) (T, error) {
	var buffer = bytes.NewBuffer(data)
	var t T
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&t)
	return t, err
}
