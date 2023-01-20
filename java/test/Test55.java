package test;

import java.util.Date;
import java.util.HashMap;

import cmd.GojoTest;
import model.t2.Tkn;

public class Test55 implements GojoTest<Tkn> {

    @Override
    public Tkn getnerateTestObject() {
        Tkn tkn = new Tkn();
        tkn.setValidUntil(new Date());
        tkn.setGeneratedDateTime(new Date(951861600000L)); // Tue Feb 29 2000 22:00:00 GMT+0000
        tkn.setMetadata(new HashMap<>());
        tkn.getMetadata().put("foo", "bar");
        tkn.setUserId("testUser");
        return tkn;
    }
}