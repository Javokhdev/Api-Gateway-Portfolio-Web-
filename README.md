# Portfolio API Gateway

#### The Portfolio API Gateway is a central point for routing requests to the various services responsible for managing user portfolios, including experiences, projects, educations, and skills. This gateway ensures that requests are properly directed to the appropriate microservices, facilitating a seamless and efficient data flow across the portfolio management system.

## Features

***• Unified Endpoint: Centralized routing for portfolio management services.***

***• Service Discovery: Dynamically discover and route requests to various microservices.***

***• Load Balancing: Distribute requests evenly across service instances.***

***• Security: Implements authentication and authorization mechanisms for secure access.***

## Installation

### 1. Clone the repository:

```sh
git clone git@github.com:Javokhdev/Portfolio-Api-Gateway.git
cd Portfolio-ApiGateway
```

### 2. Set up the environment:

***• Go (1.16+)***

***• PostgreSQL***

### 3. Install dependencies:

```
go mod tidy
```

### 4. Set up the database:

***• Create a PostgreSQL database.***

***• Run the migrations located in the migrations directory.***

### 5. Configuration:

Create a `.env` file in the root directory with the following environment variables:

```.env
HTTPPort = :7070

PostgresHost    = localhost
PostgresPort    = 5432
PostgresUser    = postgres
PostgresPassword = root
PostgresDatabase = web_portfolio
```

## Usage

### Run the service:

```
go run server/server.go
go run main.go
```


## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.

## License

This project is licensed under the MIT License.

[MIT](https://choosealicense.com/licenses/mit/)

## Acknowledgement

#### Javoxir Xasanov 
[![telegram](https://img.shields.io/badge/telegram-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://t.me/javohir_khasanov)

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/javohir-xasanov/)


