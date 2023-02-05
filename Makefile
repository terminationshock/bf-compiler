.PHONY: clean check

exe=mpibf

$(exe): .FORCE

.FORCE:
	go build -o $(exe) src/*

clean:
	rm -f $(exe)
	rm -f a.out
	rm -f *.x

check: $(exe) hello_world.test add.test rot13.test

%.test: %.x
	test ! -f test/$*.in || diff -q test/$*.out <(./$*.x < test/$*.in)
	test -f test/$*.in || diff -q test/$*.out <(./$*.x)

%.x: test/%.bf
	./mpibf -o $@ $<
