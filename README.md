# 1. USER MANAGEMENT

## 1.1. Create a new user

POST /api/users
permission: 0 - user, 1 - admin

```json
{
  "username": "anhtuyet",
  "password": "123456",
  "fullname": "Nguyễn Ánh Tuyết"
}
```

## 1.2. Get all users

GET /api/users?page=&size=

```json
{
  "paging": {
    "total_page": 2,
    "total_item": 3,
    "page_size": 2,
    "current_page": 1
  },
  "users": [
    {
      "user_id": 2,
      "username": "quanghai",
      "permission": 0,
      "provider_id": null,
      "created_at": "2025-01-08T23:42:01.492779",
      "login_provider": "Local",
      "fullname": "Nguyễn Quang Hải",
      "avatar": "https://th.bing.com/th/id/OIP.P8F796BGNue4Lu2SImT1bgAAAA$1rs=1&pid=ImgDetMain",
      "phone": null,
      "address": null
    },
    {
      "user_id": 3,
      "username": "mailinh",
      "permission": 0,
      "provider_id": null,
      "created_at": "2025-01-08T23:43:29.443218",
      "login_provider": "Local",
      "fullname": "Trần Linh Mai",
      "avatar": "https://th.bing.com/th/id/OIP.P8F796BGNue4Lu2SImT1bgAAAA$1rs=1&pid=ImgDetMain",
      "phone": null,
      "address": null
    }
  ]
}
```

## 1.3. Get user information

GET /api/users/:id

```json
{
  "id": 4,
  "username": "anhtuyet",
  "permission": 0,
  "provider_id": null,
  "created_at": "2025-01-08T16:44:02.205Z",
  "login_provider": "Local",
  "fullname": "Nguyễn Ánh Tuyết",
  "avatar": "https://th.bing.com/th/id/OIP.P8F796BGNue4Lu2SImT1bgAAAA$1rs=1&pid=ImgDetMain",
  "phone": null,
  "address": null,
  "balance" : 1000.45
}
```

## 1.4. Reset password

PUT /api/users/reset-password

```json
{
  "password": "123456789"
}
```

## 1.5. Delete user

DELETE /api/users/:id

## 1.6. Get my profile

GET /api/user/me

# 2. USER AUTHENTICATION

## 2.1. Local login

POST /api/auth/login/local

```json
{
  "username": "quanghai@12",
  "password": "123456789"
}
```

## 2.2. Facebook login

GET /api/auth/login/facebook

## 2.3. Google login

GET /api/auth/login/google

## 2.4. Logout

GET /api/auth/logout

# 3. PRODUCT MANAGEMENT

## 3.1. Get  special product

GET /api/products/home?max=5&page_size=&current_page=1

