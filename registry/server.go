package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

/*
* @Author: lyr1cs
* @Email: linyugang7295@gmail.com
* @Description:
* @Date: 2024-10-24 13:02
 */

const ServerPort = ":3000"
const ServicesURL = "http://loaclhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.Mutex),
}

// RegistryService 实际上一个实现了Hander接口的服务，后续调用Hande函数需要对应的Hander
type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	switch r.Method {
	case http.MethodPost:
		// 解析请求体
		dec := json.NewDecoder(r.Body)
		// 服务信息
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %s with URL: %s\n", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
