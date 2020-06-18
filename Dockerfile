# Start from golang v1.11 base image
FROM golang:1.14.4 as builder
LABEL maintainer="Glenn Pringle <glenn@pringle.com.au>"
ENV GO111MODULE=on
WORKDIR /app

COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o /go/bin/go-app .

######## Start a new stage from scratch #######
#FROM envoyproxy/envoy-alpine:v1.14.1
FROM scratch
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/go-app /go/bin/go-app
COPY --from=builder /app/static /go/bin/static

CMD ["/go/bin/go-app"]
