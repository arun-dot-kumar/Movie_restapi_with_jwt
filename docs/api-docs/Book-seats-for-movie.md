## Booking APIs

### `GET` `/movie/user/bookmovie` book a seats in preferred movie show

Input

```json
    {
        "userName": "arun",
        "email": "testuser@gmail.com",
        "showName": "charlie",
        "numOfSeats": 2,
        "seats": "A1,B1",
        "showTime": "10:00 AM"
    }
```

Sample Response

```json
    {
      "status": "2 tickets A1,B1 for the movie charlie at 10:00 AM has successfully booked, Total cost of Tickets are 600 and booking id is 1"
    }
```
