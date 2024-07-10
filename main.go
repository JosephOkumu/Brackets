package main

import (
    "fmt"
    "os"
)

// isBalanced function checks if the given string s represents a balanced sequence of brackets
// It returns true if the brackets are balanced, and false otherwise
func isBalanced(s string) bool {
    // Create a stack to keep track of opening brackets
    stack := []rune{}
    // Create a map to store the corresponding closing brackets for each opening bracket
    brackets := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    // Iterate over each character in the string
    for _, char := range s {
        switch char {
        // If the character is an opening bracket, push it onto the stack
        case '(', '[', '{':
            stack = append(stack, char)
        // If the character is a closing bracket
        case ')', ']', '}':
            // If the stack is empty or the top of the stack does not match the corresponding opening bracket, return false
            if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
                return false
            }
            // If the closing bracket matches the opening bracket at the top of the stack, remove the opening bracket from the stack
            stack = stack[:len(stack)-1]
        }
    }

    // If the stack is empty at the end, it means all brackets are balanced
    return len(stack) == 0
}

func main() {
    // Get the command-line arguments (excluding the program name)
    args := os.Args[1:]
    // Iterate over each argument
    for _, arg := range args {
        // Check if the argument represents a balanced sequence of brackets
        if isBalanced(arg) {
            fmt.Println("OK")
        } else {
            fmt.Println("Error")
        }
    }
}
