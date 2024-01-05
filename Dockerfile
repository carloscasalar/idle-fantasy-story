FROM golang:alpine AS build
WORKDIR /go/src/api
COPY . .
RUN go build -o /go/bin/api cmd/api/main.go

FROM scratch
COPY --from=build /go/bin/api /go/bin/api
ENTRYPOINT ["/go/bin/api"]