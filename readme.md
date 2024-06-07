# E-commerce REST API

This README provides an overview of the available endpoints for the E-commerce REST API. Each endpoint is designed to handle specific operations related to sellers, buyers, products, and orders.

## Table of Contents

- [Seller Endpoints](#seller-endpoints)
  - [Login Seller](#login-seller)
  - [Register Seller](#register-seller)
- [Buyer Endpoints](#buyer-endpoints)
  - [Login Buyer](#login-buyer)
  - [Register Buyer](#register-buyer)
- [Product Endpoints](#product-endpoints)
  - [Get All Products](#get-all-products)
  - [Get Products by Seller ID](#get-products-by-seller-id)
  - [Add Product](#add-product)
- [Order Endpoints](#order-endpoints)
  - [Get All Orders](#get-all-orders)
  - [Create Order](#create-order)
  - [Accept Order](#accept-order)
  - [Get List of Orders](#get-list-of-orders)

## Seller Endpoints

### Login Seller
- **Endpoint:** `POST /seller/login`
- **Handler:** `loginSeller`
- **Description:** Authenticates a seller and returns a token.

### Register Seller
- **Endpoint:** `POST /seller/register`
- **Handler:** `registerSeller`
- **Description:** Registers a new seller account.

## Buyer Endpoints

### Login Buyer
- **Endpoint:** `POST /buyer/login`
- **Handler:** `loginBuyer`
- **Description:** Authenticates a buyer and returns a token.

### Register Buyer
- **Endpoint:** `POST /buyer/register`
- **Handler:** `registerBuyer`
- **Description:** Registers a new buyer account.

## Product Endpoints

### Get All Products
- **Endpoint:** `GET /product/list`
- **Handler:** `getAll`
- **Description:** Retrieves a list of all products.

### Get Products by Seller ID
- **Endpoint:** `GET /product`
- **Handler:** `productBySellerID`
- **Description:** Retrieves products by a specific seller's ID.

### Add Product
- **Endpoint:** `POST /product`
- **Handler:** `addProduct`
- **Description:** Adds a new product to the catalog.

## Order Endpoints

### Get All Orders
- **Endpoint:** `GET /order/list`
- **Handler:** `getAll`
- **Description:** Retrieves a list of all orders.

### Create Order
- **Endpoint:** `POST /order`
- **Handler:** `createOrder`
- **Description:** Creates a new order.

### Accept Order
- **Endpoint:** `PUT /orders/:id`
- **Handler:** `acceptOrder`
- **Description:** Accepts an order by its ID.

### Get List of Orders
- **Endpoint:** `GET /orders/list`
- **Handler:** `getListOrder`
- **Description:** Retrieves a list of orders for the authenticated user.

---

Each endpoint is mapped to a specific handler function within the API's codebase. The handlers are responsible for processing the requests, interacting with the underlying services, and returning appropriate responses. For more detailed information on each endpoint's functionality, refer to the source code linked within the handler descriptions.
