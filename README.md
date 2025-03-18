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

`POST` / `localhost:8080/api/register`
 
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

`POST` / `localhost:8080/user/login`
  
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
    "data": "success login! Welcome murilo",
    "message": "operation from handler: login-user success"
}
```

---

### **Buscar usuário**

`GET` / `localhost:8080/api/user/:ID`

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

`GET` / `localhost:8080/api/user`

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

`DELETE` / `localhost:8080/api/user?id`
  
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

`PUT` / `localhost:8080/api/user?id`
  
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
