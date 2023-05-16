package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//а - алфавит, т.е. часть слова, а иногда и само слово, например: "в".
	a := []string{"А", "а", "Б", "б", "В", "в", "Г", "г", "Д", "д", "Е", "е", "Ё", "ё", "Ж", "ж", "З", "з", "И", "и", "Й", "й", "К", "к", "Л", "л", "М", "м", "Н", "н", "О", "о", "П", "п", "Р", "р", "С", "с", "Т", "т", "У", "у", "Ф", "ф", "Х", "х", "Ц", "ц", "Ч", "ч", "Ш", "ш", "Щ", "щ", "Ъ", "ъ", "Ы", "ы", "Ь", "ь", "Э", "э", "Ю", "ю", "Я", "я"}

	//Читаем файл.
	//Данный код очень часто встречающийся паттерн, для предотвращения ошибок. В случае такой ошибки, компилятор не прекратит работу, сообщив тип ошибки.
	file, err := os.ReadFile("words.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	//Изначально file - является срезом байт.
	fmt.Println("Срез байт: ")
	fmt.Println(file)

	//Конвертируем срез байт в тип string
	content := string(file)

	//С помощью TrimSpace удаляем пробелы в начале и в конце, а Split - "разрезает" string по символам и записывает их в срез.
	contentSliceByEverySymbol := strings.Split(strings.TrimSpace(content), "")
	fmt.Println("\nСрез символов по отдельности:")
	fmt.Println(contentSliceByEverySymbol)

	var contentSliceByWord []string
	maybeWord := ""
	enabledToAdd := false
	//Данный цикл пройдется по каждому символу. Цикл найдет все группировки букв и добавит их в срез.
	for _, symbol := range contentSliceByEverySymbol {
		//Если символ буква, то добавляем в актуальное слово и разрешаем его вставить.
		if contains(a, symbol) == true {
			enabledToAdd = true
			maybeWord = maybeWord + symbol
		} else {
			//Если символ не буква, то вставляем актуальное слово (если это доступно), обнуляем актуальное слово, запрещаем вставлять
			if enabledToAdd == true {
				contentSliceByWord = append(contentSliceByWord, maybeWord)
			}
			enabledToAdd = false
			maybeWord = ""
		}
	}
	fmt.Println("\nСрез символов по словам:")
	fmt.Println(contentSliceByWord)

}

// Простая функция, проверяющая находится ли строка в срезе.
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
