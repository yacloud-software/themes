.PHONY:	client server
all:	client server
client:
	cd client && go install ${LDFLAGS} themes-client.go
server:
	cd server && go install ${LDFLAGS} themes-server.go `ls -1 *.go|grep -v themes-server.go`
