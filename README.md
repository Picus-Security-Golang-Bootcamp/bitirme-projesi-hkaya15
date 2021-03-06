# Picus Security Golang Backend Bootcamp | E-Commerce App
This work is the final assignment of the [Picus Security](https://www.picussecurity.com) Golang Backend Web Development Bootcamp. It includes the basics of the e-commerce app backend.

## Installation
```bash
git clone https://github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-hkaya15.git
```

## Features

### SignUp
    Users can create an account with this feature. 
* Verify E-mail
* Hash Password
* Generate Token (Access - 15 min. & Refresh - 7 day)
* Set Cookie (It is used for authentication)

### Login
    Users can enter on app with this feature
* Password Hash Decode
* Decode Cookie (It is used for authentication)
* Verify Access Token
* Verify Refresh Token
* Generate Token (If expired)
* Set Cookie

### Category
    Users can upload categories with csv file or get categories list

1. Upload Category File
    * Authorization Middleware (User can upload category file if user role is admin)
    * Compare DB records & File (No duplicate records)

2. Get All Categories
    * Pagination Middleware

### Product
    Users can create, search, update, delete product or get all product list

1. Create 
    * Authorization Middleware (User can create product if user role is admin)
2. Search
    * Search by product name, description or id without case-sensitivity
3. Update
    * Authorization Middleware (User can update product if user role is admin)
4. Delete
    * Authorization Middleware (User can delete product if user role is admin)
5. Get All Products
    * Pagination Middleware    

### Cart
    Users can add product on own cart, update cart, delete product from cart or get list of cart

1. Add
    * Authentication Middleware
2. Update
    * Authentication Middleware
    * Acting as a delete (If user update product number with zero, it is working like delete)
3. Delete
    * Authentication Middleware
4. Get Cart
    * Authentication Middleware

### Order
    Users can give an order, cancel order or get all order list

1. Complete Order
    * Authentication Middleware
2. Cancel
    * Authentication Middleware
3. Get All Order
    * Authentication Middleware

## Health Check
    There is an basic health check function that controls db health with 10 seconds break
## Packages
    Swagger-go, Viper, Zap

## Credits

You can check the [Structure of API](https://app.swaggerhub.com/apis/HKaya15/e-commerce_app/1.0.0)

## Structure of the App
- cmd
    - main.go
- pkg
    - api
        - model
    - app
        - cart
            - handler
            - model
            - repository
            - service
        - category
            - handler
            - model
            - repository
            - service
        - order
            - handler
            - model
            - repository
            - service
        - product
            - handler
            - model
            - repository
            - service
        - status
            - handler
        - user
            - handler
            - model
            - repository
            - service
    - base
        - config
        - db
        - errors
        - graceful
        - helper
        - jwt
        - log
        - middleware
        - pagination
    - docs

## License
[MIT](https://mit-license.org)