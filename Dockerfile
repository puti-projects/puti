############################
# Builder image
############################
ARG GOLANG_BUILDER_VERSION=alpine
FROM golang:${GOLANG_BUILDER_VERSION} AS builder

RUN apk update && apk add --no-cache build-base git tzdata && apk add ca-certificates

# Create putiuser
RUN adduser -D -g "" putiuser

COPY . /puti
WORKDIR /puti

# Build the binary.
RUN make


############################
# Usage image
############################
FROM alpine:latest

# Import from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy useful files or filepath to path /app
WORKDIR /app/puti
# binary
COPY --from=builder /puti/puti puti
# backend html
COPY --from=builder /puti/console console
# other static file
COPY --from=builder /puti/assets assets

# these should set volume
# config
COPY --from=builder /puti/configs configs
# theme path
COPY --from=builder /puti/theme theme
# upload path
COPY --from=builder /puti/uploads uploads

RUN mkdir -p /data/puti \
    && chown -R putiuser /data /data/puti /app/puti
VOLUME ["/data"]

COPY --from=builder /puti/scripts/docker/docker-entrypoint.sh /usr/local/bin/
RUN ln -s /usr/local/bin/docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]

USER putiuser

EXPOSE 8000 8080
CMD ["./puti"]
