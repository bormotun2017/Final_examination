package incident

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

type IncidentData struct { //структура хранения данных Incident
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func IncidentSort(arrSort []IncidentData, err error) ([]IncidentData, error) { // функция сортировки инцидентов
	(sort.Slice(arrSort, func(i, j int) bool {
		return arrSort[i].Status < arrSort[j].Status
	}))
	return arrSort, err
}

func ResultData() ([]IncidentData, error) { // функция сбора данных системы инцидентов
	arrResult, err := IncidentSort(GetIncidentData())

	return arrResult, err

}

func GetIncidentData() ([]IncidentData, error) { //вункция гет запроса инцидентов

	result, err := http.Get("http://127.0.0.1:8383/accendent") // отправляем гет запрос
	if err != nil {
		fmt.Println("Ошибка гет запроса к инцидентам")
		log.Fatal(err)
	}
	body, err := io.ReadAll(result.Body) //читаем тело ответа
	result.Body.Close()

	if result.StatusCode != 200 { // проверям код ответа
		fmt.Println("Код ответа сервера не равен 200")
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", result.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var TrialIncidentData []IncidentData //создаем массив
	var IncidentData []IncidentData
	err = json.Unmarshal(body, &TrialIncidentData) // переводим тело запроса в наш массив с помощью встроенной функ
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, datum := range TrialIncidentData { //разбираем массив
		if datum.Status == "active" || datum.Status == "closed" { //проверяем соответствие статуса
			IncidentData = append(IncidentData, datum) //
		}
		//	fmt.Println(i, datum)

	}
	fmt.Println()

	return IncidentData, nil

}

func IncidentPrint() {

	arr, err := ResultData()
	if err != nil {
		fmt.Println(err)
	}

	for i, data := range arr {
		fmt.Println(i, data)

	}
}
