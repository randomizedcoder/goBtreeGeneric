#
# Makefile
#

# ldflags variables to update --version
# short commit hash
COMMIT :=$(shell /usr/bin/git describe --always)
DATE :=$(shell /bin/date -u +"%Y-%m-%d-%H:%M")
BINARY := goBtreeGeneric
BINARYSYNC := doubleSyncMap

all: clean build

test:
	go test

clean:
	[ -f ${BINARY} ] && /bin/rm -rf ./${BINARY} || true

build:
	CGO_ENABLED=0 go build -ldflags "-X main.commit=${COMMIT} -X main.date=${DATE}" -o ./${BINARY} ./${BINARY}.go

buildsync:
	CGO_ENABLED=0 go build -ldflags "-X main.commit=${COMMIT} -X main.date=${DATE}" -o ./${BINARYSYNC} ./${BINARYSYNC}.go