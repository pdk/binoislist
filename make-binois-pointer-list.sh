#!/bin/bash

OUTPUT_FILE=$1
PKG_NAME=$2
BASE_TYPE=$3
LIST_TYPE=$4

if [ -z "$LIST_TYPE" ]
then
    echo "usage: make-binois-pointer-list.sh newfile.go packagename basetype listtype"
    exit 1
fi

echo '// Code generated with make-binois-pointer-list.sh DO NOT EDIT.' > "$OUTPUT_FILE"
echo '// http://github.com/pdk/binoislist'  >> "$OUTPUT_FILE"
echo '' >> "$OUTPUT_FILE"

sed -e "s/binoislist/$PKG_NAME/g;" \
    -e "/const UnknownBinois/d;" \
    -e "s/UnknownBinois/nil/g;" \
    -e "s/UnknownBinois/Unknown$BASE_TYPE/g;" \
    -e "s/BinoisList/$LIST_TYPE/g;" \
    -e "s/BinoisPtr/\\*$BASE_TYPE/g;" \
    -e "s/Binois/$BASE_TYPE/g;" \
    < $GOPATH/src/github.com/pdk/binoislist/list.go >> "$OUTPUT_FILE"
