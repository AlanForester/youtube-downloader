package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Получаем URL видео из аргументов командной строки
	videoURL := "https://www.youtube.com/watch?v=kzcgel0ddHI" // URL по умолчанию
	if len(os.Args) > 1 {
		videoURL = os.Args[1]
	}
	outputFile := "output_video.mp4"

	// Проверяем, установлен ли yt-dlp
	_, err := exec.LookPath("yt-dlp")
	if err != nil {
		fmt.Println("yt-dlp не установлен. Установите его с помощью команды:")
		fmt.Println("brew install yt-dlp")
		os.Exit(1)
	}

	fmt.Println("Начинаем загрузку видео...")

	// Формируем команду для скачивания
	cmd := exec.Command("yt-dlp",
		"-f", "best[ext=mp4]", // Выбираем лучший формат mp4
		"-o", outputFile, // Имя выходного файла
		"--no-playlist", // Не скачивать плейлист
		"--no-warnings", // Отключаем предупреждения
		videoURL,
	)

	// Перенаправляем вывод команды в stdout и stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Запускаем команду
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Ошибка при скачивании видео: %v\n", err)
		os.Exit(1)
	}

	absPath, _ := filepath.Abs(outputFile)
	fmt.Printf("Видео успешно загружено: %s\n", absPath)
}
