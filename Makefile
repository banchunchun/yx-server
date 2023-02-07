APP=cmd\server\main.go

windows:
		GOOS=windows go build -o ${APP}
linux:
		GOOS=linux GOARCH=amd64 go build -o ${APP}