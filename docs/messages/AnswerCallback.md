# AnswerCallback

[dev.max.ru/docs-api/methods/POST/answers](https://dev.max.ru/docs-api/methods/POST/answers)

Метод предназначен для подтверждения нажатия инлайн-кнопки пользователем (убирает индикатор загрузки на клиенте).

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле                         | Тип        | Обязательный | Описание                                                                          |
|------------------------------|------------|--------------|-----------------------------------------------------------------------------------|
| `CallbackID`                 | **string** | Да           | ID коллбэка, полученного из входящего события                                     |
| `Message` | [**object**](../objects/NewMessageBody.md) | Нет  | Обновление сообщения или клавиатуры в ответ на нажатие                       |
| <nobr>`Notification`</nobr>  | **string** | Нет          | Заполните это, если хотите просто отправить одноразовое уведомление пользователю  |

### Пример запроса {#request-example}  

```go
response, err := bot.Messages.AnswerCallback(
    context.Background(), 
    models.AnswerCallbackReq{
        CallbackID: "f9LHodD0cOLW7qZQo5Yp4sWNbFSb7DnBL1K2N5O5vMYCShXQUyMx0IUn",
        Message: &m.NewMessageBody{
            Text: "Action confirmed!",
        },
    }
)
```

## Ответ {#response}

### Поля ответа {#response-parameters}

| Поле      | Тип         | Описание                                                 |
|-----------|-------------|----------------------------------------------------------|
| `success` | **boolean** | `true`, если запрос был успешным,`false` — в противном случае                       |
| `code`    | **string**  | Код ошибки операции                                      |
| `message` | **string**  | Объяснительное сообщение, если результат не был успешным |

### Пример тела ответа {#response-example-body}

**Успех:** 

```json
{
    "success": true
}
```

**Ошибка:**

```json
{
    "code": "proto.payload",
    "message": "callback_id: Callback identifier is invalid"
}
```