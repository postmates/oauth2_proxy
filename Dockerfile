FROM golang:1.8-alpine
ADD . $GOPATH/src/github.com/bitly/oauth2_proxy
WORKDIR $GOPATH/src/github.com/bitly/oauth2_proxy

# Build *really static*
ENV CGO_ENABLED 0
RUN go install -ldflags '-s' github.com/bitly/oauth2_proxy

# This requires Docker version 17.05 or newer
FROM scratch
COPY --from=0 /go/bin/oauth2_proxy /bin/oauth2_proxy
ENTRYPOINT ["/bin/oauth2_proxy"]
