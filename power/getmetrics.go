package power

import (
	"fmt"

	battery "github.com/distatus/battery"
)

func GetBatteryInfo() {
	batteries, _ := battery.GetAll()
	for _, bat := range batteries {
		fmt.Println("state: ", bat.State)
		fmt.Printf("\nBattery: %.2f%%\n", bat.Current/bat.Full*100)
		fmt.Printf("health: %.2f%%\n", bat.Full/bat.Design*100)
		fmt.Printf("voltag: %.2f%%\n", bat.Voltage/bat.DesignVoltage*100)

	}
}
