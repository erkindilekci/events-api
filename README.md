# Events API
Events API is a web application designed to manage and serve event-related data. It leverages several modern technologies to provide a robust and efficient backend service.

## Technologies

- `Go`
- `Gin Gonic`
- `JWT`
- `SQLite`

## Setting Up Events API Locally

To run the Events API application on your local machine, follow these steps:

1. **Clone the Repository:**
   Begin by cloning the Events API repository to your local environment using Git:
   ```bash
   git clone https://github.com/erkindilekci/events-api.git
   ```

2. **Install Dependencies:**
   Navigate to the project directory and install the required dependencies:
   ```bash
   cd events-api
   go mod tidy
   ```

3. **Run the Application:**
   Start the application using the Go command:
   ```bash
   go run ./cmd/eventsapi
   ```

4. **Access the Application:**
   Open your web browser and navigate to `http://localhost:8080` to access the Events API.

**Note:** Ensure you have Go installed on your system before proceeding with these steps.