# JWT TOKEN 

-----
this is a simple jwt token library

## WHAT IS JWT

-----
```
jwt.io definition, "JSON Web Token (JWT) is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed. JWTs can be signed using a secret (with the HMAC algorithm) or a public/private key pair using RSA. 
```

Huh, That's a mouth full. I'll just JWT is a way of securely verify request over HTTP without the need of sessions. This allows you to write web services that can be consumed by over the web, mobile, desktop and any other device that can communicate over via HTTP.  
It consists of three part
1. Header - base64encoded, defines the algorithm for used for hashing the signature and the type of JWT.
1. Payload- contains information about the issue and the user of the token eg expiration date(ESP)
1. Signature - contains a Hash of the header and payload Using the algorithm in the header
All these are concatenated with a dot, to form a token. 

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3OD..kwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95Or..M7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ
 ```

That's a bare bone summary, but you can check out this [article](https://medium.com/vandium-software/5-easy-steps-to-understanding-json-web-tokens-jwt-1164c0adfcec) for a detailed explanation... Afterall we're here to build a library who cares what it does smiles.

