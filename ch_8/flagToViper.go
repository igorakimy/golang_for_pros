package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	flag.Int("i", 100, "i parameter")
	// Импорт данных из пакета flag в пакет pflag.
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	// Для синтаксического анализа уже вызывается pflag.Parse().
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	// Получить значение целочисленного аргумента командной строки.
	i := viper.GetInt("i")
	fmt.Println(i)
}
