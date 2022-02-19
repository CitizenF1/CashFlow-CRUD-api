# CashFlow-CRUD-api

### Routes

#### Payments

- [GET] /payments

  - Return all current payments

- [POST] /payments

  - Add a new payment to database
  - Expected data:

  ```json
  {
    "name": "Example",
    "price": 1200,
    "category": "example"
  }
