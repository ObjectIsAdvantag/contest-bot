PROGRAM=webhook
GOAPP_CMD=/c/go_appengine/goapp.bat


default: dev

clean:
	rm -f $(PROGRAM) $(PROGRAM).exe $(PROGRAM).zip $(PROGRAM).debug

tunnel:
	lt -p 8080 -s contestbot &

dev:
	 $(GOAPP_CMD) serve -admin_port 8567

deploy:
	$(GOAPP_CMD) deploy


