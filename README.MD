### How to use:

**Run docker:**

```make docker-run```

**LoginEndpoint:**

```
curl --location --request POST 'http://127.0.0.1:10000/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
"login": "test@example.com",
"password": "12345"
}'
```
### IMPORTANT:

> For balance manipulations, use an integer value (US cents).


**Balance endpoint:**

```
curl --location --request GET 'http://127.0.0.1:10000/user/balance' \
--header 'Authorization: Bearer <jwtToken>'
```

**Transaction up endpoint:**

```
curl --location --request POST 'http://127.0.0.1:10000/transaction/up' \
--header 'Authorization: Bearer <jwtToken>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 10
}'
```

**Transaction down endpoint:**

```
curl --location --request POST 'http://127.0.0.1:10000/transaction/down' \
--header 'Authorization: Bearer <jwtToken>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 10
}'
```
