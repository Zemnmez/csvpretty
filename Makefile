SHELL:=/bin/bash -O globstar

.INTERMEDIATE: sum.Makefile

%/Makefile.include: %/Makefile
	echo "BASE:=$*/" > $@
	echo "include $<" >> $@


sum.Makefile: $(addsuffix .include,$(shell ls ./*/**/Makefile))
	cat $^ > $@
	rm $^

.PHONY: dummy

dummy: sum.Makefile
	STEP=DO $(MAKE) -$(MAKEFLAGS) -j -f $< $(MAKECMDGOALS)

%: dummy
	-
