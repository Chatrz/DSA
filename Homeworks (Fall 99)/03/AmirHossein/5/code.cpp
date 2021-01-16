x <= find(key[x]);
current <= root[T];

int total <= 0;

while (current != x)
{
    if (key[current] < key[x])
    {
        total += nodes_under[left[current]] + 1;
        current <= right[current];
    } else 
    {
        current <= left[current];
    }
}
total += nodes_under + 1;
return total;