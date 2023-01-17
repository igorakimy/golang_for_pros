Компилировать пакет main в текущем каталоге и начать его отладку:
```bash
dlv debug
```
Запуск отладки с передачей аргументов:
```bash
dlv debug -- arg1 arg2 arg3
```
Продолжить выполнение из оболочки отладчика, как если бы была введена команда `go run`:
```bash
(dlv) c
```
Перезапустить программу в отладчике для повторной отладки:
```bash
(dlv) r
```
Установить точку останова для функции с названием function:
```bash
(dlv) break function
```
Перейти к следующей операции во время отладки:
```bash
(dlv) next
```
Вывести значение переменной во время отладки:
```bash
(dlv) print i
```
Отладка Go-тестов:
```bash
dlv test
```