# SQL Questions for Backend Developers (2 Years of Experience)

# **Basic SQL Queries**
1. How would you retrieve the top 5 highest salaries from an `employees` table?

   **Answer:**
   ```sql
   SELECT * FROM employees ORDER BY salary DESC LIMIT 5;
   ```

2. Write a query to fetch all records where the `created_at` column is within the last 30 days.
   ```sql
   SELECT * FROM table_name WHERE created_at >= CURRENT_DATE - INTERVAL '30 days';
   ```

3. How would you delete duplicate rows in a table while keeping one copy?
   ```sql
   DELETE FROM table_name
   WHERE id NOT IN (
       SELECT MIN(id)
       FROM table_name
       GROUP BY column1, column2, ...
   );
   ```

4. Write a query to find employees who don’t have a manager (assuming `manager_id` can be null).
   ```sql
   SELECT * FROM employees WHERE manager_id IS NULL;
   ```

5. Write a query to fetch the first 10 rows of a table, ordered by a specific column.
   ```sql
   SELECT * FROM table_name ORDER BY column_name LIMIT 10;
   ```

6. How would you fetch all records where a string column starts with a specific prefix?
   ```sql
   SELECT * FROM table_name WHERE column_name LIKE 'prefix%';
   ```

7. Write a query to find all records where a column matches a value in a given list of values.
   ```sql
   SELECT * FROM table_name WHERE column_name IN ('value1', 'value2', 'value3');
   ```

8. How would you retrieve only unique values from a specific column?
   ```sql
   SELECT DISTINCT column_name FROM table_name;
   ```

9. Write a query to fetch all rows where the value of one column is greater than the value of another column in the same table.
   ```sql
   SELECT * FROM table_name WHERE column1 > column2;
   ```

10. How would you calculate the difference between two date columns in days?
    ```sql
    SELECT column1 - column2 AS day_difference FROM table_name;
    ```

---

# **Joins and Relationships**
1. Explain the difference between INNER JOIN, LEFT JOIN, RIGHT JOIN, and FULL OUTER JOIN.
   - **INNER JOIN**: Returns only the rows that have matching values in both tables.
   - **LEFT JOIN**: Returns all rows from the left table and the matching rows from the right table.
   - **RIGHT JOIN**: Returns all rows from the right table and the matching rows from the left table.
   - **FULL OUTER JOIN**: Returns all rows when there is a match in either table.

2. Write a query to retrieve all orders along with the customer details, even if the customer does not have any orders.
   ```sql
   SELECT * FROM customers
   LEFT JOIN orders ON customers.customer_id = orders.customer_id;
   ```

3. How would you find all products that have never been sold, using a `products` table and an `orders` table?
   ```sql
   SELECT * FROM products
   WHERE product_id NOT IN (
       SELECT product_id FROM orders
   );
   ```

4. Write a query to fetch all employees who share the same manager.
   ```sql
   SELECT * FROM employees e1
   JOIN employees e2 ON e1.manager_id = e2.manager_id
   WHERE e1.id != e2.id;
   ```

5. Write a query to retrieve data from three tables using JOIN.
   ```sql
   SELECT *
   FROM table1
   JOIN table2 ON table1.column = table2.column
   JOIN table3 ON table2.column = table3.column;
   ```

6. How would you join two tables based on multiple conditions?
   ```sql
   SELECT *
   FROM table1
   JOIN table2 ON table1.column1 = table2.column1 AND table1.column2 = table2.column2;
   ```

7. Write a query to retrieve all customers and their orders, and include a count of orders per customer.
   ```sql
   SELECT customers.customer_id, COUNT(orders.order_id) AS order_count
   FROM customers
   LEFT JOIN orders ON customers.customer_id = orders.customer_id
   GROUP BY customers.customer_id;
   ```

8. How would you join a table to itself (self-join)? Provide an example use case.
   ```sql
   SELECT e1.name AS employee, e2.name AS manager
   FROM employees e1
   JOIN employees e2 ON e1.manager_id = e2.id;
   ```

