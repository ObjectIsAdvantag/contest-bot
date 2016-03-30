PROGRAM=webhook
GOAPP_CMD=/c/go_appengine/goapp.bat


default: dev

dev: clean goget tunnel serve

production: clean goget deploy

clean:
	rm -f $(PROGRAM) $(PROGRAM).exe $(PROGRAM).zip $(PROGRAM).debug

goget:
	go get golang.org/x/net/context
	go get google.golang.org/appengine
	go get google.golang.org/appengine/log

tunnel:
	lt -p 8080 -s contestbot &

serve:
	 $(GOAPP_CMD) serve -admin_port 8567

deploy:
	$(GOAPP_CMD) deploy


