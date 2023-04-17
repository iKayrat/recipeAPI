# Рецепт приложения REST API с Golang и PostgreSQL с использованием Docker

Это пример REST API приложения рецепта, созданного с помощью Golang и PostgreSQL, и его можно легко настроить и запустить с помощью Docker Compose для легкой разработки и развертывания.

## Введение
В этом проекте я решил испытать себя и использовать впервые фреймворк Fiber. 
Fiber - это быстрый и легкий веб-фреймворк для Go, который разработан с учетом простоты использования и эффективности.
Как разработчик, уже знакомый с другими веб-фреймворками на языке Go, я хотел расширить свой набор навыков и исследовать новые варианты. Fiber заинтриговал меня своим обещанием улучшенной производительности и простым и интуитивным API.
Мой опыт работы с базой данных над запросами фильтрации ингредиентов и времени приготовления, также загвоздки при конфигурация Docker файлов, над решением которого понадобился целый день, был положительным и подтвердил мою способность адаптироваться, уделять внимание малозначительным деталям и изучать новые подходы для веб-разработки. Теперь я более уверен в своей способности работать с различными технологиями, и с нетерпением жду новых вызовов в будущих проектах.

## Функционал
- Создание, чтение, обновление и удаление рецептов (CRUD)
- Хранение и возможность извлекать рецепты в базе данных PostgreSQL
- Фильтр по списку ингредиентов.
- Аутентификация и авторизация с помощью JWT
- Docker-compose для простоты разработки и развертывания

Выполните следующие шаги, чтобы настроить и запустить приложение рецепта с помощью docker-compose:

1. Клонируйте репозиторий на свой локальный компьютер:
```
git clone https://github.com/iKayrat/recipeAPI.git
```

2. Создайте файл **.env** в корневом каталоге проекта со следующими переменными среды:
```
DBSOURCE=postgresql://user:password@localhost:5432/dbname?sslmode=disable
SECRET_KEY=secret_key
```
(примечание: переменные среды уже имеются))

3. Запустите контейнеры Docker с помощью docker-сompose:
```
docker-compose up -d
```
или
```
make run
```

Это команда запустит контейнер PostgreSQL на **порту :5432**, создаст необходимые таблицы и запустит Golang REST API (**на порту:8000**).

4. Получите доступ к API по адресу http://localhost:8080/ в веб-браузере или через клиент API, например Postman.

## API запросы

- `POST /register` :Зарегистрировать нового пользователя
- `POST /login`    :Войти существующим пользователем и получите токен JWT
- `POST /logout`   :Выйти
    
- `POST /recipes` :Записать новый рецепт
- `GET /recipes/all` :Получить список всех рецептов
- `GET /recipes/id/:id` :Получить рецепт по ID
- `POST /recipes/?ingredients=sugar,milk` :Получить рецепты по ингредиентам
- `POST /recipes/time` :Получить рецепты по общему времени принготовления
- `PUT /recipes/:id` :Обновить рецепт по ID
- `DELETE /recipes/:id` :Удалить рецепт по ID

*ссылка для тестипрования на Postman:
https://elements.getpostman.com/redirect?entityId=14424408-9114de86-8e8c-4928-9aab-9cfe61d80891&entityType=collection


## Контакты

Если у вас есть какие-либо вопросы или отзывы, пожалуйста, не стесняйтесь обращаться ко мне по адресу ik.kairat@gmail.com. Я был бы рад услышать от вас и помочь с любыми запросами.