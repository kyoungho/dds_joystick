package main

import (
	"github.com/rticommunity/rticonnextdds-connector-go"
	"github.com/rticommunity/rticonnextdds-connector-go/types"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/joystick"
	"log"
	"path"
	"runtime"
	"fmt"
)

func main() {
	// Find the file path to the XML configuration
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Panic("runtime.Caller error")
	}
	filepath := path.Join(path.Dir(filename), "./Joystick.xml")

	// Create a connector defined in the XML configuration
	connector, err := rti.NewConnector("MyParticipantLibrary::Zero", filepath)
	if err != nil {
		log.Panic(err)
	}
	// Delete the connector when this main function returns
	defer connector.Delete()

	// Get an output from the connector
	output, err := connector.GetOutput("MyPublisher::JoystickWriter")
	if err != nil {
		log.Panic(err)
	}
	var js Joystick

	joystickAdaptor := joystick.NewAdaptor()
	stick := joystick.NewDriver(joystickAdaptor, joystick.Dualshock3)

	work := func() {
		stick.On(joystick.LeftPress, func(data interface{}) {
			fmt.Println("left_press")
			js.Button = JS_LEFT
			output.Instance.Set(&js)
			output.Write()
		})
		stick.On(joystick.RightPress, func(data interface{}) {
			fmt.Println("right_press")
			js.Button = JS_RIGHT
			output.Instance.Set(&js)
			output.Write()
		})
		stick.On(joystick.UpPress, func(data interface{}) {
			fmt.Println("up_press")
			js.Button = JS_UP
			output.Instance.Set(&js)
			output.Write()
		})
		stick.On(joystick.DownPress, func(data interface{}) {
			fmt.Println("down_press")
			js.Button = JS_DOWN
			output.Instance.Set(&js)
			output.Write()
		})
		stick.On(joystick.StartPress, func(data interface{}) {
			fmt.Println("start_press")
			js.Button = JS_START
			output.Instance.Set(&js)
			output.Write()
		})
	}

	robot := gobot.NewRobot("joystickBot",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{stick},
		work,
	)

	robot.Start()
}
