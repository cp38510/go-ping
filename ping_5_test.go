package main

import (
//	"bytes"
	"flag"
	"github.com/sparrc/go-ping"
	"net"
	"fmt"
//	"sort"
	"strconv"
	"time"
//	"math/rand"
)

var usage = `
Usage:
    pping [-s start IP] [-e end IP]

Example:
    # pping -s 172.20.13.1 -e 172.20.13.225
`

//ip1 := make(chan int)
//ip21 := make(chan int)

func worker(id int, jobs <-chan string, results chan<- string) {
//	out2 := make(chan string)
	for j := range jobs {

		pinger, err := ping.NewPinger(j)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		out2 := make(chan string)
//		out3 := <- out2
//		fmt.Print(out3)
		go func() {
			out3 := <-out2
			fmt.Print(out3)
		}()

		pinger.OnFinish = func(stats *ping.Statistics) {

			var0 := fmt.Sprintf("%v", stats.PacketLoss)
			var10, err := strconv.ParseFloat(var0, 64)
			if err == nil {
			}

			var2 := float64(0)
			var3 := float64(100)


//			ip11 := <- ip1
//			ip211 := <- ip21


			if var10 == var2 {
				//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST ACTIVE")
				//				fmt.Printf("%s %s\n", pinger.Addr(), "- HOST ACTIVE")
				out0 := fmt.Sprintf ("%s %s\n", pinger.Addr(), "- HOST ACTIVE")
//				out0 := fmt.Sprintf("\"%s\",\n", pinger.Addr())
				out2 <- out0
			}else if var10 < var3 {
				//				fmt.Printf("%s %s, %v%% packet loss\n", pinger.Addr(), "- HOST WARNING", stats.PacketLoss)
				out0 := fmt.Sprintf ("%s %s, %v%% packet loss\n", pinger.Addr(), "- HOST WARNING", stats.PacketLoss)
				out2 <- out0
			}else if var10 == var3 {
				//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST DISABLED")
				//				fmt.Printf("%s %s\n", pinger.Addr(), "- HOST DISABLED")
				out0 := fmt.Sprintf ("%s %s\n", pinger.Addr(), "- HOST DISABLED")
   			    out2 <- out0
			}
//			out6 := <- out2
//			out5 := fmt.Sprintf("%s", out6)
//			fmt.Print(out5)

		}

		//		timeout := flag.Duration("t", 10000000000, "")
		//		count := flag.Int("c", 3, "")

		//		pinger.Count = *count
		//		pinger.Timeout = *timeout
		pinger.Count = 3
//		pinger.Timeout = 10000000000
		pinger.Timeout = 10000000000




//		ololo5 := make(chan []string)
/*
		go func() {
			out3 := <- out2
//			ololo1 := fmt.Sprint(out3)
//			fmt.Println("TEST")
//			time.Sleep(time.Second * 2)
//			amt := time.Duration(rand.Intn(250))
//			time.Sleep(time.Millisecond * amt)
			fmt.Print(out3)
//			ololo := []rune(ololo1)
//			ololo2 := []string{out3}
//			ololo := []string{ololo1}
//			ololo <- ololo1

		realIPs := make([]net.IP, 0, len(ololo2))

			for _, ip := range ololo2 {
				realIPs = append(realIPs, net.ParseIP(ip))
			}

			sort.Slice(realIPs, func(i, j int) bool {
				return bytes.Compare(realIPs[i], realIPs[j]) < 0
			})

			for _, ip := range realIPs {
				fmt.Printf("%s\n", ip)
			}


//			fmt.Print(ololo2)
//			out10 := fmt.Sprintf("[%s]", out3)
//			fmt.Print(out3)
		}()
*/

//		ololo5 := <- ololo
//		fmt.Print(ololo5)




		pinger.Run()

//		time.Sleep(time.Second * 10)
		results <- j
	}
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
//	fmt.Println(out2)


	for j := 1; j <= 1; j++ {

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


	for a := 1; a <= 1; a++ {
		<-results
	}


}

func main(){

	go f1()
	time.Sleep(time.Second * 15)
	return

}
