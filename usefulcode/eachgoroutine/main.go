package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// UnifedSCADA is Object of total turbines
type UnifedSCADA struct {
	TotalData map[string]WPPIS `json:"totaldata,omitempty"`
	Mu        *sync.Mutex
}

// WPPIS is Object of location turbine
type WPPIS struct {
	URL        string          `json:"url"`
	UpdateRate uint16          `json:"updaterate"`
	Time       string          `json:"time"`
	HTTPState  uint16          `json:"httpstate"`
	Turbines   locationTurbine `json:"turbines,omitempty"`
}

type locationTurbine struct {
	ActivePower       string        `json:"activepower,omitempty"`
	ReactivePower     string        `json:"reactivepower,omitempty"`
	TotalActiveEnergy string        `json:"totalactiveenergy,omitempty"`
	WindSpeed         string        `json:"windspeed,omitempty"`
	WindDirection     string        `json:"winddirection,omitempty"`
	Temperature       string        `json:"temperature,omitempty"`
	Turbine           []eachTurbine `json:"turbine,omitempty"`
}

type eachTurbine struct {
	Name          string `json:"name,omitempty"`
	WindSpeed     string `json:"windspeed,omitempty"`
	RotorSpeed    string `json:"rotorspeed,omitempty"`
	ActivePower   string `json:"activepower,omitempty"`
	WindDirection string `json:"winddirection,omitempty"`
	AlarmMsg      string `json:"alarmmsg,omitempty,omitempty"`
	TurbineStatus string `json:"turbinestatus,omitempty"`
}

type wppisXML struct {
	XMLName     xml.Name `xml:"WPPIS"`
	Time        string   `xml:"time,attr"`
	GeneralData struct {
		WindFarm struct {
			ActivePower struct {
				Text string `xml:",chardata"`
				Unit string `xml:"unit,attr"`
			} `xml:"ActivePower"`
			ReactivePower struct {
				Text string `xml:",chardata"`
				Unit string `xml:"unit,attr"`
			} `xml:"ReactivePower"`
			TotalActiveEnergy struct {
				Text string `xml:",chardata"`
				Unit string `xml:"unit,attr"`
			} `xml:"TotalActiveEnergy"`
			TotalReactiveEnergy struct {
				Text string `xml:",chardata"`
				Unit string `xml:"unit,attr"`
			} `xml:"TotalReactiveEnergy"`
			WindSpeed struct {
				Text string `xml:",chardata"`
				Unit string `xml:"unit,attr"`
			} `xml:"WindSpeed"`
			WindDirection struct {
				Text string `xml:",chardata"`
				Unit string `xml:"unit,attr"`
			} `xml:"WindDirection"`
			Temperature struct {
				Text string `xml:",chardata"`
				Unit string `xml:"unit,attr"`
			} `xml:"Temperature"`
			Frequency struct {
				Text string `xml:",chardata"`
				Unit string `xml:"unit,attr"`
			} `xml:"Frequency"`
			TotalTurbineNum     string `xml:"TotalTurbineNum"`
			OperatingTurbineNum string `xml:"OperatingTurbineNum"`
			WaitingTurbineNum   string `xml:"WaitingTurbineNum"`
			StopTurbineNum      string `xml:"StopTurbineNum"`
			EStopTurbineNum     string `xml:"EStopTurbineNum"`
			ServiceTurbineNum   string `xml:"ServiceTurbineNum"`
			OfflineTurbineNum   string `xml:"OfflineTurbineNum"`
		} `xml:"WindFarm"`
		WindTurbineSummary struct {
			TurbineSummary []struct {
				Name        string `xml:"name,attr"`
				ActivePower struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"ActivePower"`
				ReactivePower struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"ReactivePower"`
				TotalActiveEnergy struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"TotalActiveEnergy"`
				TotalReactiveEnergy struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"TotalReactiveEnergy"`
				TurbineStatus string `xml:"TurbineStatus"`
				WindSpeed     struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"WindSpeed"`
				WindDirection struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"WindDirection"`
				NacellePostion struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"NacellePostion"`
				Temperature struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"Temperature"`
				YawCableWindup struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"YawCableWindup"`
				RotorSpeed struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"RotorSpeed"`
				GeneratorSpeed struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"GeneratorSpeed"`
				Torque struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"Torque"`
				HydraulicPressure struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"HydraulicPressure"`
				HydraulicTemperature struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"HydraulicTemperature"`
				Frequency struct {
					Text string `xml:",chardata"`
					Unit string `xml:"unit,attr"`
				} `xml:"Frequency"`
				AlarmMsg struct {
					WturAlm struct {
						Text     string `xml:",chardata"`
						AlarmNum string `xml:"alarmNum,attr"`
					} `xml:"WturAlm"`
				} `xml:"AlarmMsg"`
			} `xml:"TurbineSummary"`
		} `xml:"WindTurbineSummary"`
	} `xml:"GeneralData"`
}

// Start returns json data from scada servers
func (w *WPPIS) Start(us *UnifedSCADA, isRestart chan *WPPIS) {
	// start := time.Now()
	for {
		ctx, cancel := context.WithCancel(context.Background())
		c := w.toJSON(ctx, w.toXML(ctx, w.getReq(ctx, us)))
		select {
		case i := <-c:
			// 빈 값이 오면 함수 종료
			if i.URL == "" {
				defer cancel()
				log.Println("빈 값이 와서 종료합니다 ...", i)
				isRestart <- w
				return
			}
			us.Mu.Lock()
			us.TotalData[i.URL] = i
			us.Mu.Unlock()
		}
		// fmt.Printf("URL: %s 업데이트 주기: %d 실행시간: %v Goroutine수: %d\n", w.URL, w.UpdateRate, time.Since(start), runtime.NumGoroutine())
		time.Sleep(time.Duration(w.UpdateRate) * time.Millisecond)
	}
}

