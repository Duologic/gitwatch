FROM golang:1.14 as base

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

FROM base AS builder

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .


FROM alpine/git:v2.26.2

COPY --from=builder /app/gitwatch /app/gitwatch

RUN mkdir ${HOME}/.ssh && \
    ssh-keyscan -t rsa github.com >> ${HOME}/.ssh/known_hosts

ENTRYPOINT ["/app/gitwatch"]
