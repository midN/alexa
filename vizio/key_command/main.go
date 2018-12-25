package key_command

import (
	"github.com/midn/alexa/vizio/helpers"
)

type Keylist struct {
	CodeSet int    `json:"CODESET"`
	Code    int    `json:"CODE"`
	Action  string `json:"ACTION"`
}

type RequestData struct {
	Keylists []Keylist `json:"KEYLIST"`
}

func Click() {
	keyPress(3, 2)
}

func Up() {
	keyPress(3, 8)
}

func Down() {
	keyPress(3, 0)
}

func Right() {
	keyPress(3, 7)
}

func Left() {
	keyPress(3, 1)
}

func Home() {
	keyPress(4, 15)
}

func CycleInput() {
	keyPress(7, 1)
}

func keyPress(codeset int, code int) {
	msg := &RequestData{
		Keylists: []Keylist{
			Keylist{
				CodeSet: codeset,
				Code:    code,
				Action:  "KEYPRESS",
			},
		},
	}
	helpers.GenerateRequest("PUT", "key_command/", msg, false)
}
