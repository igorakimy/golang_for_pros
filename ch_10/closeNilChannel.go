package main

func main() {
	// Попытка закрыть нулевой канал завершится паникой.
	var c chan string
	close(c)
}
