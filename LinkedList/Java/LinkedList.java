import java.util.Collections;
import java.util.Arrays;

public class LinkedList {

    // head of the list
    Node head;

    // class of nodes
    static class Node {

        int data;
        Node next;

        // constructor of Node
        // note that data which stored in list can be any type you want
        Node(int data){

            this.data = data;
            this.next = null;

        }

        @Override
        public String toString() {
            return this.data + "";
        }
    }

    /**
     * insert new value at the end of linkedlist
     * @param list
     * @param value
     * @return
     */
    public static LinkedList insert(LinkedList list, int value){

        // initialize new node
        Node newNode = new Node(value);
        
        // check if head is null insert newNode to head
        if (isEmpty(list)) {
            
            list.head = newNode;

        } else { // head is not null

            try {
                
                    // init new Node to iterate in list to find last node
                Node last = lastNode(list);

                // last node found and now we can insert newNode to next of it.
                last.next = newNode;

            } catch (Exception e) {
                System.err.println(e.getMessage());
            }

        }

        // and finally return edited list
        return list;

    }

    /**
     * accept an array and insert it at the end of linkedlist
     * this method use insert method
     * 
     * @param list
     * @param array
     * @return
     */
    public static LinkedList insertArray(LinkedList list, int[] array) {

        if (array != null && array.length != 0) {

            for(int el: array) {
                list = insert(list, el);
            }

        }

        return list;

    }

    /**
     * insert value to specific index
     * 
     * @param list
     * @param value
     * @param index
     * @return
     * @throws NegativeArraySizeException
     */
    public static LinkedList insertAt(LinkedList list, int value, int index) throws NegativeArraySizeException {

        // check validty of index
        if (index < 0) 
            throw new NegativeArraySizeException("Your index is negative!");

        Node newNode = new Node(value);

        if (index == 0) {

            newNode.next = list.head;
            list.head = newNode;

        } else {
            // init new Node to find previous Node of expected Node
            Node beforeNode = list.head;

            // iterating to find node at index of ( index - 1 )
            for (int i = 1; i < index ; i++) {

                if (beforeNode.next != null) {
                    
                    beforeNode = beforeNode.next;

                } else { // if index number be greater than length of list, newNode insert at the end of list

                    break;

                }

            }

            // make new connection between old Nodes and new one
            newNode.next = beforeNode.next;
            beforeNode.next = newNode;

        }

        return list;

    }

    /**
     * delete a node at specific index
     * 
     * @param list
     * @param index
     * @return
     * @throws NegativeArraySizeException
     */
    public static LinkedList deleteAt(LinkedList list, int index) throws NegativeArraySizeException {

        // check validty of index  
        if (index < 0) 
            throw new NegativeArraySizeException("Your index is negative!");

        // check emptiness of list
        // if list be empty, deletion doesn't make sense
        if (!isEmpty(list)) {

            // check for head deletion
            if (index == 0) {

                Node temp = list.head.next;
                list.head = temp;

            } else {

                // init new Node to find previous Node of expected Node
                Node beforeNode = list.head;

                // iterating to find node at index of ( index - 1 )
                for (int i = 1; i < index ; i++) {

                    if (beforeNode.next != null) {
                        
                        beforeNode = beforeNode.next;

                    } else { // if index number be greater than length of list, input list will be return

                        return list;

                    }

                }

                // skipping expected Node
                beforeNode.next = beforeNode.next.next;

            }

        } 

        return list;

    }

    /**
     * accept an array and delete values at those indexes
     * this method use deleteAt method to delete index by index
     *  
     * @param list
     * @param indexes
     * @return
     * @throws NegativeArraySizeException
     */
    public static LinkedList deleteArrayOfIndex(LinkedList list, int[] indexes) throws NegativeArraySizeException {

        // check emptiness of entered array
        if (indexes != null && indexes.length != 0) {

            // sort entered array in descending form
            indexes = Arrays.stream(indexes).boxed()
                                                .sorted(Collections.reverseOrder())
                                                .mapToInt(Integer::intValue)
                                                .toArray();

            // delete nodes index by index
            for (int index : indexes) {
                list = deleteAt(list, index);
            }

        }

        return list;

    }
    
    /**
     * delete a value if it is found in list
     * this method use deleteAt and indexOf method
     * 
     * @param list
     * @param value
     * @return
     */
    public static LinkedList deleteByValue(LinkedList list, int value) {
        
        // find index of value
        int index = LinkedList.indexOf(list, value);
        // if index be -1 then value has not found
        if (index >= 0) {
            
            // delete value 
            return LinkedList.deleteAt(list, index);
            
        }
        
        return list;
        
    }
    
    /**
     * accept an array and delete values of list which found in the array
     * 
     * @param list
     * @param values
     * @return
     */
    public static LinkedList deleteArrayOfValue(LinkedList list, int[] values) {

        // check emptiness of entered array
        if (values != null && values.length != 0) {

            // delete values one by one
            for (int value: values) {
                list = deleteByValue(list, value);
            }

        }

        return list;

    }
    
