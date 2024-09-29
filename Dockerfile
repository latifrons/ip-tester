FROM golang:1.22-bullseye AS builder

RUN useradd -m builder
WORKDIR /src

ADD ./go.mod ./
RUN go mod download

COPY . .

RUN chown -R builder:builder /src
USER builder

RUN go build -a -o main .

# Copy OG into basic alpine image
FROM ubuntu:24.04

RUN useradd -m user

WORKDIR /www
COPY --from=builder --chown=exchange:exchange /src/main .

RUN chown -R user:user .

USER user

ENTRYPOINT ["./main"]