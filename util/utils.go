package util

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unsafe"

	"github.com/gofiber/fiber/v2"
)

// 비밀번호 유효성 체크 (영문, 숫자, 특수문자 1개이상을 포함하는 8 ~ 15 자리)
func ValidPassword(pwd string) (bool) {
	var (
		hasMaxLen  = false
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(pwd) >= 8 {
		hasMinLen = true
		if len(pwd) <= 15 {
			hasMaxLen = true
		}
	}
	for _, c := range pwd {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return  hasMaxLen  && hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// 회원 os 정보, 접속 IP, 기기 정보 반환
func UserAgentProvider(c *fiber.Ctx) (fiber.Map) {

	userAgent := c.Get("User-Agent")
	info := strings.Split(userAgent, "/")

	if len(info) == 3 {
		os := info[0]
		ipAddr := info[1]
		device := info[2]

		if os == "" || ipAddr == "" || device == "" || (os != "ios" && os != "android") {
			return fiber.Map{}
		}

		// 어플에서 인식이 안된 상태로 올 경우 0.0.0.0 받음
		if ipAddr == "0.0.0.0" {

			clientIp := c.IP()
			arr := strings.Split(clientIp, ":")
			clientIp = arr[len(arr) - 1]

			privateIpMatch := ValidPrivateIpv4(clientIp)
			if privateIpMatch {
				ipAddr = "0.0.0.0"
			} else {
				ipAddr = clientIp
			}
		}

		ipMatch := ValidIpv4(ipAddr)
		if !ipMatch { return fiber.Map{} }

		return fiber.Map{"os": os, "ipAddr": ipAddr, "device": device}
	}

	return fiber.Map{}
}

// 사설 ip 주소 유효성 검사
func ValidPrivateIpv4(ip string) bool {
	ip = strings.Trim(ip, " ")
	re := regexp.MustCompile(`(^127\.)|(^10\.)|(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)|(^192\".168\.)`)
	if re.MatchString(ip) {
		return true
	} else {
		return false
	}
}

// ip 주소 유효성 검사
func ValidIpv4(ip string) bool {
	ip = strings.Trim(ip, " ")
	re := regexp.MustCompile(`^(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}$`)
	if re.MatchString(ip) {
		return true
	} else {
		return false
	}
}

func MakeRamdomString(n int) string {
	const (
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func GetYmd() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02")
}

func GetYmdHms() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}