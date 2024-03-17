# Go Template

Simple golang project template

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)

## Installation

1. Make sure you have Golang installed. If not, you can download it from [Golang Official Website](https://go.dev/doc/install).

2. Install 'make' if not already installed. 

    * On Debian/Ubuntu, you can use

    ```bash
    sudo apt-get update
    sudo apt-get install make
    ```

   * On macOS, you can use [Homebrew](https://brew.sh/)

    ```bash
    brew install make
    ```

   * On Windows, you can use [Chocolatey](https://chocolatey.org/)

    ```bash
    choco install make
    ```

3. Clone the repository

    ```bash
    git clone https://github.com/wildanfaz/go-template.git
    ```

4. Change to the project directory

    ```bash
    cd go-template
    ```

## Usage

1. Start the application using docker

    ```bash
    docker-compose up
    ```

## Commands

1. Install all dependencies
    ```bash
    make install
    ```

2. Start the application without docker
    ```bash
    make start
    ```