package utils

type FunctorInputData map[string]interface{}
type FunctorOutputData map[string]interface{}

type Functor func(FunctorInputData) FunctorOutputData
