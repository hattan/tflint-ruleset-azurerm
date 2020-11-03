#!/usr/bin/env bash

declare version=""
declare resourceName=""
declare DEBUG_FLAG=false

while [[ "$#" -gt 0 ]]
do
  case $1 in
    -r | --resource)
      resourceName=$2
      ;;

    -d | --debug )             
      DEBUG_FLAG=true
      ;; 
  esac
  shift
done

providerVersion=$(cat provider.tf | grep version)
providerVersion="v"$(echo "$providerVersion" | cut -d'"' -f 2)
echo "Found Provider version: $providerVersion"

if [ ! -d ./terraform-provider-azurerm ]; then
  git clone git@github.com:terraform-providers/terraform-provider-azurerm.git
fi

pushd terraform-provider-azurerm
  git fetch && git fetch --tags
  git checkout $version

  resourceFile=$(grep -Ril "${resourceName}" . --include=*.go --exclude=*_test.go --exclude=*registration.go)
  resourceFolder="${resourceFile%/*}/"
  go list -json -f "{{.ImportPath}} {{.Imports}}" github.com/terraform-providers/terraform-provider-azurerm/$resourceFolder | jq -r '.Imports[]  | select(. | contains("azure-sdk-for-go"))'
 popd