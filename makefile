TARGET = ./build/
NAME = gowakeup


.PHONY: run clean

run:
	go run main.go wake 12:34:56:78:9A:BC

linux:
	$(eval GOOS=linux)
	$(eval GOARCH=amd64)
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${TARGET}/${NAME}-${GOOS}-${GOARCH}

pi_linux:
	$(eval GOOS=linux)
	$(eval GOARCH=arm)
	$(eval GOARM=5)
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${TARGET}/${NAME}-${GOOS}-${GOARCH}-v${GOARM}

windows:
	$(eval GOOS=windows)
	$(eval GOARCH=amd64)
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${TARGET}/${NAME}-${GOOS}-${GOARCH}.exe

darwin:
	$(eval GOOS=darwin)
	$(eval GOARCH=amd64)
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${TARGET}/${NAME}-${GOOS}-${GOARCH}

clean:
	${RM} -r ${TARGET}/

all: linux pi_linux windows darwin