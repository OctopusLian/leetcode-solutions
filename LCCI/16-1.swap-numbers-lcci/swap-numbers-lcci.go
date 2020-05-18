func swapNumbers(numbers []int) []int {
    var result []int
    result = append(result,numbers[1],numbers[0])
    return result
}