9. Write a query to find all pairs of employees who work in the same department.
   ```sql
   SELECT e1.name AS employee1, e2.name AS employee2
   FROM employees e1
   JOIN employees e2 ON e1.department_id = e2.department_id
   WHERE e1.id < e2.id;
   ```

10. How would you use a CROSS JOIN to generate all possible combinations of two tables?
    ```sql
    SELECT *
    FROM table1
    CROSS JOIN table2;
    ```

---

# **Aggregation and Grouping**
1. Write a query to count the number of orders placed by each customer.
   ```sql
   SELECT customer_id, COUNT(*) AS order_count
   FROM orders
   GROUP BY customer_id;
   ```

2. How would you calculate the total revenue generated from orders in the last quarter?
   ```sql
   SELECT SUM(order_amount) AS total_revenue
   FROM orders
   WHERE order_date >= DATE_TRUNC('quarter', CURRENT_DATE) - INTERVAL '1 quarter'
     AND order_date < DATE_TRUNC('quarter', CURRENT_DATE);
   ```

3. Write a query to find the maximum, minimum, and average order amount for each customer.
   ```sql
   SELECT customer_id, MAX(order_amount) AS max_order, MIN(order_amount) AS min_order, AVG(order_amount) AS avg_order
   FROM orders
   GROUP BY customer_id;
   ```

4. How would you group data by month and calculate the total sales for each month?
   ```sql
   SELECT DATE_TRUNC('month', order_date) AS month, SUM(order_amount) AS total_sales
   FROM orders
   GROUP BY month
   ORDER BY month;
   ```

5. Write a query to find departments with more than 10 employees.
   ```sql
   SELECT department_id, COUNT(*) AS employee_count
   FROM employees
   GROUP BY department_id
   HAVING COUNT(*) > 10;
   ```

6. How would you count the number of distinct values in a column?
   ```sql
   SELECT COUNT(DISTINCT column_name) FROM table_name;
   ```

7. Write a query to calculate the average salary for employees in each department.
   ```sql
   SELECT department_id, AVG(salary) AS avg_salary
   FROM employees
   GROUP BY department_id;
   ```

8. Write a query to list the top 3 customers by total spending.
   ```sql
   SELECT customer_id, SUM(order_amount) AS total_spending
   FROM orders
   GROUP BY customer_id
   ORDER BY total_spending DESC
   LIMIT 3;
   ```

9. How would you calculate the percentage of total sales contributed by each product?
   ```sql
   SELECT product_id, SUM(order_amount) AS product_sales,
          (SUM(order_amount) / SUM(SUM(order_amount)) OVER()) * 100 AS percentage_contribution
   FROM orders
   GROUP BY product_id;
   ```

10. Write a query to find the most frequently ordered product.
    ```sql
    SELECT product_id, COUNT(*) AS order_count
    FROM orders
    GROUP BY product_id
    ORDER BY order_count DESC
    LIMIT 1;
    


# **Indexes and Optimization**

1. **What are indexes in SQL, and how do they improve query performance?**

Indexes are special database objects that improve query performance by allowing the database to locate rows more efficiently. They work by creating a data structure (usually a B-tree or hash index) to quickly look up values in a table. Without indexes, SQL queries would need to scan the entire table, which can be slow for large datasets.

2. **Explain the difference between clustered and non-clustered indexes.**

- **Clustered index**: The table rows are physically stored in the order of the index. A table can have only one clustered index.
- **Non-clustered index**: The index is stored separately from the table data, and it contains pointers to the table rows. A table can have multiple non-clustered indexes.

3. **How would you find all the indexes on a specific table?**

```sql
SELECT indexname, indexdef
FROM pg_indexes
WHERE tablename = 'your_table_name';
```

4. **If a query is running slow, how would you approach optimizing it?**

- Check the query execution plan using `EXPLAIN` to identify bottlenecks.
- Add or optimize indexes.
- Avoid using `SELECT *` and select only the necessary columns.
- Consider using joins instead of subqueries, if applicable.
- Use appropriate data types for columns.
- Analyze the data distribution and update statistics if necessary.

5. **What is a covering index, and when would you use it?**

A covering index is an index that includes all the columns needed for a query, so the database can fulfill the query using just the index, without needing to access the table itself. You would use it when you want to speed up read queries by avoiding table lookups.

6. **Write a query to check if a specific column has an index.**

```sql
SELECT *
FROM pg_indexes
WHERE indexdef LIKE '%column_name%'
  AND tablename = 'your_table_name';
