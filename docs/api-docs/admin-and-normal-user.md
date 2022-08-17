## Admin and users registration APIs

### `Post` `/admin/register`  admin registration

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
      "email": "testuser@gmail.com",
      "userId": 1,
      "username": "arun"
    }
```

### `POST` `/user/register` user registration

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
      "email": "testuser@gmail.com",
      "userId": 1,
      "username": "arun"
    }
```