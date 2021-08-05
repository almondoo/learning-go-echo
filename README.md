# learning-go-echo

## 技術スタック

### フロントエンド

[Next.js](https://github.com/tsubasa111/learning-next-js)
React/Next  
SWR

## バックエンド

Golang/Echo  
gorm  
golang-migrate/migrate  
jwt

### DB

MySQL 5.7

## 本番環境

### ホスティング

Vercel

### サーバー

GAE(スタンダード)  
F1

### NoSQL

Google Cloud Database

## ローカル開発環境構築手順

ファイル生成  
1\. make build

docker 起動  
2. make up

docker workspace に入る  
3. make bash

vendor を生成する  
4. go mod vendor

## DB 関連手順

Auto Migration
