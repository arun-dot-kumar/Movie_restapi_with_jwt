# Simulation Of Movie Ticket System Using JWT Autherization

## List of features

### Admin Service
* Add Admin User (Admin Who Maintaines Addition And Booking Of Movie Shows)
* Generate Jwt Token For Admin (Geneartes Jwt Token For Admin User)
* Add Movies (any movie that has been released)
* Show Booking List(displays the list of bookings till now)

### User Service
* Create New User (Creat User Accounts for Ticker Booking)
* Generate Jwt Token For User
* Get Show List (list of shows available for booking)
* Book Ticket With Multiple Seat Selection (Book Multiple Seats In A Screen Show)

# Running With local mysql database:

## Prerequisets
- [ ] Mysql Database to be installed with the proper username and password
- [ ] Create Database Named Movie


## for build and runniing application.
```bash
go build -o movie .
./movie -instance
```

# Running Locally using Docker

```bash
docker-compose up
```
# Documentation

* [API Docs](docs/api-docs/index.md)


access site on: http://localhost:4567/

# TODO

- [ ] Add testscases for all api's
- [ ] Fix MYSQL setup in github actions
- [ ] Add logger to each struct in services / api handlers
- [ ] Add  Theater list by cities and date
