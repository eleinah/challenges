func isPalindrome(s string) bool {
    expr := `[a-zA-Z0-9]`
    re, _ := regexp.Compile(expr)

    scrubs := re.FindAllString(s, -1)
    scrubbed := ""
    for _, scrub := range scrubs {
        scrubbed += strings.ToLower(scrub)
    }
    scrubbed = strings.ReplaceAll(scrubbed, " ", "")
    rev := reverse(scrubbed)

    return scrubbed == rev
}

func reverse(s string) (result string) {
    for _, c := range s {
        result = string(c) + result
    }
    return
}
