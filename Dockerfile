# Start from golang v1.11 base image
FROM golang:1.14.4 as builder
LABEL maintainer="Glenn Pringle <glenn@pringle.com.au>"
ENV GO111MODULE=on
WORKDIR /app

COPY go.mod .
COPY go.sum .

# Add ssh key for bitbucket
#COPY rsa_ssh.key /root/.ssh/id_rsa
#RUN chmod 400 /root/.ssh/id_rsa

#Trust bitbucket ssh-rsa
#RUN echo bitbucket.org ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAubiN81eDcafrgMeLzaFPsw2kNvEcqTKl/VqLat/MaB33pZy0y3rJZtnqwR2qOOvbwKZYKiEO1O6VqNEBxKvJJelCq0dTXWT5pbO2gDXC6h6QDXCaHo6pOHGPUy+YBaGQRGuSusMEASYiWunYN0vCAI8QaXnWMXNMdFP3jHAJH0eDsoiGnLPBlBp4TNm6rYI74nMzgz3B9IikW4WVK+dc8KZJZWYjAuORU3jc1c/NPskD2ASinf8v3xnfXeukU0sJ5N6m5E8VLjObPEO+mN2t/FZTMZLiFqPWc/ALSqnMnnhwrNi2rbfg/rd/IpL8Le3pSBne8+seeFVBoGqzHM9yXw== >> /root/.ssh/known_hosts
#RUN git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/go-app .

######## Start a new stage from scratch #######
#FROM envoyproxy/envoy-alpine:v1.14.1
FROM alpine as alpine

RUN apk --no-cache add ca-certificates
WORKDIR /root
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/go-app /usr/local/bin/go-app
#COPY --from=builder /app/envoy/envoy_template.yaml /root/envoy/envoy_template.yaml
COPY --from=builder /app/static /root/static

ENTRYPOINT /usr/local/bin/go-app
