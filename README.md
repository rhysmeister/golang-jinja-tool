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

# Code Explanation

For my own understanding here is a line-by-line explanation of the code in main.go

* 1 - Define the fiel to part of package main.
* 3-12 - Imports required packages. All of these, expect the jinja2 package, come from the [Go standard library](https://pkg.go.dev/std).
* 14-16 - Declare a few variables.
* 18-21 - The [init function](https://www.digitalocean.com/community/tutorials/understanding-init-in-go) runs automatically when the program starts. Here we use the flag package to setup a few command-line flags.
* 23-33 - We define the processVarsIntoMap function.
  * 24 - We define 'm' to be a map with string keys and 'any' value. The any keyword is basically an alias for interface{} and means that any type can be accepted.
  * 25 - Define 'ss' as a slice of strings.
  * 27 - We split the comma-seperated variables string and put the result in 'ss'. We should end up with individual values like 'key=value'.
  * 28 - We initialise the 'm' variable using the [make function](https://zetcode.com/golang/make-fun/).
  * 29 - For setup a for loop to iterate over the values in 'ss'. We use the range function to iterate over the slice. The index is placed into '_' because we don't do anything with it (the compiler complains if we don't).
    * 30 - Split the kv pair string into a key and value, e.g. key=value -> ["key", "value"].
    * 31 - Add the kv pair to the map.
    * 33 - Return the map of kv pairs.
* 36 - 63 - The program's main function.
  * 37 - Parse the cmd line flags.
  * 39 - 40 - If the file, given in the template flag, does not exist then assume the template is a string and place it into content.
  * 41 - 47 - Else read the content from the file and use that as a jinja2 template.
  * 49 - Call the function to progress the provided variables into a map.
  * 51 - 52 - Create an object to parse the Jinja2 template and variables.
  * 53 - 55 - Check for errors.
  * 56 - Defer the close of the j2 object.
  * 58 - Parse the template and get the resulting content.
  * 59 - 61 - Check for errors.
  * 62 - Print the original template and the parsed result.
