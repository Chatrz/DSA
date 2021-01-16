import java.io.BufferedReader; 
import java.io.IOException; 
import java.io.InputStreamReader;
import java.util.StringTokenizer; 
public class LostTree{
    static class Node{
        int id;
        int value; 
        int leftNodeIndex;
        int rightNodeIndex;
        public Node(int id, int leftNodeIndex, int rightNodeIndex){
            this.id=id;
            this.leftNodeIndex=leftNodeIndex;
            this.rightNodeIndex=rightNodeIndex;
        }
    }
    static class LTree{
        private Node[] nodes;
        private int[] roots;
        private int nodeNumber;
        private int valueToAssign;
        public LTree(int numberOfNodes){
            nodeNumber=numberOfNodes;
            nodes=new Node[numberOfNodes+1];//0 index is used to store the root of tree
            roots=new int[numberOfNodes+1];
            for(int i=0;i<numberOfNodes+1;i++){
                roots[i]=i;
            }
            valueToAssign=0;
        }
        public void insertNode(int id, int leftNodeIndex,int rightNodeIndex){
            nodes[id]=new Node(id, leftNodeIndex, rightNodeIndex);
            if(roots[id]==id)roots[0]=id;
            if(leftNodeIndex != -1) roots[leftNodeIndex]=roots[id];
            if(rightNodeIndex != -1)roots[rightNodeIndex]=roots[id];
            int rootIndex=roots[0];
            if(rootIndex!=roots[rootIndex])roots[0]=roots[rootIndex];
            
        }
        private void setRoot(){
            nodes[0]=nodes[roots[0]];
        }
        public void assignValue(){
            setRoot();
            assignValue(nodes[0]);
        }
        public void assignValue(Node root){
            if(root==null) return;
            if(root.leftNodeIndex!=-1) assignValue(nodes[root.leftNodeIndex]);
            valueToAssign++;
            root.value=valueToAssign;
            if(root.rightNodeIndex!=-1) assignValue(nodes[root.rightNodeIndex]);
        }
        public void printValues(){
            for(int i=1;i<nodeNumber+1;i++){
                System.out.print(nodes[i].value);
                System.out.print(i==nodeNumber?"\n":" ");
            }
        }
    }



    public static void main (String[] args)throws IOException{
        BufferedReader input=new BufferedReader(new InputStreamReader(System.in));
        StringTokenizer st=new StringTokenizer(input.readLine());
        int numberOfNodes=Integer.parseInt(st.nextToken());
        LTree tree=new LTree(numberOfNodes);
        for(int i=0;i<numberOfNodes;i++){
            st=new StringTokenizer(input.readLine());
            int id=Integer.parseInt(st.nextToken());
            int leftNodeIndex=Integer.parseInt(st.nextToken());
            int rightNodeIndex=Integer.parseInt(st.nextToken());
            tree.insertNode(id, leftNodeIndex, rightNodeIndex);
        }
        tree.assignValue();
        tree.printValues();
    }
}
