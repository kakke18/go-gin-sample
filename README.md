# go-gin-sample

gin & jsonRPC


## 動作確認
### Request
```
curl '${endpoint}/jsonrpc' -H 'content-type: application/json' --data-binary '{"jsonrpc": "2.0", "method": "App.GetUsers", "params": {}, "id": "${UUID}"}'
```

### Response
```json
{
  "jsonrpc":"2.0",
  "result":{
    "users":[
      {
        "key":"1",
        "name":"user1",
        "email_address":"user1@example.com",
        "created_at":"2020-10-01T00:00:00Z"
      },
      {
        "key":"2",
        "name":"user2",
        "email_address":"user2@example.com",
        "created_at":"2020-10-01T00:00:00Z"
      },
      {
        "key":"3",
        "name":"user3",
        "email_address":"user3@example.com",
        "created_at":"2020-10-01T00:00:00Z"
      }
    ]
  },
  "id": "${UUID}"
}
```

## GAEへのデプロイ方法
```
gcloud app deploy
```
