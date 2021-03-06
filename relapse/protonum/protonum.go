//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//Package protonum is used to rewrite field names to field numbers in relapse grammars.
package protonum

import (
	"fmt"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/parser/proto"
	"github.com/katydid/katydid/relapse/ast"
)

//FieldNamesToNumbers rewrites field names contained in the grammar to their respective field numbers found in the protocol buffer filedescriptorset.
//This allows for more speedy field comparisons in validation when used in conjunction with the ProtoNumParser.
func FieldNamesToNumbers(pkgName, msgName string, desc *descriptor.FileDescriptorSet, grammar *ast.Grammar) (*ast.Grammar, error) {
	g := grammar.Clone()
	descMap, err := proto.NewDescriptorMap(pkgName, msgName, desc)
	if err != nil {
		return nil, err
	}
	root := descMap.GetRoot()
	refs := ast.NewRefLookup(g)
	nameToNumber := &nameToNumber{
		refs:    make(map[string]*context),
		descMap: descMap,
	}
	nameToNumber.refs["main"] = &context{root, false}
	newRefs := make(map[string]*ast.Pattern)
	oldContexts := 0
	newContexts := 1
	for oldContexts != newContexts {
		oldContexts = newContexts
		for name, context := range nameToNumber.refs {
			newp, err := nameToNumber.translate(context, refs[name])
			if err != nil {
				return nil, err
			}
			newRefs[name] = newp
		}
		newContexts = len(nameToNumber.refs)
	}
	return ast.NewGrammar(newRefs), nil
}

type nameToNumber struct {
	refs    map[string]*context
	descMap proto.DescMap
}

type context struct {
	msg   *descriptor.DescriptorProto
	index bool
}

func (this *context) Equal(that *context) bool {
	return this.msg == that.msg && this.index == that.index
}

func anyErr(err1, err2 error) error {
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

type ErrDup struct {
	name string
	c1   *context
	c2   *context
}

func (this *ErrDup) Error() string {
	return fmt.Sprintf("relapse/protonum: Duplicate Reference Error: Name: %v, Context1: %v, Context2: %v", this.name, this.c1.msg.GetName(), this.c2.msg.GetName())
}

func (this *nameToNumber) translate(context *context, p *ast.Pattern) (*ast.Pattern, error) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty, *ast.LeafNode, *ast.ZAny:
		return p, nil
	case *ast.TreeNode:
		return this.translateName(context, v.GetName(), v.GetPattern())
	case *ast.Concat:
		l, err1 := this.translate(context, v.GetLeftPattern())
		r, err2 := this.translate(context, v.GetRightPattern())
		return ast.NewConcat(l, r), anyErr(err1, err2)
	case *ast.Or:
		l, err1 := this.translate(context, v.GetLeftPattern())
		r, err2 := this.translate(context, v.GetRightPattern())
		return ast.NewOr(l, r), anyErr(err1, err2)
	case *ast.And:
		l, err1 := this.translate(context, v.GetLeftPattern())
		r, err2 := this.translate(context, v.GetRightPattern())
		return ast.NewAnd(l, r), anyErr(err1, err2)
	case *ast.ZeroOrMore:
		p, err := this.translate(context, v.GetPattern())
		return ast.NewZeroOrMore(p), err
	case *ast.Reference:
		c, ok := this.refs[v.GetName()]
		if !ok {
			this.refs[v.GetName()] = context
			return p, nil
		}
		if !c.Equal(context) {
			//TODO we could probably create a new reference here
			//  for every conflicting combination of msg x repeated x referece name
			return nil, &ErrDup{v.GetName(), c, context}
		}
		return p, nil
	case *ast.Not:
		p, err := this.translate(context, v.GetPattern())
		return ast.NewNot(p), err
	case *ast.Contains:
		p, err := this.translate(context, v.GetPattern())
		return ast.NewContains(p), err
	case *ast.Optional:
		p, err := this.translate(context, v.GetPattern())
		return ast.NewOptional(p), err
	case *ast.Interleave:
		l, err1 := this.translate(context, v.GetLeftPattern())
		r, err2 := this.translate(context, v.GetRightPattern())
		return ast.NewInterleave(l, r), anyErr(err1, err2)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

type errExpectedArray struct {
	name string
	c    *context
}

func (this *errExpectedArray) Error() string {
	return fmt.Sprintf("relapse/protonum: Expected Array Error: Name: %v, Msg: %v", this.name, this.c.msg.GetName())
}

type errExpectedField struct {
	name string
	c    *context
}

func (this *errExpectedField) Error() string {
	return fmt.Sprintf("relapse/protonum: Expected Field Error: Name: %v, Msg: %v", this.name, this.c.msg.GetName())
}

type errAnyFieldNotSupported struct {
	name string
}

func (this *errAnyFieldNotSupported) Error() string {
	return fmt.Sprintf("relapse/protonum: Any Field Not Supported: Name: %v", this.name)
}

type errAnyNameExceptNotSupported struct {
	name string
}

func (this *errAnyNameExceptNotSupported) Error() string {
	return fmt.Sprintf("relapse/protonum: AnyNameExcept Not Supported Error: Name: %v", this.name)
}

type errUnknownField struct {
	name string
	c    *context
}

func (this *errUnknownField) Error() string {
	return fmt.Sprintf("relapse/protonum: Unknown Field Error: Name: %v, Msg: %v", this.name, this.c.msg.GetName())
}

func getField(m map[uint64]*descriptor.FieldDescriptorProto, name string) *descriptor.FieldDescriptorProto {
	for _, f := range m {
		if f.GetName() != name {
			continue
		}
		return f
	}
	return nil
}

func (this *nameToNumber) translateName(current *context, name *ast.NameExpr, child *ast.Pattern) (*ast.Pattern, error) {
	switch n := name.GetValue().(type) {
	case *ast.Name:
		if current.index {
			if n.IntValue == nil {
				return nil, &errExpectedArray{name.String(), current}
			}
			c := &context{current.msg, false}
			newp, err := this.translate(c, child)
			if err != nil {
				return nil, err
			}
			return ast.NewTreeNode(name, newp), nil
		}
		if n.StringValue == nil {
			return nil, &errExpectedField{name.String(), current}
		}
		f := getField(this.descMap.LookupFields(current.msg), n.GetStringValue())
		if f == nil {
			return nil, &errUnknownField{name.String(), current}
		}
		msg := this.descMap.LookupMessage(f)
		c := &context{msg, f.IsRepeated()}
		newp, err := this.translate(c, child)
		if err != nil {
			return nil, err
		}
		newName := ast.NewUintName(uint64(f.GetNumber()))
		return ast.NewTreeNode(newName, newp), nil
	case *ast.AnyName:
		if current.index {
			c := &context{current.msg, false}
			newp, err := this.translate(c, child)
			if err != nil {
				return nil, err
			}
			return ast.NewTreeNode(name, newp), nil
		} else {
			return nil, &errAnyFieldNotSupported{name.String()}
		}
	case *ast.AnyNameExcept:
		return nil, &errAnyNameExceptNotSupported{name.String()}
	case *ast.NameChoice:
		l, err1 := this.translateName(current, n.GetLeft(), child)
		r, err2 := this.translateName(current, n.GetRight(), child)
		return ast.NewOr(l, r), anyErr(err1, err2)
	}
	panic(fmt.Sprintf("unknown name typ %T", name))
}
