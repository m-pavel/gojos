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
        tkn.setGeneratedDateTime(new Date(951861600000L));
        tkn.setMetadata(new HashMap<>());
        tkn.getMetadata().put("foo", "bar");
        return tkn;
    }
}