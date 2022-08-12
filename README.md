# Company Service

## Requirements
1. Docker
2. Docker-compose

## How to run the service

Run the following to setup the database, kafka instance and to run the service

``make build``

### Endpoints

1. Get AuthToken - Used to get a jwt token for mutating operations . **GET** ``/authorize?email={email}``
<br>
2. Get Companies - Get a list of all companies. **GET** ``/companies``
<br>
3. Get Company - Get a single company. **GET** ``/company``
<br>
4. Create Company - Create a new company.  **POST** ``/company/create``  
<br>
5. Delete Company - Create a new company.  **DELETE** ``/company/delete?id={ID}``
<br>
6. Update Company - Create a new company.  **PUT** ``/company/update?id={ID}``




  
