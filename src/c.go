package main

var (
	mpiLib = `#include <mpi.h>
void mympi_init() {
  MPI_Init(NULL, NULL);
}
int mympi_rank() {
  int rank = 0;
  MPI_Comm_rank(MPI_COMM_WORLD, &rank);
  return rank;
}
void mympi_allreduce(void *data) {
  MPI_Allreduce(MPI_IN_PLACE, data, 1, MPI_INT, MPI_SUM, MPI_COMM_WORLD);
}
void mympi_finalize() {
  MPI_Finalize();
}`
	dummyLib = `
void mympi_init() {}
int mympi_rank() {
  return 0;
}
void mympi_allreduce(void *data) {}
void mympi_finalize() {}`
)

func Library(hasMpi bool) string {
	if hasMpi {
		return mpiLib
	}
	return dummyLib
}
