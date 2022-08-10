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
 * Kindly use the following emails already in the db **raymond@test.com** (has admin rights )or **gitonga@test.com** (no admin rights)
 <br><br>

2. Get Companies - Get a list of all companies. **GET**
``/companies``
  <br>
  This endpoint accepts the following optional parameters you can use for filtering
    * phone
    * id
    * name
    * code
    * country
    * website
    * phone
 <br><br>

3. Get Company - Get a single company. **GET** ``/company``
   <br>
   This endpoint accepts the following optional parameters you can use for filtering
    * phone
    * id
    * name
    * code
    * country
    * website
    * phone
    <br><br>

4. Create Company - Create a new company.  **POST** ``/company/create``
   <br>
   * This endpoint expects mandatory **Header**  ``Authorization : {JWT_TOKEN}`` and for testing purposes ``X-REAL-IP: {IP_ADDRESS}``
   * This endpoint expects the following payload
   ```
   {
    "name": "tosh",
    "code":"124",
    "country":"CY",
    "website":"www.xxx",
    "phone":"354993"
   }
   ```
   <br><br>

5. Delete Company - Create a new company.  **DELETE** ``/company/delete?id={ID}``
   <br>
    * This endpoint expects mandatory **Header**  ``Authorization : {JWT_TOKEN}`` and for testing purposes ``X-REAL-IP: {IP_ADDRESS}``
   <br><br>

6. Update Company - Create a new company.  **PUT** ``/company/update?id={ID}``
   <br>
   * This endpoint expects mandatory **Header**  ``Authorization : {JWT_TOKEN}`` and for testing purposes ``X-REAL-IP: {IP_ADDRESS}``
   <br><br>
   * This endpoint expects the following payload. All of the fields are optional
   ``
     {
     "name": "tosh 1",
     "code":"124",
     "country":"CY",
     "website":"www.xxx",
     "phone":"729320243"
     }
   ``




  
