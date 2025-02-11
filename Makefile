build:
go build -o ./cmd/main.go 

run:
./main -c config.env

docker-build:
docker build -t gw-exchanger .

docker-run: 
docker run -p 50051:50051 --env-file=config.env gw-exchanger