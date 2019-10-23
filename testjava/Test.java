public class Test {
    private int a;
    public static long b;

    public static void main(String[] args) {
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
    }
}