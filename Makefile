.PHONY: default build clean

default: build

build:
	make -C src/

clean:
	make -C src/ clean
	rm -rf goserver