```json
{
  "paging": {
    "current_page": 1,
    "page_size": 2,
    "total_pages": 3,
    "total_items": 15
  },
  "newProducts": [
    {
      "id": 45,
      "name": "Dell XPS 13 2-in-1",
      "price": "1099.99",
      "description": "Dell XPS 13 2-in-1, Intel Core i7, 13.4' FHD Touch, 16GB RAM, 512GB SSD, Windows 11 Pro",
      "stock": 6,
      "discount": "0.20",
      "created_at": "2025-01-11T17:32:25.904Z",
      "category": "Mac Games & Accessories",
      "category_id": 9,
      "manufacturer": "Sony",
      "manufacturer_id": 3,
      "tag": "new",
      "images": [
        {
          "image_url": "https://m.media-amazon.com/images/I/815uX7wkOZS._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/81divYKpeTL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/61gKkYQn6lL._AC_UY218_.jpg"
        }
      ]
    },
    {
      "id": 33,
      "name": "Acer Aspire 5 Slim Laptop",
      "price": "364.99",
      "description": "Acer Aspire 5 Slim Laptop, 15.6 inches Full HD IPS Display, AMD Ryzen 3 3200U, Vega 3 Graphics, 4GB DDR4, 128GB SSD, Backlit Keyboard, Windows 10 in S Mode, A515-43-R19L, Silver",
      "stock": 15,
      "discount": "0.00",
      "created_at": "2025-01-11T17:32:25.904Z",
      "category": "PC Game Headsets",
      "category_id": 1,
      "manufacturer": "Lenovo",
      "manufacturer_id": 6,
      "tag": "new",
      "images": [
        {
          "image_url": "https://m.media-amazon.com/images/I/71yFp1LYUGL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/61IpRGnny7L._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/71CGrFl8cjL._AC_UY218_.jpg"
        }
      ]
    }
  ],
  "bestSellingProducts": [
    {
      "id": 13,
      "name": "Logitech G915 LIGHTSPEED RGB Mechanical Gaming Keyboard",
      "price": "149.99",
      "description": "Logitech G915 LIGHTSPEED RGB Mechanical Gaming Keyboard, Low Profile GL Clicky Key Switch, LIGHTSYNC RGB, Advanced LIGHTSPEED Wireless and Bluetooth Support - Clicky,Black",
      "stock": 25,
      "discount": "0.30",
      "created_at": "2025-01-11T17:32:25.904Z",
      "category": "PC Game Headsets",
      "category_id": 1,
      "manufacturer": "Samsung",
      "manufacturer_id": 2,
      "tag": "new",
      "images": [
        {
          "image_url": "https://m.media-amazon.com/images/I/61-kTKQuDUL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/61gSpxZTZZL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/71Yp7pxBFOL._AC_UY218_.jpg"
        }
      ]
    },
    {
      "id": 35,
      "name": "Lenovo IdeaPad 3 Laptop",
      "price": "299.99",
      "description": "Lenovo IdeaPad 3 Laptop, 15.6 HD Display, AMD Ryzen 3 3250U, 4GB RAM, 128GB Storage, Windows 10 Home in S Mode, 81W1009DUS, Abyss Blue",
      "stock": 10,
      "discount": "0.00",
      "created_at": "2025-01-11T17:32:25.904Z",
      "category": "PC Game Headsets",
      "category_id": 1,
      "manufacturer": "Samsung",
      "manufacturer_id": 2,
      "tag": "new",
      "images": [
        {
          "image_url": "https://m.media-amazon.com/images/I/614Jk1dIoGL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/61jojpe4KVL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/71pTP-ll4sL._AC_UY218_.jpg"
        }
      ]
    }
  ],
  "highestDiscountProducts": [
    {
      "id": 17,
      "name": "MOFII Wireless Keyboard",
      "price": "39.99",
      "description": "MOFII Wireless Keyboard and Mouse Combo,2.4GHz Retro Full-Size Keyboard with Number Pad and Cute Mouse for Computer PC Desktops Laptop WindowsxP/7/8/10(Milk Tea Colorful-B)",
      "stock": 34,
      "discount": "0.60",
      "created_at": "2025-01-11T17:32:25.904Z",
      "category": "PlayStation 4 Headsets",
      "category_id": 4,
      "manufacturer": "Sony",
      "manufacturer_id": 3,
      "tag": "featured",
      "images": [
        {
          "image_url": "https://m.media-amazon.com/images/I/61qhXbkJvTL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/61-kTKQuDUL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/71Yp7pxBFOL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/61RM1rMoceL._AC_UY218_.jpg"
        }
      ]
    },
    {
      "id": 16,
      "name": "MageGee Portable Mechanical Gaming Keyboard",
      "price": "30.99",
      "description": "MageGee Portable 60% Mechanical Gaming Keyboard, MK-Box LED Backlit Compact 68 Keys Mini Wired Office Keyboard with Red Switch for Windows Laptop PC Mac - Black/Grey",
      "stock": 95,
      "discount": "0.50",
      "created_at": "2025-01-11T17:32:25.904Z",
      "category": "Computer Headsets",
      "category_id": 5,
      "manufacturer": "Apple",
      "manufacturer_id": 4,
      "tag": "new",
      "images": [
        {
          "image_url": "https://m.media-amazon.com/images/I/61is2ZwnHEL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/61-kTKQuDUL._AC_UY218_.jpg"
        },
        {
          "image_url": "https://m.media-amazon.com/images/I/715mkYSaJLL._AC_UY218_.jpg"
        }
      ]
    }
  ]
}
```

## 3.2. Get all products

GET /api/products?category_id=1&tag=new&&search=samsung&page_size=10&current_page=1 
tag=[new|sale|featured]
price_min=[],
price_max=[],
order=[product_id_asc|product_id_desc|product_name_asc|product_name_desc|price_asc|price_desc|create_at_asc|create_at_desc]

