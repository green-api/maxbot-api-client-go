# GetMessage

[dev.max.ru/docs-api/methods/GET/messages/-messageId-](https://dev.max.ru/docs-api/methods/GET/messages/-messageId-)

Метод предназначен для получения полной информации о конкретном сообщении по его ID.

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле        | Тип        | Обязательный | Описание                                                  |
|-------------|------------|--------------|-----------------------------------------------------------|
| `MessageID` | **string** | Да           | ID сообщения (mid), чтобы получить одно сообщение в чате  |

### Пример запроса {#request-example}  

```go
response, err := bot.Messages.GetMessage(
    context.Background(), 
    models.GetMessageReq{
        MessageID: "mid:987654321",
    }
)
```

## Ответ {#response}

[**объект `Message`**](../objects/Message.md)

| Поле            | Тип                                                    | Описание                                                                       |
|-----------------|--------------------------------------------------------|--------------------------------------------------------------------------------|
| `sender`        | [**object**](../objects/User.md)                       | Пользователь, отправивший сообщение                                            |
| `recipient`     | [**object**](../objects/Message.md#поля-объекта-recipient-recipient)          | Получатель сообщения. Может быть пользователем или чатом   |
| `timestamp`     | **integer**                                            | Время создания сообщения в формате Unix-time                                   |
| <nobr>`linked_message`</nobr> | [**object**](../objects/Message.md#поля-объекта-linkedmessage-linkedmessage) | Пересланное или ответное сообщение       |
| `body`          | [**object**](../objects/Message.md#поля-объекта-messagebody-messagebody)        | Содержимое сообщения        |
| `stat`          | [**object**](../objects/Message.md#поля-объекта-messagestat--messagestat) | Статистика сообщения. Возвращается только для постов в каналах |
| `url`           | **string**                                             | Публичная ссылка на пост на канале. Отсутствует для диалогов и групповых чатов  |


### Пример тела ответа {#response-example-body}

```json
{
    "recipient": {
        "chat_id": 123456789,
        "chat_type": "dialog",
        "user_id": 111222333
    },
    "timestamp": 1773837299858,
    "body": {
        "mid": "mid.000000000782a4ed00f130920001920e",
        "seq": 116250201283494414,
        "text": "Hello world!"
    },
    "sender": {
        "user_id": 9876543210,
        "first_name": "Jane",
        "last_name": "",
        "is_bot": false,
        "last_activity_time": 1773837347000
    }
}
```
