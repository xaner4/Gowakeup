package wol

import (
	"encoding/hex"
	"fmt"
	"net"
	"regexp"
	"strings"
)

// CreateMagicPacket makes a magic packet that can be used later
func CreateMagicPacket(MACAddress string) ([]byte, error) {
	// syncStream is the start of the magic packet
	syncStream := "FFFFFFFFFFFF"
	// delimiters that can exsist in a MAC address
	delimiter := []string{":", "-"}

	// reMAC checks if the MAC address passed in the argument is a valid MAC address
	reMAC := regexp.MustCompile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$")
	MAC := reMAC.Find([]byte(MACAddress))

	// If the MAC address is not valid, return an error
	if MAC == nil {
		return nil, fmt.Errorf("%q is not a valid MAC address", MACAddress)
	}

	// Copies the Byte string of MAC found in reqex to the MAC address Variable
	// loop over to remove any delimiters
	MACAddress = string(MAC)
	for _, v := range delimiter {
		MACAddress = strings.ReplaceAll(MACAddress, v, "")
	}

	// Repeat the mac address 16 times
	tMAC := strings.Repeat(MACAddress, 16)

	// Decodes the string to HEX
	magicPacket, err := hex.DecodeString(syncStream + tMAC)
	if err != nil {
		return nil, err
	}
	return magicPacket, nil
}

// SendMagicPacket sends a magic packet to the network to wake up a computer
func SendMagicPacket(mp []byte, addr string, port int) (err error) {
	// Checks if IP is valid
	ip := net.ParseIP(addr)
	if ip == nil {
		return fmt.Errorf("%q is not a valid IP address", addr)
	}
	// Sets the address and port to on connection string
	adress := fmt.Sprintf("%s:%d", ip, port)

	// Creates the connection
	conn, err := net.Dial("udp", adress)
	defer conn.Close()
	if err != nil {
		return err
	}

	// Sends the packet to teh connection
	_, err = conn.Write(mp)
	if err != nil {
		return err
	}

	return nil
}
