
## Jwt Token Generation APIs

### `GET` `/admin/token` Genarate token for admin

Input

```go
    type AdminUser struct {
        gorm.Model
        Name     string `json:"name"`
        Username string `json:"username" gorm:"unique"`
        Email    string `json:"email" gorm:"unique"`
        Password string `json:"password"`
    }
```

Sample Response

```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFydW4iLCJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImV4cCI6MTY2MDczNDUxOH0.JU-81FMVjivS9_j3_NZhg35kgYnkDYI7x2YDpb24MZA"
    }
```

### `GET` `/user/token` Genarate token for users

Input

```go
    type User struct {
        gorm.Model
        Name     string `json:"name"`
        Username string `json:"username" gorm:"unique"`
        Email    string `json:"email" gorm:"unique"`
        Password string `json:"password"`
    }
```

Sample Response

```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFydW4iLCJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImV4cCI6MTY2MDczNDUxOH0.JU-81FMVjivS9_j3_NZhg35kgYnkDYI7x2YDpb24MZA"
    }
```