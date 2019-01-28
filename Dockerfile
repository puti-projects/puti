############################
# Builder image
############################
ARG GOLANG_BUILDER_VERSION=alpine
FROM golang:${GOLANG_BUILDER_VERSION} AS builder

RUN apk update && apk add --no-cache build-base git tzdata && update-ca-certificates

COPY . /puti
WORKDIR /puti

# Build the binary.
RUN make


############################
# Usage image
############################
FROM scratch

# Import from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable.
COPY --from=builder /puti /puti

WORKDIR /puti

EXPOSE 8000 8080
# Run the puti binary.
CMD ["./puti"]