```json
  {
    "paging": {
      "total_page": 1,
      "total_item": 2,
      "current_page": 1,
      "page_size": 10
    },
    "query": {
      "search": "samsung",
      "category_id": 1,
      "tag": "new",
      "price_min": null,
      "price_max": null
    },
    "products": [
      {
        "id": 13,
        "name": "Logitech G915 LIGHTSPEED RGB Mechanical Gaming Keyboard",
        "price": "149.99",
        "description": "Logitech G915 LIGHTSPEED RGB Mechanical Gaming Keyboard, Low Profile GL Clicky Key Switch, LIGHTSYNC RGB, Advanced LIGHTSPEED Wireless and Bluetooth Support - Clicky,Black",
        "stock": 25,
        "discount": "0.30",
        "created_at": "2025-01-11T17:32:25.904Z",
        "category": "PC Game Headsets",
        "category_id": 1,
        "manufacturer": "Samsung",
        "manufacturer_id": 2,
        "tag": "new",
        "images": [
          {
            "image_url": "https://m.media-amazon.com/images/I/61-kTKQuDUL._AC_UY218_.jpg"
          },
          {
            "image_url": "https://m.media-amazon.com/images/I/61gSpxZTZZL._AC_UY218_.jpg"
          },
          {
            "image_url": "https://m.media-amazon.com/images/I/71Yp7pxBFOL._AC_UY218_.jpg"
          }
        ]
      },
      {
        "id": 35,
        "name": "Lenovo IdeaPad 3 Laptop",
        "price": "299.99",
        "description": "Lenovo IdeaPad 3 Laptop, 15.6 HD Display, AMD Ryzen 3 3250U, 4GB RAM, 128GB Storage, Windows 10 Home in S Mode, 81W1009DUS, Abyss Blue",
        "stock": 10,
        "discount": "0.00",
        "created_at": "2025-01-11T17:32:25.904Z",
        "category": "PC Game Headsets",
        "category_id": 1,
        "manufacturer": "Samsung",
        "manufacturer_id": 2,
        "tag": "new",
        "images": [
          {
            "image_url": "https://m.media-amazon.com/images/I/614Jk1dIoGL._AC_UY218_.jpg"
          },
          {
            "image_url": "https://m.media-amazon.com/images/I/61jojpe4KVL._AC_UY218_.jpg"
          },
          {
            "image_url": "https://m.media-amazon.com/images/I/71pTP-ll4sL._AC_UY218_.jpg"
          }
        ]
      }
    ]
  }
```

## 3.3. Get product details

GET /api/products/details?product_id=1

```json
  {
    "id": 1,
    "name": "Logitech MK270 Wireless Keyboard And Mouse Combo",
    "price": "27.99",
    "discount": "0.00",
    "stock": 100,
    "created_at": "2025-01-11T17:32:25.904Z",
    "tag": "featured",
    "category_id": 6,
    "categoryName": "Computer Keyboards, Mice & Accessories",
    "manufacturer_id": 1,
    "manufacturerName": "Logitech",
    "description": "Logitech MK270 Wireless Keyboard And Mouse Combo For Windows, 2.4 GHz Wireless, Compact Mouse, 8 Multimedia And Shortcut Keys, For PC, Laptop - Black",
    "images": [
      {
        "image_url": "https://m.media-amazon.com/images/I/61gSpxZTZZL._AC_UY218_.jpg"
      },
      {
        "image_url": "https://m.media-amazon.com/images/I/7113k0qNblL._AC_UY218_.jpg"
      },
      {
        "image_url": "https://m.media-amazon.com/images/I/61j3wQheLXL._AC_UY218_.jpg"
      },
      {
        "image_url": "https://m.media-amazon.com/images/I/61is2ZwnHEL._AC_UY218_.jpg"
      }
    ],
    "attributes": [
      {
        "name": "Connectivity",
        "value": "Wired"
      },
      {
        "name": "Key Switch Type",
        "value": "Mechanical (Cherry MX Red)"
      },
      {
        "name": "Backlight",
        "value": "RGB"
      },
      {
        "name": "Layout",
        "value": "Full-size"
      },
      {
        "name": "Compatibility",
        "value": "Windows"
      }
    ],
    "relatedProducts": [
      {
        "id": 2,
        "name": "NPET K10V3PRO Gaming Keyboard",
        "price": "12.49",
        "images": [
          {
            "image_url": "https://m.media-amazon.com/images/I/61is2ZwnHEL._AC_UY218_.jpg"
          },
          {
            "image_url": "https://m.media-amazon.com/images/I/61gSpxZTZZL._AC_UY218_.jpg"
          },
          {
            "image_url": "https://m.media-amazon.com/images/I/61j3wQheLXL._AC_UY218_.jpg"
          }
        ]
      }
    ]
  }
```

## 3.4. Create product

POST /api/products/create

