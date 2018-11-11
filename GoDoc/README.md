use 'godoc cmd/go_exercises/GoDoc/example' for documentation on the go_exercises/GoDoc/example command 

PACKAGE DOCUMENTATION

package example
    import "go_exercises/GoDoc/example"

    This is the package comment, a top-level piece of documentation used to
    explain things about the package (see json or exp/template) All godoc
    comments are in this form with no whitespace between them and what they
    accompany

CONSTANTS

const (
    CONSTA = iota
    CONSTB
    CONSTC
    CONSTD
    ANOTHER = 7
)
    Some enum examples

VARIABLES

var Default float64 = 0.7
    This is just a random variable

TYPES

type Example float64
    Example is a float used for demonstration purposes

func (e Example) Sqrt() Example
    Returns the square root of an example

type Example2 struct {
    X Example
    // contains filtered or unexported fields
}
    Example2 is also for demonstartion

func NewExample(num int) *Example2
    NewExample is used to get a ready-to-use Example2


