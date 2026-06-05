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

**Note:** Test files (`*_test.go` and `main_tests.go`) are copyrighted by Boot.dev. These are provided as part of the course material and are used for validating solutions.

All other code in this repository is licensed under the MIT License.
