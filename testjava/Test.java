import java.lang.invoke.*;

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
    }

    public static int fab(int n) {
        if (n <= 1) {
            return n;
        }
        return fab(n - 1) + fab(n - 2);
    }
}