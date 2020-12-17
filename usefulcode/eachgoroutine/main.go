package main

import (
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type UnifedSCADA struct {
	TotalData map[string]WPPIS `json:"totaldata"`
	Mu        *sync.Mutex
}

type WPPIS struct {
	URL        string          `json:"url"`
	UpdateRate uint16          `json:"updaterate"`
	Time       string          `json:"time"`
	HTTPState  uint16          `json:"httpstate"`
	Turbines   LocationTurbine `json:"turbines"`
	StartTime  time.Time       `json:"start-time"`
}
type LocationTurbine struct {
	ActivePower       string        `json:"activepower"`
	ReactivePower     string        `json:"reactivepower"`
	TotalActiveEnergy string        `json:"totalactiveenergy"`
	WindSpeed         string        `json:"windspeed"`
	WindDirection     string        `json:"winddirection"`
	Temperature       string        `json:"temperature"`
	Turbine           []EachTurbine `json:"turbine"`
}

type EachTurbine struct {
	Name          string `json:"name"`
	WindSpeed     string `json:"windspeed"`
	RotorSpeed    string `json:"rotorspeed"`
	ActivePower   string `json:"activepower"`
	WindDirection string `json:"winddirection"`
	AlarmMsg      string `json:"alarmmsg,omitempty"`
	TurbineStatus string `json:"turbinestatus"`
}

type WPPISXML struct {
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
	start := time.Now()
	ctx, c := w.ToJSON(w.ToXML(w.GetReq(us)))
	// restart := make(chan *WPPIS)
	for {
		select {
		case i := <-c:
			us.Mu.Lock()
			us.TotalData[i.URL] = i
			us.Mu.Unlock()
		case <-ctx.Done():
			fmt.Println("재시작 메시지 확인")
			fmt.Println("해당URL: ", w.URL)
			fmt.Println(w)
			isRestart <- w
			return
		}
		fmt.Printf("URL: %s 업데이트 주기: %d 실행시간: %v Goroutine수: %d\n", w.URL, w.UpdateRate, time.Since(start), runtime.NumGoroutine())
		time.Sleep(time.Duration(w.UpdateRate) * time.Millisecond)
	}
}

// Response -> byte
func (w *WPPIS) GetReq(us *UnifedSCADA) (context.Context, <-chan []byte) {
	ctx, cancel := context.WithCancel(context.Background())
	byteChan := make(chan []byte)
	go func() {
		defer close(byteChan)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client := &http.Client{Transport: tr}
		res, err := client.Get("https://" + w.URL + "/anonymous/RealtimeData.xml?type=DETAIL")
		if err != nil {
			defer cancel()
			// defer w.Start(us)
			log.Println(err, "\n다시 시작합니다 ...")
			log.Println(w.HTTPState)
			return
		}
		defer res.Body.Close()
		byteValue, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
		}
		byteChan <- byteValue
	}()
	return ctx, byteChan
}

// Byte -> XML
func (w *WPPIS) ToXML(ctx context.Context, res <-chan []byte) (context.Context, <-chan WPPISXML) {
	xmlData := WPPISXML{}
	xmlChan := make(chan WPPISXML)
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
	return ctx, xmlChan
}

// XML -> JSON
func (w *WPPIS) ToJSON(ctx context.Context, xmlData <-chan WPPISXML) (context.Context, <-chan WPPIS) {
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
			jsonData.Turbines = LocationTurbine{
				ActivePower:       location.ActivePower.Text,
				ReactivePower:     location.ReactivePower.Text,
				TotalActiveEnergy: location.TotalActiveEnergy.Text,
				WindSpeed:         location.WindSpeed.Text,
				WindDirection:     location.WindDirection.Text,
				Temperature:       location.Temperature.Text,
				Turbine:           []EachTurbine{},
			}

			for _, turbines := range x.GeneralData.WindTurbineSummary.TurbineSummary {
				jsonData.Turbines.Turbine = append(jsonData.Turbines.Turbine, EachTurbine{
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
	return ctx, jsonChan
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

	for {
		for ir := range isRestart {
			fmt.Println("재시작합니다~~~~~~~~~~~~~~~~~~~")
			go ir.Start(us, isRestart)
		}
		us.Mu.Lock()
		// str, err := json.Marshal(us.TotalData)
		// if err != nil {
		// 	log.Println(err)
		// }
		// fmt.Println(string(str))
		us.Mu.Unlock()
		time.Sleep(1000 * time.Millisecond)
	}
}
