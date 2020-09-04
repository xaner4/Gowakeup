package gowakeup

import (
	"flag"
	"fmt"

	wol "gitlab.com/xaner4/GoWakeUp/pkg"
)

func CMD() {
	MACAddress := flag.String("mac", "", "MAC address")
	ip := flag.String("ip", "255.255.255.255", "Destination IP address")
	port := flag.Int("port", 9, "Destination port")
	flag.Parse()

	mp, err := wol.CreateMagicPacket(*MACAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = wol.SendMagicPacket(mp, *ip, *port)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Magic packet sent successfully to %q on port %d\n", *ip, *port)

	}
}
