#!/usr/bin/env bash

cd public
yarn
cp node_modules/openseadragon/build/openseadragon/openseadragon.min.js js/
cp node_modules/openseadragon/build/openseadragon/openseadragon.min.js.map js/
cp -a node_modules/openseadragon/build/openseadragon/images .
cp node_modules/@recogito/annotorious-openseadragon/dist/openseadragon-annotorious.min.js js/
cp node_modules/@recogito/annotorious-openseadragon/dist/openseadragon-annotorious.min.js.map js/
cp node_modules/@recogito/annotorious-openseadragon/dist/annotorious.min.css css/
cp node_modules/@recogito/annotorious-openseadragon/dist/annotorious.min.css.map css/
rm -fr node_modules

cd ..
statik -src=./public

go build -o build/macos/Annotorious.app/Contents/MacOS/Annotorious

