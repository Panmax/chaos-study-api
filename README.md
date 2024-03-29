## 《混沌复习》

此项目为《混沌复习》的后端

> 不知道你是不是和我一样因为焦虑胡乱买了很多的课程、也胡乱学了很多课程，但过了一段时间想再回顾一遍之前学过的东西，却不知如何下手，混沌复习来帮你解决这个难题，不用再纠结每天学点什么，跟着混沌走就对了。

### 只需 3 步即可开启你的混沌复习之旅：
1. 录入你所有希望复习的课程（当然也可以是书籍）
2. 设置内每日的复习计划
3. 每日登录混沌复习，查看今日要复习的内容

## API 文档

> API 测试地址：`http://chaos.jpanj.com`

### 统一前缀 /api


### POST 创建用户

`/user`

request

```json
{
	"username": "username",
	"password": "password"
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

### 登录

`/auth/login`

request

```json
{
	"username": "username",
	"password": "password"
}
```

response

```json
{
    "code": 200,
    "expire": "2019-09-17T06:12:12+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njg2NzE5MzIsIm9yaWdfaWF0IjoxNTY4NjQzMTMyLCJ1c2VybmFtZSI6ImppYXBhbiJ9.ooM41jdWEkBb5y41KMh49g3FJ6PcVfBpVfMsVeYXrvY"
}
```

### POST 修改密码

`/user/password`

request

```json
{
	"password": "654321"
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

### POST 创建课程

`/course`

request

```json
{
	"name": "The Linux Command Line", // 课程名
	"total_chapter": 37, // 共计多少章节
	"url": "http://billie66.github.io/TLCL/book/index.html", // 在线地址（可选）
	"pick": 1  // 每日要复习的章节数
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

`/courses?page=0&size=20`

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
	"count" :2, // 每日复习课程数
	"not_repeat": true // 不允许重复
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