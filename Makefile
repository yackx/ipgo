DOCKER_IMAGE = yackx/pandoc

default: pdf

pdf:
	@echo 'Pandoc md to PDF'
	@docker run -v "`pwd`":/data $(DOCKER_IMAGE) pandoc \
	    -N --variable version=2.1 \
		--defaults=templates/pdf.yaml \
		--defaults=templates/input-files.yaml \
		--listings \
		-o build/ipgo.pdf

epub:
	@echo 'Pandoc md to epub'
	@docker run -v "`pwd`"":/data $(DOCKER_IMAGE) pandoc \
	    -N --variable version=2.1 \
		--defaults=templates/epub.yaml \
		--defaults=templates/input-files.yaml \
		--listings \
		-o build/ipgo.epub

pandoc-version:
	@docker run $(DOCKER_IMAGE) pandoc -v

cover-pdf:
# 	Imagemagick converter
	convert cover.png -quality 100 cover.pdf
