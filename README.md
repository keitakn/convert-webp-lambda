# convert-webp-lambda

S3にアップロードされた画像ファイルをWebpフォーマットに変換する為のLambda関数です。

## Getting Started

### AWSクレデンシャルの設定

[名前付きプロファイル](https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-configure-profiles.html) を利用しています。

このプロジェクトで利用しているプロファイル名は `nekochans-dev` です。

### 環境変数の設定

[direnv](https://github.com/direnv/direnv) 等を利用して環境変数を設定します。

```
export DEPLOY_STAGE=デプロイステージを設定、デフォルトは dev
export TRIGGER_BUCKET_NAME=Lambda関数実行のトリガーとなるS3バケット名を指定
export DESTINATION_BUCKET_NAME=変換済のWebp画像がアップロードされるS3バケット名を指定
```

### デプロイ

以下の手順です。

1. `docker-compose up --build -d` でコンテナを起動します
1. `docker-compose exec node bash` でコンテナに入る
1. `yarn install` で依存パッケージをインストール
1. `yarn run deploy` を実行する

必ずDockerコンテナ内でこれらの作業を行って下さい。

Webpフォーマットへの変換は [sharp](https://github.com/lovell/sharp) を利用しているので、予めビルドされたネイティブモジュール（バイナリ）がLambdaの実行環境と異なると正常に動作しません。

詳しくは [AWS Lambda(Node.js)にsharp(Native Module)をデプロイする方法](https://dev.classmethod.jp/articles/how-to-deploy-with-native-module/) をご覧下さい。

ちなみに `main` ブランチにマージされたタイミングでデプロイが自動実行されるようになっています。

## Lambda関数の仕様

`TRIGGER_BUCKET_NAME` の `uploads/` ディレクトリにアップロードされたファイルをWebpに変換し `DESTINATION_BUCKET_NAME` の `encoded/` ディレクトリに移動します。

対応している画像フォーマットは `.png` だけですが、簡単な拡張で他の画像フォーマットにも対応可能です。

## その他

ディレクトリ構成は以下の公式テンプレートを利用しています。

`sls create -t aws-nodejs-typescript -p [任意のディレクトリ名]` でプロジェクトの初期構成を作成して、そこから微調整しています。

https://github.com/serverless/serverless/tree/master/lib/plugins/create/templates/aws-nodejs-typescript
