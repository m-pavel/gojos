package cmd;

import java.io.Serializable;

public interface GojoTest<T extends Serializable> {
    T getnerateTestObject();
}
