# Update  

[dev.max.ru/docs-api/objects/Update](https://dev.max.ru/docs-api/objects/Update)

Объект Update представляет различные типы событий, произошедших в чате.

!!! warning "Условия получения события из группового чата"  
    Чтобы получать события из группового чата или канала, назначьте бота администратором.

## Поля объекта {#Update}  

| Поле                      | Тип         | Описание                                                                                   |
|---------------------------|-------------|--------------------------------------------------------------------------------------------|
| <nobr>`update_type`<nobr> | **string**  | Тип обновления. Возможные значения:                                                        |
|||                                            `bot_added`          - Бот добавлен в чат или группу                                    |
|||                                            `bot_removed`        - Бот удален из чата или группы                                    |
|||                                            `bot_started`        - Пользователь нажал кнопку «Старт» или возобновил работу с ботом  |
|||                                            `bot_stopped`        - Пользователь заблокировал бота или остановил его работу          |
|||                                            `chat_title_changed` - Изменено название группового чата                                |
|||                                            `dialog_muted`       - Уведомления в диалоге отключены пользователем                    |
|||                                            `dialog_unmuted`     - Уведомления в диалоге снова включены                             |
|||                                            `dialog_cleared`     - История сообщений в диалоге очищена, но сам чат не удален        |
|||                                            `dialog_removed`     - Диалог полностью удален из списка чатов                          |
|||                                            `message_created`    - Получено новое входящее сообщение                                |
|||                                            `message_callback`   - Обработка нажатия на инлайн-кнопку                               |
|||                                            `message_edited`     - Ранее отправленное сообщение было изменено                       |
|||                                            `message_removed`    - Сообщение было удалено из истории переписки                      |
|||                                            `user_added`         - В чат добавлен новый участник                                    |
|||                                            `user_removed`       - Участник покинул чат или был из него удален                      |
| `timestamp`              | **integer**                            | Unix-время, когда произошло событие                              |
| `message`                | [**object**](./Message.md#Message)     | Сообщение                                                        |
| `user_locale`            | **string**   | Текущий язык пользователя в формате IETF BCP 47. Доступно только в диалогах                |

## Пример объекта {#Update-example}  

```json  
    {
        "update_type": "message_created",
        "timestamp": 0,
        "message": {
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
        },
        "user_locale": "string"
    }
```
