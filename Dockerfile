############################
# Builder image
############################
ARG GOLANG_BUILDER_VERSION=1.12-alpine
FROM golang:${GOLANG_BUILDER_VERSION} AS builder

RUN apk update && apk add --no-cache build-base git tzdata ca-certificates && update-ca-certificates

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
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV GOSU_VERSION 1.11
RUN set -eux; \
	\
	apk add --no-cache --virtual .gosu-deps \
		dpkg \
		gnupg \
	; \
	\
	dpkgArch="$(dpkg --print-architecture | awk -F- '{ print $NF }')"; \
	wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch"; \
	wget -O /usr/local/bin/gosu.asc "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch.asc"; \
	\
    # verify the signature
	export GNUPGHOME="$(mktemp -d)"; \
    # for flaky keyservers, consider https://github.com/tianon/pgp-happy-eyeballs, ala https://github.com/docker-library/php/pull/666
	gpg --batch --keyserver ha.pool.sks-keyservers.net --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4; \
	gpg --batch --verify /usr/local/bin/gosu.asc /usr/local/bin/gosu; \
	command -v gpgconf && gpgconf --kill all || :; \
	rm -rf "$GNUPGHOME" /usr/local/bin/gosu.asc; \
	\
    # clean up fetch dependencies
	apk del --no-network .gosu-deps; \
	\
	chmod +x /usr/local/bin/gosu; \
    # verify that the binary works
	gosu --version; \
	gosu nobody true

WORKDIR /app/puti/
# Copy useful files or filepath to path /app
# binary
COPY --from=builder /puti/puti ./puti
# backend html
COPY --from=builder /puti/console ./console
# other static file
COPY --from=builder /puti/assets ./assets

# these should set volume 
# config
COPY --from=builder /puti/configs /app/init/configs
# theme path
COPY --from=builder /puti/theme /app/init/theme
# upload path
COPY --from=builder /puti/uploads /app/init/uploads

RUN addgroup -g 1000 -S putiuser \
    && adduser -u 1000 -S putiuser -G putiuser 

# copy script
COPY --from=builder /puti/scripts/docker/docker-entrypoint.sh /puti/scripts/docker/wait-for-db.sh /usr/local/bin/
RUN ln -s /usr/local/bin/docker-entrypoint.sh /

VOLUME ["/data"]
EXPOSE 8000 8080

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["puti"]
