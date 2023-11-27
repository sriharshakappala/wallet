## Simple Wallet Service

#### Setup Environment

This app uses MYSQL for database. For convenience a docker-compose.yml file has been added in this repo for spinning up local environment quickly.

1. Navigate to the root directory of this app
2. Run `docker-compose up`
3. Once the MYSQL image is pulled and container is started execute this command `migrate -path ./migrations -database "mysql://admin:password@tcp(localhost:3306)/wallet" up`. This will run the migrations and creates the schema & seeds with initial data.

***

#### Stack
1. App is written in Go Lang
2. Schema is managed by Go Lang migrations utility. This can be installed via `brew install golang-migrate`. Sample command to create migrations - `migrate create -ext sql -dir migrations -seq create_users`
3. This uses MYSQL database
4. Dev environment & dependencies managed by docker

***

#### Schema

<img width="950" alt="Screenshot 2023-11-27 at 6 25 32 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/b5d2ac78-f565-4c76-a855-7e11c0d175f1">

#### DBML (https://dbdiagram.io/d)

```
Table users {
  id integer [primary key]
  username varchar
  created_at timestamp
  updated_at timestamp
}

Table wallet {
  id integer [primary key]
  user_id integer
  balance float
  created_at timestamp
  updated_at timestamp
}

Table transactions {
  id integer [primary key]
  user_id integer
  txn_type string
  txn_date timestamp
  txn_amount float
  closing_balance float
  other_party_id integer
  created_at timestamp
  updated_at timestamp
}

Ref: wallet.user_id - users.id

Ref: users.id < transactions.user_id

Ref: users.id < transactions.other_party_id
```



