SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build  -o bin/cyj main.go

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build  -o bin/cyj.exe main.go