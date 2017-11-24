package utils

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/ghodss/yaml"
	"github.com/hokaccha/go-prettyjson"
)

func ConvertSwaggerFile(input_filepath string) (error, []byte, string) {
	var is_writable bool
	var output_filepath, output_file_extension string
	var result []byte
	input_file_extension := path.Ext(input_filepath)
	fmt.Printf("input_filepath: %s\n", input_filepath)
	fmt.Printf("input_file_extension: %s\n", input_file_extension)
	data, err := ioutil.ReadFile(input_filepath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err, result, ""
	}
	if IsYAML(string(data)) {
		var result_raw []byte
		result_raw, err = yaml.YAMLToJSON(data)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return err, result, ""
		}
		result, _ := prettyjson.Format(result_raw)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return err, result, ""
		}
		output_file_extension = ".json"
		is_writable = true

	}
	if IsJSON(string(data)) {
		result, err = yaml.JSONToYAML(data)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return err, result, ""
		}
		output_file_extension = ".yaml"
		is_writable = true
	}

	fmt.Printf("is_writable: %t\n", is_writable)
	fmt.Printf("output_file_extension: %s\n", output_file_extension)
	if is_writable {
		output_filepath = input_filepath[0:len(input_filepath)-len(input_file_extension)] + output_file_extension
		fmt.Printf("output_filepath: %s\n", output_filepath)
		ioutil.WriteFile(output_filepath, result, 0644)
	}

	fmt.Printf("output: %v\n", string(result))
	return nil, result, output_filepath
}