```

7. **What are some common causes of slow queries, and how can they be resolved?**

- Missing indexes: Add appropriate indexes to speed up lookups.
- Complex joins: Simplify joins or break them into smaller queries.
- Large result sets: Use pagination or filtering to reduce the amount of data processed.
- Lack of query optimization: Analyze the execution plan and refactor the query.

8. **How would you use ANALYZE or EXPLAIN to understand query performance?**

- **EXPLAIN**: Shows the execution plan of a query, detailing how tables are scanned and joined, and what indexes are used.
- **ANALYZE**: Collects statistics on tables and indexes, which helps the query planner make better decisions.

9. **Explain the impact of too many indexes on table performance.**

Too many indexes can degrade write performance (inserts, updates, deletes) because the indexes must also be updated when the data changes. Additionally, maintaining multiple indexes increases disk space usage and can slow down data modification operations.

10. **How do composite indexes work, and when should they be used?**

A composite index is an index on multiple columns. It is used when queries frequently filter or sort by more than one column. Composite indexes should be created in the order that matches the query patterns to be most effective.

---

# **Subqueries and CTEs**

1. **What is the difference between a subquery and a Common Table Expression (CTE)?**

A **subquery** is a query nested inside another query, whereas a **CTE** (Common Table Expression) is a temporary result set that you can reference within a `SELECT`, `INSERT`, `UPDATE`, or `DELETE` statement. CTEs are often more readable and reusable than subqueries.

2. **Write a query using a CTE to find the second-highest salary in an employees table.**

```sql
WITH SalaryRank AS (
  SELECT salary, RANK() OVER (ORDER BY salary DESC) AS rank
  FROM employees
)
SELECT salary
FROM SalaryRank
WHERE rank = 2;
```

3. **Write a query to list all employees who earn more than the average salary in their department.**

```sql
SELECT employee_id, name, salary
FROM employees e
WHERE salary > (
  SELECT AVG(salary)
  FROM employees
  WHERE department_id = e.department_id
);
```

4. **How would you write a subquery to find customers who have placed more than 5 orders?**

```sql
SELECT customer_id
FROM customers
WHERE customer_id IN (
  SELECT customer_id
  FROM orders
  GROUP BY customer_id
  HAVING COUNT(order_id) > 5
);
```

5. **Write a query using a correlated subquery to find employees who earn more than the average salary of their department.**

```sql
SELECT employee_id, name, salary
FROM employees e
WHERE salary > (
  SELECT AVG(salary)
  FROM employees
  WHERE department_id = e.department_id
);
```

6. **What are the advantages of using a CTE over a subquery?**

- **Readability**: CTEs can make complex queries easier to read and maintain by separating the logic into a temporary result set.
- **Reusability**: CTEs can be referenced multiple times within a query, whereas a subquery can only be used once.
- **Recursion**: CTEs allow for recursive queries, which are not possible with subqueries.

7. **Write a query using a recursive CTE to generate a hierarchy of employees and their managers.**

```sql
WITH RECURSIVE EmployeeHierarchy AS (
  SELECT employee_id, manager_id, name
  FROM employees
  WHERE manager_id IS NULL
  UNION ALL
  SELECT e.employee_id, e.manager_id, e.name
  FROM employees e
  JOIN EmployeeHierarchy eh ON e.manager_id = eh.employee_id
)
SELECT * FROM EmployeeHierarchy;
```

8. **How would you use a CTE to paginate results in a query?**

```sql
WITH PaginatedResults AS (
  SELECT *, ROW_NUMBER() OVER (ORDER BY column_name) AS row_num
  FROM your_table
)
SELECT * FROM PaginatedResults
WHERE row_num BETWEEN 11 AND 20;
```

9. **Write a query using a subquery to find products with sales below the average.**

```sql
SELECT product_id, product_name
FROM products
WHERE sales < (
  SELECT AVG(sales)
  FROM products
);
```

10. **How would you rewrite a query with multiple subqueries to improve performance?**

Instead of using multiple subqueries, you could:
- Use **joins** to bring together data from multiple tables.
- Use **CTEs** to break the query into manageable pieces and avoid repeating the same subquery logic.
- Check if materialized views can be used for complex subqueries that don't need to be computed on every query.



# **Transactions**

1. **What is a transaction in SQL, and why is it important?**

   A transaction in SQL is a set of one or more SQL queries that are executed as a single unit. The primary purpose of a transaction is to ensure data integrity and consistency by guaranteeing that either all the queries within the transaction are successfully executed (commit), or none of them are (rollback). Transactions ensure that the database remains in a consistent state even in the event of system failures or errors.

2. **Explain the purpose of COMMIT, ROLLBACK, and SAVEPOINT in SQL transactions.**

   - **COMMIT**: The COMMIT statement is used to save all changes made during the current transaction. Once the changes are committed, they become permanent in the database.
   - **ROLLBACK**: The ROLLBACK statement is used to undo all changes made during the current transaction. It returns the database to its state before the transaction began.
   - **SAVEPOINT**: A SAVEPOINT creates a point within a transaction that you can later roll back to, without affecting the entire transaction. It allows for partial rollbacks.

3. **Write a query to transfer $100 from one account to another, ensuring the transaction is atomic.**

   ```sql
   BEGIN;
   UPDATE accounts SET balance = balance - 100 WHERE account_id = 1;
   UPDATE accounts SET balance = balance + 100 WHERE account_id = 2;
   COMMIT;
