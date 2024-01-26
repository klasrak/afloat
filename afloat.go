package afloat

// noCopy is used to ensure that Float32 cannot be copied.
// Go does not have a native way to prevent copying of types.
// The `vet` can detect types that cannot be copied because they contain a `Lock()` method.
// To prevent copying of other types, embed a `noCopy` type.
// See https://github.com/golang/go/issues/8005#issuecomment-190753527 for more information.
type noCopy struct{}

// Lock is a no-op used to ensure that noCopy cannot be copied.
func (*noCopy) Lock() {}
