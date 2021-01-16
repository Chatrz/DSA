import java.util.Scanner;
import java.util.Vector;

public class Solution{

    static boolean is_found(Vector<String> vec, String temp)
    {
        for (String item : vec)
        {
            if (item.contains(temp))
                return true;
        }
        return false;
    }

    static boolean is_lost(Vector<String> vec, String temp)
    {
        for (String item : vec)
        {
            if (temp.contains(item))
                return true;
        }
        return false;
    }

    public static void main(String[] args) {
        Vector<String> found = new Vector<>();
        Vector<String> lost = new Vector<>();
	    Scanner scanner = new Scanner(System.in);
	    String main_string = scanner.nextLine();
	    int number = scanner.nextInt();
	    long count = 0;
	    for (int i = 0; i < number; i++)
        {
            String temp = scanner.next();
            if (is_found(found, temp))
            {
                count++;
            } else if (is_lost(lost, temp))
            {
            } else {
                boolean result = (main_string.indexOf(temp) != -1);
                if (result)
                {
                    found.add(temp);
                    count++;
                } else {
                    lost.add(temp);
                }
            }
        }
	    System.out.print(count);
    }
}
