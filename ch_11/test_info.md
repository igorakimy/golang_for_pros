Выполнить тесты:
```bash
go test testMe.go testMe_test.go
```
Выполнить тесты с отображением детальной информации:
```bash
go test testMe.go testMe_test.go -v
```
Выполнить только некоторые тесты по регулярному выражению:
```bash
go test testMe.go testMe_test.go -run="F2" -v
```
Проверить покрытие кода:
```bash
go test -cover -v
```
Сгенерировать профиль покрытия тестами:
```bash
go test -coverprofile=coverage.out
```
Анализ файла в котором содержится профиль покрытия тестами:
```bash
go tool cover -func=coverage.out
```
Анализ файла в котором содержится профиль покрытия тестами в браузере:
```bash
go tool cover -html=coverage.out
```
Анализ профиля покрытия тестами в браузере и сохранение результата в html файл:
```bash
go tool cover -html=coverage.out -o output.html
```
Выполнить тесты с таймаутом:
```bash
go test too_long.go too_long_test.go -v -timeout 20s
```
Выполнить бенчмарк тесты:
```bash
go test -bench=. benchmarkMe.go benchmarkMe_test.go 
```
Выполнить бенчмарк тесты с отображением статистики выделенной памяти:
```bash
go test -benchmem -bench=. benchmarkMe.go benchmarkMe_test.go 
```
Скомпилировать бинарный файл для под нужную ОС и архитектуру:
```bash
env GOOS=linux GOARCH=arm go build xCompile.go
```
Скомпилировать в бинарный файл и отобразить машинный код:
```bash
env GOSSAFUNC=main GOOS=windows GOARCH=amd64 go build -gcflags "-S" machineCode.go
```