package model.t2;

import java.io.Serializable;
import java.util.Date;
import java.util.HashMap;

public class Tkn implements Serializable{
    private String userId;
    private Date generatedDateTime;
    private Date validUntil;
    private HashMap<String, String> metadata;
    public String getUserId() {
        return userId;
    }
    public void setUserId(String userId) {
        this.userId = userId;
    }
    public Date getGeneratedDateTime() {
        return generatedDateTime;
    }
    public void setGeneratedDateTime(Date generatedDateTime) {
        this.generatedDateTime = generatedDateTime;
    }
    public Date getValidUntil() {
        return validUntil;
    }
    public void setValidUntil(Date validUntil) {
        this.validUntil = validUntil;
    }
    public HashMap<String, String> getMetadata() {
        return metadata;
    }
    public void setMetadata(HashMap<String, String> metadata) {
        this.metadata = metadata;
    }
}
