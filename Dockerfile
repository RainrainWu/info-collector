# build stage
FROM golang:alpine AS build-env
ADD . /build
RUN cd /build && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/info-collector
 
# final stage
FROM centurylink/ca-certs
COPY --from=build-env /build /
ENTRYPOINT ["/build"]
