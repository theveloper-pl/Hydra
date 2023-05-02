package net

import (
	"fmt"
	"github.com/go-ping/ping"
	"time"
	"regexp"
	"net"
	"strconv"
	"sync"
)


type Host struct {
    domain []string
    addr  []string
    cname  string
    txt  []string
	errors []error
}


var MAX_CONCURRENT_JOBS = 5

type ScanResult struct {
	Port  int
	State string
	Service string
}



func (host Host)describe(){
	fmt.Println("--------Report for host--------")


        fmt.Printf("Domain: %v\n", host.domain)
        fmt.Printf("Address: %v\n", host.addr)
        fmt.Printf("Cname: %v\n", host.cname)
        fmt.Printf("TXT: %v\n\n", host.txt)
        fmt.Printf("Errors: %v\n", host.errors)

	fmt.Println("-------------------------------")
	
}

func do_ping(host string, ping_amount int) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		fmt.Printf("Error creating pinger: %v\n", err)
		return
	}

	pinger.Count = ping_amount           // Liczba wysyłanych pakietów ping
	pinger.Timeout = 5 * time.Second     // Maksymalny czas oczekiwania na odpowiedź
	pinger.Interval = 1 * time.Second    // Czas pomiędzy kolejnymi próbami ping
	pinger.SetPrivileged(true)           // Need to be set on windows

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s Ping raport ---\n", stats.Addr)
		fmt.Printf("%d Packets transmitted, %d Packets received, %v%% Packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("Round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		fmt.Printf("Error running pinger: %v\n", err)
		return
	}
}


func do_discover(host string){
	h := Host{}
	h, err := get_ip(host)
	if err !=nil{
		fmt.Printf("Error occured while discovering host", err)
	}

	for _, s := range h.domain {
		cname, err := do_lookup_cname(s)
		if err != nil {
			h.errors = append(h.errors, err)
		}
		h.cname = cname

		txt, err := do_lookup_txt(s)
		if err != nil {
			h.errors = append(h.errors, err)
		}
		h.txt = txt

	}
	h.describe()

}


func get_ip(host string) (Host, error){
	h := Host{}
	if match, _ := regexp.MatchString("([a-z]+).([a-z]+)", host); match != false{
		value, err := do_find_addr(host)
		if err != nil {
			fmt.Printf("Error running do_find_addr: %v\n", err)
			return Host{}, err
		}
		h.addr = value
		h.domain = append(h.domain, host)

	}else{
		value, err := do_find_host(host)
		if err != nil {
			fmt.Printf("Error running do_find_host: %v\n", err)
			return Host{}, err
		}
		h.domain = value
		h.addr = append(h.addr, host)
	}
	return h, nil
}




func do_find_addr(host string) ([]string, error){
	addrs, err := net.LookupHost(host)
	return addrs, err
}

func do_find_host(host string) ([]string, error){
	names, err := net.LookupAddr(host)
	return names, err
}

func do_lookup_cname(host string) (string, error){
	cname, err := net.LookupCNAME(host)
	return cname, err
}

func do_lookup_txt(host string) ([]string, error){
	txt, err := net.LookupTXT(host)
	return txt, err
}

func do_scanPorts(host string, only_open bool, only_tcp bool, only_udp bool, start_port int, end_port int){
	h, err := get_ip(host)

	if err !=nil{
		fmt.Printf("Error occured while scanning ports", err)
	}

	fmt.Printf("Port Scanning on %s \n", host)
	results := InitialScan(h.addr[0], only_open, only_tcp, only_udp, start_port, end_port)

	for i,e := range results{
		if(e.State == "Open" && only_open){
			fmt.Printf("%d : %s - %s\n", i+start_port, e.State, e.Service)			
		}else if (!only_open) {
			fmt.Printf("%d : %s - %s\n", i+start_port, e.State, e.Service)
		}
	}
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	result.Service = protocol

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}

func InitialScan(hostname string, only_open bool, only_tcp bool, only_udp bool, start_port int, end_port int) []ScanResult {

	var results []ScanResult
	wg := &sync.WaitGroup{}
	port_range := (end_port - start_port)

	if (only_tcp || only_udp){
		wg.Add(port_range*2)
	}else{
		wg.Add(port_range)
	}
	

	for i := start_port; i <= end_port; i++ {

		if(!only_udp){
			go func(number int){
				defer wg.Done()
				results = append(results, ScanPort("tcp", hostname, number))
			}(i)			
		}
		if(!only_tcp){
			go func(number int){
				defer wg.Done()
				results = append(results, ScanPort("udp", hostname, number))
			}(i)			
		}


	}

	wg.Wait()
	return results
}