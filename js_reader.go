/*****************************************************************************
*   (c) 2005-2015 Copyright, Real-Time Innovations.  All rights reserved.    *
*                                                                            *
* No duplications, whole or partial, manual or electronic, may be made       *
* without express written permission.  Any such copies, or revisions thereof,*
* must display this notice unaltered.                                        *
* This code contains trade secrets of Real-Time Innovations, Inc.            *
*                                                                            *
*****************************************************************************/

package main

import (
	"github.com/rticommunity/rticonnextdds-connector-go"
	"github.com/rticommunity/rticonnextdds-connector-go/types"
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
	filepath := path.Join(path.Dir(filename), "./Joystick.xml")

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
					log.Println("left_press")
				case types.JS_RIGHT:
					log.Println("right_press")
				case types.JS_UP:
					log.Println("up_press")
				case types.JS_DOWN:
					log.Println("down_press")
				}
			}
		}

	}
}
