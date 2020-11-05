#!/usr/bin/env bash

declare providerVersion=""
declare resourceName=""
declare DEBUG_FLAG=false

while [[ "$#" -gt 0 ]]
do
  case $1 in
    -r | --resource)
      resourceName=$2
      ;;

    -p | --providerVersion)
      providerVersion=$2
      ;;

    -d | --debug )             
      DEBUG_FLAG=true
      ;; 
  esac
  shift
done

main() {
    
  clone_repo

  pushd terraform-provider-azurerm
    debug "git checkout $providerVersion"
    quiet_git checkout $providerVersion 

    resourceFile=$(grep -Ril "\"${resourceName}\"" . --include=*.go --exclude=*_test.go --exclude=*registration.go)
   popd

  if [ ! -z $resourceFile ]; then
    module=$(go run ./main.go -r "./terraform-provider-azurerm/$resourceFile" --json | jq -r '.importModule')
    debug "\nResource Source File:      $resourceFile"
    debug "Found Azure SDK Go Module: $module \n"
    if [ "$DEBUG_FLAG" == false ];then
      echo $module
    fi
  else
    echo "No go files found for resource $resourceName"
  fi

  cleanup
}

quiet_git() {
  if [ "$DEBUG_FLAG" == false ];then
    stdout=$(tempfile)
    stderr=$(tempfile)

    if ! git "$@" </dev/null >$stdout 2>$stderr; then
        cat $stderr >&2
        rm -f $stdout $stderr
        exit 1
    fi

    rm -f $stdout $stderr
  else 
    git "$@"
  fi
}

clone_repo() {
  quiet_git clone git@github.com:terraform-providers/terraform-provider-azurerm.git
}

cleanup() {
  rm -rf terraform-provider-azurerm
}

debug() {
  if [ "$DEBUG_FLAG" == true ];then
    echo -e $1
  fi
}
pushd () {
    command pushd "$@" > /dev/null
}

popd () {
    command popd "$@" > /dev/null
}

#############################
main
