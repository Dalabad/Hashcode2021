package com.google.hashcode.models;

import java.util.HashMap;
import java.util.Map;

/**
 * Created by neotreat on 11/02/16.
 */
public class Warehouse extends Location {

    /**
     *
     */
    private int id;

    /**
     *
     */
    private HashMap<ProductType, Integer> stock = new HashMap<ProductType, Integer>();

    /**
     *
     * @param id
     */
    public Warehouse(int id) {
        this.id = id;
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
     * Set the payload, which will replace the previous one completely!
     */
    public void setStock(HashMap<ProductType, Integer> stock) {

        this.stock = stock;
    }

    /**
     *
     * @param stock
     */
    public void addStock(HashMap<ProductType, Integer> stock) {

        // Add all products per type
        for (Map.Entry<ProductType, Integer> e : stock.entrySet()) {

            ProductType product = e.getKey();
            Integer amount = e.getValue();

            this.addProductToStock(product, amount);
        }
    }

    /**
     *
     * @param product
     */
    public void addProductToStock(ProductType product, Integer amount) {

        // Check if the product type is already present in the payload
        if(this.stock.containsKey(product)) {
            Integer currentAmount = this.stock.get(product);
            this.stock.put(product, new Integer(currentAmount.intValue() + amount.intValue()));
        }
        else {
            this.stock.put(product, amount);
        }
    }

    /**
     * Check if a product is in stock.
     */
    public boolean isProductInStock(ProductType product) {
        return PayloadHelper.isProductInMap(product, this.stock);
    }

    /**
     * Take a product from the stock.
     */
    public void retrieveProductFromStock(ProductType product, Integer amount) throws Exception {

        // Check if the item is in stock
        if (!isProductInStock(product)) {
            throw new Exception("Item out of stock!");
        }

        // Current number of products
        Integer currentAmount = this.stock.get(product);

        // Check if there are that many items available
        if(currentAmount.intValue() < amount.intValue()) {
            throw new Exception("Not that many items available");
        }

        // If there is nothing left remove the product completely
        Integer newAmount = new Integer(currentAmount.intValue() - amount.intValue());

        if (newAmount.intValue() <= 0) {
            this.stock.remove(product);
        }
        else {
            this.stock.put(product, newAmount);
        }
    }
}
