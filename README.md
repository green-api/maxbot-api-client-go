# MaxBot API Client (Go)

`maxbot-api-client-go` — это библиотека для интеграции с API MAX Bot. Этот проект предоставляет структурированный интерфейс для взаимодействия с конфигурациями бота, управления сообщениями, отправки медиафайлов и подписки на события через long-polling.

Для использования библиотеки потребуется получить токен бота в консоли разработчика MAX API.

## API

Документацию по REST API MAX можно найти по ссылке [https://dev.max.ru/docs-api]. Библиотека является оберткой для REST API, поэтому документация по указанной выше ссылке также применима к используемым здесь моделям и параметрам запроса.

## Поддержка

[![Support](https://img.shields.io/badge/support@green--api.com-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:support@green-api.com)
[![Support](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/greenapi_support_ru_bot)
[![Support](https://img.shields.io/badge/WhatsApp-25D366?style=for-the-badge&logo=whatsapp&logoColor=white)](https://wa.me/77780739095)

## Руководства и новости

[![Guides](https://img.shields.io/badge/YouTube-%23FF0000.svg?style=for-the-badge&logo=YouTube&logoColor=white)](https://www.youtube.com/@green-api)
[![News](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/green_api)
[![News](https://img.shields.io/badge/WhatsApp-25D366?style=for-the-badge&logo=whatsapp&logoColor=white)](https://whatsapp.com/channel/0029VaHUM5TBA1f7cG29nO1C)


## Установка

**Убедитесь, что у вас установлена версия Go не ниже 1.20**

```shell
go version
```

**Создайте Go модуль, если он не создан:**

```
go mod init ModuleName
```

**Установите библиотеку:**

```bash
go get github.com/green-api/maxbot-api-client-go
```

**Импорт:**

```go
import(
    "github.com/green-api/maxbot-api-client-go"
)
```

## Использование и примеры

**Как инициализировать клиент:**

```go
bot, err := api.New(client.Config{
    BaseURL: "https://platform-api.max.ru",
    Token:   "YOUR_BOT_TOKEN",
})
if err != nil {
    log.Fatal().Err(err).Msg("failed to init MAX API")
}
```

**Как получить информацию о боте:**

Ссылка на пример: [GetBot/main.go](./cmd/examples/GetBot/main.go)

```go
response, err := bot.Bots.GetBot(context.Background())
```

**Как отправить сообщение:**

Ссылка на пример: [SendMessage/main.go](./cmd/examples/SendMessage/main.go)

```go
response, err := bot.Messages.SendMessage(context.Background(), models.SendMessageReq{
    UserID: 1234567890,
    Text: "Hello world!",
})
```

**Как легко отправить файл (по ссылке или локальный):**

Ссылка на пример: [SendFile/main.go](./cmd/examples/SendFile/main.go)

```go
response, err := bot.Helpers.SendFile(context.Background(), models.SendFileReq{
    ChatID:     exampleChatID,
    FileSource: "cmd/examples/assets/file.txt",
})
```

```go
response, err := bot.Helpers.SendFile(context.Background(), models.SendFileReq{
    ChatID:     exampleChatID,
    FileSource: "https://storage.yandexcloud.net/sw-prod-03-test/ChatBot/corgi.jpg",
})
```

**Как вручную загрузить файл (для кастомных вложений):**

Ссылка на пример: [UploadFile/main.go](./cmd/examples/UploadFile/main.go)

```go
response, err := bot.Uploads.UploadFile(context.Background(), models.UploadFileReq{
    Type:     models.UploadImage,
    FilePath: "/path/to/your/image.png",
})
```

**Как получить входящее уведомление:**

Ссылка на пример: [GetUpdates/main.go](./cmd/examples/GetUpdates/main.go)

```go
response, err := bot.Subscriptions.GetUpdates(context.Background())
```

---

## Список примеров

| Описание                             | Ссылка на пример                                             |
|--------------------------------------|--------------------------------------------------------------|
| Как отправить сообщение              | [SendMessage/main.go](./cmd/examples/SendMessage/main.go)    |
| Как получить информацию о боте       | [GetBot/main.go](./cmd/examples/GetBot/main.go)              |
| Как отправить файл                   | [SendFile/main.go](./cmd/examples/SendFile/main.go)          |
| Как загрузить файл                   | [UploadFile/main.go](./cmd/examples/UploadFile/main.go)      |
| Как получить входящее уведомление    | [GetUpdates/main.go](./cmd/examples/GetUpdates/main.go)      |


## Список всех методов библиотеки

| Метод API       | Описание                | Ссылка на документацию MAX                     | Ссылка на документацию библиотеки            |
|-----------------|-------------------------|------------------------------------------------|----------------------------------------------|
| `Bots.GetBot` | Получает информацию о боте | [GetBot](https://dev.max.ru/docs-api/methods/GET/me) | [GetBot](./docs/bots/GetBot.md) |
| `Bots.PatchBot` | Изменяет информацию о боте | | [PatchBot](./docs/bots/PatchBot.md) |
| `Chats.GetChats` | Возвращает список групповых чатов, в которых участвовал бот | [GetChats](https://dev.max.ru/docs-api/methods/GET/chats) | [GetChats](./docs/chats/GetChats.md) |
| `Chats.GetChat` | Возвращает информацию о групповом чате по его ID | [GetChat](https://dev.max.ru/docs-api/methods/GET/chats/-chatId-) | [GetChat](./docs/chats/GetChat.md) |
| `Chats.EditChat` | Позволяет редактировать информацию о групповом чате | [EditChat](https://dev.max.ru/docs-api/methods/PATCH/chats/-chatId-) | [EditChat](./docs/chats/EditChat.md) |
| `Chats.DeleteChat` | Удаляет групповой чат для всех участников | [DeleteChat](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-) | [DeleteChat](./docs/chats/DeleteChat.md) |
| `Chats.SendAction` | Позволяет отправлять следующие действия бота в групповой чат | [SendAction](https://dev.max.ru/docs-api/methods/POST/chats/-chatId-/actions) | [SendAction](./docs/chats/SendAction.md) |
| `Chats.GetPinnedMessage` | Возвращает закрепленное сообщение в чате | [GetPinnedMessage](https://dev.max.ru/docs-api/methods/GET/chats/-chatId-/pin) | [GetPinnedMessage](./docs/chats/GetPinnedMessage.md) |
| `Chats.PinMessage` | Закрепляет сообщение в групповом чате | [PinMessage](https://dev.max.ru/docs-api/methods/PUT/chats/-chatId-/pin) | [PinMessage](./docs/chats/PinMessage.md) |
| `Chats.UnpinMessage` | Удаляет закрепленное сообщение в групповом чате | [UnpinMessage](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/pin) | [UnpinMessage](./docs/chats/UnpinMessage.md) |
| `Chats.GetChatMembership` | Возвращает членство бота в групповом чате | [GetChatMembership](https://dev.max.ru/docs-api/methods/GET/chats/-chatId-/members/me) | [GetChatMembership](./docs/chats/GetChatMembership.md) |
| `Chats.LeaveChat` | Удаляет бота из группового чата | [LeaveChat](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/members/me) | [LeaveChat](./docs/chats/LeaveChat.md) |
| `Chats.GetChatAdmins` | Возвращает список всех администраторов группового чата | [GetChatAdmins](https://dev.max.ru/docs-api/methods/GET/chats/-chatId-/members/admins) | [GetChatAdmins](./docs/chats/GetChatAdmins.md) |
| `Chats.SetChatAdmins` | Назначает участника группы администратором | [SetChatAdmins](https://dev.max.ru/docs-api/methods/POST/chats/-chatId-/members/admins) | [SetChatAdmins](./docs/chats/SetChatAdmins.md) |
| `Chats.DeleteAdmin` | Отменяет права администратора пользователя в групповом чате | [DeleteAdmin](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/members/admins/-userId-) | [DeleteAdmin](./docs/chats/DeleteAdmin.md) |
| `Chats.GetChatMembers` | Возвращает список участников группового чата | [GetChatMembers](https://dev.max.ru/docs-api/methods/GET/chats/-chatId-/members) | [GetChatMembers](./docs/chats/GetChatMembers.md) |
| `Chats.AddMembers` | Добавляет участников в групповой чат | [AddMembers](https://dev.max.ru/docs-api/methods/POST/chats/-chatId-/members) | [AddMembers](./docs/chats/AddMembers.md) |
| `Chats.DeleteMember` | Удаляет участника из группового чата | [DeleteMember](https://dev.max.ru/docs-api/methods/DELETE/chats/-chatId-/members) | [DeleteMember](./docs/chats/DeleteMember.md) |
| `Subscriptions.GetSubscriptions` | Возвращает список подписок на уведомления веб-хуков | [GetSubscriptions](https://dev.max.ru/docs-api/methods/GET/subscriptions) | [GetSubscriptions](./docs/subscriptions/GetSubscriptions.md) |
| `Subscriptions.Subscribe` | Настраивает доставку событий бота через веб-хук | [Subscribe](https://dev.max.ru/docs-api/methods/POST/subscriptions) | [Subscribe](./docs/subscriptions/Subscribe.md) |
| `Subscriptions.Unsubscribe` | Отменяет подписку бота на получение обновлений через веб-хук | [Unsubscribe](https://dev.max.ru/docs-api/methods/DELETE/subscriptions) | [Unsubscribe](./docs/subscriptions/Unsubscribe.md)|
| `Subscriptions.GetUpdates` | Получает входящие обновления | [GetUpdates](https://dev.max.ru/docs-api/methods/GET/updates) | [GetUpdates](./docs/subscriptions/GetUpdates.md) |
| `Upload.UploadFile` | Загружает файл на серверы MAX для последующей передачи | [UploadFile](https://dev.max.ru/docs-api/methods/POST/uploads) | [UploadFile](./docs/upload/UploadFile.md) |
| `Helpers.SendFile` | Упрощает отправку файлов, автоматически определяя URL или путь | | [SendFile](./docs/helpers/SendFile.md) |
| `Messages.GetMessages` | Возвращает информацию о сообщении или массив сообщений из чата | [GetMessages](https://dev.max.ru/docs-api/methods/GET/messages) | [GetMessages](./docs/messages/GetMessages.md) |
| `Messages.SendMessage` | Отправляет текстовое или медиа-сообщение указанному пользователю или в чат | [SendMessage](https://dev.max.ru/docs-api/methods/POST/messages) | [SendMessage](./docs/messages/SendMessage.md) |
| `Messages.EditMessage` | Редактирует текст или медиафайл ранее отправленного сообщения | [EditMessage](https://dev.max.ru/docs-api/methods/PUT/messages) | [EditMessage](./docs/messages/EditMessage.md) |
| `Messages.DeleteMessage` | Удаляет сообщение из чата | [DeleteMessage](https://dev.max.ru/docs-api/methods/DELETE/messages) | [DeleteMessage](./docs/messages/DeleteMessage.md) |
| `Messages.GetMessage` | Извлекает содержимое и метаданные конкретного сообщения по его ID | [GetMessage](https://dev.max.ru/docs-api/methods/GET/messages/-messageId-) | [GetMessage](./docs/messages/GetMessage.md) |
| `Messages.GetVideoInfo` | Возвращает подробную информацию о прикрепленном видео | [GetVideoInfo](https://dev.max.ru/docs-api/methods/GET/videos/-videoToken-) | [GetVideoInfo](./docs/messages/GetVideoInfo.md) |
| `Messages.AnswerCallback` | Отправляет ответ после того, как пользователь нажмет кнопку | [AnswerCallback](https://dev.max.ru/docs-api/methods/POST/answers) | [AnswerCallback](./docs/messages/AnswerCallback.md)|
