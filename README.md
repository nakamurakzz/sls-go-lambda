# Serverless FrameworkでGoのLambda Functionをデプロイする
## Install
serverless frameworkのインストール
```bash
npm i -g serverless
```

Goのモジュールインストール
```bash
go mod tidy
```

## Build
```bash
make build
```

## Deploy
```bash
make deploy
```

- aws-cliで設定したdefaultのprofileを使い、Lambda Functionをデプロイします
