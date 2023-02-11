.PHONY: clean check

EXE=mpibf

$(EXE): .FORCE

.FORCE:
	go build -o $(EXE) src/*.go

clean:
	rm -f $(EXE)
	rm -f a.out
	rm -f *.x
	rm -f *.s

TEST_FILES=hello_world add rot13 pi
TESTS=$(patsubst %, test/%, $(TEST_FILES))

check: $(EXE) $(TESTS)

test/%: test/%.0.x test/%.1.x
	test ! -f $@.in || diff -q $@.out <(./$< < $@.in)
	test -f $@.in || diff -q $@.out <(./$<)
	test ! -f $@.in || diff -q $@.out <(./$(word 2,$^) < $@.in)
	test -f $@.in || diff -q $@.out <(./$(word 2,$^))

test/%.0.x: test/%.bf
	./mpibf -o $@ -O0 $<

test/%.1.x: test/%.bf
	./mpibf -o $@ -O1 $<
