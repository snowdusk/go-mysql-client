package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	gomysqlclient "github.com/snowdusk/go-mysql-client"
)

func main() {
	var c gomysqlclient.Config
	var fileName string

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `Usage of %s:
gomysql -P 3306 -h localhost -u user -p password database_name\n`, os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&c.Host, "h", "localhost", "host")
	flag.UintVar(&c.Port, "P", 3306, "port")
	flag.StringVar(&c.User, "u", "", "user")
	flag.StringVar(&c.Password, "p", "", "password")
	flag.StringVar(&fileName, "f", "", "file")
	flag.Parse()
	c.Database = flag.Arg(0)
	e := []string{flag.Arg(1)}

	cli, err := gomysqlclient.NewCli(&c)
	if err != nil {
		log.Fatal(err)
	}
	if fileName != "" {
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		e = gomysqlclient.QueriesFromReader(f)
	}
	if err := cli.Run(e...); err != nil {
		log.Fatal(err)
	}
}