    /**
     * get index of an value in list if it exist else this method return -1
     * @param list
     * @param value
     * @return
     */
    public static int indexOf(LinkedList list, int value) {

        // check emptiness of list
        if (!isEmpty(list)) {

            // value is in head
            if (list.head.data == value) {
                return 0;
            } else { // value is not in head

                Node temp = list.head;          // init new Node for iteration
                int counter = 0;                // this is counter to find index of value if it exist
                
                // iteration
                while(temp.next != null) {

                    temp = temp.next;
                    counter++;

                    if (temp.data == value) {   // check equality
                        return counter;
                    }

                }

            }

        }

        return -1;                              // if this value doesn't exist in list

    }

    /**
     * get value at specific index also you must pass a default value
     * to be returned when entered index isn't valid
     * 
     * @param list
     * @param index
     * @param none
     * @return
     */
    public static int get(LinkedList list, int index, int none) {

        // check emptiness of list
        if (isEmpty(list)) {
            
            if (index == 0) {
                return list.head.data;
            } 

            Node temp = list.head;
            int counter = 0;

            // iteration
            while (temp.next != null) {

                temp = temp.next;
                counter++;

                if (counter == index) {     // check equality
                    return temp.data;
                }

            }

        }

        return none;
        
    }

    /**
     * length of list
     * 
     * @param list
     * @return
     */
    public static int length(LinkedList list) {

        // length of empty list is 0
        if (isEmpty(list)) {
            return 0;
        }
        
        Node temp = list.head;          // new temp Node for iteration
        int length = 1;                 // a counter for length of list

        // iteration
        while (temp.next != null) {

            length++;
            temp = temp.next;

        }

        return length;

    }

    @Override
    public String toString() {
        
        StringBuilder listStr = new StringBuilder("[");     // init StringBuilder for appending string
                                                            // also open bracket

        // check emptiness of list
        if (isEmpty(this)) {

            listStr.append("]");                            // close bracket for empty list

        } else { // this list in not empty

            Node temp = this.head;                          // new temp Node for iteration
            listStr.append(temp + ", ");

            // iteration
            while(temp.next != null) {

                temp = temp.next;
                if (temp.next != null) {                    // this is not last element
                    listStr.append(temp + ", ");
                } else {                                    // this is last element
                    listStr.append(temp);
                }

            }

            listStr.append("]");                            // close bracket

        }

        return listStr.toString();

    }

    /**
     * Note: to use this method internally you must before check
     * the list didn't be empty
     * 
     * @param list
     * @return last Node
     */
    private static Node lastNode(LinkedList list) throws Exception {

        // check emptiness of list
        if (isEmpty(list)) {
            throw new Exception("This list is Empty. so This method can't handle it!");
        }

        Node last = list.head;          // new Node for iteration and find last node

        // iteration
        while (last.next != null) {
            last = last.next;
        }
        
        return last;

    }

    /**
     * emptiness checking
     * 
     * @param list
     * @return
     */
    private static boolean isEmpty(LinkedList list) {

        // check nullity of head
        if (list.head == null) {
            return true;
        }

        return false;

    }

    public static void main(String[] args){

        // tests

        LinkedList list = new LinkedList();

        LinkedList.insertArray(list, new int[]{1,2,3,4,5,6});
        LinkedList.insert(list, 7);
        System.out.println(list);
        LinkedList.insertAt(list, 12, 0);
        LinkedList.insertAt(list, 14, 4);
        LinkedList.insertAt(list, 16, 20);
        System.out.println(list);
        LinkedList.deleteAt(list, 5);
        System.out.println(list);
        LinkedList.deleteAt(list, 0);
        System.out.println(list + " , " + LinkedList.length(list));
        LinkedList.deleteByValue(list, 7);
        System.out.println(list + " , " + LinkedList.length(list));
        LinkedList.deleteByValue(list, 123);
        System.out.println(list + " , " + LinkedList.length(list));

        LinkedList.insertArray(list, new int[]{21,22,23,24});
        System.out.println(list + " , " + LinkedList.length(list));

        LinkedList.insertArray(list, new int[]{});
        System.out.println(list + " , " + LinkedList.length(list));

        LinkedList.deleteArrayOfIndex(list, new int[]{0,1,2,3});
        System.out.println(list + " , " + LinkedList.length(list));

        LinkedList.deleteArrayOfIndex(list, new int[]{});
        System.out.println(list + " , " + LinkedList.length(list));

        LinkedList.deleteArrayOfValue(list, new int[]{21,22,23,24});
        System.out.println(list + " , " + LinkedList.length(list));

        LinkedList.deleteArrayOfValue(list, new int[]{});
        System.out.println(list + " , " + LinkedList.length(list));

    }
}