```json
{
  "name": "Logitech MK270 Wireless Keyboard And Mouse Combo",
  "price": 27.99,
  "description": "Logitech MK270 Wireless Keyboard And Mouse Combo For Windows, 2.4 GHz Wireless, Compact Mouse, 8 Multimedia And Shortcut Keys, For PC, Laptop - Black",
  "stock": 100,
  "discount": 0.00,
  "category_id": 6,
  "manufacturer_id": 1,
  "tag": "featured",
  "images": [
    {
      "image_url": "https://m.media-amazon.com/images/I/61gSpxZTZZL._AC_UY218_.jpg"
    },
    {
      "image_url": "https://m.media-amazon.com/images/I/7113k0qNblL._AC_UY218_.jpg"
    },
    {
      "image_url": "https://m.media-amazon.com/images/I/61j3wQheLXL._AC_UY218_.jpg"
    },
    {
      "image_url": "https://m.media-amazon.com/images/I/61is2ZwnHEL._AC_UY218_.jpg"
    }
  ],
  "attributes": [
    {
      "name": "Connectivity",
      "value": "Wired"
    },
    {
      "name": "Key Switch Type",
      "value": "Mechanical (Cherry MX Red)"
    },
    {
      "name": "Backlight",
      "value": "RGB"
    },
    {
      "name": "Layout",
      "value": "Full-size"
    },
    {
      "name": "Compatibility",
      "value": "Windows"
    }
  ]
}
```

## 3.5. Update product information

PUT /api/products/:id

```json
{
  "price": 49.99,
  "stock": 70,
  "discount": 0.2
}
```

## 3.6 Delete product

DELETE /api/products/:id





# 4. PRODUCT REVIEW

## 4.1. Post a review

POST /api/reviews

```json
{
  "content": "This product is fantastic! It has exceeded all my expectations. The performance is top-notch and it is very user-friendly. I highly recommend it to anyone looking for a high-quality product.",
  "rating": 5,
  "product_id": 1
}
```

## 4.2. Get all reviews of a product

GET /api/reviews?id=1&page=1&size=5

```json
{
  "paging": {
    "total_item": 7,
    "total_page": 3,
    "current_page": 1,
    "page_size": 3
  },
  "average_rating": 3.7,
  "reviews": [
    {
      "id": 5,
      "content": "I'm very disappointed with this product. It stopped working after a week of use. The quality is poor and not worth the price. I would not recommend this to anyone.",
      "rating": 1,
      "posted_at": "2024-12-20T09:28:52.822Z",
      "user": {
        "id": 8,
        "name": "Nguyễn Khánh Du"
      }
    },
    {
      "id": 2,
      "content": "The product is decent for its price. It performs well for basic tasks, but don't expect high-end features. The build quality is average, but it gets the job done. Overall, a good purchase if you're on a budget.",
      "rating": 3,
      "posted_at": "2024-12-20T09:28:09.021Z",
      "user": {
        "id": 8,
        "name": "Nguyễn Khánh Du"
      }
    },
    {
      "id": 6,
      "content": "Great product! It arrived on time and in perfect condition. Easy to set up and use. I am very satisfied with my purchase and would buy again.",
      "rating": 4,
      "posted_at": "2024-12-20T09:28:57.391Z",
      "user": {
        "id": 7,
        "name": "Bùi Công Anh"
      }
    }
  ]
}
```

# 5. ORDER

## 5.1. Create order

POST /api/orders

```json
{
  "total": 100000,
  "details": [
    {
      "product_id": 1,
      "quantity": 1,
      "subtotal": 20000
    },
    {
      "product_id": 2,
      "quantity": 2,
      "subtotal": 80000
    }
  ]
}
```

## 5.2. Get order by user id
GET api/orders/history?

order=[id_asc|id_desc|total_asc|total_desc|date_asc|date_desc]
date=YYYY-MM-DD
status=[completed|pending]
page=1&size=10

