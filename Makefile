
render_er_diagram_to_svg: ## ER図をSVG画像ファイルにレンダリングする
	d2 ./reference/ER_draft.d2

watch_er_diagram: ## ER図をローカルサーバで起動・閲覧する
	d2 ./reference/ER_draft.d2 --watch --host 127.0.0.1 --port 54321

render_apis_specs_to_md: ## API仕様書をMarkdown形式にレンダリングする
	npx widdershins --omitHeader --code true ./reference/fcoin-balances.yaml ./reference/fcoin-balances.md

watch_apis_specs: ## API仕様書をローカルサーバで起動・閲覧する
	npx @redocly/cli preview-docs  reference/fcoin-balances.yaml --host "127.0.0.1" --port 65535

build: ## サービス（app）をビルドする
	docker compose build --no-cache

run: ## buildされたサービス（appとdb）を起動する
	docker compose up -d

run_app: ## buildされたappを起動する
	docker compose up app -d

up: build run ## サービス（appとdb）をビルド・起動する

up_app: build run_app ## appをビルド・起動する

down: ## サービス（appとdb）を停止・削除する
	docker compose down

down_app: ## appを停止・削除する
	docker compose rm app --stop --force

reload_app: down_app up_app

rm: down rm_volume ## サービス（appとdb）を停止・削除（volumeも）する

status: ## サービス（appとdb）のステータスを確認する
	docker compose ps --all

app_logs: ## appのログを閲覧する
	docker compose logs app --no-log-prefix

db_logs: ## appのログを閲覧する
	docker compose logs db --no-log-prefix

db_in: ## 起動しているdbに接続する
	mysql -h 127.0.0.1 -u taro --password=pass fcoin-balances-db

rm_volume: ## ローカルのvolumeを削除する
	docker volume rm fcoin-balances-manager_db_data

test: ## テストを実行する
  ## go: -race requires cgo; enable cgo by setting CGO_ENABLED=1
	go test -v -race -shuffle=on -covermode=atomic ./...

help: ## makeコマンドの一覧を表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
