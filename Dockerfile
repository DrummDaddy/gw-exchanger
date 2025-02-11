FROM golang:1.22 

WORKDIR /app 

COPY go.mod ./ 
COPY go.sum ./ 
RUN go mod download 

COPY . ./ 

RUN go build -0 main ./cmd/main.go 

EXPOSE 50051 

CMD ["./main", "-c", "config.env"]