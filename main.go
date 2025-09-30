package main

import (
	"embed"
	_ "embed"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v3/pkg/application"

	"natsui/backend"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {

	if err := backend.InitLogRotate("./", "natsui.log", "info", 3, 10); err != nil {
		fmt.Printf("init log failed: %v\n", err)
	}

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "natsui",
		Description: "Nats GUI tool",
		Services: []application.Service{
			application.NewService(&GreetService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})
	log.Infof("init app")

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "NatsUI", // Title: "Window 1",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour:    application.NewRGB(27, 38, 54),
		URL:                 "/",
		Width:               1024,
		Height:              768,
		MinimiseButtonState: application.ButtonEnabled, // ButtonHidden, Setting Button States During Window Creation
		MaximiseButtonState: application.ButtonEnabled, // ButtonDisabled,
		CloseButtonState:    application.ButtonEnabled,
		// Windows: application.WindowsWindow{ // Controlling Window Style (Windows)
		// 	ExStyle: w32.WS_EX_CLIENTEDGE | w32.WS_EX_COMPOSITED | w32.WS_EX_TRANSPARENT,
		// },
	})
	log.Infof("init app main window")

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	// Run the application. This blocks until the application has been exited.
	log.Infof("app running...")
	err := app.Run()
	log.Infof("app end...")

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
