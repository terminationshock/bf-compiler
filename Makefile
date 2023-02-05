.PHONY: clean check

exe=mpibf

$(exe): .FORCE

.FORCE:
	go build -o $(exe) src/*

clean:
	rm -f $(exe)
	rm -f a.out
	rm -f *.x

check: $(exe) hello_world.x

%.x: test/%.bf
	./mpibf -o $@ $<
	diff -q $<.out <(./$@)
