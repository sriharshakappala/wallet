## Simple Wallet Service

#### Setup Environment

This app uses MYSQL for database. For convenience a docker-compose.yml file has been added in this repo for spinning up local environment quickly.

1. Navigate to the root directory of this app
2. Run `docker-compose up`
3. Once the MYSQL image is pulled and container is started execute this command `migrate -path ./migrations -database "mysql://admin:password@tcp(localhost:3306)/wallet" up`. This will run the migrations and creates the schema & seeds with initial data.

<img width="806" alt="Screenshot 2023-11-27 at 8 55 21 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/086f6b6a-081c-4028-9b2e-651e1886e5f1">

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

#### Execution

##### User creation

<img width="409" alt="Screenshot 2023-11-27 at 8 56 51 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/85b07df6-ba91-4adc-85a8-366811e1ae59">

<img width="448" alt="Screenshot 2023-11-27 at 8 57 51 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/fd04aca0-1ee1-4bdd-b2db-ed94922b7033">

<img width="490" alt="Screenshot 2023-11-27 at 8 57 56 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/9c84b7f6-c1be-4e1d-8de3-0999ad68eadd">

##### View Balances

<img width="431" alt="Screenshot 2023-11-27 at 8 59 50 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/486b37d9-01f5-43cc-a7a3-72feac82a433">

##### Transaction

<img width="372" alt="Screenshot 2023-11-27 at 9 54 26 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/ecaaee24-9370-47c1-ae2c-824db35578bd">

<img width="970" alt="Screenshot 2023-11-27 at 9 54 45 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/2196e944-67a3-4f75-bc8c-bdc2baf41c1b">

<img width="532" alt="Screenshot 2023-11-27 at 9 54 53 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/79521334-a120-4568-b224-394e57082e8d">

##### Refund

<img width="384" alt="Screenshot 2023-11-27 at 10 00 21 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/054a2200-c772-4806-afa4-949df3576bb5">

<img width="1018" alt="Screenshot 2023-11-27 at 10 01 18 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/7569196a-fec8-4e25-a64a-181cca13e0f6">

<img width="531" alt="Screenshot 2023-11-27 at 10 01 25 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/8c3ca38d-3ebe-4bae-9d6f-f80803130cd3">

##### View Statement

<img width="408" alt="Screenshot 2023-11-27 at 10 28 41 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/f2ce32e9-59ab-403f-b0bd-9afeaae09327">

<img width="358" alt="Screenshot 2023-11-27 at 10 28 51 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/8acc0230-2a4c-486c-bbdb-3ea4d47d686d">

<img width="1013" alt="Screenshot 2023-11-27 at 10 32 38 PM" src="https://github.com/sriharshakappala/wallet/assets/3955701/768e50ee-f843-44ce-bd9b-bcfdcc823dfc">
