include $(GOROOT)/src/Make.inc

TARG=semaphore
GOFILES=\
	semaphore.go \
	counter.go \
	twophasebarrier.go \
	main.go

main: main.6
	6l -o main main.6

main.6: main.go changingroom.6
	6g main.go

changingroom.6: changingroom.go semaphore.6 counter.6
	6g changingroom.go

twophasebarrier.6: twophasebarrier.go semaphore.6 counter.6
	6g twophasebarrier.go

counter.6: counter.go semaphore.6
	6g counter.go

semaphore.6: semaphore.go
	6g semaphore.go

include $(GOROOT)/src/Make.pkg