import java.lang.invoke.*;
import java.util.Arrays;

public class Test {
    private int a;
    public static long b;
    public static Test ttt;

    static {
        ttt = null;
    }

    public static void main(String[] args) throws Exception {
        Test test = new Test();
        int sum = 0;
        for (int i = 1 ; i <= 100; i ++) {
            sum += i;
        }
        System.out.println(sum);
        test.a = 1;
        test.b = 2;
        System.out.println(test.a);
        System.out.println(test.b);
        System.out.println(test instanceof Object);

//        MethodType mt = MethodType.methodType(Test.class);
//        MethodHandles.lookup().findConstructor(Test.class, mt);
        System.out.println(fab(30));
        int[] arr = new int[]{3, 2, 4, 10, 2, 9, 6};
        sort(arr);
        printArray(arr);
    }

    public static int fab(int n) {
        if (n <= 1) {
            return n;
        }
        return fab(n - 1) + fab(n - 2);
    }

    public static void sort(int[] intArray) {
        for (int i = 0; i < intArray.length - 1; i++) {
            for (int j = i + 1; j < intArray.length - i - 1; j++) {
                if (intArray[j] > intArray[j+1]) {
                    int tmp = intArray[j];
                    intArray[j] = intArray[j+1];
                    intArray[j+1] = tmp;
                }
            }
        }
    }

    public static void printArray(int[] intArray) {
        for (int a : intArray) {
            System.out.println(a);
        }
    }
}