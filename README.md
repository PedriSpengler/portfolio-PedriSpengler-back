# Envia Email - Backend em Go 📧 (Gomailer)

Este é um backend simples desenvolvido em Go que envia e-mails utilizando o pacote gomail. O projeto foi desenvolvido para demonstrar como integrar o envio de e-mails com um serviço web, utilizando a API REST. Ele também implementa a verificação de CAPTCHA utilizando o Google reCAPTCHA para garantir que a solicitação seja feita por um humano 🤖.

## Funcionalidades ✨

- Envia e-mails através de um servidor SMTP utilizando o pacote gomail.
- Verifica se os campos obrigatórios foram preenchidos na solicitação ✅.
- Valida o reCAPTCHA para garantir que o envio de e-mails seja feito por um humano 🛡️.
- Permite configurar o servidor SMTP e o reCAPTCHA através de variáveis de ambiente ⚙️.
- Protege o servidor com CORS, permitindo acesso somente de origens autorizadas 🌐.

## Estrutura do Projeto 📂

```
    envia-email/
    ├── internal/
    │ └── handler/
    │ ├── email_handler.go # Manipulador da API de envio de e-mail
    │ └── utils/
    │ ├── email_utils.go # Funções auxiliares para envio de e-mail e validação de CAPTCHA
    ├── main.go # Arquivo principal para inicializar o servidor
    ├── go.mod # Arquivo de dependências
    ├── .env # Arquivo para definir variáveis de ambiente
    └── README.md # Documentação do projeto
```

## Como Rodar o Projeto 🚀

### Pré-requisitos 📋
- Go 1.21 ou superior 🛠️
- Dependências do projeto, que serão baixadas automaticamente via go mod 📥.
- Conta de e-mail SMTP configurada (por exemplo, Gmail) 📧.
- Chave secreta do reCAPTCHA para validar solicitações de envio 🔑.

### Instalação 🛠️
1. Clone o repositório para o seu ambiente local:

   
```
   git clone https://github.com/usuario/envia-email.git
   cd envia-email
```
Instale as dependências:
go mod tidy
Crie um arquivo .env na raiz do projeto e adicione as seguintes variáveis de ambiente:
```
    SMTP_EMAIL=seu-email@gmail.com
    SMTP_PASSWORD=sua-senha-do-email
    RECAPTCHA_SECRET=sua-chave-secreta-recaptcha
    PORT=8080
```
Execute o servidor:

go run main.go
O servidor estará rodando em http://localhost:8080 ou na porta definida na variável de ambiente PORT.

## Endpoints da API 💻
POST /api/send-email
Este endpoint permite enviar um e-mail para o endereço configurado no servidor SMTP.

Request Body:
```
{
    "email": "seu-email@dominio.com",
    "subject": "Assunto do e-mail",
    "name": "Seu Nome",
    "message": "Mensagem que será enviada no corpo do e-mail",
    "captcha": "Resposta do reCAPTCHA"
}
```
Campos obrigatórios:

email: Endereço de e-mail do remetente ✉️.

subject: Assunto do e-mail 📝.

name: Nome do remetente 🧑‍💻.

message: Corpo do e-mail 📨.

captcha: Resposta do reCAPTCHA para validação de envio 🔒.

Resposta:

Em caso de sucesso, a resposta será:
```
{
    "message": "E-mail enviado com sucesso!"
}
```

Em caso de erro, a resposta será um código de erro com uma mensagem explicativa, por exemplo:
```
{
    "error": "Campos email, assunto, nome e mensagem são obrigatórios"
}
```
## Dependências 📦
github.com/joho/godotenv: Carregamento de variáveis de ambiente 🌍.

github.com/rs/cors: Configuração de CORS para permitir que apenas origens específicas possam acessar o serviço 🔐.

gopkg.in/gomail.v2: Pacote para envio de e-mails utilizando SMTP 📧.

gopkg.in/alexcesaro/quotedprintable.v3: Codificação de mensagens em formato Quoted-Printable (indireto) 🔠.

## Como Funciona 🔄
O usuário faz uma requisição POST para o endpoint /api/send-email com as informações necessárias 📥.

O servidor valida os campos obrigatórios ✅.

A resposta do reCAPTCHA é verificada através de uma chamada à API do Google 🌍.

Se os dados estiverem corretos e o reCAPTCHA for validado, o servidor envia o e-mail utilizando o pacote gomail 📧.

O servidor retorna uma resposta de sucesso ou erro, dependendo do resultado do envio 🎯.

## Configuração do Gmail 📧
Se você estiver usando o Gmail para enviar os e-mails, siga as instruções abaixo para permitir o envio de e-mails via SMTP:

Habilite o acesso a apps menos seguros ou configure uma senha de app se a autenticação de dois fatores estiver habilitada 🔐.

Nota: Para produção, recomenda-se o uso de uma conta dedicada e a configuração de uma senha de app para maior segurança 🔒.

## Contribuições 🤝
Contribuições são bem-vindas! Se você tiver sugestões ou encontrar problemas, sinta-se à vontade para abrir uma issue ou enviar um pull request 💡.

## Licença 📄
Este projeto está licenciado sob a MIT License - veja o arquivo LICENSE para mais informações.
# portfolio-PedriSpengler-back
# portfolio-PedriSpengler-back
