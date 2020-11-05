package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

const azureGoSdkRoot = "github.com/Azure/azure-sdk-for-go"

var isJSONOutput bool

type azureRmResource struct {
	FilePath     string `json:"filePath"`
	ImportModule string `json:"importModule"`
}

func main() {
	var resource azureRmResource
	parseFlags(&resource)

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, resource.FilePath, nil, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}

	err = findAzureSdkImport(file, &resource)
	if err != nil {
		panic(err)
	}

	if !isJSONOutput {
		fmt.Println(resource.ImportModule)
	} else {
		encoded, err := json.Marshal(resource)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(encoded))
	}
}

func parseFlags(resource *azureRmResource) {
	flag.StringVar(&resource.FilePath, "r", "", "path to go file")
	flag.BoolVar(&isJSONOutput, "json", false, "path to go file")
	flag.Parse()
}

func findAzureSdkImport(file *ast.File, resource *azureRmResource) error {
	for _, s := range file.Imports {
		importName := s.Path.Value
		if strings.Contains(importName, azureGoSdkRoot) {
			resource.ImportModule = s.Path.Value
			resource.ImportModule = strings.Replace(resource.ImportModule, "\"", "", -1)
			return nil
		}
	}
	return fmt.Errorf("No Import found with path %s", azureGoSdkRoot)
}
