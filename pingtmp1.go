package main

import (
	"flag"
	"github.com/sparrc/go-ping"
	"net"
	"fmt"
	"strconv"
	"time"
)

var usage = `
Usage:
    pping [-s start IP] [-e end IP]

Example:
    # pping -s 172.20.13.1 -e 172.20.13.225
`



func worker(id int, jobs <-chan string, results chan<- string) {
//	out2 := make(chan string)
	for j := range jobs {

		pinger, err := ping.NewPinger(j)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		out2 := make(chan string)

		pinger.OnFinish = func(stats *ping.Statistics) {

			var0 := fmt.Sprintf("%v", stats.PacketLoss)
			var10, err := strconv.ParseFloat(var0, 64)
			if err == nil {
			}

			var2 := float64(0)
			var3 := float64(100)

//			out2 := make(chan string)
//			out3 := <-out2
//			fmt.Println(out3)

			if var10 == var2 {
				//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST ACTIVE")
//				fmt.Printf("%s %s\n", pinger.Addr(), "- HOST ACTIVE")
				out0 := fmt.Sprintf ("%s %s\n", pinger.Addr(), "- HOST ACTIVE")
//				out0 := fmt.Sprint("lol1")
				out2 <- out0
			}else if var10 < var3 {
//				fmt.Printf("%s %s, %v%% packet loss\n", pinger.Addr(), "- HOST WARNING", stats.PacketLoss)
				out0 := fmt.Sprintf ("%s %s, %v%% packet loss\n", pinger.Addr(), "- HOST WARNING", stats.PacketLoss)
//				out0 := fmt.Sprint("lol2")
				out2 <- out0
			}else if var10 == var3 {
				//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST DISABLED")
//				fmt.Printf("%s %s\n", pinger.Addr(), "- HOST DISABLED")
				out0 := fmt.Sprintf ("%s %s\n", pinger.Addr(), "- HOST DISABLED")
//				out0 := fmt.Sprint("lol3")
				out2 <- out0
			}


		}

		//		timeout := flag.Duration("t", 10000000000, "")
		//		count := flag.Int("c", 3, "")

		//		pinger.Count = *count
		//		pinger.Timeout = *timeout
		pinger.Count = 3
		pinger.Timeout = 10000000000


		go func() {
			out3 := <- out2
			time.Sleep(time.Second * 10)
			fmt.Println(out3)
		}()

		pinger.Run()


//		out3 := <- out2


//		time.Sleep(time.Second * 10)
		results <- j
	}
//	fmt.Println(out2)
}


func f1() {
	jobs := make(chan string)
	results := make(chan string)


	start := flag.String("s", "172.20.13.1", "start IP")
	end := flag.String("e", "172.20.13.224", "end IP")

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

	ips := int(ip2[3]-ip[3])


	if ip[0] != ip2[0]{
		fmt.Printf("Error! Start and end IP's have different subnets\n")
		return
	}else if ip[1] != ip2[1]{
		fmt.Printf("Error! Start and end IP's have different subnets\n")
		return
	}else if ip[2] != ip2[2]{
		fmt.Printf("Error! Start and end IP's have different subnets\n")
		return
	}


	//ips2 := ip2.To4()[3]-ip.To4()[3]
	//i5, _ := strconv.Atoi(ips)
	//byte5 := byte(i5)

	//ips2 := int(ips)
	//fmt.Println(ips2)


	for w := 0; w <= ips; w++ {
		go worker(w, jobs, results)
	}


	for j := 1; j <= 5; j++ {

		/*
		start := flag.String("s", "172.20.13.1", "start IP")
		end := flag.String("e", "172.20.13.224", "end IP")


		flag.Parse()

		ip := net.ParseIP(*start)
		ip2 := net.ParseIP(*end)

		ip = ip.To4()
		ip2 = ip2.To4()

		*/

		//ips := (ip2[3]-ip[3])
		//fmt.Println(ips)

		//ips1 <- ips

		for i := ip[3]; i <= ip2[3]; i++ {

			test := ip[0]
			test1 := ip[1]
			test2 := ip[2]

			output := fmt.Sprintf("%d.%d.%d.%d", test, test1, test2, i)
			jobs <- output
		}

	}

	close(jobs)


	for a := 1; a <= 5; a++ {
		<-results
	}

}

func main(){
//	var out2 chan string = make(chan string)

	go f1()
//	out3 := out2
//	fmt.Println(out2)
	time.Sleep(time.Second * 25)
	return

}
