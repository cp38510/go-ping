//EXAMPLE: go run gor10.go 172.20.13.

package main

import "fmt"
import "flag"
import "github.com/sparrc/go-ping"
import "time"
import "math/rand"

func ff(from string) {
        for i := 1; i < 224; i++ {

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	host := flag.Arg(0)
        output := fmt.Sprintf("%s%d",host, i)
	pinger, err := ping.NewPinger(output)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	//pinger, err := ping.NewPinger("mail.ru")
	//pinger, err := ping.NewPinger("66.66.66.66")
	//if err != nil {
	//	fmt.Printf("ERROR: %s\n", err.Error())
	//	return
	//}

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
//	pinger.OnFinish = func(stats *ping.Statistics) {
//		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
//		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
//			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
//		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
//			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
//	}

        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Count = 3
	pinger.Timeout = 10000000000
	pinger.Run()

}
}

func main() {
//        go ff("goroutine")

    for i := 0; i < 4; i++ {
        go ff("goroutine")
    }

        var input string
        fmt.Scanln(&input)
        fmt.Println("done")
}

