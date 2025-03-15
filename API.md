# FarmQuest API Documentation

Base URL: `/api`

## Authentication
### Login
- **Endpoint**: `POST /auth/login`
- **Request Body**:
  ```json
  {
    "email": "string (required, valid email)",
    "password": "string (required, min 8 chars)"
  }
  ```
- **Response**:
  ```json
  {
    "token": "string"
  }
  ```

### Register
- **Endpoint**: `POST /auth/register`
- **Request Body**:
  ```json
  {
    "username": "string (required, min 3, max 255 chars)",
    "email": "string (required, valid email)",
    "password": "string (required, min 8 chars)"
  }
  ```
- **Response**:
  ```json
  {
    "token": "string"
  }
  ```

## Crops
### List Crops
- **Endpoint**: `GET /crops`
- **Query Parameters**:
  - `page`: integer (min: 1)
  - `size`: integer (min: 1, max: 100)
  - `search`: string (optional)
- **Response**:
  ```json
  {
    "id": "integer",
    "name": "string",
    "description": "string",
    "full_image_url": "string",
    "basic_needs": ["string"],
    "tags": ["string"],
    "price": "integer",
    "rating": "integer"
  }
  ```

### Get Crop Dropdown
- **Endpoint**: `GET /crops/dropdown`
- **Response**:
  ```json
  {
    "id": "integer",
    "name": "string",
    "cropped_image_url": "string",
    "full_image_url": "string"
  }
  ```

### Get Crop Details
- **Endpoint**: `GET /crops/:id`
- **Response**: Same as List Crops response

## Farms
### List Farms
- **Endpoint**: `GET /farms`
- **Query Parameters**:
  - `page`: integer (min: 1)
  - `size`: integer (min: 1, max: 100)
  - `search`: string (optional)
- **Response**:
  ```json
  {
    "id": "integer",
    "name": "string",
    "image": "string",
    "status": "string",
    "location": "string",
    "area": "float"
  }
  ```

### Get Farm Details
- **Endpoint**: `GET /farms/:id`
- **Response**:
  ```json
  {
    "id": "integer",
    "name": "string",
    "image": "string",
    "status": "string",
    "description": "string",
    "location": "string",
    "health": "integer",
    "area": "float",
    "planted_at": "datetime",
    "growing_at": "datetime (optional)",
    "harvest_at": "datetime (optional)",
    "growth_status": {
      "description": "string",
      "image_url": "string"
    },
    "related_news": [
      {
        "title": "string",
        "description": "string",
        "image_url": "string",
        "link": "string"
      }
    ]
  }
  ```

### Create Farm
- **Endpoint**: `POST /farms`
- **Request Body**:
  ```json
  {
    "name": "string (required, max 255 chars)",
    "crop_id": "integer (required, > 0)",
    "description": "string (required, max 1000 chars)",
    "location": "string (required, max 255 chars)",
    "area": "float (required, > 0)"
  }
  ```

### Update Farm
- **Endpoint**: `PUT /farms/:id`
- **Request Body**:
  ```json
  {
    "name": "string (optional, max 255 chars)",
    "status": "string (optional)",
    "health": "integer (optional, > 0, <= 100)",
    "area": "float (optional, > 0)",
    "description": "string (optional, max 1000 chars)",
    "location": "string (optional, max 255 chars)"
  }
  ```

### Get Farm Growth Status
- **Endpoint**: `GET /farms/:id/growth`
- **Response**:
  ```json
  {
    "stages": [
      {
        "description": "string",
        "status": "string",
        "date": "datetime (optional)"
      }
    ]
  }
  ```

## Questions
### Get Questions
- **Endpoint**: `GET /questions`
- **Response**:
  ```json
  {
    "text": "string",
    "answers": [
      {
        "text": "string"
      }
    ]
  }
  ```

## Health Check
### Get Health Status
- **Endpoint**: `GET /health`
