# Coding Challenge: Prices Tick Processor

## Overview

This repository contains a partially implemented application that generates random ticks and sends them to a Kafka publisher. The publisher writes these ticks to a Kafka topic. Your task is to complete the implementation by building the missing components and improving the existing codebase.

## Tasks

### 1. Kafka Consumer and Data Persistence
- Implement a Kafka consumer to consume price ticks from the Kafka topic.
- Persist the consumed ticks to a data store. The current setup supports Redis and Spanner. You are free to choose one or implement both.
- Ensure the consumer is robust and can handle errors gracefully.

### 2. gRPC API Implementation
- Implement the gRPC API to fetch tick data from the chosen data store.
- The scaffolding for the gRPC API is already in place. You need to add the implementation based on the data store you use.

### 3. Multi-Worker Setup
- Enhance the Kafka consumer to support multiple workers grouped by `assetID`.
- Demonstrate how the system scales with multiple workers and handles concurrency.
- Implement a graceful exit path for concurrent workers.

### 4. Testing and Mocking
- The repository includes a Docker setup for Kafka, Spanner, and Redis to help you test your implementation.
- Create unit tests to demonstrate the consumer implementation is working as expected. Mock the datastore interface to isolate the consumer logic.
- Create unit tests for datastore functions.

## Evaluation Criteria
- **Functionality**: Your implementation should work as expected and meet the requirements outlined above.
- **Code Quality**: Write clean, maintainable, and well-documented code.
- **Improvements**: Feel free to suggest and implement any improvements or bug fixes to the existing codebase.
- **Testing**: Include unit tests, or any other relevant tests to ensure the reliability of your implementation.

## Submission
- Fork this repository and work on your fork.
- Once completed, create a pull request with your changes.
- Include a `SOLUTION.md` file in your submission explaining:
  - Your approach and design decisions. 
  - Any improvements or bug fixes you implemented. 
  - Details about the tests you included and how to run them.

## Notes
- The codebase is generic and uses public libraries, so it should be straightforward to work with.
- If you have any questions or need clarification, feel free to reach out.

Good luck, and we look forward to seeing your solution!