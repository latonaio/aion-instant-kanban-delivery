# aion-instant-kanban-delivery
localからREST API経由でkanbanにデータを投入できる、開発者向けのマイクロサービスです。
任意のマイクロサービスとして任意のkanbanデータを送信できるので、開発中などの理由で依存関係にあるマイクロサービスがデプロイできない場合でも、
kanbanの動作確認を行うことができます。

### 動作環境
このマイクロサービスを利用するために Kubernetes が動作する環境が必要です。
- OS: Linux
  
- CPU: Intel64/AMD64/ARM64

- Kubernetes


### デプロイ方法
`kubectl apply -f deployment.yml`

### 環境変数
IS_DEBUG: デバッグモードを使用するかどうか。使用する場合、モッククライアントを使用し、kanbanサーバーを使用しません。

### kanbanの操作
kanbanのセットと書き込みをAPIを通じて行うことができます。
#### セット
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

