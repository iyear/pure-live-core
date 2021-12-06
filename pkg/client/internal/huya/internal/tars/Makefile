.PHONY: install
install:
	go install github.com/TarsCloud/TarsGo/tars/tools/tars2go

.PHONY: gen
gen:
	tars2go -add-servant=false -tarsPath github.com/TarsCloud/TarsGo/tars *.tars