This ensures that both operations (decreasing from one account and increasing in the other) happen together. If either fails, neither change will be committed.

## 4. How would you handle errors during a transaction in SQL?
To handle errors in SQL transactions, you can use TRY...CATCH blocks (in databases like SQL Server) or use BEGIN...EXCEPTION in PostgreSQL. If an error occurs, you can issue a ROLLBACK to undo the changes made in the transaction up to that point.
```
BEGIN;

-- Perform some operations
UPDATE accounts
SET balance = balance - 100
WHERE account_id = 1;

-- Simulating an error
-- Invalid query or constraint violation
UPDATE accounts
SET balance = balance + 100
WHERE account_id = 'nonexistent';

-- If an error occurs, rollback
ROLLBACK;
```
## 5. What is the difference between ACID properties and transaction isolation levels?

**ACID properties** refer to the four key properties of transactions:

**Atomicity:** Transactions are all-or-nothing.
**Consistency:** Transactions move the database from one valid state to another.

**Isolation:** Transactions are isolated from one another.

**Durability:** Once a transaction is committed, it will persist even in case of system failure.

**Transaction Isolation Levels** define the degree to which the operations in one transaction are isolated from the operations of other concurrent transactions.

## 6. Explain the different transaction isolation levels in SQL.
SQL defines four transaction isolation levels:

**Read Uncommitted:** Allows reading uncommitted data, meaning dirty reads are possible.

**Read Committed:** Guarantees that only committed data is read, but non-repeatable reads are possible.

**Repeatable Read:** Ensures that once a value is read, it will not change during the transaction, preventing non-repeatable reads.

**Serializable:** Provides the highest level of isolation by ensuring transactions are executed serially, preventing phantom reads.

## 7. Write a query to demonstrate the use of SAVEPOINT in a transaction.

```
BEGIN;

-- Set a savepoint
SAVEPOINT before_update;

UPDATE accounts
SET balance = balance - 50
WHERE account_id = 1;

-- If there is an error, roll back to the savepoint
-- ROLLBACK TO before_update;

UPDATE accounts
SET balance = balance + 50
WHERE account_id = 2;

-- Commit the changes if no issues
COMMIT;
```

