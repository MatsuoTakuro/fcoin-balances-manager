FROM golang:1.19.1-bullseye as builder
WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN cd cmd && go build -trimpath -ldflags="-w -s" -o "fcoin-balances-manager"

FROM gcr.io/distroless/base-debian11 as dev
COPY --from=builder opt/app/cmd/fcoin-balances-manager /fcoin-balances-manager
CMD ["/fcoin-balances-manager"]
