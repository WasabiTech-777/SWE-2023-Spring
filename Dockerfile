FROM golang:1.19.0
ENV GOPATH /GO/src/github.com/SWE-2023-Spring/api
WORKDIR /GO/src/github.com/SWE-2023-Spring/dbTest
COPY . /GO/src/github.com/SWE-2023-Spring/dbTest
RUN go mod tidy