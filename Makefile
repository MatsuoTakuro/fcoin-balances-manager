
render_ER_diagram_to_svg:
	d2 ./reference/ER_draft.d2

watch_ER_diagram_on_local_server:
	d2 ./reference/ER_draft.d2 --watch --host 127.0.0.1 --port 54321

render_APIs_specs_to_markdown:
	npx widdershins --omitHeader --code true ./reference/fcoin-balances.yaml ./reference/fcoin-balances.md

watch_APIs_specs_on_local_server:
	npx @redocly/cli preview-docs  reference/fcoin-balances.yaml --host "127.0.0.1" --port 65535
