LinkedList A[n]; // A linked list with size n
LinkedList B[m]; // A linked list with size m

reverse(A); >> Time = O(n)
reverse(B); >> Time = O(m)

a <- head(A); >> C1
b <- head(B); >> C2

while (a != nullptr && b != nullptr) >> Time = O( min(n,m) )
{
    if (a->next->key == b->next->key) >> Time = C3
    {
        a = a->next;
        b = b->next;
    } else
    {
        break;
    }
}

if (a->key == b->key) >> Time = C4
{
    return a;
} else
{
    return nullptr;
}
