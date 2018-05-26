# Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# blame: jnml, labs.nic.cz

PREFIX	:= ruler
LEX		:= golex
YACC	:= goyacc

ifeq ($(bash -c 'sed --help &> &1' | grep GNU),)
SED		:= sed -i ""
else
SED 	:= sed -i
endif

.PHONY: all test

all: grammar.go tokenizer.go

test: all
	go test

%.go: %.l
	$(LEX) -o $@ $<

%.go: %.y
	$(YACC) -o $@ $<
	$(SED) -e '/^func yyParse/s/)/, outNode **node&/' $@

clean:
	go clean
	rm -f *.tmp filter.go tokenizer.go