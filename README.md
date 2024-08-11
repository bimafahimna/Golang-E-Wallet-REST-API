# Golang-E-Wallet-REST-API
Robust e-wallet REST API using Go's Gin framework. 🔥

## Project Description
Developed with **Clean Architecture** in mind and provided with **[Postman Collection](https://documenter.getpostman.com/view/31342872/2sA3e5c7aB)** as documentation.

Featuring **JWT authentication**, **robust error handling**, and **logging with logrus**. Enables seamless fund transfers and top-ups for a great user experience.




## How to run

``git clone https://github.com/bimafahimna/Golang-E-Wallet-REST-API.git``

1. Connect with the selected postgresql database
1. Run wallet_number_seq.sql
2. Run starter.sql
3. Run dataseeding.sql
4. Run main.go 

PS: can use the Makefile

## Database
### Seed Data
- [x] 5 Users
- [x] 20 transactions
- [x] additional data

## Features
- [x] REST API standard design
 - [x] logger
 - [x] request cancelation
 - [x] graceful shutdown

### Authentication and Authorizatoin
- [x] Register endpoint
- [x] Login endpoint
- [x] Forget Password endpoint
- [x] Reset Password endpoint

### User
- [x] user detail

### List of Transactions
- [x] Get Transaction List endpoint

### Transfer
- [x] Transfer endpoint
   - [x] The target wallet must be invalid or does not exist
   - [x] User can't transfer to their own wallet
   - [x] User can't transfer when balance is insufficient
   - [x] User can't transfer from a wallet that doesn't belong to them
   - [x] description max 35 chars

### Top Up
- [x] Top-Up endpoint
 - [x] fields from client is as requirements
 - [x] minimum amount is 50,000.0000 maximum amount is 10,000,000.0000
 - [x] source of funds must be from listed options
 - [x] transaction description "Top Up from (sourceOfFunds)"
 - [x] user can only top-up to his/her own wallet
 - [x] add game attempt after every top up with amount multiple of 10,000,000.000 (not accumulation)

