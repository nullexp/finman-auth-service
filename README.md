
# FinMan Auth Service

FinMan Auth Service is a microservice responsible for handling authentication and JWT token management. This service provides functionalities to create tokens based on user credentials and validates tokens for authentication if required.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Environment Variables](#environment-variables)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

- Go 1.22 or higher
- Docker

## Environment Variables

Create a `.env` file in the root directory of the project and add the following environment variables:

```env
# .env
JWT_SECRET=eDM!":jmx2/QoHBlY'.O8e4?Uy,",9
JWT_EXPIRE_MINUTE=20
PORT=8080
IP=0.0.0.0
```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/finman-auth-service.git
   cd finman-auth-service
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

## Usage

### Running the Service

You can run the service locally using the following command:

```bash
go run main.go
```

Alternatively, you can run the service using Docker:

1. Build the Docker image:

   ```bash
   docker build -t finman-auth-service .
   ```

2. Run the Docker container:

   ```bash
   docker run --env-file .env -p 8080:8080 finman-auth-service
   ```

### Environment Variables

The service uses the following environment variables:

- `JWT_SECRET`: The secret key used to sign the JWT tokens.
- `JWT_EXPIRE_MINUTE`: The expiration time for JWT tokens in minutes.
- `PORT`: The port on which the service will run.
- `IP`: The IP address on which the service will bind.

## Testing

You can run the tests using the following command:

```bash
go test ./...
```

For testing with Docker, use the following command:

```bash
docker build -t finman-auth-service-test -f Dockerfile.test .
docker run --env-file .env finman-auth-service-test
```

