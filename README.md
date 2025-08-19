# 🍴 Restaurant Management API (Go + MongoDB)

A backend API for managing a restaurant system built with **Golang (Gin framework)** and **MongoDB**.  
This project handles restaurant operations like managing menus, orders, tables, users, and invoices.

---

## 📂 Project Structure

```text
.
├── .gitignore
├── controllers/        # Request handlers (business logic for each resource)
│   ├── foodController.go
│   ├── invoiceController.go
│   ├── menuController.go
│   ├── orderController.go
│   ├── orderItemController.go
│   ├── tableController.go
│   └── userController.go
├── database/           # MongoDB connection setup
│   └── databaseConnection.go
├── helpers/            # Helper functions (e.g., JWT token handling)
│   └── tokenHelper.go
├── middleware/         # Custom middlewares (e.g., authentication)
│   └── authMiddleware.go
├── models/             # Data models (MongoDB collections)
│   ├── foodModel.go
│   ├── invoiceModel.go
│   ├── menuModel.go
│   ├── noteModel.go
│   ├── orderItemModel.go
│   ├── orderModel.go
│   ├── tableModel.go
│   └── userModel.go
├── routes/             # API route definitions
│   ├── foodRouter.go
│   ├── invoiceRouter.go
│   ├── menuRouter.go
│   ├── orderItemRouter.go
│   ├── orderRouter.go
│   ├── tableRouter.go
│   └── userRouter.go
├── main.go             # Entry point
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
└── .env                # Environment variables (ignored in Git)
```


---

## ⚡ Features
- User authentication with JWT  
- Manage **Food items, Menus, Tables, Orders, Invoices**  
- Role-based access control with middleware  
- MongoDB integration with models  
- Modular folder structure (easy to extend)  

---

## 🛠️ Tech Stack
- **Language:** Go (Golang)  
- **Framework:** Gin  
- **Database:** MongoDB  
- **Auth:** JWT  

---

## 🚀 Getting Started

### 1. Clone the repo
```
git clone https://github.com/your-username/restaurant-management-go.git
cd restaurant-management-go

```

### 2. Install dependencies
```
go mod tidy

```

### 3. Setup environment variables
```
Create a .env file in the root directory:
MONGO_URI=mongodb://localhost:27017
PORT=port
SECRET_KEY=your-secret-key
```

### 4. Run the server
```
go run main.go

```
