package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/sparrc/go-ping"
	"math/rand"
	"os"

	//	"sort"
	//	"sync"

	//	"sort"
	//	"sync"

	//	"go/types"
	"net"
	"strconv"
	"time"
	"bufio"
)

var usage = `
Usage:
    pping [-s start IP] [-e end IP]

Example:
    # pping -s 172.20.13.1 -e 172.20.13.225
`



func main(){
	//	var res
	//	res <- f1()
	//f1()
	//	out2 := make(chan string)
	out2 := make(chan string)
	//	out2 := make([]string, 7)
	//	out2 := make(chan struct{})
	//	out2 := []chan string{make(chan string)}

	go func() {
		//func f1() (chan string){

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

		//	out2 := make(chan string)
		//	out4 := make(chan string)
		//	out5 := make(chan string)
		//	out3 := make(chan string)

		for w := 0; w <= ips; w++ {

			go func(id int, jobs <-chan string, results chan<- string) {
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

						if var10 == var2 {
							//				out0 := fmt.Sprintf ("%s %s\n", pinger.Addr(), "- HOST ACTIVE")
							out0 := fmt.Sprintf("%s", pinger.Addr())
							out2 <- out0
						} else if var10 < var3 {
							out0 := fmt.Sprintf("%s %s, %v%% packet loss\n", pinger.Addr(), "- HOST WARNING", stats.PacketLoss)
							out2 <- out0
						} else if var10 == var3 {
							out0 := fmt.Sprintf("%s %s\n", pinger.Addr(), "- HOST DISABLED")
							out2 <- out0
						}

					}

					pinger.Count = 3
					pinger.Timeout = 10000000000


					/*
												go func() string{
															//			fmt.Println ("TEST")
															out3 := <-out2
															out4 := fmt.Sprint(out3)
															fmt.Println(out4)
															return out4

														}()
					*/

					go func() {

						out3 := <-out2
						//						out4 := fmt.Sprint(out3)
//						fmt.Print(out3)

						file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
						if err != nil {
							fmt.Println("File does not exists or cannot be created")
							os.Exit(1)
						}
						defer file.Close()

						w := bufio.NewWriter(file)
						r := rand.New(rand.NewSource(time.Now().UnixNano()))
						i := r.Perm(5)
						fmt.Fprintf(w, "%v\n", out3)

						w.Flush()


						/*
												contents := []string{out3}
												sampleChan := make(chan string, 1)
												var sampleList []string
												var wg sync.WaitGroup
												for _, line := range contents {
													wg.Add(1)

													go func(line string, ch chan string, wg *sync.WaitGroup) {
														defer wg.Done()
														ch <- line

													}(line, sampleChan, &wg)
												}
												go func (channel chan string, sampleList *[]string) {
													for s := range channel {
														*sampleList = append(*sampleList, s)
													}

												}(sampleChan, &sampleList)
												wg.Wait()
												sort.Strings(sampleList)
												fmt.Println(sampleList)
						*/
						//						return sampleList
					}()


					pinger.Run()

					//				time.Sleep(time.Second * 10)
					results <- j

				}


			}(w, jobs, results)


		}


		for i := ip[3]; i <= ip2[3]; i++ {

			test := ip[0]
			test1 := ip[1]
			test2 := ip[2]

			output := fmt.Sprintf("%d.%d.%d.%d", test, test1, test2, i)
			jobs <- output
		}


		//out5 := <- out4
		//	return out2

	}()



	//	fmt.Println(<-out2)
	/*
		for v := range out2 {
	//		out4 := []string{v}
			fmt.Println([]string{v})
	//		fmt.Println(out4)
	//		fmt.Println(v)
	//		out2[v] = make(chan []string)
		}

		*/


	/*
			out3 := <-out2
			out4 := []string{out3}
			fmt.Println(out4)
	*/


	//	contents := []string{out2}
	/*
		sampleChan := make(chan string, 1)
		var sampleList []string
		var wg sync.WaitGroup
		for _, line := range out2 {
			wg.Add(1)

			go func(line string, ch chan string, wg *sync.WaitGroup) {
				defer wg.Done()
				ch <- line

			}(line, sampleChan, &wg)
		}
		go func (channel chan string, sampleList *[]string) {
			for s := range channel {
				*sampleList = append(*sampleList, s)
			}

		}(sampleChan, &sampleList)
		wg.Wait()
		sort.Strings(sampleList)
		fmt.Println(sampleList)
	*/


	/*
		res := <- f1()
		fmt.Println(res)
	*/
	//	fmt.Println(f1())

	time.Sleep(time.Second * 12)
	return

}


