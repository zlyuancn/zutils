/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/23
   Description :
-------------------------------------------------
*/

package zutils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash/fnv"
	"net/url"
)

var Crypto = new(cryptoUtil)

type cryptoUtil struct{}

func (*cryptoUtil) Md5(text string) string {
	c := md5.New()
	c.Write([]byte(text))
	return hex.EncodeToString(c.Sum(nil))
}
func (*cryptoUtil) Md5_16(text string) string {
	c := md5.New()
	c.Write([]byte(text))
	return hex.EncodeToString(c.Sum(nil))[8:24]
}

func (*cryptoUtil) Md5Bytes(text []byte) []byte {
	c := md5.New()
	c.Write(text)
	return c.Sum(nil)
}
func (*cryptoUtil) Md5_16Bytes(text []byte) []byte {
	c := md5.New()
	c.Write(text)
	return c.Sum(nil)[8:24]
}

func (*cryptoUtil) Sha1(text string) string {
	c := sha1.New()
	c.Write([]byte(text))
	return hex.EncodeToString(c.Sum(nil))
}
func (*cryptoUtil) Sha224(text string) string {
	v := sha256.Sum224([]byte(text))
	return hex.EncodeToString(v[:])
}
func (*cryptoUtil) Sha256(text string) string {
	v := sha256.Sum256([]byte(text))
	return hex.EncodeToString(v[:])
}
func (*cryptoUtil) Sha384(text string) string {
	v := sha512.Sum384([]byte(text))
	return hex.EncodeToString(v[:])
}
func (*cryptoUtil) Sha512(text string) string {
	v := sha512.Sum512([]byte(text))
	return hex.EncodeToString(v[:])
}

func (*cryptoUtil) Sha1Bytes(text []byte) []byte {
	c := sha1.New()
	c.Write(text)
	return c.Sum(nil)
}
func (*cryptoUtil) Sha224Bytes(text []byte) []byte {
	v := sha256.Sum224(text)
	return v[:]
}
func (*cryptoUtil) Sha256Bytes(text []byte) []byte {
	v := sha256.Sum256(text)
	return v[:]
}
func (*cryptoUtil) Sha384Bytes(text []byte) []byte {
	v := sha512.Sum384(text)
	return v[:]
}
func (*cryptoUtil) Sha512Bytes(text []byte) []byte {
	v := sha512.Sum512([]byte(text))
	return v[:]
}

func (*cryptoUtil) Fnv32a(text string) uint32 {
	c := fnv.New32a()
	_, _ = c.Write([]byte(text))
	return c.Sum32()
}
func (*cryptoUtil) Fnv64a(text string) uint64 {
	c := fnv.New64a()
	_, _ = c.Write([]byte(text))
	return c.Sum64()
}
func (*cryptoUtil) Fnv128a(text string) string {
	c := fnv.New128a()
	c.Write([]byte(text))
	return hex.EncodeToString(c.Sum(nil))
}

func (*cryptoUtil) Fnv32aBytes(text []byte) uint32 {
	c := fnv.New32a()
	_, _ = c.Write(text)
	return c.Sum32()
}
func (*cryptoUtil) Fnv64aBytes(text []byte) uint64 {
	c := fnv.New64a()
	_, _ = c.Write(text)
	return c.Sum64()
}
func (*cryptoUtil) Fnv128aBytes(text []byte) []byte {
	c := fnv.New128a()
	c.Write(text)
	return c.Sum(nil)
}

func (*cryptoUtil) HmacMd5(text, key string) string {
	c := hmac.New(md5.New, []byte(key))
	c.Write([]byte(text))
	return hex.EncodeToString(c.Sum(nil))
}
func (*cryptoUtil) HmacSha1(text, key string) string {
	c := hmac.New(sha1.New, []byte(key))
	c.Write([]byte(text))
	return hex.EncodeToString(c.Sum(nil))
}
func (*cryptoUtil) HmacSha256(text, key string) string {
	c := hmac.New(sha256.New, []byte(key))
	c.Write([]byte(text))
	return hex.EncodeToString(c.Sum(nil))
}
func (*cryptoUtil) HmacSha512(text, key string) string {
	c := hmac.New(sha512.New, []byte(key))
	c.Write([]byte(text))
	return hex.EncodeToString(c.Sum(nil))
}

func (*cryptoUtil) HmacMd5Bytes(text, key []byte) []byte {
	c := hmac.New(md5.New, key)
	c.Write(text)
	return c.Sum(nil)
}
func (*cryptoUtil) HmacSha1Bytes(text, key []byte) []byte {
	c := hmac.New(sha1.New, key)
	c.Write(text)
	return c.Sum(nil)
}
func (*cryptoUtil) HmacSha256Bytes(text, key []byte) []byte {
	c := hmac.New(sha256.New, key)
	c.Write(text)
	return c.Sum(nil)
}
func (*cryptoUtil) HmacSha512Bytes(text, key []byte) []byte {
	c := hmac.New(sha512.New, key)
	c.Write(text)
	return c.Sum(nil)
}

func (*cryptoUtil) Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}
func (*cryptoUtil) Base64Decode(text string) (string, error) {
	bs, err := base64.StdEncoding.DecodeString(text)
	return string(bs), err
}

func (*cryptoUtil) Base64EncodeBytes(text []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(text)))
	base64.StdEncoding.Encode(buf, text)
	return buf
}
func (*cryptoUtil) Base64DecodeBytes(text []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(text)))
	n, err := base64.StdEncoding.Decode(buf, text)
	return buf[:n], err
}

func (*cryptoUtil) UrlEncode(text string) string {
	return url.QueryEscape(text)
}
func (*cryptoUtil) UrlDecode(text string) (string, error) {
	return url.QueryUnescape(text)
}
