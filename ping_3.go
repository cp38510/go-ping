package main

import (
	"flag"
	"fmt"
	"github.com/sparrc/go-ping"
	"net"
)


func ff(from string) {

	start := flag.String("s", "172.20.13.1", "start IP")
	end := flag.String("e", "172.20.13.224", "end IP")



	flag.Parse()

	ip := net.ParseIP(*start)
	ip2 := net.ParseIP(*end)

	ip = ip.To4()
	ip2 = ip2.To4()

//	ip[3]++
//	ip2[3]++

//	fmt.Println(ip)
//	fmt.Println(ip2)

//	fmt.Println(ip2[3]-ip[3])


	for i := ip[3]; i <= ip2[3]; i++ {
	//	host := "172.20.13."
	//	host := ip[1]
	//	host := (*start)[:9]
		test := ip[0]
		test1 := ip[1]
		test2 := ip[2]
	//	output := fmt.Sprintf("%v.%d",ip[:3], i)
	//	fmt.Println(output)
		output := fmt.Sprintf("%d.%d.%d.%d",test,test1,test2, i)
	//	fmt.Println(output)


		pinger, err := ping.NewPinger(output)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		pinger.OnRecv = func(pkt *ping.Packet) {
			fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
				pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
		}

		fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
		pinger.Count = 3
		pinger.Timeout = 10000000000
		pinger.Run()
	}

}

func main() {

//	    for i := 0; i < 2; i++ {
	     go ff("goroutine")
//	    }

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

