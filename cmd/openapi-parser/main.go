package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/roscopecoltran/aor-gin-swagger/generator/utils"
	"strings"
	"text/template"
)

type Swagger struct {
	SwagVersion string `json:"swagger"`
	ObjectsFlag bool
	ApiTypeFlag bool
	Package     string
	Paths       json.RawMessage
	Definitions json.RawMessage
	Defs        []*Definitionsprops
}

var Swag Swagger

type Definitionsprops struct {
	Name       string
	Usage      string
	Properties map[string]interface{} `json:"properties"`
	Indprop    []Property
}

type Property struct {
	Name                 string
	Type                 string      `json:"type"`
	Format               string      `json:"format"`
	Items                interface{} `json:"items"`
	Enum                 interface{} `json:"enums"`
	Refs                 string      `json:"$refs"`
	AdditionalProperties interface{} `json:additionalProperties`
	Default              bool        `json:"default"`
}

type AdditionalProps struct {
	Type  string
	Refs  string
	Items interface{}
}

func HandleRefs(aInRefVal string, aInDefName string, usage string) string {

	defintion := strings.SplitAfter(aInRefVal, "#/definitions/")
	fmt.Println(defintion[1])
	def, _, _, _ := jsonparser.Get(Swag.Definitions, defintion[1])
	if defintion[1] == aInDefName {
		fmt.Println("This will cause loop")
		return defintion[1]
	}

	ParseDefintions(defintion[1], def, defintion[1], usage)
	return defintion[1]
}

func ParseDefintions(aInDefintionName string, jsonRawDef []byte, aInMetaTargetName string, usage string) {

	var vardef Definitionsprops
	_ = json.Unmarshal(jsonRawDef, &vardef)
	// IF it is first Element, Proceed without checking
	if len(Swag.Defs) > 0 {
		for i := range Swag.Defs {
			if aInDefintionName == Swag.Defs[i].Name {

				fmt.Println(aInDefintionName, "This defintion already added here")
				return
			}
		}
	}
	vardef.Name = aInMetaTargetName
	vardef.Usage = usage
	fmt.Println(vardef)
	//v,_,_,_ = jsonparser.Get(swag.Definitions,defintion[1],"properties")
	fmt.Println("================================")

	for key, val := range vardef.Properties {
		lname := key

		ltype := val.(map[string]interface{})["type"]
		if ltype == nil {
			ltype = ""
		}

		lFormat := val.(map[string]interface{})["format"]
		if lFormat == nil {
			lFormat = ""
		}
		lItems := val.(map[string]interface{})["items"]

		if lItems != nil {
			fmt.Println("lItems:", lItems)
			for keyItem, valItem := range (lItems).(map[string]interface{}) {
				if keyItem == "type" {
					if ltype == "array" {
						ltype = "Collection(" + valItem.(string) + ")"
					} else {
						ltype = valItem.(string)
					}
				}
				if keyItem == "$ref" {
					parsedRefName := HandleRefs(valItem.(string), vardef.Name, vardef.Usage)
					if ltype == "array" {
						ltype = "Collection(" + parsedRefName + ")"
					} else {
						ltype = parsedRefName
					}

				}
			}
		} else {
			lItems = ""
		}
		lEnum := val.(map[string]interface{})["enum"]
		if lEnum == nil {
			lEnum = ""
		}
		lDefault := val.(map[string]interface{})["default"]
		if lDefault == nil {
			lDefault = false
		}
		lRefs := val.(map[string]interface{})["$ref"]
		if lRefs != nil {
			parsedRefName := HandleRefs(lRefs.(string), vardef.Name, vardef.Usage)
			ltype = parsedRefName

		} else {
			lRefs = ""
		}
		lAddProps := val.(map[string]interface{})["additionalProperties"]
		if lAddProps != nil {
			fmt.Println("Add Props:", lAddProps)
			var tmpAddType string
			for keyItem, valItem := range (lAddProps).(map[string]interface{}) {
				if keyItem == "type" {
					if valItem.(string) != "" {
						tmpAddType = valItem.(string)
					} else {
						tmpAddType = ""
					}
				}

				if keyItem == "items" {
					for k, v := range (valItem).(map[string]interface{}) {
						if k == "$ref" {
							parsedRefName := HandleRefs(v.(string), vardef.Name, vardef.Usage)

							if tmpAddType == "array" {
								tmpAddType = "Collection(" + parsedRefName + ")"
							} else {
								tmpAddType = parsedRefName
							}

						}

					}
				}
				if keyItem == "$ref" {
					parsedRefName := HandleRefs(valItem.(string), vardef.Name, vardef.Usage)

					if tmpAddType == "array" {
						tmpAddType = "Collection(" + parsedRefName + ")"
					} else {
						tmpAddType = parsedRefName
					}

				}
			}
			if ltype == "object" {
				ltype = tmpAddType
			}

		} else {
			lAddProps = nil
		}

		fmt.Println("================================")
		fmt.Println("Property Name:", lname)
		fmt.Println("Type:", ltype)
		fmt.Println("Format:", lFormat)
		fmt.Println("Items:", lItems)
		fmt.Println("Enum:", lEnum)
		fmt.Println("Refs:", lRefs)
		fmt.Println("Default:", lDefault)
		fmt.Println("lAddProps:", lAddProps)

		fmt.Println("================================")

		tmpProperty := Property{Name: string(lname), Type: ltype.(string), Format: lFormat.(string), Items: lItems, Enum: lEnum, Refs: lRefs.(string), AdditionalProperties: lAddProps, Default: lDefault.(bool)}

		vardef.Indprop = append(vardef.Indprop, tmpProperty)

	}

	Swag.Defs = append(Swag.Defs, &vardef)

	fmt.Println(Swag.Defs)
}

