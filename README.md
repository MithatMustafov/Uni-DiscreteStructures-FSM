# Finite State Machine - Two 'a' Characters Recognizer

This project implements a Deterministic Finite Automaton (DFA) that recognizes strings containing exactly two 'a' characters over the alphabet Σ = {a, b}.

## Problem Description

Given a set of strings L over alphabet Σ = {a, b}, the automaton recognizes strings that contain exactly two 'a' characters. For example:
- ✅ "abab" belongs to L (contains exactly two 'a's)
- ✅ "bababbb" belongs to L (contains exactly two 'a's)
- ❌ "abbb" does not belong to L (contains only one 'a')
- ❌ "bab" does not belong to L (contains only one 'a')

## State Machine Description

The FSM consists of three states:
- Start (0): Initial state
- GotFirstA (1): State after reading first 'a'
- GotSecondA (2): Final state after reading second 'a'

### State Transitions
```
Start (0):
  - on 'a' → GotFirstA
  - on 'b' → Start

GotFirstA (1):
  - on 'a' → GotSecondA
  - on 'b' → GotFirstA

GotSecondA (2):
  - on 'a' → Error (invalid)
  - on 'b' → GotSecondA
```

## Usage

### Prerequisites
- Go 1.16 or higher

### Running the Program
```bash
go run main.go <input_string>
```

### Examples
```bash
# Example 1: Accepted string (abab)
$ go run main.go abab
■ Processing input string: abab
➤  Processing input: 'a' ➜  Transitioning to state: 1
➤  Processing input: 'b' ➜  Transitioning to state: 1
➤  Processing input: 'a' ➜  Transitioning to state: 2
➤  Processing input: 'b' ➜  Transitioning to state: 2
■ String accepted

# Example 2: Rejected string (bab - only one 'a')
$ go run main.go bab
■ Processing input string: bab
➤  Processing input: 'b' ➜  Transitioning to state: 0
➤  Processing input: 'a' ➜  Transitioning to state: 1
➤  Processing input: 'b' ➜  Transitioning to state: 1
■ String not accepted

# More examples
$ go run main.go bababbb     # Accepted (contains exactly two a's)
$ go run main.go abbb        # Not accepted (contains only one a)
```

Note: The actual output in the terminal will be color-coded:
- Yellow (■) for processing indicators
- Yellow (➤) for state transitions
- Green (■) for acceptance message
- Red (■) for rejection message

### Output Format
The program provides colored terminal output showing:
- Input string being processed
- State transitions for each character
- Final acceptance/rejection status

## Regular Expression
The language can be described by the regular expression:
```
b*ab*ab*
```

## Implementation Details
- Written in Go
- Uses ANSI color codes for terminal output
- Implements error handling for invalid inputs
- Provides detailed state transition logging

## Error Handling
- Invalid characters (not 'a' or 'b') are rejected
- More than two 'a' characters are rejected
- Empty strings are rejected
- Proper usage instructions are shown if no input is provided

## License
This project is part of the Discrete Structures course work and is available for educational purposes.