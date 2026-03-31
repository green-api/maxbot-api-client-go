# GetChat

[dev.max.ru/docs-api/methods/GET/chats](https://dev.max.ru/docs-api/methods/GET/chats)

Возвращает информацию о групповом чате по его ID

## Запрос {#request}

### Поля запроса  {#request-parameters}

| Поле     | Тип         | Обязательный | Описание                        |
|----------|-------------|--------------|---------------------------------|
| `ChatID` | **integer** | Да           | ID запрашиваемого чата          |

### Пример запроса {#request-example}  

```go
response, err := bot.Chats.GetChat(
    context.Background(), 
    &models.GetChatReq{
        ChatID: 123456789,
    }
)
```

## Ответ {#response}

### Поля ответа  {#response-parameters}

[**Объект `Chat`**](../objects/Chat.md)

| Поле                 | Тип                            | Описание                                                               |
|----------------------|--------------------------------|------------------------------------------------------------------------|
| `chat_id`            | **integer**                    | ID чата                                                                |
| `type`               | **string**                     | Для групп принимает значение `chat`                                    |
| `status`             | **string**                     | Статус чата                                                            |
| `title`              | **string**                     | Обновленное название чата (если было изменено)                         |
| `icon`               | [**object**](../objects/Chat.md#поля-объекта-imageimage)           | Иконка чата                                          |
| `last_event_time`    | **integer**                    | Время последнего события в чате                                        |
| `participants_count` | **integer**                    | Количество участников чата. Для диалогов всегда 2                      |
| `owner_id`           | **integer**                    | ID владельца чата                                                      |
| `participants`       | **object**                     | Участники чата с временем последней активности                         |
| `is_public`          | **boolean**                    | Доступен ли чат публично                                               |
| `link`               | **string**                     | Ссылка на чат                                                          |
| `description`        | **string**                     | Описание чата                                                          |
| `dialog_with_user`   | [**object**](../objects/Chat.md#поля-объекта-dialogwithuser-dialogwithuser)  | Данные о пользователе в диалоге                      |
| `chat_message_id`    | **string**                     | ID сообщения, содержащего кнопку, через которую был инициирован чат    |
| `pinned_message`     | [**object**](../objects//Message.md#message)     | Закреплённое сообщение в чате                                |


### Пример тела ответа {#response-example-body}

**Успех:**

```json
{
    "chat_id": -722703071564882,
    "type": "chat",
    "status": "active",
    "title": "Group chat",
    "last_event_time": 1774253213015,
    "participants_count": 3,
    "is_public": false,
    "owner_id": 12345789,
    "participants": {
        "123456789": 1774253340181,
        "987654321": 0,
        "0123456789": 1773988397845
    },
    "messages_count": 4
}
```
