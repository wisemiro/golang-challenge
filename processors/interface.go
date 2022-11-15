package processors

import (
	"sync"
)

type DriverRanking struct {
	// Fill in your properties here
	ID   string `json:"id"`
	Name string `json:"name"`
	Rate int64  `json:"rate"`
}

func (s *DriverRanking) String() string {
	// Implement this function
	return s.Name
}

type HotelRanking struct {
	// Fill in your properties here
	ID   string `json:"id"`
	Name string `json:"name"`
	Rate int64  `json:"rate"`
}

func (s *HotelRanking) String() string {
	// Implement this function
	return s.Name
}

type ProcessorInterface interface {
	StartProcessing() error
	GetTopRankedDriver() *DriverRanking
	GetTopRankedHotel() *HotelRanking
}

func CreateProcessorFromData(data *TripsData, wg *sync.WaitGroup) ProcessorInterface {
	// @todo Initialize your processor here
	// Open data.json
	//jsonFile, err := os.Open("data.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("Successfully Opened data.json")
	//// defer the closing of data.json, can parse it later on
	//defer func(jsonFile *os.File) {
	//	jsnErr := jsonFile.Close()
	//	if jsnErr != nil {
	//		fmt.Println(jsnErr)
	//	}
	//}(jsonFile)
	//
	////Read the data.json contents
	//byteResult, readErr := io.ReadAll(jsonFile)
	//if readErr != nil {
	//	log.Println(readErr)
	//}
	//
	//var res map[string]interface{}
	//marshErr := json.Unmarshal([]byte(byteResult), &res)
	//if marshErr != nil {
	//	log.Println(marshErr)
	//}
	//
	//fmt.Println(res["users"])

	return nil
}
