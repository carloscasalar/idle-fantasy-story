FROM golang:alpine AS build
WORKDIR /go/src/api
COPY . .
RUN go build -o /go/bin/api cmd/api/main.go
RUN mkdir -p /go/bin/data
COPY init/storage/inmemory/worlds.yml /go/bin/data/worlds.yml

FROM scratch
COPY --from=build /go/bin/api /go/bin/api
COPY --from=build /go/bin/data /go/bin/data
ENTRYPOINT ["/go/bin/api"]