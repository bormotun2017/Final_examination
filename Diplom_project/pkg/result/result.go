package result

import (
	incident "Diplom_Project/pkg/Incident"
	"Diplom_Project/pkg/billing"
	"Diplom_Project/pkg/email"
	"Diplom_Project/pkg/mms"
	sms "Diplom_Project/pkg/sms"
	"Diplom_Project/pkg/support"
	"Diplom_Project/pkg/voicecall"
	"fmt"
)

// ResultT - Структура с данными для вывода
type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

// ResultSetT - Структура данных всех сервисов
type ResultSetT struct {
	SMS       [][]sms.SMSData                `json:"sms"`
	MMS       [][]mms.MMSData                `json:"mms"`
	VoiceCall []voicecall.VoiceData          `json:"voice_call"`
	Email     map[string][][]email.EmailData `json: email"`
	Billing   billing.BillingData            `json: billing"`
	Support   []int                          `json: support"`
	Incidents []incident.IncidentData        `json:"incident"`
}

func PrintResultData() {
	arr := GetResultData()

	for _, data := range arr.Data.MMS {
		fmt.Println(data)

	}

}

func GetResultData() ResultT { // функция сбора данных со всех приложений

	sms, errSms := sms.ResultData()
	mms, errMms := mms.ResultData()
	voicecall, errVoice := voicecall.VoiceCallDataFile()
	email, errEmail := email.ResultData()
	billing, errBilling := billing.ResultData()
	support, errSupport := support.ResultData()
	incident, errIncident := incident.ResultData()

	resultErr := fmt.Sprint(errSms, errMms, errVoice, errEmail, errBilling, errSupport, errIncident)

	if errSms != nil || errMms != nil || errVoice != nil || errEmail != nil || errBilling != nil || errSupport != nil || errIncident != nil {
		return ResultT{Status: false, Data: ResultSetT{}, Error: resultErr}
	}

	return ResultT{
		Status: true,
		Data: ResultSetT{
			sms,
			mms,
			voicecall,
			email,
			billing,
			support,
			incident,
		},
		Error: "",
	}
}
