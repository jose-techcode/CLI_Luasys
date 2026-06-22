# 1. Installation

Sys tool .tar.gz package

- To install, download the .tar.gz, navigate to the directory where the .tar.gz file is located and run:

`tar -xzvf SysAdmin-Tool_Linux_x86_64.tar.gz`

- Move to /usr/local/bin/

`sudo mv sys /usr/local/bin/`

- Then, run:

`sys --help`

- **Note:**

Run the CLI tool with administrator permissions (sudo) for proper operation.

# 2. SysAdmin Tool

[DEPRECATED] The SysAdmin Tool is an CLI tool for auditing, diagnosing, and monitoring hardware resources, such as CPU, GPU, motherboard, BIOS, memory, disk and partitions, USB devices, host (kernel and operating system), net (network), and processes. Net and processes do not appear in the general command, they only work in subcommands. You need administrator or sudo permission to run the CLI tool.

# 3. Features

- `sys` - Displays general information, excluding net and process. Admin/Sudo.

- `cpu` - Shows CPU information.

- `gpu` - Shows GPU information.

- `motherboard` - Shows motherboard information. Admin/Sudo.

- `bios` - Shows BIOS information. Admin/Sudo.

- `mem` - Shows memory information.

- `disk` - Shows disk information.

- `usb` - Shows USB device information.

- `host` - Shows kernel and operating system information.

- `net` - Shows network information. It works only through subcommand.

- `processes` - Shows process information. It works only through subcommand.

# 4. Technologies

- Language: Go (1.25.7+)
- Libraries: Cobra (1.10.2+) & Gopsutil (3.24.5+)
- Environment: Linux/Debian(13+)/Fedora(42+)
- Code versioning: Git (2.53.0+)
- CI: Github Actions
- Others: dmidecode, pciutils, usbutils

# 5. Clone the Repository

- Bash

`git clone https://github.com/josewalter01/SysAdmin-Tool`

# 6. Project Folder

- Bash

`cd SysAdmin-Tool`

# 7. Contribute

Feel free to open Issues or submit Pull Requests.

# 8. License

This project is licensed under the GPL 3.0.
