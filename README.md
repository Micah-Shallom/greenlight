# greenlight

This repository contains instructions and commands for setting up and running the Greenlight project.

## Setup Instructions

### Prerequisites

- Go version:
  ```sh
  go version
  ```

### Installations

#### API
- Check version:
  ```sh
  ./bin/api -version
  ```

#### Staticcheck
- Install latest version:
  ```sh
  go install honnef.co/go/tools/cmd/staticcheck@latest
  ```

- Verify installation:
  ```sh
  which staticcheck
  ```

### Running API
- Start API server:
  ```sh
  go run ./cmd/api
  ```

#### CORS Example
- Run simple CORS example:
  ```sh
  go run ./cmd/examples/cors/simple
  ```

### Dependencies

- Install httprouter v1:
  ```sh
  go get github.com/julienschmidt/httprouter@v1
  ```

- Install hey:
  ```sh
  go install github.com/rakyll/hey@latest
  ```

- Install pq v1:
  ```sh
  go get github.com/lib/pq@v1
  ```

- Install rate limiter:
  ```sh
  go get golang.org/x/time/rate@latest
  ```

- Install bcrypt:
  ```sh
  go get golang.org/x/crypto/bcrypt@latest
  ```

- Install mail:
  ```sh
  go get github.com/go-mail/mail/v2@v2
  ```

### Database Setup

- Set environment variable for PostgreSQL DSN:
  ```sh
  export GREENLIGHT_DB_DSN='postgres://greenlight:pa55word@localhost:5432/greenlight?sslmode=disable'
  ```

- Verify environment variable:
  ```sh
  echo $GREENLIGHT_DB_DSN
  ```

#### Migrations

- Install migrate tool:
  ```sh
  curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
  sudo apt-get update
  sudo apt-get install -y migrate
  ```

- Check migrate version:
  ```sh
  migrate -version
  ```

- Create migrations:
  ```sh
  migrate create -seq -ext=.sql -dir=./migrations create_movies_table
  migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constraints
  ```

- Apply migrations:
  ```sh
  migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
  ```

- Rollback migrations:
  ```sh
  migrate -path=./migrations -database=$EXAMPLE_DSN goto 1
  migrate -path=./migrations -database=$EXAMPLE_DSN up
  migrate -path=./migrations -database=$EXAMPLE_DSN down
  ```

- Create additional migrations:
  ```sh
  migrate create -seq -ext=.sql -dir=./migrations create_users_table
  migrate create -seq -ext=.sql -dir=./migrations create_tokens_table
  migrate create -seq -ext=.sql -dir=./migrations add_movies_indexes
  migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
  ```

#### Troubleshooting

- Fix SQL Migration Errors:
  ```sh
  migrate -path=./migrations -database=$EXAMPLE_DSN force 1
  ```

  Once forced, migrations should run without errors.

## API Usage

### Health Check
- Check server status:
  ```sh
  curl localhost:4000/v1/healthcheck
  ```

### Movies Endpoint
- Create movie:
  ```sh
  curl -X POST localhost:4000/v1/movies
  ```

- Retrieve movie details:
  ```sh
  curl localhost:4000/v1/movies/123
  ```

- Update movie:
  ```sh
  curl -X PUT -d '{"title":"Black Panther","year":2018,"runtime":134,"genres":["sci-fi","action","adventure"]}' localhost:4000/v1/movies/2
  ```

- Partially update movie:
  ```sh
  curl -X PATCH -d '{"year": 1985}' localhost:4000/v1/movies/4
  ```

- Concurrent updates:
  ```sh
  xargs -I % -P8 curl -X PATCH -d '{"runtime":97}' "localhost:4000/v1/movies/4" < <(printf '%s\n' {1..8})
  ```

- Query movies:
  ```sh
  curl "localhost:4000/v1/movies?title=moana&genres=animation,adventure&page=1&page_size=5&sort=year"
  ```

### Users Endpoint
- Create user:
  ```sh
  curl -i -d '{"name": "Alice Smith", "email": "alice@example.com", "password": "pa55word"}' localhost:4000/v1/users
  ```

- Activate user:
  ```sh
  curl -X PUT -d '{"token": "ZYGQTPU5PKKJRY7SFOAMKXPGQY"}' localhost:4000/v1/users/activated
  ```

### Authentication
- Obtain authentication token:
  ```sh
  curl -i -d '{"email": "alice@example.com", "password": "pa55word"}' localhost:4000/v1/tokens/authentication
  ```

- Access restricted resources:
  ```sh
  curl -i -H "Authorization: Bearer XXXXXXXXXXXXXXXXXXXXXXXXXX" localhost:4000/v1/healthcheck
  ```

### Debugging and Monitoring
- View debug variables:
  ```sh
  curl http://localhost:4000/debug/vars
  ```

## Additional Information

### Supported Types
- Go to JSON:
  - [List of supported Go types and their JSON equivalents]

### Database Types
- PostgreSQL and Go types:
  - [Mapping of PostgreSQL types to Go types]

### SMTP Server
- For email testing:
  - [MailTrap SMTP settings]

