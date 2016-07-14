#!/bin/bash

export VTTOP=$(pwd)
export VTROOT="${VTROOT:-${VTTOP/\/src\/github.com\/flike\/kingshard/}}"
# VTTOP sanity check
if [[ "$VTTOP" == "${VTTOP/\/src\/github.com\/flike\/kingshard/}" ]]; then
  echo "WARNING: VTTOP($VTTOP) does not contain src/github.com/flike/kingshard"
fi

export GOTOP=$VTTOP
function prepend_path()
{
  if [ -d "$2" ] && [[ ":$1:" != *":$2:"* ]]; then
    echo "$2:$1"
  else
    echo "$1"
  fi
}

export GOPATH=$(prepend_path $GOPATH $VTROOT)
