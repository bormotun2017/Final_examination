# Final_examination by Egor Biziaew

Как запустить 

1. Нужно клонировать проект

   git@github.com:bormotun2017/Final_examination.git

   (необходимо чтобы папки Diplom_project и Simuliator были смежными для праильности прописания путей)

2. Запустить симулятор. go run Simuliator\skillbox-diploma\main.go

3. Запустить проект. go run Diplom_project\cmd\main.go

4. Возможно придется установить библиотеку gorilla/mux, если среда разработки ее не подключит автоматически

go get -u github.com/gorilla/mux

5. Результат работы программы увидим по адресу: `http://localhost:8585`
