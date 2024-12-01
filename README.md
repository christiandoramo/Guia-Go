## Instalação do GO (Ubuntu 24)

- Baixar go no site
- Extrair de "go1.23.3.linux-amd64.tar.gz" pasta go para /usr/local
- Configurar novo PATH: export PATH=$PATH:/usr/local/go/bin
- Testar instalação com: go version
- Rodar "echo $PATH" no vs code para ler path correto e copiar a saída
- Colar em settings "terminal.integrated.env.linux" adicionando o path no vs code
- Resetar vs code

## Iniciando projeto

- criar modulo - go mod init example/hello
- rodar código go - go run .
- go mod tidy atualiza as dependencias/pacotes go.mod e o historico de dependencias/autenticação go.sum