package com.webcalc.power;

public class Error {
    private final Boolean error;
    private final String message;

    public Error(Boolean error, String message) {
        this.error = error;
        this.message = message;
    }

    public Boolean getError() {
        return this.error;
    }

    public String getMessage() {
        return this.message;
    }

    @Override
    public String toString() {
        return "{" +
            " error='" + getError() + "'" +
            ", message='" + getMessage() + "'" +
            "}";
    }
}