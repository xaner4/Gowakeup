package wol

import (
	"encoding/hex"
	"fmt"
	"net"
	"regexp"
	"strings"
)

var syncStream string = "FFFFFFFFFFFF"

// CreateMagicPacket makes a magic packet that can be used later
func CreateMagicPacket(MACAddress string) ([]byte, error) {
	reMAC := regexp.MustCompile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$")
	MAC := reMAC.Find([]byte(MACAddress))

	if MAC == nil {
		return nil, fmt.Errorf("%q is not a valid MAC address", MACAddress)
	}

	MACAddress = strings.ReplaceAll(string(MAC), ":", "")

	tMAC := strings.Repeat(MACAddress, 16)
	magicPacket, err := hex.DecodeString(syncStream + tMAC)
	if err != nil {
		return nil, err
	}
	return magicPacket, nil
}

// SendMagicPacket sends a magic packet to the network to wake up a computer
func SendMagicPacket(magicPacket []byte, addr string, port int) (status int, err error) {
	adress := fmt.Sprintf("%s:%d", addr, port)
	conn, err := net.Dial("udp", adress)
	defer conn.Close()
	if err != nil {
		return 0, err
	}
	status, err = conn.Write(magicPacket)
	if err != nil {
		return 0, err
	}

	return status, nil
}
