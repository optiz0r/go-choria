#!/bin/bash

set -x

apt-get install wixl
find .
cd {{cpkg_name}}-{{cpkg_version}}
wixl -o {{cpkg_name}}-{{cpkg_version}}.msi dist/choria.wxs
mv {{cpkg_name}}-{{cpkg_version}}.msi ..