const (
	dirFilePerm os.FileMode = 0777

	modelDirectory string = "model"
	metaDirectory  string = "meta/"
	projectKey     string = "bitbucket-eng-sjc1.cisco.com/an"
)

func loadTemplate(name string) *template.Template {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Print("Environment variable 'GOPATH' must be set.")
	}

	fname := name
	fmt.Printf("Parsing template '%s'", fname)
	t, err := template.ParseFiles(fname)
	if err != nil {
		fmt.Print("Invalid adgen template : ", err)
	}
	return t
}

func createDir(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err := os.MkdirAll(name, dirFilePerm)
		if err != nil {
			fmt.Printf("cannot make dir '%s' : %s", name, err)
		}
	}
}

func createFile(name string) *os.File {

	f, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, dirFilePerm)
	if err != nil {
		fmt.Printf("Failed to create file '%s' : %s", name, err.Error())
	}

	return f
}

func generateModel(mg Swagger) {
	var t *template.Template
	var f *os.File
	fmt.Print("Generating Model File")
	t = loadTemplate("templates/model.gotmpl")
	createDir("generated")
	f = createFile("generated/model.yaml")
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Print("Failed to close file: ", err)
		}
	}(f)

	err := t.Execute(f, mg)
	if err != nil {
		fmt.Print("Error processing template")
	}

}

func main() {
	// add cli-flags
	configFile := "./config.yaml"
	filePath := "./tests/swagger.json"
	fmt.Printf("// reading file %s\n", filePath)
	file, err1 := ioutil.ReadFile(filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Swagger Specfile doesnt exist:", filePath)
		return
	}
	// parsing Yaml to populate structures
	if err1 != nil {
		fmt.Printf("// error while reading file %s\n", filePath)
		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}
	err2 := json.Unmarshal(file, &Swag)
	fmt.Printf("this is swag value : %s \n", Swag.SwagVersion)
	fmt.Println("================================")
	meta := utils.GenModel{}
	lMetaExist := utils.ParseYaml(configFile, &meta)
	if lMetaExist {

		Swag.Package = meta.Package
		Swag.ObjectsFlag = false
		Swag.ApiTypeFlag = false
		fmt.Println(meta)
		for i := range meta.Objects {
			fmt.Println(string(meta.Objects[i].SourceName))

			v, _, _, _ := jsonparser.Get(Swag.Paths, "/"+meta.Objects[i].SourceName, "get", "responses", "200", "schema", "$ref")
			if string(v) == "" {
				it, _, _, _ := jsonparser.Get(Swag.Paths, "/"+meta.Objects[i].SourceName, "get", "responses", "200", "schema", "type")
				if string(it) == "array" {

					v, _, _, _ = jsonparser.Get(Swag.Paths, "/"+meta.Objects[i].SourceName, "get", "responses", "200", "schema", "items", "$ref")
				}
			}
			fmt.Printf("%s\n", string(v))
			fmt.Println("================================")
			defintion := strings.SplitAfter(string(v), "#/definitions/")
			fmt.Println(defintion[1])
			def, _, _, _ := jsonparser.Get(Swag.Definitions, defintion[1])
			switch meta.Objects[i].Usage {
			case "mo":
				Swag.ObjectsFlag = true
			case "type":
				Swag.ApiTypeFlag = true
			}
			ParseDefintions(defintion[1], def, meta.Objects[i].TargetName, meta.Objects[i].Usage)
		}
		generateModel(Swag)
	}
	fmt.Print(err2)

}
