# go-c-shared
This is a POC of using a C-shared library (created from Go) in another Go program.
It's composed of 3 files:
- `lib.go` The library
- `main.go` The main Go program leveraging the C-shared library
- `Makefile` For easy running and building. Call `make run` to run and build.

If you'd like to try this in Docker:

```
docker run --rm -w /work -v $PWD:/work golang make run
```
