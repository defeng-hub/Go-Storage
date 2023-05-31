CC=mingw32-gcc
version:
	echo "go-storage"
	echo "version:v3.0.0"

run:
	go run main.go \
		-Sql_Dsn "go_gin_api:go_gin_api@tcp(159.75.37.58:3306)/go_gin_api" \
		-Oss_Type 0 \
		-Run_Url "http://127.0.0.1:3000"




build:
	set GOOS=linux
	set GOARCH=amd64
	go build -o Go-Storage main.go

clean:
	del Go-Storage