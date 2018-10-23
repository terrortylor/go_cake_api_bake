# go_cake_api_mix
A basic API for testing with.
It's part of three API's that pass JSON around and manipulate it.

It has the following endpoints:

* POST bake - bakes ingredients in bowl, returns a cake

# Dependencies
golang's dep is used for dependency management:
https://github.com/golang/dep
'''
dep ensure
'''

# Run
Either just run it:
```
go run main.go
```

or compile and then run executable:
```
go build
./go_cake_api_bake
```

# cURL examples:
## POST /heatoven
```
curl -i -X POST localhost:8000/heatoven -H 'Content-Type: application/json' -d '{"temperature": 4}'
```

## POST /bake
```
curl -i -X POST localhost:8000/bake -H 'Content-Type: application/json' -d '{"name":"test","Ingredients":[{"name":"eggs","quantity":"all the fucking eggs"}]}'
```

# Docker
If you want to build this into a docker image you must ensure that all Dependencies are available and that it's compiled for linux:
```
dep ensure
env GOOS=linux GOARCH=arm go build -v
docker image build . --tag go_cake_api_bake
```

Run the image as a container:
```
docker run -d -p 8080:8000 --name bake go_cake_api_bake
```
Note: that this exposes 8000 from the container to 8080 locally, so the curl commands examples above require changing the port to 8080

Check logs:
```
docker logs bake
```

Kill and remove container:
```
docker kill bake && docker rm bake
```

# Harbor
If you are using harbor then upload:
Note: remember to login!
```
docker login harbour.whateva
```

```
docker tag go_cake_api_bake:latest harbour.whateva/cake_factory/go_cake_api_bake:1.0.0
docker push harbour.whateva/cake_factory/go_cake_api_bake:1.0.0
```