``` json
{
  "paging": {
    "current_page": 1,
    "page_size": 10,
    "total_item": 1,
    "total_page": 1
  },
  "filter": {
    "order": "id_asc",
    "date": "2024-09-23"
  },
  "orders": [
    {
      "total_item": 1,
      "order_id": 210,
      "total": "1243",
      "status": "completed",
      "order_date": "2024-09-22T18:00:09.000Z",
      "details": [
        {
          "id": 452,
          "product": {
            "id": 32,
            "name": "HP 15 Laptop",
            "price": 599.99,
            "images": [
              "https://m.media-amazon.com/images/I/71yaP7euNAL._AC_UY218_.jpg",
              "https://m.media-amazon.com/images/I/61Wc1fDGJuL._AC_UY218_.jpg",
              "https://m.media-amazon.com/images/I/51PsNbMd-CL._AC_UY218_.jpg"
            ],
            "category": {
              "id": 1,
              "name": "PC Game Headsets"
            },
            "manufacturer": {
              "id": 5,
              "name": "Asus"
            }
          },
          "quantity": 1,
          "subtotal": 480
        },
        {
          "id": 453,
          "product": {
            "id": 74,
            "name": "Vivo X70 Pro, 256GB, Cosmic Black",
            "price": 729.99,
            "images": [
              "https://m.media-amazon.com/images/I/51fYXSnSu9L._AC_UY218_.jpg",
              "https://m.media-amazon.com/images/I/515zGEaozeL._AC_UY218_.jpg",
              "https://m.media-amazon.com/images/I/71L1ezoIH9L._AC_UY218_.jpg"
            ],
            "category": {
              "id": 10,
              "name": "Mac Gaming Keyboards"
            },
            "manufacturer": {
              "id": 5,
              "name": "Asus"
            }
          },
          "quantity": 1,
          "subtotal": 693
        },
        {
          "id": 454,
          "product": {
            "id": 30,
            "name": "SteelSeries Arctis 3 - All-Platform Gaming Headset",
            "price": 69.99,
            "images": [
              "https://m.media-amazon.com/images/I/71IvAZyaR0L._AC_UY218_.jpg",
              "https://m.media-amazon.com/images/I/61dVV8sjTLL._AC_UY218_.jpg",
              "https://m.media-amazon.com/images/I/71hGqu8tSCL._AC_UY218_.jpg"
            ],
            "category": {
              "id": 3,
              "name": "PlayStation 5 Headsets"
            },
            "manufacturer": {
              "id": 3,
              "name": "Sony"
            }
          },
          "quantity": 1,
          "subtotal": 70
        }
      ]
    }
  ]
}
```

# 6. CART
## 6.1. Get cart by user id
GET /api/carts?page=1&size=2
``` json
{
  "user_id": 1,
  "paging": {
    "total_item": 5,
    "total_page": 3,
    "current_page": 1,
    "page_size": 2
  },
  "items": [
    {
      "product": {
        "id": 1,
        "name": "Product 1",
        "price": 100000,
        "discount": 0.1,
        "quantity": 20,
        "images": [
          "https://example.com/image1.jpg",
          "https://example.com/image2.jpg"
        ]
      },
      "quantity": 2
    },
    {
      "product": {
        "id": 2,
        "name": "Apple",
        "price": 1,
        "discount": 0.1,
        "quantity": 100,
        "images": [
          "https://example.com/image3.jpg",
          "https://example.com/image4.jpg"
        ]
      },
      "quantity": 2
    }
  ]
}
```

## 6.2. Add product to cart

Add one product to cart
POST /api/carts
```json
{
    "product_id": 1,
}
```

Add multiple products to cart
POST /api/carts/items
```json 
{
    "items": [
        {
            "product_id": 1,
            "quantity": 2
        },
        {
            "product_id": 2,
            "quantity": 3
        }
    ]
}
```

## 6.3. Update product quantity in cart
PUT /api/carts/:id
```json
{
    "quantity": 3
}
```

## 6.4. Delete product from cart
DELETE /api/carts/:id

# 7. VIEW STATISTICS

## 7.1. Get new user count by month
GET /api/users/statistics/new-users

```json
[
  {
    "month": 12,
    "year": 2024,
    "quantity": 6
  }
]
```

## 7.2. Get total revenue by month
GET /api/orders/statistics/revenue

```json
[
  {
    "month": 12,
    "year": 2024,
    "sales": 200000
  },
  {
    "month": 1,
    "year": 2025,
    "sales": 100000
  }
]
```

## 7.3. Get top best-selling products
GET /api/orders/statistics/best-sellers?limit=5

```json
[
  {
    "id": 5,
    "name": "Date",
    "quantity": 4
  },
  {
    "id": 3,
    "name": "Banana",
    "quantity": 3
  },
  {
    "id": 4,
    "name": "Cherry",
    "quantity": 2
  },
  {
    "id": 2,
    "name": "Apple",
    "quantity": 2
  },
  {
    "id": 1,
    "name": "Product 1",
    "quantity": 1
  }
]
```

## 7.4. Get top customers by total spending
GET /api/orders/statistics/top-customer?limit=5

