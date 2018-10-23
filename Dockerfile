FROM golang:latest

EXPOSE 8000
WORKDIR /usr/local/bin
COPY go_cake_api_bake go_cake_api_bake

ENTRYPOINT ./go_cake_api_bake
