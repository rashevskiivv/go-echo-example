FROM golang
WORKDIR /go/src/go-fiber-api-docker
COPY . .
RUN go mod download
RUN go build -o server main.go
EXPOSE 8000
CMD [ "./server" ]