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

all: filter.go hit.go tokenizer.go validate.go
	@$(SED) -n "/^func filterParse/,/^}/p" filter.go
	@$(SED) -n "/^func validateParse/,/^}/p" validate.go

test: all
	go test

%.go: %.l
	$(LEX) -o $@ $<

%.go: %.y
	@cat yacc.spec.in $< > $<.tmp
	$(YACC) -p $(basename $@) -o $@ $<.tmp
	@rm $<.tmp
ifneq ($@,validate.go)
	@$(SED) '/func $(basename $@)Parse/s/)/, dp *DataPackage, res *bool&/' $@
endif

clean:
	go clean
	rm -f filter.go tokenizer.go hit.go validate.go
