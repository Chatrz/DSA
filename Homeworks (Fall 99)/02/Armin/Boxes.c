#include <stdio.h>
#include <stdlib.h>

struct Box {
    int key;
    struct Box *next;
};
struct Place {
    int size;
    struct Box *boxes;
};

int main() {
    int n, m, x, y, d;
    scanf("%d %d", &n, &m);
    struct Place places[n + 1];
    for (int i = 1; i <= n; i++) {
        struct Box *newBox = (struct Box *) malloc(sizeof(struct Box));
        newBox->key = i;
        newBox->next = NULL;
        places[i].boxes = newBox;
        places[i].size = 1;
    }
    for (int j = 0; j < m; j++) {
        scanf("%d %d", &x, &y);
        if (places[x].size != 0) {
            if (places[y].size == 0) {
                places[y].boxes = places[x].boxes;
                places[y].size = places[x].size;
            } else {
                struct Box *tmp;
                for (tmp = places[y].boxes; tmp->next != NULL; tmp = tmp->next);
                tmp->next = places[x].boxes;
                places[y].size += places[x].size;
            }
            places[x].boxes = NULL;
            places[x].size = 0;
        }
    }
    scanf("%d", &d);
    printf("%d ", places[d].size);
    struct Box *tmp;
    for (tmp = places[d].boxes; tmp != NULL; tmp = tmp->next) {
        printf("%d ", tmp->key);
    }
    return 0;
}
