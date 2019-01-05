package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
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

	robot.Start()
}
