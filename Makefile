SRC_DIR=src
_FILES=main.go
FILES=$(addprefix ${SRC_DIR}/,${_FILES})
CC=go build
COPTS=

all:
	$(CC) -o gimme $(COPTS) $(FILES) 

clean:
	-rm gimme *.o