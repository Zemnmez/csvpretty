
$(BASE)%.csv: .FORCE
	go run github.com/zemnmez/cardauth/misc/csvpretty -input $@ -w

.FORCE:
.PHONY: all .FORCE
all: $(wildcard $(BASE)*.csv)