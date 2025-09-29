# Coding Challenge: Prices Tick Processor

## Overview

This repository contains a partially implemented application that generates
random ticks and sends them over a channel to a tick consumer.
The tick consumer is responsible for storing the ticks to a data store and
calculating and storing candles every minute.

## Constraints

Please disable any AI code generation tools while working on this challenge.
However, you are free to use any language server code completion tools that
are available with your IDE/editor.

You are free to use any publicly available libraries or package to help you
complete the challenge. Please document any external libraries you use in
your `SOLUTION.md` file, and reasons for your choices.

You may make use of online resources to help you complete the challenge, with
the exception of using AI code generation tools.

## Tasks

### 1. Tick Consumer and Data Persistence

- Create the database tables for storing ticks and candles.
- Implement the tick persistence to the Postgres database.
- Implement the calculation of 1 minute candles from the incoming ticks,
  and store these candles in the same Postgres database.
- Ensure the consumer is robust and can handle errors gracefully.

### 2. REST API Implementation

- Implement the REST API to fetch tick data and candle data from the data store.
- Choose appropriate defaults for limit and offset for pagination.
- For candles if no interval is provided, default to 1 minute.
- For candles, the endpoint should support querying the assetID, time range and
  granularity (e.g., 1 minute, 5 minutes). By storing 1 minute candle, you can
  derive higher granularity candles using aggregation.
- Implement middleware for metrics to provide information to the health endpoint.

### 3. Multi-Worker Setup

- Enhance the consumer to support multiple workers grouped by `assetID`.
- Demonstrate how the system scales with multiple workers and handles concurrency.
- Implement a graceful exit path for concurrent workers.

### 4. Testing and Mocking

- The repository includes a Docker setup for Postgres to help you test your implementation.
- Create unit tests to demonstrate the consumer implementation is working as
  expected. Mock the datastore interface to isolate the consumer logic.
- Create unit tests for datastore functions.

## Evaluation Criteria

- **Functionality**: Your implementation should work as expected and meet the
  requirements outlined above.
- **Code Quality**: Write clean, maintainable, and well-documented code.
- **Improvements**: Feel free to suggest and implement any improvements or bug
  fixes to the existing codebase.
- **Testing**: Include unit tests, or any other relevant tests to ensure the
  reliability of your implementation.

## Submission

- Fork this repository and work on your fork.
- Once completed, create a pull request with your changes.
- Include a `SOLUTION.md` file in your submission explaining:
  - Your approach and design decisions.
  - Any improvements or bug fixes you implemented.
  - Details about the tests you included and how to run them.

## Notes

- The codebase is generic and uses public libraries, so it should be
  straightforward to work with.
- If you have any questions or need clarification, feel free to reach out.

Good luck, and we look forward to seeing your solution!
