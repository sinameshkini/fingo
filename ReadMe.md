# FinGo

**FinGo** is a simple core banking service written in Go.

---

## Features
- Manage user accounts with different account types and currencies.
- RESTFul API for seamless integration with other systems.
- Manage transactions, accounts balance and history


- [x] Configuration
  - [x] Account Type
  - [x] Currencies
  - [x] Policies
    - [x] Get
    - [x] Create
    - [x] Update
- [ ] Account
  - [x] Create new
  - [x] Get with balance
  - [x] Get list by UserID
  - [ ] Update info
- [ ] Transaction
  - [x] Transfer
  - [x] Reverse
  - [x] Inquiry
  - [x] History
  - [x] Multi Transfer
  - [x] Worker Pool
- [ ] Requests Validation


---

## Development
### Prerequisites
- [Go 1.23.1+](https://golang.org/)
- A PostgreSQL database (or other supported database via GORM)

### Start
```shell
git clone github.com/sinameshkini/fingo
cd fingo
make tidy
make test
make run
```
---

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contributing
We welcome contributions! Please fork the repository, make your changes, and submit a pull request.