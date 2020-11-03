# go-gin-sample

gin & jsonRPC


## 動作確認
### Request
```
curl 'localhost:8080/jsonrpc' -H 'content-type: application/json' --data-binary '{"jsonrpc": "2.0", "method": "App.GetMembers", "params": {}, "id": "243a718a-2ebb-4e32-8cc8-210c39e8a14b"}'
```

### Response
```json
{
  "jsonrpc":"2.0",
  "result":{
    "members":[
      {
        "key":"1",
        "name":"member1",
        "email_address":"member1@example.com",
        "created_at":"2020-10-01T00:00:00Z"
      },
      {
        "key":"2",
        "name":"member2",
        "email_address":"member2@example.com",
        "created_at":"2020-10-01T00:00:00Z"
      },
      {
        "key":"3",
        "name":"member3",
        "email_address":"member3@example.com",
        "created_at":"2020-10-01T00:00:00Z"
      }
    ]
  },
  "id":"243a718a-2ebb-4e32-8cc8-210c39e8a14b"
}
```