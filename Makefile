.PHONY: clean check

MPI_HOME=/usr/lib64/mpi/gcc/openmpi4

EXE=mpibf

$(EXE): .FORCE

.FORCE:
	go build -o $(EXE) src/*.go

clean:
	rm -f $(EXE)
	rm -f a.out
	rm -f *.x
	rm -f *.s

SERIAL_TEST_FILES=hello_world add rot13 pi
MPI_TEST_FILES=sum
SERIAL_TESTS=$(patsubst %, test/serial/%, $(SERIAL_TEST_FILES))
MPI_TESTS=$(patsubst %, test/mpi/%, $(MPI_TEST_FILES))

check: $(EXE) $(SERIAL_TESTS) $(MPI_TESTS)

test/serial/%: test/serial/%.x
	test ! -f $@.in || diff -q $@.out <(./$< < $@.in)
	test -f $@.in || diff -q $@.out <(./$<)

test/mpi/%: test/mpi/%.x
	diff -q $@.out <(mpiexec -n 4 ./$<)

test/serial/%.x: test/serial/%.bf
	./mpibf -o $@ $<

test/mpi/%.x: test/mpi/%.bf
	./mpibf -o $@ $< -L$(MPI_HOME)/lib64 -I$(MPI_HOME)/include
