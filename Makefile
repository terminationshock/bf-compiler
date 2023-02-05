#.PHONY:

exe=mpibf

$(exe): .FORCE

.FORCE:
	go build -o $(exe) src/*

clean:
	rm -f $(exe)
