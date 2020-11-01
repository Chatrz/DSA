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
        if(root==null)return null;
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
        if(root==null){
            root=new Node(value);
            return root;
        }
        Node currentNode=root;
        while(currentNode!=null){
            if(currentNode.value<value)currentNode=currentNode.right;
            else currentNode=currentNode.left;
        }
        currentNode=new Node(value);
        return currentNode;
    }
    public Node recursiveInsert(int value){
        if(root==null){
            root=new Node(value);
            return root;
        }
        return recursiveInsert(root,value);
    }
    public Node recursiveInsert(Node root,int value){
        if(root==null){
            root=new Node(value);
            return root;
        }
        if(root.value<value)root.right=recursiveInsert(root.right,value);
        else root.left=recursiveInsert(root.left,value);

        return root;
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
    public void deleteNode(int value){
        deleteNode(root,null,value);
    }
    public void deleteNode(Node currentNode,Node parentNode, int value){
        if(currentNode==null)return;
        if(currentNode.value>value)deleteNode(currentNode.left,currentNode, value);
        else if(currentNode.value<value)deleteNode(currentNode.right,currentNode,value);
        else{
            //handles the node with only one or zero child
            if(currentNode.left==null || currentNode.right==null){
                Node childNode= currentNode.left==null?currentNode.right:currentNode.left;
                if(parentNode.left==currentNode)parentNode.left=childNode;
                else parentNode.right=childNode;
            }

            //handle the node with 2 children
            Node successor=findSuccessor(currentNode);
            currentNode.value=successor.value;
            deleteNode(currentNode.right, currentNode, successor.value);
        }
    }
    //convert list to bst
    //tree traversals
    public String inorderTraversal(){
        return this.inorderTraversal(root);
    }
    public String inorderTraversal(Node root){
        if(root==null)return "";
        String repr="";
        repr+=inorderTraversal(root.left);
        repr+=Integer.toString(root.value)+ " ";
        repr+=inorderTraversal(root.right);
        return repr;
    }
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