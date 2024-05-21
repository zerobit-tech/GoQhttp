#!/usr/bin/env bash

APP_VERSION='v1.0.0'

output_name='QHttp_' 
rm ./bin/${output_name}${APP_VERSION}


echo 'Building..: '${output_name}${APP_VERSION}
go build  -ldflags="-X 'main.FeatureSet=ALL'  -X 'main.Version=${APP_VERSION}'" -o ./bin/${output_name}${APP_VERSION}  ./cmd/web



ls -al ./bin
 
docker build -t zerobittech/qhttp .
docker login
docker push zerobittech/qhttp
 