package sms

import (
	"Diplom_Project/pkg/utils"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const filename = "../Simuliator/skillbox-diploma/sms.data"

type SMSData struct { // структура дял хранения данных смс
	Сountry      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type ByProvider []SMSData // тип для сортировки по провайдеру

func (a ByProvider) Len() int           { return len(a) }
func (a ByProvider) Less(i, j int) bool { return a[i].Provider < a[j].Provider }
func (a ByProvider) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByCountry []SMSData // тип для сортировки по стране

func (a ByCountry) Len() int           { return len(a) }
func (a ByCountry) Less(i, j int) bool { return a[i].Сountry < a[j].Сountry }
func (a ByCountry) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func SmsDataSort(arr []SMSData) [][]SMSData { //функ сортировки и объединения массивов

	arrResult := make([][]SMSData, 0)
	arrSortProvider := make([]SMSData, 0)
	arrSortCountry := make([]SMSData, 0)

	for _, data := range arr {
		data.Сountry = utils.ReplacementCountryCode(data.Сountry) // замена кода страны на название
		arrSortProvider = append(arrSortProvider, data)
		arrSortCountry = append(arrSortCountry, data)

	}

	sort.Stable(ByProvider(arrSortProvider)) // сортируем по провайдеру

	sort.Stable(ByCountry(arrSortCountry)) // сортируем по стране

	arrResult = append(arrResult, arrSortProvider)

	arrResult = append(arrResult, arrSortCountry)

	return arrResult

}

func ResultData() ([][]SMSData, error) { // функция сбора данных для глобальной структуры
	arr, err := SmsDataFile()

	if err != nil {
		fmt.Println("Ошибка валидации SMS")
		log.Fatalln(err)
	}

	arrResult := SmsDataSort(arr)

	return arrResult, err

}

func SmsDataFile() ([]SMSData, error) { // функ чтения данных из файла и проверки

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	arr := make([]SMSData, 0, 0)
	var result SMSData
	s := string(file)
	s1 := strings.Split(s, "\n")
	for i := 0; i < len(s1); i++ {

		if s1[i] != "" { // проверка на пустоту строки

			configLine := strings.Split(string(s1[i]), ";") //делим строку по разделителю

			switch true {
			case len(configLine) != 4: //проверяем на 4 слова в строке
				continue
			case utils.ValidateProvider(configLine) == false: // проверка на разрешенного провайдера
				continue
			case utils.ValidateCountry(configLine[0]) == false: // проверка на существование страны

			default:
				result.Сountry = configLine[0]
				result.Bandwidth = configLine[1]
				result.ResponseTime = configLine[2]
				result.Provider = configLine[3]

				arr = append(arr, result) // заполняем массив

			}
		}
	}
	return arr, nil
}

func SmsPrint(arr [][]SMSData) { // функция вывода на экран

	for i := 0; i < len(arr); i++ {
		fmt.Println()

		for j := 0; j < len(arr[i]); j++ {

			fmt.Println(arr[i][j])
		}

	}

}
