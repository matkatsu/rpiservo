package main

import (
	"fmt"
	"net/http"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	master := gobot.NewMaster()
	adaptor := raspi.NewAdaptor()
	servo := gpio.NewServoDriver(adaptor, "12") //PWM0(18)

	work := func() {
		// 13 ~ 41 = 0度 ~ 180度
		servo.Move(uint8(27)) //center
		gobot.After(1*time.Second, func() {
			servo.Move(uint8(35))
			gobot.After(500*time.Millisecond, func() {
				servo.Move(uint8(27))
			})
		})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{adaptor},
		[]gobot.Device{servo},
		work,
	)

	// Starts the API server on default port 3000
	server := api.NewAPI(master)
	server.AddHandler(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	})
	server.Port = "3000"
	server.Start()

	servoEndpoint := master.AddRobot(robot)
	// curl -H 'Content-Type:application/json' 'http://localhost:3000/api/robots/servoBot/commands/move'
	servoEndpoint.AddCommand("move", func(params map[string]interface{}) interface{} {
		reqParam := fmt.Sprintf("Params: %+v\n", params)
		robot.Start()
		return reqParam
	})

	master.Start()
}
