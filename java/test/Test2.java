package test;

import java.util.Date;

import cmd.GojoTest;

public class Test2 implements GojoTest<Date> {
    @Override
    public Date getnerateTestObject() {
        Date d = new Date();
        System.out.println(d);
        System.out.println(d.getTime());
        return d;
    }
}