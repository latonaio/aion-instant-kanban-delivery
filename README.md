# aion-instant-kanban-delivery
## 概要
aion-instant-kanban-deliveryは、localからREST API経由でkanbanにデータを投入できる、開発者向けのマイクロサービスです。   
任意のマイクロサービスとして任意のkanbanデータを送信できるので、開発中などの理由で依存関係にあるマイクロサービスがデプロイできない場合でも、kanbanの動作確認を行うことができます。

## 動作環境
aion-instant-kanban-deliveryは、aion-coreのプラットフォーム上での動作を前提としています。   
使用する際は、事前に下記の通りAIONの動作環境を用意してください。   
* ARM CPU搭載のデバイス(NVIDIA Jetson シリーズ等)   
* OS: Linux Ubuntu OS   
* CPU: ARM64   
* Kubernetes   
* [AION](https://github.com/latonaio/aion-core)のリソース  

## セットアップ
以下のコマンドを実行して、docker imageを作成してください。
```
$ cd /path/to/aion-instant-kanban-delivery
$ bash docker-build.sh
```

## 起動方法
### 環境変数
* IS_DEBUG: デバッグモードを使用するかどうか。使用する場合、モッククライアントを使用し、kanbanサーバーを使用しません。   

### デプロイ
以下のコマンドを実行して、podを立ち上げてください。
```
$ cd /path/to/aion-instant-kanban-delivery
$ kubectl apply -f deployment.yml
```

## kanbanの操作方法
kanbanのセットと書き込みをAPIを通じて行うことができます。

### セット
`POST /kanban/set`
```
{
  "mimicked_ms_name": "なりすましたいマイクロサービスの名前(string 必須)",
  "process_num": "プロセスナンバー(int 必須)"
}
```
#### 書き込み
`POST /kanban/write`
```
{
  "connection_key": "コネクションキー(string 任意)",
  "metadata"(kanbanに投入するmetadata(配列) 必須): [
    {
      "property_name": "メタデータのプロパティ名",
      "body": "メタデータのボディ"
    }
  ]
  device_name: "デバイス名（string 任意）"
}
```

