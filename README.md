# Ethereum address generator

Simple docker image for generate Ethereum address (private key, public key, address).
Print:
```
Private key:  
Public key:   
Address:      
```
 
Save to the files:
```
./addr/privateKey
./addr/publicKey
./addr/address
```

## Run

### Docker

```sh
docker run -i -t --rm \
    ghcr.io/g1ibby/eth-address-generator
```

Save address to file

```sh
docker run -i -t --rm \
    -v $(pwd)/addr:/app/addr \
    ghcr.io/g1ibby/eth-address-generator
```

### Source

Use of the `./scripts/build` script is optional (it's main feature is
embedding git version into compiled binary), you can use usual
`go get|install|build` to get the application instead.

```
$ ./scripts/build
$ ./bin/eth-address-generator
Private key:  9f911b2cafa0cb79ab3ebb52a6afee80968f4c09b3a12e2de3666e7550deea31
Public key:   d46b9864925a758bdb28063fa3d567b1bc6dec4ae616525ed4d3fcdd3f206a52ddbc681a69f79ebe2518a71f33fcc325150a3c499c3e86bb530494b3cb822be0
Address:      0xD4A69D7b736740a15db7fc2e8ab5A5491DeF9550
```
