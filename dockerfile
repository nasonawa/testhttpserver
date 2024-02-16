FROM registry.access.redhat.com/ubi9/go-toolset:1.20.10-11 AS builder

COPY ./main.go ./main.go


RUN go mod init testserver && go mod tidy
RUN go build .

FROM redhat/ubi9-minimal

COPY --from=builder /opt/app-root/src/testserver ./testserver

EXPOSE 8080
ENTRYPOINT [ "./testserver" ]