.PHONY:	client server
all:	client server
client:
	cd client && go install themes-client.go
server:
	cd server && go install themes-server.go `ls -1 *.go|grep -v themes-server.go`
