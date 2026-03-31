# GetVideoInfo

[dev.max.ru/docs-api/methods/GET/videos/-videoToken-](https://dev.max.ru/docs-api/methods/GET/videos/-videoToken-)

Метод предназначен для получения статуса обработки и метаданных загруженного видео.

## Запрос {#request}

### Поля запроса {#request-parameters}

| Поле         | Тип        | Обязательный | Описание             |
|--------------|------------|--------------|----------------------|
| `VideoToken` | **string** | Да           | Токен видео-вложения |

### Пример запроса {#request-example}

```go
response, err := bot.Messages.GetVideoInfo(
    context.Background(), 
    models.GetVideoInfoReq{
        VideoToken: "f9LHo9qcjkTKVMC1O8b_62OLxJsPk0_8edD0clywyQ2lv0GN5P0QBI5aikR",
    }
)
```

## Ответ {#response}

### Поля ответа {#response-parameters}

| Поле        | Тип         | Описание                                                          |
|-------------|-------------|-------------------------------------------------------------------|
| `token`     | **string**  | Токен видео-вложения                                              |
| `urls`      | [**object**](#VideoUrls) | URL-ы для скачивания или воспроизведения видео       |
| `thumbnail` | [**object**](../objects/Message.md#photoattachmentpayload)  | Миниатюра видео   |
| `width`     | **integer** | Ширина видео                                                      |
| `height`    | **integer** | Высота видео                                                      |
| `duration`  | **integer** | Длина видео в секундах                                            |

### Поля объекта `VideoUrls` {#VideoUrls}

| Поле        | Тип         | Описание                                       |
|-------------|-------------|------------------------------------------------|
| `mp4_1080`  | **string**  | URL видео в разрешении 1080p, если доступно    |
| `mp4_720`   | **string**  | URL видео в разрешении 720p, если доступно     |
| `mp4_480`   | **string**  | URL видео в разрешении 480p, если доступно     |
| `mp4_360`   | **string**  | URL видео в разрешении 360p, если доступно     |
| `mp4_240`   | **string**  | URL видео в разрешении 240p, если доступно     |
| `mp4_144`   | **string**  | URL видео в разрешении 144p, если доступно     |
| `hls`       | **string**  | URL трансляции, если доступна                  |


### Пример тела ответа {#response-example-body}

```json
{
    "token": "f9LHo9qcjkTKVMC1O8b_62OLxJsPk0_8edD0clywyQ2lv0GN5P0QBI5aikR",
    "width": 556,
    "height": 1280,
    "duration": 3741,
    "urls": {
        "mp4_480": "http://vd667.okcdn.ru/?expires=177392338663"
    },
    "thumbnail": {
        "url": "https://pimg.mycdn.me/getImage?disableStub1gj0w"
    }
}
```