.PHONY: clean check

exe=mpibf

$(exe): .FORCE

.FORCE:
	go build -o $(exe) src/*

clean:
	rm -f $(exe)
	rm -f a.out
	rm -f *.x

check: $(exe) hello_world.test add.test rot13.test mpi_sum.test_mpi

%.test: %.x
	test ! -f test/$*.in || diff -q test/$*.out <(./$*.x < test/$*.in)
	test -f test/$*.in || diff -q test/$*.out <(./$*.x)

%.test_mpi: %.x
	diff -q test/$*.out <(mpiexec -n 4 ./$*.x)

%.x: test/%.bf
	./mpibf -o $@ $< -L$(LD_LIBRARY_PATH)
