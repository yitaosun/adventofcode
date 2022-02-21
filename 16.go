package main

import (
	"fmt"
	"log"
	"math"
)

var bitmasks16 = map[int]int{
	1: 0b0001,
	2: 0b0011,
	3: 0b0111,
	4: 0b1111,
}

type packet16 struct {
	version int
	typ     int
	literal int
	length  int
	packets []*packet16
}

func (p *packet16) String() string {
	if p.typ == 4 {
		return fmt.Sprintf("version=%d type=%d length=%d literal=%d",
			p.version, p.typ, p.length, p.literal)
	} else {
		return fmt.Sprintf("version=%d type=%d length=%d packets=%d",
			p.version, p.typ, p.length, len(p.packets))
	}
}

type bits16 struct {
	h string
	i int
}

func (b *bits16) readBits(n int) int {
	// log.Printf("reading %d bits", n)
	var rv int
	bitsRemaining := n
	for {
		var finalByte, origByte int
		r := b.h[b.i/4]
		if r >= '0' && r <= '9' {
			origByte = int(r - '0')
		} else {
			origByte = int(r-'A') + 10
		}
		bitsUsed := 4 - (b.i % 4)
		finalByte = origByte & bitmasks16[bitsUsed]
		// log.Printf("1. orig=%#b final=%#b bits=%d", origByte, finalByte, bitsUsed)
		if bitsRemaining <= bitsUsed {
			finalByte >>= (bitsUsed - bitsRemaining)
			bitsUsed = bitsRemaining
		}
		// log.Printf("2. orig=%#b final=%#b bits=%d", origByte, finalByte, bitsUsed)
		rv = (rv << bitsUsed) | finalByte
		bitsRemaining -= bitsUsed
		b.i += bitsUsed
		if bitsRemaining <= 0 {
			// log.Printf("n=%d i=%d rv=0%b", n, b.i, rv)
			return rv
		}
	}
}

func (b *bits16) readPacket() *packet16 {
	p := &packet16{length: 6}
	p.version = b.readBits(3)
	p.typ = b.readBits(3)
	if p.typ == 4 {
		for {
			g := b.readBits(5)
			p.literal = p.literal<<4 | (g & 0b1111)
			p.length += 5
			if (g & 0b10000) == 0 {
				break
			}
		}
	} else {
		lt := b.readBits(1)
		if lt == 0 {
			l := b.readBits(15)
			p.length += 16
			for {
				sub := b.readPacket()
				p.packets = append(p.packets, sub)
				p.length += sub.length
				l -= sub.length
				if l <= 0 {
					break
				}
			}
		} else {
			p.length += 12
			for i := b.readBits(11) - 1; i >= 0; i-- {
				sub := b.readPacket()
				p.packets = append(p.packets, sub)
				p.length += sub.length
			}
		}
	}
	// m := 4 - (p.length % 4)
	// p.length += m
	// b.readBits(m)
	return p
}

func (p *bits16) length() int {
	return len(p.h) * 4
}

func (p *packet16) totalVersion() int {
	total := p.version
	s := NewStack()
	for _, sub := range p.packets {
		s.Push(sub)
	}
	for {
		if s.Size() == 0 {
			break
		}
		sub := s.Pop().(*packet16)
		total += sub.version
		for _, subSub := range sub.packets {
			s.Push(subSub)
		}
	}
	return total
}

func (p *packet16) eval() int {
	var rv int
	switch p.typ {
	case 0:
		for _, sub := range p.packets {
			rv += sub.eval()
		}
	case 1:
		rv = 1
		for _, sub := range p.packets {
			rv *= sub.eval()
		}
	case 2:
		rv = math.MaxInt
		for _, sub := range p.packets {
			if v := sub.eval(); v < rv {
				rv = v
			}
		}
	case 3:
		rv = math.MinInt
		for _, sub := range p.packets {
			if v := sub.eval(); v > rv {
				rv = v
			}
		}
	case 4:
		rv = p.literal
	case 5:
		if p.packets[0].eval() > p.packets[1].eval() {
			rv = 1
		}
	case 6:
		if p.packets[0].eval() < p.packets[1].eval() {
			rv = 1
		}
	case 7:
		if p.packets[0].eval() == p.packets[1].eval() {
			rv = 1
		}
	default:
		log.Fatal("invalid packet type", p.typ)
	}
	return rv
}

func solve16A(input []string) int {
	for _, line := range input {
		b := &bits16{h: line}
		p := b.readPacket()
		log.Printf("Read packet %v with total version of %d",
			p, p.totalVersion())
	}
	return 0
}

func solve16B(input []string) int {
	for _, line := range input {
		b := &bits16{h: line}
		p := b.readPacket()
		log.Printf("Read packet %v that evaluates to %d",
			p, p.eval())
	}
	return 0
}
