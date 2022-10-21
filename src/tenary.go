package src

func If[T any](cond bool, v1, v0 T) T {
    if cond {
        return v1
    }
    return v0
}