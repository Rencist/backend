- Link Documentation : https://documenter.getpostman.com/view/22081311/2s946fcrxf
- Repo Front End : https://github.com/Rencist/sea-cinema-frontend

## How To Run

- Install dependencies
  ```sh
  go mod tidy
  ```
- Run server to migrate table
  ```sh
  go run main.go
  ```
- Run database seeder
  ```sh
  cd seeder | go run seeder.go
  ```
