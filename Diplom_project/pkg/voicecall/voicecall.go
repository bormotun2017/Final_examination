package voicecall

import (
	"Diplom_Project/pkg/utils"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const filename = "../Simuliator/skillbox-diploma/voice.data"

// VoiceData - Структура для хранениния данных о звонках
type VoiceData struct {
	Country             string
	CurrentLoad         int
	AverageResponseTime int
	Provider            string
	ConnectionStability float32
	PurityCommunication int
	TTFB                int
	MedianCallDuration  int
}

func VoiceCallDataFile() ([]VoiceData, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	arrVoice := make([]VoiceData, 0, 0)
	var resultVoice VoiceData
	s := string(file)
	s1 := strings.Split(s, "\n")
	for i := 0; i < len(s1); i++ {

		if s1[i] != "" { // проверка на пустоту строки

			configLine := strings.Split(string(s1[i]), ";") //делим строку по разделителю

			switch true {
			case len(configLine) != 8: //проверяем на 8 слов в строке
				continue
			case utils.ValidateProviderVoiceCall(configLine) == false: // проверка на разрешенного провайдера
				continue
			case utils.ValidateCountry(configLine[0]) == false: // проверка на существование страны

			default:
				resultVoice.Country = configLine[0]
				resultVoice.CurrentLoad = utils.StrToIntConv(configLine[1])
				resultVoice.AverageResponseTime = utils.StrToIntConv(configLine[2])
				resultVoice.Provider = configLine[3]
				resultVoice.ConnectionStability = utils.StrToFloat32Conv(configLine[4])
				resultVoice.PurityCommunication = utils.StrToIntConv(configLine[5])
				resultVoice.TTFB = utils.StrToIntConv(configLine[6])
				resultVoice.MedianCallDuration = utils.StrToIntConv(configLine[7])

				arrVoice = append(arrVoice, resultVoice) // заполняем массив

			}
		}
	}

	return arrVoice, err
}

func VoiceCallPrint(arrVoice []VoiceData) {

	for i, _ := range arrVoice {
		fmt.Println(arrVoice[i])
	}
}
