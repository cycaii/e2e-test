package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

const (
	SendMarketAdMsg   = "MarketAd"
	SendCCHMappingMsg = "CCH"
	SendJitMsg        = "JIT"
)

func TestAction(cli *cli.Context) error {
	num := cli.Int("num")
	name := cli.String("name")
	fmt.Println(num, name)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "send",
			Usage:  "send kafka message, name: MarketAd/CCH/JIT, number: int, total msg count",
			Action: SendAction,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "num",
					//Value: 1,
					Usage: "--num, total message count send to kafka.",
				},
				cli.StringFlag{
					Name: "name",
					//Value: SendMarketAdMsg,
					Usage: "--name, kafkaTopic.",
				},
			},
		},
	}

	fmt.Println("args: %v", os.Args)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func SendAction(c *cli.Context) error {
	if !c.IsSet("num") || !c.IsSet("name") {
		return errors.New("either num or name flag must be provided")
	}

	total := c.Int("num")
	typeName := c.String("name")
	fmt.Println("num:", total, ",name:", typeName, c.IsSet("num"), c.IsSet("name"))

	if typeName == "" {
		return errors.New("name flag cannot be empty")
	}

	if err := sendKafkaMsg(total, typeName); err != nil {
		return err
	}
	fmt.Println("Send kafka message successfully.")
	return nil
}

func sendKafkaMsg(total int, typeName string) error {
	fmt.Println("sendKafkaMsg.")

	switch typeName {
	case SendMarketAdMsg:
		ProduceMarketAdMsg(total)
	case SendCCHMappingMsg:
		ProduceCCHMappingMsg(total)
	case SendJitMsg:
		ProduceJittMsf(total)
	default:
		return nil
	}
	return nil
}
