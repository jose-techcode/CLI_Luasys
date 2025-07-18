# 1. CLI Luasys (admin/sudo)

A "CLI Luasys" é uma CLI Tool de auditoria, diagnóstico e monitoramento de recursos do hardware, como cpu, memória, placa-mãe, bios, memória, disco e partições, bateria, temperatura, usb, host (kernel e sistema operacional), net (rede) e processes (processos). Precisa de permissão de administrador ou sudo para executar a CLI Luasys. Ao clonar o repositório, você terá controle de uma API REST local.

# 2. Funcionalidades

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

# 3. Tecnologias

- Linguagem: Golang
- Framework: Echo
- Bibliotecas: Cobra
- Ambiente: Linux
- Versionamento de código: Git
- Outros: dmidecode, lspci, acpi, sensors e lsusb

# 4. API

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

# 5. Clone do Repositório

- Bash

git clone github.com/jose-techcode/CLI_Luasys

# 6. Pasta do Projeto

- Bash

cd CLI_Luasys

# 7. Contribuição

Sinta-se livre para abrir Issues ou enviar Pull Requests.

# 8. Licença

Este projeto está licenciado sob a licença GPL.

# 9. Observações

Não disponível para instação via apt, nem via .deb. Documentação sujeita a mais atualizações, bem como a CLI Luasys.