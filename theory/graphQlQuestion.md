# Top 50 GraphQL Interview Questions

## 1. What is GraphQL?
**Answer**: GraphQL is a query language for APIs and a runtime for executing those queries with the data. It enables clients to request exactly the data they need, making it efficient and flexible.

## 2. How does GraphQL differ from REST?
**Answer**: Unlike REST, where the server defines the structure of responses, GraphQL allows clients to specify the structure of responses. This means clients can request only the data they need, reducing over-fetching and under-fetching of data.

## 3. What is a GraphQL schema?
**Answer**: A GraphQL schema defines the types, queries, and mutations in your GraphQL API. It acts as a contract between the client and server, specifying what data can be queried or mutated.

## 4. What are types in GraphQL?
**Answer**: Types in GraphQL define the shape of the data in the schema. Common types include `Query`, `Mutation`, `ObjectType`, `Enum`, `InputType`, and `Scalar`.

## 5. What are queries in GraphQL?
**Answer**: A query in GraphQL is a request from the client to the server for data. It specifies the fields that should be returned in the response.

## 6. What are mutations in GraphQL?
**Answer**: Mutations are used to modify data on the server, such as creating, updating, or deleting records. They are similar to HTTP POST, PUT, or DELETE requests in REST.

## 7. What is a resolver in GraphQL?
**Answer**: A resolver is a function responsible for fetching data for a specific field in a query. It’s how GraphQL handles requests for data and maps them to underlying data sources.

## 8. What is a GraphQL subscription?
**Answer**: A subscription in GraphQL is a way to listen for real-time updates. It allows the client to receive notifications when data changes, such as through WebSockets.

## 9. What is the difference between query and mutation in GraphQL?
**Answer**: A `Query` fetches data from the server, while a `Mutation` modifies data on the server. Both are defined in the GraphQL schema.

## 10. What are scalar types in GraphQL?
**Answer**: Scalar types represent the leaves of the query, and they include types like `String`, `Int`, `Float`, `Boolean`, and `ID`.

## 11. What are object types in GraphQL?
**Answer**: Object types define a set of fields that represent a complex object in your schema. For example, a `User` type might have fields like `name`, `email`, and `age`.

## 12. What are input types in GraphQL?
**Answer**: Input types define the structure of data that can be passed into mutations. They are used for complex arguments to mutations or queries.

## 13. What are enums in GraphQL?
**Answer**: Enums are a special kind of scalar that represents a set of possible values. For example, you can define an `Enum` for `Status` with values `ACTIVE`, `INACTIVE`, `PENDING`.

## 14. What are interfaces in GraphQL?
**Answer**: An interface is an abstract type that defines a set of fields that must be implemented by object types. It is similar to an interface in object-oriented programming.

## 15. What are unions in GraphQL?
**Answer**: A union type allows a field to return one of several different object types. It’s useful when a field can return more than one type of object.

## 16. How do you paginate in GraphQL?
**Answer**: Pagination in GraphQL can be done using techniques like `limit-offset` or `cursor-based` pagination. Pagination is typically done in queries using `first`, `after`, `last`, and `before` arguments.

## 17. What are fragments in GraphQL?
**Answer**: Fragments are reusable pieces of a query. They allow you to define parts of queries that can be reused in different parts of your query.

## 18. What are directives in GraphQL?
**Answer**: Directives provide a way to modify the execution of a query. For example, the `@skip` directive can be used to conditionally exclude fields from a query.

## 19. What are variables in GraphQL?
**Answer**: Variables allow you to pass dynamic values to your queries. This helps prevent hardcoding values and makes your queries reusable.

## 20. How does error handling work in GraphQL?
**Answer**: Errors in GraphQL are returned in the `errors` field of the response, and the `data` field contains the successful result (if any). Errors can occur during query resolution or due to validation.

