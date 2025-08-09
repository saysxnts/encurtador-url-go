# Encurtador de URL em Go (Golang)

![Status](https://img.shields.io/badge/status-conclu%C3%ADdo-brightgreen)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

## 📖 Sobre o Projeto

Este é um microsserviço de backend desenvolvido inteiramente em Go que funciona como um encurtador de URL. O projeto foi construído com foco em simplicidade e performance, utilizando apenas a biblioteca padrão do Go (`net/http`) para criar um servidor HTTP robusto e eficiente.

---

## ✨ Funcionalidades

- **Endpoint `POST /encurtar`**: Recebe uma URL longa em formato JSON e retorna uma URL curta e única.
- **Endpoint `GET /{codigoCurto}`**: Recebe um código curto na URL e redireciona o usuário para a URL longa original.
- **Armazenamento em Memória**: Utiliza mapas (`map`) para armazenar a associação entre as URLs curtas e longas, garantindo respostas de alta velocidade.
- **Prevenção de Duplicatas**: Verifica se uma URL longa já foi encurtada para retornar o mesmo código curto, economizando recursos.

---

## 🛠️ Tecnologias e Conceitos

- **Go (Golang)**: Linguagem principal do projeto.
- **Biblioteca Padrão `net/http`**: Utilizada para criar o servidor web, registrar rotas e lidar com requisições e respostas HTTP.
- **Biblioteca Padrão `encoding/json`**: Para codificar e decodificar dados no formato JSON.
- **Rotas e Handlers**: Estruturação do código para lidar com diferentes endpoints e métodos HTTP.

---

## 🚀 Como Rodar o Projeto

Siga os passos abaixo para executar a aplicação localmente.

### **Pré-requisitos**

- **Go**: É necessário ter o Go (versão 1.18 ou superior) instalado. Você pode baixá-lo em [go.dev](https://go.dev/).

### **Instalação e Execução**

1. Clone o repositório:
   ```bash
   git clone [https://github.com/seu-usuario/seu-repositorio.git](https://github.com/seu-usuario/seu-repositorio.git)
   ```

2. Navegue até a pasta do projeto:
   ```bash
   cd encurtador-url-go
   ```

3. Inicie o servidor:
   ```bash
   go run main.go
   ```
   *O servidor estará rodando em `http://localhost:8080`.*

---

### **Como Testar a API**

Você pode usar uma ferramenta como o `curl` para testar os endpoints.

**1. Para encurtar uma URL (POST)**

Crie um arquivo `payload.json` na raiz do projeto com o seguinte conteúdo:
```json
{
    "url": "[https://github.com/google/go-github](https://github.com/google/go-github)"
}
```

Em seguida, execute no seu terminal:
```bash
curl -X POST -H "Content-Type: application/json" -d '@payload.json' http://localhost:8080/encurtar
```
A resposta será: `{"short_url":"http://localhost:8080/xxxxxx"}`

**2. Para redirecionar (GET)**

Copie a URL curta recebida na resposta acima e cole na barra de endereço do seu navegador. Você será redirecionado para a URL original.

---

## ✒️ Autor

**Guilherme**
