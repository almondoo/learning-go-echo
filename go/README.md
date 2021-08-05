# Go 基本ルールまとめ  

## Response  

| String | 内容 | 返り値 |
| --- | --- | --- |
| data | クライアントに送りたいデータ | データ |
| status | 成功か失敗か | ok or ng |
| token | もしAccessTokenが無効でRefreshTokenで認証した時に返す値 | token or null |



```go
"data": {
  interface
},
"status": code < 300 ok else ng,
"token": {
  "AccessToken": "ey...",
  "RefreshToken": "ey..."
} or null
```

## 依存関係
router  
↓  
Handler  
↓　　
usecase  
↓  
service  
↓  
repository  


## ディレクトリ構成  


```
go
├── application
│   └── usecase ユースケース
│
├── db
│   └── mysql.go DB接続系
│
├── domain
│   ├── entity テーブル構成
│   ├── repositoty リポジトリ
│   └── service ビジネスロジック
│
├── infrastructure
│   ├── auth 認証系
│   ├── cookie クッキー操作
│   ├── persistence リポジトリビジネスロジック
│   ├── security セキュリティー
│   └── storedb KVS系
│
├── interface
│   ├── context 認証系
│   ├── fileupload ファイル操作
│   ├── handler コントローラー 
│   ├── middleware ミドルウェア
│   ├── response カスタムレスポンス
│   └── validation バリデーション
│
└── main.go 
```
