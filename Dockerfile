FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./update-operator /update-operator

ENTRYPOINT ["/update-operator"]
