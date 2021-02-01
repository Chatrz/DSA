import java.io.BufferedReader; 
import java.io.IOException; 
import java.io.InputStreamReader; 
import java.util.StringTokenizer; 
public class SuperQueue {
    static class SQueue{
        infoNode root;
        infoNode tail;
        public SQueue(){
            root=null;
            tail=null;
        }
        public void enqueue(String element, String repeat){
            if(root==null){
                root=new infoNode(repeat,element);
                tail=root;
            }
            else{
                tail.next=new infoNode(repeat,element);
                tail=tail.next;
            }
        }
        public void dequeue(String numberOfRemove){
          Long removeNum=Long.parseLong(numberOfRemove);
          if(root==null)return;
          if(root.numberOfObj>removeNum)root.numberOfObj-=removeNum;
          else{
            Long difference=removeNum-root.numberOfObj;
            root=root.next;
            if(root!=null) dequeue(Long.toString(difference));
            
          }
        }
        public void peek(){
          if(root==null)System.out.println("empty");
          else System.out.println(root.obj);
        }
    }
    static class infoNode{
        Long numberOfObj;
        String obj;
        infoNode next;
        public infoNode(String number, String obj){
            numberOfObj=Long.parseLong(number);
            this.obj=obj;
            next=null;
        }
    }
    public static void main (String[] args) throws IOException {
        SQueue queue=new SQueue();
        BufferedReader input=new BufferedReader(new InputStreamReader(System.in));
        StringTokenizer st=new StringTokenizer(input.readLine());
        int requestNumber= Integer.parseInt(st.nextToken());
        for (int i=1;i<=requestNumber;i++){
            st=new StringTokenizer(input.readLine());
            String operator=st.nextToken();
            if(operator.equals("?") )queue.peek();
            else if(operator.equals("+")){
                queue.enqueue(st.nextToken(), st.nextToken());
            }else{
                queue.dequeue(st.nextToken());
            }
        }
    }
}
