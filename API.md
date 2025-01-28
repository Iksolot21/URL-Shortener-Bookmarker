# API Documentation

## Auth Endpoints

### Register
*   **Method:** `POST`
*   **URL:** `/auth/register`
*   **Body:**
    ```json
    {
        "username": "your_username",
        "password": "your_password",
        "email": "your_email@example.com"
    }
    ```
*   **Response (201 Created):**
    ```json
    {
        "message": "User registered successfully"
    }
    ```
 *   **Response (400 Bad Request):**
    ```json
    {
        "error": "Username is already in use or invalid input"
    }
    ```

# Login
*   **Method:** `POST`
*   **URL:** `/auth/login`
*   **Body:**
    ```json
    {
        "username": "your_username",
        "password": "your_password"
    }
    ```
*   **Response (200 OK):**
    ```json
    {
        "token": "jwt_token"
    }
    ```
