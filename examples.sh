#! /bin/bash

go build  -o ./handlergen .

./handlergen -file _example/advanced/advanced.yaml > _example/advanced/generated.go
./handlergen -file _example/basic/basic.yaml > _example/basic/generated.go
./handlergen -file _example/petstore/petstore.yaml -format openapi -pkg main > _example/petstore/generated.go
./handlergen -file _example/rest/rest.yaml -pkg main > _example/rest/handler.go

rm ./handlergen

echo "refreshed examples"