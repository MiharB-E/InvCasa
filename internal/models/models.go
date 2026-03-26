package models

type User struct {
    ID      int64  `json:"id"`
    Email   string `json:"email"`
    Password string `json:"-"`
    Name    string `json:"name"`
    GroupID int64  `json:"group_id"`
}

type Group struct {
    ID         int64  `json:"id"`
    Name       string `json:"name"`
    InviteCode string `json:"invite_code"`
}

type Category struct {
    ID        int64  `json:"id"`
    Name      string `json:"name"`
    IsDefault bool   `json:"is_default"`
}

type Product struct {
    ID         int64   `json:"id"`
    Name       string  `json:"name"`
    CategoryID *int64  `json:"category_id"`
    Quantity   float64 `json:"quantity"`
    Unit       string  `json:"unit"`
    Status     string  `json:"status"`
    GroupID    int64   `json:"group_id"`
    IsFavorite bool    `json:"is_favorite"`
}

type Purchase struct {
    ID        int64   `json:"id"`
    ProductID int64   `json:"product_id"`
    UserID    int64   `json:"user_id"`
    Quantity  float64 `json:"quantity"`
    Price     float64 `json:"price"`
    StoreName string  `json:"store_name"`
}

type ShoppingListItem struct {
    ID        int64  `json:"id"`
    ProductID int64  `json:"product_id"`
    Status    string `json:"status"`
}