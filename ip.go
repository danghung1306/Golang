package main

import (
	"bufio"
	f "fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const pattern string = "^(\\d+).(\\d+).(\\d+).(\\d+)/(\\d+)"

var IP uint32 = 0
var bitUsed uint8 = 0
var bitUnUsed uint8 = 0
var subNetMask uint32 = 0b11111111111111111111111111111111
var subNetMaskInvert uint32 = 0
var childIp [4]uint32

func main() {

	subIpAdress := inputIP()
	bitUsed = subIpAdress[4]
	bitUnUsed = 32 - bitUsed

	for i := 0; i < 4; i++ {
		IP = (IP << 8) | uint32(subIpAdress[i])
	}
	originIP := IP & (subNetMask << bitUnUsed)

	childIp = saparatedIP(originIP)
	for i := 0; i < 4; i++ {
		f.Printf("\n%dst subnet is: ", i+1)
		convertAddress(childIp[i])
		f.Printf("/%d", bitUsed+2)
	}
}

func convertAddress(ip uint32) {
	ipAddress := [4]uint8{}
	for i := 3; i >= 0; i-- {
		ipAddress[i] = uint8(ip)
		ip = ip >> 8
	}
	for i := 0; i < 4; i++ {
		if i != 3 {
			f.Printf("%d.", ipAddress[i])
		} else {
			f.Printf("%d", ipAddress[i])
		}
	}
	return
}

func saparatedIP(parentIP uint32) (childIp [4]uint32) {
	childIp[0] = parentIP
	childIp[1] = parentIP | ((subNetMaskInvert + 0b01) << (bitUnUsed - 2))
	childIp[2] = parentIP | ((subNetMaskInvert + 0b10) << (bitUnUsed - 2))
	childIp[3] = parentIP | ((subNetMaskInvert + 0b11) << (bitUnUsed - 2))
	return
}

func inputIP() (ipArr [5]uint8) {
	f.Print("\nInput your IP Address: ")
	reader := bufio.NewReader(os.Stdin)

	buffer, _ := reader.ReadString('\n')

	buffer = strings.Replace(buffer, "\n", "", -1)

	if flag, _ := regexp.MatchString(pattern, buffer); flag {
		obj := regexp.MustCompile("(\\d+)")
		str := obj.FindAllString(buffer, -1)
		for i := 0; i < len(str); i++ {
			if len(str[i]) < 4 {
				tmp, _ := strconv.Atoi(str[i])
				if tmp < 0 || tmp > 255 {
					f.Println("This sub IP is not valid, x,y,z,w range into [0:255]")
					os.Exit(0)
				} else {
					ipArr[i] = uint8(tmp)
				}
				if i == 4 {
					if ipArr[i] > 31 || ipArr[i] < 0 {
						f.Println("Subnet mask is not valid, it should be range into [0:32]")
						os.Exit(0)
					} else {
						//do nothing
					}
				}
			} else {
				f.Println("This sub IP is not valid, format xxx.yyy.zzz.www/t")
				os.Exit(0)
			}
		}

	} else {
		f.Println("This IP is not valid, format xxx.yyy.zzz.www/t")
		os.Exit(0)
	}
	return
}
