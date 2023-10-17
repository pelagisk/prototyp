# File Archive for Prototyp

## Structure

- backend: API written in Golang Gin with SQLite database
- frontend: React and Axios

## How to install and run

For backend, install `golang >= 1.21.1` and run

```bash
cd backend
go mod tidy
go run .
```

For frontend, run

```bash
cd frontend
npm install
npm start
```

Note that both are running in dev mode. I had to allow CORS to let them communicate on localhost. 

## Improvements

The bonus features:
- Renaming file when uploading
- List sorting
- Pagination

Additionally:
- Containerization (Docker)
- Test/production mode
- Frontend testing
- Allow greater flexibility of adding new fields to metatata/database? Right now, it takes a lot of manual work to change anything in frontend resp. backend. Perhaps using Golang `reflect` package and iterating over struct fields?
