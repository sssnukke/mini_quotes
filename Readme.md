# Цитатник - REST API сервис на Go

Мини-сервис для хранения и управления цитатами.

## Установка и запуск

1. Клонируйте репозиторий:
   ```bash
   git clone git@github.com:sssnukke/mini_quotes.git

## Запуск через Docker

   1. Убедитесь, что у вас установлен Docker
   2. Выполните в терминале:

   ```bash
   # Сборка образа
   docker build -t mini_quotes .
   
   # Запуск контейнера
   docker run -dp 8080:8080 quotebook