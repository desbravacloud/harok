Harok
======

Harok is a SRE tool created to automate and manage your infrastructure

# Install

Binary downloads of the harok can be found on the Releases page.

Unpack the harok binary and add it to your PATH and you are good to go!

If you want to build the binary:

- `go build`

# Running Migrations

Add your credentials in database.yml (you can create this file from database.yml.example at the root of the project)

Download on your machine Soda cli
```shell
go get github.com/gobuffalo/pop/...
```

Check if the command has been installed on your machine
```shell
ls -l ~/go/bin/soda
```

For more information, check on [official page](https://gobuffalo.io/documentation/database/pop/)

After the steps above, you can run the migrations on your machine with:

```shell
soda migrate up
```

# Authentication to Database

Create a file in $HOME/.harok/config.json

```json
{
  "credentials": {
    "db_user": "postgres",
    "db_password": "postgres",
    "db_endpoint": "localhost",
    "db_database": "example",
    "db_port": "5432",
    "github_token" : "12345678",
    "github_org" : "spotlitebr",
    "jenkins_address": "https://jenkins.teste.io",
    "jenkins_user": "jenkinsUser",
    "jenkins_token": "1234567890"
  }
}
```