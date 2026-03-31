# GetMessages

[dev.max.ru/docs-api/methods/GET/messages](https://dev.max.ru/docs-api/methods/GET/messages)

Метод предназначен для получения списка сообщений по ID чата или точному списку ID сообщений.

## Запрос {#request}

### Поля запроса  {#request-parameters}

| Поле         | Тип          | Обязательный | Описание                                                                                  |
|--------------|--------------|--------------|-------------------------------------------------------------------------------------------|
| `ChatID`     | **integer**  | **Да***      | Идентификатор чата <br>*(*обязательно, если не указан `MessageIDs`)*                      |
| <nobr>`MessageIDs`</nobr>   | **array**    | **Да*** | Массив точных идентификаторов сообщений <br>*(*обязательно, если нет `ChatID`)* |
| `From`       | **integer**  | Нет          | Время начала для запрашиваемых сообщений (в формате Unix timestamp)                       |
| `To`         | **integer**  | Нет          | Время окончания для запрашиваемых сообщений (в формате Unix timestamp)                    |
| `Count`      | **integer**  | Нет          | Максимальное количество сообщений в ответе (по умолчанию 50)                              |

### Пример запроса {#request-example}  

**Получение всех сообщений чата:**

```go
response, err := bot.Messages.GetMessages(
    context.Background(), 
    models.GetMessagesReq{
        ChatID: 123456,
    }
)
```

**Получение списка выбранных сообщений:**

```go
response, err := bot.Messages.GetMessages(
    context.Background(), 
    models.GetMessagesReq{
        MessageIDs: [
            "mid.000000000782a4e0019d002654056aed",
            "mid.000000000782a4e0019d0004e58330de",
            "mid.000000000782a4e0019d0004e3424fbd"
        ]
    }
)
```

## Ответ {#response}

### Поля ответа {#response-parameters}

| Поле       | Тип        | Описание                                               |
|------------|------------|--------------------------------------------------------|
| `messages` | [**array**](../objects/Message.md)  |    Массив объектов сообщений  |


### Пример тела ответа {#response-example-body}

```json
{
            "recipient": {
                "chat_id": 123456789,
                "chat_type": "dialog",
                "user_id": 111222333
            },
            "timestamp": 1773835378770,
            "body": {
                "mid": "mid.000000000782a4e0019d00d3e0525a17",
                "seq": 116250075383093783,
                "text": "Hello!"
            },
            "sender": {
                "user_id": 987654321,
                "first_name": "Jane",
                "last_name": "",
                "is_bot": false,
                "last_activity_time": 1773838157000,
                "name": "Jane"
            }
        },
        {
            "recipient": {
                "chat_id": 123456789,
                "chat_type": "dialog",
                "user_id": 987654321
            },
            "timestamp": 1773824005332,
            "body": {
                "mid": "mid.000000000782a4e0019d002654d46aec",
                "seq": 116249330013465324,
                "text": "What's up?"
            },
            "sender": {
                "user_id": 111222333,
                "first_name": "GREEN-API test",
                "username": "id5047259512_1_bot",
                "is_bot": true,
                "last_activity_time": 1773838164201,
                "name": "GREEN-API test"
            }
        },
        {
            "recipient": {
                "chat_id": 123456789,
                "chat_type": "dialog",
                "user_id": 111222333
            },
            "timestamp": 1773824005125,
            "body": {
                "mid": "mid.000000000782a4e0019d002654056aed",
                "seq": 116249329999899373,
                "text": "Wanna hang out?"
            },
            "sender": {
                "user_id": 987654321,
                "first_name": "Jane",
                "last_name": "",
                "is_bot": false,
                "last_activity_time": 1773838157000,
                "name": "Jane"
            }
        }
```