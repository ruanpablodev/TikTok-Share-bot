package main

import (
	"fmt"
	"math/rand"
	"time"

	tm "github.com/buger/goterm"
	randomUserAgent "github.com/corpix/uarand"
	"github.com/valyala/fasthttp"
)

var (
	client fasthttp.Client

	errors int = 0
	sent   int = 0
	rpm    int = 0
)

func addShare(itemID string) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.Header.SetMethod("POST")
	req.SetRequestURI(generateURL())
	req.SetBody([]byte(fmt.Sprintf("item_id=%v&share_delta=1", itemID)))

	req.Header.Set("User-Agent", randomUserAgent.GetRandom())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	if err := client.Do(req, res); err != nil {
		errors++
		return
	}

	sent++
}

func rpmCounter() {
	for {
		before := sent
		time.Sleep(100 * time.Millisecond)
		after := sent

		rpm = (after - before) * 600
	}
}

func statusPrinter() {
	for {
		fmt.Printf("[+] Enviados: %v | Requests por minuto: %v | Erros: %v\r", sent, rpm, errors)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	tm.Clear()

	fmt.Print("[+] instagram.com/rp.xyz\n\n")

	var threads int
	var itemID string

	fmt.Print("[>] Threads: ")
	fmt.Scanln(&threads)
	fmt.Print("O id do video se pega aki https://www.tiktok.com/@USERNAME/video/<ESSE NUMERO AKI> \n")
	fmt.Print("[>] ID do video: ")
	fmt.Scanln(&itemID)

	fmt.Print("\n\n")

	go rpmCounter()
	go statusPrinter()

	for i := 0; i < threads; i++ {
		go func() {
			for {
				addShare(itemID)
			}
		}()
	}

	select {}
}
