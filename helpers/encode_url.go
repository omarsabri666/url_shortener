package helpers

var Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncodeBase62(num int64) string {
	if num == 0 {
		return "0"
	}

	result := ""
	for num > 0 {
		remainder := num % 62                            // 1️⃣ get remainder
		result = string(Base62Chars[remainder]) + result // 2️⃣ map to a character
		num /= 62                                        // 3️⃣ divide & continue
	}

	return result
}
