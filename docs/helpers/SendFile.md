# SendFile

Универсальный метод для упрощенной отправки файлов

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле          | Тип        | Обязательный | Описание                                                                                          |
|---------------|------------|--------------|---------------------------------------------------------------------------------------------------|
| `FileSource`  | **string** | Да           | URL-адрес файла (начинается с http/https) или путь к локальному файлу                             |
| `ChatID`      | **integer**  | Нет*         | ID чата (обязательно, если не передан `UserID`)                                                   |
| `DisableLinkPreview` | **boolean**  | Нет | Если `false`, сервер не будет генерировать превью для ссылок в тексте сообщения                   |
| `UserID`      | **integer**  | Нет*         | ID пользователя (обязательно, если не передан `ChatID`)                                           |
| `Text`        | **string** | Нет          | Новый текст сообщения (до 4000 символов)                                                          |
| `Attachments` | **array**  | Нет          | Вложения сообщения. Если пусто, все вложения будут удалены                                        |
| `Link`        | [**object**](../objects/NewMessageBody.md#поля-объекта-newmessagelinknewmessagelink) | Нет          | Ссылка на другое сообщение  |
| `Notify`      | **boolean**   | Нет          | Отправлять ли пуш-уведомление пользователю (`true`/`false`)                                       |
| `Format`      | **string** | Нет          | Формат разметки текста (`HTML`, `Markdown`)                                                       |


### Пример запроса {#request-example}

**Пример отправки локального файла:**

```go
response, err := bot.Helpers.SendFile(context.Background(), models.SendFileReq{
    ChatID:     123456789,
    Text:       "Вот отчет за этот месяц",
    FileSource: "./reports/march.pdf", 
})
```

```go
response, err := bot.Helpers.SendFileByUpload(context.Background(), models.SendFileReq{
    UserID:     55555,
    Text:       "Голосовое сообщение",
    FileSource: "./media/voice_memo.ogg",
})
```

**Пример отправки файла по ссылке:**

```go
response, err := bot.Helpers.SendFile(context.Background(), models.SendFileReq{
    UserID:     987654321,
    Text:       "Зацени картинку!",
    FileSource: "https://example.com/image.png", 
})
```

```go
response, err := bot.Helpers.SendFileByUrl(context.Background(), models.SendFileReq{
    ChatID:     123456789,
    Text:       "Красивый пейзаж",
    FileSource: "https://example.com/nature.jpg?width=1080",
})
```

## Ответ {#response}

### Поля ответа (объект `Message`) {#response-parameters}

| Поле        | Тип        | Описание                                                                          |
|-------------|------------|-----------------------------------------------------------------------------------|
| `Message`   | **object** | Закреплённое сообщение - [объект `Message`](../objects/Message.md) |

### Пример тела ответа {#response-example-body}

```json
{
    "sender": {
        "user_id": 111222,
        "first_name": "MyBot",
        "is_bot": true
    },
    "recipient": {
        "chat_id": 123456789,
        "chat_type": "group"
    },
    "timestamp": 1679051234,
    "body": {
        "mid": "msg-111222333",
        "seq": 42,
        "text": "Вот отчет за этот месяц",
        "attachments": [
            {
                "type": "file",
                "payload": {
                    "token": "file_token_xyz",
                    "filename": "march.pdf"
                }
            }
        ]
    }
}
```
