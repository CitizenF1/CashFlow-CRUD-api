# CashFlow-CRUD-api

### Routes

#### Payments

- [GET] /payments

  - Return all current payments

- [POST] /payments

  - Add a new payment to database
  - Expected data:

  ```json{
    "data": {
        "_id": "6210f98db2450971ec98aca1",
        "category": "example",
        "name": "Example",
        "price": 1200
    },
    "status": 201
}
  
- [DELETE] /:id

  - delete payment by id
