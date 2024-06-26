PANDOC ?= pandoc

default: pdf

pdf:
	@echo 'Pandoc md to PDF'
	@$(PANDOC) \
	    -N --variable version=2.1 \
		--defaults=templates/pdf.yaml \
		--defaults=templates/input-files.yaml \
		--listings \
		-o build/ipgo.pdf

epub:
	@echo 'Pandoc md to epub'
	@$(PANDOC) \
	    -N --variable version=2.1 \
		--defaults=templates/epub.yaml \
		--defaults=templates/input-files.yaml \
		--listings \
		-o build/ipgo.epub

cover-pdf:
# 	Imagemagick converter
	convert content/cover-bw.png -quality 100 content/cover-bw.pdf

clean:
	@echo 'Cleaning build directory'
	@rm -rf build
