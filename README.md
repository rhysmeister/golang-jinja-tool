# golang-jinja-tool
A simple cmd-line tool to parse and display Jinja templates. Just a simple project to learn some of the basics of Go.

Setup project...

```bash
go mod init jinja2
go mod tidy
```

Run a very basic hello world example...

```bash
go run main.go -t "Hello {{ name }}" -v 'name=World'
```

Basic example using a file template...

```bash
cat << EOF > /tmp/hello_world.j2
Hello {{ name }}
EOF

go run main.go -t /tmp/hello_world.j2 -v 'name=World'
```