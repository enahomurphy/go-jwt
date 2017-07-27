package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Payload contains details of the user and issuer
// this might a subdivided claims
// reserved claims , public claims and private
// This jwt will encode reserved and public claims
// SUB - Subject might contain user id
// ISS - issuer signature eg (url)
// Exp - expiration date for token in milliseconds
// AUD - Audience
// PUBLIC - contains additional information for the user
type Payload struct {
	Sub    string      `json:"sub,omitempty"`
	Exp    int64       `json:"exp,omitempty"`
	Iss    string      `json:"iss,omitempty"`
	Aud    string      `json:"aud,omitempty"`
	Public interface{} `json:"public,omitempty"`
}

// jwt token contains three segment that makes a whole
// these segments consist of the header, payload and signature
// the header consist of the hashing algorithm used for
// creating the token, this is encoded using a base64
// encoding algorithm. The payload contains user and issuers
// information, and also the time the the token is to last
// this is also encoded using a base64 encoding algorithm.
// the final part of the token is the signature, this consist
// of the encoded header and payload hashed using the hashing
// algorithm specified in the encoded header

// Base64Encode takes in a string and returns
// a bases64 encode string
func Base64Encode(src string) string {
	data := []byte(src)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

// Base64Decode takes in an Encoded string and returns
// the decoded value of that string if there is an error
// decoding the string, an error is returned with an
// empty string
func Base64Decode(src string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		errMsg := fmt.Errorf("Decoding Error %s", err)
		return "", errMsg
	}
	return string(decoded), nil
}

// Hmac256 generates a Hmac256 hash of a string using
// a specified secret
// NB the hash is irreversible but the value passed in
// can be evaluated by comoaring the message with the
// returned hash
func Hmac256(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// CompareHmac compares a Hmac256 hash against a message
// by hasing the message using the key and checking
// if the hash of that message equals the hash passed
// in. the resulting value is a boolean true or false
func CompareHmac(message string, messageHmac string, secret string) bool {
	key := []byte(secret)
	mac1 := hmac.New(sha256.New, key)
	mac1.Write([]byte(message))
	expectedMac := base64.StdEncoding.EncodeToString(mac1.Sum(nil))
	return expectedMac == messageHmac
}

// getHeader this creates the jwt header.
// Alg: algorithm used
// Typ: of token
func getHeader() string {
	type Header struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	str, _ := json.Marshal(header)
	return Base64Encode(string(str))
}

// Encode creates a jwt token by hashing the encoded header
// containing the hasing algorithms and the encoded payload
// of the user using the HS256 algorithm. this method returns
// a token concatenated with the base64 encoded payload and header
// {header}.{payload}.{signature}
func Encode(payload Payload, secret string) string {
	header := getHeader()
	encodedPayload, _ := json.Marshal(payload)
	signatureValue := header + "." + Base64Encode(string(encodedPayload))
	return signatureValue + "." + Hmac256(signatureValue, secret)
}

// Decode validates and decodes user payload doing the following
// sliting the token into three parts and decoding the payload
// the resulting value is converted to struct
// the EXP on the decoded value is verified by checking if the
// current time exceeds whats in the token. This shows the token has expired
// if that checks out, the signature is tested with the secret -
// key to validate the token if that passes the payload is returned
func Decode(jwt string, secret string) (interface{}, error) {
	token := strings.Split(jwt, ".")

	// check if the jwt token contains
	// header, payload and token
	if len(token) != 3 {
		splitErr := errors.New("Invalid token: token should contain header, payload and secret")
		return nil, splitErr
	}
	// decode payload
	decodedPayload, PayloadErr := Base64Decode(token[1])
	if PayloadErr != nil {
		return nil, fmt.Errorf("Invalid payload: %s", PayloadErr.Error())
	}
	payload := Payload{}

	// parses payload from string to a struct
	ParseErr := json.Unmarshal([]byte(decodedPayload), &payload)
	if ParseErr != nil {
		return nil, fmt.Errorf("Invalid payload: %s", ParseErr.Error())
	}

	if payload.Exp != 0 && time.Now().Unix() > payload.Exp {
		return nil, errors.New("Expired token: token has expired")
	}

	signatureValue := token[0] + "." + token[1]

	// verifies if the header and signature is exactly whats in
	// the signature
	if CompareHmac(signatureValue, token[2], secret) == false {
		return nil, errors.New("Invalid token")
	}

	return payload, nil
}
