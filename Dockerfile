FROM golang:1.18.4 as build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY . /workspace
WORKDIR /workspace
RUN go build -o out/wol -mod vendor

FROM scratch
COPY --from=build /workspace/out/wol /
ENTRYPOINT [ "/wol"]
