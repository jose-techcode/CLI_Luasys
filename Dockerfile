FROM ubuntu:22.04

# Update and install every utility of CLI tool Luasys

RUN apt update && apt install -y \
    dmidecode \
    pciutils \
    acpi \
    lm-sensors \
    usbutils \
    && apt clean

# Copy the binary

COPY luasys /usr/local/bin/luasys
RUN chmod +x /usr/local/bin/luasys

# Command of binary

ENTRYPOINT ["/usr/local/bin/luasys"]