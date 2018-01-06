package model.t1;

import java.io.Serializable;

public class Parent implements Serializable{
    private static final long serialVersionUID = -7748096251001970577L;
    private int intfld;
    public int getIntfld() {
        return intfld;
    }
    public void setIntfld(int intfld) {
        this.intfld = intfld;
    }
}
