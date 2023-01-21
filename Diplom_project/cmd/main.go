package main

import (
	"Diplom_Project/pkg/result"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	// Если хочется посмотреть результаты выполнения отдельных функция, нужно просто убрать комментарии
	// c нужного куска кода

	// S M S

	//	arr, err := sms.SmsDataFile()
	//	fmt.Println(err)
	//	arrResult := sms.SmsDataSort(arr)
	//	sms.SmsPrint(arrResult)

	// M M S

	//arrResult, err := mms.ResultData()
	//fmt.Println(err)
	//mms.MmsPrint(arrResult)

	//Voice Call

	//arrVoicecall, err := voicecall.VoiceCallDataFile()
	//fmt.Println(err)
	//voicecall.VoiceCallPrint(arrVoicecall)

	// E M A I L

//	arrEmail, err := email.ResultData()
//	fmt.Println(err)
//	email.EmailPrint(arrEmail)

	// B I L L I N G
	//arrBilling, err := billing.BillingDataFile()
	//fmt.Println(err)
	//billing.PrintBilling(arrBilling)

	//S U P P O R T
	//support.SupportPrint()

	// I N C I D E N T
	//incident.IncidentPrint()

}

func main() {

	r := mux.NewRouter()                        // задаем роутер
	r.HandleFunc("/", Logger(handleConnection)) // завернули хэндлфунк в логгер

	s := &http.Server{ //наш сервер с настройками
		Addr:           "localhost:8585", // ,будет на любом ip с этим портом например 127.0.0.1:8080
		Handler:        r,                //указываем хэндлер
		ReadTimeout:    10 * time.Second, //всякие полезные таймы обработки
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20, //1*2 ^20 - 128 kb
	}

	go func() { //запускаем сервер в отдельной го рутине
		log.Printf("Listening on http://%s\n", s.Addr)
		log.Fatal(s.ListenAndServe()) // слушаем сервер
	}()

	graceful(s, 5*time.Second) // в течении 5 секунд еще обрабатываем входящие запросы, потом дообрабатываем все что собрали

}

// handleConnection Возврвт ответа
func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res []byte
	res, _ = json.Marshal(result.GetResultData())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)

	if err != nil {
		log.Fatal(err)
	}
}

func graceful(hs *http.Server, timeout time.Duration) { // graceful shutdown
	stop := make(chan os.Signal, 1) // создаем канал

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM) // завершит канал когда получит сигнал

	<-stop // получили сигнал

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Printf("nShutdown with timeout: %s\n", timeout)

	if err := hs.Shutdown(ctx); err != nil {
		log.Printf("Error: %v\n", err)

	} else {
		log.Println("Server stopped")
	}

}

func Logger(next http.HandlerFunc) http.HandlerFunc { //логгер который будет перехватывать midlewear
	return func(w http.ResponseWriter, r *http.Request) { //возвращаем функцию
		w.Header().Set("Content-Type", "application/json") // обясняем хэдэру, что будем работать с json

		log.Printf("server [net/http] method [%s] connecion from [%v]", r.Method, r.RemoteAddr)

		next.ServeHTTP(w, r)
	}
}