```json
[
  {
    "id": 8,
    "name": "Nguyễn Khánh Du",
    "total": 200000
  },
  {
    "id": 7,
    "name": "Bùi Công Anh",
    "total": 100000
  }
]
```

## 7.5. Statistic product quantity by category
GET /api/products/statistics/category

```json
[
  {
    "id": 1,
    "name": "PC Game Headsets",
    "quantity": 528
  },
  {
    "id": 2,
    "name": "Computers & Tablets",
    "quantity": 372
  },
  {
    "id": 6,
    "name": "Computer Keyboards, Mice & Accessories",
    "quantity": 249
  },
  {
    "id": 3,
    "name": "PlayStation 5 Headsets",
    "quantity": 211
  }
]
```

## 7.6. Statistic product quantity by manufacturer
GET /api/products/statistics/manufacturer

```json
[
  {
    "id": 1,
    "name": "Logitech",
    "quantity": 391
  },
  {
    "id": 2,
    "name": "Samsung",
    "quantity": 384
  },
  {
    "id": 3,
    "name": "Sony",
    "quantity": 380
  }
]
```
# 8. CATEGORY MANAGEMENT

## 8.1. Get all categories

GET /api/categories

```json
{
  "categories": [
    {
      "category_id": 1,
      "name": "PC Game Headsets",
      "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
      "description": null,
      "super_category_id": null,
      "created_at": "2025-01-11T17:32:25.891Z",
      "updated_at": "2025-01-11T17:32:25.891Z",
      "product_in_category": "17",
      "index": 1,
      "children": [
        {
          "category_id": 4,
          "name": "PlayStation 4 Headsets",
          "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
          "description": null,
          "super_category_id": 1,
          "created_at": "2025-01-11T17:32:25.891Z",
          "updated_at": "2025-01-11T17:32:25.891Z",
          "product_in_category": "8",
          "index": 2
        },
        {
          "category_id": 2,
          "name": "Computers & Tablets",
          "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
          "description": null,
          "super_category_id": 1,
          "created_at": "2025-01-11T17:32:25.891Z",
          "updated_at": "2025-01-11T17:32:25.891Z",
          "product_in_category": "14",
          "index": 3,
          "children": [
            {
              "category_id": 3,
              "name": "PlayStation 5 Headsets",
              "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
              "description": null,
              "super_category_id": 2,
              "created_at": "2025-01-11T17:32:25.891Z",
              "updated_at": "2025-01-11T17:32:25.891Z",
              "product_in_category": "10",
              "index": 4,
              "children": [
                {
                  "category_id": 6,
                  "name": "Computer Keyboards, Mice & Accessories",
                  "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
                  "description": null,
                  "super_category_id": 3,
                  "created_at": "2025-01-11T17:32:25.891Z",
                  "updated_at": "2025-01-11T17:32:25.891Z",
                  "product_in_category": "7",
                  "index": 5
                },
                {
                  "category_id": 9,
                  "name": "Mac Games & Accessories",
                  "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
                  "description": null,
                  "super_category_id": 3,
                  "created_at": "2025-01-11T17:32:25.891Z",
                  "updated_at": "2025-01-11T17:32:25.891Z",
                  "product_in_category": "3",
                  "index": 6,
                  "children": [
                    {
                      "category_id": 10,
                      "name": "Mac Gaming Keyboards",
                      "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
                      "description": null,
                      "super_category_id": 9,
                      "created_at": "2025-01-11T17:32:25.891Z",
                      "updated_at": "2025-01-11T17:32:25.891Z",
                      "product_in_category": "3",
                      "index": 7
                    }
                  ]
                }
              ]
            },
            {
              "category_id": 8,
              "name": "Computer Keyboard & Mouse Combos",
              "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
              "description": null,
              "super_category_id": 2,
              "created_at": "2025-01-11T17:32:25.891Z",
              "updated_at": "2025-01-11T17:32:25.891Z",
              "product_in_category": "7",
              "index": 8
            }
          ]
        },
        {
          "category_id": 7,
          "name": "Computer Keyboards",
          "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
          "description": null,
          "super_category_id": 1,
          "created_at": "2025-01-11T17:32:25.891Z",
          "updated_at": "2025-01-11T17:32:25.891Z",
          "product_in_category": "5",
          "index": 9
        },
        {
          "category_id": 5,
          "name": "Computer Headsets",
          "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
          "description": null,
          "super_category_id": 1,
          "created_at": "2025-01-11T17:32:25.891Z",
          "updated_at": "2025-01-11T17:32:25.891Z",
          "product_in_category": "6",
          "index": 10
        }
      ]
    }
  ]
}
```

