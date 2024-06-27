
# API SERVIS Go-CRUD-JWT 3 Table


## Create User

Request :
- Method : POST
- Endpoint : `localhost:3000/user`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```yaml
form-data:
  name:test
  phone:12345678
  email:test@test.dev
  username:test
  password:test

```

Response :

```json 
{
    "data": {
        "id": 3,
        "name": "test",
        "phone": "12345678",
        "email": "test@test.dev",
        "username": "test",
        "password": "test",
        "CreatedAt": "2021-07-17 20:52:28.205326"
    }
}
```

## Login

Request :
- Method : POST
- Endpoint : `localhost:3000/login`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```yaml
form-data:
  username:test
  password:test

```

Response :

```json 
{
    "message": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.9xaLm7gtDvsvJ4wgmsAvZoOZfAgN7HUAc2htDY9hT2g",
    "data": [
        {
            "id": 3,
            "name": "test",
            "phone": "12345678",
            "email": "test@test.com",
            "username": "test",
            "gender": false
        }
    ]
}
```


## Get List User

Request :
- Method : GET
- Endpoint : `localhost:3000/users`
- Header :
    - Authorization: token
    - Accept: application/json

- Params :

```yaml
Query Params:
  limit: 2
  offset: 2
  name: "test"
  
```

Response :

```json 
{
    "count": 1,
    "data": [
        {
            "id": 1,
            "name": "test",
            "phone": "12345678",
            "email": "test@test.dev",
            "username": "test",
            "gender": false
        }
    ]
}
```



## Delete User

Request :
- Method : DELETE
- Endpoint : `localhost:3000/user?id=1`
- Header :
    - Authorization: token

Response :

```json 
{
    "data": {
        "id": 2,
        "DeletedAt": "2021-07-17 20:52:28"
    }
}
```


## Update User

Request :
- Method : PUT
- Endpoint : `localhost:3000/user`
- Header :
    - Authorization: token

- Body :

```yaml
form-data:
  name:test
  phone:12345678
  email:test2@test.dev
  username:test
  password:test

```

Response :

```json 
{
    "data": {
        "id": 3,
        "name": "test",
        "phone": "12345678",
        "email": "test2@test.dev",
        "username": "test",
        "password": "test",
        "CreatedAt": "2021-07-17 20:52:28.205326"
    }
}
```

## Create, Update, Delet, View Pasien

Request :
- Method : POST, DELETE, PUT, GET
- Endpoint : `localhost:3000/pasien`
- Header :
    - Authorization: token
- Body :

```yaml
form-data:
  nama:pasien
  alamat:jl.pasien
  tanggal_lahir:2021-07-17
  nomor_telepon:12345678

```

Response :

```json 
{
    "data": {
        "id": 1,
        "nama": "pasien",
        "alamat": "jl.pasien",
        "tanggal_lahir": "test@test.dev",
        "nomor_telepon": "12345678",
        "CreatedAt": "2021-07-17 20:52:28.205326"
    }
}
```

## Create, Update, Delet, View Dokter

Request :
- Method : POST, DELETE, PUT, GET
- Endpoint : `localhost:3000/dokter`
- Header :
    - Authorization: token
- Body :

```yaml
form-data:
  nama:dokter
  spesialisasi:jantung
  nomor_telepon:12345678

```

Response :

```json 
{
    "data": [
        {
            "id": 1,
            "nama": "dokter",
            "spesialisasi": "jantung",
            "nomor_telepon": "12345678",
            "created_at": "2024-06-27T12:34:46+07:00"
        }
    ]
}
```

## Create, Update, Delet, View Kunjungan

Request :
- Method : POST, DELETE, PUT, GET
- Endpoint : `localhost:3000/dokter`
- Header :
    - Authorization: token
- Body :

```yaml
form-data:
  id_dokter:1
  id_pasien:1
  tanggal_kunjungan:2006-01-02T22:04:05+07:00
  diagnosa:sakit perut

```

Response :

```json 
{
    "data": [
        {
            "id": 1,
            "id_dokter": 1,
            "id_pasien": 1,
            "nama_dokter": "qw",
            "nama_pasien": "qw",
            "tanggal_kunjungan": "2006-01-02T22:04:05+07:00",
            "diagnosa": "qw",
            "created_at": "2024-06-27T23:51:42+07:00"
        }
    ]
}
```


## Setup Database
create database db_klinik