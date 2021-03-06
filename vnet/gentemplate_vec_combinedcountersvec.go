// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=vnet -id combinedCountersVec -d VecType=CombinedCountersVec -d Type=CombinedCounters github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vnet

import (
	"github.com/platinasystems/go/elib"
)

type CombinedCountersVec []CombinedCounters

func (p *CombinedCountersVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]CombinedCounters, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *CombinedCountersVec) validate(new_len uint, zero CombinedCounters) *CombinedCounters {
	old_cap := uint(cap(*p))
	old_len := uint(len(*p))
	if new_len <= old_cap {
		// Need to reslice to larger length?
		if new_len > old_len {
			*p = (*p)[:new_len]
			for i := old_len; i < new_len; i++ {
				(*p)[i] = zero
			}
		}
		return &(*p)[new_len-1]
	}
	return p.validateSlowPath(zero, old_cap, new_len, old_len)
}

func (p *CombinedCountersVec) validateSlowPath(zero CombinedCounters, old_cap, new_len, old_len uint) *CombinedCounters {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]CombinedCounters, new_cap, new_cap)
		copy(q, *p)
		for i := old_len; i < new_cap; i++ {
			q[i] = zero
		}
		*p = q[:new_len]
	}
	if new_len > old_len {
		*p = (*p)[:new_len]
	}
	return &(*p)[new_len-1]
}

func (p *CombinedCountersVec) Validate(i uint) *CombinedCounters {
	var zero CombinedCounters
	return p.validate(i+1, zero)
}

func (p *CombinedCountersVec) ValidateInit(i uint, zero CombinedCounters) *CombinedCounters {
	return p.validate(i+1, zero)
}

func (p *CombinedCountersVec) ValidateLen(l uint) (v *CombinedCounters) {
	if l > 0 {
		var zero CombinedCounters
		v = p.validate(l, zero)
	}
	return
}

func (p *CombinedCountersVec) ValidateLenInit(l uint, zero CombinedCounters) (v *CombinedCounters) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *CombinedCountersVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p CombinedCountersVec) Len() uint { return uint(len(p)) }
