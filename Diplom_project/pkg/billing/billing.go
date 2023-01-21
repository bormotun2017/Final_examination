package billing

import (
	"Diplom_Project/pkg/utils"
	"fmt"
	"io/ioutil"
	"log"
)

const filename = "../Simuliator/skillbox-diploma/billing.data"

type BillingData struct { // структура для Биллинга
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func BillingDataFile() (BillingData, error) {

	file, err := ioutil.ReadFile(filename) // читаем файл
	if err != nil {
		fmt.Println("ошибка чтения billing.data")
		log.Fatalln(err)
	}

	var resultBilling BillingData
	//arrBilling := make([]BillingData, 0, 0)

	resultBilling.CreateCustomer = utils.ByteToBool(file[5])
	resultBilling.Purchase = utils.ByteToBool(file[4])
	resultBilling.Payout = utils.ByteToBool(file[3])
	resultBilling.Recurring = utils.ByteToBool(file[2])
	resultBilling.FraudControl = utils.ByteToBool(file[1])
	resultBilling.CheckoutPage = utils.ByteToBool(file[0])

	//arrBilling = append(arrBilling, resultBilling)

	return resultBilling, err
}

func PrintBilling(arrBilling BillingData) {
	file, err := ioutil.ReadFile(filename) // читаем файл
	if err != nil {
		fmt.Println("ошибка чтения billing.data")
		log.Fatalln(err)
	}
	fmt.Println(file)
	fmt.Println(arrBilling)
}

func ResultData() (BillingData, error) { // функция сбора данных для глобальной структуры
	arrResult, err := BillingDataFile()

	if err != nil {
		fmt.Println("Ошибка валидации данных Email")
		log.Fatalln(err)
	}

	return arrResult, err
}
