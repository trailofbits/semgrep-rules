
func LuhnCheck(_ pan: Int, check: Int?=nil) -> (Bool, Int?) {
    var num: Int!
    num = pan
    var cd: Int! = nil
    var sum = 0, digit = 0
    var shouldDouble: Bool? = false
    if let ch = check {
        cd = ch
    } else {
        cd = num % 10
        num /= 10
    }
    while num > 0 {
        digit = num % 10
        num /= 10
        if let d = shouldDouble {
            if d {
                digit *= 2
            }
            shouldDouble!.toggle()
        }
        sum += digit
    }
    let correct_check = 10 - (sum % 10)
    // ok: assign-nil-to-implicitly-unwrapped-optional
    shouldDouble = nil
    if cd == correct_check {
        // ruleid: assign-nil-to-implicitly-unwrapped-optional
        num = nil
        return (true, cd)
    } else {
        // ruleid: assign-nil-to-implicitly-unwrapped-optional
        cd = nil
        return (false, nil)
    }
}

print("Hello, world!")
print(LuhnCheck(4429000011112224))
print(LuhnCheck(4429000011112223))
