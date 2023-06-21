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

Further tests with multiple variables

```bash
go run main.go -t "Hello {{ name }} {{ second }}" -v 'name=World,second=Blah'
go run main.go -t "Hello {{ name }} {{ second }} {{ num |int * 2 }}" -v 'name=World,second=Blah,num=99'
```

```bash
cat << EOF > /tmp/hello_world.j2
Hello {{ first_name }} {{ last_name }}, you are {{ age }}
EOF

go run main.go -t /tmp/hello_world.j2 -v 'first_name=Rhys,last_name=Campbell,age=40'
```