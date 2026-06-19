[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/NDlq6Zcp)
# File Database

![shell](https://tdwi.org/articles/2019/03/11/-/media/TDWI/TDWI/BITW/database1.jpg)

## Overview

In this assignment you will write a single shell script to implement a tool to lookup, insert, and remove records in a table.

The database table will be stored as one record per line in a file.  Multiple tables are supported, one table per file.  See the [small](./small.csv) for the required format.  It is a CSV file.

While this task can be implemented more easily with a single record per file, the point of this assignment is to understand how to use the shell to effectively manipulate a single file.  So your database must use a single file to store the data (one file per table).  Secondary/temporary files can and should be used as needed.

## Requirements

Write a single POSIX compliant shell script that satisfies the following requirements:

- `dbfile create [table] [field1,field2,...]` create a new database table. If the table already exists then issue an error and do not create a new table.
- `dbfile list` must list the names of all the tables in the database in alphabetical order.

- `dbfile add [table] [value1,value2,...]` will add a single record to the table.
- `dbfile lookup [table] [query]` will show records that match the query.  We want a somewhat performant database so this must only read the entire contents of the file/table once and only store one line/record at a time in memory.
- `dbfile delete [table] [query]` must remove all records that match the query
- `dbfile update [table] [query] [value1,value2,...]` must replace the record that matches the query with the new record.

- The database must be stored in a directory given by the environment variable `DBFILE_DATABASE` if provided, otherwise it is stored in the `db` directory in the current working directory.
- The query language is simple (not SQL, thank goodness).  The format is "Field=Value".  A record matches if it has a provided field's value matches the provided value.  An example query is "Country=United States".
- Spaces are allowed in field names and in values.
- Do not hard code the field names in your shell script.  This is part of the database table's schema.  The tool must be able handle any number of columns (field names).
- You must use commas to delimit values (so that we have a CSV file for each table). To simplify parsing, commas are not allowed in field names and values.
- The script must define and use at lease one shell function.
- All the code must be contained in the shell script at `src/dbfile`.

## Hints

For this assignment we will use the ShellCheck tool for linting.  It can be installed by following these [instructions](https://github.com/koalaman/shellcheck#installing).  Shellcheck can also be installed as a VSCode extension, [here](https://marketplace.visualstudio.com/items?itemName=timonwong.shellcheck).

Before you start the assignment please consider reading the manual pages for the following commands:

- `grep`
- `cut`
- `sed`

You might find some useful hints in Chapter 13 of the "Shell Programming" book.

Since the tables are stored as CSV files the VS Code extension [Rainbow CSV](https://marketplace.visualstudio.com/items?itemName=mechatroner.rainbow-csv) might be useful.

## Process

Then create the file `src/dbfile`, commit, ensure `make test -B` passes and then push your code to GitHub.
Make sure your GitHub action passes as expected.  Please do not modify any files in this project other than file in `src`.
