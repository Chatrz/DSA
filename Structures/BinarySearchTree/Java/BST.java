public class BST{
    private Node root;
    public BST(){
        root=null;
    }
    public BST(int value){
        root=new Node(value);
    }
    public BST(Node root){
        this.root=root;
    }



    //search
    public Node search(int value){
        Node currentNode=root;
        while(currentNode!=null){
            if(currentNode.value==value) return currentNode;
            else if(currentNode.value<value)currentNode=currentNode.right;
            else currentNode=currentNode.left;
        }
        return null;
    }
    public Node recursiveSearch(int value){
        return recursiveSearch(root, value);
    }
    public Node recursiveSearch(Node root,int value){
        //this is how java handles default value for parameters
        if(root.value==value||root==null)return root;
        else if(root.value<value)return recursiveSearch(root.right, value);
        else return recursiveSearch(root.left, value);
    }
    //insert
    public Node Insert(int value){
        Node currentNode=root;
        while(currentNode!=null){
            if(currentNode.value<value)currentNode=currentNode.right;
            else currentNode=currentNode.left;
        }
        currentNode=new Node(value);
        return currentNode;
    }
    public Node recursiveInsert(int value){
        return recursiveInsert(root,value);
    }
    public Node recursiveInsert(Node root,int value){
        if(root==null){
            root=new Node(value);
            return root;
        }

        if(root.value<value)return recursiveInsert(root.right,value);
        else return recursiveInsert(root.left,value);
    }
    public Node findMin(Node root){
        Node currentNode=root;
        while(currentNode.left!=null)currentNode=currentNode.left;
        return currentNode;
    }
    public Node findMax(Node root){
        Node currentNode=root;
        while(currentNode.right!=null)currentNode=currentNode.right;
        return currentNode;
    }
    //delete
    //convert list to bst
    //tree traversals
    //presessor 
    public Node findPredecessor(Node targetNode){
        if(targetNode.left!=null) return findMax(targetNode.left);
        int baseValue=targetNode.value;
        Node currentNode=root;
        Node predecessor=new Node(-1);
        while(currentNode!=null){
            if(currentNode.value>baseValue)currentNode=currentNode.left;
            else if(currentNode.value<baseValue){
                predecessor=currentNode;
                currentNode=currentNode.right;
            }
            else break;
        }
        return predecessor;
    }
    //successor
    public Node findSuccessor(Node targetNode){
        if (targetNode.right!=null) return findMin(targetNode.right);
        int baseValue=targetNode.value;
        Node currentNode=root;
        Node successor=new Node(-1);
        while(currentNode!=null){
            if(currentNode.value<baseValue)currentNode=currentNode.right;
            else if(currentNode.value>baseValue){
                successor=currentNode;
                currentNode=currentNode.left;
            }
            else break;
        }
        return successor;
    }
}