## 21. What is batching in GraphQL?
**Answer**: Batching is a technique where multiple GraphQL requests are grouped into a single HTTP request to reduce the number of network round-trips.

## 22. What is dataloader in GraphQL?
**Answer**: Dataloader is a utility for batching and caching database queries in GraphQL. It helps to prevent the N+1 query problem by batching requests and caching results.

## 23. How do you implement authentication in GraphQL?
**Answer**: Authentication in GraphQL is typically handled through HTTP headers, where a token (like JWT) is passed and verified in the resolver layer to authenticate the user.

## 24. What is the difference between GraphQL and RESTful APIs?
**Answer**: In REST, the server defines the structure of responses, whereas in GraphQL, the client specifies the structure of the response. GraphQL also allows for more granular queries and reduces over-fetching and under-fetching of data.

## 25. How do you handle authorization in GraphQL?
**Answer**: Authorization in GraphQL is typically handled by checking the user’s permissions within resolvers, often based on roles or access control lists (ACLs).

## 26. What are the common security concerns with GraphQL?
**Answer**: Common concerns include unauthorized access to sensitive data, denial of service through expensive queries, and exposing too much data via introspection.

## 27. What is introspection in GraphQL?
**Answer**: Introspection is a feature of GraphQL that allows clients to query the schema for details about the types, queries, mutations, and subscriptions available in the API.

## 28. How do you optimize GraphQL queries?
**Answer**: Optimization can be achieved by using techniques like batching, caching, limiting query depth, and avoiding overly complex queries that return large datasets.

## 29. What are the advantages of using GraphQL?
**Answer**: Advantages of GraphQL include flexible queries, reduced over-fetching, real-time updates with subscriptions, and the ability for clients to request exactly the data they need.

## 30. What is Apollo Server?
**Answer**: Apollo Server is a popular implementation of a GraphQL server that allows you to build a GraphQL API quickly and easily with features like schema stitching, caching, and federation.

## 31. What is Apollo Client?
**Answer**: Apollo Client is a library that helps you interact with GraphQL APIs from the client side. It provides features like caching, pagination, and optimistic UI updates.

## 32. What is schema stitching in GraphQL?
**Answer**: Schema stitching is the process of combining multiple GraphQL schemas into a single unified schema. It allows you to split your GraphQL API into smaller, more manageable parts.

## 33. What is GraphQL Federation?
**Answer**: Federation is a technique for composing multiple GraphQL services into one unified API. It is often used with Apollo Server for building microservices.

## 34. What is the GraphQL `@defer` directive?
**Answer**: The `@defer` directive allows a GraphQL query to return parts of the response immediately, while other parts can be deferred and sent later. This is useful for reducing latency in complex queries.

## 35. What is the `@include` directive in GraphQL?
**Answer**: The `@include` directive allows fields to be conditionally included in the query. For example, a field can be included only if a variable is true.

## 36. How do you handle caching in GraphQL?
**Answer**: Caching in GraphQL can be handled using techniques like client-side caching (with Apollo Client) or server-side caching, such as using DataLoader, HTTP caching, or persisted queries.

## 37. What are persisted queries in GraphQL?
**Answer**: Persisted queries are pre-registered queries that are sent to the server as identifiers rather than the full query text, which can reduce the size of the request and improve performance.

## 38. How do you version a GraphQL API?
**Answer**: GraphQL APIs are usually versionless, but schema changes can be managed by deprecating old fields or introducing new fields in the schema.

## 39. What are some common use cases for GraphQL?
**Answer**: Common use cases for GraphQL include web and mobile applications, real-time data feeds, and microservices.

## 40. How does GraphQL handle batch processing?
**Answer**: Batch processing in GraphQL can be done by grouping multiple requests together in one API call, or by using data loaders to batch database queries.

## 41. How do you monitor a GraphQL server?
**Answer**: Monitoring a GraphQL server can be done through logging, performance tracking (e.g., using Apollo Engine), and error tracking tools like Sentry or New Relic.

