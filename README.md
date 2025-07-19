# 1. Instalação

Pacote .deb da ferramenta Luasys.

Para instalar, baixe o arquivo .deb e rode:

- sudo dpkg -i luasys_1.0.2.deb

Depois, execute:

- luasys --help

Observação:

Execute a CLI tool com permissão de administrador (sudo) para funcionar corretamente.

# 2. CLI Luasys (admin/sudo)

A "CLI Luasys" é uma CLI tool multifuncional de auditoria, diagnóstico e monitoramento de recursos do hardware, como cpu, memória, placa-mãe, bios, memória, disco e partições, bateria, temperatura, usb, host (kernel e sistema operacional), net (rede) e processes (processos). Precisa de permissão de administrador ou sudo para executar a CLI Luasys. Ao clonar o repositório, você terá controle de uma API REST local.

# 3. Funcionalidades

- luasys - Exibe informações gerais, com exceção de net e process.

- api - Inicia servidor HTTP da API. Localhost.

- cpu - Exibe informações da CPU.

- gpu - Exibe informações da GPU.

- motherboard - Exibe informações da placa-mãe.

- bios - Exibe informações da BIOS.

- mem - Exibe informações da memória.

- disk - Exibe informações do disco.

- battery - Exibe informações da bateria.

- temp - Exibe informações da temperatura.

- usb - Exibe informações de dispositivos USB.

- host - Exibe informações do kernel e do sistema operacional.

- net - Exibe informações da rede.

- processes - Exibe informações de processos.

# 4. Tecnologias

- Linguagem: Golang.
- Framework: Echo.
- Bibliotecas: Cobra.
- Ambiente: Linux.
- Versionamento de código: Git.
- Containerização: Docker.
- Outros: dmidecode, lspci, acpi, sensors e lsusb.

# 5. API

- GET: Exibe informações da auditoria, diagnóstico e monitoramento via servidor HTTP por API.

- / - Rota geral.

- /cpu - Rota da CPU.

- /gpu - Rota da GPU.

- /motherboard - Rota da placa-mãe.

- /bios - Rota da BIOS.

- /mem - Rota da memória.

- /disk - Rota do disco.

- /battery - Rota da bateria.

- /temperature - Rota da temperature.

- /usb - Rota de dispositivos USB.

- /host - Rota do kernel e do sistema operacional.

- /net - Rota da rede.

- /processes - Rota de processos.

# 6. Clone do Repositório

- Bash

git clone github.com/jose-techcode/CLI_Luasys

# 7. Pasta do Projeto

- Bash

cd CLI_Luasys

# 8. Rodar em Docker

- Build da imagem

docker build -t luasys:latest.

- Rodar a imagem

docker run --rm luasys:latest

- Rodar a imagem com privilégio (recomendado)

docker run --rm --privileged luasys:latest (opcional: < subcomando >)

# 9. Contribuição

Sinta-se livre para abrir Issues ou enviar Pull Requests.

# 10. Licença

Este projeto está licenciado sob a licença GPL.