############################
# Builder image
############################
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git make ca-certificates tzdata && update-ca-certificates
# Create putiuser.
RUN adduser -D -g '' putiuser

WORKDIR /puti
COPY . /puti

# Build the binary.
RUN make


############################
# Usage image
############################
FROM scratch

# Import from the builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable.
COPY --from=builder /puti/puti /puti
# Use an unprivileged user.
USER putiuser

EXPOSE 8000
EXPOSE 8080
# Run the puti binary.
CMD ["puti"]