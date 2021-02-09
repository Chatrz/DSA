import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int stagesNumber, commandsNumber;
        stagesNumber = scanner.nextInt();
        commandsNumber = scanner.nextInt();
        Stage[] stages = new Stage[stagesNumber];
        for (int i = 0; i < stagesNumber; i++) {
            stages[i] = new Stage(i + 1);
        }
        for (int i = 0; i < commandsNumber; i++) {
            int src = scanner.nextInt(), des = scanner.nextInt();
            stages[src - 1].moveBoxesTo(stages[des - 1]);
        }

        stages[scanner.nextInt() - 1].print();
    }
}

class Stage {
    int size;
    Box top, bottom;

    public Stage(int number) {
        top = new Box(number);
        bottom = top;
        size = 1;
    }

    public void print() {
        System.out.print(size + " ");
        Box bottomNode = bottom;
        for (int i = 0; i < size; ++i, bottomNode = bottomNode.next) {
            System.out.print(bottomNode.value + " ");
        }
    }

    public void moveBoxesTo(Stage des) {
        if (size != 0) {
            if (des.bottom == null) {
                des.bottom = bottom;
            } else {
                des.top.next = bottom;
            }
            des.top = top;
            top = bottom = null;
            des.size += size;
            this.size = 0;
        }
    }
}

class Box {
    int value;
    Box next;

    public Box(int value) {
        this.value = value;
        next = null;
    }
}
