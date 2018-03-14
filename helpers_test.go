package dialogflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-test/deep"
)

func JSONEqual(a, b []byte) error {
	var j, j2 interface{}

	d := json.NewDecoder(bytes.NewReader(a))
	if err := d.Decode(&j); err != nil {
		return err
	}

	d = json.NewDecoder(bytes.NewReader(b))
	if err := d.Decode(&j2); err != nil {
		return err
	}

	if diff := deep.Equal(j, j2); diff != nil {
		return fmt.Errorf("%s", diff)
	}
	return nil
}

func JSONStringsEqual(a, b string) error {
	var j, j2 interface{}

	d := json.NewDecoder(strings.NewReader(a))
	if err := d.Decode(&j); err != nil {
		return err
	}

	d = json.NewDecoder(strings.NewReader(b))
	if err := d.Decode(&j2); err != nil {
		return err
	}

	if diff := deep.Equal(j, j2); diff != nil {
		return fmt.Errorf("%s", diff)
	}
	return nil
}

func PayloadTester(req interface{}, expected []byte) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	return JSONEqual(payload, expected)
}
