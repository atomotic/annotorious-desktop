#!/usr/bin/env bash

cd public
yarn
cp node_modules/openseadragon/build/openseadragon/openseadragon.min.js js/
cp node_modules/openseadragon/build/openseadragon/openseadragon.min.js.map js/
cp -av node_modules/openseadragon/build/openseadragon/images .
cp node_modules/@recogito/annotorious-openseadragon/dist/openseadragon-annotorious.min.js js/
cp node_modules/@recogito/annotorious-openseadragon/dist/openseadragon-annotorious.min.js.map js/
rm -fr node_modules

cd ..
statik -src=./public

go build
