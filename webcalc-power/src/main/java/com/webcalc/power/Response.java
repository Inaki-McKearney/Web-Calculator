package com.webcalc.power;

public class Response {
    private final Boolean error;
    private final String string;
    private final Long answer;

    public Response(Boolean error, String string, Long answer) {
        this.error = error;
        this.string = string;
        this.answer = answer;
    }

    public Boolean getError() {
        return this.error;
    }

    public String getString() {
        return this.string;
    }

    public Long getAnswer() {
        return this.answer;
    }

    @Override
    public String toString() {
        return "{" +
            " error='" + getError() + "'" +
            ", string='" + getString() + "'" +
            ", answer='" + getAnswer() + "'" +
            "}";
    }
}
