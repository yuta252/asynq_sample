## 内容

Goによる非同期処理の実装サンプル

https://github.com/hibiken/asynq

## 起動方法

1. workerを起動する
   ```
   make start-worker
   ```

2. clientからキューにタスクを追加する

   ```
   make start-client
   ```


## Queueのモニタリング

- 以下のコマンドでサーバーを実行

   ```
   make start-monitoring
   ```

- http://localhost:8080/monitoring/ にアクセス

## 参照

Asynq: https://github.com/hibiken/asynq

Asynqmon: https://github.com/hibiken/asynqmon
