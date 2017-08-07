// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=pg -id stream -d PoolType=stream_pool -d Type=Streamer -d Data=elts github.com/platinasystems/go/elib/pool.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pg

import (
	"github.com/platinasystems/go/elib"
)

type stream_pool struct {
	elib.Pool
	elts []Streamer
}

func (p *stream_pool) GetIndex() (i uint) {
	l := uint(len(p.elts))
	i = p.Pool.GetIndex(l)
	if i >= l {
		p.Validate(i)
	}
	return i
}

func (p *stream_pool) PutIndex(i uint) (ok bool) {
	return p.Pool.PutIndex(i)
}

func (p *stream_pool) IsFree(i uint) (v bool) {
	v = i >= uint(len(p.elts))
	if !v {
		v = p.Pool.IsFree(i)
	}
	return
}

func (p *stream_pool) Resize(n uint) {
	c := elib.Index(cap(p.elts))
	l := elib.Index(len(p.elts) + int(n))
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]Streamer, l, c)
		copy(q, p.elts)
		p.elts = q
	}
	p.elts = p.elts[:l]
}

func (p *stream_pool) Validate(i uint) {
	c := elib.Index(cap(p.elts))
	l := elib.Index(i) + 1
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]Streamer, l, c)
		copy(q, p.elts)
		p.elts = q
	}
	if l > elib.Index(len(p.elts)) {
		p.elts = p.elts[:l]
	}
}

func (p *stream_pool) Elts() uint {
	return uint(len(p.elts)) - p.FreeLen()
}

func (p *stream_pool) Len() uint {
	return uint(len(p.elts))
}

func (p *stream_pool) Foreach(f func(x Streamer)) {
	for i := range p.elts {
		if !p.Pool.IsFree(uint(i)) {
			f(p.elts[i])
		}
	}
}

func (p *stream_pool) ForeachIndex(f func(i uint)) {
	for i := range p.elts {
		if !p.Pool.IsFree(uint(i)) {
			f(uint(i))
		}
	}
}

func (p *stream_pool) Reset() {
	p.Pool.Reset()
	if len(p.elts) > 0 {
		p.elts = p.elts[:0]
	}
}
