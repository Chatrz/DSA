
import java.io.BufferedReader; 
import java.io.IOException; 
import java.io.InputStreamReader; 
import java.util.StringTokenizer; 
// import java.util.Scanner;
public class Boxes{
	public static void main (String[] args)throws IOException{
		BufferedReader br=new BufferedReader(new InputStreamReader(System.in));
		StringTokenizer st=new StringTokenizer(br.readLine());
		// Scanner input=new Scanner(System.in);
		int boxNum=Integer.parseInt(st.nextToken());
		int moveNum=Integer.parseInt(st.nextToken());
		PlaceHolder holder=new PlaceHolder(boxNum);
		for(int i=1;i<=moveNum;i++){
			st=new StringTokenizer(br.readLine());
			int b1=Integer.parseInt(st.nextToken());
			int b2=Integer.parseInt(st.nextToken());
			holder.moveBox(b1,b2);
		}
		st=new StringTokenizer(br.readLine());
		int placeToPrint=Integer.parseInt(st.nextToken());

		holder.printHolder(placeToPrint);

	}
	static class PlaceHolder{
		Node[] boxHolders;
		public PlaceHolder(int boxNum){
			boxHolders=new Node[boxNum+1];
			for(int i=1;i<boxNum+1;i++){
				boxHolders[i]=new Node(i);
			}
		}
		public void moveBox(int iHolder,int jHolder){
			Node nodeToMove=boxHolders[iHolder];
			if(nodeToMove==null)return;
			Node baseNode=boxHolders[jHolder];
			if(baseNode==null){
				boxHolders[jHolder]=nodeToMove;
				boxHolders[iHolder]=null;
				return;
			}
			baseNode.tail.next=nodeToMove;
			baseNode.nodeAhead+=nodeToMove.nodeAhead;
			baseNode.tail=nodeToMove.tail;
			boxHolders[iHolder]=null;

    	}
    	public void printHolder(int index){
    		Node currentNode=boxHolders[index];
    		System.out.print(currentNode==null?"0":currentNode.nodeAhead+" ");
    		while(currentNode!=null){
    			System.out.print(currentNode.id+" ");
    			currentNode=currentNode.next;
    		}
       	}
	}
	static class Node{
		int id;
		int nodeAhead;
		Node next;
		Node tail;

		public Node (int id){
			this.id=id;
			nodeAhead=1;
			next=null;
			tail=this;
		}
	}
}