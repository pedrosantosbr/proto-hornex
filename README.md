# proto-hornex

An application that register users for a tournament platform

- register

* auth
* search
* tournaments
* teams

- framework
- routers
- database

- application
- services
- repositories

- domain
- models

RegisterUserUseCase (UML)

-> input
`{
  "firstName": "",
  "lastName": "",
  "email": "",
  "dateOfBirth": "",
}`

==================

- controllers (UserParams)
- services (UserParams)
- repositories (UserParams) User

-> ouput
`{
  "id": ""
}`

# Challenge

Build a API REST in Go

# Rules

- Entity (Use Cases)
- Repository (DB)
- Service (Business Logic)
- Router (Handlers) (HTTP)
- Validator (Checks for errors)

# Entities

- User

* firstName
* lastName
* email
* dateOfBirth

# Use Cases (CRUD)

- Create User Use Case
- Update User Use Case
- Delete User Use Case
- List Users Use Case
- Get User Use Case
