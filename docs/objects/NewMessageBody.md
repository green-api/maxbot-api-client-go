# NewMessageBodyy  

[dev.max.ru/docs-api/objects/NewMessageBody](https://dev.max.ru/docs-api/objects/NewMessageBody)

## Поля объекта {#NewMessageBody}  

Массив объектов с полями:

| Поле          | Тип                                    | Описание                                                                                 |
|---------------|----------------------------------------|------------------------------------------------------------------------------------------|
| `text`        | **string**                             | Новый текст сообщения (до 4000 символов)                                                 |
| <nobr>`attachments`</nobr> | [**array**](./Message.md#Attachments)  | Вложения сообщения. Если пусто, все вложения будут удалены                  |
| `link`        | [**object**](#NewMessageLink)          | Ссылка на другое сообщение                                                               |
| `notify`      | **boolean**                            | Если `false`, участники чата не будут уведомлены (по умолчанию `true`)                   |
| `format`      | **string**                             | Если установлен, текст сообщения будет форматирован данным способом (`html`, `markdown`) |


### Поля объекта `NewMessageLink`{#NewMessageLink}

| Поле   | Тип        | Обязательный | Описание                                                |
|--------|------------|--------------|---------------------------------------------------------|
| `type` | **string** | Нет          | Тип ссылки сообщения (`reply`, `forward`)               |
| `mid`  | **string** | Нет          | ID сообщения исходного сообщения                        |


## Пример объекта {#NewMessageBody-example}  

```json  
    {
        "text": "string",
        "attachments": [{
            "type": "image",
            "payload": {
                "url": "string",
                "token": "string",
                "photos": {
                    "token": "string"
                }
            }
        }],
        "link": {
            "type": "forward",
            "mid": "string"
        },
        "notify": true,
        "format": "markdown"
    }
```