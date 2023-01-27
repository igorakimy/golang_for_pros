package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: URL\n")
		return
	}

	URL := os.Args[1]
	client := http.Client{}

	req, _ := http.NewRequest("GET", URL, nil)
	// Отслеживание HTTP-запроса. Объект httptrace.ClientTrace определят события,
	// которые нас интересуют.
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}

	// Функция httptrace.WithClientTrace() возвращает новый контекст, полученный
	// на основе заданного родительского контекста.
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("Requesting data from server!")
	// Метод http.DefaultTransport.RoundTrip() служит оберткой для
	// объекта http.DefaultTransport.RoundTrip, чтобы он отслеживал текущий запрос.
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Фактическое выполнение запроса к веб-серверу с помощью функции Do().
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Получение HTTP-данных и вывод их на экран.
	_, _ = io.Copy(os.Stdout, response.Body)
}
