# 💳 Gateway de Pagamentos


<p align="center">
  <img src="https://img.shields.io/badge/build-passing-brightgreen?style=for-the-badge" alt="Build Passing" />
  <img src="https://img.shields.io/badge/docker-ready-blue?style=for-the-badge&logo=docker&logoColor=white" />
  <img src="https://img.shields.io/badge/license-MIT-blue?style=for-the-badge" alt="License: MIT" />
  <img src="https://img.shields.io/badge/made%20with-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
</p>


**Uma solução moderna para processamento de pagamentos com análise antifraude em tempo real!**  
Desenvolvido com Go, Next.js, NestJS, Kafka e PostgreSQL. 🔥

<p align="center">
  <img src="./GatewaydePagamento/img/projeto.png" alt="Diagrama Arquitetura" width="600"/>
</p>

---

## 🚀 Descrição

Este projeto simula um fluxo real de **gateway de pagamentos**.  
Nele, um cliente pode:
- Criar contas
- Consultar saldo
- Processar pagamentos
- Passar por um serviço de **análise antifraude** baseada em eventos Kafka antes da aprovação.

O objetivo é demonstrar **boas práticas de arquitetura de microsserviços**, **comunicação assíncrona** e **segurança de transações**.

---

## 🛠️ Tecnologias Utilizadas

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/NestJS-E0234E?style=for-the-badge&logo=nestjs&logoColor=white" />
  <img src="https://img.shields.io/badge/Kafka-231F20?style=for-the-badge&logo=apachekafka&logoColor=white" />
  <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" />
  <img src="https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white" />
</p>

---

## 📦 Estrutura do Projeto

```
├── cmd/                # Main application entry point
├── internal/
│   ├── domain/         # Entidades e regras de domínio (Account)
│   ├── repository/     # Persistência de dados (PostgreSQL)
│   ├── service/        # Regras de negócio (AccountService)
│   ├── web/
│       ├── handlers/   # Camada HTTP
│       ├── server/     # Inicialização do servidor web
├── migrations/         # Scripts SQL para o banco
├── docker-compose.yml  # Orquestração de containers
├── README.md
```

---

## 🔄 Fluxo de Funcionamento

<p align="center">
  <img src="./GatewaydePagamento/img/fluoxograma (2).png" alt="Fluxo Pagamento" width="600"/>
</p>

1. **Usuário** envia uma requisição para criar ou consultar uma conta via API.
2. **Gateway API** gerencia a conta e publica eventos no Kafka.
3. **Serviço de Antifraude** consome e analisa os eventos.
4. **Decisão de fraude** retorna via Kafka para o Gateway.
5. **Usuário** recebe resposta do status da transação.

---

## 🐳 Docker Compose

Suba todo o ambiente local com apenas:

```bash
docker-compose up -d
```

Isso iniciará:
- PostgreSQL
- Kafka (em configuração separada)

*Scripts de migração SQL incluídos em `/migrations` para criar as tabelas necessárias.*

---

## 🧪 Teste a API

Exemplos de testes com client HTTP:

### Criar Conta

```http
POST http://localhost:8080/accounts
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@doe.com"
}
```

### Consultar Conta

```http
GET http://localhost:8080/accounts
X-API-Key: <sua-api-key-recebida-na-criação>
```

*(Arquivos de testes automáticos incluídos.)*

---

## 🛡️ Segurança

- Todas as APIs utilizam **autenticação via X-API-Key** nos headers.
- Controle de concorrência nas atualizações de saldo (`SELECT FOR UPDATE`).
- Manipulação de erros centralizada e padronizada.

---

## ✅ Status do Projeto

| Componente | Status |
|:-----------|:------:|
| Frontend (Next.js) | ✅ Finalizado |
| Gateway API (Go) | ✅ Finalizado |
| Apache Kafka | ✅ Configurado |
| Antifraude (Nest.js) | ✅ Finalizado |
| Docker Compose | ✅ Finalizado |
| Migrations SQL | ✅ Finalizado |
| Testes de API | ✅ Finalizado |

---

## 📄 Licença

Distribuído sob a licença **MIT**.  
Veja o arquivo [LICENSE](LICENSE) para mais informações.

---

✨ Desenvolvido por
Feito com 💙 por Viviane Aguiar

<p align="center"> <a href="https://www.linkedin.com/in/vivianezzt/" target="_blank"> <img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" alt="LinkedIn" /> </a> &nbsp; <a href="https://www.instagram.com/vivianezzt/" target="_blank"> <img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white" alt="Instagram" /> </a> </p>

---
