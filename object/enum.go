package object

import "strings"

type Enum struct {
	Branches map[string]Object
}

func (en *Enum) Type() ObjectType { return EnumObj }
func (en *Enum) Inspect() string {
	out := strings.Builder{}
	out.WriteString("{")
	str := make([]string, len(en.Branches))
	i := 0
	for s, v := range en.Branches {
		outBranches := strings.Builder{}
		outBranches.WriteString("case ")
		outBranches.WriteString(s)
		outBranches.WriteString(" : ")
		outBranches.WriteString(v.Inspect())
		str[i] = outBranches.String()
		i++
	}
	out.WriteString(strings.Join(str, "; "))
	out.WriteString("}")
	return out.String()
}
