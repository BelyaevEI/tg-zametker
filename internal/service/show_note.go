package service

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/BelyaevEI/tg-zametker/internal/utils"
)

func (s *serv) showNotes(userID int64) (string, error) {
	var buf bytes.Buffer

	notes, err := s.repository.ShowNotes(userID)
	if err != nil {
		return "", err
	}

	// Создаем writer с табуляцией
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	// Определяем разделитель
	separator := "------------------------------------------"

	// Печатаем разделитель перед таблицей
	fmt.Fprintln(w, separator)

	// Сформируем таблицу с заголовком без вертикальных разделителей
	fmt.Fprintln(w, "№   \tЗаметка")
	fmt.Fprintln(w, separator)

	// Печатаем строки с данными, без вертикальных разделителей
	for i, v := range notes {
		// Разбиваем длинные строки на части (максимум 40 символов в строке)
		wrappedNote := utils.WrapText(v, 40)
		lines := strings.Split(wrappedNote, "\n")
		for j, line := range lines {
			if j == 0 {
				// Для первой строки выводим номер и заметку
				fmt.Fprintf(w, "%-3v \t%-40v\n", strconv.Itoa(i+1)+".", line)
			} else {
				// Для последующих строк добавляем пустую колонку для номера
				fmt.Fprintf(w, "    \t%-40v\n", line)
			}
		}
		fmt.Fprintln(w, separator) // Разделитель между строками
	}

	w.Flush()

	return buf.String(), nil
}