## 8. How would you ensure data consistency when multiple users access the same table?
To ensure data consistency, you can use transactions with appropriate isolation levels. For example, using Repeatable Read or Serializable isolation levels prevents other users from modifying the data while one transaction is in progress. Additionally, locking mechanisms can be used to control concurrent access to critical data.

## 9. What is a deadlock, and how can it be avoided in transactions?
A deadlock occurs when two or more transactions are waiting on each other to release locks, causing a standstill. To avoid deadlocks:

- Ensure transactions are kept as short as possible.
- Access tables in a consistent order.
- Use lower isolation levels where possible to reduce lock contention.
- Use the NOWAIT or SKIP LOCKED options in certain databases to handle locking issues.

## 10. How would you test the behavior of a transaction under high concurrency?
To test a transaction under high concurrency, you can simulate multiple users attempting to perform operations on the same data simultaneously. Tools like pgbench (for PostgreSQL) or load testing frameworks can be used. It's important to check how the transaction behaves with multiple concurrent reads and writes, ensuring data integrity and checking for deadlocks or race conditions.
```

This markdown structure provides concise answers to each question with clear explanations and code examples where applicable.
```


# **Advanced SQL Topics**

## 1. **What is a database view, and how would you use one?**

A **database view** is a virtual table that is based on the result set of a SQL query. It does not store data physically but provides a way to simplify complex queries or aggregate data. Views are useful for encapsulating complex joins, filters, or calculations and can be used to provide a simplified interface to users or applications.

### Example:

```sql
CREATE VIEW active_employees AS
SELECT id, name, department, salary
FROM employees
WHERE status = 'active';
```

---

## 2. **Explain the difference between SQL’s HAVING and WHERE clauses with examples.**

- **WHERE**: The WHERE clause is used to filter rows before any grouping or aggregation occurs. It filters individual records from the database.
- **HAVING**: The HAVING clause is used to filter groups after aggregation. It is used in combination with GROUP BY to filter aggregated results.

### Example:

```sql
-- WHERE clause
SELECT name, salary
FROM employees
WHERE salary > 50000;

-- HAVING clause
SELECT department, AVG(salary) AS avg_salary
FROM employees
GROUP BY department
HAVING AVG(salary) > 60000;
```

---

## 3. **Write a query to fetch all records where a specific column has a JSON value containing a specific key or value.**

In PostgreSQL, you can use the `->` or `->>` operators to extract values from a JSON column. The `jsonb` type allows for efficient querying of JSON data.

### Example:

```sql
SELECT *
FROM employees
WHERE employee_data->>'job_title' = 'Developer';
```

This query fetches all records where the `employee_data` column (containing JSON data) has a key `job_title` with the value `'Developer'`.

---

## 4. **What are window functions, and how would you use one to calculate a running total?**

Window functions perform calculations across a set of table rows that are related to the current row. They are often used for running totals, rankings, and other cumulative or comparative metrics without collapsing rows.

### Example:

```sql
SELECT employee_id, salary,
       SUM(salary) OVER (ORDER BY employee_id) AS running_total
FROM employees;
```

This query calculates the running total of salaries ordered by `employee_id`.

---

## 5. **Write a query using a window function to rank employees based on their salaries within their department.**

You can use the `RANK()` or `ROW_NUMBER()` function to rank employees based on their salary within their department.

### Example:

```sql
SELECT employee_id, name, salary, department,
       RANK() OVER (PARTITION BY department ORDER BY salary DESC) AS salary_rank
FROM employees;
```

This query ranks employees within each department based on their salary, with the highest salary ranked first.

---

## 6. **How would you store hierarchical data in a relational database?**

Hierarchical data can be stored in relational databases using the following approaches:
- **Adjacency List Model**: Each record has a `parent_id` column referring to its parent.
- **Nested Set Model**: Uses `left` and `right` values to represent hierarchy.
- **Path Enumeration**: Stores the full path of each node.
- **Common Table Expressions (CTE)**: Recursive CTEs allow querying hierarchical data dynamically.

### Example (Adjacency List Model):

