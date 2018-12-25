package main

import (
	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/midn/alexa/vizio/key_command"
	"github.com/midn/alexa/vizio/youtube"
)

type Keylist struct {
	CodeSet int    `json:"CODESET"`
	Code    int    `json:"CODE"`
	Action  string `json:"ACTION"`
}

type Sample struct {
	Keylists []Keylist `json:"KEYLIST"`
}

// DispatchIntents dispatches each intent to the right handler
func DispatchIntents(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "hello":
		response = handleHello()
	case "click":
		response = handleClick()
	case "home":
		response = handleHome()
	case "cartoons":
		response = handleCartoons()
	case "next_video":
		response = handleNextVideo()
	// case "volume":
	// 	response = handleVolume()
	case alexa.HelpIntent:
		response = handleHelp()
	default:
		response = handleHelp()
	}

	return response
}

func handleHello() alexa.Response {
	return alexa.NewSimpleResponse("Saying Hello", "Hello, World")
}

func handleClick() alexa.Response {
	key_command.Click()
	return alexa.NewSimpleResponse("Vizio", "Clicked")
}

func handleHome() alexa.Response {
	key_command.Home()
	return alexa.NewSimpleResponse("Vizio", "Navigated home")
}

func handleCartoons() alexa.Response {
	youtube.OpenAndPlay()
	return alexa.NewSimpleResponse("Vizio", "Playing cartoons")
}

func handleNextVideo() alexa.Response {
	youtube.NextVideo()
	return alexa.NewSimpleResponse("Vizio", "Playing next video")
}

func handleHelp() alexa.Response {
	return alexa.NewSimpleResponse("Help for Vizio", "To receive a greeting, ask Vizio to say hello")
}

// func handleVolume() alexa.Response {
// 	volume.SetVolume(8)
// 	return alexa.NewSimpleResponse("Vizio", "Volume set")
// }

// Handler is the lambda hander
func Handler(request alexa.Request) (alexa.Response, error) {
	return DispatchIntents(request), nil
}

func main() {
	lambda.Start(Handler)
}
