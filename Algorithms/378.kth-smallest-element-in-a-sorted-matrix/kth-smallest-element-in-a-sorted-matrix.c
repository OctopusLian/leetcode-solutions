int cmp(const void *a, const void *b) { return (*(int *)a - *(int *)b); }

int kthSmallest(int **matrix, int matrixSize, int *matrixColSize, int k) {
    //解法1，直接排序
    int *rec = (int *)malloc(matrixSize * matrixSize * sizeof(int));

    int num = 0;
    for (int i = 0; i < matrixSize; i++) {
        for (int j = 0; j < matrixColSize[i]; j++) {
            rec[num++] = matrix[i][j];
        }
    }
    qsort(rec, num, sizeof(int), cmp);

    return rec[k - 1];
}