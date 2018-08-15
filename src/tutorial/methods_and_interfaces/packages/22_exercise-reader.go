package methods_and_interfaces

import "tutorial/go_packages"

type MyReader struct{}

//Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 65
	}
	return len(bytes), nil
}

func ReadersExercise22() {
	go_packages.Validate(MyReader{})
}
