# Proxmox API Wrapper
```markdown
 ______   ______     ______     __  __     __    __     ______     __  __     ______     ______   __    
/\  == \ /\  == \   /\  __ \   /\_\_\_\   /\ "-./  \   /\  __ \   /\_\_\_\   /\  __ \   /\  == \ /\ \   
\ \  _-/ \ \  __<   \ \ \/\ \  \/_/\_\/_  \ \ \-./\ \  \ \ \/\ \  \/_/\_\/_  \ \  __ \  \ \  _-/ \ \ \  
 \ \_\    \ \_\ \_\  \ \_____\   /\_\/\_\  \ \_\ \ \_\  \ \_____\   /\_\/\_\  \ \_\ \_\  \ \_\    \ \_\ 
  \/_/     \/_/ /_/   \/_____/   \/_/\/_/   \/_/  \/_/   \/_____/   \/_/\/_/   \/_/\/_/   \/_/     \/_/ 
                                                                                                        
```

Proxmox API Wrapper is a Golang wrapper for interacting with the Proxmox API in a simpler and more intuitive way.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you can use the Proxmox API Wrapper, you need to ensure that Go is installed on your machine. If it's not, you can download it from [here](https://golang.org/dl/).

### Installing

First, clone the repository using the following command:

```bash
git clone https://github.com/devanbenz/proxmox-api-wrapper.git
```

Within the repo, install the dependencies:

```bash
go mod download
```

### Building
To compile the application, you can use the go build command:

```bash
go build -o proxmox-api-wrapper
```

This will create a binary named proxmox-api-wrapper in the current directory.

### Usage
The main.go file is set up to make an example set of API calls to Proxmox.

You need to replace <ip>, <username>, and <password> placeholders in the file with your Proxmox server details.

After that, you can run the application:

```bash
./proxmox-api-wrapper
```