## 42. What is the role of `graphql-tag`?
**Answer**: `graphql-tag` is a JavaScript library used to parse GraphQL queries written as tagged template literals. It helps in writing and parsing GraphQL queries in JavaScript code.

## 43. How do you test GraphQL queries?
**Answer**: GraphQL queries can be tested using tools like Postman, GraphiQL, or Apollo Client’s testing utilities. You can also write unit tests for resolvers and schema validation.

## 44. What is the difference between GraphQL and WebSockets?
**Answer**: WebSockets are a communication protocol that allows for real-time two-way communication, while GraphQL is a query language for APIs. GraphQL subscriptions are typically implemented over WebSockets.

## 45. What is a GraphQL Playground?
**Answer**: GraphQL Playground is an interactive IDE for exploring GraphQL APIs. It provides an interface for writing queries, inspecting schemas, and testing API endpoints.

## 46. What is `graphql-js`?
**Answer**: `graphql-js` is the official JavaScript implementation of GraphQL. It is used for building a GraphQL server in Node.js and provides all the tools needed to define schemas, resolvers, and handle requests.

## 47. How do you implement caching in GraphQL?
**Answer**: Caching in GraphQL can be done at the query level (using Apollo Client’s cache) or at the response level (using HTTP caching, persistent queries, or a reverse proxy).

## 48. What are query complexity and depth analysis in GraphQL?
**Answer**: Query complexity and depth analysis are used to calculate the potential cost of a query based on its structure, to prevent abuse of the API by overly complex queries.

## 49. How do you handle file uploads in GraphQL?
**Answer**: File uploads in GraphQL can be handled using custom scalars or specialized libraries like `graphql-upload`, which supports multipart file uploads.

## 50. What is a schema-first approach in GraphQL?
**Answer**: A schema-first approach is when you define the schema first and then generate the resolvers based on the schema. This approach encourages better design and collaboration.


# Advance Question

# Top 50 GraphQL Interview Questions for 2 Years of Experience

## 1. What is GraphQL and how does it work?
**Answer**: GraphQL is a query language for APIs and a runtime for executing those queries with data. It allows clients to request exactly the data they need, minimizing over-fetching and under-fetching.

## 2. What are the key differences between GraphQL and REST APIs?
**Answer**: GraphQL gives clients the ability to request exactly what they need, whereas REST APIs typically return predefined responses. GraphQL also allows for real-time updates through subscriptions.

## 3. Can you explain the concept of a schema in GraphQL?
**Answer**: A GraphQL schema defines the types of data available in an API. It consists of object types, queries, mutations, and subscriptions that define the shape and structure of the data clients can interact with.

## 4. What are queries in GraphQL?
**Answer**: A query is a request for data from the GraphQL server. It allows clients to request specific fields and avoid over-fetching unnecessary data.

## 5. What are mutations in GraphQL?
**Answer**: Mutations are used for modifying data on the server, such as creating, updating, or deleting records. They are similar to POST, PUT, or DELETE requests in REST.

## 6. What are resolvers in GraphQL?
**Answer**: Resolvers are functions that handle fetching the data for specific fields in a query. They resolve the data to be returned to the client.

## 7. How do you handle errors in GraphQL?
**Answer**: Errors are returned in the `errors` field in the response, and successful results are returned in the `data` field. You can handle validation or execution errors in resolvers.

## 8. What are the different types of GraphQL types?
**Answer**: The common GraphQL types are `Object Types`, `Scalars` (e.g., `String`, `Int`), `Enums`, `Input Types`, and `Interfaces`.

## 9. What are scalar types in GraphQL?
**Answer**: Scalar types are the most basic types in GraphQL and represent primitive values. These include `String`, `Int`, `Float`, `Boolean`, and `ID`.

## 10. What are input types in GraphQL?
**Answer**: Input types define the structure of input data passed into mutations or queries. They allow complex arguments to be passed to fields or mutations.

