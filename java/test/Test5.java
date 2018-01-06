package test;

import cmd.GojoTest;
import model.t1.Child1;

public class Test5 implements GojoTest<Child1> {
    @Override
    public Child1 getnerateTestObject() {
        Child1 c1 = new Child1();
        c1.setChildString(null);
        c1.setIntfld(325);
        return c1;
    }
}