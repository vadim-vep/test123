package main

import (
	webview "github.com/webview/webview_go"
	//"github.com/zserge/webview"
	// "github.com/webview/webview"
	"log"
)

func main() {
	log.Println("webview0 Start")
	// Create a new webview window
	w := webview.New(true) // Enable debugging for easier development
	defer w.Destroy()

	// Load the HTML content
	w.Navigate("https://www.example.com")

	// Handle window closing gracefully
	w.SetTitle("My WebView Window")
	w.SetSize(800, 600, webview.HintNone)
	w.Run()
	log.Println("webview0 End")
}
