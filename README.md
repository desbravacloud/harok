Rocket
======

Rocket is a SRE tool created to automate and manage your infrastructure

# Install

Binary downloads of the Rocket can be found on the Releases page.

Unpack the rocket binary and add it to your PATH and you are good to go!

If you want to build the binary:

- `go build`

Authentication to Database
======

Create a file in $HOME/.harok/config.json

´
{
  "credentials": {
    "db_user": "postgres",
    "db_password": "postgres",
    "db_endpoint": "localhost",
    "db_database": "example",
    "db_port": "5432"
  }
}
´