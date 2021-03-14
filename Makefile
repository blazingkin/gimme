SRC_DIR=src
_FILES=main.go
FILES=$(addprefix ${SRC_DIR}/,${_FILES})
CC=go build
COPTS=

all:
	$(CC) $(COPTS) $(FILES)

clean:
	-rm main *.o