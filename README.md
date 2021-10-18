# aion-instant-kanban-delivery
aion-instant-kanban-deliveryは、ローカルからREST API経由でkanbanにデータを投入するための開発者用のマイクロサービスです。  
aion-instant-kanban-deliveryでは、任意のマイクロサービスとして任意のkanbanデータを送信できるので、開発中などの理由で依存関係にあるマイクロサービスがデプロイできない場合でも、
kanbanの動作確認を行うことができます。

# 使い方
### デプロイ
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

`POST /kanban/write`
```
{
  "connection_key": "コネクションキー(string　任意)",
  "metadata"(kanbanに投入するmetadata(配列) 必須): [
    {
      "property_name": "メタデータのプロパティ名",
      "body": "メタデータのボディ"
    }
  ]
  device_name: "デバイス名（string 任意）"
}
```

