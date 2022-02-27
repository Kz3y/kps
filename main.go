package main

import (
	"os"
	"runtime"

	"github.com/urfave/cli"

	"PortScan/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "port_scanner"
	app.Author = "k3zy"
	app.Email = "x@qq.com"
	app.Version = "1.2.0"
	app.Usage = "tcp syn/connect port scanner"
	app.Commands = []cli.Command{cmd.Scan}
	app.Flags = append(app.Flags, cmd.Scan.Flags...)
	err := app.Run(os.Args)
	_ = err
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
