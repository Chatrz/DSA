import java.util.*;

//  a noobie answer for https://www.codechef.com/LRNDSA02/problems/INPSTFIX
/*
	takes number of test cases as first line
	and in second line takes the length of the input expression
	in third line it takes the infix expression and convert it to postfix form

*/
public class InfixToPostFix{

	public static void main(String[] args){
		Scanner input =new Scanner(System.in);
		String validPattern="[A-Z*()+^/-]*";
		int caseNum=Integer.parseInt(input.nextLine());
		String[] results=new String[caseNum];
		for(int i=0;i<caseNum;i++){
			int length=Integer.parseInt(input.nextLine());
			String line=input.nextLine();
			if(line.matches(validPattern)){
				char[] character=line.toCharArray();
				results[i]=convert(character,length);
			}else{
				System.out.println("You have entered unvalid operator operands");
			}
		}
		for(String result:results){
			System.out.println(result);
		}
	}
	static boolean isPrecedent(char nextOp,char stackOp){
		if(nextOp==')'||nextOp=='n')return true;
		if(nextOp=='*'&& stackOp=='/')return false;
		if(nextOp=='+'&& stackOp=='-')return false;
		String operators="^*/+-";
     	return operators.indexOf(nextOp)>operators.indexOf(stackOp);
		
	}
	static String convert(char[] elements,int size){
		String operand="^*/+-";
		Stack<String> symbols=new Stack<String>();
		Stack<Character> operators=new Stack<Character>();
		for(int i=0;i<size;i++){
			char element=elements[i];
			if(operand.indexOf(element)>=0|| element=='('){
				operators.push(element);
			}else if(element==')'){
				if(operators.peek()=='(')operators.pop();
				else{
					while(operators.peek()!='('){
						String top=symbols.pop();
						String secondTop=symbols.pop();
						symbols.push(secondTop+top+operators.pop());
					}
					// in this stage the top of operators stack would be '('
					operators.pop();
				}
			}else{
				char nextEl=(i+1)<size-1?elements[i+1]:'n';
				// when symbols stack is empty operators would be empty too and peeking empty stack throws an error (ig)
				if(symbols.empty())symbols.push(Character.toString(element));
				else if(operators.peek()=='('||!isPrecedent(nextEl,operators.peek())){
					symbols.push(Character.toString(element));

				}else {
					String newStr=symbols.pop()+Character.toString(element)+operators.pop();
					symbols.push(newStr);
				}
			}
		}
		while(operators.size()>0){
			String top=symbols.pop();
			String secondTop=symbols.pop();
			symbols.push(secondTop+top+operators.pop());
		}
		return symbols.pop();
	}
}