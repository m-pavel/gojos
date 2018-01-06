package test;

import cmd.GojoTest;
import model.t1.Child1;

public class Test1 implements GojoTest<Child1> {
    @Override
    public Child1 getnerateTestObject() {
        Child1 c1 = new Child1();
        c1.setChildString("stringVaaalue");
        c1.setIntfld(325);
        return c1;
    }
}