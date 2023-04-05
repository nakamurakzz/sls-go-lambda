# Serverless FrameworkでGoのLambda Functionをデプロイする
## Install
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

- aws-cliで設定したdefaultのprofileを使い、Lambda Functionをデプロイする