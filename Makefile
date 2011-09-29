include $(GOROOT)/src/Make.inc

TARG=geometry

GOFILES=\
	line2d.go\
	point2d.go\
	point3d.go\

include $(GOROOT)/src/Make.pkg
