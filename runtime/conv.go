package runtime

func Fstring(s string) *String {
    return &String{s, objectBase()}
}
