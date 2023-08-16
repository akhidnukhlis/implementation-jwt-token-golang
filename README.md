# implementation-jwt-token-golang

This repository provides a simple guide on how to generate an ECDSA key pair (Private Key and Public Key) using OpenSSL.

## Prerequisites

Before you begin, ensure that you have OpenSSL installed on your system. If not, you can download and install it from the official OpenSSL website or use your package manager.

## Steps to Generate ECDSA Key Pair

1. Open a terminal window.

2. To generate a Private Key, run the following command:

    ``` 
    openssl ecparam -name prime256v1 -genkey -noout -out private_key.pem
    ``` 

    This command will create a private key file named `private_key.pem` using the prime256v1 curve.

3. To extract the Public Key from the Private Key, run the following command:

    ``` 
    openssl ec -in private_key.pem -pubout -out public_key.pem
    ```

    This command will generate a corresponding public key file named `public_key.pem`.

4. You now have your ECDSA key pair generated and stored in the `private_key.pem` (Private Key) and `public_key.pem` (Public Key) files.

## Usage

You can use these keys for various purposes such as digital signatures, authentication, and encryption.

## Disclaimer

Keep your private key secure and never share it publicly. It's recommended to use proper key management practices to safeguard your keys.

## Contact

If you have any questions, suggestions, or need assistance with this key generation process, feel free to reach out to me:

📧 Email: [nukhlis@gmail.com](mailto:nukhlis@gmail.com)
💬 LinkedIn: [@akhidnukhlis](https://www.linkedin.com/in/akhidnukhlis/)

Let's connect and collaborate!

## License

This project is licensed under the [MIT License](LICENSE).
