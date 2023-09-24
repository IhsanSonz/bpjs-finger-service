package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gin-gonic/gin"
)

var serviceStatus = "Bruh is the service running or not?\nHello from github.com/IhsanSonz"

func main() {
	r := gin.Default()

	// Define an API endpoint to identify that the service is running
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, serviceStatus)
	})

	// // Define an API endpoint to activate BPJS fingerprint app
	// r.GET("/activate-bpjs", func(c *gin.Context) {
	// 	// Activate the BPJS fingerprint window by its title
	// 	result := robotgo.ActiveName("After.exe")
	// 	// result := robotgo.ActiveName("Notepad")

	// 	if result == nil {
	// 		c.String(http.StatusOK, "BPJS Fingerprint activated")
	// 	} else {
	// 		c.String(http.StatusNotFound, "BPJS Fingerprint activation failed")
	// 	}
	// })

	r.GET("/activate-bpjs", func(c *gin.Context) {
		// Replace "path/to/your/exefile.exe" with the actual path to your executable.
		exePath := "./BPJSForeground.exe"

		// Create a new command to run the executable.
		cmd := exec.Command(exePath)

		// Set the standard output and error to os.Stdout and os.Stderr, respectively.
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the command.
		if err := cmd.Run(); err != nil {
			// fmt.Printf("Error running the executable: %v\n", err)
			c.String(http.StatusNotFound, "BPJS Fingerprint activation failed")
			return
		}

		// fmt.Println("Executable ran successfully.")
		c.String(http.StatusOK, "BPJS Fingerprint activated")
	})

	// Start the API server in a separate Goroutine
	go func() {
		err := r.Run(":8089") // Change the port as needed
		if err != nil {
			fmt.Println("Error starting the API server:", err)
		} else {
			fmt.Println("Service running in port 8089")
		}
	}()

	// Create a Fyne GUI
	myApp := app.New()
	myWindow := myApp.NewWindow("Service Status")
	myLabel := widget.NewLabel(serviceStatus)

	// Create a container for the GUI elements
	content := container.NewVBox(
		myLabel,
	)

	myWindow.SetContent(content)
	myWindow.SetFixedSize(true)
	myWindow.Hide()
	// myWindow.SetAlwaysOnTop(true)

	myWindow.ShowAndRun()
}
