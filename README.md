# Carteira-Banco Digital

Esse projeto é uma implementação de um desafio para uma vaga de backend, o objetivo do desafio é implementar um serviço de carteira digital onde os usuários poderão transferir valores, realizar pagamentos em estabelecimentos, optei por desenvolver um frontend para praticar outras skills utilizando o mesmo projeto.

### 🕵️‍♂️ Link para o frontend:

### 🧾 Diagrama de sequência do fluxo de transação:

[Fluxo de transação](./api/docs/Diagrama%20fluxo%20de%20transação.pdf)

### Requisitos:

- ~~Para ambos tipos de usuário, precisamos do Nome Completo, CPF, e-mail e Senha. CPF/CNPJ e e-mails devem ser únicos no sistema. Sendo assim, seu sistema deve permitir apenas um cadastro com o mesmo CPF ou endereço de e-mail.~~

- ~~Usuários podem enviar dinheiro (efetuar transferência) para lojistas e entre usuários.~~

- ~~Lojistas só recebem transferências, não enviam dinheiro para ninguém.~~

- ~~Validar se o usuário tem saldo antes da transferência.~~

- ~~Antes de finalizar a transferência, deve-se consultar um serviço autorizador externo, use este mock para simular (https://run.mocky.io/v3/5794d450-d2e2-4412-8131-73d0293ac1cc).~~

- ~~A operação de transferência deve ser uma transação (ou seja, revertida em qualquer caso de inconsistência) e o dinheiro deve voltar para a carteira do usuário que envia.~~

- ~~No recebimento de pagamento, o usuário ou lojista precisa receber notificação (envio de email, sms) enviada por um serviço de terceiro e eventualmente este serviço pode estar indisponível/instável. Use este mock para simular o envio (https://run.mocky.io/v3/54dc2cf1-3add-45b5-b5a9-6bf7e7f1f4a6).~~

- ~~Este serviço deve ser RESTFul.~~

### 🛠 Tecnologias utilizadas:

###  Frontend:
- [Next JS](https://nextjs.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Redux Toolkit](https://redux-toolkit.js.org/)
- [React-i18next](https://react.i18next.com/)

### Backend:
 - [Go](https://go.dev/)
 - [Gin](https://gin-gonic.com)
 - [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
 - [JWT-Go](https://github.com/dgrijalva/jwt-go)