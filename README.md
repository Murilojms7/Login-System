## **Documentação do Endpoint: Go Login**

- User
    - [Login usuário](#login-usuário)
    - [Registrar usuário](#registrar-usuário)
    - [Buscar usuário](#buscar-usuário)
    - [Listar usuários](#listar-usuários)
    - [Deletar usuário](#deletar-usuário)
    - [Atualizar usuário](#atualizar-usuário)
---

### **Registrar usuário**

`POST` / `localhost:8000/api/register`
 
Este endpoint permite que um usuário seja registrado.

**Corpo da requisição**

|Campo|Tipo|Obrigatório|Descrição
|-----|----|:-----------:|---------
|Name|String|✅|Nome do usuário
|Email|String|✅|Email do usuário
|Password|String|✅|Senha do usuário
|Phone|Integer|✅|telefone do usuário
|Birth|String|✅|data de nascimento do usuário

```js
{   
    "name": "murilo",
    "email": "murilo@gmail.com",
    "password": "123",
    "phone": 119581231231,
    "Birth": "19/08/2004"
}
```

**Exemplo de Resposta**

```js
{
    "data": "user: murilo@gmail.com created successfully",
    "message": "operation from handler: createUser success"
}
```

---

### **login usuário**

`POST` / `localhost:8000/user/login`
  
Este endpoint permite que um usuário consiga fazer login na aplicação.

**Corpo da requisição**

|Campo|Tipo|Obrigatório|Descrição
|-----|----|:-----------:|---------
|Email|String|✅|Email do usuário
|Password|String|✅|Senha do usuário

```js
{
    "email": "Pedro@gmail.com",
    "password": "Pedro12345",
}
```

**Exemplo de Resposta**

```js
{
    "status": "SUCCESS",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI3NzJkNWFmNy1lYTc5LTQ4MjktOTBkNS0xODAyOTU2YWE2ODciLCJ0ZWFtSWQiOiI2NzM0Nzc0Ny1jZmUxLTRlN2QtOWYxNi04ZDBiZmFjZTZjZmQiLCJhY2Nlc3NMZXZlbCI6MiwiaWF0IjoxNzM5Nzk5OTAxLCJleHAiOjE3Mzk4MzIzMDF9.x_FcTVQRdfFqFegsnOCeFJKfsQcUw-kVZz-RwYFYQvc",
    "user": {
        "name": "Pedro",
        "email": "Pedro@gmail.com",
        "accessLevel": 2,
        "teamId": "67347747-cfe1-4e7d-9f16-8d0bface6cfd",
        "teamName": "Itaú"
    }
}
```

---

### **Buscar usuário**

`GET` / `localhost:8000/api/user/:ID`

Este endpoint permite que seja buscado usuário pelo ID.

**Parâmetros da requisição**
|Parâmetro|Obrigatório|Descrição
|-----|:-----------:|---------
|ID|✅|Id do usuário

**Exemplo de Resposta**

```js
{
    "data": {
        "birth": "19/08/2004",
        "email": "murilo@gmail.com",
        "id": "294be115-7f74-4258-a364-bd2c96c85b44",
        "name": "murilo",
        "phone": 119581231231
    },
    "message": "operation from handler: show-user success"
}
```

---

### **Listar usuários**

`GET` / `localhost:8000/api/user`

Este endpoint permite que seja listado todos os usuários.

**Exemplo de Resposta**

```js
{
    "data": [
        {
            "birth": "19/08/2004",
            "email": "murilo@gmail.com",
            "id": "294be115-7f74-4258-a364-bd2c96c85b44",
            "name": "henry",
            "phone": 0
        },
        {
            "birth": "19/08/2004",
            "email": "murilo@gmail.com",
            "id": "b11cc8d5-6056-44fc-a380-b2edda220474",
            "name": "murilo",
            "phone": 119581231231
        }
    ],
    "message": "operation from handler: list-users success"
}
```

---

### **Deletar usuário**

`DELETE` / `localhost:8000/api/user?id`
  
Este endpoint permite que um usuário seja removido da aplicação.

**Parâmetros da requisição**
|Parâmetro|Obrigatório|Descrição
|-----|:-----------:|---------
|ID|✅|Id do usuário

**Exemplo de Resposta**

```js
{
    "data": {
        "Birth": "19/08/2004",
        "email": "murilo@gmail.com",
        "name": "henry",
        "phone": 0
    },
    "message": "operation from handler: delete-user success"
}
```

---

### **Atualizar usuário**

`PUT` / `localhost:8000/api/user?id`
  
Este endpoint permite que seja atualizado um usuário na aplicação.

**Corpo da requisição**

|Campo|Tipo|Obrigatório|Descrição
|-----|----|:-----------:|---------
|Name|String|🚫|Nome do usuário
|Email|String|🚫|Email do usuário
|Password|String|🚫|Senha do usuário
|Phone|Integer|🚫|telefone do usuário
|Birth|String|🚫|data de nascimento do usuário

```js
{   
    "name": "murilo",
    "email": "murilo@gmail.com",
    "password": "123",
    "phone": 119581231231,
    "Birth": "19/08/2004"
}
```

**Exemplo de Resposta**

```js
{
    "data": {
        "birth": "19/08/2004",
        "email": "murilo@gmail.com",
        "id": "294be115-7f74-4258-a364-bd2c96c85b44",
        "name": "murilo",
        "phone": 119581231231
    },
    "message": "operation from handler: show-user success"
}
```