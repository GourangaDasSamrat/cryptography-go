# Learn Cryptography with Golang

This is a course from [Boot.dev](https://www.boot.dev/courses/learn-cryptography-golang)

## Chapter 1: Symmetric Encryption

### Levels

- **Level 1**: Cryptography Is About Encryption (Coding Challenge)
- **Level 2**: What Is Cryptography? (Quiz)
- **Level 3**: What Is Cryptography? (Quiz)
- **Level 4**: Symmetric vs. Asymmetric (Coding Challenge)
- **Level 5**: Symmetric vs. Asymmetric (Quiz)
- **Level 6**: Symmetric vs. Asymmetric (Quiz)
- **Level 7**: Better Keys (Coding Challenge)
- **Level 8**: Cryptology, Cryptography and Cryptanalysis (Quiz)
- **Level 9**: Cryptology, Cryptography and Cryptanalysis (Quiz)
- **Level 10**: Cryptology, Cryptography and Cryptanalysis (Quiz)
- **Level 11**: Single-Use Keys (Coding Challenge)
- **Level 12**: Single-Use Keys Review (Quiz)
- **Level 13**: Single-Use Keys Review (Quiz)

## Chapter 2: Encoding

### Levels

- **Level 1**: Encoding Bytes (Coding Challenge)
- **Level 2**: Encoding != Encryption (Quiz)
- **Level 3**: Encoding != Encryption (Quiz)
- **Level 4**: Encoding != Encryption (Quiz)
- **Level 5**: Formatting (Coding Challenge)
- **Level 6**: Decoding (Coding Challenge)

## Chapter 3: Brute Force

### Levels

- **Level 1**: Alphabet Formula (Coding Challenge)
- **Level 2**: Brute Force Attacks (Quiz)
- **Level 3**: Brute Force Attacks (Quiz)
- **Level 4**: Brute Force Attacks (Quiz)
- **Level 5**: Crack an Insecure Key (Coding Challenge)
- **Level 6**: Entropy (Quiz)
- **Level 7**: Entropy (Quiz)

## Chapter 4: Caesar Cipher

### Levels

- **Level 1**: What Is a Cipher? (Quiz)
- **Level 2**: What Is a Cipher? (Quiz)
- **Level 3**: Caesar Cipher (Coding Challenge)
- **Level 4**: Caesar Cipher Security (Quiz)
- **Level 5**: Caesar Cipher Security (Quiz)

## Chapter 5: XOR

### Levels

- **Level 1**: Exclusive or (XOR) (Coding Challenge)
- **Level 2**: XOR Quiz (Quiz)
- **Level 3**: XOR Quiz (Quiz)
- **Level 4**: One Time Pad (Coding Challenge)
- **Level 5**: Perfect Security (Quiz)
- **Level 6**: Perfect Security (Quiz)
- **Level 7**: Perfect Security (Quiz)

## Chapter 6: Stream Ciphers

### Levels

- **Level 1**: Stream Ciphers (Coding Challenge)
- **Level 2**: Stream Ciphers (Quiz)
- **Level 3**: Stream Ciphers (Quiz)
- **Level 4**: State Ciphers (Coding Challenge)
- **Level 5**: State Ciphers (Quiz)
- **Level 6**: State Ciphers (Quiz)

## Chapter 7: Block Ciphers

### Levels

- **Level 1**: Block Ciphers (Coding Challenge)
- **Level 2**: Block vs. Stream (Quiz)
- **Level 3**: Block vs. Stream (Quiz)
- **Level 4**: Block Sizes (Coding Challenge)
- **Level 5**: Chunks Review (Quiz)
- **Level 6**: Chunks Review (Quiz)
- **Level 7**: Chunks Review (Quiz)
- **Level 8**: Key Schedules (Coding Challenge)
- **Level 9**: Key Schedule Review (Quiz)
- **Level 10**: Key Schedule Review (Quiz)

## Running Tests

Use the `justfile` to run tests for a specific chapter and level:

```bash
just test <chapter> <level>
```

**Examples:**

- `just test 1 1` - Run tests for Chapter 1, Level 1
- `just test 1 4` - Run tests for Chapter 1, Level 4

## License

This repository contains solutions for the [Boot.dev Cryptography with Golang](https://www.boot.dev/courses/learn-cryptography-golang) course.

**Note:** Test files (`*_test.go`, `main_tests.go` and `helpers.go`) are copyrighted by Boot.dev. These are provided as part of the course material and are used for validating solutions.

All other code in this repository is licensed under the MIT License.
