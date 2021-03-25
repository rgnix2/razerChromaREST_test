package main

import (
	"fmt"

	razer "github.com/rgnix2/goRazerChromaSDK"
)

// type Author struct {
// 	Name    string `json:"name"`
// 	Contact string `json:"contact"`
// }
// type AppInfo struct {
// 	Title           string `json:"title"`
// 	Description     string `json:"description"`
// 	Author          `json:"author"`
// 	DeviceSupported []string `json:"device_supported"`
// 	Category        string   `json:"category"`
// }
// type StaticColor struct {
// 	Effect string `json:"effect"`
// 	// Param  struct {
// 	// 	Color int `json:"color"`
// 	// } `json:"param"`
// }

func main() {
	// info := razer.AppInfo{"Stock RGB",
	// 	"This app is for stock moving to rgb colors", razer.Author{"OneNix", "email@email.com"},
	// 	[]string{"chromalink", "keyboard"},
	// 	"application"}

	// razerSession, err := razer.GetSession(info)
	// if err != nil {
	// 	print(err)
	// }
	// payload, err := json.Marshal(map[string]interface{}{
	// 	"effect": "CHROMA_NONE",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 2.
	// client := &http.Client{}
	// url := razerSession.URI + "/mousepad"
	// println(url)
	// time.Sleep(1 * time.Second)
	// // 3.
	// req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	// req.Header.Set("Content-Type", "application/json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 4.
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 5.
	// defer resp.Body.Close()

	// // 6.
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(body))

	//////////////
	// links := []string{
	// 	"http://localhost:54235/razer/chromasdk/keyboard",
	// 	"http://localhost:54235/razer/chromasdk",
	// }

	// checkUrls(links)

	// WORKING STUFF MAYBE BELOW!!!
	// Flieds in order
	// Title, Description. Author.Name, Author.Email,  []devices as string "chromalink", "keyboard", "etc", then application leave as is.
	// will fix this warning later...
	info := razer.AppInfo{"Stock RGB",
		"This app is for stock moving to rgb colors", razer.Author{"OneNix", "email@email.com"},
		[]string{"chromalink", "keyboard"},
		"application"}

	razerSession, err := razer.GetSession(info)
	if err != nil {
		print(err)
	}
	fmt.Printf("\nsessionID: %d \nURI: %s\n", razerSession.Sessionid, razerSession.URI)

	color := razer.StaticColor{Effect: "CHROMA_STATIC", Param: struct {
		Color int "json:\"color\""
	}{14444}}
	//razer.KeyboardStatic(info, color)
	//razer.MouseStatic(info, color)
	resColor, err := razer.MouseStatic(info, color)
	if err != nil {
		print(err)
	}
	fmt.Printf("result: %d ", resColor.Result)

	mouse, err := razer.KeyboardStaticPUT(info, color)
	if err != nil {
		print(err)
	}
	fmt.Printf("result: %d ", mouse.Result)

	razerVersion, err := razer.GetVersion()
	if err != nil {
		print(err)
	}

	fmt.Printf("\nversion: %s", razerVersion.Version)

	//println(blah, err)

}
