# REST API server (Go, gin, gorm) 

This project contains rest api server built with Go

---
## Getting started
1. download vendors: _go mod vendor_
2. edit "app.env" file (server's port, use tls or not, DB configuration, etc)
2. build docker image using script in ./deployments/
> #### cd ./deployments
> #### sh deploy.sh
3. start container using script in
> #### sh restart.sh
4. send requests

### API:

---

#### GET USER
[GET] {{server_host}}/api/v1/user/{{users_uuid}}

##### Response (200):
```
{
    "created_at": "2022-10-05T18:53:40.351205804Z",
    "deleted_at": null,
    "id": 4,
    "login": "ledorubik",
    "updated_at": "2022-10-05T18:53:40.351205804Z",
    "uuid": "859100a8-0a2f-48c2-9e4f-74e4258b5034"
}
```

#### CREATE USER
[POST] {{server_host}}/api/v1/user
```
{
    "login": "ledorubik"
}
```

##### Response (201):
```
{
    "created_at": "2022-10-05T18:53:40.351205804Z",
    "deleted_at": null,
    "id": 4,
    "login": "ledorubik",
    "updated_at": "2022-10-05T18:53:40.351205804Z",
    "uuid": "859100a8-0a2f-48c2-9e4f-74e4258b5034"
}
```

#### UPDATE USER
[PUT] {{server_host}}/api/v1/user/{{users_uuid}}
```
{
    "login": "super_ledorubik"
}
```

##### Response (200):
```
{
    "login": "super_ledorubik",
    "uuid": "859100a8-0a2f-48c2-9e4f-74e4258b5034"
}
```

#### DELETE USER
[DELETE] {{server_host}}/api/v1/user/{{users_uuid}}
##### Response (200):
```
{}
```