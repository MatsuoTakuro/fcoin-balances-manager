# fcoin-balances-manager

- 独自コイン（Fcoin = fan coin）の残高を管理するシステム

## 起動方法

 1. サービスを起動する

    > **make up（`Makefile`を使用）**
    >
    > - Webアプリ（`fcoin-balances-manager`）は`localhost:8080`で起動
    >
    > - DB（MySQL、`fcoin-balances-db`）は`localhost:3306`で起動

    - 例

    ```sh
    make up
    docker compose build
    [+] Building 13.6s (14/15)
    => [internal] load build definition from Dockerfile 0.0s
    => => transferring dockerfile: 32B
    ...
    ...
    => => naming to docker.io/library/fcoin-balances-manager

    Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
    docker compose up -d
    [+] Running 3/3
    ⠿ Network fcoin-balances-manager_default  Created     0.0s
    ⠿ Container fcoin-balances-db             Healthy     10.9s
    ⠿ Container fcoin-balances-manager        Started     11.2s

    # 正常に起動しているかを確認
    docker compose ps
    NAME                     COMMAND                  SERVICE             STATUS              PORTS
    fcoin-balances-db        "docker-entrypoint.s…"   db                  running (healthy)   0.0.0.0:3306->3306/tcp, 33060/tcp
    fcoin-balances-manager   "/fcoin-balances-man…"   app                 running             0.0.0.0:8080->8080/tcp
    ```

 2. ユーザーとして新規登録する（Fcoin残高を作成する）

      > **curl -XPOST localhost:8080/user -d '{"name": "{your name}"}'**

    - 例

    ```sh
    curl -XPOST localhost:8080/user -d '{"name": "taro"}'
    {"user_id":1,"name":"taro","balance":{"amount":0}}%
    ```

 3. Fcoin残高に5000コインを追加する

    > curl -XPATCH localhost:8080/users/{user_id} -d '{"amount": 5000}'

    - 例

    ```sh
    curl -XPATCH localhost:8080/users/1 -d '{"amount": 5000}'
    {"balance_trans_id":1,"user_id":1,"balance_id":1,"amount":5000,"processed_at":"2023-01-30T16:01:33.132934637Z"}%
    ```

 4. Fcoin残高から3000コインを消費する

    > **curl -XPATCH localhost:8080/users/{user_id} -d '{"amount": -3000}'**

    - 例

    ```sh
    curl -XPATCH localhost:8080/users/1 -d '{"amount": -3000}'
    {"balance_trans_id":2,"user_id":1,"balance_id":1,"amount":-3000,"processed_at":"2023-01-30T16:02:22.89337084Z"}%
    ```

 5. 別のユーザーに（Fcoin残高を消費して）1000コインを転送する

    > **curl -XPOST localhost:8080/users/{user_id}/transfer -d '{"user_id": {someones user_id}, "amount": 1000}'**

    - 例

    ```sh
    # 事前に、転送先のユーザー（花子（"hanako"））を登録する
    curl -XPOST localhost:8080/user -d '{"name": "hanako"}'
    {"user_id":2,"name":"hanako","balance":{"amount":0}}%

    # 太郎（"taro"）から花子（"hanako"）に1000コインを転送する
    curl -XPOST localhost:8080/users/1/transfer -d '{"user_id": 2, "amount": 1000}'
    {"id":3,"user_id":1,"balance_id":1,
     "transfer": {"id":1,"from_user":1,"from_balance":1,"to_user":2,"to_balance":2,"amount":1000,"processed_at":"2023-02-04T01:48:10.526119417Z"},
     "amount":-1000,"processed_at":"2023-02-04T01:48:10.538264958Z"}%
    ```

 6. 現在のFcoin残高とその取引履歴を確認する

    > **curl localhost:8080/users/{user_id}**

    - 例

    ```sh
    curl localhost:8080/users/1
    {"balance":{"id":1,"amount":1000},
     "history":[
      {"id":1,"user_id":1,"balance_id":1,"amount":5000,"processed_at":"2023-02-04T01:41:00.228165Z"},
      {"id":2,"user_id":1,"balance_id":1,"amount":-3000,"processed_at":"2023-02-04T01:41:09.800373Z"},
      {"id":3,"user_id":1,"balance_id":1,
       "transfer":{"id":1,"from_user":1,"from_balance":1,"to_user":2,"to_balance":2,"amount":1000,"processed_at":"2023-02-04T01:48:10.526119Z"},
       "amount":-1000,"processed_at":"2023-02-04T01:48:10.538265Z"}
     ]
    }%
    ```

 7. サービスを停止する（コンテナも削除）

    > **make down（`Makefile`を使用）**

    - 例

    ```sh
    make down
    docker compose down
    [+] Running 3/3
    ⠿ Container fcoin-balances-manager        Removed         0.2s
    ⠿ Container fcoin-balances-db             Removed         1.6s
    ⠿ Network fcoin-balances-manager_default  Removed         0.0s
    ```

## 参考資料

### API仕様（on draft）

- [fcoin-balances APIs (markdown)](/reference/fcoin-balances.md)
- [fcoin-balances APIs (stoplight's workspace)](https://retail-ai.stoplight.io/docs/fcoin-balances-manager/m82708lnwhw7z-fcoin-balances) *emailでの招待が必要
- 以下のmakeコマンドでlocalサーバからも閲覧可能

  ```sh
  make watch_APIs_specs
  ```

### ER図（on draft）

![ER Diagram on draft](/reference/ER_draft.svg "ER Diagram on draft")

- 以下のmakeコマンドでlocalサーバからも閲覧可能

  ```sh
  make watch_ER_diagram
  ```
