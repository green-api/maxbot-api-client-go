# EditMessage

[dev.max.ru/docs-api/methods/PUT/messages](https://dev.max.ru/docs-api/methods/PUT/messages)

Метод предназначен для редактирования текста или кнопок ранее отправленного сообщения.

## Запрос {#request}

### Поля запроса  {#request-parameters}

| Поле          | Тип          | Обязательный   | Описание                                                                                 |
|---------------|--------------|----------------|------------------------------------------------------------------------------------------|
| `MessageID`   | **string**   | Да             | ID редактируемого сообщения                                                              |
| `Text`        | **string**   | Нет            | Новый текст сообщения (до 4000 символов)                                                 |
| <nobr>`Attachments`</nobr> | **array** | Нет  | Вложения сообщения. Если пусто, все вложения будут удалены                               |
| `Link`        | [**object**](../objects/NewMessageBody.md#поля-объекта-newmessagelinknewmessagelink) | Нет          | Ссылка на другое сообщение      |
| `Notify`      | **boolean**  | Нет            | Если `false`, участники чата не будут уведомлены (по умолчанию `true`)                   |
| `Format`      | **string**   | Нет            | Если установлен, текст сообщения будет форматирован данным способом (`html`, `markdown`) |

### Пример запроса {#request-example}  

```go
response, err := bot.Messages.EditMessage(
    context.Background(), 
    models.EditMessageReq{
        MessageID: "mid.000000000782a4e0019d00d3ef744e91",
        Text:      "Updated message text!",
        Notify: false
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
    "success": false,
    "message": "error.edit.wrong.author"
}
```