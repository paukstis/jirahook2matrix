OUT = jirahook2matrix
# for BSD make
#VERSION != git describe --always --long
# for GNU make
VERSION = $(shell git describe --always --long)

all: build

build:
	go build -i -v -o ${OUT} -ldflags="-X main.version=${VERSION}"

static:
	go build -i -v -o ${OUT}-v${VERSION} -tags netgo -ldflags="-extldflags \"-static\" -w -s -X main.version=${VERSION}" ${PKG}

run: build
	./${OUT}

clean:
	-@rm ${OUT} ${OUT}-v*

.PHONY: run build static
