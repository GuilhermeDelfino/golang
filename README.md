# Golang Anotations - Delfino

#### Starting a Golang project

`go mod init <source/project-name>`

#### Folder Structure (example)

```
/my-golang-project
  ├── cmd
  │   └── myapp
  │       └── main.go
  ├── pkg
  │   ├── mypackage1
  │   │   └── ...
  │   └── mypackage2
  │       └── ...
  ├── internal
  │   └── myinternalpackage
  │       └── ...
  ├── api
  │   └── myapi
  │       └── ...
  ├── web
  │   ├── static
  │   │   └── ...
  │   ├── templates
  │   │   └── ...
  │   └── main.go
  ├── scripts
  │   └── ...
  ├── test
  │   └── ...
  ├── docs
  │   └── ...
  ├── LICENSE
  ├── README.md
  └── go.mod
```

###### Explanation:

- /cmd: This directory is for your application-specific code. Each application can have its own subdirectory (e.g., myapp) with a main.go file.
- /pkg: This directory is for libraries and packages that are meant to be reused across different applications.
- /internal: Similar to /pkg, but the code here is only accessible to the main application and its subpackages.
- /api: If your project involves an API, you can structure it here.
- /web: This is for web application code. You can have subdirectories for static files, templates, etc.
- /scripts: Any utility or script files.
- /test: Test files for your project.
- /docs: Documentation files.

#### Starting

We need a function with main name to be entrypoint of our application

```
  package main

  func main(){
    // Code to start here
    ...
  }
```
