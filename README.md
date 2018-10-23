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
