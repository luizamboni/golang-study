hexagon architecture POC
===
A little POC of hexagonal architecture how i understood. 


- **core** package contains domain logic, and not depends of anyone.
- **adapters** depends on **core** package.
- **http server** is a adapter that lead income http requests to domain logic.
- **logger** is inside **core** because everyone use it.