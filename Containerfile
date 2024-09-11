FROM docker.io/library/golang:1.23.1-bookworm AS builder
COPY . /src
WORKDIR /src
ENV CGO_ENABLED=0
RUN go build -o gud-git ./cmd/gud-git

FROM scratch
COPY --from=builder /src/gud-git /gud-git
ENTRYPOINT ["/gud-git"]
