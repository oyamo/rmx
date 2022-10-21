package src

import "os"

const (
   VERSION = "0.0.1"
)

type Engine struct {
   Input *Input
   Cwd string
}


func NewEngine() (*Engine, error) {
   cwd, err := os.Getwd()
   if err != nil {
      return nil, err
   }
   return &Engine{
      Input: nil,
      Cwd: cwd,
   } , nil
}

func (e *Engine) Run ()  error {
   input, err := NewInput()

   if err != nil {
	  UsageError()
	  return err
   }

   e.Input = input

   if input.HasFlag(FLAG_HELP) {
	  Usage()
	  return nil
   }

   if input.HasFlag(FLAG_V) {
      Version()
      return nil
   }

   return e.delete()
}

func (e *Engine) delete() error {
   for _, file := range e.Input.Files {
     err := e.DeleteItem(file)
     if err != nil {
       return err
     }
   }

   return nil
}