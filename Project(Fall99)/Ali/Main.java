import java.util.*;

public class Main {
    static Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) {
        int commandsNum = scanner.nextInt();
        scanner.nextLine();
        String init = scanner.nextLine();
        Calculator calculator = new Calculator(init);
        for (int i = 0; i < commandsNum; i++) {
            String command = scanner.nextLine();
            calculator.doCommand(command);
        }

    }

    static class Calculator {
        private Node head, tail, cursor;
        private Node temp;
        private static final int MAX_VAL = (int) Math.pow(10, 9) + 7;
        private HashMap<String, Long> resHashMap;

        public Calculator(String initialStr) {
            temp = new Node("#");
            for (String str : initialStr.split(" ")) {
                plus(str);
            }
            resetCursor();
            resHashMap = new HashMap<>();
        }

        private void insertAtCursor(String data) {
            Node newNode = new Node(data);
            if (head == null) {
                head = newNode;
                tail = head;
                cursor = head;
                temp.next = head;
                head.prev = temp;
                return;
            }
            if (tail != cursor) {
                newNode.next = cursor.next;
                cursor.next.prev = newNode;
                if (cursor == temp) {
                    head = newNode;
                }
            } else {
                tail = newNode;
            }
            cursor.next = newNode;
            newNode.prev = cursor;
            moveCursorForward();

        }

        private void delete() {
            if (cursor == temp)
                return;
            if (head == tail) {
                head = tail = null;
            } else {
                cursor.prev.next = cursor.next;
                if (cursor == tail) {
                    tail = cursor.prev;
                } else {
                    cursor.next.prev = cursor.prev;
                    if (cursor == head) {
                        head = cursor.next;
                        cursor = temp;
                        return;
                    }
                }
            }
            moveCursorBackward();

        }

        private void moveCursorForward() {
            if (cursor != tail)
                cursor = cursor.next;
        }

        private void moveCursorBackward() {
            if (cursor != temp)
                cursor = cursor.prev;
        }

        private void resetCursor() {
            cursor = temp;
        }

        private String getInputWithCursor() {
            if (head == null)
                return "|";
            Node n = head;
            String str = "";
            if (cursor == temp) {
                str += "|";
            }
            while (n != null) {
                str += n.data;
                if (n == cursor) {
                    str += "|";
                }
                n = n.next;
            }
            return str;
        }

        private String getInputWithSpace() {
            if (head == null)
                return "";
            Node n = head;
            String str = "";
            String last = n.data;
            while (n != null) {
                String data = n.data;
                if (data.equals("(") && last.equals(")"))
                    str += "* ( ";
                else if (isNumber(data) && last.equals(")"))
                    str += (" * " + data);
                else if (isNumber(last) && data.equals("("))
                    str += " * ( ";
                else if (data.equals("(") || data.equals(")")) {
                    str += " " + data + " ";
                } else if ((isNumber(data) || isOperator(last)) && (isNumber(data) || isOperator(data))) {
                    str += data;
                } else str += " " + data + " ";
                last = data;
                n = n.next;
            }
            return str;
        }

        private boolean isOperator(String str) {
            return str.equals("+") || str.equals("*") || str.equals("-");
        }


        private int getPrecedence(String str) {
            if (str.equals("*"))
                return 3;
            if (str.equals("+") || str.equals("-"))
                return 2;
            if (str.equals("("))
                return 1;
            return 0;
        }

        private String infixToPostfix(String infixexpr) {
            Stack<String> opStack = new Stack();
            LinkedList<String> postfixList = new LinkedList<>();
            String[] tokenList = infixexpr.split("\\s+");
            for (String token : tokenList) {
                if (isNumber(token)) {
                    if (token.length() > 8)
                        postfixList.add(mod(token, MAX_VAL));
                    else
                        postfixList.add(token);
                } else if (token.equals("(")) {
                    opStack.push(token);
                } else if (token.equals(")")) {
                    String topToken = opStack.pop();
                    while (!topToken.equals("(")) {
                        postfixList.add(topToken);
                        topToken = opStack.pop();
                    }
                } else {
                    while (!opStack.isEmpty() && (getPrecedence(opStack.peek()) >= getPrecedence(token))) {
                        postfixList.add(opStack.pop());
                    }
                    opStack.push(token);
                }
            }

            while (!opStack.isEmpty()) {
                postfixList.add(opStack.pop());
            }
            return arrayToString(postfixList);
        }

        private String postfixEval(String postfixExpr) {
            Stack<String> operandStack = new Stack();
            String[] tokenList = postfixExpr.split("\\s+");
            for (String token : tokenList) {
                if (isNumber(token)) {
                    operandStack.push(token);
                } else {
                    String operand2 = operandStack.pop();
                    String operand1 = operandStack.pop();
                    String result = doMath(token, operand1, operand2);
                    operandStack.push(result);
                }
            }
            return operandStack.pop();
        }

        private String doMath(String op, String op1, String op2) {

            switch (op) {
                case "*":
                    return Long.toString(((Long.parseLong(op1) % MAX_VAL) * (Long.parseLong(op2) % MAX_VAL)) % MAX_VAL);
                case "+":
                    return Long.toString(((Long.parseLong(op1) % MAX_VAL) + (Long.parseLong(op2) % MAX_VAL)) % MAX_VAL);
                case "-":
                    return Long.toString(((Long.parseLong(op1) % MAX_VAL) - (Long.parseLong(op2) % MAX_VAL)) % MAX_VAL);
                default:
                    return null;
            }
        }

        private String arrayToString(LinkedList<String> arr) {
            String res = "";
            for (String s : arr) {
                res += (s + " ");
            }
            return res;
        }

        private boolean isNumber(String str) {
            try {
                Long.parseLong(str);
            } catch (Exception e) {
                return false;
            }
            return true;
        }

        public void doCommand(String command) {
            String[] strings = command.split(" ");
            switch (strings[0]) {
                case "+":
                    plus(strings[1]);
                    break;
                case "-":
                    minus();
                    break;
                case ">":
                    greaterThan();
                    break;
                case "<":
                    lessThan();
                    break;
                case "?":
                    questionMark();
                    break;
                case "!":
                    exclamationMark();
                    break;
                default:
                    return;
            }
        }

        private void plus(String newNode) {
            for (char c : newNode.toCharArray()) {
                insertAtCursor(c + "");
            }
        }

        private void minus() {
            delete();
        }

        private void greaterThan() {
            moveCursorForward();
        }

        private void lessThan() {
            moveCursorBackward();
        }

        private void questionMark() {
            System.out.println(getInputWithCursor());
        }

        private void exclamationMark() {
            String input = getInputWithSpace();
            Long res_from_hash = checkMap(input);
            if (res_from_hash != null) {
                System.out.println(res_from_hash);
                return;
            }
            long res = Long.parseLong(postfixEval(infixToPostfix(input)));
            while (res < 0) {
                res += MAX_VAL;
            }
            resHashMap.put(input, res);
            System.out.println(res);
        }

        private Long checkMap(String key) {
            return resHashMap.get(key);
        }

        private String mod(String num, long a) {
            long res = 0;
            char[] charArr = num.toCharArray();
            for (int i = 0; i < num.length(); i++)
                res = (res * 10 + (int) charArr[i] - '0') % a;
            return Long.toString(res);
        }

        private class Node {
            String data;
            Node next;
            Node prev;

            public Node(String data) {
                this.data = data;
            }
        }
    }
}