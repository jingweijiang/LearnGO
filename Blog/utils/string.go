package utils

import "math/rand"

// rune 和 string 的区别
// rune 是 Go 语言的基本类型，表示一个 Unicode 码点，可以用一个或多个 rune 组成一个 string。
// string 是 Go 语言的内置类型，是一个不可变的序列，由若干个字节组成，每个字节表示一个 Unicode 码点。
// 因此，string 类型可以包含任意的 Unicode 字符，而 rune 类型只能包含单个 Unicode 码点。
// 例如，"Hello, 世界" 是一个 string，而它的每个字符都是一个 rune。

func randomString(n int) string {

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
