FROM golang:1.19.0

WORKDIR /dbTest
COPY . .
RUN go mod tidy
ENTRYPOINT ["C:\Users\marin\go\bin\dlv.exe"]