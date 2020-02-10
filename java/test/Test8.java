package test;

import cmd.GojoTest;
import java.io.Serializable;

public class Test8 implements GojoTest<Test8.C> {
    public static class A implements Serializable {
        String value;
    }
    public static class B implements Serializable {
        A a;
    }
    public static class C implements Serializable {
        B b;
        A a;
    }
    @Override
    public C getnerateTestObject() {
        A a = new A();
        a.value = "just string";
        B b = new B();
        C c = new C();
        b.a = a;
        c.a = a;
        c.b = b;
        return c;
    }
}