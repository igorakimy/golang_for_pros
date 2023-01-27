package main

import (
	"fmt"
	"net"
)

func main() {
	// Получить все сетевые интерфейсы текущего компьютера в виде среза,
	// содержащего элементы net.Interface.
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Обойти все сетевые интерфейсы.
	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)
		// Получить сетевой интерфейс по имени.
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		// Получить срез сетевых адресов интерфейса.
		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
	}
}
