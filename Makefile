DOCKER_IMAGE = yackx/pandoc

default: pdf

pdf:
	@echo 'Pandoc md to PDF'
	@docker run -v `pwd`:/data $(DOCKER_IMAGE) pandoc \
	    -N --variable version=2.1 \
		--defaults=pdf.yaml \
		--defaults=input-files.yaml \
		--listings \
		-o build/ipgo.pdf

epub:
	@echo 'Pandoc md to epub'
	@docker run -v `pwd`:/data $(DOCKER_IMAGE) pandoc \
	    -N --variable version=2.1 \
		--defaults=epub.yaml \
		--defaults=input-files.yaml \
		--listings \
		-o build/ipgo.epub

pdf-sample:
	@echo 'Pandoc md to PDF sample'
	@docker run -v `pwd`:/data $(DOCKER_IMAGE) pandoc \
	    -N --variable version=2.1 \
		--defaults=pdf.yaml \
		--defaults=input-samples.yaml \
		--listings \
		-o build/ipgo-sample.pdf

epub-sample:
	@echo 'Pandoc md to epub sample'
	@docker run -v `pwd`:/data $(DOCKER_IMAGE) pandoc \
	    -N --variable version=2.1 \
		--defaults=epub.yaml \
		--defaults=input-samples.yaml \
		--listings \
		-o build/ipgo-sample.epub

pandoc-version:
	@docker run $(DOCKER_IMAGE) pandoc -v

cover-pdf:
# 	Imagemagick converter
	convert cover.png -quality 100 cover.pdf
