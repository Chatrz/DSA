Stack main_stack = [10]
Stack holder_stack = [10]


enqueue(key)
{
    while (main_stack != nullptr)
    {
        holder = main_stack.pop();
        holder_stack.push(holder);
    }
    holder_stack.push(key);
    while (holder_stack != nullptr)
    {
        holder = holder_stack.pop();
        main_stack.push(holder);
    }
}

dequeue()
{
    return main_stack.pop();
}