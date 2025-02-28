# SQL Questions for Backend Developers (2 Years of Experience)

## **Basic SQL Queries**
1. How would you retrieve the top 5 highest salaries from an `employees` table?
 ## Answer;
           select *from employee

2. Write a query to fetch all records where the `created_at` column is within the last 30 days.
3. How would you delete duplicate rows in a table while keeping one copy?
4. Write a query to find employees who don’t have a manager (assuming `manager_id` can be null).
5. Write a query to fetch the first 10 rows of a table, ordered by a specific column.
6. How would you fetch all records where a string column starts with a specific prefix?
7. Write a query to find all records where a column matches a value in a given list of values.
8. How would you retrieve only unique values from a specific column?
9. Write a query to fetch all rows where the value of one column is greater than the value of another column in the same table.
10. How would you calculate the difference between two date columns in days?

---

## **Joins and Relationships**
1. Explain the difference between INNER JOIN, LEFT JOIN, RIGHT JOIN, and FULL OUTER JOIN.
2. Write a query to retrieve all orders along with the customer details, even if the customer does not have any orders.
3. How would you find all products that have never been sold, using a `products` table and an `orders` table?
4. Write a query to fetch all employees who share the same manager.
5. Write a query to retrieve data from three tables using JOIN.
6. How would you join two tables based on multiple conditions?
7. Write a query to retrieve all customers and their orders, and include a count of orders per customer.
8. How would you join a table to itself (self-join)? Provide an example use case.
9. Write a query to find all pairs of employees who work in the same department.
10. How would you use a CROSS JOIN to generate all possible combinations of two tables?

---

## **Aggregation and Grouping**
1. Write a query to count the number of orders placed by each customer.
2. How would you calculate the total revenue generated from orders in the last quarter?
3. Write a query to find the maximum, minimum, and average order amount for each customer.
4. How would you group data by month and calculate the total sales for each month?
5. Write a query to find departments with more than 10 employees.
6. How would you count the number of distinct values in a column?
7. Write a query to calculate the average salary for employees in each department.
8. Write a query to list the top 3 customers by total spending.
9. How would you calculate the percentage of total sales contributed by each product?
10. Write a query to find the most frequently ordered product.

---

## **Indexes and Optimization**
1. What are indexes in SQL, and how do they improve query performance?
2. Explain the difference between clustered and non-clustered indexes.
3. How would you find all the indexes on a specific table?
4. If a query is running slow, how would you approach optimizing it?
5. What is a covering index, and when would you use it?
6. Write a query to check if a specific column has an index.
7. What are some common causes of slow queries, and how can they be resolved?
8. How would you use `ANALYZE` or `EXPLAIN` to understand query performance?
9. Explain the impact of too many indexes on table performance.
10. How do composite indexes work, and when should they be used?

---

## **Subqueries and CTEs**
1. What is the difference between a subquery and a Common Table Expression (CTE)?
2. Write a query using a CTE to find the second-highest salary in an `employees` table.
3. Write a query to list all employees who earn more than the average salary in their department.
4. How would you write a subquery to find customers who have placed more than 5 orders?
5. Write a query using a correlated subquery to find employees who earn more than the average salary of their department.
6. What are the advantages of using a CTE over a subquery?
7. Write a query using a recursive CTE to generate a hierarchy of employees and their managers.
8. How would you use a CTE to paginate results in a query?
9. Write a query using a subquery to find products with sales below the average.
10. How would you rewrite a query with multiple subqueries to improve performance?

---

## **Transactions**
1. What is a transaction in SQL, and why is it important?
2. Explain the purpose of COMMIT, ROLLBACK, and SAVEPOINT in SQL transactions.
3. Write a query to transfer $100 from one account to another, ensuring the transaction is atomic.
4. How would you handle errors during a transaction in SQL?
5. What is the difference between ACID properties and transaction isolation levels?
6. Explain the different transaction isolation levels in SQL.
7. Write a query to demonstrate the use of SAVEPOINT in a transaction.
8. How would you ensure data consistency when multiple users access the same table?
9. What is a deadlock, and how can it be avoided in transactions?
10. How would you test the behavior of a transaction under high concurrency?

---

## **Advanced Topics**
1. What is a database view, and how would you use one?
2. Explain the difference between SQL’s `HAVING` and `WHERE` clauses with examples.
3. Write a query to fetch all records where a specific column has a JSON value containing a specific key or value.
4. What are window functions, and how would you use one to calculate a running total?
5. Write a query using a window function to rank employees based on their salaries within their department.
6. How would you store hierarchical data in a relational database?
7. Write a query to pivot rows into columns in SQL.
8. What is a materialized view, and how is it different from a regular view?
9. How would you use the `JSONB` data type in PostgreSQL for storing semi-structured data?
10. What is sharding, and when would you use it?

---

## **Database Design**
1. What are the differences between normalization and denormalization? When would you use each?
2. How would you design a database schema to handle users, roles, and permissions?
3. What is a foreign key constraint, and what are the different ON DELETE behaviors?
4. How would you enforce uniqueness across multiple columns in a table?
5. What are some strategies for designing a schema to handle multi-tenancy?
6. How would you design a database to handle soft deletes?
7. Explain the purpose of surrogate keys versus natural keys in database design.
8. How would you implement a many-to-many relationship in a database?
9. What are the advantages and disadvantages of using stored procedures?
10. How would you design a schema to support versioning of records in a table?

---

## **Common Real-World Scenarios**
1. How would you handle a situation where your table is growing rapidly in size and causing performance issues?
2. What strategies would you use to safely migrate a large database table without downtime?
3. Write a query to paginate results from a table, showing 10 results per page.
4. Explain how you would design a database to handle soft deletes.
5. How would you manage schema changes in a production database?
6. What steps would you take to troubleshoot slow queries in a production database?
7. Write a query to merge two tables with duplicate data.
8. How would you handle data integrity issues caused by an application bug?
9. What techniques would you use to archive old data from a large table?
10. How would you monitor and ensure database health over time?

---

## **Debugging and Maintenance**
1. What is the purpose of the `EXPLAIN` or `EXPLAIN PLAN` statement in SQL?
2. How do you identify and fix deadlocks in a database?
3. How would you detect missing indexes or unused indexes in a database?
4. Write a query to identify the most resource-intensive queries running on a database.
5. How would you monitor database performance in real-time?
6. What tools would you use for database profiling?
7. How would you handle log growth for a transactional database?
8. How do you ensure database backups are reliable and restorable?
9. What is the difference between database repair and optimization?
10. How would you test changes to a database schema in a development environment?

---
