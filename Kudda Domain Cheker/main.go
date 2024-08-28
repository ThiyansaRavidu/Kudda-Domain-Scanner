package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv" // <-- Import this package
	"strings"
	"time"
)

// API descriptions and endpoints
var apis = map[int]string{
	1:  "DNS Lookup",
	2:  "Reverse DNS Lookup",
	3:  "IP Geolocation Lookup",
	4:  "Reverse IP Lookup",
	5:  "HTTP Headers",
	6:  "Page Links",
	7:  "AS Lookup",
	8:  "DNS Lookup Advanced",
	9:  "Response Checker Advanced",
	10: "NS Lookup",
}

var apiEndpoints = map[int]string{
	1:  "https://api.hackertarget.com/dnslookup/?q=%s",
	2:  "https://api.hackertarget.com/reversedns/?q=%s",
	3:  "https://api.hackertarget.com/ipgeo/?q=%s",
	4:  "https://api.hackertarget.com/reverseiplookup/?q=%s",
	5:  "https://api.hackertarget.com/httpheaders/?q=%s",
	6:  "https://api.hackertarget.com/pagelinks/?q=%s",
	7:  "https://api.hackertarget.com/aslookup/?q=%s",
}

// Record structure to hold DNS records
type Record struct {
	A     []string `json:"A"`
	AAAA  []string `json:"AAAA"`
	MX    []string `json:"MX"`
	NS    []string `json:"NS"`
	TXT   []string `json:"TXT"`
	CNAME []string `json:"CNAME"`
	SOA   []string `json:"SOA"`
}

type ResponseDetails struct {
	StatusCode int
	Status     string
	Header     map[string][]string
}

type TLSDetails struct {
	Sni           string
	Enabled       bool
	TlsVersion    string
	Alpn          interface{}
	Istlssucsess  bool
}

type SpeedtestDetails struct {
	MaxSpeed     float64
	Failed       bool
	FailedReason string
}

type DomainInfo struct {
	Speedtest          SpeedtestDetails
	IsClientFailed     bool
	FailedReason       string
	RemoteIP           string
	Host               string
	TLS                TLSDetails
	IsResponseReceived bool
	Headers            map[string][]string
	Method             string
	HTTPVersion        string
	HTTPPath           string
	ErrType            string
	HTTPPort           int
	Response           ResponseDetails
}

func getTerminalWidth() int {
	// Default terminal width
	return 80
}
func displayBanner() {
	terminalWidth := getTerminalWidth()

	// Print the banner
	fmt.Println(" ██ ▄█▀ █    ██ ▓█████▄ ▓█████▄  ▄▄▄      ")
	fmt.Println(" ██▄█▒  ██  ▓██▒▒██▀ ██▌▒██▀ ██▌▒████▄    ")
	fmt.Println("▓███▄░ ▓██  ▒██░░██   █▌░██   █▌▒██  ▀█▄  ")
	fmt.Println("▓██ █▄ ▓▓█  ░██░░▓█▄   ▌░▓█▄   ▌░██▄▄▄▄██ ")
	fmt.Println("▒██▒ █▄▒▒█████▓ ░▒████▓ ░▒████▓  ▓█   ▓██▒")
	fmt.Println("▒ ▒▒ ▓▒░▒▓▒ ▒ ▒  ▒▒▓  ▒  ▒▒▓  ▒  ▒▒   ▓▒█░")
	fmt.Println("░ ░▒ ▒░░░▒░ ░ ░  ░ ▒  ▒  ░ ▒  ▒   ▒   ▒▒ ░")
	fmt.Println("░ ░░ ░  ░░░ ░ ░  ░ ░  ░  ░ ░  ░   ░   ▒   ")
	fmt.Println("░  ░      ░        ░       ░          ░  ░")
	fmt.Println("                 ░       ░                ")
	fmt.Println()

	// Center "Made By Thiyansa" and "Version 1.0" below the banner
	printCentered("Made By Thiyansa", terminalWidth)
	printCentered("Version 1.0", terminalWidth)
	fmt.Println()
}

func printCentered(text string, width int) {
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	fmt.Println(strings.Repeat(" ", padding) + text)
}

func displayMenu() {
	fmt.Println("Select an API to query:")
	for number, description := range apis {
		fmt.Printf("%d. %s\n", number, description)
	}
	fmt.Println() // Adds a newline for better readability
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter domain or IP address: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func fetchAPIData(apiURL string) (string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}
	return string(body), nil
}

