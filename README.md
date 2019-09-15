## 《混沌复习》

此项目为《混沌复习》的后端

> 不知道你是不是和我一样因为焦虑胡乱买了很多的课程、也胡乱学了很多课程，但过了一段时间想再回顾一遍之前学过的东西，却不知如何下手，混沌复习来帮你解决这个难题

### 只需 3 步即可开启你的混沌复习之旅：
1. 录入你所有希望复习的课程（当然也可以是书籍）
2. 设置内每日的复习计划
3. 每日登录混沌复习，查看今日要复习的内容

## API 文档

### 统一前缀 /api

### POST 创建课程

`/course`

request

```json
{
	"name": "The Linux Command Line",
	"total_chapter": 37,
	"url": "http://billie66.github.io/TLCL/book/index.html",
	"pick": 1
}
```

response

```json
{
    "code": 0,
    "message": "success",
    "data": true
}
```

### GET 获取课程列表

`/courses?limit=20&offset=0`

response

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "total": 5,
        "results": [
            {
                "id": 1,
                "name": "代码精进之路",
                "total_chapter": 45,
                "url": "",
                "pick": 1,
                "created_at": 1568474643,
                "updated_at": 1568474643
            },
            {
                "id": 2,
                "name": "ElasticSearch 核心技术与实战",
                "total_chapter": 46,
                "url": "",
                "pick": 2,
                "created_at": 1568474674,
                "updated_at": 1568474674
            },
            {
                "id": 3,
                "name": "Go语言从入门到实战",
                "total_chapter": 55,
                "url": "",
                "pick": 1,
                "created_at": 1568474693,
                "updated_at": 1568474693
            },
            {
                "id": 4,
                "name": "从 0 开始学大数据",
                "total_chapter": 45,
                "url": "",
                "pick": 1,
                "created_at": 1568474715,
                "updated_at": 1568474715
            },
            {
                "id": 5,
                "name": "The Linux Command Line",
                "total_chapter": 37,
                "url": "http://billie66.github.io/TLCL/book/index.html",
                "pick": 1,
                "created_at": 1568522753,
                "updated_at": 1568522753
            }
        ]
    }
}
```

### GET 根据ID获取课程

`/course/1`

response

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "id": 1,
        "name": "代码精进之路",
        "total_chapter": 45,
        "url": "",
        "pick": 1,
        "created_at": 1568474643,
        "updated_at": 1568474643
    }
}
```

### DELETE 删除课程

`/course/1`

response

```json
{
    "code": 0,
    "message": "success",
    "data": true
}
```

### PUT 更新课程

`/course/1`

request

```json
{
    "name": "代码精进之路",
    "total_chapter": 45,
    "url": "",
    "pick": 2
}
```

response

```json
{
    "code": 0,
    "message": "success",
    "data": true
}
```

### GET 获取复习计划

`/plan`

response

```json
{
    "code": 0,
    "message": "success",
    "data": {
        "count": 1,
        "not_repeat": false
    }
}
```

### PUT 设置复习计划

`/plan`

request

```json
{
	"count" :2,
	"not_repeat": true
}
```

response

```json
{
    "code": 0,
    "message": "success",
    "data": true
}
```

### GET 查看今日复习安排

`/courses/pick`

response

```json
{
    "code": 0,
    "message": "success",
    "data": [
        {
            "course": {
                "id": 4,
                "name": "从 0 开始学大数据",
                "total_chapter": 45,
                "url": "",
                "pick": 1,
                "created_at": 1568474715,
                "updated_at": 1568474715
            },
            "chapters": [
                20
            ]
        },
        {
            "course": {
                "id": 1,
                "name": "代码精进之路",
                "total_chapter": 45,
                "url": "",
                "pick": 2,
                "created_at": 1568474643,
                "updated_at": 1568537031
            },
            "chapters": [
                21,
                38
            ]
        }
    ]
}
```