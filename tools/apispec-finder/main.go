package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
)

const workDir string = "/tmp/apispec-finder"
const azurermProviderRepo string = "https://github.com/terraform-providers/terraform-provider-azurerm.git"

func cloneAzurermProvidersource() {
	fmt.Println(fmt.Sprintf("Cloning Repo: %s", azurermProviderRepo))

	_, err := git.PlainClone(workDir, false, &git.CloneOptions{
		URL:      azurermProviderRepo,
		Progress: os.Stdout,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func prepWorkDir() {
	if _, err := os.Stat(workDir); !os.IsNotExist(err) {
		err := os.RemoveAll(workDir)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func findSourceFile(resouceName string) (string, error) {
	command := fmt.Sprintf("grep -Ril \"\\\"%s\\\"\" %s --include=*.go --exclude=*_test.go --exclude=*registration.go", resouceName, workDir)
	out, err := exec.Command("grep", "-Ril", "\"\\\"%s\\\"\"", "%s", "--include=*.go", "--exclude=*_test.go", "--exclude=*registration.go").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)

	return command, err
}

func main() {
	//prepWorkDir()
	//cloneAzurermProvidersource()
	findSourceFile("azurerm_subnet")

	fset := token.NewFileSet() // positions are relative to fset

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "/home/hattan/projects/terraform-provider-azurerm/azurerm/internal/services/network/subnet_resource.go", nil, parser.ImportsOnly)
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
