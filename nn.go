package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli"
)

type record struct {
	lastFired int64
	number    int
}

func getRecord(configDir, id string) *record {
	b, err := ioutil.ReadFile(filepath.Join(configDir, id+".txt"))
	r := &record{}
	if err != nil {
		return r
	}
	fmt.Sscanf(string(b), "{%d,%d}", &r.lastFired, &r.number)
	return r
}

func calcDateDelta(last, now int64) int64 {
	delta := now - last
	return delta
}

func updateRecord(id string, now int64, num int) error {
	str := fmt.Sprintf("{%d,%d}", now, num)
	err := ioutil.WriteFile(filepath.Join(getConfigDir(), id+".txt"), []byte(str), 0644)
	if err != nil {
		return err
	}
	return nil
}

func getConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "nn")
}

func showDelta(delta int64, num int, format string) error {
	fmt.Printf(format, delta/86400, num)
	return nil
}

func defaultCommand(c *cli.Context) {
	err := os.MkdirAll(getConfigDir(), 0755)
	if err != nil {
		log.Printf("failed create config directory. err: %s", err.Error())
		os.Exit(1)
	}

	id := c.GlobalString("id")
	format := c.GlobalString("format")

	now := time.Now().Unix()
	last := getRecord(getConfigDir(), id)
	var delta int64
	if last.lastFired != 0 {
		delta = calcDateDelta(last.lastFired, now)
	}

	err = showDelta(delta, last.number+1, format)
	if err != nil {
		fmt.Println("failed to show. confirm format validity")
		os.Exit(1)
	}

	err = updateRecord(id, now, last.number+1)
	if err != nil {
		fmt.Println("failed to update record. err: %s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

const (
	appName    = "nn"
	appUsage   = "show nn"
	appVersion = "0.1"
)

var globalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "id",
		Value: "default",
		Usage: "specify id of nn",
	},
	cli.StringFlag{
		Name:  "f, format",
		Value: "%d 日ぶり %d 回目",
		Usage: "specify format of nn",
	},
}

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.Flags = globalFlags
	app.Action = defaultCommand
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [options] [arguments...]

VERSION:
   {{.Version}}{{if or .Author .Email}}

AUTHOR:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`
	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
