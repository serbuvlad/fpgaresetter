#!/bin/bash

set -ex

go build -buildvcs=false .
mkdir -p /usr/local/bin
mv fpgaresetter /usr/local/bin
mkdir -p /usr/local/share/fpgaresetter
cp printtargets.tcl reset.tcl /usr/local/share/fpgaresetter
