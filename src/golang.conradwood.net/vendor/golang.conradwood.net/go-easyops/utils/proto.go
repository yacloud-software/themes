package utils

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
)

// take a proto and convert it into a base64 string
func Marshal(req proto.Message) (string, error) {
	data, err := proto.Marshal(req)
	if err != nil {
		return "", err
	}
	b64 := base64.StdEncoding.EncodeToString(data)
	return b64, nil
}

// take a base64 string and convert it into the proto
func Unmarshal(b64string string, req proto.Message) error {
	pdata, err := base64.StdEncoding.DecodeString(b64string)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(pdata, req)
	if err != nil {
		return err
	}

	return nil
}
