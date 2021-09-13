FROM golang:1.17.1-bullseye AS builder
COPY . /src
WORKDIR /src
ENV CGO_ENABLED=0
RUN go build -o gud-git ./cmd/gud-git

FROM scratch
COPY --from=builder /src/gud-git /gud-git
ENTRYPOINT ["/gud-git"]
