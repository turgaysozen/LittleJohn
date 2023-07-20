# Stock-Portfolio Management API
LittleJohn is a fictitious brokerage service API that allows users to manage and view their stock portfolios. This README provides an overview of the project and instructions for running it.

## Project Overview
The LittleJohn project aims to provide an HTTP API for users to access their portfolios and retrieve historical stock prices. The API supports the following endpoints:

1- Get Personal Portfolio: Retrieves the authenticated user's portfolio, which contains a list of stocks with their ticker symbols and current prices.

Endpoint: ```GET /tickers```

Example Response:
```
[
  {
    "symbol": "AAPL",
    "price": "154.30"
  },
  {
    "symbol": "AMZN",
    "price": "3478.05"
  }
]

```

2- Get Historical Prices: Retrieves the historical price fluctuations for a specific stock ticker over the last 90 days. The prices are returned in descending order by date.

Endpoint: 

- ```GET /tickers/{symbol}/history```
- ```GET /tickers/{symbol}/history?page={page}&pageSize={pageSize}```

Example Response:
```
[
  {
    "date": "2021-09-01",
    "price": "150.41"
  }
]

```

## Running the Project
To run the LittleJohn project, follow these steps:

Make sure you have Docker installed on your machine.

Clone the project repository to your local machine.

Open a terminal and navigate to the project's root directory.

Build the Docker image using the following command:

```docker build -t littlejohn . && docker run -p 8080:8080 littlejohn```

This command builds the Docker image and runs the tests inside a Docker container. It ensures that the API functions as expected.

## Predefined Users and Authentication: 
The LittleJohn API supports Basic Authentication. There are three predefined users available: "user1", "user2", and "Turgay". Use the corresponding username as the access token and leave the password empty. Include the Basic Authentication header in your API requests for authentication

## Pagination for Historical Prices: 
To retrieve the historical prices of stocks, you can use the following endpoint:

```GET /tickers/{symbol}/history?page={page}&pageSize={pageSize}```

- {symbol}: The ticker symbol of the stock.
- {page}: (Optional) The page number for pagination. Default value is 1 if not provided.
- {pageSize}: (Optional) The number of items per page. Default value is 10 if not provided.

For example, to retrieve the historical prices of AAPL stock with page 1 and pageSize 2, you can use the following API call:

```http://localhost:8080/tickers/AAPL/history?page=1&pageSize=2```

## Postman Collection:
 The project includes a Postman collection in the main root directory. You can import this collection into Postman to explore and test the available API endpoints. Simply import the collection into Postman and start making API calls to LittleJohn.

## Notes
The project uses dummy data for portfolios and historical stock prices. The data is generated dynamically and may vary for different users.

Authentication is performed using an access token sent via HTTP Basic Authentication. The token should be provided in the Authorization header with an empty password.

The portfolios can contain a minimum of 1 and a maximum of 10 stock tickers.

The project does not use any persistence mechanism, such as a database. The data is consistent across server restarts, but it is not stored persistently.

The project is designed to emulate real-life stock data behavior, where the same stock ticker may have different prices for different users.

## Conclusion
The LittleJohn project provides a basic HTTP API for managing stock portfolios and retrieving historical stock prices. You can use this README as a guide to build, run, and test the project. Feel free to extend the project and enhance its functionality to meet your specific requirements.
