.PHONY: default test clean 

default: build

build:
	make -C src/

test:
	make -C test/ clean 
	make -C test/ test

clean:
	make -C src/ clean
	rm -rf goserver
