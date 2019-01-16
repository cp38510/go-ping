package main

import (
	"flag"
	"github.com/sparrc/go-ping"
	"net"
	"fmt"
	"strconv"
//	"strings"
	"time"
)

var usage = `
Usage:
    pping [-s start IP] [-e end IP]

Example:
    # pping -s 172.20.13.1 -e 172.20.13.225
`

func worker(id int, jobs <-chan string, results chan<- string) {
	for j := range jobs {

		pinger, err := ping.NewPinger(j)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}


		pinger.OnFinish = func(stats *ping.Statistics) {

			var0 := fmt.Sprintf("%v", stats.PacketLoss)
			var10, err := strconv.ParseFloat(var0, 64)
			if err == nil {
			}

			var2 := float64(0)
			var3 := float64(100)

			out2 := make(chan string)

			go func() {
//			fmt.Println ("TEST")
			out3 := <-out2
			fmt.Print(out3)

			}()

			if var10 == var2 {
//				out0 := fmt.Sprintf ("%s %s\n", pinger.Addr(), "- HOST ACTIVE")
				out0 := fmt.Sprintf ("%s\n",pinger.Addr())
				out2 <- out0
			}else if var10 < var3 {
//				out0 := fmt.Sprintf ("%s %s, %v%% packet loss\n", pinger.Addr(), "- HOST WARNING", stats.PacketLoss)
//				out2 <- out0
			}else if var10 == var3 {
//				out0 := fmt.Sprintf ("%s %s\n", pinger.Addr(), "- HOST DISABLED")
//				out2 <- out0
			}

/*			if var10 == var2 {
				//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST ACTIVE")
				fmt.Printf("%s %s\n", pinger.Addr(), "- HOST ACTIVE")
			}else if var10 < var3 {
				fmt.Printf("%s %s, %v%% packet loss\n", pinger.Addr(), "- HOST WARNING", stats.PacketLoss)
			}else if var10 == var3 {
				//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST DISABLED")
				fmt.Printf("%s %s\n", pinger.Addr(), "- HOST DISABLED")
			}
*/
		}

		pinger.Count = 3
		pinger.Timeout = 10000000000


		pinger.Run()

		time.Sleep(time.Second * 10)
		results <- j
	}



	}

func f2() {

	func f1() {
		jobs := make(chan string)
		results := make(chan string)
		//	out2 := make(chan string)

		//	start := flag.String("s", "172.20.13.1", "start IP")
		//	end := flag.String("e", "172.20.13.224", "end IP")

		start := flag.String("s", "8.8.8.8", "start IP")
		end := flag.String("e", "8.8.8.16", "end IP")

		hhelp := flag.String("h", "", "")
		hhelp1 := flag.String("help", "", "")

		flag.Usage = func() {
			fmt.Printf(usage)
		}

		flag.Parse()

		if *hhelp == "h" {
			flag.Usage()
			return
		}

		if *hhelp1 == "help" {
			flag.Usage()
			return
		}

		ip := net.ParseIP(*start)
		ip2 := net.ParseIP(*end)

		ip = ip.To4()
		ip2 = ip2.To4()

		ips := int(ip2[3] - ip[3])

		if ip[0] != ip2[0] {
			fmt.Printf("Error! Start and end IP's have different subnets\n")
			return
		} else if ip[1] != ip2[1] {
			fmt.Printf("Error! Start and end IP's have different subnets\n")
			return
		} else if ip[2] != ip2[2] {
			fmt.Printf("Error! Start and end IP's have different subnets\n")
			return
		}

		for w := 0; w <= ips; w++ {
			//		go worker(w, jobs, results)
			go worker(w, jobs, results)
		}

		for i := ip[3]; i <= ip2[3]; i++ {

			test := ip[0]
			test1 := ip[1]
			test2 := ip[2]

			output := fmt.Sprintf("%d.%d.%d.%d", test, test1, test2, i)
			jobs <- output
		}

		//	close(jobs)
		//return

	}
go f1()
}


func main(){
//out5 := make(chan string)
//	var ret string
//	go func() {
//		ret = f1()
//	}()
go f2()
//go f1()
//	ret = f1()
//	out5 := f1()
//	fmt.Println(f1())
//f1()
	time.Sleep(time.Second * 12)
	return

}

