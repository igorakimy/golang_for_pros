Для того чтобы скомпилировать исходный файл Go, нужно воспользоваться командой:
```bash
go tool compile unsafe.go
```
В результате получится объектный файл - файл с расширением **.o**
#
Чтобы создать архивный файл, при выполнении компиляции нужно добавить флаг *-pack*
```bash
go tool compile -pack unsafe.go
```
В результате получится файл с расширением **.a**
#
Чтобы создать C-библиотеку, необходимо выполнить следующие команды:
```bash
ls -l callClib/
gcc -c callClib/*.c
ls -l callC.o
file callC.o
/usr/bin/ar rs callC.a *.o
ls -l callC.a
file callC.a
rm callC.o
```
#
Чтобы сгенерировать из Go-кода общую библиотеку C, нужно выполнить следующую команду:
```bash
go build -o usedByC.o -buildmode=c-shared usedByC.go
```
#
Чтобы скомпилировать C-код, который вызывает Go-функции, нужно выполнить команду:
```bash
gcc -o willUseGo willUseGo.c .usedByC.o
```
#
Чтобы создать код WebAssembly, необходимо выполнить следующую команду:
```bash
GOOS=js GOARCH=wasm go build -o main.wasm toWasm.go
```