#!/usr/bin/env bash

set -eo pipefail
echo "Generating proto code"
cd proto
buf generate

cd ..

# move proto files to the right places
rm -f gen/coudaang/openapi.pb.go
cp gen/openapiv2.swagger.json docs/swagger-ui/swagger.json
cp -r gen/coudaang/* packages/api/

rm -rf gen