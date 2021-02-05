
bk: main.go
		go build
install: 
		cp $(CURDIR)/bk /usr/local/bin/bk
		cp $(CURDIR)/jp.fish ~/.config/fish/functions
#		cp $(CURDIR)/jp.sh /usr/local/bin/jp
clean:
		rm /usr/local/bin/bk
		rm ~/.config/fish/functions/jp.fish
#		rm /usr/local/bin/jp
