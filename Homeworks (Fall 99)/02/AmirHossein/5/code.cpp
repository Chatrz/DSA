Stack buffer[n]; // Create a stack for keeping the nodes addresses

inorder_tree_walk(Tree A)
{
    current = root[A]; // Start from head

    if (current == nullptr)
    {
        return; // Empty tree
    }

    while (true)
    {

        if (left[current] != nullptr)
        {
            buffer.push(current); // Keep in stack
            current = left[current]; // Switch to left child
            continue;
        }

        print(key[current]); // Print current

        if (right[current] != nullptr)
        {
            current = right[current]; // Switch to right child
            continue;
        }

        if (buffer == empty)
        {
            break; // Exit condition
        }

        while (buffer != empty)
        {
            current = buffer.pop(); // Buffer clearing
            print(key[current]);

            if (right[current] != nullptr)
            {
                current = right[current];
                break;
            }
        }
    }
}