<h1 align="center"> :closed_lock_with_key: Telegram-бот для хранения паролей </h1>

<p align="center"> Проект разработан в рамках прохождения отбора на стажировку VK "Разработчик системы кластеризации Tarantool" </p>

<hr>

## Навигация

* [Техническое задание](#chapter-0)
* [Описание разработанного решения](#chapter-1)
* [Запуск](#chapter-2)

<a id="chapter-0"></a>

## :page_with_curl: Техническое задание

Реализовать Telegram бота, который обладает функционалом персонального хранилища паролей.

Должны быть поддержаны следующие команды (как именно они будут работать, остается на ваше усмотрение):

1. `/set` - **добавляет логин и пароль к сервису**
2. `/get` - **получает логин и пароль по названию сервиса**
3. `/del` - **удаляет значения для сервиса**

### Требования к реализации:

1. Бот должен быть написан на одном из трех языков на выбор: Golang,
   Python, Lua
2. Чтобы обеспечить небольшую безопасность, пароли не должны
   оставаться в чате долго, бот должен их удалять по истечении некоторого
   времени.
3. Поэтому для каждого из пользователей должно быть свое
   пространство.

### Будет плюсом:

1. Развернуть бота в облаке и приложить на него ссылку.
2. Использовать Docker
3. Мы не хотим, чтобы наши пользователи расстроились работой нашего
   сервиса, поэтому если уборщица выдернула провод из розетки нашего
   сервера, мы бы не хотели потерять данные.

<a id="chapter-1"></a>

## :page_facing_up: Описание разработанного решения

:heavy_exclamation_mark: **Бот запущен и доступен по адресу: https://t.me/yuleo_password_bot**

Решение разработано с использованием языка **Go** и библиотеки [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api).

Структура проекта основана на [go-clean-template](https://github.com/evrone/go-clean-template).

В качестве СУБД выбрана **PostgreSQL**.

При запуске исполняемого файла первым аргументом командрой строки передаётся путь до файла конфигурации `.env` (аналогичного файлу `example.env`). 

<a id="chapter-2"></a>

## :zap: Запуск

0. Создать файл конфигурации `.env` (аналогичный файлу `example.env`) и задать в нем значения параметров.

1. В docker-контейнере

```bash 
docker-compose up --build
```

Аналогично с использованием `make`:
```bash 
make up
```

2. Локальный запуск

```bash 
docker-compose up postgres
docker-compose up init-db
go run ./... ./dev.env 
```
