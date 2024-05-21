#!/usr/bin/env bash

APP_VERSION='v1.0.0'


 
output_name='QHttp_' 
echo 'Building..: '${output_name}${APP_VERSION}
go build  -ldflags="-X 'main.FeatureSet=ALL'  -X 'main.Version=${APP_VERSION}'" -o ./bin/${output_name}${APP_VERSION}  ./cmd/web


docker build --build-arg="APP_VERSION=${APP_VERSION}"  -t zerobittech/qhttp:${APP_VERSION} .
docker push zerobittech/qhttp:${APP_VERSION}




# GOOS='windows'
# GOARCH='amd64'

 

# output_name='QHttp_win_' 
# echo 'Building..: '${output_name}${APP_VERSION}
# env GOOS=$GOOS GOARCH=$GOARCH go build  -ldflags="-X 'main.FeatureSet=ALL' -X 'main.Version=${APP_VERSION}'" -o ./bin/${output_name}${APP_VERSION}.exe ./cmd/web
