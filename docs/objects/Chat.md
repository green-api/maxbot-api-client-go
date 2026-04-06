# Chat  

[dev.max.ru/docs-api/objects/Chat](https://dev.max.ru/docs-api/objects/Chat)

Объект содержит общую информацию о групповом чате  

## Поля объекта {#Chat}  

| Поле                 | Тип                            | Описание                                                               |
|----------------------|--------------------------------|------------------------------------------------------------------------|
| `chat_id`            | **integer**                    | ID чата                                                                |
| `type`               | **string**                     | Для групп принимает значение `chat`                                    |
| `status`             | **string**                     | Статус чата. Возможные значения:                                       |
|                      |                                | `active` — Бот является активным участником чата.                      |
|                      |                                | `removed` — Бот был удалён из чата.                                    |
|                      |                                | `left` — Бот покинул чат.                                              |
|                      |                                | `closed` — Чат был закрыт.                                             |
| `title`              | **string**                     | Обновленное название чата (если было изменено)                         |
| `icon`               | [**object**](#Image)           | Иконка чата                                                            |
| `last_event_time`    | **integer**                    | Время последнего события в чате                                        |
| <nobr>`participants_count`</nobr> | **integer**       | Количество участников чата. Для диалогов всегда 2                      |
| `owner_id`           | **integer**                    | ID владельца чата                                                      |
| `participants`       | **object**                     | Участники чата с временем последней активности                         |
| `is_public`          | **boolean**                    | Доступен ли чат публично                                               |
| `link`               | **string**                     | Ссылка на чат                                                          |
| `description`        | **string**                     | Описание чата                                                          |
| `dialog_with_user`   | [**object**](#DialogWithUser)  | Данные о пользователе в диалоге                                        |
| `chat_message_id`    | **string**                     | ID сообщения, содержащего кнопку, через которую был инициирован чат    |
| `pinned_message`     | [**object**](./Message.md)     | Закреплённое сообщение в чате                                          |


### Поля объекта `Image`{#Image}

| Поле          | Тип        | Описание                   |
|---------------|------------|----------------------------|
| `url`         | **string** |  URL изображения           |


### Поля объекта `DialogWithUser` {#DialogWithUser}

| Поле                 | Тип         | Описание                                                                             |
|----------------------|-------------|--------------------------------------------------------------------------------------|
| `user_id`            | **integer** | Идентификатор пользователя или бота                                                  |
| `first_name`         | **string**  | Отображаемое имя пользователя или бота                                               |
| `username`           | **string**  | Никнейм бота или уникальное публичное имя пользователя                               |
| `is_bot`             | **boolean** | `true`, если это бот                                                                 |
| <nobr>`last_activity_time`<nobr> | **integer** | Время последней активности пользователя или бота в MAX (Unix-время в миллисекундах)  |
| `description`        | **string**  | Описание пользователя или бота (до 16000 символов)                                   |
| `avatar_url`         | **string**  | URL аватара пользователя или бота в уменьшенном размере                              |
| `full_avatar_url`    | **string**  | URL аватара пользователя или бота в полном размере                                   |

## Пример объекта {#Chat-example}  

```json  
    {
        "chat_id": 0,
        "type": "chat",
        "status": "active",
        "title": "string",
        "icon": {
            "url": "string"
        },
        "last_event_time": 0,
        "participants_count": 0,
        "owner_id": 0,
        "participants": object,
        "is_public": true,
        "link": "string",
        "description": "string",
        "dialog_with_user": {
            "user_id": 0,
            "first_name": "string",
            "last_name": "string",
            "username": "string",
            "is_bot": true,
            "last_activity_time": 0,
            "name": "string",
            "description": "string",
            "avatar_url": "string",
            "full_avatar_url": "string"
        },
        "chat_message_id": "string",
        "pinned_message": {
            "sender": {
                "user_id": 0,
                "first_name": "string",
                "last_name": "string",
                "username": "string",
                "is_bot": true,
                "last_activity_time": 0,
                "name": "string"
            },
            "recipient": {
                "chat_id": 0,
                "chat_type": "chat",
                "user_id": 0
            },
            "timestamp": 0,
            "link": {
                "type": "forward",
                "sender": {
                    "user_id": 0,
                    "first_name": "string",
                    "last_name": "string",
                    "username": "string",
                    "is_bot": true,
                    "last_activity_time": 0,
                    "name": "string"
                },
                "chat_id": 0,
                "message": {
                    "mid": "string",
                    "seq": 0,
                    "text": "string",
                    "attachments": [{ ... }],
                    "markup": [{ ... }]
                }
            },
            "body": {
                "mid": "string",
                "seq": 0,
                "text": "string",
                "attachments": [{
                    "type": "image",
                    "payload": {
                        "photo_id": 0,
                        "token": "string",
                        "url": "string"
                    }
                }],
                "markup": [{
                    "type": "strong",
                    "from": 0,
                    "length": 0
                }]
            },
            "stat": {
                "views": 0
            },
            "url": "string"
        }
    }
```
