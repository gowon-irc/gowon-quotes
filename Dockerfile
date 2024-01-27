FROM golang:alpine as build-env
COPY . /src
WORKDIR /src
RUN go build -o gowon-quotes

FROM alpine:3.19.1
WORKDIR /app
COPY --from=build-env /src/gowon-quotes /app/
ENTRYPOINT ["./gowon-quotes"]
