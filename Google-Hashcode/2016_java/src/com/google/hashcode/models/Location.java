package com.google.hashcode.models;

/**
 * Created by neotreat on 11/02/16.
 */
public abstract class Location {

    public abstract int getId();

    /**
     *
     */
    public int row;

    /**
     *
     */
    public int column;

    /**
     *
     */
    public Location () {
        this.row = 0;
        this.column = 0;
    }

    /**
     *
     * @param row
     * @param column
     * @return
     */
    public Location (int row, int column) {
        this.row = row;
        this.column = column;
    }
}
