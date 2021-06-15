package main

func main() error {
	val, err := f()
	if err != nil {
		return errors.Wrap(err, "f")
	}
}