```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    parent_id INT REFERENCES categories(id)
);
```

---

## 7. **Write a query to pivot rows into columns in SQL.**

Pivoting data involves transforming rows into columns to make it more readable.

### Example (Using `CASE WHEN`):

```sql
SELECT employee_id,
       SUM(CASE WHEN month = 'January' THEN salary END) AS January,
       SUM(CASE WHEN month = 'February' THEN salary END) AS February
FROM salaries
GROUP BY employee_id;
```

This converts row-based salary records into columns for each month.

---

## 8. **What is a materialized view, and how is it different from a regular view?**

A **materialized view** is a database object that stores the result of a query physically, unlike a regular view, which is virtual and recalculates data on every query.

- **Regular View**: Always fetches fresh data but may be slower.
- **Materialized View**: Faster but needs to be refreshed to get updated data.

### Example:

```sql
CREATE MATERIALIZED VIEW sales_summary AS
SELECT product_id, SUM(sales) AS total_sales
FROM sales
GROUP BY product_id;
```

To refresh:

```sql
REFRESH MATERIALIZED VIEW sales_summary;
```

---

## 9. **How would you use the `JSONB` data type in PostgreSQL for storing semi-structured data?**

The `JSONB` data type in PostgreSQL allows efficient storage and querying of JSON data.

### Example:

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT,
    attributes JSONB
);

INSERT INTO products (name, attributes)
VALUES ('Laptop', '{"brand": "Dell", "ram": "16GB"}');
```

Querying JSONB data:

```sql
SELECT * FROM products WHERE attributes->>'brand' = 'Dell';
```

---

## 10. **What is sharding, and when would you use it?**

Sharding is a database partitioning technique that divides a large dataset into smaller, more manageable pieces (shards) distributed across multiple database instances. This improves performance and scalability.

### When to Use Sharding:
- When dealing with **large-scale** data.
- When **horizontal scaling** is needed.
- When reducing **query load** on a single database server.

### Example:

- Users from different regions can be stored in different shards.
- An application can route user queries to the correct shard based on user ID.

```sql
-- Shard 1 (users from North America)
CREATE TABLE users_na (
    id SERIAL PRIMARY KEY,
    name TEXT,
    region TEXT
);

-- Shard 2 (users from Europe)
CREATE TABLE users_eu (
    id SERIAL PRIMARY KEY,
    name TEXT,
    region TEXT
);
```
---

#  **Database Design**

### 1. What are the differences between normalization and denormalization? When would you use each?
Normalization is the process of organizing data to reduce redundancy and improve integrity. It involves dividing large tables into smaller ones and defining relationships between them. Denormalization, on the other hand, is the process of combining tables to improve read performance at the cost of some redundancy.

- Use **normalization** when ensuring data integrity and minimizing redundancy are priorities.
- Use **denormalization** when optimizing for read-heavy operations where performance is more important than data redundancy.

### 2. How would you design a database schema to handle users, roles, and permissions?
A common approach is to use three tables:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE user_roles (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    role_id INT REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);
```
This structure allows assigning multiple roles to users while maintaining referential integrity.

### 3. What is a foreign key constraint, and what are the different ON DELETE behaviors?
A foreign key constraint ensures referential integrity by linking a column to another table's primary key. ON DELETE behaviors define what happens when a referenced record is deleted:

- **CASCADE**: Deletes related records automatically.
- **SET NULL**: Sets the foreign key to NULL.
- **SET DEFAULT**: Sets the foreign key to a default value.
- **RESTRICT/NO ACTION**: Prevents deletion if there are dependent records.

### 4. How would you enforce uniqueness across multiple columns in a table?
You can enforce uniqueness using a **unique constraint** across multiple columns:

```sql
CREATE TABLE employee_assignments (
    employee_id INT,
    project_id INT,
    assignment_date DATE,
    UNIQUE (employee_id, project_id)
);
```
This ensures that an employee can be assigned to a project only once.

