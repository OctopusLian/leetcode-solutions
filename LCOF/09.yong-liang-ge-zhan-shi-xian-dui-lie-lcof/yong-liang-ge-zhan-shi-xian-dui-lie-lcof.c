typedef struct {
    int len;
    int top1;
    int top2;
    int *s1;//栈1，入栈=入队
    int *s2;//栈2，出栈=出队
} CQueue;


CQueue* cQueueCreate() {
    CQueue *queue = malloc(sizeof(CQueue));
    queue->len = 10000;
    queue->top1 = -1;
    queue->top2 = -1;
    queue->s1 = malloc(queue->len * sizeof(int));
    queue->s2 = malloc(queue->len * sizeof(int));
    return queue;
}

void cQueueAppendTail(CQueue* obj, int value) {
    if(obj->top1 == -1)
        while(obj->top2 != -1)
            obj->s1[++(obj->top1)] = obj->s2[obj->top2--];
    obj->s1[++(obj->top1)] = value;
}

int cQueueDeleteHead(CQueue* obj) {
    if(obj->top2 == -1)
        while(obj->top1 != -1)
            obj->s2[++(obj->top2)] = obj->s1[obj->top1--];
    return obj->top2==-1 ? -1 : obj->s2[obj->top2--];
}

void cQueueFree(CQueue* obj) {
    free(obj->s1);
    free(obj->s2);
    free(obj);
}

/**
 * Your CQueue struct will be instantiated and called as such:
 * CQueue* obj = cQueueCreate();
 * cQueueAppendTail(obj, value);
 
 * int param_2 = cQueueDeleteHead(obj);
 
 * cQueueFree(obj);
*/