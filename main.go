package main

import (
	"fmt"

	razer "github.com/rgnix2/goRazerChromaSDK"
)

func main() {

	// INFO you application information
	// Flieds in order
	// Title, Description. Author.Name, Author.Email,  []devices as string "chromalink", "keyboard", "end points", "etc", then application leave as is.

	// load info
	info := razer.AppInfo{Title: "Stock RGB",
		Description: "This app is for stock moving to rgb colors", Author: razer.Author{Name: "OneNix", Contact: "email@email.com"},
		DeviceSupported: []string{"chromalink", "keyboard", "mouse", "mousepad", "headset", "keypad"},
		Category:        "application"}

	// get session and {start heartbeat}
	razerSession, err := razer.GetSession(info)
	if err != nil {
		print(err)
	}
	fmt.Printf("\nsessionID: %d \nURI: %s\n", razerSession.Sessionid, razerSession.URI)
	go razer.HeartBeat()

	// // load color for static color
	// color := razer.StaticColor{Effect: "CHROMA_STATIC", Param: struct {
	// 	Color int "json:\"color\""
	// }{14444}}

	// // test static mouse
	// resColor, err := razer.MouseStatic(info, color)
	// if err != nil {
	// 	print(err)
	// }
	// fmt.Printf("result: %d ", resColor.Result)

	// // test static kb
	// mouse, err := razer.KeyboardStaticPUT(info, color)
	// if err != nil {
	// 	print(err)
	// }
	// fmt.Printf("result: %d ", mouse.Result)

	// get version
	razerVersion, err := razer.GetVersion()
	if err != nil {
		print(err)
	}

	fmt.Printf("\nversion: %s", razerVersion.Version)

}
