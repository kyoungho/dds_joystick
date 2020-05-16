package main

import (
	"github.com/rticommunity/rticonnextdds-connector-go"
	"github.com/kyoungho/dds_joystick/types"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"log"
	"path"
	"runtime"
)

func main() {
	// Find the file path to the XML configuration
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Panic("runtime.Caller error")
	}
	filepath := path.Join(path.Dir(filename), "../Joystick.xml")

	// Create a connector defined in the XML configuration
	connector, err := rti.NewConnector("MyParticipantLibrary::Zero", filepath)
	if err != nil {
		log.Panic(err)
	}
	// Delete the connector when this main function returns
	defer connector.Delete()

	// Get an input from the connector
	input, err := connector.GetInput("MySubscriber::JoystickReader")
	if err != nil {
		log.Panic(err)
	}

	raspiAdaptor := raspi.NewAdaptor()
	gpg3 := gopigo3.NewDriver(raspiAdaptor)

	work := func() {
	// Get DDS samples
	for {
		connector.Wait(-1)
		input.Take()
		numOfSamples := input.Samples.GetLength()
		for j := 0; j < numOfSamples; j++ {
			if input.Infos.IsValid(j) {
				var js types.Joystick
				err := input.Samples.Get(j, &js)
				if err != nil {
					log.Println(err)
				}
				switch js.Button{
				case types.JS_LEFT:
					gpg3.SetMotorDps(gopigo3.MOTOR_RIGHT, 1000)
					gpg3.SetMotorDps(gopigo3.MOTOR_LEFT, 0)
					log.Println("left_press")
				case types.JS_RIGHT:
					gpg3.SetMotorDps(gopigo3.MOTOR_LEFT, 1000)
					gpg3.SetMotorDps(gopigo3.MOTOR_RIGHT, 0)
					log.Println("right_press")
				case types.JS_UP:
					gpg3.SetMotorDps(gopigo3.MOTOR_LEFT + gopigo3.MOTOR_RIGHT, 1000)
					log.Println("up_press")
				case types.JS_DOWN:
					gpg3.SetMotorDps(gopigo3.MOTOR_LEFT + gopigo3.MOTOR_RIGHT, -1000)
					log.Println("down_press")
				case types.JS_START:
					gpg3.SetMotorDps(gopigo3.MOTOR_LEFT + gopigo3.MOTOR_RIGHT, 0)
					log.Println("start_press")
				}
			}
		}
	}

	}

	robot := gobot.NewRobot("gopigo3",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gpg3},
		work,
	)

	robot.Start()
}
