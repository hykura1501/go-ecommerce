FROM golang:1.24

WORKDIR /be-ecommerce

COPY go.mod go.sum ./
RUN go mod download


COPY . .
RUN go build -v -o main ./main.go

CMD [ "./main" ]