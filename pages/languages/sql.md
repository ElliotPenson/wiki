# SQL

SQL (Structured Query Language) is a language for interacting with relational
databases in an RDBMS. This page was created with reference to Ben Forta's
[book on SQL](http://forta.com/books/0672336073/).

## Terms

- A **table** is a structured list of data of a specific, consistent type.
- A **schema** describes table layout (and other information about the
  database).
- A table consists of **columns** and **rows**.
  - Each column is a single field with an associated data type.
  - Each row is a record in a table.
- Every row in a table should have some column (or set of columns) that uniquely
  identified it. This is called a **primary key**. Primary keys must be unique
  and shouldn't be reused.

## `SELECT`

The `SELECT` statement retrieves information from one or more tables. At
minimum, you must specify what you want to select and from where you want to
select it.

```sql
-- Retrieve a single column.
SELECT first_name FROM People;

-- Retrieve multiple columns.
SELECT id, first_name, last_name FROM People;

-- Retrieve all columns.
SELECT * FROM People;
```

### Unique Values

Use the `DISTINCT` keyword to find unique rows. Note that the keyword applies to
all columns, not just the one it precedes.

```sql
SELECT DISTINCT first_name FROM People;
```

### Limit Results

Use the `LIMIT` keyword to control the number of rows returned. Note that SQL
implementations differ here.

```sql
SELECT first_name FROM People LIMIT 5;

-- Start from row 5.
SELECT first_name FROM People LIMIT 5 OFFSET 5;
```

### Sorting

Use the `ORDER BY` clause to sort data retrieved using a `SELECT`. Be sure that
the `ORDER BY` clause is the last clause in your `SELECT` statement.

```sql
SELECT first_name FROM People ORDER BY first_name;

-- Sort by multiple column names.
SELECT id FROM People ORDER BY first_name, last_name;

-- Sort in descending order (Z to A).
SELECT email FROM People ORDER BY first_name DESC;
```

### Filtering

Use the `WHERE` clause to filter data retrieved by the `SELECT` statement. The
`WHERE` clause is specified right after the table name. `AND`, `OR`, and `NOT`
may be used in where clauses. SQL (like most languages) processes `AND`
operators before `OR` operators. Use parenthesis to explicitly group related
operators.

```sql
SELECT first_name, last_name FROM People WHERE age = 21;

SELECT first_name FROM People WHERE last_name = 'Penson' AND age <= 30;

SELECT first_name FROM People WHERE last_name = 'Penson' OR last_name = 'Bland';

SELECT first_name FROM People WHERE NOT last_name = 'Penson';
```

The `IN` operator is used to specify a range of conditions, any of which can be
matched.

```sql
SELECT last_name FROM People WHERE zip IN ('20005', '20036') ORDER BY last_name;
```

Various operators are allowed. Note that single quotes should be used for string
literals.

| Operator              | Description                  |
| --------------------- | ---------------------------- |
| `=`                   | Equality                     |
| `<>` or `!=`          | Non-equality                 |
| `<`                   | Less than                    |
| `<=`                  | Less than or equal to        |
| `!<`                  | Not less than                |
| `>`                   | Greater than                 |
| `>=`                  | Greater than or equal to     |
| `!>`                  | Not greater than             |
| `BETWEEN <a> AND <b>` | Between two specified values |
| `IS NULL`             | Is a `NULL` value.           |

### Wildcards

To use wildcards in search clauses, the `LIKE` operator must be used. Wildcard
searching can only be used with text fields. The `%` sign matches any number of
occurrences of any character. The underscore (`_`) matches any single character.

```sql
SELECT email FROM People WHERE email LIKE '%@gmail.com';

-- Find People (babies) with single digit heights.
SELECT height FROM People WHERE height LIKE '_ inches';
```

### Calculated Fields

A calculated field is created on-the-fly within a SQL `SELECT` statement.
Concatenate strings with the `||` operator. SQL supports mathematical operators
like `+`, `-`, `*`, and `/`.

```sql
SELECT last_name || ', ' || first_name FROM People;

SELECT item_price * quantity FROM Orders;
```

An **alias** is an alternate name for a field or value. Aliases are assigned
with the `AS` keyword.

```sql
SELECT last_name || ', ' || first_name AS name FROM People ORDER BY name;
```

SQL also provides a variety of functions for manipulating values. Unfortunately,
functions tend to be very DBMS specific. Unlike SQL statements, SQL functions
are not portable. Be sure to look at documentation for your specific SQL
implementation!

### Aggregation

SQL supports five efficient aggregation functions.

| Function | Description                           |
| -------- | ------------------------------------- |
| AVG()    | Return a column's average value       |
| COUNT()  | Return the number of rows in a column |
| MAX()    | Return a column's highest value       |
| MIN()    | Return a column's lowest value        |
| SUM()    | Return the sum of a column's values   |

The `COUNT` function can be used in a few different ways.

```sql
-- Count the number of rows in a table.
SELECT COUNT(*) FROM People;

-- Count the number of rows that have values in a specific column.
SELECT COUNT(email) FROM People;

-- Count unique values.
SELECT COUNT(DISTINCT age) FROM People;
```

### Grouping

`GROUP BY` divides data into logical sets so that you can perform aggregate
calculations on each group. Note that groups may be nested by specifying
multiple columns in the `GROUP BY` clause. Filter groups with the `HAVING`
clause. `HAVING` filters groups, `WHERE` filters rows.

```sql
SELECT age, COUNT(*) AS num_people FROM People GROUP BY age;

SELECT zip_code, COUNT(*) AS num_people FROM People GROUP BY age HAVING COUNT(*) > 2;
```

### Subqueries

**Subqueries** are queries embedded into other queries. SQL imposes no limit on
the number of subqueries that can be nested. Note that subquery `SELECT`
statements can only retrieve a single column. Attempting to retrieve multiple
columns will return an error.

```sql
SELECT last_name, first_name FROM People
WHERE id IN (SELECT customer_id FROM Orders WHERE product_id = 32);
```

Joins are usually more performant than subqueries.

### Joins

Relational tables are designed so that information is split into multiple
tables, one for each data type. The tables are related to each other through
common values. A **join** is a mechanism used to associate tables within a
`SELECT` statement.

Note that you must use the fully qualified column name (table and column
separated by a period) whenever there is a possible ambiguity about which column
you are referring to.

#### Inner Join

The common **inner join** (or **equijoin**) is a join based on the testing of
equality between two tables.

```sql
-- Implicit inner join.
SELECT vendor_name, product_name FROM Vendors, Products
WHERE Vendors.id = Products.vendor_id;

-- Explicit inner join (preferred).
SELECT vendor_name, product_name FROM Vendors
INNER JOIN Products ON Vendors.id = Products.vendor_id;
```

#### Outer Join

Suppose, you're hoping to find the number of products sold by each vendor. Some
of these counts might be zero. The **outer join** lets you include rows with no
mutual relation.

```sql
SELECT vendor_name, COUNT(*) FROM Vendors
LEFT OUTER JOIN Products ON Vendors.id = Products.vendor_id
```

The `RIGHT` or `LEFT` keyword must be used to specify the table from which to
include all rows. Some SQL implementations also provide a `FULL` keyword that
retrieves all rows from both tables.

### Combine Results

A `UNION` is composed of two or more `SELECT` statements, each separated by the
keyword `UNION`. Each query in a `UNION` must contain the same columns,
expressions, or aggregate functions.

```sql
SELECT first_name, last_name FROM People WHERE state IN ('MA', 'IL')
UNION
SELECT first_name, last_name FROM People WHERE age = 50;
```

Note that this specific query could have been performed with multiple `WHERE`
clauses instead.

## `INSERT`

`INSERT` adds rows to a database table. Complete or partial rows may be
inserted.

```sql
-- Insert a complete row.
INSERT INTO People
VALUES ('Elliot', 'Penson', 26, 'DC', 'elliotpenson@gmail.com', NULL);

-- Insert a complete row with explicit columns (highly preferred).
INSERT INTO People(first_name, last_name, age, state, email, favorite_color)
VALUES ('Elliot', 'Penson', 26, 'DC', 'elliotpenson@gmail.com', NULL);

-- Insert a partial row.
INSERT INTO People(first_name, last_name) VALUES ('Elliot', 'Penson');
```

## `UPDATE`

The `UPDATE` statement modifies data in a table. The format includes the table
name, column names/values, and a filter condition. Omitting the filter condition
will update all rows in the table.

```sql
-- Update a column.
UPDATE People SET age = 27 WHERE first_name = 'Elliot';

-- Update multiple columns.
UPDATE People SET state = 'CA', email = 'elliot@company.com'
WHERE first_name = 'Elliot';
```

## `DELETE`

Remove rows from a table with `DELETE`. Be careful to include a `WHERE` clause!

```sql
-- Delete all Toms.
DELETE FROM People WHERE first_name = 'Tom';

-- Delete all rows!
DELETE FROM People;
```

## `CREATE TABLE`

Use the `CREATE TABLE` statement to create a table. Specify the name and a
series of column definitions.

```sql
CREATE TABLE People (
    id         INT      NOT NULL  PRIMARY KEY,
    first_name TEXT     NOT NULL,
    last_name  TEXT     NOT NULL,
    age        INT      NOT NULL  DEFAULT 0,
    state      CHAR(2)  NOT NULL,
    email      TEXT     NULL,
    birthdate  DATETIME NOT NULL  DEFAULT NOW()
);
```

`NULL` columns permit `NULL` values, `NOT NULL` columns do not accept rows with
no value. `NULL`. The `DEFAULT` keyword is used to specify default values.

Standard SQL types include `INT`, `SMALLINT`, `REAL`, `FLOAT`, `CHAR(N)`,
`TEXT`, `DATE`, `TIME`, `TIMESTAMP`. Note that some implementations (like
PostgresSQL) provide other types.

## Constraints

Relational databases store related data in multiple tables. DBMSs enforce
**referential integrity** by imposing constraints on database tables.
**Constraints** are rules that govern how database data is inserted or
manipulated.

### Primary Keys

A **primary key** is a special constraint used to ensure that values in a column
(or set of columns) are unique and never change. No two rows may have the same
primary key value. Use the `PRIMARY KEY` keyword in a table definition to define
a primary key.

### Foreign Keys

A **foreign key** is a column whose values must be listed in the primary key of
another table. Foreign keys are essential for ensuring referential integrity.
Use the `REFERENCES` keyword to define a foreign key.

```sql
CREATE TABLE Orders (
    order_number INT      NOT NULL  PRIMARY KEY,
    order_date   DATETIME NOT NULL,
    customer_id  INT      NOT NULL REFERENCES People(id)
);
```

### Unique Constraints

Unique constraints are used to ensure that all data in a column is unique.
Unique constraints are similar to primary keys. In contrast, though, a table can
contain multiple constraints, but only one primary key. Unique constraints can
also be modified or updated. Use the `UNIQUE` keyword to define a unique
constraint.

### Check Constraints

Check constraints ensure that data in a column (or set of columns) meet a set of
criteria. Common criteria include minimum or maximum values, a range, and
specific values. Use the `CHECK` keyword to define a check constraint.

```sql
CREATE TABLE Order (
   order_number INT      NOT NULL PRIMARY KEY,
   product_id   CHAR(10) NOT NULL,
   quantity     INT      NOT NULL CHECK (quantity > 0),
   price        MONEY    NOT NULL
);
```

## `ALTER TABLE`

```sql
-- Add a column.
ALTER TABLE People ADD city TEXT;

-- Remove a column.
ALTER TABLE People DROP COLUMN city;
```

## `DROP TABLE`

```sql
DROP TABLE People;
```

## Transactions

Transaction processing is used to maintain database integrity by ensuring that
batches of SQL operations execute completely or not at all. If no error occurs,
the entire set of statements is **committed**. If an error occurs, a
**rollback** can restore the database to a known and safe state. Note that
individual statements are implicitly a single transaction.

Unfortunately, transaction processing is implemented differently in each DBMS.
In PostgreSQL, a transaction is set up by surrounding the SQL commands with
`BEGIN` and `COMMIT`. A `ROLLBACK` will cancel the updates.

```sql
-- PostgreSQL transaction.
BEGIN;
...
COMMIT;
```

**Savepoints** allow you to selectively discard parts of a transaction while
committing the rest.

```sql
BEGIN;
...
SAVEPOINT my_savepoint;
...
ROLLBACK TO my_savepoint;
...
COMMIT;
```

## Indexes

Searching for specific column values can often be inefficient. The DBMS might
have to read every row in the table looking for matches. **Indexes** are a copy
of selected columns organized logically to improve the speed of searching and
sorting. These data structures can improve lookup from linear to logarithmic or
even constant time.

While indexes improve the performance of retrieval operations, they degrade the
performance of data insertion, modification, and deletion. Index data can also
take up significant storage space.

```sql
CREATE INDEX last_name_index ON People (last_name);

DROP INDEX last_name_index;
```

## PostgreSQL

PostgreSQL (or Postgres) is an open-source SQL implementation. See
[the documentation](https://www.postgresql.org/docs/11/index.html) for more
information.

On macOS, PostgreSQL can be installed via homebrew (i.e.
`brew install postgres`). PostgreSQL uses a client/server model. Run
`brew services start postgresql` to begin the server. Create a new database with
`createdb <db-name>`, delete a database with `dropdb <db-name>`.

`psql` is the PostgreSQL terminal client. This application allows one to
interactively enter, edit, and execute SQL commands. The command takes the form
`psql <db-name>`.

| Meta-Command        | Description                                 |
| ------------------- | ------------------------------------------- |
| \list               | List all databases.                         |
| \connect <database> | Connect to a database.                      |
| \dt                 | Display all tables in the current database. |
| \dn                 | List schemas (namespaces).                  |
| \d <̣table>          | Describe a table (show columns)             |
| \q                  | Quit psql                                   |

## Tricks and Examples

### Remove Duplicates

Suppose we have a `People` table with `Id` and `Email` columns. The following
command will delete all duplicate email entries keeping entries based on the
smallest ID.

```sql
DELETE * FROM People p1
INNER JOIN People p2
WHERE p1.Email = p2.Email AND p1.Id > p2.Id;
```
