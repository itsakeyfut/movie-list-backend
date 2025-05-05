## Gin ではどうやって、JSON を解釈しているのか？

`c.JSON()` は、`http.ResponseWriter` に対して `Content-Type` ヘッダーを `application/json` に設定し、

`json.Marshal()` を使って JSON に変換し、その結果をレスポンスとして返却します。

Gin 内部で行われていることを簡単に示すと以下のようになります：

```go
// Ginの内部実装
func (c *Context) JSON(code int, obj interface{}) {
    c.SetHeader("Content-Type", "application/json")
    c.Status(code)
    if obj != nil {
        jsonBytes, _ := json.Marshal(obj) // ここでMarshalされる
        c.Writer.Write(jsonBytes)
    }
}
```

このように、`json.Marshal()` を使って自動でデータを JSON 形式に変換し、クライアントに送信するため、Gin を使う際に自分で `Marshal` を書く必要はありません。

リクエストでも自動でデシリアライズします。
