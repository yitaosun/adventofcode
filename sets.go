package main

type Set map[interface{}]struct{}

func NewSet(vals ...interface{}) Set {
	s := Set{}
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

func (s Set) Add(k interface{}) Set {
	s[k] = struct{}{}
	return s
}

func (s Set) Remove(k interface{}) Set {
	delete(s, k)
	return s
}

func (s Set) Has(k interface{}) bool {
	_, ok := s[k]
	return ok
}

func (s Set) AddAll(o Set) Set {
	for k := range o {
		s[k] = struct{}{}
	}
	return s
}

func (s Set) RemoveAll(o Set) Set {
	for k := range o {
		delete(s, k)
	}
	return s
}

func (s Set) KeepAll(o Set) Set {
	for k := range s {
		if _, ok := o[k]; !ok {
			delete(s, k)
		}
	}
	return s
}

func (s Set) ToSlice() []interface{} {
	rv := []interface{}{}
	for k := range s {
		rv = append(rv, k)
	}
	return rv
}

func (s Set) Equals(o Set) bool {
	if len(s) != len(o) {
		return false
	}
	for k := range s {
		if _, ok := o[k]; !ok {
			return false
		}
	}
	return true
}
