#!/bin/bash

OUTPUT_FILE=$1
PKG_NAME=$2
BASE_TYPE=$3
LIST_TYPE=$4
UNKNOWN_VALUE=$5

if [ -z "$UNKNOWN_VALUE" ]
then
    echo "usage: make-binois-list.sh newfile.go packagename basetype listtype unknownvalue"
    exit 1
fi

echo '// Code generated with make-binois-list.sh DO NOT EDIT.' > "$OUTPUT_FILE"
echo '// http://github.com/pdk/binoislist'  >> "$OUTPUT_FILE"
echo >> "$OUTPUT_FILE"

sed -e "s/binoislist/$PKG_NAME/g;" \
    -e "/const UnknownBinois/d;" \
    -e "s/UnknownBinois/$UNKNOWN_VALUE/g;" \
    -e "s/BinoisList/$LIST_TYPE/g;" \
    -e "s/BinoisPtr/$BASE_TYPE/g;" \
    -e "s/Binois/$BASE_TYPE/g;" \
    < $GOPATH/src/github.com/pdk/binoislist/list.go >> "$OUTPUT_FILE"