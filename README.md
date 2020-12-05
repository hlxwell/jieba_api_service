# 结巴中文分词 API 服务

仅仅只是把 Go Jieba (够结巴) 给包装了成了一个 web api service，方便其他语言的程序调用。

## How to run

```
go build -v .
```

## How to use

```
GET /api/v1/word_divider?sentence=中华人民共和国万岁。
```

return

```
{
  "result": [
    "中华人民共和国",
    "万岁",
    "！"
  ]
}
```