### 5. What are some strategies for designing a schema to handle multi-tenancy?
- **Single Database, Shared Schema:** All tenants share tables, with a `tenant_id` column.
- **Single Database, Separate Schemas:** Each tenant has a dedicated schema.
- **Separate Databases:** Each tenant has its own database.

The best approach depends on scalability and isolation requirements.

### 6. How would you design a database to handle soft deletes?
Use a `deleted_at` column to track soft deletions instead of removing records:

```sql
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP;
```
Queries should filter out soft-deleted records:

```sql
SELECT * FROM users WHERE deleted_at IS NULL;
```

### 7. Explain the purpose of surrogate keys versus natural keys in database design.
- **Surrogate Key**: An artificial primary key (e.g., auto-incremented ID) that has no business meaning.
- **Natural Key**: A primary key derived from meaningful business data (e.g., email or SSN).

Surrogate keys are preferred for consistency and stability, while natural keys can be useful when a unique business identifier already exists.

### 8. How would you implement a many-to-many relationship in a database?
Use a join table to connect the two entities:

```sql
CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL
);

CREATE TABLE author_books (
    author_id INT REFERENCES authors(id),
    book_id INT REFERENCES books(id),
    PRIMARY KEY (author_id, book_id)
);
```
This allows multiple authors per book and multiple books per author.

### 9. What are the advantages and disadvantages of using stored procedures?
**Advantages:**
- Encapsulates business logic within the database.
- Improves performance by reducing network traffic.
- Enhances security by restricting direct table access.

**Disadvantages:**
- Harder to version-control compared to application code.
- Can lead to vendor lock-in due to database-specific syntax.
- Increased complexity in debugging and maintenance.

### 10. How would you design a schema to support versioning of records in a table?
One approach is to store versions of records in an audit table:

```sql
CREATE TABLE document_versions (
    id SERIAL PRIMARY KEY,
    document_id INT,
    version INT,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (document_id, version)
);
```
This ensures that each document can have multiple versions, allowing rollback or tracking changes over time.

---

# **Common Real-World Scenarios**

### 1. How would you handle a situation where your table is growing rapidly in size and causing performance issues?
To handle large table growth, consider these strategies:
- **Indexing**: Add indexes to optimize queries.
- **Partitioning**: Split large tables into smaller partitions.
- **Archiving**: Move old data to an archive table.
- **Sharding**: Distribute data across multiple databases.
- **Compression**: Use data compression techniques.

### 2. What strategies would you use to safely migrate a large database table without downtime?
- **Blue-Green Deployment**: Set up a new schema and gradually migrate.
- **Rolling Migrations**: Apply changes in phases.
- **Replication**: Create a copy, switch traffic when ready.
- **Backups**: Always take full backups before migration.
- **Database Views**: Abstract changes to avoid application disruptions.

### 3. Write a query to paginate results from a table, showing 10 results per page.
```sql
SELECT * FROM employees
ORDER BY id
LIMIT 10 OFFSET 20; -- Fetch page 3 (offset = (page-1) * limit)
```

### 4. Explain how you would design a database to handle soft deletes.
- Add a `deleted_at` column (NULL for active rows, timestamp when deleted).
- Use indexes to exclude deleted rows in queries.
- Implement database views to filter out deleted records.
- Periodically clean up soft-deleted data.

```sql
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP NULL;
UPDATE users SET deleted_at = NOW() WHERE id = 1;
```

### 5. How would you manage schema changes in a production database?
- **Use Migrations**: Track changes with migration scripts.
- **Feature Flags**: Roll out changes progressively.
- **Backward Compatibility**: Ensure old and new versions can coexist.
- **Zero-Downtime Deployments**: Use online schema changes.

### 6. What steps would you take to troubleshoot slow queries in a production database?
- **EXPLAIN ANALYZE**: Identify slow query steps.
- **Indexes**: Ensure queries use proper indexing.
- **Query Optimization**: Rewrite inefficient queries.
- **Caching**: Use caching mechanisms (Redis, Memcached).
- **Database Monitoring**: Analyze query logs and performance metrics.

