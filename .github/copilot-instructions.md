## Layout of the project
- /database contains all the DAO files
  - /initializers contains all the database initialization files
  - /managers contains all the database manager files that actually manages the active instance of each DAO. Service files should call the manager files to retrieve data from each DAO.
- /rest contains all the rest files. This is where we actually serve REST endpoints, validate input, and call the appropriate service classes to perform work. In this folder, we should handle all HTTP and status code reading/setting. When we call out to other methods in different directories, we shouldn't ever pass any HTTP request or expect a response in return. The only place we should be importing net/http should be in this directory
- /service contains functions that perform meaningful work. Calling out to external services, orchestrating data manipulation and filtering, etc.
- /vo contains value objects. These are the files that define all the structs, including response bodies, data definitions, and in-memory cached data definitions.

## Project guidelines
- Always follow SOLID principles. All methods and functions should do one singular thing. Examples of method names:
  - buildResponseBundle
  - filterOutDuplicates
  - isValid
- Unit tests should always be created and updated when adding or modifying non-trivial code. Any code that is touched or created should be unit tested. If we follow SOLID principles, unit tests should be short and sweet, requiring only 2 or 3 tests each. If you need more tests than that, in most cases the code should be split up into more logical methods
- Simple is always better. Always stick to simple code flows and easy-to-read methods. Don't use hard to read ternary operators or unnecessarily long method chaining. Maintainability is key

## Agent Guidelines
- Keep it simple stupid
- Don't overengineer
- Value quick actions rather than long complex project updating
- Perform exactly what I ask and restrict yourself to the structures that I have defined