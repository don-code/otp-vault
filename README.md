OTP Vault
=========

This application allows one-time password (OTP) secrets to be securely stored and used to generate tokens. It borrows *heavily* from 99designs' [AWS Vault](https://github.com/99designs/aws-vault).

Like AWS Vault, [99designs' keyring](https://github.com/99designs/keyring) is used to provide multi-backend storage of OTP secret keys. [go-otp](https://github.com/hgfischer/go-otp) provides OTP token generation.

## Usage
### Adding keys
```bash
# Add a TOTP key named 'my-totp' with secret key ABCD1234
> otp-vault add my-totp ABCD1234

# Add an HOTP key named 'my-hotp' with secret key ABCD1234, counter initialized to 20
> otp-vault add my-hotp ABCD1234 hotp 20

# Add a TOTP key with an 8-character length and 5-second period
> otp-vault add my-secure-totp ABCD1234 totp 0 15 8

```

### Generating tokens
```bash
# Get a key for my-totp
> otp-vault get my-totp
< 123456 (2017-09-17 18:13:34.144377561 -0400 EDT)
```

### Listing all keys
```bash
> otp-vault list
my-totp
my-hotp
my-secure-totp

```

### Removing a key
```bash
> otp-vault rm my-totp
< Delete credentials for profile "my-totp"? (Y|n)
> Y
```
