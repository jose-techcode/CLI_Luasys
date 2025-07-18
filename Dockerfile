FROM ubuntu:22.04

# Atualiza e instala todos os utilitários necessários da CLI tool Luasys
RUN apt update && apt install -y \
    dmidecode \
    pciutils \
    acpi \
    lm-sensors \
    usbutils \
    && apt clean

# Copia o binário
COPY luasys /usr/local/bin/luasys
RUN chmod +x /usr/local/bin/luasys

# Comando do binário
ENTRYPOINT ["/usr/local/bin/luasys"]