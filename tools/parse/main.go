package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "/home/hattan/projects/terraform-provider-azurerm/azurerm/internal/services/analysisservices/analysis_services_server_resource.go", nil, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		importName := s.Path.Value
		if strings.Contains(importName, "github.com/Azure/azure-sdk-for-go") {
			fmt.Println(s.Path.Value)
		}

	}

}
