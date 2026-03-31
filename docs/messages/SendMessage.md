# SendMessage

[dev.max.ru/docs-api/methods/POST/messages](https://dev.max.ru/docs-api/methods/POST/messages)

Метод предназначен для отправки текстового сообщения или медиафайла пользователю или в чат.

## Запрос {#request}

### Поля запроса  {#request-parameters}

| Поле                 | Тип          | Обязательный    | Описание                                                                        |
|----------------------|--------------|-----------------|---------------------------------------------------------------------------------|
| `UserID`             | **integer**  | **Да***         | ID пользователя-получателя <br>* *обязателен, если не указан `ChatID`*          |
| `ChatID`             | **integer**  | **Да***         | ID чата-получателя <br>* *обязателен, если не указан `UserID`*                  |
| <nobr>`DisableLinkPreview`</nobr> | **boolean** | Нет | Если `false`, сервер не будет генерировать превью для ссылок в тексте сообщения |
| `Text`               | **string**   | Нет             | Новый текст сообщения (до 4000 символов)                                        |
| `Attachments`        | **array**    | Нет             | Вложения сообщения. Если пусто, все вложения будут удалены                      |
| `Link`               | [**object**](../objects/NewMessageBody.md#поля-объекта-newmessagelinknewmessagelink)  | Нет          | Ссылка на другое сообщение  |
| `Notify`             | **boolean**  | Нет             | Отправлять ли пуш-уведомление пользователю (`true`/`false`)                     |
| `Format`             | **string**   | Нет             | Формат разметки текста (`HTML`, `Markdown`)                                     |

### Пример запроса {#request-example}

```go
response, err := bot.Messages.SendMessage(
    context.Background(), 
    models.SendMessageReq{
        ChatID: 123456,
        Text:   "Hello, world!",
        Notify: true,
    }
)
```

## Ответ  {#response}

[**объект `Message`**](../objects/Message.md)

| Поле            | Тип                                                    | Описание                                                                       |
|-----------------|--------------------------------------------------------|--------------------------------------------------------------------------------|
| `sender`        | [**object**](../objects/User.md)                       | Пользователь, отправивший сообщение                                            |
| `recipient`     | [**object**](../objects/Message.md#поля-объекта-recipient-recipient)          | Получатель сообщения. Может быть пользователем или чатом  |
| `timestamp`     | **integer**                                            | Время создания сообщения в формате Unix-time                                   |
| <nobr>`linked_message`</nobr> | [**object**](../objects/Message.md#поля-объекта-linkedmessage-linkedmessage) | Пересланное или ответное сообщение    |
| `body`          | [**object**](../objects/Message.md#поля-объекта-messagebody-messagebody)        | Содержимое сообщения     |
| `stat`          | [**object**](../objects/Message.md#поля-объекта-messagestat--messagestat) | Статистика сообщения. Возвращается только для постов в каналах |
| `url`           | **string**                                             | Публичная ссылка на пост на канале. Отсутствует для диалогов и групповых чатов  |

### Пример тела ответа {#response-example-body}

```json
{
    "message": {
        "recipient": {
            "chat_id": 123456789,
            "chat_type": "dialog",
            "user_id": 111222333
        },
        "timestamp": 1773838491240,
        "body": {
            "mid": "mid.000000000782a4e0019d01035e6843dd",
            "seq": 116250279361922013,
            "text": "Hello world!"
        },
        "sender": {
            "user_id": 9876543210,
            "first_name": "John",
            "username": "id5047259512_1_bot",
            "is_bot": true,
            "last_activity_time": 1773838491258
        }
    }
}
```