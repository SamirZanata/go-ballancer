# Load Balancer Round Robin em Go

Um balanceador de carga Round Robin implementado em Go, demonstrando conceitos avançados de concorrência, proxy reverso e distribuição de requisições HTTP.

## Características

- **Round Robin**: Distribuição sequencial e equilibrada de requisições
- **Thread-Safe**: Uso de operações atômicas para segurança em ambientes concorrentes
- **Alta Performance**: Complexidade O(1) para seleção de servidor
- **Proxy Reverso**: Utiliza `net/http/httputil` para encaminhamento eficiente
- **Concorrência Nativa**: Aproveitamento de Goroutines para alta escalabilidade

## Arquitetura

```
Cliente → Load Balancer (Porta 8080) → Servidores Upstream
                                    ├─→ localhost:8001
                                    ├─→ localhost:8002
                                    └─→ localhost:8003
```

## Como Usar

### Pré-requisitos

- Go 1.21 ou superior

### Instalação

```bash
# Clone o repositório
git clone <seu-repositorio>
cd load-balancer

# Instale as dependências
go mod download
```

### Execução

1. **Inicie os servidores upstream** (em terminais separados):

```bash
# Servidor 1
go run examples/server.go -port 8001

# Servidor 2
go run examples/server.go -port 8002

# Servidor 3
go run examples/server.go -port 8003
```

2. **Inicie o Load Balancer**:

```bash
go run load_balancer.go
```

3. **Teste o balanceador**:

```bash
# Faça múltiplas requisições
curl http://localhost:8080
curl http://localhost:8080
curl http://localhost:8080
```

## Justificativa Técnica

### Concorrência (Go)
O uso de **Goroutines** e **Channels** (em uma implementação completa com Health Check) é a principal vantagem de Go. Goroutines são threads leves gerenciadas pelo runtime de Go, permitindo que o Load Balancer lide com dezenas de milhares de conexões simultâneas com baixo consumo de memória, uma característica ideal para infraestrutura.

### Segurança Concorrente
O campo `Current` (índice de Round Robin) é acessado e modificado por múltiplas goroutines concorrentemente. Usar o pacote `sync/atomic` garante que a operação de incremento seja segura (thread-safe) e eficiente, prevenindo race conditions sem o overhead de mutexes mais pesados.

### Performance
A complexidade da seleção do próximo servidor no Round Robin é **O(1)**. A biblioteca nativa `net/http/httputil` para Proxy Reverso é altamente otimizada, garantindo baixa latência no encaminhamento da requisição.

## Estrutura do Projeto

```
.
├── load_balancer.go
├── go.mod
├── README.md
└── examples/
    └── server.go
```

## Configuração

Para alterar os servidores upstream, edite o array `serverURLs` na função `main()`:

```go
serverURLs := []string{
    "http://localhost:8001",
    "http://localhost:8002",
    "http://localhost:8003",
}
```

## Exemplo de Uso

```bash
# Terminal 1: Servidor 1
$ go run examples/server.go -port 8001
Servidor iniciado na porta 8001

# Terminal 2: Servidor 2
$ go run examples/server.go -port 8002
Servidor iniciado na porta 8002

# Terminal 3: Servidor 3
$ go run examples/server.go -port 8003
Servidor iniciado na porta 8003

# Terminal 4: Load Balancer
$ go run load_balancer.go
Load Balancer iniciado na porta 8080. Servidores: [http://localhost:8001 http://localhost:8002 http://localhost:8003]

# Terminal 5: Testes
$ curl http://localhost:8080
Resposta do servidor: 8001

$ curl http://localhost:8080
Resposta do servidor: 8002

$ curl http://localhost:8080
Resposta do servidor: 8003
```

## Tecnologias Utilizadas

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)

- **Go 1.21+**: Linguagem de programação
- **net/http**: Servidor HTTP nativo
- **net/http/httputil**: Proxy reverso
- **sync/atomic**: Operações atômicas thread-safe

## Licença

Este projeto é de código aberto e está disponível para uso educacional.

## Autor

**Samir Zanata Jr.**

📧 Email: [samirzanata@icloud.com](mailto:samirzanata@icloud.com)

Desenvolvido como parte do portfólio técnico, demonstrando conhecimento em:
- Programação concorrente em Go
- Arquitetura de sistemas distribuídos
- Balanceamento de carga
- Proxy reverso

