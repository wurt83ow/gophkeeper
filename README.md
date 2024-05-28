### GophKeeper Server

GophKeeper is a client-server system designed to securely store and manage user credentials, binary data, and other private information. This repository contains the server-side implementation of GophKeeper, providing the backend services for user authentication, data storage, and synchronization.

#### Key Features

- **User Registration, Authentication, and Authorization**: Manages user accounts, securely authenticates users, and authorizes access to private data.
- **Secure Data Storage**: Stores encrypted private data, including login credentials, arbitrary text, binary data, and credit card information.
- **Data Synchronization**: Ensures data consistency across multiple authorized clients of the same user.
- **Data Retrieval**: Provides secure endpoints for users to retrieve their stored private data.

#### Project Structure

- **cmd/portfolio**: Contains the main server application entry point.

  - `main.go`: Entry point for the server application.
  - `Dockerfile`: Docker configuration for building the server application container.

- **static**: Contains static files for the server.

  - `dist`: Distribution files for the frontend, if applicable.
  - `icons`: Icons used in the application.
  - `technology`: Technology icons used in the application.

- **.vscode**: Visual Studio Code configuration.

  - `settings.json`: Settings for the development environment.

- **api-spec**: (Presumed location for API specifications, not explicitly listed)

  - API specifications for server endpoints.

- **config**: Configuration files for the server.

  - `default.conf`: Default server configuration.
  - `nginx.conf`: Nginx configuration for reverse proxy setup.

- **docker-compose.yml**: Docker Compose configuration for setting up the server environment.
- **sectionsData.json**: JSON data file for sections (likely configuration or initial data setup).
- **go.mod**: Go module dependencies.
- **go.sum**: Checksums for Go module dependencies.
- **cleanup_and_rebuild.sh**: Script for cleaning up and rebuilding the project.

#### Getting Started

1. **Installation**:

   - Clone the repository:
     ```sh
     git clone https://github.com/yourusername/gophkeeper-server.git
     cd gophkeeper-server
     ```

2. **Build and Run with Docker**:

   - Build the Docker image:
     ```sh
     docker build -t gophkeeper-server -f cmd/portfolio/Dockerfile .
     ```
   - Run the server with Docker Compose:
     ```sh
     docker-compose up
     ```

3. **Configuration**:
   - Ensure the `default.conf` and `nginx.conf` are properly configured for your environment.

#### API Endpoints

The server exposes various API endpoints for client interactions, including:

- **User Registration**: Endpoint to register new users.
- **User Authentication**: Endpoint to authenticate existing users.
- **Data Storage**: Endpoints to store various types of private data.
- **Data Retrieval**: Endpoints to retrieve stored data.
- **Data Synchronization**: Endpoints to synchronize data across clients.

For detailed API specifications, refer to the API documentation (assumed to be in the `api-spec` directory).

#### Testing

Unit tests are implemented to ensure the functionality and reliability of the server. Run the tests using:

```sh
go test ./...
```

#### Contribution

Contributions are welcome! Please submit pull requests or open issues to discuss potential improvements or report bugs.

#### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

For more detailed documentation, please refer to the [README](README.md) file in the repository.
