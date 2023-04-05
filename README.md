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

## Local Run
```bash
make local
```

## Deploy
```bash
make deploy AWS_PROFILE={AWS PROFILE名}
```

- aws-cliで設定したprofile名を指定して、Lambda Functionをデプロイします

## Reference
- main関数のサンプル
https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-handler.html

- ランタイム
https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/lambda-runtimes.html

- Serverless Framework
https://www.serverless.com/framework/docs/providers/aws/guide/serverless.yml