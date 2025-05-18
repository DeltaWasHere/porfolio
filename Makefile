#watch for tailwind files changes to regenerate 'em
watch_tailwind:
	echo "watching temple"
	npx @tailwindcss/cli -i ./static/css/styles.css -o ./static/css/output.css --watch
#watch for templ files changes to regenrate em
watch_temple:
	echo "watching temple"
	templ generate --watch --cmd="go run ." --proxy="http://localhost:3000"
# run with make -j2(indicates number of jobs) watch
watch: watch_tailwind watch_temple