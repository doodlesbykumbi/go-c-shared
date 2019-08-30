all:
	go build -buildmode=c-shared -o libgo.so lib.go
	go build -o main main.go
run: all
	LD_LIBRARY_PATH=. ./main
