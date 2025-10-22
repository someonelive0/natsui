package main

import (
	"log"
	"os"

	"natsui/backend"

	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

var App = grumble.New(&grumble.Config{
	Name:                  "kafcli",
	Description:           "An CLI of Nats",
	HistoryFile:           "kafcli.hist",
	Prompt:                "kafcli » ",
	PromptColor:           color.New(color.FgGreen, color.Bold),
	HelpHeadlineColor:     color.New(color.FgGreen),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,

	Flags: func(f *grumble.Flags) {
		// f.String("f", "config", "kafui.toml", "set config file")
		f.String("a", "address", "127.0.0.1:9092", "set nats host")
		f.String("u", "user", "DEFAULT", "set nats user")
		f.String("p", "password", "", "set nats password")
		f.Bool("v", "verbose", false, "enable verbose mode")
	},
})

var natstool *backend.NatsTool

func init() {
	App.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println("  Natsui CLI ")
		// a.Println()
	})

	quitCommand := &grumble.Command{
		Name:     "quit",
		Help:     "quit this cli",
		LongHelp: "quit this cli",
		Aliases:  []string{"q"},
		Run: func(c *grumble.Context) error {
			os.Exit(0)
			return nil
		},
	}
	App.AddCommand(quitCommand)
}

func main() {
	myconfig, err := backend.LoadConfig("natsui.toml")
	if err != nil {
		log.Fatalf("LoadConfig failed: %s", err)
	}
	log.Printf("myconfig nats servers %v", myconfig.Nats.Servers)
	natstool = backend.NewNatsTool(&myconfig.Nats)

	grumble.Main(App)
}
