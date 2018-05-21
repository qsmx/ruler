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

all: grammar.go tokenizer.go

test: all
	go test

%.go: %.l
	$(LEX) -o $@ $<

grammar.go: grammar.y
	$(YACC) -p $(PREFIX) -o $@ $<
	$(SED) '/func $(PREFIX)Parse/s/)/, dp *DataPackage, res *bool&/' $@

clean:
	go clean
	rm -f grammar.go tokenizer.go

