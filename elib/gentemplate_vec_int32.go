// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id Int32 -d VecType=Int32Vec -d Type=int32 vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

type Int32Vec []int32

func (p *Int32Vec) Resize(n uint) {
	c := Index(cap(*p))
	l := Index(len(*p)) + Index(n)
	if l > c {
		c = NextResizeCap(l)
		q := make([]int32, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *Int32Vec) validate(new_len uint, zero int32) *int32 {
	c := Index(cap(*p))
	lʹ := Index(len(*p))
	l := Index(new_len)
	if l <= c {
		// Need to reslice to larger length?
		if l > lʹ {
			*p = (*p)[:l]
			for i := lʹ; i < l; i++ {
				(*p)[i] = zero
			}
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *Int32Vec) validateSlowPath(zero int32, c, l, lʹ Index) *int32 {
	if l > c {
		cNext := NextResizeCap(l)
		q := make([]int32, cNext, cNext)
		copy(q, *p)
		for i := c; i < cNext; i++ {
			q[i] = zero
		}
		*p = q[:l]
	}
	if l > lʹ {
		*p = (*p)[:l]
	}
	return &(*p)[l-1]
}

func (p *Int32Vec) Validate(i uint) *int32 {
	var zero int32
	return p.validate(i+1, zero)
}

func (p *Int32Vec) ValidateInit(i uint, zero int32) *int32 {
	return p.validate(i+1, zero)
}

func (p *Int32Vec) ValidateLen(l uint) (v *int32) {
	if l > 0 {
		var zero int32
		v = p.validate(l, zero)
	}
	return
}

func (p *Int32Vec) ValidateLenInit(l uint, zero int32) (v *int32) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *Int32Vec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p Int32Vec) Len() uint { return uint(len(p)) }
