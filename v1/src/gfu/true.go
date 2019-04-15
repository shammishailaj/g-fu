package gfu

import (
  //"log"
  "strings"
)

type True struct {
  BasicVal
}

func (t *True) Init(g *G) *True {
  t.BasicVal.Init(&g.TrueType, t)
  return t
}

func (_ *True) Bool(g *G) bool {
  return true
}

func (_ *True) Dump(out *strings.Builder) {
  out.WriteRune('T')
}

func (t *True) Eq(g *G, rhs Val) bool {
  return t == rhs
}

func (t *True) Eval(g *G, task *Task, env *Env) (Val, E) {
  return t, nil
}

func (t *True) Is(g *G, rhs Val) bool {
  return t == rhs
}

func (t *True) Quote(g *G, task *Task, env *Env) (Val, E) {
  return t, nil
}

func (t *True) Splat(g *G, out Vec) Vec {
  return append(out, t)
}

func (t *True) String() string {
  return "T"
}
