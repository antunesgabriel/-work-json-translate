package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestMakeEmptyJson(t *testing.T) {
	t.Run("clone json with empty field values", func(t *testing.T) {
		pathSource := "./test.json"
		newJsonFileName := "./new_test.json"
		outputDir := "."

		source := []byte(`
			"foo": {
				"bass": "yeah",
				"nes": {
					"ted": "yes"
				}
			}
		`)

		expected := `{"foo": {"bass": "","nes": {"ted": ""}}`

		f, err := os.Create("./test.json")

		if err != nil {
			t.Errorf("failed on create test json file")
			return
		}

		if _, err := f.Write(source); err != nil {
			t.Errorf("failed on create test json file")
			return
		}

		if MakeEmptyJson(pathSource, newJsonFileName, outputDir) != nil {
			t.Error("failed")
		}

		defer os.RemoveAll(pathSource)
		defer os.RemoveAll(newJsonFileName)

		resultFile, err := ioutil.ReadFile(newJsonFileName)

		if err != nil {
			t.Errorf("failed on reader output test json file")
			return
		}

		if r, err := jsonBytesEqual([]byte(expected), resultFile); err != nil || !r {
			t.Errorf("failed")
			return
		}
	})
}

func jsonBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}

	fmt.Println(j, j2)

	return reflect.DeepEqual(j2, j), nil
}
