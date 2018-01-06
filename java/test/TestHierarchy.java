package test;

import java.io.Serializable;

import cmd.GojoTest;

public class TestHierarchy implements GojoTest<TestHierarchy.C2>{
    public static class P1 implements Serializable{
        protected int p1int;
    }
    public static class C1 extends P1 {
        private String c1str;
    }
    
    public static class C2 extends C1 {
        private String c2str ;
    }

    @Override
    public C2 getnerateTestObject() {
        C2 c2 = new C2();
        c2.c2str = "C1";
        c2.p1int = 55;
        return c2;
    }
}
