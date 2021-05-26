# Mini Kraicklist

In this exercise we will learn how to create simple REST API using Go.

We are trying to create API for imaginary classified ads website called Kraicklist (pun of Craigslist). The idea is user could post new ads & get latest ads from it.

Please check upcoming sections for details of API design.

Your tasks:

1. Implement the server API using Go & MySQL
2. Containerize your server API using Docker
3. Create deployment template using Docker Compose which include the deployment script for your server & MySQL server

---

## Post New Ad

POST: `/ads`

This API is used for posting new ad.

**Request Body:**

- `title`, String => title of the ad
- `body`, String => body of the ad contains descriptions
- `tags`, Array of String, _optional_ => related tags for the ad if any

**Sample Request:**

```json
{
  "title": "Porsche Carrera GT 2016 for Sale",
  "body": "Clean body & great engine!",
  "tags": ["car", "porsche", "porsche 2016"]
}
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "success": true,
    "data": {
        "id": 1282323,
        "title": "Porsche Carrera GT 2016 for Sale",
        "body": "Clean body & great engine!",
        "tags": [
            "car",
            "porsche",
            "porsche 2016"
        ],
        "created_at": 1581931121
    }
}
```

**Error Responses:**

- Bad Request

  ```json
  HTTP/1.1 400 Bad Request
  Content-Type: application/json

  {
      "success": false,
      "err": "ERR_BAD_REQUEST",
      "message": "field `title` cannot be empty"
  }
  ```

[Back to Top](#mini-kraicklist)

---

## Get Latest Ads

GET: `/ads?limit=<limit>`

This API is used for getting latest ads from database.

If `limit` value is not specified, by default the value is `5`.

**Sample Request:**

```json
GET: /ads?limit=10
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "success": true,
    "data": {
        "ads": [
            {
                "id": 1282323,
                "title": "Porsche Carrera GT 2016 for Sale",
                "body": "Clean body & great engine!",
                "tags": [
                    "car",
                    "porsche",
                    "porsche 2016"
                ],
                "created_at": 1581931121,
                "updated_at": 1581931121
            }
        ]
    }
}
```

**Error Responses:**

No specific error responses

[Back to Top](#mini-kraicklist-extended)

---

## Update Ad

PUT: `/ads/<ad_id>`

This API is used for updating ad.

**Request Body:**

- `title`, String, _optional_ => updated title of the ad
- `body`, String, _optional_ => updated body of the ad
- `tags`, Array of String, _optional_ => updated related tags for the ad

**Sample Request:**

```json
PUT: /ads/1282323
Content-Type: application/json

{
    "title": "Porsche Carrera GT 2016 & 2017 for Sale!",
    "body": "Both of them have clean body & great engine!",
    "tags": [
        "car",
        "porsche",
        "porsche 2016",
        "porsche 2017"
    ]
}
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "success": true,
    "data": {
        "id": 1282323,
        "title": "Porsche Carrera GT 2016 & 2017 for Sale!",
        "body": "Both of them have clean body & great engine!",
        "tags": [
            "car",
            "porsche",
            "porsche 2016",
            "porsche 2017"
        ],
        "updated_at": 1582000607
    }
}
```

**Error Responses:**

- Bad Request

  ```json
  HTTP/1.1 400 Bad Request
  Content-Type: application/json

  {
      "success": false,
      "err": "ERR_BAD_REQUEST",
      "message": "at least one parameter must be specified"
  }
  ```

- Not Found

  ```json
  HTTP/1.1 404 Not Found
  Content-Type: application/json

  {
      "success": false,
      "err": "ERR_NOT_FOUND",
      "message": "data is not found"
  }
  ```

[Back to Top](#mini-kraicklist-extended)

---

## Delete Ad

DELETE: `/ads/<ad_id>`

This API is used for deleting ad. The ad will be literally deleted from database.

**Sample Request:**

```json
DELETE /ads/1282323
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "success": true,
    "data": {
        "id": 1282323
    }
}
```

**Error Responses:**

- Not Found

  ```json
  HTTP/1.1 404 Not Found
  Content-Type: application/json

  {
      "success": true,
      "err": "ERR_NOT_FOUND",
      "message": "data is not found"
  }
  ```

[Back to Top](#mini-kraicklist-extended)

---

## Get Ad Statistics

GET: `/stats`

This API is used for fetching ads statistic. Currently we will only show total ads in statistic.

**Sample Request:**

```json
GET /stats
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "success": true,
    "data": {
        "stats": {
            "total_ads": 1
        }
    }
}
```

**Error Responses:**

No specific error response

[Back to Top](#mini-kraicklist-extended)

---