## 11. How would you use GraphQL subscriptions?
**Answer**: Subscriptions allow clients to listen to real-time updates from the server. They are often implemented using WebSockets and are useful for applications like chat systems or live feeds.

## 12. Can you explain the difference between a query and a mutation in GraphQL?
**Answer**: A `Query` fetches data from the server, while a `Mutation` modifies data on the server (e.g., adding or updating records).

## 13. What are fragments in GraphQL?
**Answer**: Fragments are reusable pieces of query logic that allow you to avoid repeating fields in multiple parts of a query.

## 14. How do you handle authentication in GraphQL?
**Answer**: Authentication in GraphQL is handled by passing an authorization token (e.g., JWT) in HTTP headers. This token is then validated in resolvers.

## 15. What are variables in GraphQL and how do you use them?
**Answer**: Variables are used to pass dynamic values to GraphQL queries or mutations. They prevent hardcoding values and allow queries to be more reusable.

## 16. How do you paginate data in GraphQL?
**Answer**: Pagination is typically done using cursor-based pagination, where the query includes `first`, `after`, `last`, and `before` arguments to fetch chunks of data.

## 17. How would you implement caching in GraphQL?
**Answer**: Caching in GraphQL can be implemented with client-side libraries like Apollo Client, which caches data after queries, or by using server-side caching for query results.

## 18. How do you handle file uploads in GraphQL?
**Answer**: File uploads in GraphQL can be handled using a specialized library like `graphql-upload`, which supports multipart file uploads.

## 19. What is the `@defer` directive in GraphQL?
**Answer**: The `@defer` directive allows the server to defer parts of a query’s response, returning the initial data while other parts can be returned later.

## 20. How do you manage different environments in GraphQL (e.g., dev, staging, production)?
**Answer**: Environment management can be handled by using environment variables or configuration files to configure the GraphQL endpoint and other settings specific to each environment.

## 21. What is the purpose of the `@include` directive in GraphQL?
**Answer**: The `@include` directive allows a field to be conditionally included in a query based on a boolean argument.

## 22. What are some common security issues with GraphQL APIs?
**Answer**: Common security issues include exposing sensitive data, denial of service through complex queries, and issues with query introspection. It’s important to implement authentication, authorization, and query complexity analysis.

## 23. What is Apollo Server?
**Answer**: Apollo Server is a popular open-source GraphQL server implementation that supports various features like schema stitching, subscriptions, and middleware integrations.

## 24. What is Apollo Client and how does it work?
**Answer**: Apollo Client is a library used to interact with GraphQL APIs from the client side. It provides features like caching, pagination, and state management.

## 25. How does Apollo Client handle caching?
**Answer**: Apollo Client caches query results and allows you to configure how caching should be done, ensuring that repeated queries can be served from the cache, improving performance.

## 26. How do you perform query optimization in GraphQL?
**Answer**: Query optimization can be achieved by limiting query depth, using batching to group multiple requests, and limiting the amount of data returned by queries.

## 27. How do you handle versioning in GraphQL?
**Answer**: GraphQL APIs are often versionless, with changes handled through schema evolution. Deprecating old fields and introducing new fields avoids the need for versioning.

## 28. What is schema stitching in GraphQL?
**Answer**: Schema stitching involves combining multiple GraphQL schemas into a single unified schema, allowing for modular GraphQL APIs across microservices.

## 29. What is a GraphQL Federation?
**Answer**: GraphQL Federation allows you to build a distributed architecture for GraphQL by composing multiple independent GraphQL services into a single schema, often used in microservices environments.

## 30. How do you implement authorization in GraphQL?
**Answer**: Authorization can be implemented in GraphQL by checking user roles or permissions within resolvers or middleware, ensuring that users can only access authorized data.

## 31. What are query complexity and depth analysis in GraphQL?
**Answer**: Query complexity and depth analysis are used to calculate the cost of a query based on its structure, preventing clients from making expensive or deep queries that could overload the server.

