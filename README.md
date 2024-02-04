# About This Service

```
  SERVER
  LOCAL: [http](http://localhost:8111)
  DEVELOPMENT: [http](http://dealls.muhammadikhsan.my.id)

  DATABASE
  LOCAL: [http](http://localhost:5432)
  DEVELOPMENT: [http](http://103.127.97.209:5432)
```

## Structure
This service adopts a clean architecture that has been adjusted to support development needs that are faster and easier to understand with a structure :
  - Main Structure
    - Delivery
      - Router
      - Contract
      - Handler
    - Domain
      - UseCase
        - Model
        - Service
      - Data
        - Entity
        - Repository
  - Support Structure
    - Material
      - Client
        - PostgreSQL
      - Contract
      - Generator
      - Helper
      - Middleware
      - Modules
      - Secret
      - Static

## Tech Stach
  - Customized Golang Chi
    - Great performance and easy to customize
  - GORM
    - Can help to make interaction with the database faster
  - PostgreSQL
    - One of the popular and familiar databases
    - Has better security than similar databases

## Requirements System
  ### Functional
  - Authentication & Authorization
    - JWT Token
    - Basic Authentication
    - Http Cookies
  ### Non Functional
  - 

## How To Run
  ### preparation
  - create `.env` file
    - if run local create `.env.local`
    - if run development create `.env.development`
    - if run production create `.env.production`
    - if run test create `.env.test`

  ### Execution
  - open bash terminal (required)
  - run `make docker-up` to install postgreql (local only)
  - run `make migrate` to migrate database (local only)
  - run `go mod download` to download all package
  - run `go mod tidy` to tidy all package
  - run `go work sync` to sync all workspace
  - run `make serve env=LOCAL`
    - if run local `env=LOCAL`
    - if run development `env=DEVELOPMENT`
    - if run production `env=PRODUCTION`

  ### Testing
  - open bash terminal (required)
  - run `make tester`

  ### Build Docker Image
  - open bash terminal (required)
  - run `make docker-build`

  ### Deployment
  - VM from Biznet Gio
  - OS Debian 11
  - Containers directly using Docker