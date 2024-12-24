# CipherKit 🚀

Welcome to **CipherKit**! This powerful library is designed to enhance your Go applications with high-performance features and easy-to-use functionalities. Whether you're a seasoned Go developer or just starting out, this library has something for everyone. 

## Features ✨
- **Written in Go** for optimal speed and performance.
- Supports multiple formats: **base64**, **SHA256**, **RSA**, **AES256**, **hex**, and **binary**.
- Comes with convenient helper methods to simplify development and data type conversions.
- Includes a **CLI tool** for enhanced efficiency, perfect for those who may not be familiar with Go or wish to streamline daily tasks.

## Installation 📦

To get started with **My Go Library**, you need to install it using Go's package manager. Run the following command in your terminal:
```bash
go get github.com/mmdali-dev/cipherkit 

```
or use this for install CipherKit CLI for easy use in os:
```bash
go install github.com/mmdali-dev/cipherkit 
```
Ensure you have Go installed and properly set up on your machine before proceeding.

## Usage 🛠️

After the installation, you can easily integrate the library into your project. Here’s a basic usage example:
```golang
package main

import (
    "fmt"
    "github.com/mmdali-dev/cipherkit/base64"
)

func main() {
    // Example usage of the library
    data := []byte("Hello, World!")
    encoded_text := base64.Encode(data)
    println(encoded_text)
}

```
You can explore various methods provided for encoding, hashing, and encryption in the documentation.

## CLI Tool 🖥️

For those who prefer a command-line interface, **My Go Library** includes a CLI tool to assist you. Here’s how to use it:

1. **Install the CLI**:
   Follow the same installation procedure as above.

2. **Run the CLI command**:
   You can run the CLI with the following syntax:
```bash
      cipherkit <command> [options]
   
```
   For example, to encode a string in Base64:

      cipherkit -mod base64 -text "Hello, World!"
   

3. **Available mods**:
   - base64: Encode/Decode data in Base64.
   - sha256: Generate SHA256 hashes.
   - rsa: Perform RSA encryption/decryption.
   - And more! Check the help command for detailed usage.
### also you can encrypt and decrypt any file 🗃️
like this:
```bash
cipherkit -mod aes256 -file myvideo.mp4
```
or
```bash
cipherkit -mod aes256 -file myvideo.mp4 -output new.mp4
```
## Contributing 🤝

We welcome contributions! If you have suggestions or improvements, feel free to open an issue or submit a pull request.

## License 📄

This project is licensed under the MIT License. See the LICENSE file for details.

---

Thank you for using **CipherKit**! We hope it makes your development experience smoother and more enjoyable. 🎉

---
## Contact Us🙋
- Telegram : @MmdAli_Dev
- Email : core.my.activity@gmail.com