## 32. How do you test GraphQL queries?
**Answer**: GraphQL queries can be tested using tools like GraphiQL, Apollo Client’s testing utilities, or by writing unit tests for the resolvers and schemas.

## 33. What are persisted queries in GraphQL?
**Answer**: Persisted queries allow clients to send only the query identifier, rather than the full query string, improving efficiency and security by reducing the size of requests.

## 34. How do you handle subscription messages in GraphQL?
**Answer**: Subscription messages in GraphQL are handled over WebSockets, where clients can subscribe to certain events and receive real-time updates from the server.

## 35. What are the common use cases for GraphQL?
**Answer**: Common use cases for GraphQL include real-time data feeds (e.g., live chat or stock updates), web and mobile applications, and handling complex data relationships between microservices.

## 36. What are fragments in GraphQL used for?
**Answer**: Fragments are reusable parts of GraphQL queries that allow you to extract common fields from multiple queries or mutations, reducing redundancy.

## 37. What is GraphQL introspection?
**Answer**: Introspection allows clients to query a GraphQL server for its schema. This helps clients discover available types, queries, and mutations.

## 38. How do you prevent N+1 query problems in GraphQL?
**Answer**: The N+1 query problem can be prevented by using data loader libraries to batch and cache database queries in GraphQL resolvers.

## 39. What are the main advantages of using GraphQL over REST?
**Answer**: GraphQL allows clients to specify exactly the data they need, reducing over-fetching and under-fetching, and it provides more flexible querying, especially for complex data relationships.

## 40. How do you handle timeouts and retries in GraphQL?
**Answer**: Timeouts and retries in GraphQL can be handled using middleware or client libraries like Apollo Client, which support retry policies for failed requests.

## 41. How do you deal with large responses in GraphQL?
**Answer**: Large responses can be handled through pagination, query limits, or by breaking down queries into smaller chunks that return manageable amounts of data.

## 42. What is `graphql-tag` used for?
**Answer**: `graphql-tag` is a JavaScript library that parses GraphQL queries written as tagged template literals into executable query objects, making it easier to use GraphQL in JavaScript.

## 43. What are the differences between GraphQL and gRPC?
**Answer**: GraphQL is a query language for APIs, while gRPC is a high-performance RPC framework. GraphQL uses HTTP and JSON, while gRPC uses HTTP/2 and Protocol Buffers for serialization.

## 44. How do you implement search functionality in GraphQL?
**Answer**: Search functionality can be implemented using a query that accepts a search string and returns matching results from the server, potentially using full-text search libraries or database indexing.

## 45. How do you manage large GraphQL APIs?
**Answer**: Large GraphQL APIs can be managed through modular schemas, schema stitching, or Federation, breaking the API into smaller, more maintainable parts.

## 46. What are the most commonly used GraphQL tools and libraries?
**Answer**: Common tools and libraries include Apollo Server, Apollo Client, GraphQL.js, Prisma, GraphiQL, and GraphQL Playground.

## 47. How would you troubleshoot a slow GraphQL query?
**Answer**: Troubleshooting a slow query involves checking the query’s complexity, using query profiling tools, optimizing resolvers, reducing the size of returned data, and checking for database or network issues.

## 48. What is the role of middleware in a GraphQL server?
**Answer**: Middleware in GraphQL is used for various purposes, including authentication, logging, error handling, and adding additional headers or context to requests.

## 49. How does a GraphQL server handle caching?
**Answer**: GraphQL servers handle caching by utilizing caching mechanisms at the server, resolver, or query level. Common libraries like Apollo Server provide built-in caching features.

## 50. How do you handle rate limiting in GraphQL?
**Answer**: Rate limiting in GraphQL can be handled using middleware that tracks the number of requests and restricts access after a certain limit, often integrated with a Redis-based store.
