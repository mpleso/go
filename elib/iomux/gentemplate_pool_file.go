// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=iomux -id file -d Data=files -d PoolType=filePool -d Type=Filer github.com/platinasystems/go/elib/pool.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iomux

import (
	"github.com/platinasystems/go/elib"
)

type filePool struct {
	elib.Pool
	files []Filer
}

func (p *filePool) GetIndex() (i uint) {
	l := uint(len(p.files))
	i = p.Pool.GetIndex(l)
	if i >= l {
		p.Validate(i)
	}
	return i
}

func (p *filePool) PutIndex(i uint) (ok bool) {
	return p.Pool.PutIndex(i)
}

func (p *filePool) IsFree(i uint) (v bool) {
	v = i >= uint(len(p.files))
	if !v {
		v = p.Pool.IsFree(i)
	}
	return
}

func (p *filePool) Resize(n uint) {
	c := elib.Index(cap(p.files))
	l := elib.Index(len(p.files) + int(n))
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]Filer, l, c)
		copy(q, p.files)
		p.files = q
	}
	p.files = p.files[:l]
}

func (p *filePool) Validate(i uint) {
	c := elib.Index(cap(p.files))
	l := elib.Index(i) + 1
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]Filer, l, c)
		copy(q, p.files)
		p.files = q
	}
	if l > elib.Index(len(p.files)) {
		p.files = p.files[:l]
	}
}

func (p *filePool) Elts() uint {
	return uint(len(p.files)) - p.FreeLen()
}

func (p *filePool) Len() uint {
	return uint(len(p.files))
}

func (p *filePool) Foreach(f func(x Filer)) {
	for i := range p.files {
		if !p.Pool.IsFree(uint(i)) {
			f(p.files[i])
		}
	}
}

func (p *filePool) ForeachIndex(f func(i uint)) {
	for i := range p.files {
		if !p.Pool.IsFree(uint(i)) {
			f(uint(i))
		}
	}
}

func (p *filePool) Reset() {
	p.Pool.Reset()
	if len(p.files) > 0 {
		p.files = p.files[:0]
	}
}
