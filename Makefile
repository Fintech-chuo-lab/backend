# Makefile


# serve at local
.PHONY: serve
serve:
	goapp serve ./appengine


# deploy appengine
.PHONY: deploy
deploy:
	goapp deploy ./appengine
