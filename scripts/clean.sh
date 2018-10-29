#!/bin/bash
if [ -f $GOPATH/bin/$@ ]; then
    rm $GOPATH/bin/$@
fi
