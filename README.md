# TinyVM-Go
---
TinyVM-Go is a virtual machine with the goal of having a small footprint.
Translating the source code into bytecodes that we can operate.

TinyVM-Go is inspired by the project [tinyvm] (https://github.com/jakogut/tinyvm) and implemented it in pure go.

## Building

Run
```bash
go get -v github.com/EricYT/tinyvm-go
cd $GOPATH/src/github.com/EricYT/tinyvm-go
make
or
make build
```
to complie these codes. This will use a docker image to build it, with
the current directory volume-mounted into place. This will store incrmental
state for the fastest possible build.
Run `make all-build` to build for all architectures.
Run `make clean` to clean up.

## Running
The program source file `./programs/tinyvm/jsr.vm`
```assembly
print_eax:
        push ebp
        mov ebp, esp
        prn eax
        pop ebp
        ret

start:
        mov eax, 42
        call print_eax

        mov eax, 23
        call print_eax
```
Run
```bash
./bin/amd64/tvm ./programs/tinvm/jsr.vm
2018/01/11 21:26:10 Prepare to interpret the file programs/tinyvm/jsr.vm
42
23
```
## Debuging
There a simple debug tool named `tdb` inspired by `gdb`.
```bash
$ cat ./programs/tinyvm/jsr.vm
print_eax:
        push ebp
        mov ebp, esp
        prn eax
        pop ebp
        ret

start:
        mov eax, 42
        call print_eax

        mov eax, 23
        call print_eax
$ ./bin/amd64/tdb programs/tinyvm/jsr.vm
2018/01/11 21:32:10 Prepare to interpret the file programs/tinyvm/jsr.vm
tdb >: b 1
WARNING: "b" is not a valid command.
tdb >: break 1
tdb >: run
BreakPoint hit at address: 1
tdb >: step
Advancing instruction pointer to 2
tdb >: step
42
Advancing instruction pointer to 3
tdb >: continue
BreakPoint hit at address: 1
tdb >: step
Advancing instruction pointer to 2
tdb >: continue
23
End of program readched.
```

## TODO
- debug print information of registers
