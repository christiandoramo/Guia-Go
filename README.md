guias iniciais: https://go.dev/doc/tutorial/

## Instalação do GO (Ubuntu 24)

- Baixar go no site
- Extrair de "go1.23.3.linux-amd64.tar.gz" pasta go para /usr/local
- Configurar novo PATH: export PATH=$PATH:/usr/local/go/bin
- Testar instalação com: go version
- Rodar "echo $PATH" no vs code para ler path correto e copiar a saída
- Colar em settings "terminal.integrated.env.linux" adicionando o path no vs code
- Resetar vs code

## Iniciando projeto 1

- criar modulo - go mod init example/hello
- rodar código go - go run .
- go get "caminho" pega pacote online
- go mod tidy atualiza as dependencias/pacotes go.mod e o historico de dependencias/autenticação go.sum

## Iniciando projeto 2

- go mod init example/greetings
- sempre que criar novo modulo editar o place para não dar problema em produção:go mod edit -replace example.com/greetings=../greetings
- go test -v mostra testes - arquivos com "\_test.go" no final e funções iniciadas com Test
- go build gera um executável (rodar com ./NomeDoArquivo)
- configurar o install path para buildar num local cohenecido: go list -f '{{.Target}}' vê onde deve instalar
- export PATH=$PATH:/home/christiandoramo/go/bin/hello salva o path de instalação
- depois de exportar o path do build para o local, instalar o build na msm pasta com gi install
- após isso, ao digitar o nome do arquivo em qualquer parte do terminal executará o programa instalado. ex: hello

## Iniciando projeto 3

- usar na pasta fora de hello - go work init ./hello (inicia um workspace)
- com isso pode rodar o modulo inteiro com - go run ./hello
- go work use ./example/hello (adiciona ao workspace o codigo de ./example.hello importado)
