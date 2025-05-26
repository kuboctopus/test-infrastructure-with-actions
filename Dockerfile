FROM golang:1.24-alpine

WORKDIR /test

COPY go.mod go.sum ./
COPY zone_test.go ./
RUN go mod download && go mod verify

RUN apk add --no-cache bind

COPY --chown=named:named zonefiles /var/cache/bind
COPY --chown=named/named named.conf /etc/bind/named.conf
CMD ["named", "-c", "/etc/bind/named.conf", "-g", "-u", "named"]