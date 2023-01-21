package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func CodeCountry() []string { // получение кодов стран
	arr := ReadCsvFile()
	arrCode := make([]string, 0)
	for i := 0; i < len(arr); i++ { // пробегаем по массиву
		arrCode = append(arrCode, arr[i][1]) //заполняем кодами
	}
	//	for i, _ := range arrCode {
	//		fmt.Println(arrCode[i])
	//	}
	return arrCode
}

func NameCountry() []string { // получение имен стран
	arr := ReadCsvFile()
	arrName := make([]string, 0)
	for i := 0; i < len(arr); i++ { // пробегаем по массиву
		arrName = append(arrName, arr[i][0]) //заполняем именами
	}
	for i, _ := range arrName {
		fmt.Println(arrName[i])
	}
	return arrName
}

//func ValidateCountry(line []string) bool { // функция проверки страны
//	arr := ReadCsvFile()
//	for i := 0; i < len(arr); i++ { // пробегаем по массиву
//		for j := 0; j < len(arr[i]); j++ {
//			if arr[i][j] == line[0]+"\r" || arr[i][j] == line[0] { // /r-перенос строки, без него не видит 2й символ

//				return true
//			}
//			//fmt.Println(arr[i][j])
//		}
//	}
//	return false

// }

func ReplacementCountryCode(line string) string {
	arr := ReadCsvFile()

	for i := 0; i < len(arr); i++ { // пробегаем по массиву
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == line+"\r" || arr[i][j] == line { //
				line = arr[i][0]

			}
		}
	}
	return line
}

func ValidateCountry(line string) bool { // функция проверки страны
	arr := ReadCsvFile()
	for i := 0; i < len(arr); i++ { // пробегаем по массиву
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == line+"\r" || arr[i][j] == line { //

				return true
			}
			//fmt.Println(arr[i][j])
		}
	}
	return false

}

func ValidateProvider(configline []string) bool { // функция проверки провайдера
	arr := []string{"Topolo", "Rond", "Kildy"}
	for i, _ := range arr {
		if arr[i] == configline[3] {
			return true
		}

	}
	return false
}

func ValidateProviderEmail(configline []string) bool { // функция проверки провайдера для Email
	arr := []string{
		"Gmail",
		"Yahoo",
		"Hotmail",
		"MSN",
		"Orange",
		"Comcast",
		"AOL",
		"Live",
		"RediffMail",
		"GMX",
		"Proton Mail",
		"Yandex",
		"Mail.ru",
	}
	for i, _ := range arr {
		if arr[i] == configline[1] {
			return true
		}

	}
	return false
}

func ValidateProviderVoiceCall(configline []string) bool { // функция проверки провайдера для VoiceCall
	arr := []string{"TransparentCalls", "E-Voice", "JustPhone"}
	for i, _ := range arr {
		if arr[i] == configline[3] {
			return true
		}

	}
	return false
}

func ValidateProviderMMS(configline string) bool { // функция проверки провайдера для MMS
	arr := []string{"Topolo", "Rond", "Kildy"}
	for i, _ := range configline {
		if arr[i] == configline {
			return true
		}
	}

	return false
}

func ReadCsvFile() [][]string { // функция чтения файла со странами

	file, err := ioutil.ReadFile("pkg/utils/countries-codes.csv") // читаем файл
	if err != nil {
		log.Fatalln(err)
	}
	out := make([][]string, 0) // создаем массив
	if err != nil {
		log.Fatalln(err)
	}
	s := strings.Split(string(file), "\n") // сплитуем построчно

	for _, line := range s { // читаем строку
		if len(line) > 0 { //отсекаем кавычки
			lastind := strings.LastIndex(line, ",")
			arr := []string{line[:lastind], line[lastind+1:]}
			out = append(out, arr)
		}
	}
	return out
}

func StrToIntConv(s string) int { // конвертирует str в int
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return res
}

func StrToFloat32Conv(s string) float32 { // конвертирует str в int
	res, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Fatalln(err)
	}

	return float32(res)
}

func ByteToBool(x byte) bool {
	if x == 48 {
		return false
	}
	return true
}
