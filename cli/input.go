package cli

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func acceptmessage() (target string) {
	// Создаем сканер, который читает из стандартного ввода (консоли)
	scanner := bufio.NewScanner(os.Stdin)

	// Ждем, пока пользователь введет текст и нажмет Enter
	if scanner.Scan() {
		// Записываем результат по адресу нашей переменной
		target = scanner.Text()
		// (Опционально) Удаляем случайные пробелы на концах строки
		target = strings.TrimSpace(target)
	}

	// Проверяем, не случилось ли ошибки при чтении
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка чтения: %v", err)
	}
	return target
}
