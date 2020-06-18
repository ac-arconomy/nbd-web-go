## Start from golang v1.11 base image
#FROM golang:1.14.4 as builder
#LABEL maintainer="Glenn Pringle <glenn@pringle.com.au>"
#ENV GO111MODULE=on
#WORKDIR /app
#
#COPY go.mod .
#COPY go.sum .
#
## Get dependancies - will also be cached if we won't change mod/sum
#RUN go mod download
#
#COPY . .
#
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
#      -ldflags='-w -s -extldflags "-static"' -a \
#      -o /go/bin/go-app .
#
######### Start a new stage from scratch #######
##FROM envoyproxy/envoy-alpine:v1.14.1
#FROM scratch
## Copy the Pre-built binary file from the previous stage
#COPY --from=builder /go/bin/go-app /go/bin/go-app
#COPY --from=builder /app/static /go/bin/static
#
#CMD ["/go/bin/go-app"]


FROM golang:alpine as golang
WORKDIR /go/src/app
COPY . .
# Static build required so that we can safely copy the binary over.
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"'

FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
# -0 means no compression.  Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -q -r -0 /zoneinfo.zip .

FROM scratch
# the test program:
COPY --from=golang /go/bin/app /app
# the timezone data:
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine /zoneinfo.zip /
# the tls certificates:
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app"]