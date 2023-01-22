
render_ER_diagram_to_svg:
	d2 ./reference/ER_draft.d2

watch_ER_diagram:
	d2 ./reference/ER_draft.d2 --watch --host 127.0.0.1 --port 54321
