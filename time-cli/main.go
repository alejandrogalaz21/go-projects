package main

import (
	"fmt"
	"time"
)


	func main() {
		now := time.Now()
		fmt.Println("🌍 Hora actual en diferentes zonas:")
		fmt.Println("🕒 UTC:", now.Format("15:04:05"))
		fmt.Println("🗽 New York:", now.In(time.FixedZone("EST", -5*3600)).Format("15:04:05"))
		fmt.Println("🌆 Los Angeles:", now.In(time.FixedZone("PST", -8*3600)).Format("15:04:05"))
		fmt.Println("🇪🇸 Madrid:", now.In(time.FixedZone("CET", 1*3600)).Format("15:04:05"))
	}