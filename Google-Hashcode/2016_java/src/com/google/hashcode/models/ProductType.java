package com.google.hashcode.models;

/**
 * Created by neotreat on 11/02/16.
 */
public class ProductType implements Comparable<ProductType> {

    /**
     *
     */
    private int id;

    /**
     *
     */
    private int weight;

    /**
     *
     */
    public ProductType(int id, int weight) {
        this.id = id;
        this.weight = weight;
    }



    /**
     *
     * @return
     */
    public int getId() {
        return id;
    }

    /**
     *
     * @param id
     */
    public void setId(int id) {
        this.id = id;
    }



    /**
     *
     * @return
     */
    public int getWeight() {
        return weight;
    }

    /**
     *
     * @param weight
     */
    public void setWeight(int weight) {
        this.weight = weight;
    }

    /**
     *
     * @param o
     * @return
     */
    @Override
    public int compareTo(ProductType o) {

        if(this.weight < o.getWeight())
        {
            return -1;
        }
        else if (this.weight == o.getWeight()) {
            return 0;
        }
        else {
            return 1;
        }
    }
}
