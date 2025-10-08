# particard

A CLI tool for managing parti members.
particard is a command-line interface (CLI) tool designed to manage parti members.
It allows you to create, retrieve, update, and remove member records efficiently.


Then, you can run the setup command to initialize your database:

```bash
particard setup --db sqlite
```
This will create the necessary database directory and file, and configure the application.

## Commands

### `particard setup`

Initialize the particard database and set up environment variables.

The 'setup' command initializes the necessary directory structure for the particard
database and saves the database path in the configuration file.

Example:
  particard setup --db sqlite

### `particard new`

Create a new parti member.

The 'new' command creates a new parti member record in the database.
You can specify the parti, name, and role using flags. If not provided,
default values will be used for parti ("German Legion") and role ("member").

Example:
  particard new --name "John Doe" --parti "New Alliance" --role "Recruit"
  particard new -n "Jane Smith"

Flags:
  -n, --name string   name of parti member
  -p, --parti string  parti name (default "German Legion")
  -r, --role string   role of parti member (default "member")

### `particard get`

Get a parti member by ID.

The 'get' command retrieves and displays the information of a parti member
identified by their unique ID. You must provide a valid UUID as the argument.

Example:
  particard get 123e4567-e89b-12d3-a456-426614174000

Flags:
  -i, --ident uint   ident size for new lines (default 2)

### `particard update`

Update information of a parti member.

The 'update' command modifies the information of an existing parti member.
You must provide the member's unique ID as an argument, and then use flags
to specify the fields you wish to update (parti, name, or role).

Example:
  particard update 123e4567-e89b-12d3-a456-426614174000 --name "New Name" -p "New Parti"
  particard update 123e4567-e89b-12d3-a456-426614174000 -r "New Role"

Flags:
  -n, --name string   New name of the member
  -p, --parti string  New parti of the member
  -r, --role string   New role of the member

### `particard remove`

Remove a parti member by ID.

The 'remove' command deletes a parti member record from the database
identified by their unique ID. You must provide a valid UUID as the argument.

Example:
  particard remove 123e4567-e89b-12d3-a456-426614174000