### 7. Write a query to merge two tables with duplicate data.
```sql
INSERT INTO customers (id, name, email)
SELECT id, name, email FROM customers_backup
ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, email = EXCLUDED.email;
```

### 8. How would you handle data integrity issues caused by an application bug?
- **Audit Logs**: Track changes to identify errors.
- **Data Validation**: Apply constraints and triggers.
- **Fix Data**: Write correction scripts.
- **Prevent Recurrence**: Patch application bugs, add tests.

### 9. What techniques would you use to archive old data from a large table?
- **Partitioning**: Move old records to separate partitions.
- **Archival Tables**: Store historical data in a separate table.
- **ETL Processes**: Extract old data and store in data lakes.
- **Scheduled Jobs**: Automate archival with cron jobs.

### 10. How would you monitor and ensure database health over time?
- **Performance Metrics**: Track query execution times.
- **Index Maintenance**: Rebuild and optimize indexes.
- **Disk Usage**: Monitor and optimize storage.
- **Backup Strategy**: Ensure automated and tested backups.
- **Security Audits**: Regularly check for vulnerabilities.

---

# **Debugging and Maintenance**

### 1. What is the purpose of the `EXPLAIN` or `EXPLAIN PLAN` statement in SQL?
The `EXPLAIN` statement provides insight into how a SQL query will be executed, including details on indexes, join strategies, and sorting methods. It helps identify inefficiencies in queries and optimize them.

#### Example:
```sql
EXPLAIN ANALYZE SELECT * FROM employees WHERE department = 'HR';
```

### 2. How do you identify and fix deadlocks in a database?
Deadlocks occur when two transactions wait for each other to release resources. To identify deadlocks:
- Use database logs (`SHOW ENGINE INNODB STATUS;` in MySQL, `pg_stat_activity` in PostgreSQL).
- Monitor deadlock frequency.

To fix deadlocks:
- Use consistent transaction ordering.
- Implement shorter transactions.
- Reduce row-level locking where possible.

### 3. How would you detect missing indexes or unused indexes in a database?
- Use `EXPLAIN` to check if queries scan large amounts of data.
- Monitor slow query logs.
- In PostgreSQL, use:
  ```sql
  SELECT * FROM pg_stat_user_indexes WHERE idx_scan = 0;
  ```
  to find unused indexes.

### 4. Write a query to identify the most resource-intensive queries running on a database.
```sql
SELECT query, total_exec_time
FROM pg_stat_statements
ORDER BY total_exec_time DESC
LIMIT 10;
```
This retrieves the top 10 most time-consuming queries.

### 5. How would you monitor database performance in real-time?
- Use built-in tools like MySQL’s `SHOW PROCESSLIST`, PostgreSQL’s `pg_stat_activity`.
- Utilize monitoring tools like Prometheus, New Relic, or pgAdmin.

### 6. What tools would you use for database profiling?
- PostgreSQL: `pg_stat_statements`
- MySQL: `pt-query-digest`
- SQL Server: Query Store, Profiler
- Oracle: AWR (Automatic Workload Repository)

### 7. How would you handle log growth for a transactional database?
- Set log rotation policies.
- Compress or archive old logs.
- Configure database settings to auto-purge logs.
- Use `pg_logrotate` for PostgreSQL or `logrotate` for Linux systems.

### 8. How do you ensure database backups are reliable and restorable?
- Automate regular backups using cron jobs or built-in schedulers.
- Store backups in multiple locations.
- Regularly test restoring backups in a staging environment.
- Use checksums to verify backup integrity.

### 9. What is the difference between database repair and optimization?
- **Database repair**: Fixes corruption issues using commands like `REPAIR TABLE` (MySQL) or `pg_resetwal` (PostgreSQL).
- **Database optimization**: Improves performance by reorganizing indexes, vacuuming tables, or defragmenting storage.

### 10. How would you test changes to a database schema in a development environment?
- Use migration scripts (e.g., Flyway, Liquibase).
- Test changes in a local/staging database before applying them to production.
- Validate changes using integration tests and performance benchmarks.
- Use feature flags to enable incremental rollouts.














