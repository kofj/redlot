CC=`which go`
CCCOLOR="\033[34m"
LINKCOLOR="\033[34;1m"
SRCCOLOR="\033[33m"
BINCOLOR="\033[37;1m"
MAKECOLOR="\033[32;1m"
ENDCOLOR="\033[0m"

run:
	@rm -fr ./bin/redlot
	@$(QUIET_BUILD)$(CC) build -ldflags "-w -s" -o ./bin/redlot ./example/server/main.go$(CCLINK)
	./bin/redlot -log_dir="./var/log" -logtostderr=true
