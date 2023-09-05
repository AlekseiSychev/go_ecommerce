# Mini e-commerce app
Проект использует базу данных PostgreSQL и включает следующие таблицы:

- Users: Хранит информацию о пользователях, такую как их идентификационные данные и контактную информацию.

- Items: Содержит информацию о продуктах, доступных для покупки.

- Orders: Хранит информацию о заказах, включая товары, которые были приобретены в каждом заказе.

Для каждой из таблиц реализованы CRUD-операции.

## Запуск приложения с Docker и Docker Compose
Проект включает файл docker-compose.yml, который позволяет легко развернуть приложение с использованием Docker и Docker Compose. Для запуска приложения выполните следующие шаги:

1. Убедитесь, что у вас установлен Docker и Docker Compose.

2. Склонируйте репозиторий проекта на свой компьютер.

3. В командной строке перейдите в корневую папку проекта.

4. Чтобы создать и запустить приложения в контейнерах используйте команду: 
```
docker-compose up --build
```


#### После успешного запуска контейнеров, приложение должно быть доступно по адресу http://localhost:8080/user , где 8080 - это порт, указанный в docker-compose.yml.

## Документация swagger доступна по адресу 

- http://localhost:8080/docs/index.html#


## Завершение работы
Чтобы остановить и удалить контейнеры, выполните следующую команду в корневой папке проекта:

```
docker-compose down
```