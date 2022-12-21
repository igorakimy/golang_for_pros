package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	// Получить аргумент командной строки
	arguments := os.Args
	// Если аргумент не был передан
	if len(arguments) == 1 {
		fmt.Println("Please provide one test file to process!")
		// Завершить программу
		os.Exit(1)
	}

	// Получить первый аргумент
	filename := arguments[1]
	// Попытаться открыть файл с названием указанным в аргументе
	f, err := os.Open(filename)
	// Если открыть не удалось
	if err != nil {
		fmt.Printf("error opening file %s", err)
		// Завершить программу
		os.Exit(1)
	}
	// Отложить закрытие файла до конца выполнения функции
	defer f.Close()

	// Хранение количества строк входного файла, которые не соответствуют ни одному из
	// двух регулярных выражений программы.
	notAMatch := 0
	// Считать файловый процесс в буфер для построчного чтения
	r := bufio.NewReader(f)
	for {
		// Попытаться считать строку, пока не встретится символ переноса строки
		line, err := r.ReadString('\n')
		// Если достигнут конец файла
		if err == io.EOF {
			// Остановить цикл
			break
			// Если произошла ошибка при чтении файла
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
		}

		// Скомпилировать шаблон регулярного выражения
		r1 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)
		// Проверить, есть ли совпадение с шаблоном в строке
		if r1.MatchString(line) {
			// Получить срез строк подмасок заключенных в () в регулярном выражении
			match := r1.FindStringSubmatch(line)
			// Спарсить время из строки с которой найдено совпадение
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			// Если парсинг времени успешен
			if err == nil {
				// Привести дату к определенному новому формату
				newFormat := d1.Format(time.Stamp)
				// Заменить в исходной строке время на новый формат и вывести строку
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				// Если соответствий не было найдено, увеличить счетчик
				notAMatch++
			}
			// Продолжить выполнение цикла
			continue
		}

		r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\].*`)
		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}
	}
	fmt.Println(notAMatch, "lines did not match!")
}
