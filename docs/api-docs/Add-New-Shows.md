## New shows APIs

### `POST` `/movie/admin/addShow` Add new show

Input

```go
    type Show struct {
        gorm.Model
        ShowName        string         `json:"showName"`
        GoldClassSeats  datatypes.JSON `json:"goldClassSeats"`
        BalconySeats    datatypes.JSON `json:"balconySeats"`
        FirstClassSeats datatypes.JSON `json:"firstClassSeats"`
        NumberOfSeats   int            `json:"numberOfSeats"`
        ShowTime        string         `json:"showTime"`
    }
```

Sample Response

```json
    {
        "ShowId": 0,
        "ShowName": "charlie",
        "ShowTime": [
            "10:00 AM",
            "2:00 PM",
            "5:00 PM",
            "9:00 PM"
        ]
    }
```
