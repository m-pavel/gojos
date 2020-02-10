package test;

import cmd.GojoTest;
import java.io.Serializable;

public class Test9 implements GojoTest<Test9.C> {
    public static class C implements Serializable {
        java.util.Date d1;
        java.util.Date d2;
        String s;
    }
    @Override
    public C getnerateTestObject() {
        C c = new C();
        c.d1 = new java.util.Date();
        c.d2 = c.d1;
        c.s = "error";
        return c;
    }
}