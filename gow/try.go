package gow

// Block struct for try-catch
type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Exception interface in case of catch
type Exception interface{}

// Throw method
func Throw(ex Exception) {
	panic(ex)
}

// Do Try-Catch-Finally block
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
