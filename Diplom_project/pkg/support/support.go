package support

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SupportData struct { // структура хранения для Support
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func ResultData() (arrResult []int, err error) { // функция сбора данных о системме Support
	arr, err := GetSupportData()

	if err != nil {
		fmt.Println("Ошибка функции Get запроса к Support")
		log.Fatalln(err)
	}
	var tickets, workload, responseTime int

	for _, item := range arr {
		tickets += item.ActiveTickets
	}

	switch {
	case tickets < 9:
		workload = 1
	case tickets <= 16:
		workload = 2
	default:
		workload = 3
	}

	responseTime = int((float64(60) / float64(18)) * float64(tickets))
	arrResult = append(arrResult, workload)
	arrResult = append(arrResult, responseTime)

	return arrResult, err
}

func GetSupportData() ([]SupportData, error) { // функ гет запроса к серверу
	var SupportData []SupportData

	result, err := http.Get("http://127.0.0.1:8383/support") //направляем запрос на адрес
	if err != nil {
		fmt.Println("Ошибка запроса к серверу support")
		return SupportData, errors.New("Ошибка запроса к серверу support")
	}
	defer result.Body.Close()
	if result.StatusCode != 200 { //проверяем код ответа
		fmt.Println("Ошибка состояния сервера, код ответа не равенн 200")
		return SupportData, errors.New("Ошибка состояния сервера, код ответа не равенн 200")
	}
	body, err := ioutil.ReadAll(result.Body) // считываем тело ответа в массив байт

	if err != nil {
		log.Fatalln(err)
	}

	if json.Unmarshal(body, &SupportData) != nil { // преобразование байт в массив с нашей структурой
		fmt.Println("Ошибка чтения данных json об Support")
		return SupportData, errors.New("Ошибка чтения данных json о Support")
	}

	return SupportData, nil

}

func SupportPrint() {

	arrSupport, err := GetSupportData()

	arrResult, err := ResultData()
	if err != nil {
		log.Fatalln(err)
	}

	for _, data := range arrSupport { // печатаем массив get запроса
		fmt.Println(data)

	}

	fmt.Println()
	fmt.Println("загруженность Support")
	fmt.Println("среднее время ожидания ответа")

	for i, _ := range arrResult { // печатаем Result массив
		fmt.Println(arrResult[i])
	}

}
