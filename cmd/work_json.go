package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
)

func MakeEmptyJson(sourceFilePath, outputFileName, outputFolder string) error {
	fmt.Println("-> Read JSON FILE")

	jsonFile, err := ioutil.ReadFile(sourceFilePath)

	if err != nil {
		return err
	}

	var sourceData map[string]interface{}

	fmt.Println("-> Parsing JSON FILE")

	json.Unmarshal(jsonFile, &sourceData)

	fmt.Println("-> Map JSON FILE")

	result := clearJSONField(sourceData)

	jsonData, err := json.Marshal(result)

	if err != nil {
		return err
	}

	fmt.Println("-> Create new JSON FILE")

	if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(path.Join(outputFolder, outputFileName))

	if err != nil {
		return err
	}

	defer file.Close()

	file.Write(jsonData)

	return nil
}

func getNestedValue(data interface{}) interface{} {
	toMap := data.(map[string]interface{})

	return clearJSONField(toMap)
}

func clearJSONField(data map[string]interface{}) map[string]interface{} {
	var d = make(map[string]interface{})

	for key, value := range data {
		fmt.Println(reflect.TypeOf(key), reflect.TypeOf(value))

		switch value.(type) {
		case map[string]interface{}:
			d[key] = getNestedValue(data[key])
		default:
			d[key] = ""
		}
	}

	return d
}
