package email

import (
	"Diplom_Project/pkg/utils"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const filename = "../Simuliator/skillbox-diploma/email.data"

type EmailData struct { //структура для Email
	Country      string
	Provider     string
	DeliveryTime int
}

func ResultData() (map[string][][]EmailData, error) { // функция сбора данных для глобальной структуры
	arr, err := EmailDataFile()

	if err != nil {
		fmt.Println("Ошибка валидации данных Email")
		log.Fatalln(err)
	}

	arrResult := SortEmail(arr)

	return arrResult, err

}

func SortEmail(arr []EmailData) (map[string][][]EmailData) { // функ сортировки и слияния в 2 массива

	arrResult := make(map[string][][]EmailData)
	arrKey := make(map[string][]EmailData)

	for _, data := range arr { // создаем мапу с ключем в виде страны
		key := data.Country

		arrKey[key] = append(arrKey[key], data)

	}

	for _, data := range arrKey { // сортируем мапу по скорости провайдера
		sort.SliceStable(data, func(i, j int) bool {
			return data[i].DeliveryTime < data[j].DeliveryTime
		})
		fast := 3

		if fast > len(data) {
			fast = len(data)
		}

		slow := fast

		fastProviders := data[:fast]
		slowProviders := data[len(data)-slow:]

		for _, value := range data {
			arrResult[value.Country] = [][]EmailData{fastProviders, slowProviders} // заполняем итоговую мапу
		}

	}

	return arrResult

}

func EmailPrint(arrResult map[string][][]EmailData) {

	for _, data := range arrResult {

		fmt.Println(data)
	}

}

func EmailDataFile() ([]EmailData, error) { // функция чтения файла и создания массива
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	arrEmail := make([]EmailData, 0, 0)
	var resultVoice EmailData
	s := string(file)
	s1 := strings.Split(s, "\n")
	for i := 0; i < len(s1); i++ {

		if s1[i] != "" { // проверка на пустоту строки

			configLine := strings.Split(string(s1[i]), ";") //делим строку по разделителю

			switch true {
			case len(configLine) != 3: //проверяем на 3 словa в строке
				continue
			case utils.ValidateProviderEmail(configLine) == false: // проверка на разрешенного провайдера
				continue
			case utils.ValidateCountry(configLine[0]) == false: // проверка на существование страны

			default:
				resultVoice.Country = configLine[0]
				resultVoice.Provider = configLine[1]
				resultVoice.DeliveryTime = utils.StrToIntConv(configLine[2])

				arrEmail = append(arrEmail, resultVoice) // заполняем массив

			}
		}
	}
	return arrEmail, err

}


