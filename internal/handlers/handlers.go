package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

 var Logger *log.Logger
/////////////////////////////////////////////////////////////////////////////////////////////////
//Реализация хэндлера по запросу "/"/////
func Handler1(res http.ResponseWriter, req *http.Request) {
//открыть файл 
data, err := os.ReadFile("../index.html")
    if err != nil {
        Logger.Fatal(err)
    }

//возвращается строка в res.Write([]byte(s))
res.Write([]byte(data))
}
///////////////////////////////////////////////////////////////////////////////////////////////////
// Реализация хэндлера по запросу "/upload"
func Handler2(res http.ResponseWriter, req *http.Request) {
// Парсинг формы////
       if err := req.ParseForm(); err != nil {
        http.Error(res,"Ошибка при парсинге формы", http.StatusInternalServerError)
        Logger.Fatal(err)
        return
    }
// Получение файла из формы////////////////////////////////////////////////////////////////////////////
    file,header, err := req.FormFile("myFile") // "myFile" - название поля файла в форме
    if err != nil { 
        http.Error(res, "Не удалось получить файл", http.StatusInternalServerError)
        Logger.Fatal(err)
        return
    }
    defer file.Close() //Закрываем файл/////
// Чтение данных из файла////////////////////
    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(res, "Ошибка при чтении файла", http.StatusInternalServerError)
        Logger.Fatal(err)
        return
    }
//  конвертация строки//////////////////
   convString:=service.Texttomorse_and_revers(string(data))

// Создаем новый файл///////////////////
    now := time.Now()
    fmt.Println(now)
    currentTime1 := now.String()[:11]
    currentTime2 := strings.ReplaceAll(now.String()[11:37], ":", "-", )
    fileExtantion := filepath.Ext(header.Filename)
    file2, err := os.Create(currentTime1+currentTime2+fileExtantion)       
    if err != nil {
        http.Error(res, "Ошибка при создании файла", http.StatusInternalServerError)
        Logger.Fatal(err)
    }
//Записываем результат конвертации строки в файл file2////
    _, err = file2.WriteString(convString)
    if err != nil {
         http.Error(res, "Ошибка при записи результата конвертации строки в файл", http.StatusInternalServerError)
        log.Fatal(err)
    }
defer file2.Close()
    _,err=res.Write([]byte(convString))
    if err !=   nil {
        http.Error(res, "Ошибка при передаче файла", http.StatusInternalServerError)
        Logger.Fatal(err)
    }
}