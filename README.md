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
