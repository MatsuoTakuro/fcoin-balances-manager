# fcoin-balances-manager

- 独自コイン（Fcoin = fan coin）の残高を管理するシステム

## 起動方法

 1. サービスを起動する

    > `make up`（`Makefile`を使用）
     - 例

    ```sh
    make up
    docker compose build --no-cache
    [+] Building 13.6s (14/15)
    => [internal] load build definition from Dockerfile 0.0s
    => => transferring dockerfile: 32B
    ...
    ...
    => => naming to docker.io/library/fcoin-balances-manager

    Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
    docker compose up -d
    [+] Running 3/3
    ⠿ Network fcoin-balances-manager_default  Created 0.0s
    ⠿ Container fcoin-balances-db             Healthy 10.9s
    ⠿ Container fcoin-balances-manager        Started
    ```

 2. ユーザーとして新規登録する（Fcoin残高を作成する）

      > `curl -XPOST localhost:8080/user -d '{"name": "{your name}"}'`

    - 例

    ```sh
    curl -XPOST localhost:8080/user -d '{"name": "taro"}'
    {"user_id":1,"name":"taro","balance":{"amount":0}}%
    ```

 3. Fcoin残高に5000を追加する

    > `curl -XPATCH localhost:8080/users/{user_id} -d '{"amount": 5000}'`

    - 例

    ```sh
    curl -XPATCH localhost:8080/users/1 -d '{"amount": 5000}'
    {"balance_trans_id":1,"user_id":1,"balance_id":1,"amount":5000,"processed_at":"2023-01-30T16:01:33.132934637Z"}%
    ```

 4. Fcoin残高から3000を消費する

    > `curl -XPATCH localhost:8080/users/{user_id} -d '{"amount": -3000}'`

    - 例

    ```sh
    curl -XPATCH localhost:8080/users/1 -d '{"amount": -3000}'
    {"balance_trans_id":2,"user_id":1,"balance_id":1,"amount":-3000,"processed_at":"2023-01-30T16:02:22.89337084Z"}%
    ```

 5. 別のユーザーに（Fcoin残高を消費して）1000を転送する

    > `curl -XPOST localhost:8080/users/{user_id}/transfer -d '{"user_id": {someones user_id}, "amount": 1000}'`

    - 例

    ```sh
    curl -XPOST localhost:8080/users/{user_id}/transfer -d '{"user_id": {someones user_id}, "amount": 1000}'
    ```

 6. 現在のFcoin残高とその取引履歴を確認する

    > `curl GET localhost:8080/users/{user_id}`

    - 例

    ```sh
    curl GET localhost:8080/users/{user_id}
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
