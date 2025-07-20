# 1. Installation

Luasys tool .deb package.

To install, download the .deb file and run:

- sudo dpkg -i luasys_1.0.25.deb

Then, run:

- luasys --help

Note:

Run the CLI tool with administrator permissions (sudo) for proper operation.

# 2. Luasys CLI (admin/sudo)

The "CLI Luasys" is a multifunctional CLI tool for auditing, diagnosing, and monitoring hardware resources, such as CPU, GPU, memory (ram and swap), motherboard, BIOS, disk and partitions, battery, temperature, USB devices, host (kernel and operating system), net (network), and processes. You need administrator or sudo permission to run the Luasys CLI. By cloning the repository, you will have control of a local REST GET API.

# 3. Features

- luasys - Displays general information, excluding net and process.

- api - Starts the API HTTP server. Localhost.

- cpu - Displays CPU information.

- gpu - Displays GPU information.

- motherboard - Displays motherboard information.

- bios - Displays BIOS information.

- mem - Displays memory information.

- disk - Displays disk information.

- battery - Displays battery information.

- temp - Displays temperature information.

- usb - Displays USB device information.

- host - Displays kernel and operating system information.

- net - Displays network information.

- processes - Displays process information.

# 4. Technologies

- Language: Golang.
- Framework: Echo.
- Libraries: Cobra.
- Environment: Linux. 
- Code versioning: Git.
- Containerization: Docker.
- Others: dmidecode, lspci, acpi, sensors, and lsusb.

# 5. API

- GET: Displays auditing, diagnostic, and monitoring information via an HTTP server via API.

- / - General path.

- /cpu - CPU path.

- /gpu - GPU path.

- /motherboard - Motherboard path.

- /bios - BIOS path.

- /mem - Memory path.

- /disk - Disk path.

- /battery - Battery path.

- /temperature - Temperature path.

- /usb - USB device path.

- /host - Kernel and operating system path.

- /net - Network path.

- /processes - Process path.

# 6. Clone the Repository

- Bash

git clone github.com/jose-techcode/CLI_Luasys

# 7. Project Folder

- Bash

cd CLI_Luasys

# 8. Run in Docker

- Build the image

docker build -t luasys:latest .

- Run the image

docker run --rm luasys:latest

- Run the image with privileges (recommended)

docker run --rm --privileged luasys:latest (optional: < subcommand >)

# 9. Contribute

Feel free to open Issues or submit Pull Requests.

# 10. License

This project is licensed under the GPL.