
all : nutsender.exe

nutsender.exe : main.go window.go config.go rsrc.syso
	go build -ldflags="-H windowsgui"

#rsrc.syso : main.manifest nutsender.ico
#	rsrc -manifest main.manifest -ico nutsender.ico -o rsrc.syso

rsrc.syso : main.manifest
	rsrc -manifest main.manifest -o rsrc.syso

clean :
	-rm *.syso *.exe config.xml

