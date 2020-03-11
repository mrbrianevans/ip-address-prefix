package main

import (
	"fmt"
	"strings"
)

func preIP(ip string) string {
	prefixedIP := ip // default return is the input
	if strings.Count(ip, ".")==3 { // ensures there are 3 dots in the IP address before running
		for i:=len(ip)-1 ; i>=0 ; i-- {
			if ip[i] == 46 { // 46 is the ascii code for .  (dot)
				prefixedIP = ip[0:i]
				break // exit the loop once the first dot is found
			}
		}
	}
	return prefixedIP
}

func main(){
	var ips [5] string = [5]string{"69.160.160.133", "130.43.180.51", "92.40.200.148", "5.151.93.193", "34.240.50.53"}

	fmt.Println(ips)

	for i:=0; i<5; i++{
		fmt.Println(preIP(ips[i]))
	}
}


