.PHONY: clean check

EXE=bf

$(EXE): .FORCE

.FORCE:
	go build -o $(EXE) src/*.go

clean:
	rm -f $(EXE)
	rm -f a.out
	rm -f *.x
	rm -f *.s
	rm -f OUT.*
	rm -f test/*.x
	rm -f test/*.x.s

TEST_FILES=hello_world add rot13 pi
TESTS=$(patsubst %, test/%, $(TEST_FILES))

check: $(EXE) $(TESTS) test/test_opt

test/%: test/%.0.x test/%.1.x
	test ! -f $@.in || ./$< < $@.in > OUT.0
	test -f $@.in || ./$< > OUT.0
	diff -q $@.out OUT.0
	test ! -f $@.in || ./$(word 2,$^) < $@.in > OUT.1
	test -f $@.in || ./$(word 2,$^) > OUT.1
	diff -q $@.out OUT.1
	test ! -f $@.s || diff -q $@.s $(word 2,$^).s

test/test_opt: test/test_opt.bf
	./$(EXE) -S -O1 $<
	diff -q $@.s a.out.s

test/%.0.x: test/%.bf
	./$(EXE) -o $@ -O0 $<

test/%.1.x: test/%.bf
	./$(EXE) -o $@ -O1 -S $<
