package service

import "strconv"

func (s *serv) showNotes(userID int64) (string, error) {
	var msg string
	notes, err := s.repository.ShowNotes(userID)
	if err != nil {
		return "", err
	}

	//Сформируем таблицу с заголовком и индексами
	msg = "№" + " | Текст заметки\n"
	for i, v := range notes {
		msg += strconv.Itoa(i+1) + " | " + v + "\n"
	}
	return msg, nil
}