func performDNSLookup(domain string) (Record, error) {
	var result Record

	// A Records
	aRecords, err := net.LookupHost(domain)
	if err == nil {
		result.A = aRecords
	}

	// AAAA Records
	aaaaRecords, err := net.LookupIP(domain)
	if err == nil {
		for _, ip := range aaaaRecords {
			if ip.To4() == nil {
				result.AAAA = append(result.AAAA, ip.String())
			}
		}
	}

	// MX Records
	mxRecords, err := net.LookupMX(domain)
	if err == nil {
		for _, mx := range mxRecords {
			result.MX = append(result.MX, mx.Host)
		}
	}

	// NS Records
	nsRecords, err := net.LookupNS(domain)
	if err == nil {
		for _, ns := range nsRecords {
			result.NS = append(result.NS, ns.Host)
		}
	}

	// TXT Records
	txtRecords, err := net.LookupTXT(domain)
	if err == nil {
		result.TXT = txtRecords
	}

	// CNAME Records (only applicable if the domain has a CNAME)
	cname, err := net.LookupCNAME(domain)
	if err == nil && cname != domain {
		result.CNAME = append(result.CNAME, cname)
	}

	// SOA Records (this requires a custom resolver as it's not available in the net package)
	soaRecords, err := lookupSOA(domain)
	if err == nil {
		result.SOA = soaRecords
	}

	return result, nil
}

func lookupSOA(domain string) ([]string, error) {
	var soaRecords []string
	return soaRecords, fmt.Errorf("SOA record lookup not implemented")
}

func getDomainInfo(domain string) (*DomainInfo, error) {
	info := &DomainInfo{
		Host:       domain,
		Method:     "GET",
		HTTPVersion: "1.1",
		HTTPPath:   "/",
		HTTPPort:   443,
	}

	// Resolve IP address
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, fmt.Errorf("failed to get IP addresses: %v", err)
	}
	info.RemoteIP = ips[0].String()

	// Fetch HTTP response
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://" + domain)
	if err != nil {
		return nil, fmt.Errorf("failed to get HTTP response: %v", err)
	}
	defer resp.Body.Close()

	info.IsResponseReceived = true
	info.Response.StatusCode = resp.StatusCode
	info.Response.Status = resp.Status
	info.Response.Header = resp.Header

	// Fill TLS details (mocked)
	info.TLS = TLSDetails{
		Sni: "mocked.sni",
		Enabled: true,
		TlsVersion: "mocked.version",
		Alpn: nil,
		Istlssucsess: true,
	}

	// Fill speed test details (mocked)
	info.Speedtest = SpeedtestDetails{
		MaxSpeed: 1.051802979631625,
		Failed: false,
		FailedReason: "",
	}

	return info, nil
}

func performNSLookup(domain string) {
	ips, err := net.LookupNS(domain)
	if err != nil {
		fmt.Printf("Error performing NS lookup: %v\n", err)
		return
	}

	fmt.Println("NS Records:")
	for _, ns := range ips {
		fmt.Println(ns.Host)
	}
}

func main() {
	displayBanner()
	displayMenu()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of your choice: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	if apiURL, exists := apiEndpoints[choice]; exists {
		input := getInput()
		apiURL = fmt.Sprintf(apiURL, input)
		data, err := fetchAPIData(apiURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	} else if choice == 8 {
		domain := getInput()
		result, err := performDNSLookup(domain)
		if err != nil {
			fmt.Printf("Error performing DNS lookup: %v\n", err)
			return
		}

		resultJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Printf("Error formatting DNS lookup result: %v\n", err)
			return
		}

		fmt.Println(string(resultJSON))
	} else if choice == 9 {
		domain := getInput()
		domainInfo, err := getDomainInfo(domain)
		if err != nil {
			fmt.Printf("Error getting domain info: %v\n", err)
			return
		}

		resultJSON, err := json.MarshalIndent(domainInfo, "", "  ")
		if err != nil {
			fmt.Printf("Error formatting domain info: %v\n", err)
			return
		}

		fmt.Println(string(resultJSON))
	} else if choice == 10 {
		domain := getInput()
		performNSLookup(domain)
	} else {
		fmt.Println("Invalid choice.")
	}
}