## 8.2. Get category infomation

GET /api/categories/1

```json
{
    "category_id": 1,
    "name": "PC Game Headsets",
    "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
    "description": null,
    "super_category_id": null,
    "created_at": "2025-01-11T17:32:25.891Z",
    "updated_at": "2025-01-11T17:32:25.891Z"
}
```

## 8.3. Get all products in category

GET api/categories/products/1

```json
{
    "products": [
        {
            "id": 3,
            "name": "SteelSeries Arctis Nova 1 Multi-System Gaming Headset",
            "price": "59.99",
            "description": "SteelSeries Arctis Nova 1 Multi-System Gaming Headset — Hi-Fi Drivers — 360° Spatial Audio — Comfort Design — Durable — Ultra Lightweight — Noise-Cancelling Mic — PC, PS5/PS4, Switch, Xbox - Black",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 80,
            "discount": "0.00",
            "manufacturer": "Sony",
            "category": "PC Game Headsets",
            "tag": "new",
            "images": [
                "https://m.media-amazon.com/images/I/61+zbpS+6iL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/61jBnY6paeL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71Ly9zSMVnL._AC_UY218_.jpg"
            ]
        },
        {
            "id": 7,
            "name": "Logitech K270 Wireless Keyboard for Windows",
            "price": "24.95",
            "description": "Logitech K270 Wireless Keyboard for Windows, 2.4 GHz Wireless, Full-Size, Number Pad, 8 Multimedia Keys, 2-Year Battery Life, Compatible with PC, Laptop, Black",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 67,
            "discount": "0.00",
            "manufacturer": "Asus",
            "category": "PC Game Headsets",
            "tag": "featured",
            "images": [
                "https://m.media-amazon.com/images/I/51jkxo3a7bL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/61RM1rMoceL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/715mkYSaJLL._AC_UY218_.jpg"
            ]
        },
        {
            "id": 8,
            "name": "Dell Wired Keyboard",
            "price": "15.99",
            "description": "Dell Wired Keyboard - Black KB216 (580-ADMT)",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 70,
            "discount": "0.00",
            "manufacturer": "Lenovo",
            "category": "PC Game Headsets",
            "tag": "new",
            "images": [
                "https://m.media-amazon.com/images/I/51jkxo3a7bL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/61Q56A7UfNL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71ehwfAM4-L._AC_UY218_.jpg"
            ]
        },
        {
            "id": 12,
            "name": "Arteck 2.4G Wireless Keyboard",
            "price": "49.99",
            "description": "Arteck 2.4G Wireless Keyboard Stainless Steel Ultra Slim Full Size Keyboard with Numeric Keypad for Computer/Desktop/PC/Laptop/Surface/Smart TV and Windows 10/8/ 7 Built in Rechargeable Battery",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 65,
            "discount": "0.00",
            "manufacturer": "Samsung",
            "category": "PC Game Headsets",
            "tag": "featured",
            "images": [
                "https://m.media-amazon.com/images/I/61is2ZwnHEL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/61j3wQheLXL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71Yp7pxBFOL._AC_UY218_.jpg"
            ]
        },
        {
            "id": 13,
            "name": "Logitech G915 LIGHTSPEED RGB Mechanical Gaming Keyboard",
            "price": "149.99",
            "description": "Logitech G915 LIGHTSPEED RGB Mechanical Gaming Keyboard, Low Profile GL Clicky Key Switch, LIGHTSYNC RGB, Advanced LIGHTSPEED Wireless and Bluetooth Support - Clicky,Black",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 25,
            "discount": "0.30",
            "manufacturer": "Samsung",
            "category": "PC Game Headsets",
            "tag": "new",
            "images": [
                "https://m.media-amazon.com/images/I/61-kTKQuDUL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/61gSpxZTZZL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71Yp7pxBFOL._AC_UY218_.jpg"
            ]
        },
        {
            "id": 18,
            "name": "WallarGe Wireless Headphones",
            "price": "40.99",
            "description": "WallarGe Wireless Headphones for TV Watching, TV Headphones Wireless for Seniors, Easy Setup and Comfortable Bluetooth Headphones, 25 Hours Play and No Audio Delay",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 45,
            "discount": "0.00",
            "manufacturer": "Asus",
            "category": "PC Game Headsets",
            "tag": "new",
            "images": [
                "https://m.media-amazon.com/images/I/51FRJHB7XOL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/61j6ey6mBpL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71Ly9zSMVnL._AC_UY218_.jpg"
            ]
        },
        {
            "id": 19,
            "name": "Bose QuietComfort 35 II Wireless Headphones",
            "price": "29.00",
            "description": "Bose QuietComfort 35 II Wireless Bluetooth Headphones, Noise-Cancelling, with Alexa Voice Control - Black",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 15,
            "discount": "0.20",
            "manufacturer": "Lenovo",
            "category": "PC Game Headsets",
            "tag": "sale",
            "images": [
                "https://m.media-amazon.com/images/I/61o3FbQ6OJL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71IL4SsThNL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71yqtewaJKL._AC_UY218_.jpg"
            ]
        },
        {
            "id": 31,
            "name": "HP 14 Ultral Light Laptop for Students and Business",
            "price": "349.99",
            "description": "HP 14 Ultral Light Laptop for Students and Business, Intel Quad-Core, 8GB RAM, 192GB Storage(64GB eMMC+128GB Ghost Manta SD Card), 1 Year Office 365, USB C, Win 11 S",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 30,
            "discount": "0.00",
            "manufacturer": "Asus",
            "category": "PC Game Headsets",
            "tag": "new",
            "images": [
                "https://m.media-amazon.com/images/I/61gKkYQn6lL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/815uX7wkOZS._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/81divYKpeTL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/81yG4KTmOpL._AC_UY218_.jpg"
            ]
        },
        {
            "id": 32,
            "name": "HP 15 Laptop",
            "price": "599.99",
            "description": "HP 15 Laptop, 11th Gen Intel Core i5-1135G7 Processor, 8 GB RAM, 256 GB SSD Storage, 15.6” Full HD IPS Display, Windows 11 Home, HP Fast Charge, Lightweight Design (15-dy2021nr, 2020)",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 20,
            "discount": "0.20",
            "manufacturer": "Asus",
            "category": "PC Game Headsets",
            "tag": "sale",
            "images": [
                "https://m.media-amazon.com/images/I/51PsNbMd-CL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/61Wc1fDGJuL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71yaP7euNAL._AC_UY218_.jpg"
            ]
        },
        {
            "id": 33,
            "name": "Acer Aspire 5 Slim Laptop",
            "price": "364.99",
            "description": "Acer Aspire 5 Slim Laptop, 15.6 inches Full HD IPS Display, AMD Ryzen 3 3200U, Vega 3 Graphics, 4GB DDR4, 128GB SSD, Backlit Keyboard, Windows 10 in S Mode, A515-43-R19L, Silver",
            "added": "2025-01-11T17:32:25.904Z",
            "stock": 15,
            "discount": "0.00",
            "manufacturer": "Lenovo",
            "category": "PC Game Headsets",
            "tag": "new",
            "images": [
                "https://m.media-amazon.com/images/I/61IpRGnny7L._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71CGrFl8cjL._AC_UY218_.jpg",
                "https://m.media-amazon.com/images/I/71yFp1LYUGL._AC_UY218_.jpg"
            ]
        }
    ],
    "paging": {
        "total_item": 17,
        "total_page": 2,
        "current_page": 1,
        "page_size": 10
    }
}
```

## 8.4. Create category

POST api/categories/

```json
{
    "name": "PC Game Headsets",
    "thumbnail": "https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c=",
    "description": "Category for PC game headsets",
    "super_category_id": null
}
```

## 8.5 Update category information

PUT api/categories/1

```json
{
  "name": "PC Game Keyboard",
  "supper_category": 3,
  "description": " Category for PC Game Keyboard"
}
```

## 8.6 Delete category

DELETE api/categories/1


# 9. MANAFACTURER MANAGEMENT

## 9.1 Get all manufacturer

GET api/manufacturers

```json
{
    "manufacturers": [
        {
            "id": 1,
            "name": "Logitech"
        },
        {
            "id": 2,
            "name": "Samsung"
        },
        {
            "id": 3,
            "name": "Sony"
        },
        {
            "id": 4,
            "name": "Apple"
        },
        {
            "id": 5,
            "name": "Asus"
        },
        {
            "id": 6,
            "name": "Lenovo"
        }
    ]
}
```

## 9.2. Get manufacturer information

GET /api/manufacturers/1

```json
{
    "manufacturer_id": 1,
    "manufacturer_name": "Logitech"
}
```

## 9.3  Delete manufacturer

DELETE /api/manufacturers/1
