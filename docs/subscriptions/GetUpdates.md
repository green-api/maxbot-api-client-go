# GetUpdates

[dev.max.ru/docs-api/methods/GET/updates](https://dev.max.ru/docs-api/methods/GET/updates)

Метод предназначен для получения входящих уведомлений

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле           | Тип         | Обязательный | Описание                                                                                                           |
|----------------|-------------|--------------|--------------------------------------------------------------------------------------------------------------------|
| `Limit`        | **integer** | Нет          | Максимальное количество обновлений для получения (по умолчанию 100)                                                |
| `Timeout`      | **integer** | Нет          | Тайм-аут в секундах для долгого опроса (по умолчанию 30)                                                           |
| `Marker`       | **integer** | Нет          | Если передан, бот получит обновления, которые еще не были получены. Если не передан, получит все новые обновления  |
| `Types`        | **string**  | Нет          | Список типов обновлений, которые бот хочет получить. Возможные значения:                                           |
||||                                           `bot_added`          - Бот добавлен в чат или группу                                                                |
||||                                           `bot_removed`        - Бот удален из чата или группы                                                                |
||||                                           `bot_started`        - Пользователь нажал кнопку «Старт» или возобновил работу с ботом                              |
||||                                           `bot_stopped`        - Пользователь заблокировал бота или остановил его работу                                      |
||||                                           `chat_title_changed` - Изменено название группового чата                                                            |
||||                                           `dialog_muted`       - Уведомления в диалоге отключены пользователем                                                |
||||                                           `dialog_unmuted`     - Уведомления в диалоге снова включены                                                         |
||||                                           `dialog_cleared`     - История сообщений в диалоге очищена, но сам чат не удален                                    |
||||                                           `dialog_removed`     - Диалог полностью удален из списка чатов                                                      |
||||                                           `message_created`    - Получено новое входящее сообщение                                                            |
||||                                           `message_callback`   - Обработка нажатия на инлайн-кнопку                                                           |
||||                                           `message_edited`     - Ранее отправленное сообщение было изменено                                                   |
||||                                           `message_removed`    - Сообщение было удалено из истории переписки                                                  |
||||                                           `user_added`         - В чат добавлен новый участник                                                                |
||||                                           `user_removed`       - Участник покинул чат или был из него удален                                                  |

### Пример запроса {#request-example}

```go
response, err := bot.Updates.GetUpdates(context.Background(), nil)
```

**Пример запроса с параметрами:**

```go
response, err := bot.Updates.GetUpdates(
    context.Background(), 
    models.GetUpdatesReq{
        Limit: 10,
        Timeout: 10,
        Types: [
            "bot_added",
            "bot_removed",
            "bot_started",
            "bot_stopped"
        ]
    }
)
```

## Ответ {#response}

### Поля ответа {#response-parameters}

| Поле      | Тип         | Описание                                         |
|-----------|-------------|--------------------------------------------------|
| `updates` | [**array**](../objects/Update.md)  | Массив обновлений         |
| `marker`  | **integer** | Указатель на следующую страницу данных           |


### Пример тела ответа {#response-example-body}

```json
{
    "updates": [
        {
            "message": {
                "recipient": {
                    "chat_id": 123456789,
                    "chat_type": "dialog",
                    "user_id": 987654321
                },
                "timestamp": 1773835382644,
                "body": {
                    "mid": "mid.000000000782a4e0019d00d3ef744e91",
                    "seq": 116250075636977297,
                    "text": "Hello"
                },
                "sender": {
                    "user_id": 111222333,
                    "first_name": "Jane",
                    "last_name": "",
                    "is_bot": false,
                    "last_activity_time": 1773835377000,
                    "name": "Jane"
                }
            },
            "timestamp": 1773835382644,
            "user_locale": "ru",
            "update_type": "message_created"
        },
        {
            "callback": {
                "timestamp": 111222333,
                "callback_id": "f9LHodD0cOLW7qZQo5Yp4sWNbFSb7DnBL1K2N5O5vMYCShXQUyMx0IUn",
                "user": {
                    "user_id": 19706786,
                    "first_name": "Jane",
                    "last_name": "",
                    "is_bot": false,
                    "last_activity_time": 1773836012000
                },
                "payload": "/Action1"
            },
        }
    ],
    "marker": 44510937
}
```