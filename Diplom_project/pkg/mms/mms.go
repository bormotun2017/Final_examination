package mms

import (
	"Diplom_Project/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type ByProvider []MMSData // тип для сортировки по провайдеру

func (a ByProvider) Len() int           { return len(a) }
func (a ByProvider) Less(i, j int) bool { return a[i].Provider < a[j].Provider }
func (a ByProvider) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByCountry []MMSData // тип для сортировки по стране

func (a ByCountry) Len() int           { return len(a) }
func (a ByCountry) Less(i, j int) bool { return a[i].Country < a[j].Country }
func (a ByCountry) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func MmsDataSort(arr []MMSData) [][]MMSData { //функ сортировки и объединения массивов

	arrResult := make([][]MMSData, 0)
	arrSortProvider := make([]MMSData, 0)
	arrSortCountry := make([]MMSData, 0)

	for _, data := range arr {
		data.Country = utils.ReplacementCountryCode(data.Country) // замена кода страны на название
		arrSortProvider = append(arrSortProvider, data)
		arrSortCountry = append(arrSortCountry, data)

	}

	sort.Stable(ByProvider(arrSortProvider)) // сортируем по провайдеру

	sort.Stable(ByCountry(arrSortCountry)) // сортируем по стране

	arrResult = append(arrResult, arrSortProvider)

	arrResult = append(arrResult, arrSortCountry)

	return arrResult

}

func ResultData() ([][]MMSData, error) { // функция сбора данных для глобальной структуры
	arr, err := MmsDataValidate()

	if err != nil {
		fmt.Println("Ошибка валидации данных MMS")
		log.Fatalln(err)
	}

	arrResult := MmsDataSort(arr)

	return arrResult, err

}

func GetMmsData() ([]MMSData, error) { // функ гет запроса к серверу
	var mmsData []MMSData

	result, err := http.Get("http://127.0.0.1:8383/mms") //направляем запрос на адрес
	if err != nil {
		fmt.Println("Ошибка запроса к серверу mms")
		return mmsData, errors.New("Ошибка запроса к серверу mms")
	}
	defer result.Body.Close()
	if result.StatusCode != 200 { //проверяем код ответа
		fmt.Println("Ошибка состояния сервера, код ответа не равенн 200")
		return mmsData, errors.New("Ошибка состояния сервера, код ответа не равенн 200")
	}
	body, err := ioutil.ReadAll(result.Body) // считываем тело ответа в массив байт

	if err != nil {
		log.Fatalln(err)
	}

	if json.Unmarshal(body, &mmsData) != nil { // преобразование байт в массив с нашей структурой
		fmt.Println("Ошибка чтения данных json об mms")
		return mmsData, errors.New("Ошибка чтения данных json об mms")
	}

	return mmsData, nil
}

func MmsDataValidate() ([]MMSData, error) {
	var result MMSData
	arr := make([]MMSData, 0)
	mmsData, err := GetMmsData() // получаем массив данных из гет запроса
	if err != nil {
		log.Fatalln(err)
	}

	for i, s := range mmsData { // циклим массив

		switch true {
		case utils.ValidateProviderMMS(s.Provider) == false: // проверка на разрешенного провайдера
			continue
		case utils.ValidateCountry(s.Country) == false: // проверка на существование страны

		default:
			result.Country = mmsData[i].Country
			result.Provider = mmsData[i].Provider
			result.Bandwidth = mmsData[i].Bandwidth
			result.ResponseTime = mmsData[i].ResponseTime

			arr = append(arr, result) // заполняем новый  массив после проверки на провайдера и страну

		}
	}

	return mmsData, nil

}

func MmsPrint(arr [][]MMSData) { // функция вывода на экран

	for i := 0; i < len(arr); i++ {
		fmt.Println()

		for j := 0; j < len(arr[i]); j++ {

			fmt.Println(arr[i][j])
		}

	}

}