func (w *WPPIS) Restart(us *UnifedSCADA, isRestart chan *WPPIS) {
	// start := time.Now()
	log.Printf("%s의 재시작 요청을 시작됩니다.\n", w.URL)
	for {
		ctx, cancel := context.WithCancel(context.Background())
		c := w.toJSON(ctx, w.toXML(ctx, w.getReq(ctx, us)))
		select {
		case i := <-c:
			// 빈 값이 오면 함수 종료
			if i.URL == "" {
				defer cancel()
				log.Println("빈 값이 와서 종료합니다 ...")
				isRestart <- w
				return
			}
			log.Println(i)
		}
		// fmt.Printf("URL: %s 업데이트 주기: %d 실행시간: %v Goroutine수: %d\n", w.URL, w.UpdateRate, time.Since(start), runtime.NumGoroutine())
		time.Sleep(time.Duration(w.UpdateRate) * time.Millisecond)
	}
}

// Response -> byte
func (w *WPPIS) getReq(ctx context.Context, us *UnifedSCADA) <-chan []byte {
	byteChan := make(chan []byte)
	go func() {
		defer close(byteChan)
		select {
		case <-ctx.Done():
			return
		default:
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			}
			client := &http.Client{
				Transport: tr,
				Timeout:   5000 * time.Millisecond,
			}
			log.Printf("\n\n%s에 대한 리퀘스트 시작 바로 직전 ...\n", w.URL)
			res, err := client.Get("https://" + w.URL + "/anonymous/RealtimeData.xml?type=DETAIL")
			if err, ok := err.(net.Error); ok && err.Timeout() {
				log.Println("Timeout error!!!")
				us.Mu.Lock()
				us.TotalData[w.URL] = WPPIS{
					URL:        w.URL,
					UpdateRate: w.UpdateRate,
					HTTPState:  http.StatusRequestTimeout,
				}
				us.Mu.Unlock()
				return
			}
			if err != nil {
				log.Println("getReq에서의 에러", err)
				return
			}
			log.Printf("%s의 함수는 진행 중 ...\n", w.URL)
			defer res.Body.Close()
			byteValue, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println(err)
			}
			byteChan <- byteValue
		}
	}()
	log.Printf("%s는 캔슬 요청을 보냈어도 여기까진 실행됨 ...\n", w.URL)
	return byteChan
}

// Byte -> XML
func (w *WPPIS) toXML(ctx context.Context, res <-chan []byte) <-chan wppisXML {
	xmlData := wppisXML{}
	xmlChan := make(chan wppisXML)
	go func() {
		defer close(xmlChan)
		for r := range res {
			err := xml.Unmarshal(r, &xmlData)
			if err != nil {
				log.Println(err)
			}
			select {
			case xmlChan <- xmlData:
			case <-ctx.Done():
				return
			}
		}
	}()
	return xmlChan
}

// XML -> JSON
func (w *WPPIS) toJSON(ctx context.Context, xmlData <-chan wppisXML) <-chan WPPIS {
	jsonData := WPPIS{}
	jsonChan := make(chan WPPIS)
	go func() {
		defer close(jsonChan)
		for x := range xmlData {
			w.HTTPState = http.StatusOK
			location := x.GeneralData.WindFarm
			jsonData.Time = x.Time
			jsonData.URL = w.URL
			jsonData.UpdateRate = w.UpdateRate
			jsonData.Turbines = locationTurbine{
				ActivePower:       location.ActivePower.Text,
				ReactivePower:     location.ReactivePower.Text,
				TotalActiveEnergy: location.TotalActiveEnergy.Text,
				WindSpeed:         location.WindSpeed.Text,
				WindDirection:     location.WindDirection.Text,
				Temperature:       location.Temperature.Text,
				Turbine:           []eachTurbine{},
			}

			for _, turbines := range x.GeneralData.WindTurbineSummary.TurbineSummary {
				jsonData.Turbines.Turbine = append(jsonData.Turbines.Turbine, eachTurbine{
					Name:          turbines.Name,
					WindSpeed:     turbines.WindSpeed.Text,
					RotorSpeed:    turbines.RotorSpeed.Text,
					ActivePower:   turbines.ActivePower.Text,
					WindDirection: turbines.WindDirection.Text,
					AlarmMsg:      turbines.AlarmMsg.WturAlm.Text,
					TurbineStatus: turbines.TurbineStatus,
				})
			}
			select {
			case jsonChan <- jsonData:
			case <-ctx.Done():
				return
			}
		}
	}()
	return jsonChan
}

var us *UnifedSCADA

func main() {
	servers := []*WPPIS{
		{
			URL:        "galapagos.scada.unison.co.kr",
			UpdateRate: 5000,
		},
		{
			URL:        "ansan.scada.unison.co.kr",
			UpdateRate: 2000,
		},
	}

	us = &UnifedSCADA{
		TotalData: make(map[string]WPPIS),
		Mu:        &sync.Mutex{},
	}

	isRestart := make(chan *WPPIS)
	for _, s := range servers {
		go s.Start(us, isRestart)
	}

	go func() {
		for {
			for ir := range isRestart {
				log.Printf("%s의 리퀘스트 재시작 요청 ...\n\n", ir.URL)
				log.Println("goroutine 갯수: ", runtime.NumGoroutine())
				go ir.Restart(us, isRestart)
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}()

	for {
		us.Mu.Lock()
		str, err := json.Marshal(us.TotalData)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(str))
		us.Mu.Unlock()
		time.Sleep(1000 * time.Millisecond)
	}
}
