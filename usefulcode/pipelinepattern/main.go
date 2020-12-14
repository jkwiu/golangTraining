package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
	"unsafe"
)

func main() {

	ansan := &WPPIS{URL: "https://ansan.scada.unison.co.kr/anonymous/RealtimeData.xml?type=DETAIL", UpdateRate: 3000, Start: time.Now()}
	yeonggwang := &WPPIS{URL: "https://yeonggwang.scada.unison.co.kr/anonymous/RealtimeData.xml?type=DETAIL", UpdateRate: 5000, Start: time.Now()}

	var servers []*WPPIS
	servers = append(servers, ansan, yeonggwang)

	for _, server := range servers {
		go server.start()
	}
	time.Sleep(20000 * time.Second)
}

type WPPIS struct {
	URL        string `json:"url"`
	UpdateRate int    `json:"udpate-rate"`
	Data       string `json:"data"`
	Start      time.Time
}

type UnifiedScada struct {
	TotalData []WPPIS `json:"total-data"`
}

func (w *WPPIS) start() {
	for {
		fmt.Println("start again")
		for ele := range w.getReq() {
			w.Data = string(ele[0])
			rst, err := json.Marshal(w)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("All:", runtime.NumGoroutine(), "runnable:", NumRunnableGoroutine(), "running", NumRunningGoroutine())
			fmt.Println(string(rst))
			fmt.Println("실행시간: ", time.Since(w.Start))
		}
		time.Sleep(time.Duration(w.UpdateRate) * time.Millisecond)
	}
}

// Make request and send byte data to channel
func (w *WPPIS) getReq() <-chan []byte {
	byteChan := make(chan []byte)
	go func() {
		defer close(byteChan)
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client := &http.Client{Transport: tr}
		res, err := client.Get(w.URL)
		if err != nil {
			log.Println(err)
		}
		defer res.Body.Close()
		byteValue, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
		}
		byteChan <- byteValue
	}()
	return byteChan
}

func (w *WPPIS) toStr(in <-chan []byte) string {
	var rst string
	for ele := range in {
		rst = string(ele)
		fmt.Println(rst)
	}
	return rst
}

func (w *WPPIS) stop() {

}

//go:linkname readgstatus runtime.readgstatus
//go:nosplit
func readgstatus(gp unsafe.Pointer) uint32

//go:linkname allgs runtime.allgs
var allgs []unsafe.Pointer

type mutex struct {
	key uintptr
}

//go:linkname allglock runtime.allglock
var allglock mutex

//go:linkname lock runtime.lock
func lock(l *mutex)

//go:linkname unlock runtime.unlock
func unlock(l *mutex)

const (
	_Grunnable = 1
	_Grunning  = 2
)

func NumRunnableGoroutine() (num int) {
	lock(&allglock)
	for _, g := range allgs {
		if readgstatus(g)&^0x1000 == _Grunnable {
			num++
		}
	}
	unlock(&allglock)
	return
}

func NumRunningGoroutine() (num int) {
	lock(&allglock)
	for _, g := range allgs {
		if readgstatus(g)&^0x1000 == _Grunning {
			num++
		}
	}
	unlock(&allglock)

	return
}
